package middlewares

import (
	"net/http"
	"v3/utils"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {
	// Get token from the header
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, role, expPeriod, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if utils.IsAuthenticationLevelExpired(expPeriod) {
		// Access token expired -> Check valid refresh token
		// If it is invalid or expired return error base on them
	}

	c.Set("userId", userId)
	c.Set("role", role)
	c.Next()
}
