package controllers

import (
	"fmt"
	dto "go-feToDo/dtos"
	"go-feToDo/services"
	"go-feToDo/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Login handles user login by validating credentials and generating JWT tokens
func Login(c *gin.Context) {
	var loginRequest dto.LoginRequestDTO

	// Parse the login request
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrInvalidReqPayload.Error()})
		return
	}
	if err := validate.Struct(loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": utils.ParseValidationErrors(err)})
		return
	}

	// Authenticate the user and generate tokens
	loginResponse, err := services.Login(&loginRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the login response with the tokens
	c.JSON(http.StatusOK, loginResponse)
}

// Refresh handles refreshing the access token using the refresh token
func Refresh(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": dto.ErrJWTMissingAuthHeader.Error(),
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": dto.ErrJWTInvalidFormat.Error(),
		})
		return
	}

	accessToken := parts[1]

	// Parse the refresh token from the request body
	var requestBody dto.RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	refreshToken := requestBody.RefreshToken

	// Decode both access and refresh tokens
	accessPayload, accessErr := utils.DecodeToken(accessToken, false)

	if accessErr != nil {
		fmt.Print(accessErr)
	} else {
		c.JSON(http.StatusOK, &dto.LoginResponseDTO{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
		return
	}

	refreshPayload, refreshErr := utils.DecodeToken(refreshToken, true)

	// Case 2: Access token is expired, validate refresh token and issue a new access token
	if refreshErr == nil {
		// Validate the payloads of both tokens
		if accessPayload.Id != refreshPayload.Id || accessPayload.Username != refreshPayload.Username {
			c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTTokenMismatch.Error()})
			return
		}

		// Convert the ID from string (as per your DTO) to uint
		userId, err := utils.ConvId(refreshPayload.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dto.ErrAuthIdConv.Error()})
			return
		}

		// Generate a new access token using the refresh payload
		newAccessToken, err := utils.CreateToken(userId, refreshPayload.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
			return
		}

		// Respond with the new access token and the original refresh token
		c.JSON(http.StatusOK, &dto.LoginResponseDTO{
			AccessToken:  newAccessToken,
			RefreshToken: refreshToken,
		})
		return
	}
	// Case 3: Both tokens are invalid or expired
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Both tokens Invalid or expired tokens"})
}
