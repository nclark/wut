package main

import (
	"fmt"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nclark/wut/internal/app"
)

var version = "dev" // Will be set by build process

func main() {
	rand.Seed(time.Now().UnixNano())

	p := tea.NewProgram(
		app.InitialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}