package services

import (
	"avito-backend-assignment/internal/models"
	"avito-backend-assignment/internal/repositories/postgre/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

var (
	logRepository *mocks.LogRepository
	logService    *TokenService
)

func setup(t *testing.T) {
	logRepository = mocks.NewLogRepository(t)
	logService = NewTokenService(logRepository)
}

func TestGenerate(t *testing.T) {
	setup(t)

	req := models.GenerateTokenRequest{
		Type:   1,
		Length: 10,
	}
	url := "http"
	userAgent := "Test Agent"

	t.Run("When Success to generate and store Token", func(t *testing.T) {
		logRepository.Mock.On("Generate", mock.Anything, mock.AnythingOfType("*models.Log")).Return(models.Log{}, errors.New("generation is failed")).Once()
		result, statusCode, err := logService.Generate(req, url, userAgent)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, statusCode)

	})
}
