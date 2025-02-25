package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/codeShare/config"
	"www.github.com/NirajSalunke/codeShare/helpers"
	"www.github.com/NirajSalunke/codeShare/models"
)

func GetAllRooms(c *gin.Context) {
	var rooms []models.Room

	result := config.DB.Preload("Files").Find(&rooms, &models.Room{
		IsPrivate: false,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success":     true,
		"publicRooms": rooms,
	})
}

func CreateNewRoom(c *gin.Context) {
	var room models.Room

	if err := c.Bind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	hashedPass, err := helpers.HashPassword(room.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	room.Password = hashedPass

	currTime := time.Now()
	expirationTime := currTime.Add(5 * 24 * time.Hour)
	room.ExpiresAt = &expirationTime
	room.Files = []models.File{}
	if result := config.DB.Create(&room); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Room Created Successfully",
		"room":    room,
	})
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")

	res := config.DB.Delete(&models.Room{}, id)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Room Deleted Successfully",
	})
}

func GetRoomById(c *gin.Context) {
	id := c.Param("id")
	var room models.Room

	if res := config.DB.Preload("Files").First(&room, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Room not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"room":    room,
	})
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	var room models.Room
	if err := c.Bind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var currRoom models.Room
	if res := config.DB.Preload("Files").First(&currRoom, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Room not found",
		})
		return
	}

	if res := config.DB.Model(&currRoom).Updates(room); res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Room Updated Successfully",
	})
}

func GetAllFilesOfRoom(c *gin.Context) {
	id := c.Param("id")
	var room models.Room
	if res := config.DB.Preload("Files").First(&room, id); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Room not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"files":   room.Files,
	})
}
