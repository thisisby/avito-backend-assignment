package routes

import (
	"avito-backend-assignment/internal/handlers"
	"avito-backend-assignment/internal/repositories/postgre"
	"avito-backend-assignment/internal/services"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type TokenRouter struct {
	tokenHandler *handlers.TokenHandler
	e            *echo.Group
}

func NewTokenRouter(
	conn *sqlx.DB,
	e *echo.Group,
) *TokenRouter {
	logRepository := postgre.NewPostgreLogRepository(conn)
	tokenService := services.NewTokenService(logRepository)
	tokenHandler := handlers.NewTokenHandler(tokenService)

	return &TokenRouter{
		tokenHandler: tokenHandler,
		e:            e,
	}
}

func (r *TokenRouter) Register() {
	tokenGroup := r.e.Group("/token")

	tokenGroup.POST("/generate", r.tokenHandler.Generate)
	tokenGroup.GET("/retrieve", r.tokenHandler.Retrieve)
}
