package runner

import (
	"flag"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/kitabisa/teler-proxy/common"
	"github.com/mattn/go-colorable"
)

func ParseOptions() *common.Options {
	opt := &common.Options{}
	cfg := &common.Config{}
	tls := &common.TLS{}

	opt.Logger = log.NewWithOptions(
		colorable.NewColorableStderr(),
		log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
		},
	)

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
