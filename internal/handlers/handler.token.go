package handlers

import (
	"avito-backend-assignment/internal/helpers"
	"avito-backend-assignment/internal/models"
	"avito-backend-assignment/internal/services"
	"avito-backend-assignment/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TokenHandler struct {
	tokenService *services.TokenService
}

func NewTokenHandler(tokenService *services.TokenService) *TokenHandler {
	return &TokenHandler{
		tokenService: tokenService,
	}
}

func (h *TokenHandler) Generate(ctx echo.Context) error {
	var generateTokenRequest models.GenerateTokenRequest

	err := helpers.BindAndValidate(ctx, &generateTokenRequest)
	if err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	token, statusCode, err := h.tokenService.Generate(generateTokenRequest)
	if err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	generateTokenResponse := models.GenerateTokenResponse{
		ID:    token.ID,
		Token: token.Token,
	}

	return NewSuccessResponse(ctx, statusCode, "Token successfully generated", generateTokenResponse)

}

func (h *TokenHandler) Retrieve(ctx echo.Context) error {
	tokenId := ctx.QueryParam("token-id")

	logger.ZeroLogger.Info().Msgf("tokeId: %s: ", tokenId)
	token, statusCode, err := h.tokenService.FindByTokenId(tokenId)
	if err != nil {
		logger.ZeroLogger.Error().Msg(err.Error())
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	generateTokenResponse := models.GenerateTokenResponse{
		ID:    token.ID,
		Token: token.Token,
	}

	return NewSuccessResponse(ctx, statusCode, "Token successfully found", generateTokenResponse)

}
