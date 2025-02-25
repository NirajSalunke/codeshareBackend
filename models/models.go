package models

import "www.github.com/NirajSalunke/codeShare/config"

func MigrateModels() {
	config.DB.AutoMigrate(&User{})
	config.DB.AutoMigrate(&Room{})
	config.DB.AutoMigrate(&File{})
}
