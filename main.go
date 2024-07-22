package main

import (
	"golang-jwt/controllers"
	"golang-jwt/initializers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.GetEnvVariables()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/employees", controllers.GetEmployees)
	r.GET("/employee/:id", controllers.GetEmployeeById)
	r.POST("/employee", controllers.PostEmployee)
	r.PUT("/employee/:id", controllers.UpdateEmployeeById)
	r.DELETE("/employee/:id", controllers.DeleteEmployeeById)

	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET, POST, PUT, DELETE, OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "X-HTTP-Method-Override", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	})
}
