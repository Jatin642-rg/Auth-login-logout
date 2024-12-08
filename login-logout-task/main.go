package main

import (
	"Desktop/login-logout-task/handlers"
	"Desktop/login-logout-task/middleware"
	"Desktop/login-logout-task/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitializeStore()
	router := gin.Default()

	router.POST("/signup", handlers.SignUp)
	router.POST("/signin", handlers.SignIn)
	router.GET("/authorize", middleware.Authorize, handlers.Authorize)
	router.POST("/revoke", middleware.Authorize, handlers.RevokeToken)
	router.POST("/refresh", middleware.Authorize, handlers.RefreshToken)

	router.Run(":8080")
}
