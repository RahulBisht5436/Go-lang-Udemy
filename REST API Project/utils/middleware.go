package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context keys used to stash authenticated-user info on the gin.Context.
// Centralised as constants so handlers and middleware can't disagree on
// the spelling.
const (
	ContextEmailKey  = "authEmail"
	ContextUserIDKey = "authUserId"
)

// AuthMiddleware validates the JWT in the Authorization header and, on
// success, stores the authenticated user's email and id on the gin.Context
// (under ContextEmailKey and ContextUserIDKey) so downstream handlers can
// pull them out with context.GetString / context.GetInt64.
//
// On any failure — missing header, malformed token, wrong signature,
// expired token — it aborts the request with HTTP 401 and the chain
// stops; the handler is never called.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		tokenString, err := ExtractToken(raw)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication required",
				"error":   err.Error(),
			})
			return
		}

		email, userId, err := ExtractUserInfo(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid or expired token",
				"error":   err.Error(),
			})
			return
		}

		c.Set(ContextEmailKey, email)
		c.Set(ContextUserIDKey, userId)
		c.Next()
	}
}
