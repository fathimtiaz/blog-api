package handler

import "github.com/gin-gonic/gin"

type userHandler struct {
}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (h *userHandler) Register(c *gin.Context) {

}

func (h *userHandler) Login(c *gin.Context) {

}
