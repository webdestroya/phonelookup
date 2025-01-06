package utils

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	ErrorStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("204"))
	WarningStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("192"))
)

func PrintError(err any) {
	dispStr := ""
	switch v := err.(type) {
	case error:
		dispStr = v.Error()
	case string:
		dispStr = v
	case nil:
		return
	default:
		dispStr = fmt.Sprint(v)
	}

	fmt.Fprintln(os.Stderr, ErrorStyle.Render("Error:", dispStr))
}
