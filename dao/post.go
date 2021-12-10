package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
	"strconv"
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

	sqlStr = "drop table post?;" //都删除留言了，再留着这个表没啥意义
	_, err = dB.Exec(sqlStr, PostID)
	return nil
}

func ChangePost(newPost string, PostID int) error {
	sqlStr := "update userPost set userPost = ? where postID = ?"
	_, err := dB.Exec(sqlStr, newPost, PostID)
	if err != nil {
		return err
	}

	strNum := "post" + strconv.Itoa(PostID)
	sqlStrM := "update " + strNum + " set txt = ? where id = 1"
	sqlStr = sqlStrM

	_, err = dB.Exec(sqlStr, PostID, newPost)
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

func SelectCommentsSection(PostID int) (error, []model.Post) {
	var posts []model.Post

	strNum := "post" + strconv.Itoa(PostID)
	sqlStrM := "select * from " + strNum
	sqlStr := sqlStrM

	rows, err := dB.Query(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, posts
		}
		return err, posts
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.PostID, &post.Username, &post.Txt, &post.LikeNum)
		if err != nil {
			if err == sql.ErrNoRows {
				return err, posts
			}
		}

		posts = append(posts, post)
	}
	return err, posts
}
