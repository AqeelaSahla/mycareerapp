package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Env struct {
	AppPort int `env:"APP_PORT"`

	DBUsername string `env:"DB_USERNAME"`
	DBPassword string `env:"DB_PASSWORD"`
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBName     string `env:"DB_NAME"`

	JwtSecret  string `env:"JWT_SECRET"`
	JwtExpired int    `env:"JWT_EXPIRED"`
}

func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	_env := new(Env)
	err = env.Parse(_env)
	if err != nil {
		return nil, err
	}

	return _env, nil
}
