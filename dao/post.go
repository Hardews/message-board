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

func GetPost(username string) (error, []string, []string) {
	var output, time []string
	var user model.Post
	var Time string
	sqlStr := "select userPost,time from userPost where username = ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return err, output, time
		} else {
			return err, output, time
		}
	}
	for rows.Next() {
		err := rows.Scan(&user.Txt, &Time)
		if err != nil {
			return rows.Err(), output, time
		}
		output = append(output, user.Txt)
		time = append(time, Time)
	}
	return nil, output, time
}

func DeletePost(username, userPost string) error {
	sqlStr := "delete userPost from userPost where username=? and userPost=?"
	_, err := dB.Exec(sqlStr, username, userPost)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPost() (error, []string, []string, []string) {
	var user model.Post
	var username, txt, Time []string
	var time string
	sqlStr1 := "select username,userPost,time from userPost"
	rows, err := dB.Query(sqlStr1)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, username, txt, Time
		}
		return err, username, txt, Time
	}
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Txt, &time)
		if err != nil {
			return err, username, txt, Time
		}
		username = append(username, user.Username)
		txt = append(txt, user.Txt)
		Time = append(Time, time)
	}
	return err, username, txt, Time
}

func ChangePost(username, newPost, oldUserPost string) error {
	sqlStr := "update userPost set userPost = ? where username = ? and userPost = ?"
	_, err := dB.Exec(sqlStr, newPost, username, oldUserPost)
	if err != nil {
		return err
	}
	return nil
}

func SelectPost(postUsername, post string) (model.User, error) {
	var u model.User
	sqlStr := "select username,userPost from userPost where username=? and userPost=?"
	err := dB.QueryRow(sqlStr, postUsername, post).Scan(&u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}

func SelectByPostId(postName, userPost string) (int, error) {
	var u model.Post
	sqlStr := "select id from userPost where username=? and userPost=?"
	err := dB.QueryRow(sqlStr, postName, userPost).Scan(&u.PostID)
	if err != nil {
		return u.PostID, err
	}
	return u.PostID, err
}
