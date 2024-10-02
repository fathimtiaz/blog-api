package api

import (
	"blog-api/api/handler"
	"blog-api/internal/service"

	"github.com/gin-gonic/gin"
)

func Router(
	userService *service.UserService,
) (router *gin.Engine) {
	router = gin.Default()

	// user
	userHandler := handler.NewUserHandler(userService)
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// post
	postHandler := handler.NewPostHandler()
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
