package tunnel

import (
	"strings"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kitabisa/teler-waf"
	"github.com/kitabisa/teler-waf/option"
)

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
			return nil, errTelerOptFormatUnd
		default:
			return nil, errTelerOptFormatInv
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
