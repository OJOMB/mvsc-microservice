package services

import (
	"encoding/json"

	"github.com/OJOMB/mvsc-microservice/models"
	"github.com/OJOMB/mvsc-microservice/utils"
	"github.com/google/uuid"
)

var (
	// UsersService is an exported UserServiceInterface
	UsersService usersServiceInterface
)

type usersServiceInterface interface {
	GetUser(id string) (*models.User, *utils.ApplicationError)
	CreateUser(userData []byte) (*models.User, *utils.ApplicationError)
}

func init() {
	UsersService = &usersService{}
}

type usersService struct{}

// GetUser calls the model interface to retrieve the user with the given unique id
func (u *usersService) GetUser(id string) (*models.User, *utils.ApplicationError) {
	return models.UserDAO.GetUser(id)
}

// CreateUser calls the model interface to create a new user
func (u *usersService) CreateUser(userData []byte) (*models.User, *utils.ApplicationError) {
	var user models.User
	json.Unmarshal(userData, &user)

	// give new user a uuid
	user.ID = uuid.New().String()

	if err := models.UserDAO.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
