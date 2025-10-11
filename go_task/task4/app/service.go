package app

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repo
	Sec  string
}

func (svc *Service) Init(sec string) error {
	svc.Sec = sec
	dsn := "root:root@tcp(127.0.0.1:3306)/task4?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// for test
	// err = db.Migrator().DropTable(&User{}, &Post{}, &Comment{})

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})

	if err != nil {
		return err
	}
	svc.repo = &Repo{db: db}
	return nil
}

func (svc *Service) Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := svc.repo.AddUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (svc *Service) Login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	var storedUser User
	if storedUser, err = svc.repo.FindUserByName(user.Username); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(svc.Sec))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// 剩下的逻辑...
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// 获取所有的文章信息
func (svc *Service) GetPosts(c *gin.Context) {
	repo := svc.repo
	posts, err := repo.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// 获取某篇文章
func (svc *Service) GetPostById(c *gin.Context) {
	repo := svc.repo
	id := c.Param("id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}

	post, err := repo.GetPostById(uint(post_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// 添加文章
func (svc *Service) AddPost(c *gin.Context) {
	repo := svc.repo
	userid := c.GetUint("userid")
	//username := c.GetString("username")

	var post Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post.UserId = userid
	fmt.Println(post)
	if err := repo.AddPost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

func (svc *Service) UpdatePost(c *gin.Context) {
	repo := svc.repo
	userid := c.GetUint("userid")
	post_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	post, err := repo.GetPostById(uint(post_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "找不到文章"})
		return
	}

	if post.UserId != userid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无权限修改"})
		return
	}

	if err := c.ShouldBind(post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	post.UserId = userid
	if err := repo.UpdatePost(post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// 删除文章
func (svc *Service) DeletePost(c *gin.Context) {
	repo := svc.repo
	userId := c.GetUint("userid")
	postId, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param"})
		return
	}

	if err := repo.DeletePost(userId, uint(postId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

// 添加新评论
func (svc *Service) AddComment(c *gin.Context) {
	uid := c.GetUint("userid")
	repo := svc.repo

	var comment Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.UserId = uid
	if err := repo.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}

// 获取某篇文章的所有评论
func (svc *Service) GetComments(c *gin.Context) {
	repo := svc.repo
	post_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章评论失败"})
		return
	}

	if comments, err := repo.GetCommentsByPostId(uint(post_id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章评论失败"})
		return
	} else {
		c.JSON(http.StatusOK, comments)
	}
}
