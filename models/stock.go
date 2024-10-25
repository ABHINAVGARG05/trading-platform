package models

import "time"

type Stock struct {
    ID        uint      `gorm:"primaryKey"`
    Symbol    string    `gorm:"unique;not null"`
    Price     float64   `gorm:"not null"`
    Volume    uint      `gorm:"not null"` // Track volume for demand/supply-based pricing
    CreatedAt time.Time
    UpdatedAt time.Time
}
