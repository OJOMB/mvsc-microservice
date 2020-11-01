package controllers

import (
	"github.com/gin-gonic/gin"
)

// Threads is the interface to the threads world
var Threads ThreadsInterface

func init() {
	Threads = &threads{}
}

// ThreadsInterface describes the controllers required for Threads as an entity
type ThreadsInterface interface {
	GetThread() gin.HandlerFunc
	CreateThread() gin.HandlerFunc
	UpdateThread() gin.HandlerFunc
	DeleteThread() gin.HandlerFunc
}

// Threads is a concrete implementation of ThreadsInterface
type threads struct{}

// GetThread gets a thread
func (t *threads) GetThread() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// CreateThread creates a thread
func (t *threads) CreateThread() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// UpdateThread gets a thread
func (t *threads) UpdateThread() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// DeleteThread deletes a thread
func (t *threads) DeleteThread() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}
