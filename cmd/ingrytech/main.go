package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ingrytech/internal/api/rest/handler"
	"ingrytech/internal/api/rest/server"
	"ingrytech/internal/config"
	"ingrytech/internal/service/app"
	"ingrytech/internal/storage/postgres"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

const (
	prod = "prod"
)

func main() {
	cfg := config.MustLoad()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	l, err := initLogger(cfg.Env)
	logger := l.Sugar()
	defer func() {
		err := logger.Sync()
		if err != nil {
			log.Print(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("logger creating error %v", err)
	}

	logger.Info(slog.String("env", cfg.Env))
	logger.Debug("debug messages are enabled")

	// postgres
	db, err := postgres.New(ctx, cfg.PostgresURI, logger)
	if err != nil {
		logger.Fatal("Failed to init db: ", err)
	}

	// app
	app := app.New(db)

	// go
	g, gCtx := errgroup.WithContext(ctx)

	// rest server
	handler := handler.New(app)

	rest := server.New(cfg.HTTPServer.Port, handler, logger)

	logger.Info("starting REST server")
	g.Go(func() error {
		err := rest.Start()
		if err != nil {
			logger.Error(err.Error())
			return err
		}
		return nil
	})
	logger.Info("REST server started")

	g.Go(func() error {
		<-gCtx.Done()
		logger.Info("stopping REST server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return rest.Shutdown(shutdownCtx)
	})

	if err = g.Wait(); err != nil {
		logger.Error("errors from errorGroup: %v", err.Error())
	}

	logger.Info("REST server stopped")
	logger.Info("main done")
}

func initLogger(env string) (*zap.Logger, error) {
	var zapConfig zap.Config

	switch env {
	case prod:
		zapConfig = zap.NewProductionConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	zapConfig.EncoderConfig.TimeKey = "time"
	l, err := zapConfig.Build()

	return l, err
}
