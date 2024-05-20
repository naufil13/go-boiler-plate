package main

import (
	"go-first-project/initializers"
	"go-first-project/routes"
)

func init() {
	initializers.LoadVarEnv()
	initializers.ConnectToDatabase()
	initializers.ConnectToRedis()
}

func main() {
	routes.Webroutes()
}
