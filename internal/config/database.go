package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db_config := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		GetEnv("DB_USERNAME", "root"),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "3306"),
		GetEnv("DB_NAME", "victoria_property"),
	)

	db, err := sql.Open("mysql", db_config)
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("DB ping error:", err)
	}

	log.Println("Database connected")
	return db
}
