package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"itec.chat/internal/handlers/httpHandler"
	"itec.chat/internal/wsHub"

	"itec.chat/pkg/logging"

	"itec.chat/pkg/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. %s", err.Error())
	}

	logger := logging.GetLogger()

	/*db, err := postgres.NewPostgresDB(&postgres.PostgresDB{
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
	repository := repository.New(db, logger)*/

	logger.Info("Initializing httprouter...")
	var hub = wsHub.NewHub()
	go hub.Run()
	handler := httpHandler.NewHandler(logger, hub /*, repository*/)

	server := server.NewServer(logger, *handler, os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		server.Shutdown(context.Background())

		log.Println("shutting down")
		os.Exit(0)
	}()

	if err := server.Start(); err != http.ErrServerClosed {
		logger.Panicf("Error while starting server:%s", err)
	}
	<-idleConnsClosed
}
