package service

import (
	"blog-api/config"
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
)

type UserRepo interface {
	InsertUser(context.Context, *domain.User) error
	GetUser(context.Context, repository.UserQuery) (domain.User, error)
}

type userService struct {
	cfg      config.Config
	userRepo UserRepo
}

func NewUserService(cfg config.Config, userRepo UserRepo) *userService {
	return &userService{cfg, userRepo}
}

func (s *userService) Register(ctx context.Context, name, email, password string) (user domain.User, err error) {
	if user, err = domain.NewUser(name, email, password); err != nil {
		return
	}

	if err = s.userRepo.InsertUser(ctx, &user); err != nil {
		return
	}

	return
}

func (s *userService) Login(ctx context.Context, email, password string) (token string, err error) {
	var user domain.User

	if user, err = s.userRepo.GetUser(ctx, repository.UserQuery{Email: email}); err != nil {
		return
	}

	if err = user.ComparePassword(password); err != nil {
		return
	}

	return user.GenerateToken(s.cfg.JWT.Secret.String(), s.cfg.JWT.HourExpired.Int())
}
