package runner

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"net/http"
	"os/signal"

	"github.com/charmbracelet/log"
	"github.com/kitabisa/teler-proxy/common"
	"github.com/kitabisa/teler-proxy/pkg/tunnel"
)

type Runner struct {
	*http.Server
}

func New(opt *common.Options) error {
	reachable := isReachable(opt.Destination, 5*time.Second)
	if !reachable {
		return errDestUnreachable
	}

	tun, err := tunnel.NewTunnel(opt.Port, opt.Destination, opt.Config.Path, opt.Config.Format)
	if err != nil {
		return err
	}

	logger := log.StandardLog(log.StandardLogOptions{
		ForceLevel: log.ErrorLevel,
	})

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", opt.Port),
		Handler:  tun,
		ErrorLog: logger,
	}

	run := &Runner{Server: server}
	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go func() {
		if err := run.notify(sig); err != nil {
			log.Fatal("Something went wrong", "err", err)
		}
	}()

	log.Info("Server started", "port", opt.Port, "pid", os.Getpid())
	return run.start()
}

func (r *Runner) start() error {
	err := r.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (r *Runner) shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return r.Server.Shutdown(ctx)
}

func (r *Runner) restart() error {
	if err := r.shutdown(); err != nil {
		return err
	}

	return r.start()
}

func (r *Runner) notify(sigCh chan os.Signal) error {
	for {
		sig := <-sigCh
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Warn("Interrupted. Shutting down...")
			return r.shutdown()
		case syscall.SIGUSR1:
			log.Info("Restarting server...")
			return r.restart()
		}
	}
}
