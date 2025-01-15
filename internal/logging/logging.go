package logging

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func NewLogger(mode string, level string) *slog.Logger {
	var logger *slog.Logger

	handlerOptions := &slog.HandlerOptions{
		Level:       getLogLevel(level),
		AddSource:   true,
		ReplaceAttr: replaceAttr(mode),
	}

	if mode == "dev" {
		logger = slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, handlerOptions))
	}

	return logger
}

func NewLoggerFromEnv() *slog.Logger {
	mode := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_MODE")))
	level := strings.TrimSpace(strings.ToLower(os.Getenv("LOG_LEVEL")))
	return NewLogger(mode, level)
}

type slogAttr func(groups []string, attr slog.Attr) slog.Attr

func replaceAttr(mode string) slogAttr {
	return func(groups []string, attr slog.Attr) slog.Attr {
		if attr.Key == slog.TimeKey {
			attr.Key = "time"
			attr.Value = slog.TimeValue(attr.Value.Time().UTC())
		}
		if attr.Key == slog.MessageKey {
			attr.Key = "message"
		}
		if attr.Key == slog.SourceKey {
			source := attr.Value.Any().(*slog.Source)
			attr.Key = "caller"
			if mode == "dev" {
				attr.Value = slog.StringValue(fmt.Sprintf("%s:%d", source.Function, source.Line))
			} else {
				attr.Value = slog.StringValue(fmt.Sprintf("%s:%d", source.File, source.Line))
			}
		}
		return attr
	}
}

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

func getLogLevel(level string) slog.Level {
	switch strings.ToUpper(strings.TrimSpace(level)) {
	case levelDebug:
		return slog.LevelDebug
	case levelInfo:
		return slog.LevelInfo
	case levelWarn:
		return slog.LevelWarn
	case levelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
