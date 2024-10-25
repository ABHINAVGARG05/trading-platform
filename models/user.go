package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Balance   float64   `gorm:"not null;default:10000"` // Initial balance for trading
    CreatedAt time.Time
    UpdatedAt time.Time
}
