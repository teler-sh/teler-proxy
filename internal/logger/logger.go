package logger

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/mattn/go-colorable"
)

func New() *log.Logger {
	log.ValueStyles["id"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	log.ValueStyles["threat"] = lipgloss.NewStyle().Foreground(lipgloss.Color("192"))
	log.ValueStyles["request"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	log.ValueStyles["options"] = log.ValueStyles["threat"]

	return log.NewWithOptions(
		colorable.NewColorableStderr(),
		log.Options{
			ReportTimestamp: true,
			TimeFormat:      time.DateTime,
		},
	)
}
