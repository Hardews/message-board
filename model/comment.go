package model

type Comment struct {
	CommentId int
	Txt       string
	Username  string
	PostID    int
	Time      string
	LikeNum   int
}
