package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Init() (*sql.DB, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panic(err)
	}

	config := mysql.Config{
		User:                 os.Getenv("USERNAME"),
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")),
		DBName:               os.Getenv("DATABASE"),
		AllowNativePasswords: true,
	}

	return sql.Open("mysql", config.FormatDSN())
}
