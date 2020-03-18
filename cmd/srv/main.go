package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/fenollp/noterepl/pkg"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var grpcHost = os.Getenv("GRPC_HOST")

func init() {
	pkg.MustSetupLogging()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log := pkg.NewLogFromCtx(ctx, true)
	log.Info("listening on", zap.String("grpcHost", grpcHost))

	log.Info("starting runtime logic...")
	srv, err := NewServer(ctx)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}
	defer srv.Close(ctx)

	var (
		lis   net.Listener
		sOpts []grpc.ServerOption
		s     *grpc.Server
	)
	if lis, err = net.Listen("tcp", grpcHost); err != nil {
		log.Fatal("failed to listen", zap.Error(err))
	}
	defer lis.Close()
	s = grpc.NewServer(sOpts...)
	defer s.Stop()
	pkg.RegisterNoteREPLServer(s, srv)

	go func() {
		// Cuts ctx
		defer cancel()
		// Cuts Prometheus metrics
		defer func() {
			if err := httProm.Shutdown(ctx); err != nil {
				log.Error("", zap.Error(err))
			}
		}()
		// Starves gRPC clients
		defer s.GracefulStop()

		die := make(chan os.Signal, 1)
		signal.Notify(die, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-die:
			log.Info("dying", zap.String("sig", sig.String()))
		case <-ctx.Done():
			log.Info("background context DONE")
		}
	}()

	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve gRPC", zap.Error(err))
	}
}
