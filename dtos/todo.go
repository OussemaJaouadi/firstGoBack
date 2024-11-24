package dto

import (
	"go-feToDo/enums"
	"time"
)

// create todo payload.
type CreateTodoDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description,omitempty"`
}

// update todo payload.
type UpdateTodoDTO struct {
	Title       *string           `json:"title,omitempty" validate:"omitempty"`
	Description *string           `json:"description,omitempty"`
	Status      *enums.TodoStatus `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed"`
}

// response structure for a todo item.
type TodoResponseDTO struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description,omitempty"`
	Status      enums.TodoStatus `json:"status"`
	AuthorID    uint             `json:"author_id"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}
