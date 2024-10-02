package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *userHandler {
	return &userHandler{userService}
}

type registerRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (h *userHandler) Register(c *gin.Context) {
	var req registerRequest
	var err error
	var user domain.User

	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if user, err = h.userService.Register(c, req.Email, req.Name, req.Password); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, user)
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *userHandler) Login(c *gin.Context) {
	var req loginRequest
	var err error
	var token string

	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if token, err = h.userService.Login(c, req.Email, req.Password); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
