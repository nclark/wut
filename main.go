package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// States
type state int

const (
	menuState state = iota
	countdownState
	finishedState
)

// Wu-Tang quotes
var wuTangQuotes = []string{
	"I bomb atomically",
	"Wu-Tang is for the children",
	"WHAT YALL THOUGH YA WASN'T GONNA SEE ME",
	"Cash rules everything around me",
	"Bring da ruckus",
	"Wu-Tang Clan ain't nuthing ta f' wit",
	"Protect ya neck",
	"36 chambers of death",
	"Raw I'mma give it to ya",
	"Enter the Wu-Tang",
	"Shaolin shadowboxing",
	"Method Man on the left",
	"Tiger style",
	"Liquid swords",
	"Killer bees on the swarm",
	"Triumph",
	"Wu wear",
	"Staten Island stand up",
	"C.R.E.A.M.",
	"Diversify yo bonds",
}

// Particle system
type Particle struct {
	X, Y    float64
	VX, VY  float64
	Life    int
	MaxLife int
	Symbol  rune
	Color   lipgloss.Color
	Size    int
}

type FloatingQuote struct {
	Text    string
	X, Y    float64
	VX, VY  float64
	Life    int
	MaxLife int
	Color   lipgloss.Color
	Wobble  float64
	Phase   float64
}

type MatrixRain struct {
	X         int
	Chars     []rune
	Speed     float64
	Color     lipgloss.Color
	Intensity float64
}

type Explosion struct {
	X, Y      float64
	Particles []Particle
	Life      int
	MaxLife   int
}

type model struct {
	state         state
	width         int
	height        int
	timeRemaining int
	totalTime     int
	progress      progress.Model
	particles     []Particle
	quotes        []FloatingQuote
	matrixRains   []MatrixRain
	explosions    []Explosion
	frame         int
	menuSelection int
	glitchEffect  bool
	fireworksMode bool
	menuItems     []string
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Background(lipgloss.Color("#FF4500")).
			Bold(true).
			Padding(1, 2).
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#FFD700"))

	menuStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Bold(true).
			Padding(0, 2)

	selectedMenuStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#FFD700")).
				Bold(true).
				Padding(0, 2)

	timerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#000000")).
			Bold(true).
			Padding(1, 3).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FFD700")).
			Align(lipgloss.Center)

	quoteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			Bold(true)

	explosionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF4500")).
			Bold(true)

	glitchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF00FF")).
			Background(lipgloss.Color("#00FF00")).
			Bold(true)
)

func initialModel() model {
	p := progress.New(progress.WithDefaultGradient())
	return model{
		state:         menuState,
		totalTime:     900, // 15 minutes
		timeRemaining: 900,
		progress:      p,
		particles:     make([]Particle, 0),
		quotes:        make([]FloatingQuote, 0),
		matrixRains:   make([]MatrixRain, 0),
		explosions:    make([]Explosion, 0),
		menuSelection: 0,
		menuItems: []string{
			"🔥 15 MINUTE WU-TANG COUNTDOWN 🔥",
			"⚡ 5 MINUTE SHAOLIN SPECIAL ⚡",
			"💀 1 MINUTE DEATH CHAMBER 💀",
			"🎯 CUSTOM TIME (ENTER MINUTES)",
			"❌ EXIT THE CHAMBER",
		},
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

type tickMsg time.Time

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = msg.Width - 10

		// Initialize matrix rain
		if len(m.matrixRains) == 0 {
			for i := 0; i < m.width/3; i++ {
				m.matrixRains = append(m.matrixRains, MatrixRain{
					X:         rand.Intn(m.width),
					Chars:     generateMatrixChars(rand.Intn(15) + 5),
					Speed:     rand.Float64()*0.5 + 0.2,
					Color:     lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", 0, rand.Intn(255), 0)),
					Intensity: rand.Float64(),
				})
			}
		}

	case tea.KeyMsg:
		switch m.state {
		case menuState:
			switch msg.String() {
			case "up", "k":
				m.menuSelection = (m.menuSelection - 1 + len(m.menuItems)) % len(m.menuItems)
			case "down", "j":
				m.menuSelection = (m.menuSelection + 1) % len(m.menuItems)
			case "enter":
				switch m.menuSelection {
				case 0:
					m.totalTime = 900
					m.timeRemaining = 900
				case 1:
					m.totalTime = 300
					m.timeRemaining = 300
				case 2:
					m.totalTime = 60
					m.timeRemaining = 60
				case 3:
					m.totalTime = 120 // Default 2 minutes for demo
					m.timeRemaining = 120
				case 4:
					return m, tea.Quit
				}
				m.state = countdownState
				m.fireworksMode = true
				return m, tea.Batch(tickCmd(), m.spawnInitialEffects())
			case "q", "ctrl+c":
				return m, tea.Quit
			}

		case countdownState:
			switch msg.String() {
			case "q", "esc":
				m.state = menuState
				m.particles = m.particles[:0]
				m.quotes = m.quotes[:0]
				m.explosions = m.explosions[:0]
			case "g":
				m.glitchEffect = !m.glitchEffect
			case "f":
				m.fireworksMode = !m.fireworksMode
				// Spawn extra particles when toggling fireworks on
				if m.fireworksMode {
					for i := 0; i < 20; i++ {
						m.spawnParticles()
					}
				}
			case "ctrl+c":
				return m, tea.Quit
			}

		case finishedState:
			switch msg.String() {
			case "r":
				m.state = menuState
				m.particles = m.particles[:0]
				m.quotes = m.quotes[:0]
				m.explosions = m.explosions[:0]
				m.frame = 0
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}

	case tickMsg:
		m.frame++

		if m.state == countdownState {
			// Update timer every second (20 frames at 50ms)
			if m.frame%20 == 0 && m.timeRemaining > 0 {
				m.timeRemaining--
				if m.timeRemaining == 0 {
					m.state = finishedState
					return m, m.createMassiveExplosion()
				}
			}

			// Spawn effects
			if m.frame%80 == 0 { // Spawn quotes less frequently (was 40)
				m.spawnQuote()
			}
			if m.frame%8 == 0 && m.fireworksMode { // More frequent particles in fireworks mode
				m.spawnParticles()
			} else if m.frame%20 == 0 { // Normal particle spawn rate
				m.spawnParticles()
			}
			if m.frame%50 == 0 {
				m.spawnExplosion()
			}
		}

		// Update all effects
		m.updateParticles()
		m.updateQuotes()
		m.updateMatrixRain()
		m.updateExplosions()

		return m, tickCmd()
	}

	return m, nil
}

func (m *model) spawnInitialEffects() tea.Cmd {
	// ALWAYS start with the legendary quote
	m.quotes = append(m.quotes, FloatingQuote{
		Text:    "WHAT YALL THOUGH YA WASN'T GONNA SEE ME",
		X:       float64(m.width/2 - 20), // Center it initially
		Y:       float64(m.height / 4),   // Upper area
		VX:      0.3,                     // Slow drift right
		VY:      0.1,                     // Slow drift down
		Life:    800,                     // Extra long life for the opening quote
		MaxLife: 800,
		Color:   lipgloss.Color("#FFD700"), // Gold for the legendary opening
		Phase:   0,
	})

	// Spawn initial other quotes and particles
	for i := 0; i < 3; i++ {
		m.spawnQuote()
	}
	for i := 0; i < 20; i++ {
		m.spawnParticles()
	}
	return nil
}

func (m *model) spawnQuote() {
	if len(m.quotes) > 15 {
		return
	}

	quote := wuTangQuotes[rand.Intn(len(wuTangQuotes))]
	colors := []lipgloss.Color{"#FFD700", "#FF4500", "#00FFFF", "#FF1493", "#32CD32", "#9400D3"}

	m.quotes = append(m.quotes, FloatingQuote{
		Text:    quote,
		X:       rand.Float64() * float64(m.width-len(quote)),
		Y:       rand.Float64() * float64(m.height-5),
		VX:      (rand.Float64() - 0.5) * 0.5, // Much slower X velocity
		VY:      (rand.Float64() - 0.5) * 0.3, // Much slower Y velocity
		Life:    400 + rand.Intn(600),         // Live longer (was 200-500, now 400-1000)
		MaxLife: 400 + rand.Intn(600),
		Color:   colors[rand.Intn(len(colors))], // Actually use the colors array
		Phase:   rand.Float64() * math.Pi * 2,
	})
}

func (m *model) spawnParticles() {
	symbols := []rune{'★', '✦', '✧', '◆', '◇', '●', '◉', '▲', '▼', '♦'}
	colors := []lipgloss.Color{"#FFD700", "#FF4500", "#00FFFF", "#FF1493", "#32CD32"}

	for i := 0; i < 3; i++ {
		m.particles = append(m.particles, Particle{
			X:       rand.Float64() * float64(m.width),
			Y:       rand.Float64() * float64(m.height),
			VX:      (rand.Float64() - 0.5) * 4,
			VY:      (rand.Float64() - 0.5) * 4,
			Life:    50 + rand.Intn(100),
			MaxLife: 50 + rand.Intn(100),
			Symbol:  symbols[rand.Intn(len(symbols))],
			Color:   colors[rand.Intn(len(colors))],
			Size:    1 + rand.Intn(3),
		})
	}
}

func (m *model) spawnExplosion() {
	if len(m.explosions) > 5 {
		return
	}

	explosion := Explosion{
		X:       rand.Float64() * float64(m.width),
		Y:       rand.Float64() * float64(m.height),
		Life:    60,
		MaxLife: 60,
	}

	// Create explosion particles
	for i := 0; i < 20; i++ {
		angle := float64(i) * math.Pi * 2 / 20
		speed := rand.Float64()*3 + 1
		explosion.Particles = append(explosion.Particles, Particle{
			X:       explosion.X,
			Y:       explosion.Y,
			VX:      math.Cos(angle) * speed,
			VY:      math.Sin(angle) * speed,
			Life:    30 + rand.Intn(30),
			MaxLife: 30 + rand.Intn(30),
			Symbol:  []rune{'*', '+', 'x', '◦'}[rand.Intn(4)],
			Color:   lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", 255, rand.Intn(255), 0)),
		})
	}

	m.explosions = append(m.explosions, explosion)
}

func (m *model) createMassiveExplosion() tea.Cmd {
	// Clear existing effects
	m.particles = m.particles[:0]
	m.quotes = m.quotes[:0]
	m.explosions = m.explosions[:0]

	// Create massive explosion
	centerX := float64(m.width) / 2
	centerY := float64(m.height) / 2

	for i := 0; i < 200; i++ {
		angle := rand.Float64() * math.Pi * 2
		speed := rand.Float64()*8 + 2
		m.particles = append(m.particles, Particle{
			X:       centerX,
			Y:       centerY,
			VX:      math.Cos(angle) * speed,
			VY:      math.Sin(angle) * speed,
			Life:    100 + rand.Intn(200),
			MaxLife: 100 + rand.Intn(200),
			Symbol:  []rune{'★', '✦', '◆', '●', '▲', '♦'}[rand.Intn(6)],
			Color:   lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", 255, rand.Intn(255), rand.Intn(255))),
			Size:    1 + rand.Intn(4),
		})
	}

	return nil
}

func generateMatrixChars(length int) []rune {
	chars := []rune("ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇﾍ012345789ZXCVBNMASDFGHJKLQWERTYUIOP")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return result
}

func (m *model) updateParticles() {
	for i := len(m.particles) - 1; i >= 0; i-- {
		p := &m.particles[i]
		p.X += p.VX
		p.Y += p.VY
		p.VY += 0.1 // Gravity
		p.Life--

		if p.Life <= 0 || p.X < 0 || p.X >= float64(m.width) || p.Y >= float64(m.height) {
			m.particles = append(m.particles[:i], m.particles[i+1:]...)
		}
	}
}

func (m *model) updateQuotes() {
	for i := len(m.quotes) - 1; i >= 0; i-- {
		q := &m.quotes[i]
		q.Phase += 0.05                  // Slower wobble effect (was 0.1)
		q.Wobble = math.Sin(q.Phase) * 2 // Slightly more wobble but slower
		q.X += q.VX + q.Wobble*0.1       // Reduced wobble influence
		q.Y += q.VY
		q.Life--

		// Bounce off walls with some damping
		if q.X <= 0 || q.X >= float64(m.width-len(q.Text)) {
			q.VX = -q.VX * 0.8 // Reduce velocity on bounce
		}
		if q.Y <= 0 || q.Y >= float64(m.height-1) {
			q.VY = -q.VY * 0.8 // Reduce velocity on bounce
		}

		if q.Life <= 0 {
			// Create mini explosion when quote dies
			for j := 0; j < 5; j++ {
				m.particles = append(m.particles, Particle{
					X:       q.X + float64(len(q.Text)/2),
					Y:       q.Y,
					VX:      (rand.Float64() - 0.5) * 3,
					VY:      (rand.Float64() - 0.5) * 3,
					Life:    20 + rand.Intn(20),
					MaxLife: 20 + rand.Intn(20),
					Symbol:  '✧',
					Color:   q.Color,
				})
			}
			m.quotes = append(m.quotes[:i], m.quotes[i+1:]...)
		}
	}
}

func (m *model) updateMatrixRain() {
	for i := range m.matrixRains {
		rain := &m.matrixRains[i]
		if rand.Float64() < 0.1 {
			// Change characters occasionally
			for j := range rain.Chars {
				if rand.Float64() < 0.3 {
					chars := []rune("ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇﾍ012345689")
					rain.Chars[j] = chars[rand.Intn(len(chars))]
				}
			}
		}
	}
}

func (m *model) updateExplosions() {
	for i := len(m.explosions) - 1; i >= 0; i-- {
		explosion := &m.explosions[i]
		explosion.Life--

		// Update explosion particles
		for j := len(explosion.Particles) - 1; j >= 0; j-- {
			p := &explosion.Particles[j]
			p.X += p.VX
			p.Y += p.VY
			p.VX *= 0.98 // Friction
			p.VY *= 0.98
			p.Life--

			if p.Life <= 0 {
				explosion.Particles = append(explosion.Particles[:j], explosion.Particles[j+1:]...)
			}
		}

		if explosion.Life <= 0 {
			m.explosions = append(m.explosions[:i], m.explosions[i+1:]...)
		}
	}
}

func (m model) View() string {
	switch m.state {
	case menuState:
		return m.menuView()
	case countdownState:
		return m.countdownView()
	case finishedState:
		return m.finishedView()
	}
	return ""
}

func (m model) menuView() string {
	var s strings.Builder

	// Matrix rain background
	matrix := m.renderMatrixRain()
	s.WriteString(matrix)

	// Title
	title := titleStyle.Render("🐉 WU-TANG COUNTDOWN CHAMBER 🐉")
	s.WriteString(lipgloss.Place(m.width, 5, lipgloss.Center, lipgloss.Center, title))
	s.WriteString("\n\n")

	// Subtitle
	subtitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true).
		Render("Enter the 36 Chambers of Time Management")
	s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, subtitle))
	s.WriteString("\n\n")

	// Menu items
	for i, item := range m.menuItems {
		style := menuStyle
		if i == m.menuSelection {
			style = selectedMenuStyle
			item = "▶ " + item + " ◀"
		} else {
			item = "  " + item + "  "
		}

		rendered := style.Render(item)
		s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, rendered))
		s.WriteString("\n")
	}

	s.WriteString("\n\n")
	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render("Use ↑↓ or j/k to navigate • Enter to select • q to quit")
	s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, instructions))

	return s.String()
}

func (m model) countdownView() string {
	// Create a simple grid to track quote positions
	lines := make([]string, m.height)
	for i := range lines {
		lines[i] = strings.Repeat(" ", m.width)
	}

	// Place floating quotes at their positions
	for _, q := range m.quotes {
		x := int(q.X + q.Wobble*5)
		y := int(q.Y)

		// Keep within bounds
		if x < 0 {
			x = 0
		}
		if x > m.width-len(q.Text) {
			x = m.width - len(q.Text)
		}
		if y < 0 {
			y = 0
		}
		if y >= len(lines)-8 { // Leave room for timer and UI
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

		// Insert the styled quote into the line
		if y >= 0 && y < len(lines) {
			line := []rune(lines[y])
			styledText := style.Render(q.Text)

			// Simple insertion - just put spaces and then the quote
			if x < len(line) {
				prefix := strings.Repeat(" ", x)
				if len(prefix)+len(q.Text) < m.width {
					lines[y] = prefix + styledText
				}
			}
		}
	}

	// Calculate timer position (center of screen)
	timerY := m.height / 2
	if timerY >= len(lines) {
		timerY = len(lines) - 4
	}

	// Build timer
	minutes := m.timeRemaining / 60
	seconds := m.timeRemaining % 60
	timerText := fmt.Sprintf("⏰ %02d:%02d ⏰", minutes, seconds)

	if m.glitchEffect && m.frame%10 < 3 {
		timerText = glitchStyle.Render(timerText)
	} else {
		timerText = timerStyle.Render(timerText)
	}

	// Center the timer line
	centeredTimer := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, timerText)
	if timerY >= 0 && timerY < len(lines) {
		lines[timerY] = centeredTimer
	}

	// Add particles line if fireworks mode
	if m.fireworksMode && len(m.particles) > 0 && timerY+2 < len(lines) {
		particleStr := ""
		for i, p := range m.particles {
			if i > 20 {
				break
			}
			styled := lipgloss.NewStyle().Foreground(p.Color).Bold(true).Render(string(p.Symbol))
			particleStr += styled + " "
		}
		centeredParticles := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, particleStr)
		lines[timerY+2] = centeredParticles
	}

	// Add progress bar near bottom - CENTERED
	if len(lines) >= 3 {
		progressPercent := float64(m.totalTime-m.timeRemaining) / float64(m.totalTime)
		progressBar := m.progress.ViewAs(progressPercent)
		centeredProgress := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, progressBar)
		lines[len(lines)-3] = centeredProgress
	}

	// Add status line at bottom - CENTERED
	if len(lines) >= 1 {
		statusLine := ""
		if m.glitchEffect {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true).Render("[GLITCH ON] ")
		}
		if m.fireworksMode {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render("[FIREWORKS ON] ")
		}
		statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render("g: glitch • f: fireworks • esc: menu • q: quit")

		centeredStatus := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, statusLine)
		lines[len(lines)-1] = centeredStatus
	}

	// Join all lines
	return strings.Join(lines, "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m model) finishedView() string {
	var s strings.Builder

	// Massive explosion overlay
	explosionText := m.renderMassiveExplosion()
	s.WriteString(explosionText)

	// Flashing completion message
	var message string
	if m.frame%20 < 10 {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#FFFF00")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("🔥🔥🔥 TIME'S UP! 🔥🔥🔥\nWU-TANG FOREVER!\n36 CHAMBERS COMPLETE")
	} else {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Background(lipgloss.Color("#FF0000")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("🔥🔥🔥 TIME'S UP! 🔥🔥🔥\nWU-TANG FOREVER!\n36 CHAMBERS COMPLETE")
	}

	s.WriteString(lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, message))

	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Render("Press 'r' to return to menu • 'q' to quit")
	s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Bottom, instructions))

	return s.String()
}

func (m model) renderMatrixRain() string {
	// Simple matrix background
	return ""
}

func (m model) renderParticles() string {
	// Particles rendering would be complex in this context
	// In a real implementation, you'd render to a 2D buffer
	return ""
}

func (m model) renderQuotes() string {
	// Quote rendering would overlay on the buffer
	return ""
}

func (m model) renderExplosions() string {
	// Explosion rendering
	return ""
}

func (m model) renderMassiveExplosion() string {
	var s strings.Builder

	// Create simpler explosion pattern
	for i := 0; i < 10; i++ {
		line := ""
		for j := 0; j < m.width/10; j++ {
			if rand.Float64() < 0.4 {
				symbols := []string{"★", "✦", "◆", "●", "▲", "♦", "*", "+", "x"}
				symbol := symbols[rand.Intn(len(symbols))]
				colors := []string{"#FF0000", "#FF4500", "#FFD700", "#FF1493", "#00FFFF"}
				color := colors[rand.Intn(len(colors))]
				line += lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Bold(true).Render(symbol + " ")
			} else {
				line += "  "
			}
		}
		s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, line))
		s.WriteString("\n")
	}

	return s.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
