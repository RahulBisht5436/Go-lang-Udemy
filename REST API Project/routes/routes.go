package routes

import "github.com/gin-gonic/gin"

func ManageRoutes(Server *gin.Engine) {
	Server.GET("/events", getEvents)
	Server.POST("/events", createEvent)
	Server.GET("/event/:id", getEventById)
	Server.DELETE("/event/:id", deleteEventByIdFunction)
}
