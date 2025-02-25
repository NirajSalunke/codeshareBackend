package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ClerkID string `gorm:"unique;not null" json:"clerkid"`
	Email   string `gorm:"unique;not null" json:"email"`
	Rooms   []Room `gorm:"foreignKey:OwnerID" json:"rooms"`
}
