package utils

import "github.com/gin-gonic/gin"

// Respond responds with the appropriate data format
func Respond(ctx *gin.Context, HTTPStatusCode int, data interface{}) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(HTTPStatusCode, data)
	} else if ctx.GetHeader("Accept") == "application/yaml" {
		ctx.YAML(HTTPStatusCode, data)
	} else {
		ctx.JSON(HTTPStatusCode, data)
	}
}

// RespondError responds with an error in the appropriate format.
func RespondError(ctx *gin.Context, appErr *ApplicationError) {
	Respond(ctx, appErr.Status, appErr)
}
