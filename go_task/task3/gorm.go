package main

import (
	"fmt"
	"log"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Posts    []Post `gorm:"foreignkey:UserRefer"`
	PostNum  int
}

type Post struct {
	gorm.Model
	Title        string `gorm:"not null"`
	Content      string `gorm:"not null"`
	UserRefer    int    `gorm:"not null"`
	CommentState string
	Comments     []Comment `gorm:"foreignkey:PostRefer"`
}

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null"`
	PostRefer int    `gorm:"not null"`
}

func checkErr(err error) {
	if err != nil {
		_, _, line, _ := runtime.Caller(1)
		log.Fatalf("error happens: %v \nline:%v", err, line)
	}
}

func (u *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	err = tx.First(&user, u.UserRefer).Error
	if err == nil {
		user.PostNum = user.PostNum + 1
		tx.Save(&user)
	}

	return err
}

func (u *Post) AfterDelete(tx *gorm.DB) (err error) {
	var user User
	err = tx.First(&user, u.UserRefer).Error
	if err == nil {
		user.PostNum = user.PostNum - 1
		tx.Save(&user)
	}
	return err
}

func (u *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_refer=?", u.PostRefer).Count(&count).Error
	if err == nil {
		state := "无评论"
		if count > 0 {
			state = fmt.Sprintf("评论数 %d", count)
		}
		err = tx.Model(&Post{}).Where("id=?", u.PostRefer).UpdateColumn("comment_state", state).Error
	}
	return err
}

func (u *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	err = tx.Model(&Comment{}).Where("post_refer=?", u.PostRefer).Count(&count).Error
	if err == nil {
		state := "无评论"
		if count > 0 {
			state = fmt.Sprintf("评论数 %d", count)
		}
		err = tx.Model(&Post{}).Where("id=?", u.PostRefer).UpdateColumn("comment_state", state).Error
	}
	return err
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	checkErr((err))

	// 删除表，保持运行结果一致性
	err = db.Migrator().DropTable(&User{}, &Post{}, &Comment{})
	checkErr(err)
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	checkErr(err)

	user1 := User{Username: "Tom", PostNum: 0}
	user2 := User{Username: "Jack", PostNum: 0}
	db.Create(&user1)
	db.Create(&user2)

	post1 := Post{Title: "title1", Content: "tmp filed1", UserRefer: 1, CommentState: "无评论"}
	post2 := Post{Title: "title2", Content: "tmp filed2", UserRefer: 1, CommentState: "无评论"}
	db.Create(&post1)
	db.Create(&post2)

	comment1 := Comment{Content: "comment1", PostRefer: 1}
	comment2 := Comment{Content: "comment2", PostRefer: 1}
	db.Create(&comment1)
	db.Create(&comment2)

	var users []User
	err = db.Preload("Posts").Preload("Posts.Comments").Find(&users, 1).Error
	checkErr(err)

	for _, user := range users {
		fmt.Printf("UserName:%v PostNum:%v\n", user.Username, user.PostNum)
		for _, post := range user.Posts {
			fmt.Printf("\t Post Title:%v CommentState:%v\n", post.Title, post.CommentState)
			for _, comment := range post.Comments {
				fmt.Printf("\t\t Comment:%v\n", comment.Content)
			}
		}
	}
}
