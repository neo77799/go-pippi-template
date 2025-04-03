package di

import (
	"github.com/gin-gonic/gin"
	"github.com/shibu/omikuji-app/internal/api"
	"github.com/shibu/omikuji-app/internal/repository"
	"github.com/shibu/omikuji-app/internal/service"
	"gorm.io/gorm"
)

type Container struct {
	Router         *gin.Engine
	OmikujiHandler *api.OmikujiHandler
}

func NewContainer(db *gorm.DB) *Container {
	omikujiRepo := repository.NewOmikujiRepository(db)
	omikujiService := service.NewOmikujiService(omikujiRepo)
	omikujiHandler := api.NewOmikujiHandler(omikujiService)

	return &Container{
		Router:         api.SetupRouter(omikujiHandler),
		OmikujiHandler: omikujiHandler,
	}
}
