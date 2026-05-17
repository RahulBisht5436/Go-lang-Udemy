package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/rest-api/models"
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

	Server.GET("/events", getEvents)
	Server.POST("/events", createEvent)
	fmt.Printf("Server Started on the PORT : %v \n", port)
	Server.Run(":" + port)
}

func createEvent(context *gin.Context) {
	var event models.Event

	//here we bind the entered Body into the event variable
	err := context.BindJSON(&event)

	//if any error in the binding , Handle the error
	if err != nil {
		fmt.Println("error in parsing Data")
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid param passed",
			"errror":  err.Error(),
		})
		return
	}
	// Creating a validation check for the entries
	if event.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Name is required Field",
		})
		return
	} else if event.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "ID is required Field",
		})
		return
	} else if event.UserId == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "UserId is required Field",
		})
		return
	}

	event.Save()
	//Here we are sending the response
	context.JSON(http.StatusOK, gin.H{
		"Message":     "Data Saved successFully",
		"Name":        event.Name,
		"Description": event.Description,
		"Location":    event.Location,
		"ID":          event.ID,
		"UserId":      event.UserId,
	})

}
func getEvents(context *gin.Context) {
	allEvents := models.GetAllEvent()
	context.JSON(http.StatusOK, gin.H{
		"events": allEvents,
	})
}
