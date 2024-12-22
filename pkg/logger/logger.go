package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func SetupLoger(logLevel string) *slog.Logger {
    var log *slog.Logger
    switch logLevel {
    case "DEBUG":
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    case "INFO":
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    case "WARN":
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    case "ERROR":
        log = slog.New(
            slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    }
    
    return log
}
