package runner

import (
	"context"
	"fmt"
	"os"
	"time"

	"net/http"
	"os/signal"

	"github.com/charmbracelet/log"
	"github.com/kitabisa/teler-proxy/common"
	"github.com/kitabisa/teler-proxy/pkg/tunnel"
)

func New(opt *common.Options) error {
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		log.Warn("Interuppted. Shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server shutdown error", "error", err)
		}
	}()

	log.Info("Server started", "port", opt.Port, "pid", os.Getpid())

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
