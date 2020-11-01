package controllers

import (
	"github.com/gin-gonic/gin"
)

// Topics is the interface to the topics world
var Topics TopicsInterface

func init() {
	Topics = &topics{}
}

// TopicsInterface describes the controllers required for Topics as an entity
type TopicsInterface interface {
	GetTopic() gin.HandlerFunc
	CreateTopic() gin.HandlerFunc
	UpdateTopic() gin.HandlerFunc
	DeleteTopic() gin.HandlerFunc
}

// Topics is a concrete implementation of TopicsInterface
type topics struct{}

// GetTopic gets a topic
func (t *topics) GetTopic() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// CreateTopic creates a topic
func (t *topics) CreateTopic() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// UpdateTopic gets a topic
func (t *topics) UpdateTopic() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// DeleteTopic deletes a topic
func (t *topics) DeleteTopic() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}
