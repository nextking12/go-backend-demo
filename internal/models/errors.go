package models

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrNameRequired  = errors.New("name is required")
	ErrEmailRequired = errors.New("email is required")
	ErrInvalidUserID = errors.New("invalid user ID")
)
