package runner

import (
	"flag"
	"os"

	"github.com/teler-sh/teler-proxy/common"
	"github.com/teler-sh/teler-proxy/internal/logger"
)

func ParseOptions() *common.Options {
	opt := &common.Options{}
	cfg := &common.Config{}
	tls := &common.TLS{}

	opt.Logger = logger.New()

	flag.IntVar(&opt.Port, "p", 1337, "")
	flag.IntVar(&opt.Port, "port", 1337, "")

	flag.StringVar(&opt.Destination, "d", "", "")
	flag.StringVar(&opt.Destination, "dest", "", "")

	flag.StringVar(&cfg.Path, "c", "", "")
	flag.StringVar(&cfg.Path, "conf", "", "")

	flag.StringVar(&cfg.Format, "f", "yaml", "")
	flag.StringVar(&cfg.Format, "format", "yaml", "")

	flag.StringVar(&tls.CertPath, "cert", "", "")
	flag.StringVar(&tls.KeyPath, "key", "", "")

	flag.BoolVar(&version, "V", false, "")
	flag.BoolVar(&version, "version", false, "")

	flag.Usage = func() {
		common.PrintBanner()
		common.PrintUsage()
	}
	flag.Parse()

	if version {
		common.PrintVersion()
		os.Exit(1)
	}

	if opt.Destination == "" {
		common.PrintBanner()
		opt.Logger.Fatal("Something went wrong", "err", "missing destination address")
	}

	common.PrintBanner()

	opt.Config = cfg
	opt.TLS = tls

	return opt
}
