package service

import (
	"database/sql"
	"message-board/dao"
)

func CheckPassword(username,password string)  (bool,error){
	 err,user := dao.SelectByUsername(username)
	 if err!=nil{
		 if err==sql.ErrNoRows {
			 return false,nil
		 }else {
			 return false,err
		 }
	 }
	if user.Password!=password {
		return false,err
	}
	return true,nil
}


