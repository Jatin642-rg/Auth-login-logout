package middleware

import (
	"Desktop/login-logout-task/storage"
	"Desktop/login-logout-task/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	authHeader := c.GetHeader("Authorization") // Mechanism of sending token along with a request from client to service
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		c.Abort()
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.Parse_Token(tokenStr)
	if err != nil || storage.IsTokenRevoked(tokenStr) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"}) //Should check for expiry and Invalid token
		c.Abort()
		return
	}

	email, ok := (*claims)["email"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token structure"})
		c.Abort()
		return
	}

	c.Set("email", email)
	c.Set("token", tokenStr)
	c.Next()
}
