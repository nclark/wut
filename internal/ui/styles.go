package ui

import "github.com/charmbracelet/lipgloss"

// Styles
var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#FF4500")).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#FFD700"))

	MenuStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Bold(true).
			Padding(0, 2)

	SelectedMenuStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#FFD700")).
				Bold(true).
				Padding(0, 2)

	TimerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#000000")).
			Bold(true).
			Padding(1, 3).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFD700")).
			Align(lipgloss.Center)

	QuoteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	ExplosionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF4500")).
				Bold(true)

	GlitchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF00FF")).
			Background(lipgloss.Color("#00FF00")).
			Bold(true)
)