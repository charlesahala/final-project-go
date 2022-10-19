package controllers

import (
	"final-project-go/database"
	"final-project-go/helpers"
	"final-project-go/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func CreateSocMed(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	SocialMedia := models.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	if SocialMedia.UserID == userID {
		err := db.Debug().Create(&SocialMedia).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaURL,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetSocMed(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := []models.SocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "Email", "Username")
	}).Where("user_id = ?", userID).Find(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func UpdateSocMed(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err = db.Debug().Where("id=?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	socialMediaID, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert",
		})
	}

	userID := uint(userData["id"].(float64))
	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err = db.Debug().Where("id = ?", socialMediaID).Delete(&SocialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
