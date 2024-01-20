package logger

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/mattn/go-colorable"
)

func New() *log.Logger {
	logger := log.NewWithOptions(
		colorable.NewColorableStderr(),
		log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.DateTime,
		},
	)

	styles := log.DefaultStyles()
	styles.Values["id"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	styles.Values["threat"] = lipgloss.NewStyle().Foreground(lipgloss.Color("192"))
	styles.Values["request"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["options"] = styles.Values["threat"]

	logger.SetStyles(styles)

	log.TimestampKey = "ts"

	return logger
}
