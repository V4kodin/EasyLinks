package errors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	StatusCode int
	Message    string
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
var ErrorMapString = map[error]int{
	ErrorMap[0]:  0,
	ErrorMap[1]:  1,
	ErrorMap[2]:  2,
	ErrorMap[3]:  3,
	ErrorMap[4]:  4,
	ErrorMap[5]:  5,
	ErrorMap[6]:  6,
	ErrorMap[7]:  7,
	ErrorMap[8]:  8,
	ErrorMap[9]:  9,
	ErrorMap[10]: 10,
	ErrorMap[11]: 11,
	ErrorMap[12]: 12,
	ErrorMap[13]: 13,
	ErrorMap[14]: 14,
	ErrorMap[15]: 15,
	ErrorMap[16]: 16,
}

//	ErrBadRequest = errors.New("bad request")
//	ErrNotFound   = errors.New("not found")
//	ErrCollision  = errors.New("failed to create short link due to collision")
//)

func ErrorResponse(c *gin.Context, statusCode int, message string, inernalStatusCode int) {
	log.Println(message + " " + ErrorMap[inernalStatusCode].Error())
	c.JSON(statusCode, Error{
		StatusCode: inernalStatusCode,
		Message:    message})
}
