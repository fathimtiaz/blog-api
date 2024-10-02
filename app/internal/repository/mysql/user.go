package mysql

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
)

func (db *sqlDB) SaveUser(ctx context.Context, user *domain.User) (err error) {
	return db.QueryRowContext(ctx, `
		INSERT INTO user_ (name_, email, password_hash, created_at) VALUES ($1, $2, $3, $4);
		SELECT LAST_INSERT_ID();
	`, user.Name, user.Email, user.PasswordHash, user.CreatedAt).Scan(user.Id)
}

func (db *sqlDB) GetUser(ctx context.Context, query repository.UserQuery) (user domain.User, err error) {
	err = db.QueryRowContext(ctx, `
		SELECT id, name_, email, password_hash, created_at, updated_at
		FROM user_
		WHERE email
	`, query.Email).Scan(
		&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt,
	)

	return
}
