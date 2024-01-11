package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

type Config struct {
	LimitPersent int
}

func NewConfig() (*Config, error) {
	limit, ok := os.LookupEnv("LIMIT_PERCENT")
	if !ok {
		return nil, errors.New("Empty .env file")
	}
	res, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}
	return &Config{
		LimitPersent: res,
	}, nil
}
