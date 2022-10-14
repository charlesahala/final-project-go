package router

import (
	"final-project-go/controllers"

	"github.com/gin-gonic/gin"
)

func StartService() *gin.Engine {
	router := gin.Default()

	router.POST("/users/register", controllers.CreateUser)
	router.GET("/orders", controllers.GetUser)
	router.PUT("/users", controllers.UpdateUser)
	router.DELETE("/users", controllers.DeleteUser)

	router.POST("/photos", controllers.CreatePhoto)
	router.GET("/photos", controllers.GetPhoto)
	router.PUT("/photos/:photoId", controllers.UpdatePhoto)
	router.DELETE("/photos/:photoId", controllers.DeletePhoto)

	router.POST("/comments", controllers.CreateComment)
	router.GET("/comments", controllers.GetComment)
	router.PUT("/comments/:commentId", controllers.UpdateComment)
	router.DELETE("/comments/:commentId", controllers.DeleteComment)

	router.POST("/socialmedias", controllers.CreateSocialMedia)
	router.GET("/socialmedias", controllers.GetSocialMedia)
	router.PUT("/socialmedias/:socialMediaId", controllers.UpdateSocialMedia)
	router.DELETE("/socialmedias/:socialMediaId", controllers.DeleteSocialMedia)

	return router
}
