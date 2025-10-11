package app

import (
	"errors"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (repo *Repo) AddUser(user *User) error {
	return repo.db.Create(user).Error
}

func (repo *Repo) FindUserByName(name string) (user User, err error) {
	err = repo.db.Where("username = ?", name).First(&user).Error
	return user, err
}

func (repo *Repo) GetPosts() ([]Post, error) {
	var posts []Post
	err := repo.db.Find(&posts).Error
	return posts, err
}

func (repo *Repo) GetPostById(id uint) (*Post, error) {
	var post Post
	err := repo.db.First(&post, id).Error
	return &post, err
}

func (repo *Repo) AddPost(post *Post) error {
	return repo.db.Create(post).Error
}

func (repo *Repo) UpdatePost(post *Post) error {
	return repo.db.Save(post).Error
}

func (repo *Repo) DeletePost(userid uint, postid uint) error {
	var post Post
	err := repo.db.First(&post, postid).Error
	if err != nil {
		return err
	}
	if post.ID != userid {
		return errors.New("无权限")
	}
	return repo.db.Delete(&post).Error
}

func (repo *Repo) AddComment(c *Comment) error {
	return repo.db.Create(c).Error
}

func (repo *Repo) GetCommentsByPostId(id uint) ([]Comment, error) {
	var comments []Comment
	err := repo.db.Preload("Post").Where("post_id", id).Find(&comments).Error
	return comments, err
}
