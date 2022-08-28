package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenduc2906/go-rest-api/config"
	"github.com/nguyenduc2906/go-rest-api/controller"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	router.Run()
}
