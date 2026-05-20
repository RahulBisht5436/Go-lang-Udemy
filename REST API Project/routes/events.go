package routes

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

// authedUser is a tiny helper that reads the values stashed on the gin
// context by utils.AuthMiddleware. Centralising the lookup means handlers
// don't have to know the magic string keys.
func authedUser(c *gin.Context) (email string, userId int64) {
	email = c.GetString(utils.ContextEmailKey)
	userId = c.GetInt64(utils.ContextUserIDKey)
	return
}

func createEvent(context *gin.Context) {
	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event payload",
			"error":   err.Error(),
		})
		return
	}

	// Owner fields come from the JWT, NEVER from the request body —
	// otherwise any logged-in user could create events on someone else's
	// behalf just by lying about userId in the JSON.
	email, userId := authedUser(context)
	event.UserId = userId
	event.UserEmail = email

	if err := db.InsertData(&event); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to Save Data",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Data Saved successFully",
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

	_, userId := authedUser(context)

	if err := db.DeleteEventById(id, userId); err != nil {
		respondOwnershipError(context, err, "Not Able to delete the Entry")
		return
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
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to Update Information",
			"error":   err.Error(),
		})
		return
	}

	_, userId := authedUser(context)

	if err := db.UpdateEvent(id, event, userId); err != nil {
		respondOwnershipError(context, err, "Data was not Updated")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event Updated Successfully",
		"Event":   event,
	})
}

// respondOwnershipError maps the small set of domain errors that update
// and delete can return into appropriate HTTP responses:
//   - ErrEventNotFound → 404
//   - ErrForbidden     → 403  (authenticated but not the owner)
//   - anything else    → 500  (with the fallback message)
func respondOwnershipError(context *gin.Context, err error, fallback string) {
	switch {
	case errors.Is(err, db.ErrEventNotFound):
		context.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
	case errors.Is(err, db.ErrForbidden):
		context.JSON(http.StatusForbidden, gin.H{
			"message": "You are not the owner of this event",
		})
	default:
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": fallback,
			"error":   err.Error(),
		})
	}
}
