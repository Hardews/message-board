package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
)

func Post(username, userPost string) error {
	sqlStr := "insert into userPost (username,userPost) values (?,?)"
	_, err := dB.Exec(sqlStr, username, userPost)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(PostID int, userPost string) error {
	sqlStr := "update userPost set userPost = ? where id=?"
	userPost = userPost + "(已被删除)"
	_, err := dB.Exec(sqlStr, userPost, PostID)
	if err != nil {
		return err
	}
	return nil
}

func ChangePost(newPost string, PostID int) error {
	sqlStr := "update userPost set userPost = ? where postID = ?"
	_, err := dB.Exec(sqlStr, newPost, PostID)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllByPostId(postName, userPost string) (model.Post, error) {
	var u model.Post
	sqlStr := "select * from userPost where username= ? and userPost= ?"
	err := dB.QueryRow(sqlStr, postName, userPost).Scan(&u.PostID, &u.Username, &u.Txt, &u.LikeNum, &u.PostTime)
	if err != nil {
		return u, err
	}
	return u, err
}

func GetCommentsSection(PostID int) (error, []model.Comment) {
	var comments []model.Comment

	sqlStr := "select id,comment,time,commentName,commentLikeNum from usercomment where PostID = ?"
	rows, err := dB.Query(sqlStr, PostID)
	if err != nil {
		return err, comments
	}

	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.CommentId, &comment.Txt, &comment.Time, &comment.Username, &comment.LikeNum)
		if err != nil {
			return err, comments
		}
		comments = append(comments, comment)
	}
	return err, comments
}
