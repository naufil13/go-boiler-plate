package main

import (
	"go-boiler-plate/initializers"
	"go-boiler-plate/models"
)

func init() {
	initializers.LoadVarEnv()
	initializers.ConnectToDatabase()
}

func main() {
	Post()
	User()
}

func Post() {
	initializers.DB.AutoMigrate(&models.Post{})
}

func User() {
	initializers.DB.AutoMigrate(&models.User{})
}
