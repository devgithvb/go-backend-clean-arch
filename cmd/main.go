package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/grpc"
	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/http"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/cmd/migrations"
	"go.uber.org/zap"
)

func main() {
	// Bootstrap
	app, err := bootstrap.App(configs.Development)
	if err != nil {
		log.Fatalf("bootstrap app: %v", err)
	}

	//if app.Config.Application.Env == configs.Development {
	//	log.Println("The app is running in development")
	//}

	// Log
	app.Logger.Set().Named("Main").Info("Config", zap.Any("config", app.Config))

	// Migrations
	if err = migrations.Up(app); err != nil {
		app.Logger.Set().Named("Main").Fatal("Migrations.Up", zap.Error(err))
	}

	// Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start HTTP server
	hs := http.New(app)

	go func() {
		if err = hs.Run(); err != nil {
			app.Logger.Set().Named("Main").Fatal("Server.HTTP.Run", zap.Error(err))
		}
	}()

	// Start gRPC server
	gs := grpc.New(app)

	go func() {
		if err = gs.Run(); err != nil {
			app.Logger.Set().Named("Main").Fatal("Server.GRPC.Run", zap.Error(err))
		}
	}()

	// Start Pprof server

	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.Application.GracefulShutdownTimeout)
	defer cancel()

	if err = hs.Router.Shutdown(ctxWithTimeout); err != nil {
		app.Logger.Set().Named("Main").Error("Server.HTTP.Shutdown", zap.Error(err))
	}

	app.Logger.Set().Named("Main").Info("Received.Interrupt.Signal.For.Shutting.Down.Gracefully")

	defer func(app *bootstrap.Application) {
		err = app.CloseRedisClientConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Redis.Connection", zap.Error(err))
		}
	}(app)

	defer func(app *bootstrap.Application) {
		err = app.CloseMysqlConnection()
		if err != nil {
			app.Logger.Set().Named("Main").Error("Close.Mysql.Connection", zap.Error(err))
		}
	}(app)

	// app.ClosePostgresqlConnection() // Or etc..

	<-ctxWithTimeout.Done()
}
