package database

import (
	// "log"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var database *sql.DB

// 192.168.0.134:3309
// var db_host = "172.17.0.3:3306";
var db_host = "127.0.0.1:3306"

func Open() *sql.DB {

	fmt.Println("Shell:", os.Getenv("SHELL"))

	// parseTime=true https://github.com/go-sql-driver/mysql#timetime-support
	db := getEnv("DATABASE")
	database, err := sql.Open("mysql", "dev-sacos:riuc6Y5Rend$W5%BSG*S$@tcp("+db_host+")/sacos_dev?parseTime=true")
	if err != nil {
		panic(err)

	}
	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	return database
}

func Close(conn *sql.DB) {
	conn.Close()
}
