package main

import (
	"context"
	"fmt"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver"
	"go-backend-clean-arch-according-to-go-standards-project-layout/configs"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	// Bootstrap
	app := bootstrap.App(configs.Development)

	// Logger
	logger.Logger.Named("main").Info("config", zap.Any("config", app.Config))

	// Start server
	server := httpserver.New(app)
	go func() {
		server.Serve()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)
	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.Config.HTTPServer.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		fmt.Println("http server shutdown error", err)
	}

	fmt.Println("received interrupt signal, shutting down gracefully..")
	// Close all db connection, etc
	app.ClosePostgresqlConnection()
	app.CloseRedisClientConnection()

	<-ctxWithTimeout.Done()
}
