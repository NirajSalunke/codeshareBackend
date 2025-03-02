package controllers

import (
	"net/http"

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

	if err := config.DB.Create(&currFile).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to save file",
			"message": err.Error(),
		})
		return
	}

	room.Files = append(room.Files, currFile)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File created successfully",
		"file":    currFile,
		"room":    room,
	})
}

func UpdateFile(c *gin.Context) {
	id := c.Param("id")
	var oldFile models.File
	if res := config.DB.First(&oldFile, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "File not found",
			"error":   res.Error,
		})
		return
	}
	var currFile models.File
	if err := c.BindJSON(&currFile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false, "message": "Failed to retrieve from body"})
		return
	}
	if res := config.DB.Model(&oldFile).Updates(currFile); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update file",
			"error":   res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File updated successfully",
	})
}

func DeleteFile(c *gin.Context) {
	id := c.Param("id")
	var file models.File

	if res := config.DB.First(&file, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "File not found",
			"error":   res.Error.Error(),
		})
		return
	}

	if res := config.DB.Delete(&file, id); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete file",
			"error":   res.Error.Error(),
		})
		return
	}

	var room models.Room
	if res := config.DB.Preload("Files").First(&room, file.RoomID); res.Error == nil {

		var updatedFiles []models.File
		for _, f := range room.Files {
			if f.ID != file.ID {
				updatedFiles = append(updatedFiles, f)
			}
		}
		room.Files = updatedFiles
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File deleted successfully",
	})
}

func GetFileById(c *gin.Context) {
	id := c.Param("id")
	var file models.File
	if res := config.DB.First(&file, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "File not found",
			"error":   res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "File found",
		"file":    file,
	})
}
