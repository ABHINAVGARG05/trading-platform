package models

import "time"

type Stock struct {
    ID        uint      `gorm:"primaryKey"`
    Symbol    string    `gorm:"unique;not null"`
    Price     float64   `gorm:"not null"`
    Volume    uint      `gorm:"not null"`
    Supply    uint      `gorm:"default:0"`
    Demand    uint      `gorm:"default:0"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
