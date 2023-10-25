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

func (w *logWriter) Write(p []byte) (n int, err error) {
	var d data

	n = len(p)

	err = json.Unmarshal(p, &d)
	if err != nil {
		return 0, err
	}

	logger := w.WithPrefix("teler-waf")
	w.Logger = logger
	w.Logger.With("ts", d["ts"])

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
	w.Debug(d["msg"], "ts", d["ts"])
}

func (w *logWriter) writeInfo(d data) {
	if opt, ok := d["options"].(data); ok {
		w.Info(d["msg"],
			"ts", d["ts"],
			"options", opt,
		)
	}
}

func (w *logWriter) writeWarn(d data, r []byte) {
	w.Warn(d["msg"],
		"ts", d["ts"],
		"id", d["id"],
		"threat", d["category"],
		"request", string(r),
	)
}

func (w *logWriter) writeError(d data) {
	w.Error(d["msg"],
		"ts", d["ts"],
		"source", d["caller"],
	)
}

func (w *logWriter) writeFatal(d data) {
	w.writeError(d)
}
