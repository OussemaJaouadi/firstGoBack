package dto

import "time"

// CreateUserDTO represents the payload for creating a user.
type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserDTO represents the payload for updating a user.
type UpdateUserDTO struct {
	Username *string `json:"username,omitempty" validate:"omitempty,min=3,max=32"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
}

// UserResponseDTO represents the response structure for a user.
type UserResponseDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
