package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/service"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *postHandler {
	return &postHandler{postService}
}

type createPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *postHandler) Create(c *gin.Context) {
	var req createPostRequest
	var err error
	var post domain.Post

	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if post, err = h.postService.CreatePost(c, req.Title, req.Content); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func parsePostQuery(query url.Values) (page, limit int) {
	var err error
	var pageQ = query.Get("page")
	var limitQ = query.Get("limit")

	if page, err = strconv.Atoi(pageQ); pageQ != "" && err != nil {
		log.Print("error parsing query limit", err.Error())

	}

	if limit, err = strconv.Atoi(limitQ); limitQ != "" && err != nil {
		log.Print("error parsing query page", err.Error())
	}

	return
}

func (h *postHandler) List(c *gin.Context) {
	var posts []domain.Post
	var page, limit = parsePostQuery(c.Request.URL.Query())
	var err error

	if posts, err = h.postService.ListPost(c, page, limit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *postHandler) Get(c *gin.Context) {
	var err error
	var postId int
	var post domain.Post

	if postId, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if post, err = h.postService.GetPost(c, postId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandler) Update(c *gin.Context) {
	var err error
	var post domain.Post

	if err = c.ShouldBindJSON(&post); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err = h.postService.UpdatePost(c, post); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandler) Delete(c *gin.Context) {
	var err error
	var postId int

	if postId, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = h.postService.DeletePost(c, postId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

type addCommentRequest struct {
	Content string `json:"content"`
}

func (h *postHandler) AddComment(c *gin.Context) {
	var req addCommentRequest
	var err error
	var postId int
	var comment domain.Comment

	if postId, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if comment, err = h.postService.AddComment(c, postId, req.Content); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *postHandler) ListComments(c *gin.Context) {
	var comments []domain.Comment
	var page, limit = parsePostQuery(c.Request.URL.Query())
	var postId int
	var err error

	if postId, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if comments, err = h.postService.GetComments(c, postId, page, limit); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, comments)
}
