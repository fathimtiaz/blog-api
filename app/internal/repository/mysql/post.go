package mysql

import (
	"blog-api/internal/domain"
	"blog-api/internal/repository"
	"context"
	"database/sql"
	"time"
)

func (db *sqlDB) SavePost(ctx context.Context, post *domain.Post) (err error) {
	_, err = db.ExecContext(ctx, `
		INSERT INTO post_ (author_id, title, content, created_at) VALUES (?,?,?,?)
	`, post.AuthorId, post.TItle, post.Content, post.CreatedAt)

	return
}

func (db *sqlDB) GetPosts(ctx context.Context, query repository.PostQuery) (posts []domain.Post, err error) {
	var rows *sql.Rows

	if rows, err = db.QueryContext(ctx, `
		SELECT id, author_id, title, content, created_at, updated_at
		FROM post_
		WHERE deleted IS NOT TRUE
		LIMIT ? OFFSET ?
	`, query.Limit, query.Page); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var post domain.Post
		var createdAt, updatedAt sql.NullTime

		if err = rows.Scan(
			&post.Id,
			&post.AuthorId,
			&post.TItle,
			&post.Content,
			&createdAt,
			&updatedAt,
		); err != nil {
			return
		}

		post.CreatedAt = createdAt.Time
		post.UpdatedAt = updatedAt.Time

		posts = append(posts, post)
	}

	return
}

func (db *sqlDB) GetPost(ctx context.Context, query repository.PostQuery) (post domain.Post, err error) {
	var createdAt, updatedAt sql.NullTime

	err = db.QueryRowContext(ctx, `
		SELECT id, author_id, title, content, created_at, updated_at
		FROM post_
		WHERE id = ?
	`, query.Id).Scan(
		&post.Id, &post.AuthorId, &post.TItle, &post.Content, &createdAt, &updatedAt,
	)

	post.CreatedAt = createdAt.Time
	post.CreatedAt = createdAt.Time

	return
}

func (db *sqlDB) UpdatePost(ctx context.Context, post domain.Post, authorId int) (err error) {
	_, err = db.ExecContext(ctx, `
		UPDATE post_ SET title = ?, content = ?, updated_at = ? WHERE id = ? and author_id = ?
	`, post.TItle, post.Content, post.UpdatedAt, post.Id, authorId)

	return
}

func (db *sqlDB) DeletePost(ctx context.Context, id, authorId int, updatedAt time.Time) (err error) {
	_, err = db.ExecContext(ctx, `
		UPDATE post_ SET deleted = TRUE, updated_at = ? WHERE id = ? and author_id = ?
	`, updatedAt, id, authorId)

	return
}

func (db *sqlDB) SavePostComment(ctx context.Context, comment *domain.Comment) (err error) {
	_, err = db.ExecContext(ctx, `
		INSERT INTO post_comment_ (post_id, author_name, content, created_at) VALUES (?,?,?,?)
	`, comment.PostId, comment.AuthorName, comment.Content, comment.CreatedAt)

	return
}

func (db *sqlDB) GetPostComments(ctx context.Context, query repository.CommentQuery) (comments []domain.Comment, err error) {
	var rows *sql.Rows

	if rows, err = db.QueryContext(ctx, `
		SELECT id, post_id, author_name, content, created_at, updated_at
		FROM post_comment_
		WHERE post_id = ?
		LIMIT ? OFFSET ?
	`, query.PostId, query.Limit, query.Page); err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var comment domain.Comment
		var createdAt, updatedAt sql.NullTime

		if err = rows.Scan(
			&comment.Id,
			&comment.PostId,
			&comment.AuthorName,
			&comment.Content,
			&createdAt,
			&updatedAt,
		); err != nil {
			return
		}

		comment.CreatedAt = createdAt.Time
		comment.UpdatedAt = updatedAt.Time

		comments = append(comments, comment)
	}

	return
}
