package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/ABHINAVGARG05/trading-platform/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=postgres password=new_password dbname=trading_platform port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB.AutoMigrate(&models.User{}, &models.Stock{}, &models.Order{})
}
