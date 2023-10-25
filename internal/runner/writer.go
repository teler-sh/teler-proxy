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

	w.Logger.With("ts", d["ts"], "msg", d["msg"])
	w.write(d)

	return
}

func (w *logWriter) write(d data) {
	switch level := d["level"].(string); level {
	case "debug":
		w.writeDebug(d)
	case "info":
		w.writeInfo(d)
	case "warn":
		w.writeWarn(d)
	case "error":
		w.writeError(d)
	case "fatal":
		w.writeFatal(d)
	}
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

func (w *logWriter) writeWarn(d data) {
	w.Warn(d["msg"],
		"ts", d["ts"],
		"id", d["id"],
		"threat", d["category"],
		"request", d["request"],
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
