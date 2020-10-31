package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"github.com/OJOMB/mvsc-microservice/services"
	"github.com/OJOMB/mvsc-microservice/utils"
	"github.com/gorilla/mux"
)

// GetUser returns User data from the model for a given user id
func (c *Controller) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get user id from mux vars
		id := mux.Vars(r)["id"]

		// check we got a valid uuid4 string
		if _, err := uuid.Parse(id); err != nil {
			c.logger.Printf("user id not valid uuid4: %s", id)
			appErr := &utils.ApplicationError{
				Msg:    fmt.Sprintf("user id must be valid uuid4 string: %s", id),
				Status: http.StatusBadRequest,
				Code:   "bad_request",
			}
			respData, mErr := json.Marshal(appErr)
			if mErr != nil {
				c.logger.Printf("Failed to marshal error response: %v", mErr)
				http.Error(w, mErr.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(appErr.Status)
			w.Write(respData)
			return
		}

		// request user from server
		user, err := services.UsersService.GetUser(id)

		if err != nil {
			c.logger.Printf(err.Msg)
			respData, mErr := json.Marshal(err)
			if mErr != nil {
				c.logger.Printf("Failed to marshal error response: %v", mErr)
				http.Error(w, mErr.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(err.Status)
			w.Write(respData)
			return
		}

		c.logger.Printf("Successfully retrieved user with id: %s", id)
		respData, mErr := json.Marshal(user)
		if mErr != nil {
			c.logger.Printf("Failed to marshal response: %v", mErr)
			http.Error(w, mErr.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(respData)
	}
}

// CreateUser persists a new user
func (c *Controller) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData, err := ioutil.ReadAll(r.Body)

		// case when we fail to read post body
		if err != nil {
			c.logger.Printf("Received unreadable request body")
			appErr := &utils.ApplicationError{
				Msg:    fmt.Sprint("Received unreadable request body"),
				Status: http.StatusBadRequest,
				Code:   "bad_request",
			}
			respData, mErr := json.Marshal(appErr)
			if mErr != nil {
				c.logger.Printf("Failed to marshal error response: %v", mErr)
				http.Error(w, mErr.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(appErr.Status)
			w.Write(respData)
			return
		}

		// use data from request to create new user
		user, appErr := services.UsersService.CreateUser(userData)
		if err != nil {
			respBody, _ := json.Marshal(appErr)
			w.WriteHeader(appErr.Status)
			w.Write(respBody)
		}
		respBody, _ := json.Marshal(user)
		w.WriteHeader(http.StatusCreated)
		w.Write(respBody)
	}
}
