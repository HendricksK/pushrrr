package database

import (
	// "log"
	"database/sql"
	"fmt"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"

	helpers "github.com/HendricksK/pushrrr/app/helpers"
)

var dbHost = "pushrrr-mariadb:3306"

func Open() *sql.DB {

	// parseTime=true https://github.com/go-sql-driver/mysql#timetime-support
	dbName := helpers.GetEnvVar("db.database")
	dbUser := helpers.GetEnvVar("db.user")
	dbPass := helpers.GetEnvVar("db.password")

	fmt.Println(dbName)

	database, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?parseTime=true")
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		helpers.Log(err.Error(), filename, line)
		panic(err)
	}

	// See "Important settings" section.
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	return database
}

func Close(conn *sql.DB) {
	defer conn.Close()
}
