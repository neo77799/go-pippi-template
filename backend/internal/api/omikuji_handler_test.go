package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shibu/omikuji-app/internal/api"
	"github.com/shibu/omikuji-app/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOmikujiService struct {
	mock.Mock
}

func (m *MockOmikujiService) GetAllResults(c *gin.Context) ([]models.OmikujiResult, error) {
	args := m.Called(c)
	return args.Get(0).([]models.OmikujiResult), args.Error(1)
}

func (m *MockOmikujiService) GetRandomResult(c *gin.Context) (*models.OmikujiResult, error) {
	args := m.Called(c)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.OmikujiResult), args.Error(1)
}

func TestOmikujiHandler_GetResults(t *testing.T) {
	t.Run("正常系: おみくじが取得できる", func(t *testing.T) {
		mockService := new(MockOmikujiService)
		mockResult := &models.OmikujiResult{Result: "大吉"}
		mockService.On("GetRandomResult", mock.Anything).Return(mockResult, nil)

		handler := api.NewOmikujiHandler(mockService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/omikuji", nil)

		handler.GetResults(c)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"result":"大吉"}`, w.Body.String())
	})

	t.Run("異常系: サービスエラーが発生", func(t *testing.T) {
		mockService := new(MockOmikujiService)
		mockService.On("GetRandomResult", mock.Anything).Return(nil, assert.AnError)

		handler := api.NewOmikujiHandler(mockService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/omikuji", nil)

		handler.GetResults(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.JSONEq(t, `{"code":500,"message":"Failed to get omikuji result"}`, w.Body.String())
	})
}
