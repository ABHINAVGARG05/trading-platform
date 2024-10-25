package models

import "time"

type Portfolio struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`     // Reference to the user
    StockID   uint      `gorm:"not null"`     // Reference to the stock
    Quantity  uint      `gorm:"not null"`     // Number of shares owned
    AvgPrice  float64   `gorm:"not null"`     // Average price per share
    CreatedAt time.Time
    UpdatedAt time.Time
}
