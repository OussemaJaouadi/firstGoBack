package dto

// payload for login.
type LoginRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// structure of the JWT payload.
type JWTPayloadDTO struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Expires  int64  `json:"expires"`
}

// response after login.
type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// refresh token body
type RequestBody struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
