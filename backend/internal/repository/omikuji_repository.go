package repository

import (
	"github.com/shibu/omikuji-app/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OmikujiRepository interface {
	FindAll(c *gin.Context) ([]models.OmikujiResult, error)
	FindRandom(c *gin.Context) (*models.OmikujiResult, error)
}

type omikujiRepository struct {
	db *gorm.DB
}

func NewOmikujiRepository(db *gorm.DB) OmikujiRepository {
	return &omikujiRepository{db: db}
}

func (r *omikujiRepository) FindAll(c *gin.Context) ([]models.OmikujiResult, error) {
	var results []models.OmikujiResult
	if err := r.db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *omikujiRepository) FindRandom(c *gin.Context) (*models.OmikujiResult, error) {
	var result models.OmikujiResult
	if err := r.db.Order("RANDOM()").First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
