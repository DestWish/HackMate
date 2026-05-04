package models

type User struct {
	login string
	email string
	passwordHash string
	isVerified bool
	status string
	created_at [5]int
}