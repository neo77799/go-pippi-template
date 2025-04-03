package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *OmikujiHandler) *gin.Engine {
	r := gin.Default()
	r.Use(SetupMiddleware())
	r.Use(LoggerMiddleware())

	r.GET("/api/omikuji", handler.GetResults)
	return r
}
