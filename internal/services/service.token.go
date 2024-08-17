package services

import (
	"avito-backend-assignment/internal/constants"
	"avito-backend-assignment/internal/errs"
	"avito-backend-assignment/internal/helpers"
	"avito-backend-assignment/internal/models"
	"avito-backend-assignment/internal/repositories"
	"errors"
	"fmt"
	"net/http"
)

type TokenService struct {
	logRepository repositories.LogRepository
}

func NewTokenService(logRepository repositories.LogRepository) *TokenService {
	return &TokenService{
		logRepository: logRepository,
	}
}

type TokenPayload struct {
	ID    string
	Token string
}

func (s *TokenService) Generate(generateTokenRequest models.GenerateTokenRequest) (TokenPayload, int, error) {
	var logModel models.Log
	token, err := helpers.GenerateToken(constants.MapIntToToken[generateTokenRequest.Type], generateTokenRequest.Length)
	if err != nil {
		if errors.Is(err, errs.ErrUnsupportedTokenType) {
			return TokenPayload{}, http.StatusBadRequest, errs.ErrUnsupportedTokenType
		}
		return TokenPayload{}, http.StatusInternalServerError, fmt.Errorf("TokenService.Generate - helpers.GenerateToken: %w", err)
	}

	// TODO get actual user agent and url
	logModel = models.Log{
		Token:     token,
		UserAgent: "agent",
		Url:       "url",
	}

	id, err := s.logRepository.Save(logModel)
	if err != nil {
		return TokenPayload{}, http.StatusInternalServerError, fmt.Errorf("TokenService.Generate - s.logRepository.Save: %w", err)
	}

	return TokenPayload{
		ID:    id,
		Token: token,
	}, http.StatusCreated, nil
}

func (s *TokenService) FindByTokenId(tokenId string) (TokenPayload, int, error) {
	token, err := s.logRepository.FindById(tokenId)
	if err != nil {
		if errors.Is(err, errs.ErrTokenNotFound) || errors.Is(err, errs.ErrInvalidUUIDFormat) {
			return TokenPayload{}, http.StatusBadRequest, err
		}
		return TokenPayload{}, http.StatusInternalServerError, fmt.Errorf("TokenService.FindByTokenId - s.logRepository.FindByID: %w", err)
	}

	return TokenPayload{
		ID:    token.TokenID,
		Token: token.Token,
	}, http.StatusOK, nil
}
