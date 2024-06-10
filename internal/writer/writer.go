package writer

import (
	"encoding/json"

	"github.com/charmbracelet/log"
	"github.com/teler-sh/teler-proxy/internal/logger"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type logWriter struct {
	*log.Logger
}

type data = map[string]any

func New() *logWriter {
	w := new(logWriter)
	w.Logger = logger.New().WithPrefix("teler-waf")

	return w
}

func (w *logWriter) Write(p []byte) (n int, err error) {
	var d data

	n = len(p)

	err = json.Unmarshal(p, &d)
	if err != nil {
		return 0, err
	}

	err = w.write(d)

	return
}

func (w *logWriter) write(d data) error {
	switch level := d["level"].(string); level {
	case "debug":
		w.writeDebug(d)
	// case "info":
	// 	w.writeInfo(d)
	case "warn":
		w.writeWarn(d)
	case "error":
		w.writeError(d)
	case "fatal":
		w.writeFatal(d)
	}

	return nil
}

func (w *logWriter) writeDebug(d data) {
	w.Debug(d["msg"])
}

func (w *logWriter) writeInfo(d data) {
	// if opt, ok := d["options"].(data); ok {
	// 	o, err := json.Marshal(opt)
	// 	if err != nil {
	// 		return
	// 	}

	// 	w.Info(d["msg"], "options", string(o))
	// } else {
	w.Info(d["msg"])
	// }
}

func (w *logWriter) writeWarn(d data) {
	if req, ok := d["request"].(data); ok {
		r, err := json.Marshal(req)
		if err != nil {
			return
		}

		w.Warn(d["msg"],
			"id", d["id"],
			"threat", d["category"],
			"request", string(r),
		)
	} else {
		w.Warn(d["msg"])
	}
}

func (w *logWriter) writeError(d data) {
	w.Error(d["msg"],
		"source", d["caller"],
	)
}

func (w *logWriter) writeFatal(d data) {
	w.writeError(d)
}
