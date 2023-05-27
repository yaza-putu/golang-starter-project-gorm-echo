package core

import (
	"database/sql"
	"fmt"
	"gitlab.com/putuyaza/antrian/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Mysql connection
func Mysql() *gorm.DB {
	dbConfig := config.Database()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	return db
}
