package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OJOMB/mvsc-microservice/utils"
)

func readUserFile() (map[string]User, error) {
	jsonString, err := ioutil.ReadFile("./models/user_data.json")
	if err != nil {
		return nil, err
	}
	var users map[string]User
	json.Unmarshal([]byte(jsonString), &users)

	return users, nil
}

func writeUserFile(users map[string]User) error {
	data, err := json.Marshal(&users)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./models/user_data.json", data, 0644)
	return err
}

// GetUser retrieves the user with the given unique id
func GetUser(id string) (*User, *utils.ApplicationError) {
	users, err := readUserFile()
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
