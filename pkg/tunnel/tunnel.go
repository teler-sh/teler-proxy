/*
Package tunnel provides functionality for creating HTTP tunnels
and reverse proxies with [teler] WAF capabilities.

The main components of this package include the [Tunnel] type,
which represents a tunneling configuration, and related functions
and methods for tunnel setup and HTTP request handling.

To create a new tunnel, use the [NewTunnel] function, specifying
the local port, destination address, and optional configuration
parameters. The [Tunnel] type also provides the [Tunnel.ServeHTTP] method
for handling incoming HTTP requests and proxying them to the
destination, analyzing the incoming HTTP request from threats using
the [teler.Teler] middleware.

Additional configuration options can be loaded from YAML or JSON
files, allowing for customizing the [teler] WAF behavior.
*/
package tunnel

import (
	"io"
	"strings"

	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/teler-sh/teler-proxy/common"
	"github.com/teler-sh/teler-waf"
	"github.com/teler-sh/teler-waf/option"
)

type Tunnel struct {
	*httputil.ReverseProxy
	*teler.Teler

	Options teler.Options
}

// NewTunnel creates a new [Tunnel] instance for proxying HTTP traffic.
//
// Parameters:
//   - `port`: The local port on which the tunnel will listen for incoming requests.
//   - `dest`: The destination address to which incoming requests will be forwarded.
//   - `cfgPath`: The path to a configuration file for additional tunnel options.
//   - `optFormat`: The format of the configuration file ("yaml" or "json").
//   - `writer`: An optional [io.Writer] where tunnel log output will be written.
//     Pass nil to use default [teler] logging only.
//
// Please be aware that when you pass a custom `writer`, the [teler.Options.NoStderr]
// option value will be forcibly set to `true`, regardless of the `no_stderr` value
// that might be loaded from additional configuration options.
func NewTunnel(port int, dest, cfgPath, optFormat string, writer io.Writer) (*Tunnel, error) {
	if dest == "" {
		return nil, common.ErrDestAddressEmpty
	}

	// NOTE(dwisiswant0): should we accept the input `dest` parameter
	// as pointer of url.URL directly instead of string?
	destURL, err := url.Parse(dest)
	if err != nil {
		return nil, err
	}

	var opt teler.Options

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

		tun.Options = opt
	}

	if writer != nil {
		opt.LogWriter = writer
		opt.NoStderr = true
	}

	tun.Teler = teler.New(opt)

	return tun, nil
}

// ServeHTTP is a method of the [Tunnel] type, which allows
// the [Tunnel] to implement the [http.Handler] interface.
//
// This method forwards the incoming HTTP request to the
// [httputil.ReverseProxy.ServeHTTP] method, while also
// analyzing the incoming HTTP request from threats using
// the [teler.Teler] middleware.
func (t *Tunnel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Teler.HandlerFuncWithNext(w, r, t.ReverseProxy.ServeHTTP)
}
