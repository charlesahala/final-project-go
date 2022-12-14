package router

import (
	"final-project-go/controllers"
	"final-project-go/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.UserAuthentication(), controllers.UserPut)
		userRouter.DELETE("/", middlewares.UserAuthentication(), controllers.UserDelete)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.UserAuthentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.UserAuthentication(), middlewares.CommentAuthorization())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentId", controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", controllers.DeleteComment)
	}

	socialMediasRouter := router.Group("/socialmedias")
	{
		socialMediasRouter.Use(middlewares.UserAuthentication(), middlewares.SocMedAuthorization())
		socialMediasRouter.POST("/", controllers.CreateSocMed)
		socialMediasRouter.GET("/", controllers.GetSocMed)
		socialMediasRouter.PUT("/:socialMediaId", controllers.UpdateSocMed)
		socialMediasRouter.DELETE("/:socialMediaId", controllers.DeleteSocMed)
	}

	return router
}
