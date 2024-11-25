package database

import (
	"fmt"
	"log"

	"github.com/phn00dev/go-crud-temp/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() {
	configs := config.Get()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		configs.Database.Host,
		configs.Database.Username,
		configs.Database.Password,
		configs.Database.DbName,
		configs.Database.Port,
		configs.Database.SllMode,
		configs.Database.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connection error : %v", err)
		return
	}
	DB = db
}
