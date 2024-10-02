package domain

type Post struct {
	Id       int
	AuthorId int
	TItle    string
	Content  string
	Deleted  bool

	Timestamp
}

type Comment struct {
	Id         int
	PostId     int
	AuthorName string
	Content    string

	Timestamp
}
