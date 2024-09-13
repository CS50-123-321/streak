package main

import (
	"familyFormUi/be"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	be.InitRedis()
	be.InitRoutes()
}
