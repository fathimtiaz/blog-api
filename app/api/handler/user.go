package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/service"
	"log"
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
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, setResponse(ErrBadRequest, nil))
		return
	}

	if user, err = h.userService.Register(c, req.Email, req.Name, req.Password); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, setResponse(ErrServer, nil))
	}

	c.JSON(http.StatusOK, setResponse(nil, user))
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (h *userHandler) Login(c *gin.Context) {
	var req loginRequest
	var err error
	var resp loginResponse

	if err = c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, setResponse(ErrBadRequest, nil))
		return
	}

	if resp.Token, err = h.userService.Login(c, req.Email, req.Password); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, setResponse(ErrServer, nil))
		return
	}

	c.JSON(http.StatusOK, setResponse(nil, resp))
}
