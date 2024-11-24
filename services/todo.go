package services

import (
	"errors"
	"go-feToDo/database"
	dto "go-feToDo/dtos"
	"go-feToDo/models"
	"go-feToDo/utils"

	"gorm.io/gorm"
)

// Get All ToDos of a user
func GetUserAllToDo(authorID string) ([]*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID from string to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	var todos []models.Todo

	// Query to fetch all to-dos belonging to the specified author
	err = db.Unscoped().Where("author_id = ?", authorIDUint).Find(&todos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dto.ErrToDoNotFound
		}
		return nil, err
	}

	// Convert todos to DTOs
	var todoDTOs []*dto.TodoResponseDTO
	for _, todo := range todos {
		todoDTOs = append(todoDTOs, utils.ToTodoResponseDTO(&todo))
	}

	return todoDTOs, nil
}

// Get Active ToDos of a user
func GetUserActiveToDo(authorID string) ([]*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID from string to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	var todos []models.Todo

	// Query to fetch all to-dos belonging to the specified author
	err = db.Where("author_id = ? AND deleted_at IS NULL", authorIDUint).Find(&todos).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dto.ErrToDoNotFound
		}
		return nil, err
	}

	// Convert todos to DTOs
	var todoDTOs []*dto.TodoResponseDTO
	for _, todo := range todos {
		todoDTOs = append(todoDTOs, utils.ToTodoResponseDTO(&todo))
	}

	return todoDTOs, nil
}

// GetTodoById retrieves a todo by its ID, ensuring it belongs to the author.
func GetTodoById(todoID string, authorID string) (*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}
	todoIDUint, err := utils.ConvId(todoID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	var todo models.Todo
	err = db.Where("id = ?", todoIDUint).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, dto.ErrToDoNotFound
	}
	if err != nil {
		return nil, err
	}

	// Check if the author owns the todo
	if todo.AuthorID != authorIDUint {
		return nil, dto.ErrUnauthToDo
	}

	return utils.ToTodoResponseDTO(&todo), nil
}

// GetTrashedTodos retrieves all soft-deleted todos for the author.
func GetTrashedTodos(authorID string) ([]*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	var todos []models.Todo
	if err := db.Unscoped().Where("author_id = ? AND deleted_at IS NOT NULL", authorIDUint).Find(&todos).Error; err != nil {
		return nil, dto.ErrToDoTrash
	}

	// Map to DTOs
	var todoDTOs []*dto.TodoResponseDTO
	for _, todo := range todos {
		todoDTOs = append(todoDTOs, utils.ToTodoResponseDTO(&todo))
	}

	return todoDTOs, nil
}

// CreateTodo adds a new todo with a unique title for each user.
func CreateTodo(todoDTO *dto.CreateTodoDTO, authorID string) (*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	// Check if a todo with the same title already exists for the user
	var existingTodo models.Todo
	err = db.Where("title = ? AND author_id = ?", todoDTO.Title, authorIDUint).First(&existingTodo).Error
	if err == nil {
		// If an existing to-do is found, return an error indicating the title is not unique
		return nil, dto.ErrToDoTitleAlreadyExists
	}
	if err != gorm.ErrRecordNotFound {
		// Handle unexpected errors
		return nil, err
	}

	// Proceed with creating the new to-do
	todo := &models.Todo{
		Title:       todoDTO.Title,
		Description: todoDTO.Description,
		AuthorID:    authorIDUint,
	}

	if err := db.Create(todo).Error; err != nil {
		return nil, dto.ErrToDoCreate
	}

	return utils.ToTodoResponseDTO(todo), nil
}

// UpdateTodo updates an existing todo, ensuring it belongs to the author.
func UpdateTodo(todoID string, updateDTO *dto.UpdateTodoDTO, authorID string) (*dto.TodoResponseDTO, error) {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}
	todoIDUint, err := utils.ConvId(todoID)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	var todo models.Todo
	err = db.Where("id = ?", todoIDUint).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, dto.ErrToDoNotFound
	}
	if err != nil {
		return nil, err
	}

	// Check if the author owns the todo
	if todo.AuthorID != authorIDUint {
		return nil, dto.ErrUnauthToDo
	}

	// Update fields if provided
	if updateDTO.Title != nil {
		todo.Title = *updateDTO.Title
	}
	if updateDTO.Description != nil {
		todo.Description = *updateDTO.Description
	}
	if updateDTO.Status != nil {
		todo.Status = *updateDTO.Status
	}

	if err := db.Save(&todo).Error; err != nil {
		return nil, dto.ErrToDoUpdate
	}

	return utils.ToTodoResponseDTO(&todo), nil
}

// SoftDeleteTodo marks a todo as deleted without permanently removing it.
func SoftDeleteTodo(todoID string, authorID string) error {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return dto.ErrAuthIdConv
	}
	todoIDUint, err := utils.ConvId(todoID)
	if err != nil {
		return dto.ErrAuthIdConv
	}

	var todo models.Todo
	err = db.Where("id = ?", todoIDUint).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.ErrToDoNotFound
	}
	if err != nil {
		return err
	}

	// Check if the author owns the todo
	if todo.AuthorID != authorIDUint {
		return dto.ErrUnauthToDo
	}

	return db.Delete(&todo).Error
}

// DeleteTodo permanently deletes a todo from the database.
func DeleteTodo(todoID string, authorID string) error {
	db := database.GetDB()

	// Convert authorID to uint
	authorIDUint, err := utils.ConvId(authorID)
	if err != nil {
		return dto.ErrAuthIdConv
	}
	todoIDUint, err := utils.ConvId(todoID)
	if err != nil {
		return dto.ErrAuthIdConv
	}

	var todo models.Todo
	err = db.Unscoped().Where("id = ?", todoIDUint).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dto.ErrToDoNotFound
	}
	if err != nil {
		return err
	}

	// Check if the author owns the todo
	if todo.AuthorID != authorIDUint {
		return dto.ErrUnauthToDo
	}

	return db.Unscoped().Delete(&todo).Error
}
