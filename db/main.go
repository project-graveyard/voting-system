package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
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

	conn, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Cannot create the db connection: %s", err)
	}

	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatalf("Ping Error: %s", err)
	}

	fmt.Println("Database connected!")
}
