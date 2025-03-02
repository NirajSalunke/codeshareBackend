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

	if err := c.BindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	// fmt.Printf("%+v", room)
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
	name := c.Param("name")

	res := config.DB.Delete(&models.Room{}, models.Room{
		Name: name,
	})
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

func GetRoomByName(c *gin.Context) {
	name := c.Param("name")
	var room models.Room

	if res := config.DB.Preload("Files").First(&room, models.Room{
		Name: name,
	}); res.Error != nil {
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

func CheckRoomPass(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var neededRoom models.Room
	if res := config.DB.Preload("Files").First(&neededRoom, models.Room{
		Name: input.Name,
	}); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Room not found",
		})
		return
	}

	if !helpers.CheckPasswordHash(input.Password, neededRoom.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Incorrect password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"room":    neededRoom,
	})
}

func UpdateRoom(c *gin.Context) {
	name := c.Param("name")
	var room models.Room
	if err := c.Bind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var currRoom models.Room
	if res := config.DB.Preload("Files").First(&currRoom, models.Room{
		Name: name,
	}); res.Error != nil {
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
	name := c.Param("name")
	var room models.Room
	if res := config.DB.Preload("Files").First(&room, models.Room{
		Name: name,
	}); res.Error != nil {
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
