package controllers

import (
	dto "go-feToDo/dtos"
	"go-feToDo/services"
	"go-feToDo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserById handles GET requests to fetch the authenticated user's information
func GetUserById(c *gin.Context) {
	userID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	user, err := services.GetUserById(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validate.Struct(userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": utils.ParseValidationErrors(err)})
		return
	}

	user, err := services.CreateUser(&userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT requests to update the authenticated user's information
func UpdateUser(c *gin.Context) {
	userID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	var userDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := validate.Struct(userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": utils.ParseValidationErrors(err)})
		return
	}

	user, err := services.UpdateUser(userID.(string), &userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE requests to delete the authenticated user
func DeleteUser(c *gin.Context) {
	userID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	err := services.DeleteUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
