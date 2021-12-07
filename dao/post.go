package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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
			err = sql.ErrNoRows
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

func GetAllPost() (error, []string) {
	sqlStr := "select userPost from userPost"
	var output []string
	var middle string
	rows, err := dB.Query(sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			err = sql.ErrNoRows
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
