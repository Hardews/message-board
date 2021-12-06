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

func CheckUsername(username string)  (error,bool){
     err,user := dao.SelectByUsername(username)
	 if err != nil{
		 if err == sql.ErrNoRows{
			 err = nil
			 return err,true
		 }else {
			 return err,false
		 }
	 }
	 if user.Username == username{
		 return nil,false
	 }
	 return nil,true
}


