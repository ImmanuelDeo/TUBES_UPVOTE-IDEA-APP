package model

const MaxUsers = 100

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}