package domain

import "time"

type Post struct {
	Id       int
	AuthorId int
	TItle    string
	Content  string
	Deleted  bool

	Timestamp
}

func NewPost(authorId int, title, content string) Post {
	return Post{
		AuthorId:  authorId,
		TItle:     title,
		Content:   content,
		Timestamp: Timestamp{CreatedAt: time.Now()},
	}
}

type Comment struct {
	Id         int
	PostId     int
	AuthorName string
	Content    string

	Timestamp
}

func NewPostComment(postId int, authorName, content string) Comment {
	return Comment{
		PostId:     postId,
		AuthorName: authorName,
		Content:    content,
		Timestamp:  Timestamp{CreatedAt: time.Now()},
	}
}
