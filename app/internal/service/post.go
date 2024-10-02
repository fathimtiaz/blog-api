package service

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
)

type PostRepo interface {
	SavePost(context.Context, *domain.Post) error
	GetPosts(context.Context, repository.PostQuery) ([]domain.Post, error)
	GetPost(context.Context, repository.PostQuery) (domain.Post, error)
	UpdatePost(context.Context, domain.Post) error
	DeletePost(ctx context.Context, id int) error
}
