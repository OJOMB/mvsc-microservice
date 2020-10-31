package services

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/OJOMB/mvsc-microservice/models"
	"github.com/OJOMB/mvsc-microservice/utils"
)

var (
	userDAOMock        UserDAOMock
	getUserFunction    func(userID string) (*models.User, *utils.ApplicationError)
	createUserFunction func(user *models.User) *utils.ApplicationError
)

func init() {
	models.UserDAO = &UserDAOMock{}
}

type UserDAOMock struct{}

func (u *UserDAOMock) GetUser(id string) (*models.User, *utils.ApplicationError) {
	return getUserFunction(id)
}

func (u *UserDAOMock) CreateUser(user *models.User) *utils.ApplicationError {
	return createUserFunction(user)
}

func TestGetUserUserFound(t *testing.T) {
	u := &models.User{
		ID:        "id",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email",
	}
	getUserFunction = func(userID string) (*models.User, *utils.ApplicationError) {
		return u, nil
	}

	resultUser, resultAppErr := UsersService.GetUser("fake-uuid")

	if !reflect.DeepEqual(resultUser, u) {
		t.Errorf(
			"Expected:\n	user: %v, ApplicationError: %v\nGot:\n    user: %v, ApplicationError: %v\n",
			nil, resultAppErr,
			resultUser, resultAppErr,
		)
	}
}

func TestGetUserUserNotFound(t *testing.T) {
	getUserFunction = func(userID string) (*models.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Msg:    fmt.Sprintf("User with id: %s not found", userID),
			Status: http.StatusNotFound,
			Code:   "user_not_found",
		}
	}

	resultUser, resultAppErr := UsersService.GetUser("fake-uuid")

	expectedAppErr := &utils.ApplicationError{
		Msg:    "User with id: fake-uuid not found",
		Status: http.StatusNotFound,
		Code:   "user_not_found",
	}

	if resultUser != nil || !reflect.DeepEqual(resultAppErr, expectedAppErr) {
		t.Errorf(
			"Expected:\n	user: %v, ApplicationError: %v\nGot:\n    user: %v, ApplicationError: %v\n",
			nil, resultAppErr,
			resultUser, resultAppErr,
		)
	}
}

func TestGetUserReadFromDBFailed(t *testing.T) {
	err := &utils.ApplicationError{
		Msg:    "open fail",
		Status: http.StatusInternalServerError,
		Code:   "internal_server_error",
	}
	getUserFunction = func(userID string) (*models.User, *utils.ApplicationError) {
		return nil, err
	}

	resultUser, resultAppErr := UsersService.GetUser("fake-uuid")

	if resultUser != nil || !reflect.DeepEqual(resultAppErr, err) {
		t.Errorf(
			"Expected:\n	user: %v, ApplicationError: %v\nGot:\n    user: %v, ApplicationError: %v\n",
			nil, resultAppErr,
			resultUser, resultAppErr,
		)
	}
}
