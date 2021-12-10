package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func CheckPassword(username, password string) (bool, error) {
	err, user := dao.SelectByUsername(username)
	if err != nil {
		return false, err
	}
	_, flag := Interpretation(user.Password, password)
	if !flag {
		return false, nil
	}
	return true, nil
}

func CheckUsername(username string) (error, bool) {
	err, user := dao.SelectByUsername(username)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return err, true
		}
		return err, false
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

func ChangeInfo(newInfo model.UserInfo) (bool, error) {
	err := dao.ChangeUserInfo(newInfo)
	if err != nil {
		return false, err
	}
	return true, err
}

func ChangePassword(username, newPassword string) error {
	err := dao.ChangePassword(username, newPassword)
	if err != nil {
		return err
	}
	return err
}

func CheckInputInfo(username string, newUserInfo model.UserInfo) (error, model.UserInfo) {
	UserInfo, err := GetInfo(username)

	newUserInfo.Id = UserInfo.Id

	if err != nil {
		return err, UserInfo
	}
	if len(newUserInfo.Name) == 0 {
		newUserInfo.Name = UserInfo.Name
	}
	if len(newUserInfo.Professional) == 0 {
		newUserInfo.Professional = UserInfo.Professional
	}
	if len(newUserInfo.Specialty) == 0 {
		newUserInfo.Specialty = UserInfo.Specialty
	}
	if len(newUserInfo.School) == 0 {
		newUserInfo.School = UserInfo.School
	}
	return nil, newUserInfo
}
