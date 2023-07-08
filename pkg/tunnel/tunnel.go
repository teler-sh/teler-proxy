package tunnel

import (
	"strings"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kitabisa/teler-proxy/common"
	"github.com/kitabisa/teler-waf"
	"github.com/kitabisa/teler-waf/option"
)

type Tunnel struct {
	*teler.Teler
	ReverseProxy *httputil.ReverseProxy
}

func NewTunnel(port int, dest, cfgPath, optFormat string) (*Tunnel, error) {
	dest = "http://" + dest
	destURL, err := url.Parse(dest)
	if err != nil {
		return nil, err
	}

	var opt teler.Options

	if dest == "" {
		return nil, common.ErrDestAddressEmpty
	}

	tun := &Tunnel{}
	tun.ReverseProxy = httputil.NewSingleHostReverseProxy(destURL)

	if cfgPath != "" {
		switch strings.ToLower(optFormat) {
		case "yaml":
			opt, err = option.LoadFromYAMLFile(cfgPath)
		case "json":
			opt, err = option.LoadFromJSONFile(cfgPath)
		case "":
			return nil, common.ErrCfgFileFormatUnd
		default:
			return nil, common.ErrCfgFileFormatInv
		}

		if err != nil {
			return nil, err
		}

		tun.Teler = teler.New(opt)
	} else {
		tun.Teler = teler.New()
	}

	return tun, nil
}

func (t *Tunnel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Teler.HandlerFuncWithNext(w, r, t.ReverseProxy.ServeHTTP)
}
