package app

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nclark/wut/internal/effects"
	"github.com/nclark/wut/internal/utils"
	"github.com/nclark/wut/internal/wutang"
)

type TickMsg time.Time

func TickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		TickCmd(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Progress.Width = msg.Width - 10

		// Initialize matrix rain
		if len(m.MatrixRains) == 0 {
			for i := 0; i < m.Width/3; i++ {
				m.MatrixRains = append(m.MatrixRains, effects.MatrixRain{
					X:         rand.Intn(m.Width),
					Chars:     effects.GenerateMatrixChars(rand.Intn(15) + 5),
					Speed:     rand.Float64()*0.5 + 0.2,
					Color:     lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", 0, rand.Intn(255), 0)),
					Intensity: rand.Float64(),
				})
			}
		}

	case tea.KeyMsg:
		switch m.State {
		case MenuState:
			switch msg.String() {
			case "up", "k":
				m.MenuSelection = utils.Max(0, m.MenuSelection-1)
			case "down", "j":
				m.MenuSelection = utils.Min(len(m.MenuItems)-1, m.MenuSelection+1)
			case "enter":
				return m.handleMenuSelection()
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case CountdownState:
			if m.InputMode {
				return m.handleCustomTimeInput(msg)
			}
			return m.handleCountdownKeys(msg)

		case FinishedState:
			switch msg.String() {
			case "enter", "space":
				return InitialModel(), nil
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}

	case TickMsg:
		return m.handleTick()
	}

	return m, nil
}

func (m Model) handleMenuSelection() (tea.Model, tea.Cmd) {
	switch m.MenuSelection {
	case 0: // 15 minutes
		m.TotalTime = 900
		m.TimeRemaining = 900
		m.State = CountdownState
		return m, m.spawnInitialEffects()
	case 1: // 5 minutes
		m.TotalTime = 300
		m.TimeRemaining = 300
		m.State = CountdownState
		return m, m.spawnInitialEffects()
	case 2: // 1 minute
		m.TotalTime = 60
		m.TimeRemaining = 60
		m.State = CountdownState
		return m, m.spawnInitialEffects()
	case 3: // 30 seconds - For The Children
		m.TotalTime = 30
		m.TimeRemaining = 30
		m.State = CountdownState
		return m, m.spawnInitialEffects()
	case 4: // 15 seconds - Protect Ya Neck
		m.TotalTime = 15
		m.TimeRemaining = 15
		m.State = CountdownState
		return m, m.spawnInitialEffects()
	case 5: // Custom time
		m.InputMode = true
		m.CustomTimeInput = ""
		return m, nil
	case 6: // Member mode
		// Cycle through members
		currentIndex := 0
		for i, member := range wutang.Members {
			if member == m.CurrentMember {
				currentIndex = i
				break
			}
		}
		nextIndex := (currentIndex + 1) % len(wutang.Members)
		m.CurrentMember = wutang.Members[nextIndex]
		m.MenuItems[6] = "🐉 MEMBER MODE: " + m.CurrentMember + " 🐉"
		return m, nil
	case 7: // Effects menu (placeholder)
		return m, nil
	case 8: // Exit
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) handleCustomTimeInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if minutes, err := strconv.Atoi(m.CustomTimeInput); err == nil && minutes > 0 {
			m.TotalTime = minutes * 60
			m.TimeRemaining = minutes * 60
			m.State = CountdownState
			m.InputMode = false
			return m, m.spawnInitialEffects()
		}
		m.InputMode = false
		return m, nil
	case "escape":
		m.InputMode = false
		return m, nil
	case "backspace":
		if len(m.CustomTimeInput) > 0 {
			m.CustomTimeInput = m.CustomTimeInput[:len(m.CustomTimeInput)-1]
		}
	default:
		if len(msg.String()) == 1 && msg.String() >= "0" && msg.String() <= "9" {
			m.CustomTimeInput += msg.String()
		}
	}
	return m, nil
}

func (m Model) handleCountdownKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "escape":
		m.State = MenuState
		return m, nil
	case "q", "ctrl+c":
		return m, tea.Quit
	case "g":
		m.GlitchEffect = !m.GlitchEffect
	case "f":
		m.FireworksMode = !m.FireworksMode
	case "e":
		m.EmojiRain = !m.EmojiRain
	case "w":
		m.spawnWuTangLogo()
	case "s":
		m.SpinningText = !m.SpinningText
	case "t":
		m.StrobeEffect = !m.StrobeEffect
	case "m":
		// Switch member
		currentIndex := 0
		for i, member := range wutang.Members {
			if member == m.CurrentMember {
				currentIndex = i
				break
			}
		}
		nextIndex := (currentIndex + 1) % len(wutang.Members)
		m.CurrentMember = wutang.Members[nextIndex]
		m.TriggerScreenShake(10)
	case "space":
		m.spawnExplosion()
		m.TriggerScreenShake(15)
	}
	return m, nil
}

func (m Model) handleTick() (tea.Model, tea.Cmd) {
	m.Frame++
	
	// Update screen shake
	if m.ScreenShake > 0 {
		m.ScreenShake--
		if m.ScreenShake > 0 {
			m.ScreenShakeX = rand.Intn(m.ScreenShake*2) - m.ScreenShake
			m.ScreenShakeY = rand.Intn(m.ScreenShake*2) - m.ScreenShake
		}
	}

	// Update beat pulse
	m.BeatPulse = math.Sin(float64(m.Frame) * 0.1)

	if m.State == CountdownState {
		// Countdown logic
		if m.Frame%10 == 0 { // Reduce from every frame to every 10 frames (1 second)
			m.TimeRemaining--
			if m.TimeRemaining <= 0 {
				m.State = FinishedState
				return m, m.createMassiveExplosion()
			}
		}

		// Update all effects
		m.updateAllEffects()

		// Spawn periodic effects
		m.spawnPeriodicEffects()
	}

	return m, TickCmd()
}

func (m *Model) updateAllEffects() {
	// Update particles
	m.Particles = effects.UpdateParticles(m.Particles, m.Width, m.Height)
	
	// Update quotes and collect new particles
	var newParticles []effects.Particle
	m.Quotes, newParticles = effects.UpdateQuotes(m.Quotes, m.Width, m.Height)
	m.Particles = append(m.Particles, newParticles...)
	
	// Update Wu-Tang logos
	m.WuTangLogos = effects.UpdateWuTangLogos(m.WuTangLogos, m.Width, m.Height)
	
	// Update matrix rain
	m.MatrixRains = effects.UpdateMatrixRain(m.MatrixRains, m.Width, m.Height)
	
	// Update explosions
	m.Explosions = effects.UpdateExplosions(m.Explosions, m.Width, m.Height)
}

func (m *Model) spawnPeriodicEffects() {
	// Spawn quotes even more frequently - 25% more
	if rand.Intn(23) == 0 { // Increased from 30 to 23 for 25% more quotes
		m.spawnQuote()
	}

	// Spawn particles when fireworks mode is on
	if m.FireworksMode && rand.Intn(50) == 0 { // Less frequent
		m.spawnParticles()
	}

	// Spawn emoji rain
	if m.EmojiRain && rand.Intn(30) == 0 { // Less frequent
		m.spawnEmojiRain()
	}

	// Auto-spawn Wu-Tang logos
	if m.AutoWuLogos && rand.Intn(400) == 0 { // Less frequent
		m.spawnWuTangLogo()
	}

	// Random explosions
	if rand.Intn(800) == 0 { // Less frequent
		m.spawnExplosion()
		m.TriggerScreenShake(8)
	}
}

func (m *Model) spawnInitialEffects() tea.Cmd {
	// Clear any existing effects
	m.Particles = m.Particles[:0]
	m.Quotes = m.Quotes[:0]
	m.Explosions = m.Explosions[:0]
	m.WuTangLogos = m.WuTangLogos[:0]

	// Spawn an ASSLOAD of initial quotes for immediate Wu-Tang chaos
	for i := 0; i < 20; i++ {
		m.spawnQuote()
	}

	// Spawn some initial particles
	if m.FireworksMode {
		m.spawnParticles()
	}

	// Spawn initial Wu-Tang logo
	if m.AutoWuLogos {
		m.spawnWuTangLogo()
	}

	return nil
}

func (m *Model) spawnQuote() {
	quote := effects.SpawnQuote(m.Width, m.Height, m.Frame)
	m.Quotes = append(m.Quotes, quote)
}

func (m *Model) spawnParticles() {
	particles := effects.SpawnParticles(m.Width, m.Height, m.Frame, 5)
	m.Particles = append(m.Particles, particles...)
}

func (m *Model) spawnExplosion() {
	explosion := effects.SpawnExplosion(m.Width, m.Height, m.Frame)
	m.Explosions = append(m.Explosions, explosion)
}

func (m *Model) createMassiveExplosion() tea.Cmd {
	// Clear existing effects
	m.Particles = m.Particles[:0]
	m.Quotes = m.Quotes[:0]
	m.Explosions = m.Explosions[:0]

	// Create massive explosion
	particles := effects.CreateMassiveExplosion(m.Width, m.Height)
	m.Particles = append(m.Particles, particles...)
	
	// Massive screen shake
	m.TriggerScreenShake(50)

	return nil
}

func (m *Model) spawnWuTangLogo() {
	logo := effects.SpawnWuTangLogo(m.Width, m.Height, m.Frame)
	m.WuTangLogos = append(m.WuTangLogos, logo)
}

func (m *Model) spawnEmojiRain() {
	particles := effects.SpawnEmojiRain(m.Width, m.Frame)
	m.Particles = append(m.Particles, particles...)
}