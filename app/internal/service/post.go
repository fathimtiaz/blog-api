package service

import (
	"blog-api/config"
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
	"time"
)

type PostRepo interface {
	SavePost(context.Context, *domain.Post) error
	GetPosts(context.Context, repository.PostQuery) ([]domain.Post, error)
	GetPost(context.Context, repository.PostQuery) (domain.Post, error)
	UpdatePost(ctx context.Context, post domain.Post, authorId int) error
	DeletePost(ctx context.Context, id, authorId int, updatedAt time.Time) error

	SavePostComment(context.Context, *domain.Comment) error
	GetPostComments(context.Context, repository.CommentQuery) ([]domain.Comment, error)

	GetUser(context.Context, repository.UserQuery) (domain.User, error)
}

type PostService struct {
	cfg      config.Config
	postRepo PostRepo
}

func NewPostService(cfg config.Config, postRepo PostRepo) *PostService {
	return &PostService{cfg, postRepo}
}

func (s *PostService) CreatePost(ctx context.Context, title, content string) (post domain.Post, err error) {
	var userEmail string
	var user domain.User

	if userEmail, err = domain.AuthdUserEmail(ctx); err != nil {
		return
	}

	if user, err = s.postRepo.GetUser(ctx, repository.UserQuery{Email: userEmail}); err != nil {
		return
	}

	post = domain.NewPost(user.Id, title, content)

	if err = s.postRepo.SavePost(ctx, &post); err != nil {
		return
	}

	return
}

func (s *PostService) ListPost(ctx context.Context, page, limit int) (post []domain.Post, err error) {

	return s.postRepo.GetPosts(ctx, repository.PostQuery{
		Pagination: repository.Pagination{Page: page, Limit: limit},
	})
}

func (s *PostService) GetPost(ctx context.Context, id int) (post domain.Post, err error) {
	return s.postRepo.GetPost(ctx, repository.PostQuery{Id: id})
}

func (s *PostService) UpdatePost(ctx context.Context, post domain.Post) (err error) {
	var userEmail string
	var user domain.User

	if userEmail, err = domain.AuthdUserEmail(ctx); err != nil {
		return
	}

	if user, err = s.postRepo.GetUser(ctx, repository.UserQuery{Email: userEmail}); err != nil {
		return
	}

	post.UpdatedAt = time.Now()

	return s.postRepo.UpdatePost(ctx, post, user.Id)
}

func (s *PostService) DeletePost(ctx context.Context, id int) (err error) {
	var userEmail string
	var user domain.User

	if userEmail, err = domain.AuthdUserEmail(ctx); err != nil {
		return
	}

	if user, err = s.postRepo.GetUser(ctx, repository.UserQuery{Email: userEmail}); err != nil {
		return
	}

	return s.postRepo.DeletePost(ctx, id, user.Id, time.Now())
}

func (s *PostService) AddComment(ctx context.Context, postId int, content string) (comment domain.Comment, err error) {
	var userEmail string
	var user domain.User

	if userEmail, err = domain.AuthdUserEmail(ctx); err != nil {
		return
	}

	if user, err = s.postRepo.GetUser(ctx, repository.UserQuery{Email: userEmail}); err != nil {
		return
	}

	comment = domain.NewPostComment(postId, user.Name, content)

	if err = s.postRepo.SavePostComment(ctx, &comment); err != nil {
		return
	}

	return
}

func (s *PostService) GetComments(ctx context.Context, postId, page, limit int) (comments []domain.Comment, err error) {
	return s.postRepo.GetPostComments(ctx, repository.CommentQuery{
		PostId:     postId,
		Pagination: repository.Pagination{Page: page, Limit: limit},
	})
}
