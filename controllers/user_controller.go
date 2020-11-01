package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/OJOMB/mvsc-microservice/services"
	"github.com/OJOMB/mvsc-microservice/utils"
)

var (
	// Users is the interface to the world of users
	Users UsersInterface
)

func init() {
	Users = &users{}
}

// UsersInterface describes the controllers required for Users as an entity
type UsersInterface interface {
	GetUser() gin.HandlerFunc
	CreateUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
}

type users struct{}

// GetUser returns User data from the model for a given user id
func (u *users) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user id from context params
		id := ctx.Param("userID")

		// check we got a valid uuid4 string
		if _, err := uuid.Parse(id); err != nil {
			log.Printf("user id not valid uuid4: %s", id)
			appErr := &utils.ApplicationError{
				Msg:    fmt.Sprintf("user id must be valid uuid4 string: %s", id),
				Status: http.StatusBadRequest,
				Code:   "bad_request",
			}
			ctx.JSON(appErr.Status, appErr)
			return
		}

		// request user from server
		user, appErr := services.UsersService.GetUser(id)

		if appErr != nil {
			log.Printf(appErr.Msg)
			ctx.JSON(appErr.Status, appErr)
			return
		}

		log.Printf("Successfully retrieved user with id: %s", id)
		ctx.JSON(http.StatusOK, user)
	}
}

// CreateUser persists a new user
func (u *users) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, err := ioutil.ReadAll(ctx.Request.Body)

		// case when we fail to read user body
		if err != nil {
			log.Printf("Received unreadable request body")
			appErr := &utils.ApplicationError{
				Msg:    fmt.Sprint("Received unreadable request body"),
				Status: http.StatusBadRequest,
				Code:   "bad_request",
			}
			ctx.JSON(appErr.Status, appErr)
			return
		}

		// use data from request to create new user
		user, appErr := services.UsersService.CreateUser(userData)
		if err != nil {
			ctx.JSON(appErr.Status, appErr)
			return
		}
		ctx.JSON(http.StatusCreated, user)
	}
}

// UpdateUser gets a user
func (u *users) UpdateUser() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// DeleteUser deletes a user
func (u *users) DeleteUser() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}
