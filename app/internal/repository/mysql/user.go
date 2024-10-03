package mysql

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
	"database/sql"
)

func (db *sqlDB) SaveUser(ctx context.Context, user *domain.User) (err error) {
	_, err = db.ExecContext(ctx, `
		INSERT INTO user_ (name_, email, password_hash, created_at) VALUES (?,?,?,?)
	`, user.Name, user.Email, user.PasswordHash, user.CreatedAt)

	return
}

func (db *sqlDB) GetUser(ctx context.Context, query repository.UserQuery) (user domain.User, err error) {
	var createdAt, updatedAt sql.NullTime

	err = db.QueryRowContext(ctx, `
		SELECT id, name_, email, password_hash, created_at, updated_at
		FROM user_
		WHERE email = ?
	`, query.Email).Scan(
		&user.Id, &user.Name, &user.Email, &user.PasswordHash, &createdAt, &updatedAt,
	)

	user.CreatedAt = createdAt.Time
	user.CreatedAt = createdAt.Time

	return
}
