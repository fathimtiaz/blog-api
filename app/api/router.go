package api

import (
	"blog-api/api/handler"
	"blog-api/api/middleware"
	"blog-api/config"
	"blog-api/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(
	cfg config.Config,
	userService *service.UserService,
	postService *service.PostService,
) (router *gin.Engine) {
	router = gin.Default()

	router.Use(middleware.SetAuthdUserCtx(cfg.JWT.Secret.String()))

	// user
	userHandler := handler.NewUserHandler(userService)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// post
	postHandler := handler.NewPostHandler(postService)
	router.POST("/post", postHandler.Create)
	router.GET("/post", postHandler.List)
	router.GET("/post/:id", postHandler.Get)
	router.PUT("/post/:id", postHandler.Update)
	router.DELETE("/post/:id", postHandler.Delete)

	// post comment
	router.POST("/post/:id/comment", postHandler.AddComment)
	router.GET("/post/:id/comment", postHandler.ListComments)

	return
}
