package repository

type UserQuery struct {
	Email string
}

type PostQuery struct {
	Id int

	Pagination
}

type CommentQuery struct {
	PostId int

	Pagination
}
