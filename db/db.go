// db package handles the creation of a database connection
package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	"github.com/go-sql-driver/mysql" // mysql driver
)

// env holds environment variables
type env struct {
	DB db_cfg `json:"db"`
}

// db_cfg hold database config from env
type db_cfg struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Addr     string `json:"addr"`
	Database string `json:"database"`
}

// Init creates a new database connection
func Init() (*sql.DB, error) {
	file, err := os.ReadFile("../.env.json")
	if err != nil {
		log.Panicf("Reading file: %s", err)
	}

	var os_env env // holds environment variables

	// Unmarshal env file into os_env
	err = json.Unmarshal(file, &os_env)
	if err != nil {
		log.Panicf("Unmarshal json: %s", err)
	}

	// define mysql connection config
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
