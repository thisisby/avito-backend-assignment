package config

import (
	"avito-backend-assignment/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	Env string `env:"env"`

	Host string `yaml:"host"`
	Port string `yaml:"port"`

	DBHost     string `yaml:"db_host"`
	DBPort     string `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`

	RedisHost     string `yaml:"redis_host"`
	RedisPort     string `yaml:"redis_port"`
	RedisPassword string `yaml:"redis_password"`
	RedisExpires  string `yaml:"redis_expires"`

	JwtSecretKey             string `yaml:"jwt_secret_key"`
	JwtIssuer                string `yaml:"jwt_issuer"`
	JwtAccessTokenExpiresIn  int    `yaml:"jwt_access_token_expires_in"`
	JwtRefreshTokenExpiresIn int    `yaml:"jwt_refresh_token_expires_in"`
}

var Config config

func (c *config) MustInitializeConfig() {
	err := cleanenv.ReadConfig("./internal/config/config.yml", &Config)
	if err != nil {
		logger.ZeroLogger.Fatal().Msgf("config -> MustInitializeConfig -> cleanenv.ReadConfig: %v", err)
	}

}
