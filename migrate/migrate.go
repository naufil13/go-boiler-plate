package main

import (
	"go-first-project/initializers"
	"go-first-project/models"
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
