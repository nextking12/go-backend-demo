package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *CreateUserRequest) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (u *UpdateUserRequest) Validate() error {
	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Email == "" {
		return ErrEmailRequired
	}
	return nil
}
