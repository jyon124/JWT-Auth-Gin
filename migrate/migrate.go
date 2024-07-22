package main

import (
	"golang-jwt/initializers"
	"golang-jwt/models"
)

func init() {
	initializers.GetEnvVariables()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{})
}
