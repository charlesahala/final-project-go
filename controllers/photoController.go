package controllers

import (
	"final-project-go/database"
	"final-project-go/helpers"
	"final-project-go/models"
	"strconv"

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
		"User":  userDatas,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert photoId",
		})
	}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	if err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&Photo, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	photoData := models.Photo {
		ID: Photo.ID,
		Title: Photo.Title,
		Caption: Photo.Caption,
		PhotoURL: Photo.PhotoURL,
		UserID: Photo.UserID,
		UpdatedAt: Photo.UpdatedAt,
	}

	c.JSON(http.StatusOK, photoData)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", userID).Delete(&Photo).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
