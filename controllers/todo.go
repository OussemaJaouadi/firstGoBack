package controllers

import (
	dto "go-feToDo/dtos"
	"go-feToDo/services"
	"go-feToDo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllToDo handles GET requests to fetch all todos for a user
func GetAllToDo(c *gin.Context) {
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todos, err := services.GetUserAllToDo(authorID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// GetActiveToDo handles GET requests to fetch active todos for a user
func GetActiveToDo(c *gin.Context) {
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todos, err := services.GetUserActiveToDo(authorID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func GetTrashToDo(c *gin.Context) {
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todos, err := services.GetTrashedTodos(authorID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// GetTodoById handles GET requests to fetch a specific todo by ID for a user
func GetTodoById(c *gin.Context) {
	todoID := c.Param("todoID")
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todo, err := services.GetTodoById(todoID, authorID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// CreateTodo handles POST requests to create a new todo
func CreateTodo(c *gin.Context) {
	var todoDTO dto.CreateTodoDTO
	if err := c.ShouldBindJSON(&todoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrInvalidReqPayload.Error()})
		return
	}
	if err := validate.Struct(todoDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": utils.ParseValidationErrors(err)})
		return
	}

	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todo, err := services.CreateTodo(&todoDTO, authorID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo handles PUT requests to update an existing todo
func UpdateTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	var updateDTO dto.UpdateTodoDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": dto.ErrInvalidReqPayload.Error()})
		return
	}
	if err := validate.Struct(updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_errors": utils.ParseValidationErrors(err)})
		return
	}

	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	todo, err := services.UpdateTodo(todoID, &updateDTO, authorID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// SoftDeleteTodo handles DELETE requests to soft delete a todo
func SoftDeleteTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	err := services.SoftDeleteTodo(todoID, authorID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// DeleteTodo handles DELETE requests to permanently delete a todo
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoID")
	authorID, exists := c.Get("authorID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": dto.ErrJWTUnauthorizedAccess.Error()})
		return
	}

	err := services.DeleteTodo(todoID, authorID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
