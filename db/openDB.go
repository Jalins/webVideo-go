package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:Hzl123**@tcp(127.0.0.1:3306)/video_server?charset-utf8")

	if err != nil {
		log.Panic(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Panic("mysql连接错误", err)
	}

}
