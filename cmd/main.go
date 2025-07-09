package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/winslowb/onebill-kanban/data"
	"github.com/winslowb/onebill-kanban/ui"
)

func main() {
	if err := data.InitStore(); err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}

	model, err := ui.NewBoardModel()
	if err != nil {
		log.Fatalf("failed to create board model: %v", err)
	}

	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		log.Fatalf("error running program: %v", err)
		os.Exit(1)
	}
}

