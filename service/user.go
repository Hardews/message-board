package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func CheckPassword(username, password string) (bool, error) {
	err, user := dao.SelectByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
	}
	if user.Password != password {
		return false, err
	}
	return true, nil
}

func CheckUsername(username string) (error, bool) {
	err, user := dao.SelectByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, false
		} else {
			return err, false
		}
	}
	if user.Username == username {
		return err, false
	} else {
		return nil, true
	}
}

func WriteIn(username, password string) error {
	err := dao.WriteIn(username, password)
	if err != nil {
		return err
	}
	return nil
}

func CheckLength(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func WriteInfo(userInfo model.UserInfo, username string) error {
	err := dao.WriteUserInfoIN(userInfo, username)
	if err != nil {
		return err
	}
	return nil
}

func GetInfo(username string) (model.UserInfo, error) {
	userInfo, err := dao.GetUserInfo(username)
	if err != nil {
		return userInfo, err
	}
	return userInfo, err
}

func ChangeInfo(newInfo model.UserInfo, username string) (bool, error) {
	err := dao.ChangeUserInfo(newInfo, username)
	if err != nil {
		return false, err
	}
	return true, err
}
