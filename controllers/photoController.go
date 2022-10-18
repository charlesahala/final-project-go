package controllers

import (
	"final-project-go/database"
	"final-project-go/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}

	if Photo.UserID == userID {
		err := db.Debug().Create(&Photo).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	err := c.ShouldBindJSON(&Photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, Photo)
		// gin.H{
		// "id":         Photo.ID,
		// "title":      Photo.Title,
		// "caption":    Photo.Caption,
		// "photo_url":  Photo.PhotoURL,
		// "user_id":    Photo.UserID,
		// "created_at": Photo.CreatedAt,
	// }
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}
	User := models.User{}

	err := db.Model(&Photo).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	var PhotoDatas = []models.Photo{}
	userDatas := models.User{
		Email:    User.Email,
		Username: User.Username,
	}

	condition := false
	photoDatas := models.Photo{}
	for i, Photo := range PhotoDatas {
		if Photo.UserID == userID {
			condition = true
			photoDatas = PhotoDatas[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "data not found",
			"message": "photos data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Photo": photoDatas,
		"User": userDatas,
	})
}

func PutPhoto(c *gin.Context) {

}

func DeletePhoto(c *gin.Context) {

}
