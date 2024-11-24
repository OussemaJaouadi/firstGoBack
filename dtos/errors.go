package dto

import "errors"

// JWT-related errors
var (
	ErrJWTExpiredRefresh          = errors.New("JWT token has expired and can't be refreshed")
	ErrJWTInvalidFormat           = errors.New("authorization header format is invalid")
	ErrJWTUsernameClaimMissing    = errors.New("username claim missing or invalid")
	ErrJWTExpiresClaimMissing     = errors.New("expires claim missing or invalid")
	ErrJWTMissingAuthHeader       = errors.New("authorization header missing")
	ErrJWTNeedsRefresh            = errors.New("token needs to be refreshed")
	ErrJWTIDClaimMissing          = errors.New("ID claim missing or invalid")
	ErrJWTUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrJWTDecodeError             = errors.New("error decoding JWT token")
	ErrJWTExpiredToken            = errors.New("JWT token has expired")
	ErrJWTInvalidCreds            = errors.New("invalid credentials")
	ErrJWTInvalidToken            = errors.New("invalid JWT token")
	ErrJWTTokenMismatch           = errors.New("token Missmatch")
	ErrJWTUnauthorizedAccess      = errors.New("unauthorized access")
)

// CRUD User Errors
var (
	ErrUserNotFound          = errors.New("user Not Found")
	ErrUserCreate            = errors.New(`user Creating Error`)
	ErrUserDelete            = errors.New("user Deleting Error")
	ErrUserUpdate            = errors.New("user Updating Error")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrUsernameAlreadyExists = errors.New("username already exists")
)

// CRUD ToDO Errors
var (
	ErrToDoNotFound           = errors.New("To-Do Not Found")
	ErrToDoCreate             = errors.New("To-Do Creating Error")
	ErrToDoDelete             = errors.New("To-Do Deleting Error")
	ErrToDoTrash              = errors.New("To-Do Trashing Error")
	ErrToDoUpdate             = errors.New("To-Do Updating Error")
	ErrAuthIdConv             = errors.New("failed to parse User Id")
	ErrUnauthToDo             = errors.New("unauthorized to access this ToDo")
	ErrToDoTitleAlreadyExists = errors.New("you have already a To Do with that title")
)

// other
var (
	ErrPassMiss          = errors.New("password is incorrect")
	ErrInvalidReqPayload = errors.New("invalid request payload")
	ErrBycFail           = errors.New("an error occurred while verifying the password")
)
