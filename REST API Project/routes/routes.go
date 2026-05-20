package routes

import (
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func ManageRoutes(Server *gin.Engine) {
	// Public, read-only routes — anyone can browse events.
	Server.GET("/events", getEvents)
	Server.GET("/event/:id", getEventById)

	// Protected routes. AuthMiddleware runs first; if the JWT is
	// missing/invalid the request is short-circuited with 401 and the
	// handler is never called. Inside each handler we can rely on
	// authedUser() returning a real user.
	authed := Server.Group("/")
	authed.Use(utils.AuthMiddleware())
	{
		authed.POST("/events", createEvent)
		authed.PUT("/event/:id", updateEvent)
		authed.DELETE("/event/:id", deleteEventByIdFunction)
	}

	// user routes (auth endpoints themselves are public)
	Server.POST("/signup", signup)
	Server.POST("/login", login)
}
