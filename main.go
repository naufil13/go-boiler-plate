package main

import (
	"go-boiler-plate/initializers"
	"go-boiler-plate/routes"
)

func init() {
	initializers.LoadVarEnv()
	initializers.ConnectToDatabase()
	initializers.ConnectToRedis()
}

func main() {
	routes.Webroutes()
}
