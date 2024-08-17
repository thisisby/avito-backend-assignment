package routes

import (
	"avito-backend-assignment/internal/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type HealthCheckRouter struct {
	HealthCheckHandler *handlers.HealthCheckHandler
	Conn               *sqlx.DB
	echo               *echo.Group
}

func NewHealthCheck(conn *sqlx.DB, e *echo.Group) *HealthCheckRouter {

	healthCheckHandler := handlers.NewHealthCheckHandler()

	return &HealthCheckRouter{
		HealthCheckHandler: healthCheckHandler,
		Conn:               conn,
		echo:               e,
	}
}

func (r *HealthCheckRouter) Register() {
	healthCheckGroup := r.echo.Group("/health-check")

	healthCheckGroup.GET("", r.HealthCheckHandler.HealthCheck)
}
