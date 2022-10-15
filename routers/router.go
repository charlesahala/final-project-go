package router

import (
	"final-project-go/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/", controllers.UserUpdate)
		userRouter.DELETE("/", controllers.UserDelete)
	}

	// photoRouter := router.Group("/photos")
	// {
	// 	photoRouter.POST("/", controllers.CreatePhoto)
	// 	photoRouter.GET("/", controllers.GetPhoto)
	// 	photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
	// 	photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	// }

	// commentRouter := router.Group("/comments")
	// {
	// 	commentRouter.POST("/", controllers.CreateComment)
	// 	commentRouter.GET("/", controllers.GetComment)
	// 	commentRouter.PUT("/:commentId", controllers.UpdateComment)
	// 	commentRouter.DELETE("/:commentId", controllers.DeleteComment)
	// }

	// socialMediasRouter := router.Group("/socialmedias")
	// {
	// 	socialMediasRouter.POST("/", controllers.CreateSocialMedia)
	// 	socialMediasRouter.GET("/", controllers.GetSocialMedia)
	// 	socialMediasRouter.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
	// 	socialMediasRouter.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	// }

	return router
}
