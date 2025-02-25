package main

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/config"
	"www.github.com/NirajSalunke/codeShare/models"
	"www.github.com/NirajSalunke/codeShare/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectToDatabase()
	models.MigrateModels()
}

func main() {
	r := gin.Default()
	routes.LoadRoutes(r)
	r.Run()
}
