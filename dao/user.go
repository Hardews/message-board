package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
)

func SelectByUsername(username string) (error, model.User) {
	var user model.User
	sqlStr := "SELECT username,password FROM userInfo WHERE username = ?;"
	rows, err := dB.Query(sqlStr, username)

	if err != nil {
		return err, user
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Password)
		if err != nil {
			return err, user
		}
	}

	return nil, user
}

func WriteIn(username, password string) error {
	sqlStr := "insert into userInfo (username,password) values (?,?);"
	_, err := dB.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}

func ChangePassword(username, password string) error {
	sqlStr := "update userInfo set password = ? where username = ?;"
	_, err := dB.Exec(sqlStr, password, username)
	if err != nil {
		return err
	}
	return nil
}
