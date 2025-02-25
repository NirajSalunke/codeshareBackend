package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/controllers"
)

func LoadRoomRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllRooms)
	r.GET("/:id", controllers.GetRoomById)
	r.GET("/:id/files", controllers.GetAllFilesOfRoom)
	r.POST("/", controllers.CreateNewRoom)
	r.PATCH("/:id", controllers.UpdateRoom)
	r.DELETE("/:id", controllers.DeleteRoom)

}
