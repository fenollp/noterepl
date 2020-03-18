package main

import (
	"context"
	"errors"
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
	ptr := req.GetPtr()
	log.Info("handling Print")
	start := time.Now()

	if ptr == 0 {
		err = errors.New("empty")
		log.Error("", zap.Error(err))
		return
	}
	object := pkg.Get(ptr)

	rep = &pkg.PrintRep{Object: object}
	log.Info("handled Print", zap.Duration("in", time.Since(start)))
	return
}
