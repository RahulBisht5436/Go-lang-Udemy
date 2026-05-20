package routes

import (
	"errors"
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid signup payload",
			"error":   err.Error(),
		})
		return
	}

	if err := db.SaveUser(&user); err != nil {
		if errors.Is(err, db.ErrUserEmailExists) {
			context.JSON(http.StatusConflict, gin.H{
				"message": "Email already registered",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create user",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"id":      user.ID,
		"email":   user.Email,
	})
}

func login(context *gin.Context) {
	var creds models.LoginRequest
	if err := context.ShouldBindJSON(&creds); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and password are required",
			"error":   err.Error(),
		})
		return
	}

	user, err := db.AuthenticateUser(creds.Email, creds.Password)
	if err != nil {
		if errors.Is(err, db.ErrInvalidCredentials) {
			// Same response for "no such email" and "wrong password"
			// so attackers can't probe which emails exist.
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid email or password",
			})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Login failed",
			"error":   err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Note able to generate the token",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"id":      user.ID,
		"email":   user.Email,
		"token":   token,
	})
}
