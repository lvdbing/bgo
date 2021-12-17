package global

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/pkg/logger"
)

var (
	Logger *BgoLogger
)

type BgoLogger struct {
	*logger.Logger
	ctx context.Context
}

func NewLogger(w io.Writer, prefix string, flag int) *BgoLogger {
	l := logger.NewLogger(w, prefix, flag)
	return &BgoLogger{Logger: l}
}

func (l *BgoLogger) clone() *BgoLogger {
	nl := *l
	return &nl
}

func (l *BgoLogger) WithContext(ctx context.Context) *BgoLogger {
	nl := l.clone()
	nl.ctx = ctx
	return nl
}

func (l *BgoLogger) WithFields(f logger.Fields) *BgoLogger {
	nl := l.clone()
	nl.Logger = nl.Logger.WithFields(f)
	return nl
}

func (l *BgoLogger) WithCaller(skip int) *BgoLogger {
	nl := l.clone()
	nl.Logger = nl.Logger.WithCaller(skip)
	return nl
}

func (l *BgoLogger) WithCallersFrames() *BgoLogger {
	nl := l.clone()
	nl.Logger = nl.Logger.WithCallersFrames()
	return nl
}

func (l *BgoLogger) WithTrace() *BgoLogger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(logger.Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span-id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

func (l *BgoLogger) Debug(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Debug(v...)
}

func (l *BgoLogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Debugf(format, v...)
}

func (l *BgoLogger) Info(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Info(v...)
}

func (l *BgoLogger) Infof(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Infof(format, v...)
}

func (l *BgoLogger) Warn(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Warn(v...)
}

func (l *BgoLogger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Warnf(format, v...)
}

func (l *BgoLogger) Error(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Error(v...)
}

func (l *BgoLogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Errorf(format, v...)
}

func (l *BgoLogger) Fatal(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Fatal(v...)
}

func (l *BgoLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Fatalf(format, v...)
}

func (l *BgoLogger) Panic(ctx context.Context, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Panic(v...)
}

func (l *BgoLogger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l = l.WithContext(ctx).WithTrace()
	l.Logger.Panicf(format, v...)
}
