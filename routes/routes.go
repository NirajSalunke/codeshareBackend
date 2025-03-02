package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	r.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	roomRoutes := r.Group("/room")
	fileRoutes := r.Group("/file")
	userRoutes := r.Group("/user")
	LoadRoomRoutes(roomRoutes)
	LoadFileRoutes(fileRoutes)
	LoadUserRoutes(userRoutes)
}
