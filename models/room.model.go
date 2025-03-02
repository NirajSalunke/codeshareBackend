package models

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	OwnerID   string     `gorm:"default:"" json:"ownerId"`
	Name      string     `gorm:"not null;" json:"name"`
	Password  string     `gorm:"" json:"password"`
	IsPrivate bool       `gorm:"default:false" json:"isPrivate"`
	ExpiresAt *time.Time `gorm:"index" json:"expiresAt"`
	Files     []File     `gorm:"not null" json:"files"`
}
