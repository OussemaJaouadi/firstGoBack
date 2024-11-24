package services

import (
	"go-feToDo/database"
	dto "go-feToDo/dtos"
	"go-feToDo/models"
	"go-feToDo/utils"
)

func Login(loginData *dto.LoginRequestDTO) (*dto.LoginResponseDTO, error) {
	var user models.User
	db := database.GetDB()

	if err := db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		return nil, dto.ErrUserNotFound
	}
	err := user.CheckPassword(loginData.Password)
	if err != nil {
		return nil, err
	}
	accessToken, err := utils.CreateToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	refreshToken, err := utils.CreateRefreshToken(user.ID, user.Username)
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
