package api

import "github.com/gin-gonic/gin"

// Controller - interface
type Controller interface {
	RegisterAPIs(public *gin.RouterGroup, closed *gin.RouterGroup)
}

// ErrorResponse - error response model
type ErrorResponse struct {
	Error string `json:"message"`
}
