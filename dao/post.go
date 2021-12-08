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

func GetPost(username string) (error, []string) {
	var output []string
	var middle string
	sqlStr := "select userPost from userPost where username = ?"
	rows, err := dB.Query(sqlStr, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return err, output
		} else {
			return err, output
		}
	}
	for rows.Next() {
		err := rows.Scan(&middle)
		if err != nil {
			return rows.Err(), output
		}
		output = append(output, middle)
	}
	return nil, output
}

func DeletePost(username, userPost string) error {
	sqlStr := "delete userPost from userPost where username=? and userPost=?"
	_, err := dB.Exec(sqlStr, username, userPost)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPost() (error, []string, []string) {
	var output, user []string
	var middle, username string
	sqlStr1 := "select username,userPost from userPost"
	rows, err := dB.Query(sqlStr1)
	if err != nil {
		if err == sql.ErrNoRows {
			err = sql.ErrNoRows
			return err, user, output
		} else {
			return err, user, output
		}
	}
	for rows.Next() {
		err := rows.Scan(&username, &middle)
		user = append(user, username)
		output = append(output, middle)
		if err != nil {
			return rows.Err(), user, output
		}
	}
	return nil, user, output
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
	err := dB.QueryRow(sqlStr, postUsername, post).Scan(&u.Username, &u.Post)
	if err != nil {
		return u, err
	}
	return u, nil
}
