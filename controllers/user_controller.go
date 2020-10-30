package controllers

import (
	"encoding/json"
	"fmt"
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
		user, err := services.GetUser(id)

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
