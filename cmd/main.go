package main

import (
	"github.com/joho/godotenv"
	"itec.chat/pkg/logging"
	"itec.chat/pkg/repositories/postgres"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. %s", err.Error())
	}

	logger := logging.GetLogger()

	db, err := postgres.NewPostgresDB(&postgres.PostgresDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
		Logger:   logger,
	})
	if err != nil {
		logger.Panicf("Error while initialisation database:%s", err)
	}

}
