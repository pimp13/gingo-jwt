package bootstrap

import (
	"Gin-GORM-Project/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func App() {
	app := gin.Default()

	// initialize router
	routes.InitRouter(app)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	// start server
	if err := app.Run(":" + PORT); err != nil {
		log.Fatalf("ERROR: server cannot connect - %s", err)
		return
	}
}
