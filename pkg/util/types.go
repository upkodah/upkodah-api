package util

import (
	"time"
)

// stringify Array
type UrlArrayStr string
type ArrayStr string
type Url string

// DB Default Model
type DBModel struct {
	ID        uint       `json:"id" gorm:"primary_key; autoIncrement"`
	CreatedAt time.Time  `json:"createdAt" gorm:"default:CURRENT_TIMESTAMP;autoCreateTime;not null"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime;not null"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}
