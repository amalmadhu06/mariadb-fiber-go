package db

import (
	"fmt"
	"github.com/amalmadhu06/mariadb-fiber-go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	//dbInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	dbInfo := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName)
	fmt.Println("connection string: ", dbInfo)
	//test := "root:1234@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	return db, err
}
