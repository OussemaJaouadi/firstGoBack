package utils

import (
	dto "go-feToDo/dtos"
	"go-feToDo/models"
	"strconv"
)

// ToUserResponseDTO converts a User model to a UserResponseDTO
func ToUserResponseDTO(user *models.User) *dto.UserResponseDTO {
	return &dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToTodoResponseDTO converts a Todo model to a TodoResponseDTO
func ToTodoResponseDTO(todo *models.Todo) *dto.TodoResponseDTO {
	return &dto.TodoResponseDTO{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Status:      todo.Status,
		AuthorID:    todo.AuthorID,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}
}

// ConvId converts a string to a uint and returns an error if it fails
func ConvId(authorId string) (uint, error) {
	parsed, err := strconv.ParseUint(authorId, 10, 32)
	if err != nil {
		return 0, dto.ErrAuthIdConv
	}
	return uint(parsed), nil
}
