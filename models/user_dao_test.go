package models

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/OJOMB/mvsc-microservice/utils"
	"github.com/google/uuid"
)

func TestTableGetUser(t *testing.T) {
	testTable := map[string]struct {
		Input         string
		ExpectedUser  *User
		ExpectedError *utils.ApplicationError
	}{
		// success case
		"test1": {
			"2fa0b8ef-cfeb-4055-8270-299cfdef8934",
			&User{
				ID:        "2fa0b8ef-cfeb-4055-8270-299cfdef8934",
				FirstName: "Oscar",
				LastName:  "Oram",
				Email:     "oscarm-b@tutamail.com",
			},
			nil,
		},
		// fail case
		"test2": {
			"00000000-0000-0000-0000-000000000000",
			nil,
			&utils.ApplicationError{
				Msg:    "User with id: 00000000-0000-0000-0000-000000000000 not found",
				Status: http.StatusNotFound,
				Code:   "user_not_found",
			},
		},
	}

	for testName, test := range testTable {
		expectedUser := test.ExpectedUser
		expectedError := test.ExpectedError
		resultUser, resultErr := GetUser(test.Input)

		if !reflect.DeepEqual(resultUser, expectedUser) || !reflect.DeepEqual(resultErr, expectedError) {
			t.Errorf(
				"\nTEST FAILED: %s\nExpected:\n    user: %v and error: %v\nInstead got:\n    user: %v and error: %v",
				testName,
				test.ExpectedUser, test.ExpectedError,
				resultUser, resultErr,
			)
		}
	}
}

func TestTableCreateUser(t *testing.T) {
	testUUID := uuid.New().String()
	t.Logf("Test UUID: %s", testUUID)

	testTable := map[string]struct {
		Input         *User
		ExpectedError *utils.ApplicationError
	}{
		// success case
		"test1": {
			&User{
				ID:        testUUID,
				FirstName: "Oscar",
				LastName:  "Oram",
				Email:     "oscarm-b@tutamail.com",
			},
			nil,
		},
		// fail case
		"test2": {
			&User{
				ID:        "2fa0b8ef-cfeb-4055-8270-299cfdef8934",
				FirstName: "Oscar",
				LastName:  "Oram",
				Email:     "oscarm-b@tutamail.com",
			},
			&utils.ApplicationError{
				Msg:    "Failed to create user with id: 2fa0b8ef-cfeb-4055-8270-299cfdef8934 because it already exists",
				Status: http.StatusConflict,
				Code:   "user_already_exists",
			},
		},
	}

	for testName, test := range testTable {
		expectedError := test.ExpectedError
		resultErr := CreateUser(test.Input)

		if !reflect.DeepEqual(resultErr, expectedError) {
			t.Errorf(
				"\nTEST FAILED: %s\nExpected:\n    error: %v\nInstead got:\n    error: %v",
				testName,
				test.ExpectedError,
				resultErr,
			)
		}

		if testName == "test1" {
			users, _ := readUsersFromFile()
			entry, ok := users[test.Input.ID]
			if !ok {
				t.Errorf(
					"\nTEST FAILED: %s\nFailed to find entry in users with id: %s\nWrite operation must have failed",
					testName, test.Input.ID,
				)
			} else if !reflect.DeepEqual(entry, *test.Input) {
				t.Errorf(
					"\nTEST FAILED: %s\nWritten entry for user with id: %s did not match the test input\nWrite operation must have failed",
					testName, test.Input.ID,
				)
			}
		}
	}
}
