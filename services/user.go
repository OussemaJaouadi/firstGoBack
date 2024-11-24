package services

import (
	"errors"
	"go-feToDo/database"
	dto "go-feToDo/dtos"
	"go-feToDo/models"
	"go-feToDo/utils"

	"gorm.io/gorm"
)

// GetUserById fetches a user by their ID and returns the UserResponseDTO
func GetUserById(id string) (*dto.UserResponseDTO, error) {
	var user models.User
	db := database.GetDB()
	iDUint, err := utils.ConvId(id)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}
	if err := db.First(&user, iDUint).Error; err != nil {
		return nil, dto.ErrUserNotFound
	}

	userDTO := utils.ToUserResponseDTO(&user)
	return userDTO, nil
}

// GetUserByUsername fetches a user by their username and returns the UserResponseDTO
func GetUserByUsername(username string) (*dto.UserResponseDTO, error) {
	var user models.User
	db := database.GetDB()

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, dto.ErrUserNotFound
	}

	userDTO := utils.ToUserResponseDTO(&user)

	return userDTO, nil
}

// GetUserByEmail fetches a user by their email and returns the UserResponseDTO
func GetUserByEmail(email string) (*dto.UserResponseDTO, error) {
	var user models.User
	db := database.GetDB()

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, dto.ErrUserNotFound
	}

	userDTO := utils.ToUserResponseDTO(&user)

	return userDTO, nil
}

// CreateUser creates a new user with the provided data and returns the UserResponseDTO
func CreateUser(userADD *dto.CreateUserDTO) (*dto.UserResponseDTO, error) {
	db := database.GetDB()

	// Check if email already exists
	var userExist models.User
	if err := db.Where("email = ?", userADD.Email).First(&userExist).Error; err == nil {
		return nil, dto.ErrEmailAlreadyExists // Email already exists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // Unexpected error while checking email
	}

	// Check if username already exists
	if err := db.Where("username = ?", userADD.Username).First(&userExist).Error; err == nil {
		return nil, dto.ErrUsernameAlreadyExists // Username already exists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err // Unexpected error while checking username
	}

	// Proceed with user creation
	user := models.User{
		Username: userADD.Username,
		Email:    userADD.Email,
		Password: userADD.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, dto.ErrUserCreate // Error creating user
	}

	// Convert to DTO for response
	userDTO := utils.ToUserResponseDTO(&user)

	return userDTO, nil
}

// DeleteUser deletes a user by their ID
func DeleteUser(id string) error {
	db := database.GetDB()
	iDUint, err := utils.ConvId(id)
	if err != nil {
		return dto.ErrAuthIdConv
	}
	if err := db.Delete(&models.User{}, iDUint).Error; err != nil {
		return dto.ErrUserDelete
	}

	return nil
}

// UpdateUser updates a user’s details and returns the UserResponseDTO
func UpdateUser(id string, userUpdate *dto.UpdateUserDTO) (*dto.UserResponseDTO, error) {
	var user models.User
	db := database.GetDB()
	iDUint, err := utils.ConvId(id)
	if err != nil {
		return nil, dto.ErrAuthIdConv
	}

	// Find the user by ID
	if err := db.First(&user, iDUint).Error; err != nil {
		return nil, dto.ErrUserNotFound
	}

	// Validate if email or username already exists before updating
	if userUpdate.Email != nil {
		var existingUserByEmail models.User
		if err := db.Where("email = ?", *userUpdate.Email).First(&existingUserByEmail).Error; err == nil {
			// Email already exists
			return nil, dto.ErrEmailAlreadyExists
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Email = *userUpdate.Email
	}

	if userUpdate.Username != nil {
		var existingUserByUsername models.User
		if err := db.Where("username = ?", *userUpdate.Username).First(&existingUserByUsername).Error; err == nil {
			// Username already exists
			return nil, dto.ErrUsernameAlreadyExists
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		user.Username = *userUpdate.Username
	}

	// Save the updated user
	if err := db.Save(&user).Error; err != nil {
		return nil, dto.ErrUserUpdate
	}

	// Convert to DTO for response
	userDTO := utils.ToUserResponseDTO(&user)

	return userDTO, nil
}
