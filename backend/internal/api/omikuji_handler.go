package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shibu/omikuji-app/internal/models"
	"github.com/shibu/omikuji-app/internal/service"
)

type OmikujiHandler struct {
	service service.OmikujiService
}

func NewOmikujiHandler(service service.OmikujiService) *OmikujiHandler {
	return &OmikujiHandler{service: service}
}

func (h *OmikujiHandler) GetResults(c *gin.Context) {
	result, err := h.service.GetRandomResult(c)
	if err != nil || result == nil {
		c.JSON(http.StatusInternalServerError, models.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Failed to get omikuji result",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result.Result})
}
