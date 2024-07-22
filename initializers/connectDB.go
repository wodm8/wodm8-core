package initializers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDb() {
	var cfg = Cfg
	var err error

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	DB, err = gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})
	if err != nil {
		log.Fatal("Error db connection", err)
	}
}
