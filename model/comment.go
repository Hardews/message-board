package model

import "time"

type Comment struct {
	CommentId int
	Txt       string
	Username  string
	PostID    int
	Time      time.Time
}
