package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func ConecDB() {
	dsn := "root:123456(127.0.0.1:3306)/todo_list"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Lỗi kêt nối", err)
	}
	log.Println("Kết nối thành công ")
}
