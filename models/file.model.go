package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	RoomID   uint   `gorm:"not null;index;constraint:OnDelete:CASCADE;" json:"roomId"`
	Name     string `gorm:"not null" json:"name"`
	FileType string `gorm:"not null" json:"fileType"`
	Content  string ` json:"content"`
}
