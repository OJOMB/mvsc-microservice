package services

import (
	"github.com/OJOMB/mvsc-microservice/models"
	"github.com/OJOMB/mvsc-microservice/utils"
)

// GetUser calls the model interface to retrieve the user with the given unique id
func GetUser(id string) (*models.User, *utils.ApplicationError) {
	return models.GetUser(id)
}
