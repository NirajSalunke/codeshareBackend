package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/config"
	"www.github.com/NirajSalunke/codeShare/models"
)

func CreateFile(c *gin.Context) {
	id := c.Param("id")
	var currFile models.File
	var room models.Room

	if err := c.BindJSON(&currFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false, "message": "Failed to create file object!"})
		return
	}

	if res := config.DB.Preload("Files").First(&room, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Room not found",
		})
		return
	}

	currFile.RoomID = room.ID
	currFile.FileType = strings.Split(currFile.Name, ".")[1]

	if err := config.DB.Create(&currFile); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to save file",
			"message": err.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File created successfully",
		"file":    currFile,
		"room":    room,
	})
}
