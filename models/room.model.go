package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	OwnerID   *uint      `gorm:"index" json:"ownerId"`
	Password  string     `gorm:"not null" json:"password"`
	IsPrivate bool       `gorm:"default:false" json:"isPrivate"`
	ExpiresAt *time.Time `gorm:"index" json:"expiresAt"`
	Files     []File     `gorm:"not null" json:"files"`
}
