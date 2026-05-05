package models

import (
	"gorm.io/datatypes"
)

type User struct {
	Login string
	Email string
	PasswordHash string
	IsVerified bool
	Role string
	Created_at datatypes.Date
}


type UserCreateRequest struct {
	Login string
	Email string
	Password string
} 