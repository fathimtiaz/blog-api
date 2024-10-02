package handler

import "github.com/gin-gonic/gin"

type postHandler struct {
}

func NewPostHandler() *postHandler {
	return &postHandler{}
}

func (h *postHandler) Create(c *gin.Context) {

}

func (h *postHandler) List(c *gin.Context) {

}

func (h *postHandler) Get(c *gin.Context) {

}

func (h *postHandler) Update(c *gin.Context) {

}

func (h *postHandler) Delete(c *gin.Context) {

}

func (h *postHandler) AddComment(c *gin.Context) {

}

func (h *postHandler) ListComments(c *gin.Context) {

}
