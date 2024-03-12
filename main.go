package main

import (
	"context"
	"fmt"
	"moneda/evaluation/global"
	"moneda/evaluation/modules/auth"
	"moneda/evaluation/modules/flights"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/http2"
)

func main() {
  err := godotenv.Load()

  if err != nil {
    panic("Unable to load .env file")
  }

  dbData := global.LoadDB()

	httpPort := os.Getenv("PORT")

	e := echo.New()

  apiGroup := e.Group("/api")
  auth.RegisterAuthRoutes(apiGroup)
  flights.RegisterFlightsRoutes(apiGroup, dbData)

	server := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}

	go func() {
		if err := e.StartH2CServer(fmt.Sprintf(":%s", httpPort), server); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(fmt.Sprintf("Shutting down the server: %s", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
