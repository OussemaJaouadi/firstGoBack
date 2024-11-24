package models

import (
	dto "go-feToDo/dtos"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"` // "-" ensures this field isn't exposed in JSON
	Todos     []Todo    `json:"todos" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

// BeforeCreate GORM hook to set timestamps and hash password.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	// Hash the password before saving
	return u.HashPassword(u.Password)
}

// BeforeUpdate GORM hook to update the timestamp.
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}

// HashPassword hashes the user's password.
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares a plain-text password with the hashed password.
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// Differentiating between a hash mismatch and other potential bcrypt errors
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return dto.ErrJWTInvalidCreds
		}
		return dto.ErrBycFail
	}
	return nil
}
