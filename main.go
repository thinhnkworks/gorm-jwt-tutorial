package main

import (
	"gin-jwt-tutorial/controllers"
	"gin-jwt-tutorial/initializers"
	"gin-jwt-tutorial/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectPostgres()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
