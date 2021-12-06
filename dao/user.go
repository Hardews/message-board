package dao

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func InitDB()  {
	DB,err := sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/user")
	if err != nil{
		fmt.Println(err)
	}

	db = DB
}
