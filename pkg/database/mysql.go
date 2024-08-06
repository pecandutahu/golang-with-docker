package database

import (
	"log"
	"os"
	"product/internal/domain"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL() (*gorm.DB, error) {
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDBName := os.Getenv("MYSQL_DBNAME")

	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// 	mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDBName)
	dsn := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to open MySQL connection: %s", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		sqlDB, _ := db.DB()
		if err = sqlDB.Ping(); err != nil {
			log.Printf("Failed to ping MySQL: %s", err.Error())
			time.Sleep(2 * time.Second)
			continue
		}
		log.Println("Connected to MySQL!")

		// Migrasi
		db.AutoMigrate(&domain.Product{})

		return db, nil
	}
	return nil, err
}
