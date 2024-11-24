package utils

import (
	"fmt"
	"time"

	"go-feToDo/config"
	dto "go-feToDo/dtos"

	"github.com/golang-jwt/jwt/v4"
)

var (
	cfg              = config.LoadConfig()
	jwtSecretKey     = cfg.JwtSecret
	refreshSecretKey = cfg.RefreshSecret
	tokenExpiry      = time.Duration(cfg.TokenExpiry * int(time.Second))
	refreshExpiry    = time.Duration(cfg.RefreshExpiry) * time.Second
)

// CreateToken generates a new JWT access token with customizable duration
func CreateToken(id uint, username string, duration ...time.Duration) (string, error) {
	expirationTime := time.Now().Add(getDuration(duration, tokenExpiry))

	claims := jwt.MapClaims{
		"id":       fmt.Sprint(id),
		"username": username,
		"expires":  expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecretKey))
}

// CreateRefreshToken generates a new JWT refresh token with customizable duration
func CreateRefreshToken(id uint, username string, duration ...time.Duration) (string, error) {
	expirationTime := time.Now().Add(getDuration(duration, refreshExpiry))

	claims := jwt.MapClaims{
		"id":       fmt.Sprint(id),
		"username": username,
		"expires":  expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshSecretKey))
}

// DecodeToken validates and decodes a JWT token, returning the claims
func DecodeToken(tokenString string, isRefresh bool) (*dto.JWTPayloadDTO, error) {
	secretKey := []byte(getSecret(isRefresh))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, dto.ErrJWTUnexpectedSigningMethod
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, dto.ErrJWTInvalidToken
	}

	// Extract claims
	username, ok := claims["username"].(string)
	if !ok {
		return nil, dto.ErrJWTUsernameClaimMissing
	}
	id, ok := claims["id"].(string)
	if !ok {
		return nil, dto.ErrJWTIDClaimMissing
	}
	expires, ok := claims["expires"].(float64)
	if !ok {
		return nil, dto.ErrJWTExpiresClaimMissing
	}

	return &dto.JWTPayloadDTO{
		Id:       id,
		Username: username,
		Expires:  int64(expires),
	}, nil
}

// RefreshToken generates a new access token using a valid refresh token
func RefreshToken(refreshToken string) (string, error) {
	payload, err := DecodeToken(refreshToken, true)
	if err != nil {
		return "", err
	}

	// Check if the refresh token has expired
	if time.Now().Unix() > payload.Expires {
		return "", dto.ErrJWTExpiredRefresh
	}

	// Convert ID to uint for token creation
	idUint, err := ConvId(payload.Id)
	if err != nil {
		return "", err
	}

	return CreateToken(idUint, payload.Username)
}

// getDuration returns a specified duration or the default value if not provided
func getDuration(durations []time.Duration, defaultDuration time.Duration) time.Duration {
	if len(durations) > 0 {
		return durations[0]
	}
	return defaultDuration
}

// getSecret determines whether to use the access or refresh secret key
func getSecret(isRefresh bool) string {
	if isRefresh {
		return refreshSecretKey
	}
	return jwtSecretKey
}
