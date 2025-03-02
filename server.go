package main

import (
	"time"

	"github.com/gin-contrib/cors"
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

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowOrigins:    []string{"*"}, // Allow frontend origin
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	routes.LoadRoutes(r)

	r.Run()
}
