package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type env struct {
	DB db_cfg `json:"db"`
}

type db_cfg struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Addr     string `json:"addr"`
	Database string `json:"database"`
}

func Init() (*sql.DB, error) {
	file, err := os.ReadFile("../.env.json")
	if err != nil {
		log.Panic(err)
	}

	var os_env env

	err = json.Unmarshal(file, &os_env)
	if err != nil {
		log.Panic(err)
	}

	config := mysql.Config{
		User:                 os_env.DB.Username,
		Passwd:               os_env.DB.Passwd,
		Net:                  "tcp",
		Addr:                 os_env.DB.Addr,
		DBName:               os_env.DB.Database,
		AllowNativePasswords: true,
	}

	return sql.Open("mysql", config.FormatDSN())
}
