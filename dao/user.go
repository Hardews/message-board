package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
	"strconv"
)

func SelectByUsername(username string) (error, model.User) {
	var user model.User
	sqlStr := "SELECT username,password FROM userInfo WHERE username = ?;"
	err := dB.QueryRow(sqlStr, username).Scan(&user.Username, &user.Password)

	if err != nil {
		return err, user
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
	err := dB.QueryRow(sqlStr, username).Scan(&user.Id, &username, &user.Name, &user.Professional, &user.School, &user.Specialty)
	if err != nil {
		return user, err
	}
	return user, err
}

func ChangeUserInfo(userInfo model.UserInfo) error {
	sqlStr := "update userExtraInfo set Name = ? ,Professional=? ,School=? ,Specialty=? where id = ?"
	_, err := dB.Exec(sqlStr, userInfo.Name, userInfo.Professional, userInfo.School, userInfo.Specialty, userInfo.Id)
	if err != nil {
		return err
	}
	return nil
}

func CreateCommentsSection(ID int, post model.Post) error {
	IdStr := strconv.Itoa(ID)
	sqlStrM := "CREATE TABLE `Post" + IdStr + "` (    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,  `username` VARCHAR(20) DEFAULT '', `txt` VARCHAR(20) DEFAULT '', `likeNum` int(20) DEFAULT 0,  PRIMARY KEY(`id`))ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
	sqlStr := sqlStrM
	_, err := dB.Exec(sqlStr)
	if err != nil {
		return err
	}

	sqlStrM = "insert into post" + IdStr + " (username,txt,likeNum) values (?,?,?)"
	sqlStr = sqlStrM
	_, err = dB.Exec(sqlStr, post.Username, post.Txt, 0)
	if err != nil {
		return err
	}
	return nil
}
