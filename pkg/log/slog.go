package log

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

type Handler interface {
	InfoCtx(ctx context.Context, msg string, keyvals ...interface{})
	ErrorCtx(ctx context.Context, msg string, keyvals ...interface{})
	WarnCtx(ctx context.Context, msg string, keyvals ...interface{})
	DebugCtx(ctx context.Context, msg string, keyvals ...interface{})
}

type Handle struct {
	json bool // if true, output json format
	log  *slog.Logger
}

type HandlerOption func(*Handle)

func WithJSONFormat() HandlerOption {
	return func(h *Handle) {
		h.json = true
	}
}

type level int

const (
	LevelDebug level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func Level(level level) slog.Level {
	switch level {
	case LevelDebug:
		return slog.LevelDebug
	case LevelInfo:
		return slog.LevelInfo
	case LevelWarn:
		return slog.LevelWarn
	case LevelError:
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func NewHandler(level level, opts ...HandlerOption) Handler {
	h := &Handle{}
	for _, opt := range opts {
		opt(h)
	}
	if h.json {
		h.log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: Level(level),
		}))
	} else {
		h.log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: Level(level),
		}))
	}
	return h
}

func (h *Handle) InfoCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	h.log.Log(ctx, slog.LevelInfo, msg, keyvals...)
}

func (h *Handle) ErrorCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	h.log.Log(ctx, slog.LevelError, msg, keyvals...)
}

func (h *Handle) WarnCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	h.log.Log(ctx, slog.LevelWarn, msg, keyvals...)
}

func (h *Handle) DebugCtx(ctx context.Context, msg string, keyvals ...interface{}) {
	h.log.Log(ctx, slog.LevelDebug, msg, keyvals...)
}
