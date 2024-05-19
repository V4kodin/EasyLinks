package handler

import (
	"EasyLinks/server/pkg/errors"
	"EasyLinks/server/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgerrcode"
	"github.com/lib/pq"
	"io"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/createlink", h.createShortLink)
		//links.POST("/getlink", h.createLink)
	}
	return router
}

func (h *Handler) createShortLink(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//userID := GetUserID(c)

	shortURL, err := h.service.AddURL(string(body))

	if err, ok := err.(*pq.Error); ok {
		if err.Code == pgerrcode.UniqueViolation {
			c.String(http.StatusConflict, fmt.Sprintf("%s/%s", shortURL))
			return
		}
	}

	if err != nil {
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusCreated, fmt.Sprintf("%s/%s", shortURL))
}
