package utils

import (
	"avito-backend-assignment/internal/config"
	"avito-backend-assignment/pkg/postgre"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupDefaultPostgreConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable timezone=Asia/Jakarta",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBHost,
	)

	defaultDriverOptions := postgre.NewSqlxDriverOptions(
		"postgres",
		dsn,
		100,
		10,
		15*time.Minute,
	)

	conn, err := defaultDriverOptions.Connect()
	if err != nil {
		return nil, fmt.Errorf("SetupDefaultPostgreConnection - defaultDriverOptions.Connect: %w", err)
	}

	return conn, nil
}
