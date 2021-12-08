package model

import "time"

type Post struct {
	PostID   int
	Txt      string
	Username string
	PostTime time.Time
}
