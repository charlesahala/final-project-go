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

	err := c.ShouldBindJSON(&SocialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	SocialMedia := models.SocialMedia{}
	User := models.User{}

	err := db.Model(&SocialMedia).Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	var SocMedDatas = []models.SocialMedia{}
	userDatas := models.User{
		ID:       User.ID,
		Username: User.Username,
		// ProfileImageURL: User.ProfileImageURL,
	}

	condition := false
	socmedDatas := models.SocialMedia{}
	for i, SocialMedia := range SocMedDatas {
		if SocialMedia.UserID == userID {
			condition = true
			socmedDatas = SocMedDatas[i]
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
		"SocMed": socmedDatas,
		"User":    userDatas,
	})
}

func UpdateSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, err := strconv.Atoi(c.Param("socialMediaId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert photoId",
		})
	}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	if err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaURL: SocialMedia.SocialMediaURL}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&SocialMedia, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	socmedData := models.SocialMedia{
		ID: SocialMedia.ID,
		Name: SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
		UserID: SocialMedia.UserID,
		UpdatedAt: SocialMedia.UpdatedAt,
	}

	c.JSON(http.StatusOK, socmedData)
}

func DeleteSocMed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedia := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", userID).Delete(&SocialMedia).Error

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