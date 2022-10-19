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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	if Comment.UserID == userID {
		err := db.Debug().Create(&Comment).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}
	User := models.User{}
	Photo := models.Photo{}

	err := db.Model(&Comment).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	var CommentDatas = []models.Comment{}
	userDatas := models.User{
		ID:       User.ID,
		Email:    User.Email,
		Username: User.Username,
	}
	photoDatas := models.Photo{
		ID:       Photo.ID,
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoURL: Photo.PhotoURL,
		UserID:   Photo.UserID,
	}

	condition := false
	commentDatas := models.Comment{}
	for i, Comment := range CommentDatas {
		if Comment.UserID == userID {
			condition = true
			commentDatas = CommentDatas[i]
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
		"Comment": commentDatas,
		"User":    userDatas,
		"Photo":   photoDatas,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "failed to convert photoId",
		})
	}

	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	if err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{Message: Comment.Message}).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().First(&Comment, userID).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	commentData := models.Comment{
		ID: Comment.ID,
		Message: Comment.Message,
		PhotoID: Comment.PhotoID,
		UserID: Comment.UserID,
		UpdatedAt: Comment.UpdatedAt,
	}

	c.JSON(http.StatusOK, commentData)
// 		gin.H{
// 		"id": Comment.ID,
// 		"message": Comment.Message,
// 		"photo_id": Comment.PhotoID,
// 		"user_id": Comment.UserID,
// 		"updated_at": Comment.UpdatedAt,
// 	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comment := models.Comment{}

	userID := uint(userData["id"].(float64))

	err := db.Debug().Where("id = ?", userID).Delete(&Comment).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
