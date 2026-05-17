package main

import (
	"fmt"
	"os"

	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	///Loading the env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Not able to load the Env Variables")
		return
	}

	// Get a specifiv Env from the loaded Env
	port := os.Getenv("PORT")

	// Behind the screen it setup a http server
	Server := gin.Default()
	db.InitDB()
	//Server is Passed to the manage Routes
	routes.ManageRoutes(Server)

	fmt.Printf("Server Started on the PORT : %v \n", port)
	Server.Run(":" + port)
}
