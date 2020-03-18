package main

import (
	"context"
	"time"

	"github.com/fenollp/noterepl/pkg"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
)

func (srv *Server) prepare(ctx context.Context, fs ...pkg.AuthOption) (context.Context, func(), error) {
	opts := &pkg.AuthOptions{}
	for _, f := range fs {
		f(opts)
	}
	return opts.Deadline(ctx, ksuid.New().String())
}

// Server holds connections to our services accessible by gRPC rpcs.
type Server struct{}

// Close ...
func (srv *Server) Close(ctx context.Context) {
	log := pkg.NewLogFromCtx(ctx)
	log.Info("down")
	// Shutdown server's services here
}

// NewServer opens connections to our services
func NewServer(ctx context.Context) (srv *Server, err error) {
	log := pkg.NewLogFromCtx(ctx)
	srv = &Server{}
	start := time.Now()

	log.Info("server ready", zap.Duration("in", time.Since(start)))
	return
}
