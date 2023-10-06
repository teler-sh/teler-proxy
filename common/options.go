package common

import "github.com/charmbracelet/log"

type Options struct {
	Port        int
	Destination string
	*Config
	*TLS
	*log.Logger
}
