package middleware

import (
	"errors"
	dto "go-feToDo/dtos"
	"go-feToDo/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": dto.ErrJWTMissingAuthHeader.Error(),
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": dto.ErrJWTInvalidFormat.Error(),
			})
			c.Abort()
			return
		}

		// Decode the token
		tokenString := parts[1]
		payload, err := utils.DecodeToken(tokenString, false)
		if err != nil {
			if errors.Is(err, dto.ErrJWTUnexpectedSigningMethod) {
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			} else if errors.Is(err, dto.ErrJWTInvalidToken) || errors.Is(err, dto.ErrJWTDecodeError) {
				c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
			c.Abort()
			return
		}

		// Check if the token is expired
		if time.Now().Unix() > payload.Expires {
			c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTExpiredToken.Error()})
			c.Abort()
			return
		}

		// Token is valid, proceed with the request
		c.Set("username", payload.Username)
		c.Set("authorID", payload.Id)
		c.Next()
	}
}
