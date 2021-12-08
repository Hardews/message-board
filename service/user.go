package service

import (
	"database/sql"
	"message-board/dao"
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
