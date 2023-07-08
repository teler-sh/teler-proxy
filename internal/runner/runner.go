package runner

import (
	"fmt"

	"net/http"

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

	log.Info("Server started", "port", opt.Port)

	return server.ListenAndServe()
}
