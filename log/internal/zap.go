package internal

import (
	"errors"
	"io/fs"
	"log/slog"

	slogzap "github.com/samber/slog-zap/v2"

	"go.uber.org/zap"
)

func RunZap() {
	zapLogger, err := zap.NewProduction(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCaller(),
	)
	if err != nil {
		panic(err)
	}

	defer func(zapLogger *zap.Logger) {
		if err := zapLogger.Sync(); err != nil {
			// When you use sync on console, it will return an error
			var pathErr *fs.PathError
			if ok := errors.As(err, &pathErr); ok {
				return
			}
			panic(err)
		}
	}(zapLogger)

	logger := slog.New(slogzap.Option{
		Level:     slog.LevelDebug,
		Logger:    zapLogger,
		AddSource: true,
	}.NewZapHandler())

	logger.Info("Zap logger is running",
		slog.String("key", "value"),
	)

	//If you use slogzap, you can't use stack trace
	err = errors.New("error message")
	logger.
		With("error", err).
		Error("Error message")

}
