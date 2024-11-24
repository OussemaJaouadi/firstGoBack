package models

import (
	"go-feToDo/enums"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint             `json:"id" gorm:"primaryKey"`
	Title       string           `json:"title" gorm:"not null"`
	Description string           `json:"description"`
	Status      enums.TodoStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
	AuthorID    uint             `json:"author_id" gorm:"not null"`
	Author      User             `json:"author" gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"-" gorm:"index"`
}

// TableName specifies the table name for the Todo model
func (Todo) TableName() string {
	return "todos"
}

// BeforeCreate hook to handle timestamps and default status
func (t *Todo) BeforeCreate(tx *gorm.DB) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	if t.Status == "" {
		t.Status = enums.TodoStatusInProgress
	}
	return nil
}

// BeforeUpdate hook to handle timestamps
func (t *Todo) BeforeUpdate(tx *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}
