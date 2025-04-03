package models

import (
	"time"
)

type OmikujiResult struct {
	ID        uint      `gorm:"primaryKey"`
	Result    string    `gorm:"size:10;not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
