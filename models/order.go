package models

import "time"

type Order struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    StockID   uint      `gorm:"not null"`
    Type      string    `gorm:"not null"` 
    Quantity  uint      `gorm:"not null"`
    Price     float64   `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
