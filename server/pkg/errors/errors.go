package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message string
}

var ErrorMap = map[int]error{
	0:  errors.New("OK"),
	1:  errors.New("CANCELLED"),
	2:  errors.New("UNKNOWN"),
	3:  errors.New("INVALID_ARGUMENT"),
	4:  errors.New("DEADLINE_EXCEEDED"),
	5:  errors.New("NOT_FOUND"),
	6:  errors.New("ALREADY_EXISTS"),
	7:  errors.New("PERMISSION_DENIED"),
	8:  errors.New("RESOURCE_EXHAUSTED"),
	9:  errors.New("FAILED_PRECONDITION"),
	10: errors.New("ABORTED"),
	11: errors.New("OUT_OF_RANGE"),
	12: errors.New("UNIMPLEMENTED"),
	13: errors.New("INTERNAL"),
	14: errors.New("UNAVAILABLE"),
	15: errors.New("DATA_LOSS"),
	16: errors.New("UNAUTHENTICATED"),
}

//	ErrBadRequest = errors.New("bad request")
//	ErrNotFound   = errors.New("not found")
//	ErrCollision  = errors.New("failed to create short link due to collision")
//)

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.JSON(statusCode, Error{Message: message})
}
