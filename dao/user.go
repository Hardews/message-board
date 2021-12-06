package dao

import (
	"database/sql"
	"fmt"
	"message-board/model"
)

var db *sql.DB

func InitDB()  {
	DB,err := sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/user")
	if err != nil{
		fmt.Println(err)
	}

	db = DB
}

func SelectByUsername(username string)(error,model.User){
	var user model.User
     sqlStr := "select id,password from userInfo where username=?"
	 rows := db.QueryRow(sqlStr,username)
	 if rows.Err()!=nil{
		 return rows.Err(),user
	 }
	 err := rows.Scan(&user.Id,&user.Password)
	 if err!=nil{
		 return err,user
	 }
	 return nil,user
}
