package core

import (
	"encoding/json"
	"fmt"
	"os"
	"Tubes-IdeaManagerCLI/model"
)

const userFilePath = "data/users.json"

func LoadUsers() ([]model.User, error) {
	file, err := os.Open(userFilePath)
	if err != nil {
		return []model.User{}, nil
	}
	defer file.Close()

	var users []model.User
	err = json.NewDecoder(file).Decode(&users)
	return users, err
}

func SaveUsers(users []model.User) error {
	file, err := os.Create(userFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(users)
}

func RegisterUser(username, password string) error {
	users, err := LoadUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Username == username {
			return fmt.Errorf("username '%s' sudah digunakan", username)
		}
	}

	users = append(users, model.User{Username: username, Password: password})
	return SaveUsers(users)
}

func LoginUser(username, password string) bool {
	users, err := LoadUsers()
	if err != nil {
		fmt.Println("Gagal membaca data user.")
		return false
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}