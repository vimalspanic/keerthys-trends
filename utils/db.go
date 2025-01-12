package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vimalspanic/keerthys-trends/config"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.GetConfig("db_user"),
		config.GetConfig("db_password"),
		config.GetConfig("db_host"),
		config.GetConfig("db_port"),
		config.GetConfig("db_name"),
	)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}
