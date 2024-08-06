package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySQL() (*sql.DB, error) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDBName := os.Getenv("MYSQL_DBNAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDBName)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Printf("Failed to open MySQL connection: %s", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		if err = db.Ping(); err != nil {
			log.Printf("Failed to ping MySQL: %s", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		log.Println("Connected to MySQL!")
		return db, nil
	}
	return nil, err
}
