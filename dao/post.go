package dao

import (
	"database/sql"
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

func DeletePost(PostID int) error {
	sqlStr := "delete userPost from userPost where id=?"
	_, err := dB.Exec(sqlStr, PostID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPost() (error, []model.Post, []string) {
	var users []model.Post
	var user model.Post
	var Time []string
	var time string
	sqlStr1 := "select username,userPost,time from userPost"
	rows, err := dB.Query(sqlStr1)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, users, Time
		}
		return err, users, Time
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Txt, &time)
		if err != nil {
			return err, users, Time
		}
		users = append(users, user)
		Time = append(Time, time)
	}
	return err, users, Time
}

func ChangePost(newPost string, PostID int) error {
	sqlStr := "update userPost set userPost = ? where postID = ?"
	_, err := dB.Exec(sqlStr, newPost, PostID)
	if err != nil {
		return err
	}
	return nil
}

func SelectByPostId(postName, userPost string) (int, error) {
	var u model.Post
	sqlStr := "select id from userPost where username= ? and userPost= ?"
	err := dB.QueryRow(sqlStr, postName, userPost).Scan(&u.PostID)
	if err != nil {
		return u.PostID, err
	}
	return u.PostID, err
}

func GetPost(username string) error {
	var u model.Post
	sqlStr := "select userPost from userPost where username=? "
	err := dB.QueryRow(sqlStr, username).Scan(&u.Txt)
	if err != nil {
		return err
	}
	return err
}

func SelectPostAndCommentByPostID(postId int) (error, []model.Post, []model.Comment) {
	var posts []model.Post
	var comments []model.Comment
	sqlStr := "select username,userPost,commentName,comment from userComment,userPost where userComment.PostID=?=userPost.ID;"
	rows, err := dB.Query(sqlStr, postId)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, posts, comments
		}
		return err, posts, comments
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var comment model.Comment
		err := rows.Scan(&post.Username, &post.Txt, &comment.Username, &comment.Txt)
		if err != nil {
			if err == sql.ErrNoRows {
				return err, posts, comments
			}
		}

		posts = append(posts, post)
		comments = append(comments, comment)
	}
	return err, posts, comments
}
