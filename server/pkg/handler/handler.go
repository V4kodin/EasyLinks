package handler

import (
	"EasyLinks/server/pkg/errors"
	"EasyLinks/server/pkg/service"
	"github.com/gin-gonic/gin"
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
		api.GET("/getlink", h.getShortLink)
	}
	return router
}

func (h *Handler) createShortLink(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error(), errors.ErrorMapString[err])
		return
	}

	ShortURL, err := h.service.AddURL(string(body))
	switch {
	// if err nil send shortURL to user
	case err == nil:
		c.JSON(http.StatusOK, ShortURL)
	case err == errors.ErrorMap[6]:
		c.JSON(http.StatusOK, ShortURL)
	default:
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error(), errors.ErrorMapString[err])
	}
}

func (h *Handler) getShortLink(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		errors.ErrorResponse(c, http.StatusBadRequest, err.Error(), errors.ErrorMapString[err])
		return
	}
	ShortURL, err := h.service.GetURL(string(body))
	switch {
	case ShortURL.Error == 0:
		c.JSON(http.StatusOK, ShortURL)
	default:
		errors.ErrorResponse(c, http.StatusInternalServerError, err.Error(), errors.ErrorMapString[err])
	}

}
