package common

type Options struct {
	Port        int
	Destination string
	*Config
	*TLS
}
