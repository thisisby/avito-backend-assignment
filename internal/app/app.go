package app

import (
	"avito-backend-assignment/internal/config"
	"avito-backend-assignment/pkg/httpserver"
	"avito-backend-assignment/pkg/logger"
	"avito-backend-assignment/pkg/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func MustRun() {

	logger.ZeroLogger.Info().Msg("Setting up default postgre connection...")
	conn, err := utils.SetupDefaultPostgreConnection()
	if err != nil {
		logger.ZeroLogger.Fatal().Msg(fmt.Errorf("app - MustRun - utils.SetupDefaultPostgreConnection: %w", err).Error())
	}
	defer conn.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time}, host=${host}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	e.Validator = utils.NewValidator()

	v1 := e.Group("/api/v1")

	setupRouters(conn, v1)

	// running server
	log.Info().Msg("Starting http server...")
	httpServer := httpserver.New(e, httpserver.Port(config.Config.Host))

	// waiting signal
	log.Info().Msg("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msg(fmt.Sprintf("app - Run - signal: " + s.String()))
	case err = <-httpServer.Notify():
		log.Info().Msg(fmt.Errorf("app - Run - httpServer.Notify: %w", err).Error())
	}

	// Graceful shutdown
	log.Info().Msg("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatal().Msg(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err).Error())
	}

}

func setupRouters(conn *sqlx.DB, e *echo.Group) {

}
