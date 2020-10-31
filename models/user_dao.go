package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/OJOMB/mvsc-microservice/utils"
)

// UserDAO is an exported instance of userDAO
var UserDAO userDAOInterface

type userDAOInterface interface {
	GetUser(id string) (*User, *utils.ApplicationError)
	CreateUser(user *User) *utils.ApplicationError
}

func init() {
	UserDAO = &userDAO{}
}

type userDAO struct{}

// GetUser retrieves the user with the given unique id
func (u *userDAO) GetUser(id string) (*User, *utils.ApplicationError) {
	users, err := u.readUsersFromFile()
	if err != nil {
		return nil, &utils.ApplicationError{
			Msg:    err.Error(),
			Status: http.StatusInternalServerError,
			Code:   "internal_server_error",
		}
	}
	user, ok := users[id]
	if !ok {
		return nil, &utils.ApplicationError{
			Msg:    fmt.Sprintf("User with id: %s not found", id),
			Status: http.StatusNotFound,
			Code:   "user_not_found",
		}
	}
	return &user, nil
}

// CreateUser creates a user entry in the model
func (u *userDAO) CreateUser(user *User) *utils.ApplicationError {
	users, err := u.readUsersFromFile()
	if err != nil {
		return &utils.ApplicationError{
			Msg:    err.Error(),
			Status: http.StatusInternalServerError,
			Code:   "internal_server_error",
		}
	}
	if _, ok := users[user.ID]; ok {
		return &utils.ApplicationError{
			Msg:    fmt.Sprintf("Failed to create user with id: %s because it already exists", user.ID),
			Status: http.StatusConflict,
			Code:   "user_already_exists",
		}
	}
	users[user.ID] = *user

	err = u.writeUsersToFile(users)
	if err != nil {
		return &utils.ApplicationError{
			Msg:    err.Error(),
			Status: http.StatusInternalServerError,
			Code:   "internal_server_error",
		}
	}
	return nil
}

func (u *userDAO) getPathToJSONFile() (path string, err error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return
	}
	_, dir := filepath.Split(workingDir)
	fmt.Println(dir)
	if dir == "models" {
		path = "user_data.json"
	} else {
		path = "./models/user_data.json"
	}
	return
}

func (u *userDAO) readUsersFromFile() (users map[string]User, err error) {
	path, err := u.getPathToJSONFile()
	if err != nil {
		return
	}

	jsonString, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	json.Unmarshal([]byte(jsonString), &users)

	return
}

func (u *userDAO) writeUsersToFile(users map[string]User) (err error) {
	path, err := u.getPathToJSONFile()
	if err != nil {
		return
	}

	data, err := json.Marshal(&users)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, data, 0644)
	return
}
