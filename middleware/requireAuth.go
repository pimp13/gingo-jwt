package middleware

import (
	"Gin-GORM-Project/database"
	"Gin-GORM-Project/helpers"
	"Gin-GORM-Project/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie of the request
	tokenCookie, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Missing tokenParsed",
		})
		return
	}

	// Parse the tokenParsed from the cookie value
	tokenParsed, err := helpers.ValidateJWT(tokenCookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is invalid",
		})
		return
	}
	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		// Check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Time out",
			})
			return
		}

		// Find the user auth with tokenParsed sub
		var user models.User
		database.DB.First(&user, "email = ?", claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User not found",
			})
			return
		}

		// Attach the user to the context
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is invalid",
		})
		return
	}
}
