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

func WriteUserInfoIN(userInfo model.UserInfo, username string) error {
	sqlStr := "insert into userExtraInfo (username,Name,Professional,School,Specialty) values (?,?,?,?,?);"
	_, err := dB.Exec(sqlStr, username, userInfo.Name, userInfo.Professional, userInfo.School, userInfo.Specialty)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(username string) (model.UserInfo, error) {
	var user model.UserInfo
	sqlStr := "select * from userExtraInfo where username = ?"
	err := dB.QueryRow(sqlStr, username).Scan(&user.Id, &user.Name, &user.Professional, &user.School, &user.Specialty)
	if err != nil {
		return user, err
	}
	return user, err
}

func ChangeUserInfo(userInfo model.UserInfo, username string) error {
	sqlStr := "update userExtraInfo set Name=? and Professional=? and School=? and Specialty=? where username = ?"
	_, err := dB.Exec(sqlStr, userInfo.Name, userInfo.Professional, userInfo.School, userInfo.Specialty, username)
	if err != nil {
		return err
	}
	return nil
}
