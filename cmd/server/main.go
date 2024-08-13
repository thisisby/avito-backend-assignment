package main

import (
	"avito-backend-assignment/internal/app"
	"avito-backend-assignment/internal/config"
	"avito-backend-assignment/pkg/logger"
)

func init() {
	logger.InitZeroLogger()
	config.Config.MustInitializeConfig()
}

func main() {
	app.MustRun()
}
