package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kitabisa/teler-waf"
	"github.com/kitabisa/teler-waf/option"
)

var config = []byte(`excludes:
    - 4
    - 5
whitelists:
    - request.Headers matches "(curl|Go-http-client|okhttp)/*" && threat == BadCrawler
    - request.URI startsWith "/wp-login.php"
    - request.IP in ["127.0.0.1", "::1", "0.0.0.0"]
    - request.Headers contains "authorization" && request.Method == "POST"
customs_from_file: ""
log_file: /tmp/teler.log
no_stderr: false
no_update_check: false
development: false
in_memory: false
falcosidekick_url: ""`)

type Tunnel struct {
	*teler.Teler
	LocalPort    int
	Destination  string
	ReverseProxy *httputil.ReverseProxy
}

func NewTunnel(port int, dest, telerOpts, optFormat string) (*Tunnel, error) {
	destURL, err := url.Parse(dest)
	if err != nil {
		return nil, err
	}

	var opt teler.Options

	tun := &Tunnel{}
	tun.Teler = teler.New()
	tun.LocalPort = port
	tun.Destination = dest
	tun.ReverseProxy = httputil.NewSingleHostReverseProxy(destURL)

	if telerOpts != "" {
		switch strings.ToLower(optFormat) {
		case "yaml":
			opt, err = option.LoadFromYAMLString(telerOpts)
		case "json":
			opt, err = option.LoadFromJSONString(telerOpts)
		case "":
			return nil, errors.New("undefined teler option format")
		default:
			return nil, errors.New("invalid teler option format")
		}

		if err != nil {
			return nil, err
		}

		tun.Teler = teler.New(opt)
	}

	return tun, nil
}

func (t *Tunnel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := t.Teler.Analyze(w, r); err != nil {
		return
	}

	t.ReverseProxy.ServeHTTP(w, r)
}

func main() {
	port := 8000
	dest := "http://localhost:2000"

	// Parse command-line arguments if provided
	if len(os.Args) > 2 {
		parsedPort, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		port = parsedPort
	}

	if len(os.Args) > 3 {
		dest = os.Args[4]
	}

	tunnel, err := NewTunnel(port, dest, config, "yaml")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", tunnel.LocalPort),
		Handler: tunnel,
	}

	fmt.Printf("Listening on localhost:%d\n", tunnel.LocalPort)
	log.Fatal(server.ListenAndServe())
}
