package main

import (
	"context"
	"time"

	"github.com/fenollp/noterepl/pkg"
	"go.uber.org/zap"
)

// Print runs code
func (srv *Server) Print(ctx context.Context, req *pkg.PrintReq) (rep *pkg.PrintRep, err error) {
	ctx, cancel, err := srv.prepare(ctx)
	defer cancel()
	if err != nil {
		return
	}
	log := pkg.NewLogFromCtx(ctx)
	path := req.GetPath()
	log.Info("handling Print", zap.Int("path depth", len(path)))
	start := time.Now()

	log.Info("handled Print", zap.Duration("in", time.Since(start)))
	return
}
