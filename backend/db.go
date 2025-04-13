package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(localhost:3306)/academic"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Tidak bisa terhubung ke database:", err)
	}
}
