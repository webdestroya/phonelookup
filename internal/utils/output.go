package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.ANSIColor(31)).SetString("Error:")
)

func PrintError(err any) {
	switch v := err.(type) {
	case error:
		fmt.Fprintln(os.Stderr, errorStyle.Render(v.Error()))
	case string:
		fmt.Fprintln(os.Stderr, errorStyle.Render(v))
	default:
		fmt.Fprintln(os.Stderr, errorStyle.Render(fmt.Sprint(v)))
	}
}
