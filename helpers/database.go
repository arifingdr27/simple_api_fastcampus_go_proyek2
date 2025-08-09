package helpers

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySql() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetEnv("DB_USERNAME", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_HOST", ""), GetEnv("DB_PORT", ""))
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connection database")
	}

	logrus.Info("successfully connect to database")
}
