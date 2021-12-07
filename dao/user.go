package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"message-board/model"
)

var db,err = sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/userdata")

func InitDB()  {
	if err != nil{
		fmt.Println(err)
	}
}

func SelectByUsername(username string)(error,model.User){
	 var user model.User
     sqlStr := "select username,password from userInfo where username=?"
	 err := db.QueryRow(sqlStr,username).Scan(&user.Username,&user.Password)
	 if err!=nil{
		 return err,user
	 }
	 return nil,user
}

func WriteIn(username,password string)  error{
	 sqlStr := "insert into userInfo (username,password) values (?,?)"
	 _,err := db.Exec(sqlStr,username,password)
	 if err != nil{
		 return err
	 }
	 return nil
}

func ChangePassword(username,password string)  error{
	sqlStr := "update userInfo set password = ? where username = ?"
	_,err := db.Exec(sqlStr,password,username)
	if err != nil {
		return err
	}
	return nil
}
