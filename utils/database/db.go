package database

import (
	"fmt"
	apps_config "go-simple/utils/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	var db *gorm.DB

	if apps_config.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", apps_config.DB_USER, apps_config.DB_PASSWORD, apps_config.DB_HOST, apps_config.DB_PORT, apps_config.DB_NAME)
		db, _ = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if apps_config.DB_DRIVER == "pgsql" {
		dsnPostgresql := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", apps_config.DB_HOST, apps_config.DB_USER, apps_config.DB_PASSWORD, apps_config.DB_NAME, apps_config.DB_PORT)
		db, _ = gorm.Open(postgres.Open(dsnPostgresql), &gorm.Config{})
	}

	if db == nil {
		panic("Can't connect to the database")
	}

	log.Println("Connected to database")
	return db
}
