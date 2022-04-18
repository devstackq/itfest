package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bimbo/internal/config"
	"bimbo/internal/server"
)

func Start() {
	var err error
	var app *server.App

	cfg := config.GetConfig()

	if app, err = server.NewApp(cfg); err != nil {
		log.Fatalf("could not create app instance: %s\n", err.Error())
	}

	app.Initialize()

	ctx, shutdown := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		shutdown()
	}()

	app.Logger.Info("starting application...")
	app.Run(ctx)
}
