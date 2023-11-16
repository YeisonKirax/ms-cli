package config

import (
	"os"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment struct {
	Port string `env:"PORT,required"`
	Env  string `env:"ENV,required"`
}

func Load() Environment {
	environment := os.Getenv("ENV")
	if environment == "local" || environment == "" {
		dotenvErr := godotenv.Load()
		if dotenvErr != nil {
			logrus.Fatalf("Error loading .env file")
		}
	}

	envConfig := Environment{}
	err := env.Parse(&envConfig)
	if err != nil {
		logrus.Fatalf("unable to parse environment variables: %e", err)
	}
	return envConfig
}
