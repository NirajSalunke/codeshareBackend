package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/controllers"
)

func LoadFileRoutes(r *gin.RouterGroup) {
	r.POST("/create/:id", controllers.CreateFile)
	r.PUT("/:id", controllers.UpdateFile)
	r.DELETE("/:id", controllers.DeleteFile)
}
