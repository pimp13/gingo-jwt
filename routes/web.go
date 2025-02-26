package routes

import (
	"Gin-GORM-Project/controllers/user_controller"
	"Gin-GORM-Project/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter(app *gin.Engine) {
	route := app

	err := route.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatalf("Proxy unauthorized: %s", err)
		return
	}

	// Routes
	//route.GET("/", user_controller.GetAllUsers)

	userRoutes := route.Group("/users")
	{
		userRoutes.GET("", user_controller.GetAllUsers)
		userRoutes.POST("", user_controller.CreateUser)
		userRoutes.POST("/register", user_controller.Register)
		userRoutes.POST("/login", user_controller.Login)
		userRoutes.GET("/validate", middleware.RequireAuth, user_controller.Validate)
		userRoutes.POST("/logout", user_controller.Logout)
		//userRoutes.GET("/:id", user_controller.GetUser)
		//userRoutes.PUT("/:id", user_controller.UpdateUser)
		//userRoutes.DELETE("/:id", user_controller.DeleteUser)
	}
}
