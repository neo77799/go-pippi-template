package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shibu/omikuji-app/internal/models"
	"github.com/shibu/omikuji-app/internal/repository"
)

type OmikujiService interface {
	GetAllResults(c *gin.Context) ([]models.OmikujiResult, error)
	GetRandomResult(c *gin.Context) (*models.OmikujiResult, error)
}

type omikujiService struct {
	repo repository.OmikujiRepository
}

func NewOmikujiService(repo repository.OmikujiRepository) OmikujiService {
	return &omikujiService{repo: repo}
}

func (s *omikujiService) GetAllResults(c *gin.Context) ([]models.OmikujiResult, error) {
	return s.repo.FindAll(c)
}

func (s *omikujiService) GetRandomResult(c *gin.Context) (*models.OmikujiResult, error) {
	return s.repo.FindRandom(c)
}
