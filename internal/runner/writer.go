package runner

import (
	"encoding/json"

	"github.com/charmbracelet/log"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type logWriter struct {
	*log.Logger
}

type data map[string]any

func (r *Runner) NewWriter() *logWriter {
	w := new(logWriter)
	w.Logger = r.Options.Logger.WithPrefix("teler-waf")

	return w
}

func (w *logWriter) Write(p []byte) (n int, err error) {
	var d data

	n = len(p)

	err = json.Unmarshal(p, &d)
	if err != nil {
		return 0, err
	}

	w.Logger = w.With("ts", d["ts"])

	err = w.write(d)

	return
}

func (w *logWriter) write(d data) error {
	switch level := d["level"].(string); level {
	case "debug":
		w.writeDebug(d)
	case "info":
		w.writeInfo(d)
	case "warn":
		r, err := json.Marshal(d["request"])
		if err != nil {
			return err
		}

		w.writeWarn(d, r)
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
	if opt, ok := d["options"].(data); ok {
		w.Info(d["msg"],
			"options", opt,
		)
	}
}

func (w *logWriter) writeWarn(d data, r []byte) {
	w.Warn(d["msg"],
		"id", d["id"],
		"threat", d["category"],
		"request", string(r),
	)
}

func (w *logWriter) writeError(d data) {
	w.Error(d["msg"],
		"source", d["caller"],
	)
}

func (w *logWriter) writeFatal(d data) {
	w.writeError(d)
}
