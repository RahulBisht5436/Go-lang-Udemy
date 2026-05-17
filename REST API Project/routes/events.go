package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

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
	err = db.InsertData(event.Name, event.Description, event.Location, event.DateTime.Format(time.RFC3339), event.UserId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to Save Data",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"Message": "Data Saved successFully",
		"event":   event,
	})
}
func getEventById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID entered",
		})
		return
	}

	result, err := db.GetIdDbEvents(id)
	if err != nil {
		if errors.Is(err, db.ErrEventNotFound) {
			context.JSON(http.StatusNotFound, gin.H{
				"message": "Event not found",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Query Operation Failed",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Data Fetch SuccessFul",
		"result":  result,
	})
}
func getEvents(context *gin.Context) {
	allEvents, err := db.GetAllDbEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to fetch events",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"events": allEvents,
	})
}

func deleteEventByIdFunction(context *gin.Context) {

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID entered",
		})
		return
	}

	errDatabase := db.DeleteEventById(id)
	if errDatabase != nil {

		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Not Able to delete the Entry",
		})

	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event deleted Successfully",
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil || id <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID entered",
		})
		return
	}
	var event models.Event
	errEvent := context.BindJSON(&event)
	if errEvent != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to Update Information",
		})
		return
	}

	errUpdate := db.UpdateEvent(id, event)
	if errUpdate != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Data was not Updated",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event Updated Successfully",
		"Event":   event,
	})
}
