package routers

import (
	"MyGarm/controllers"
	_ "MyGarm/docs"
	"MyGarm/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My Garm API
// @description This is a sample server for a book store.
// @version 1.0.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		// Register
		userRouter.POST("/register", controllers.UserRegister)
		// Login
		userRouter.POST("/login", controllers.UserLogin)
	}

	socialMediaRouter := r.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		// Create
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		// Get All
		socialMediaRouter.GET("/", controllers.GetAllSocialMedias)
		// Get One
		socialMediaRouter.GET("/:socialMediaId", controllers.GetSocialMedia)
		// Update
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		// Delete
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		// Create
		photoRouter.POST("/", controllers.CreatePhoto)
		// Get All
		photoRouter.GET("/", controllers.GetAllPhotos)
		// Get One
		photoRouter.GET("/:photoId", controllers.GetPhoto)
		// Update
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		// Delete
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		// Create
		commentRouter.POST("/", controllers.CreateComment)
		// Get All
		commentRouter.GET("/", controllers.GetAllComments)
		// Get One
		commentRouter.GET("/:commentId", controllers.GetComment)
		// Update
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		// Delete
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
