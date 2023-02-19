package main

import (
	"github.com/bolatbek1997/tokenGeneration/controller"
	"github.com/bolatbek1997/tokenGeneration/middlewares"
	"github.com/bolatbek1997/tokenGeneration/models"
	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
