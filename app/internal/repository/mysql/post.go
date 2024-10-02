package mysql

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
)

func (db *sqlDB) SavePost(ctx context.Context, post *domain.Post) (err error) {
	return
}

func (db *sqlDB) GetPosts(ctx context.Context, query repository.PostQuery) (posts []domain.Post, err error) {
	return
}

func (db *sqlDB) GetPost(ctx context.Context, query repository.PostQuery) (post domain.Post, err error) {
	return
}

func (db *sqlDB) UpdatePost(ctx context.Context, post domain.Post) (err error) {
	return
}

func (db *sqlDB) DeletePost(ctx context.Context, id int) (err error) {
	return
}
