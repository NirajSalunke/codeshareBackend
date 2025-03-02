package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/controllers"
)

func LoadRoomRoutes(r *gin.RouterGroup) {

	r.GET("/", controllers.GetAllRooms)
	r.POST("/join", controllers.CheckRoomPass)
	r.GET("/:name", controllers.GetRoomByName)
	r.GET("/:name/files", controllers.GetAllFilesOfRoom)
	r.POST("/", controllers.CreateNewRoom)
	r.PATCH("/:name", controllers.UpdateRoom)
	r.DELETE("/:name", controllers.DeleteRoom)

}
