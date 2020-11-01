package controllers

import (
	"github.com/gin-gonic/gin"
)

// Posts is the interface to the posts world
var Posts PostsInterface

func init() {
	Posts = &posts{}
}

// PostsInterface describes the controllers required for Posts as an entity
type PostsInterface interface {
	GetPost() gin.HandlerFunc
	CreatePost() gin.HandlerFunc
	UpdatePost() gin.HandlerFunc
	DeletePost() gin.HandlerFunc
}

// Posts is a concrete implementation of PostsInterface
type posts struct{}

// GetPost gets a post
func (t *posts) GetPost() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// CreatePost creates a post
func (t *posts) CreatePost() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// UpdatePost gets a post
func (t *posts) UpdatePost() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}

// DeletePost deletes a post
func (t *posts) DeletePost() gin.HandlerFunc {
	// TODO: implement
	return func(ctx *gin.Context) {}
}
