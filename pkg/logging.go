package pkg

import (
	"context"
	"time"

	"go.uber.org/zap"
)

var baseLog *zap.Logger

func MustSetupLogging() {
	var err error
	if baseLog == nil {
		baseLog, err = zap.NewDevelopment(
			zap.AddCaller(),
			// zap.Fields(
			// 	zap.Int("pid", os.Getpid()),
			// 	zap.String("exe", path.Base(os.Args[0])),
			// ),
		)
		if err != nil {
			panic(err)
		}
	}
}

type logingContextKey int

const (
	uniqueIdKey logingContextKey = iota
)

func CtxUID(ctx context.Context) string { return ctx.Value(uniqueIdKey).(string) }

func NewLogFromCtx(ctx context.Context, noId ...bool) *zap.Logger {
	l := baseLog
	if ctx != nil {
		if uniqueId, ok := ctx.Value(uniqueIdKey).(string); ok {
			l = l.With(zap.String("", uniqueId))
		} else if len(noId) == 1 && !noId[0] {
			panic("reqid unset")
		}
	}
	return l
}

type AuthOptions struct {
	noDeadline bool
}

func (opts *AuthOptions) Deadline(ctx context.Context, uid string) (
	context.Context,
	func(),
	error,
) {
	cancel := func() {}
	if !opts.noDeadline {
		ctx, cancel = context.WithTimeout(ctx, 3*500*time.Millisecond)
	}
	ctx = context.WithValue(ctx, uniqueIdKey, uid)
	return ctx, cancel, nil
}

type AuthOption func(*AuthOptions)

func OptNoDeadline() AuthOption { return func(a *AuthOptions) { a.noDeadline = true } }
