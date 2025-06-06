package app

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nclark/wut/internal/effects"
	"github.com/nclark/wut/internal/ui"
	"github.com/nclark/wut/internal/utils"
	"github.com/nclark/wut/internal/wutang"
)

func (m Model) View() string {
	switch m.State {
	case MenuState:
		return m.menuView()
	case CountdownState:
		return m.countdownView()
	case FinishedState:
		return m.finishedView()
	}
	return ""
}

// menuView renders the main menu interface
func (m Model) menuView() string {
	var s strings.Builder

	// Matrix rain background
	matrix := m.renderMatrixRain()
	s.WriteString(matrix)

	// Apply screen shake
	if m.ScreenShake > 0 {
		s.WriteString(strings.Repeat("\n", m.ScreenShakeY))
		s.WriteString(strings.Repeat(" ", m.ScreenShakeX))
	}

	// Title with spinning effect
	title := "🐉 WU-TANG COUNTDOWN CHAMBER 🐉"
	if m.SpinningText && m.Frame%20 < 10 {
		title = "🐲 CHAMBER COUNTDOWN TANG-WU 🐲"
	}
	titleRendered := ui.TitleStyle.Render(title)
	s.WriteString(lipgloss.Place(m.Width, 5, lipgloss.Center, lipgloss.Center, titleRendered))
	s.WriteString("\n\n")

	// Subtitle
	subtitle := lipgloss.NewStyle().
		Foreground(effects.GetRainbowColor(float64(m.Frame) * 0.05)).
		Bold(true).
		Render("Enter the 36 Chambers of Time Management")
	s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, subtitle))
	s.WriteString("\n\n")

	// Menu items
	for i, item := range m.MenuItems {
		style := ui.MenuStyle
		if i == m.MenuSelection {
			style = ui.SelectedMenuStyle
			item = "▶ " + item + " ◀"
		} else {
			item = "  " + item + "  "
		}

		// Show custom time input
		if i == 5 && m.InputMode {
			item = "  🎯 ENTER MINUTES: " + m.CustomTimeInput + "_  "
		}

		rendered := style.Render(item)
		s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, rendered))
		s.WriteString("\n")
	}

	// Effects status
	s.WriteString("\n")
	effectsStatus := fmt.Sprintf("🎭 Emoji Rain: %v | 💫 Spinning: %v | ⚡ Strobe: %v | 🔥 Wu-Logos: %v",
		m.EmojiRain, m.SpinningText, m.StrobeEffect, m.AutoWuLogos)
	effectsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, effectsStyle.Render(effectsStatus)))

	s.WriteString("\n\n")
	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render("Use ↑↓ or j/k to navigate • Enter to select • q to quit")
	s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, instructions))

	return s.String()
}

// countdownView renders the countdown timer view with all effects
func (m Model) countdownView() string {
	// Apply strobe effect
	if m.StrobeEffect && m.Frame%10 < 2 {
		return strings.Repeat(" ", m.Width*m.Height)
	}

	// Create a simple grid to track quote positions
	lines := make([]string, m.Height)
	for i := range lines {
		lines[i] = strings.Repeat(" ", m.Width)
	}

	// Apply screen shake offset
	offsetX := m.ScreenShakeX
	offsetY := m.ScreenShakeY

	// Place Wu-Tang logos
	for _, logo := range m.WuTangLogos {
		x := int(logo.X) + offsetX
		y := int(logo.Y) + offsetY
		
		logoLines := strings.Split(logo.Text, "\n")
		for lineIdx, logoLine := range logoLines {
			lineY := y + lineIdx
			if lineY >= 0 && lineY < len(lines) && x >= 0 && x < m.Width-22 {
				style := lipgloss.NewStyle().Foreground(logo.Color).Bold(true)
				styledLine := style.Render(logoLine)
				if lineY < len(lines) {
					prefix := strings.Repeat(" ", x)
					lines[lineY] = prefix + styledLine
				}
			}
		}
	}

	// Place floating quotes at their positions
	for _, q := range m.Quotes {
		x := int(q.X+q.Wobble*5) + offsetX
		y := int(q.Y) + offsetY

		// Keep within bounds
		if x < 0 {
			x = 0
		}
		if x > m.Width-len(q.Text) {
			x = m.Width - len(q.Text)
		}
		if y < 0 {
			y = 0
		}
		if y >= len(lines)-8 {
			y = len(lines) - 9
		}

		// Style the quote
		alpha := float64(q.Life) / float64(q.MaxLife)
		style := lipgloss.NewStyle().Foreground(q.Color)
		if alpha < 0.3 {
			style = style.Faint(true)
		} else {
			style = style.Bold(true)
		}

		// Apply spinning text effect
		text := q.Text
		if m.SpinningText && rand.Float64() < 0.4 { // Increased from 0.1 to 0.4 for more visibility
			// Reverse random words
			words := strings.Split(text, " ")
			if len(words) > 1 && rand.Float64() < 0.6 { // Increased from 0.3 to 0.6
				idx := rand.Intn(len(words))
				words[idx] = utils.ReverseString(words[idx])
			}
			text = strings.Join(words, " ")
		}

		// Insert the styled quote into the line
		if y >= 0 && y < len(lines) {
			styledText := style.Render(text)
			if x < m.Width {
				prefix := strings.Repeat(" ", x)
				if len(prefix)+len(text) < m.Width {
					lines[y] = prefix + styledText
				}
			}
		}
	}

	// Calculate timer position (center of screen)
	timerY := m.Height / 2
	if timerY >= len(lines) {
		timerY = len(lines) - 4
	}

	// Build timer with member theme
	minutes := m.TimeRemaining / 60
	seconds := m.TimeRemaining % 60
	timerText := fmt.Sprintf("⏰ %02d:%02d ⏰", minutes, seconds)
	
	if m.MemberMode {
		timerText = fmt.Sprintf("🎤 %s | %02d:%02d 🎤", m.CurrentMember, minutes, seconds)
	}

	// Apply effects to timer
	if m.GlitchEffect && m.Frame%10 < 3 {
		timerText = ui.GlitchStyle.Render(timerText)
	} else {
		timerStyle := ui.TimerStyle.Foreground(effects.GetRainbowColor(float64(m.Frame) * 0.05))
		timerText = timerStyle.Render(timerText)
	}

	// Center the timer line
	centeredTimer := lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, timerText)
	if timerY >= 0 && timerY < len(lines) {
		lines[timerY] = centeredTimer
	}

	// Add particles line if fireworks mode OR emoji rain
	if (m.FireworksMode || m.EmojiRain) && len(m.Particles) > 0 && timerY+2 < len(lines) {
		particleStr := ""
		for i, p := range m.Particles {
			if i > 20 {
				break
			}
			
			if p.Emoji != "" {
				// Emoji particles
				particleStr += p.Emoji + " "
			} else {
				// Regular particles with spin effect
				symbol := p.Symbol
				if m.SpinningText && rand.Float64() < 0.7 { // Much more visible
					symbols := []rune{'/', '-', '\\', '|'}
					symbol = symbols[int(p.Spin)%4]
				}
				styled := lipgloss.NewStyle().Foreground(p.Color).Bold(true).Render(string(symbol))
				particleStr += styled + " "
			}
		}
		centeredParticles := lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, particleStr)
		lines[timerY+2] = centeredParticles
	}

	// Add progress bar near bottom - CENTERED with pulse effect
	if len(lines) >= 3 {
		progressPercent := float64(m.TotalTime-m.TimeRemaining) / float64(m.TotalTime)
		
		// Pulse the progress bar width
		pulseWidth := m.Progress.Width
		if m.BeatPulse > 0 {
			pulseWidth = int(float64(pulseWidth) * (1 + m.BeatPulse*0.1))
		}
		
		// Create a temporary progress model to avoid modifying the original
		progress := m.Progress
		progress.Width = pulseWidth
		progressBar := progress.ViewAs(progressPercent)
		centeredProgress := lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, progressBar)
		lines[len(lines)-3] = centeredProgress
	}

	// Add status line at bottom - CENTERED
	if len(lines) >= 1 {
		statusLine := ""
		if m.GlitchEffect {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true).Render("[GLITCH] ")
		}
		if m.FireworksMode {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render("[FIREWORKS] ")
		}
		if m.EmojiRain {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true).Render("[EMOJI RAIN] ")
		}
		if m.StrobeEffect {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true).Render("[STROBE] ")
		}
		if m.SpinningText {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#00FFFF")).Bold(true).Render("[SPINNING] ")
		}
		if m.AutoWuLogos {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render("[WU-LOGOS] ")
		}
		
		statusLine += "\n"
		statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).
			Render("g:glitch f:fireworks e:emoji w:wu-logo s:spin t:strobe m:member space:explode")

		centeredStatus := lipgloss.Place(m.Width, 2, lipgloss.Center, lipgloss.Center, statusLine)
		lines[len(lines)-2] = centeredStatus[:len(centeredStatus)-1] // Remove trailing newline
		lines[len(lines)-1] = centeredStatus[len(centeredStatus)-1:]
	}

	// Join all lines
	return strings.Join(lines, "\n")
}

// finishedView renders the completion screen
func (m Model) finishedView() string {
	var s strings.Builder

	// Massive explosion overlay
	explosionText := m.renderMassiveExplosion()
	s.WriteString(explosionText)

	// Apply screen shake to finished message
	offsetX := m.ScreenShakeX
	offsetY := m.ScreenShakeY
	
	if offsetY > 0 {
		s.WriteString(strings.Repeat("\n", offsetY))
	}
	if offsetX > 0 {
		s.WriteString(strings.Repeat(" ", offsetX))
	}

	// Flashing completion message
	var message string
	if m.Frame%20 < 10 {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#FFFF00")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("🔥🔥🔥 TIME'S UP! 🔥🔥🔥\nWU-TANG FOREVER!\n36 CHAMBERS COMPLETE\n" + m.CurrentMember + " APPROVES!")
	} else {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Background(lipgloss.Color("#FF0000")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("💀💀💀 SHAOLIN STYLE! 💀💀💀\nBRING DA RUCKUS!\nKILLA BEEZ ON ATTACK\n" + m.CurrentMember + " IN THE HOUSE!")
	}

	s.WriteString(lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, message))

	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Render("Press Enter to return to menu • 'q' to quit")
	s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Bottom, instructions))

	return s.String()
}

// renderMatrixRain renders matrix background effect
func (m Model) renderMatrixRain() string {
	// Simple matrix background
	return ""
}

// renderMassiveExplosion renders the final explosion animation
func (m Model) renderMassiveExplosion() string {
	var s strings.Builder

	// Create epic explosion pattern with emojis and colors
	for i := 0; i < 15; i++ {
		line := ""
		for j := 0; j < m.Width/5; j++ {
			if rand.Float64() < 0.6 {
				// Mix of emojis and symbols
				if rand.Float64() < 0.5 {
					emoji := wutang.CrazyEmojis[rand.Intn(len(wutang.CrazyEmojis))]
					line += emoji + " "
				} else {
					symbols := []string{"★", "✦", "◆", "●", "▲", "♦", "*", "+", "x"}
					symbol := symbols[rand.Intn(len(symbols))]
					color := effects.GetRainbowColor(float64(i*j) * 0.1)
					line += lipgloss.NewStyle().Foreground(color).Bold(true).Render(symbol + " ")
				}
			} else {
				line += "  "
			}
		}
		s.WriteString(lipgloss.Place(m.Width, 1, lipgloss.Center, lipgloss.Center, line))
		s.WriteString("\n")
	}

	return s.String()
}