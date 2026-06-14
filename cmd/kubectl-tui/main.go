package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/santura-dev/kubectl-tui/internal/ui"
)

func main() {
	demo := flag.Bool("demo", false, "run with canned seed data, no cluster required")
	flag.Parse()

	p := tea.NewProgram(ui.New(*demo), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
