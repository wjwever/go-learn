package app

import (
	"task4/middleware"

	"github.com/gin-gonic/gin"
)

type App struct {
	SecretKey string
	service   Service
}

// database operation
func (app *App) Init() error {
	return app.service.Init(app.SecretKey)
}

func (app *App) Start() {
	router := gin.Default()
	router.POST("/register", app.service.Register)
	router.POST("/login", app.service.Login)

	{
		apis := router.Group("/api")
		apis.Use(middleware.JWTAuth(app.SecretKey))
		// 获取所有文章信息
		apis.GET("/post", app.service.GetPosts)
		// 获取某篇文章信息
		apis.GET("/post/:id", app.service.GetPostById)
		// 添加文章
		apis.POST("/post", app.service.AddPost)
		// 删除文章
		apis.DELETE("/post/:id", app.service.DeletePost)
		// 跟新文章
		apis.PUT("/post/:id", app.service.UpdatePost)

		// 添加评论
		apis.GET("/post/:id/comment", app.service.GetComments)
		apis.POST("/comment", app.service.AddComment)
	}

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
