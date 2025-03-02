package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"www.github.com/NirajSalunke/codeShare/helpers"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		helpers.PrintRed("Error:- Failed to Load env.")
	}
}

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		helpers.PrintRed("Error:- Failed to connect to database:- ")
		log.Fatal(err.Error())
	}
	helpers.PrintGreen("Connected to Database Successfully.")
}
