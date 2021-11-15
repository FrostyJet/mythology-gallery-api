package models

import (
	"fmt"
	"log"

	"github.com/frostyjet/mythology-gallery-api/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&God{})

	ares := God{Name: "Ares"}
	db.Create(&ares)

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}
