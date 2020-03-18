package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/fenollp/noterepl/pkg"
	"go.uber.org/zap"
)

// Eval runs code
func (srv *Server) Eval(ctx context.Context, req *pkg.EvalReq) (rep *pkg.EvalRep, err error) {
	ctx, cancel, err := srv.prepare(ctx)
	defer cancel()
	if err != nil {
		return
	}
	log := pkg.NewLogFromCtx(ctx)
	language := req.GetLanguage()
	code := req.GetCode()
	log.Info("handling Eval", zap.String("code", code))
	start := time.Now()

	var cmd *exec.Cmd
	switch language {
	case "id":
		cmd = exec.CommandContext(ctx, "./cmd/srv/id.sh")
	case "python3":
		cmd = exec.CommandContext(ctx, "docker", "run", "--rm", "-i", "python:3-alpine", "python3", "-c", "eval(input())")
	default:
		err = fmt.Errorf("unsupported language: %q", language)
		log.Error("", zap.Error(err))
		return
	}

	var stderr string
	var stdout []byte
	if stderr, stdout, err = eval(ctx, cmd, code); err != nil {
		return
	}

	rep = &pkg.EvalRep{
		Stderr: stderr,
		Result: pkg.ObjectFrom(stdout),
	}
	log.Info("handled Eval", zap.Duration("in", time.Since(start)))
	return
}

func eval(ctx context.Context, cmd *exec.Cmd, code string) (stderr string, stdout []byte, err error) {
	log := pkg.NewLogFromCtx(ctx)

	var stdinPipe io.WriteCloser
	if stdinPipe, err = cmd.StdinPipe(); err != nil {
		log.Error("", zap.Error(err))
		return
	}
	go func() {
		defer stdinPipe.Close()
		io.WriteString(stdinPipe, code)
	}()

	var stdoutPipe, stderrPipe io.ReadCloser
	if stdoutPipe, err = cmd.StdoutPipe(); err != nil {
		log.Error("", zap.Error(err))
		return
	}
	if stderrPipe, err = cmd.StderrPipe(); err != nil {
		log.Error("", zap.Error(err))
		return
	}

	if err = cmd.Start(); err != nil {
		log.Error("", zap.Error(err))
		return
	}

	var errbytes, outbytes []byte
	if errbytes, err = ioutil.ReadAll(stderrPipe); err != nil {
		log.Error("", zap.Error(err))
		return
	}
	if outbytes, err = ioutil.ReadAll(stdoutPipe); err != nil {
		log.Error("", zap.Error(err))
		return
	}

	if err = cmd.Wait(); err != nil {
		log.Error("", zap.Error(err))
		return
	}

	return string(errbytes), outbytes, nil
}
