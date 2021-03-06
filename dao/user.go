package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
)

func SelectByUsername(username string) (error, model.User) {
	var user model.User
	sqlStr := "SELECT username,password FROM user_Info WHERE username = ?;"
	err := dB.QueryRow(sqlStr, username).Scan(&user.Username, &user.Password)

	if err != nil {
		return err, user
	}

	return nil, user
}

func WriteIn(username, password string) error {
	sqlStr := "insert into user_Info (username,password) values (?,?);"
	_, err := dB.Exec(sqlStr, username, password)
	if err != nil {
		return err
	}
	return nil
}

func ChangePassword(username, password string) error {
	sqlStr := "update user_Info set password = ? where username = ?;"
	_, err := dB.Exec(sqlStr, password, username)
	if err != nil {
		return err
	}
	return nil
}

func WriteUserInfoIN(userInfo model.UserInfo, username string) error {
	sqlStr := "insert into user_extra_info (username,Name,Professional,School,Specialty) values (?,?,?,?,?);"
	_, err := dB.Exec(sqlStr, username, userInfo.Name, userInfo.Professional, userInfo.School, userInfo.Specialty)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(username string) (model.UserInfo, error) {
	var user model.UserInfo
	sqlStr := "select * from user_extra_info where username = ?"
	err := dB.QueryRow(sqlStr, username).Scan(&user.Id, &username, &user.Name, &user.Professional, &user.School, &user.Specialty)
	if err != nil {
		return user, err
	}
	return user, err
}

func ChangeUserInfo(userInfo model.UserInfo) error {
	sqlStr := "update user_extra_info set Name = ? ,Professional=? ,School=? ,Specialty=? where id = ?"
	_, err := dB.Exec(sqlStr, userInfo.Name, userInfo.Professional, userInfo.School, userInfo.Specialty, userInfo.Id)
	if err != nil {
		return err
	}
	return nil
}
