package dao

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func InitDB()  {
	db,err := sql.Open("mysql","root:lmh123@tcp(127.0.0.1:3306)/userdata")
	if err != nil{
		fmt.Println(err)
	}
	Db = db
}
