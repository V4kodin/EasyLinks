package errors

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message})
}
