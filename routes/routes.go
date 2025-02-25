package routes

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	roomRoutes := r.Group("/room")
	fileRoutes := r.Group("/file")
	userRoutes := r.Group("/user")
	LoadRoomRoutes(roomRoutes)
	LoadFileRoutes(fileRoutes)
	LoadUserRoutes(userRoutes)
}
