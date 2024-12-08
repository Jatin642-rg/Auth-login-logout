package handlers

import (
	"Desktop/login-logout-task/storage"
	"Desktop/login-logout-task/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) { //Authorization of token
	email := c.GetString("email")
	c.JSON(http.StatusOK, gin.H{"message": "Token is valid", "email": email})
}

func RevokeToken(c *gin.Context) { //Revocation of token
	token := c.GetString("token")
	storage.RevokeToken(token) //Mechanism of revoking a token from backend
	c.JSON(http.StatusOK, gin.H{"message": "Token revoked"})
}

func RefreshToken(c *gin.Context) { //Mechanism to refresh a token
	oldToken := c.GetString("token")
	email := c.GetString("email")

	newToken, err := utils.Generate_Token(email) //Client should be able to renew the token before it expires
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	storage.RevokeToken(oldToken)
	c.JSON(http.StatusOK, gin.H{"new_token": newToken})
}
