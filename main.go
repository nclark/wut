package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
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

// Wu-Tang ASCII Art
var wuTangLogo = []string{
	"    ╔═══╗   ╔═══╗    ",
	"   ╔╝   ╚╗ ╔╝   ╚╗   ",
	"  ╔╝  ╔╗ ╚═╝ ╔╗  ╚╗  ",
	" ╔╝  ╔╝╚╗   ╔╝╚╗  ╚╗ ",
	"╔╝  ╔╝  ╚═══╝  ╚╗  ╚╗",
	"╚══╝     WU      ╚══╝",
}

// Wu-Tang members for special modes
var wuTangMembers = []string{
	"RZA", "GZA", "Method Man", "Raekwon", "Ghostface Killah",
	"Inspectah Deck", "U-God", "Masta Killa", "Ol' Dirty Bastard",
}

// Member-specific colors
var memberColors = map[string]lipgloss.Color{
	"RZA":              lipgloss.Color("#FFD700"), // Gold
	"GZA":              lipgloss.Color("#00FFFF"), // Cyan
	"Method Man":       lipgloss.Color("#FF0000"), // Red
	"Raekwon":          lipgloss.Color("#800080"), // Purple
	"Ghostface Killah": lipgloss.Color("#FFA500"), // Orange
	"Inspectah Deck":   lipgloss.Color("#00FF00"), // Green
	"U-God":            lipgloss.Color("#0000FF"), // Blue
	"Masta Killa":      lipgloss.Color("#FFFF00"), // Yellow
	"Ol' Dirty Bastard": lipgloss.Color("#FF1493"), // Deep Pink
}

// Wu-Tang quotes - EXPANDED
var wuTangQuotes = []string{
	// Classic quotes
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
	// New additions
	"Killa beez on the swarm",
	"Shame on a n***a",
	"Suuuuuuu",
	"Brooklyn Zoo",
	"Clan in da front",
	"Da mystery of chessboxin'",
	"Criminology",
	"Ice cream",
	"Guillotine swordz",
	"4th Chamber",
	"Shadowboxin'",
	"Reunited",
	"It's Yourz",
	"Hellz Wind Staff",
	"Severe Punishment",
	"Older Gods",
	"A Better Tomorrow",
	"Wu-Tang Forever",
	"Gravel Pit",
	"Uzi (Pinky Ring)",
}

// Emojis for maximum ridiculousness
var crazyEmojis = []string{
	"🔥", "💯", "🐉", "⚡", "🎯", "💀", "👹", "🤯", "🎆", "✨",
	"🌟", "💫", "🔮", "💎", "🏆", "🎸", "🎤", "🎧", "📢", "🔊",
	"🚀", "💣", "🌈", "🦾", "👑", "🗡️", "⚔️", "🛡️", "🏴‍☠️", "🎭",
	"🌪️", "🌊", "🌋", "⚡", "🔥", "💥", "✨", "🎇", "🎆", "🎉",
}

// Flame ASCII characters
var flameChars = []string{
	"🔥", "火", "炎", "燃", "焔", "灬", "㊋", "◢◤", "▲", "△",
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
	Emoji   string // For emoji particles
	Spin    float64
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
	state             state
	width             int
	height            int
	timeRemaining     int
	totalTime         int
	progress          progress.Model
	particles         []Particle
	quotes            []FloatingQuote
	matrixRains       []MatrixRain
	explosions        []Explosion
	frame             int
	menuSelection     int
	glitchEffect      bool
	fireworksMode     bool
	menuItems         []string
	screenShake       int
	screenShakeX      int
	screenShakeY      int
	rainbowMode       bool
	beatPulse         float64
	currentMember     string
	memberMode        bool
	spinningText      bool
	strobeEffect      bool
	wuTangLogos       []FloatingQuote // Reuse FloatingQuote for logos
	flameAnimations   []Particle
	emojiRain         bool
	customTimeInput   string
	inputMode         bool
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
		state:           menuState,
		totalTime:       900, // 15 minutes
		timeRemaining:   900,
		progress:        p,
		particles:       make([]Particle, 0),
		quotes:          make([]FloatingQuote, 0),
		matrixRains:     make([]MatrixRain, 0),
		explosions:      make([]Explosion, 0),
		wuTangLogos:     make([]FloatingQuote, 0),
		flameAnimations: make([]Particle, 0),
		menuSelection:   0,
		rainbowMode:     false, // Start calmer
		emojiRain:       false, // Start calmer
		beatPulse:       0,
		menuItems: []string{
			"🔥 15 MINUTE WU-TANG COUNTDOWN 🔥",
			"⚡ 5 MINUTE SHAOLIN SPECIAL ⚡",
			"💀 1 MINUTE DEATH CHAMBER 💀",
			"🎯 CUSTOM TIME (ENTER MINUTES)",
			"🐉 MEMBER MODE: " + wuTangMembers[0] + " 🐉",
			"🌈 TOGGLE EFFECTS MENU 🌈",
			"❌ EXIT THE CHAMBER",
		},
		currentMember: wuTangMembers[0],
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		tickCmd(),
	)
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
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
				if !m.inputMode {
					m.menuSelection = (m.menuSelection - 1 + len(m.menuItems)) % len(m.menuItems)
				}
			case "down", "j":
				if !m.inputMode {
					m.menuSelection = (m.menuSelection + 1) % len(m.menuItems)
				}
			case "enter":
				if m.inputMode {
					// Parse custom time
					if minutes, err := strconv.Atoi(m.customTimeInput); err == nil && minutes > 0 {
						m.totalTime = minutes * 60
						m.timeRemaining = m.totalTime
						m.state = countdownState
						m.fireworksMode = true
						m.inputMode = false
						return m, tea.Batch(tickCmd(), m.spawnInitialEffects())
					}
				} else {
					switch m.menuSelection {
					case 0:
						m.totalTime = 900
						m.timeRemaining = 900
						m.state = countdownState
						m.fireworksMode = true
						return m, tea.Batch(tickCmd(), m.spawnInitialEffects())
					case 1:
						m.totalTime = 300
						m.timeRemaining = 300
						m.state = countdownState
						m.fireworksMode = true
						return m, tea.Batch(tickCmd(), m.spawnInitialEffects())
					case 2:
						m.totalTime = 60
						m.timeRemaining = 60
						m.state = countdownState
						m.fireworksMode = true
						return m, tea.Batch(tickCmd(), m.spawnInitialEffects())
					case 3:
						// Custom time input
						m.inputMode = true
						m.customTimeInput = ""
					case 4:
						// Cycle through members
						currentIdx := 0
						for i, member := range wuTangMembers {
							if member == m.currentMember {
								currentIdx = i
								break
							}
						}
						m.currentMember = wuTangMembers[(currentIdx+1)%len(wuTangMembers)]
						m.menuItems[4] = "🐉 MEMBER MODE: " + m.currentMember + " 🐉"
						m.memberMode = true
					case 5:
						// Effects toggle - could open submenu
						m.rainbowMode = !m.rainbowMode
						m.emojiRain = !m.emojiRain
						m.spinningText = !m.spinningText
						m.strobeEffect = !m.strobeEffect
					case 6:
						return m, tea.Quit
					}
				}
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				if m.inputMode {
					m.customTimeInput += msg.String()
				}
			case "backspace":
				if m.inputMode && len(m.customTimeInput) > 0 {
					m.customTimeInput = m.customTimeInput[:len(m.customTimeInput)-1]
				}
			case "escape", "esc":
				if m.inputMode {
					m.inputMode = false
					m.customTimeInput = ""
				}
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
				m.wuTangLogos = m.wuTangLogos[:0]
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
			case "r":
				m.rainbowMode = !m.rainbowMode
			case "e":
				m.emojiRain = !m.emojiRain
			case "w":
				m.spawnWuTangLogo()
			case "s":
				m.spinningText = !m.spinningText
			case "t":
				m.strobeEffect = !m.strobeEffect
			case "m":
				// Quick member switch
				currentIdx := 0
				for i, member := range wuTangMembers {
					if member == m.currentMember {
						currentIdx = i
						break
					}
				}
				m.currentMember = wuTangMembers[(currentIdx+1)%len(wuTangMembers)]
				m.memberMode = true
				m.triggerScreenShake(10)
			case "space", " ":
				// Epic explosion on spacebar
				m.spawnExplosion()
				m.triggerScreenShake(20)
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
				m.wuTangLogos = m.wuTangLogos[:0]
				m.frame = 0
			case "q", "ctrl+c":
				return m, tea.Quit
			}
		}

	case tickMsg:
		m.frame++

		// Update beat pulse
		m.beatPulse = math.Sin(float64(m.frame) * 0.1)

		// Update screen shake
		if m.screenShake > 0 {
			m.screenShake--
			if m.screenShake > 0 {
				m.screenShakeX = rand.Intn(m.screenShake*2) - m.screenShake
				m.screenShakeY = rand.Intn(m.screenShake*2) - m.screenShake
			} else {
				m.screenShakeX = 0
				m.screenShakeY = 0
			}
		}

		if m.state == countdownState {
			// Update timer every second (10 frames at 100ms)
			if m.frame%10 == 0 && m.timeRemaining > 0 {
				m.timeRemaining--
				if m.timeRemaining == 0 {
					m.state = finishedState
					return m, m.createMassiveExplosion()
				}

				// Spawn effects at specific times
				if m.timeRemaining%10 == 0 {
					m.triggerScreenShake(3)
				}
			}

			// Spawn effects - much slower now
			if m.frame%50 == 0 { // Spawn quotes less frequently
				m.spawnQuote()
			}
			if m.frame%80 == 0 { // Spawn logos occasionally
				m.spawnWuTangLogo()
			}
			if m.emojiRain && m.frame%30 == 0 { // Much slower emoji rain
				m.spawnEmojiRain()
			}
			if m.frame%20 == 0 && m.fireworksMode { // Slower particles in fireworks mode
				m.spawnParticles()
			} else if m.frame%40 == 0 { // Much slower normal particle spawn
				m.spawnParticles()
			}
			if m.frame%100 == 0 { // Much slower explosions
				m.spawnExplosion()
			}
		}

		// Update all effects
		m.updateParticles()
		m.updateQuotes()
		m.updateMatrixRain()
		m.updateExplosions()
		m.updateWuTangLogos()

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
	
	// Spawn initial Wu-Tang logo
	m.spawnWuTangLogo()
	
	// Trigger initial screen shake
	m.triggerScreenShake(15)
	
	return nil
}

func (m *model) spawnQuote() {
	if len(m.quotes) > 15 {
		return
	}

	quote := wuTangQuotes[rand.Intn(len(wuTangQuotes))]
	
	// Member-specific quotes
	if m.memberMode {
		quote = "[" + m.currentMember + "] " + quote
	}

	m.quotes = append(m.quotes, FloatingQuote{
		Text:    quote,
		X:       rand.Float64() * float64(m.width-len(quote)),
		Y:       rand.Float64() * float64(m.height-5),
		VX:      (rand.Float64() - 0.5) * 0.5,
		VY:      (rand.Float64() - 0.5) * 0.3,
		Life:    400 + rand.Intn(600),
		MaxLife: 400 + rand.Intn(600),
		Color:   m.getMemberColor(),
		Phase:   rand.Float64() * math.Pi * 2,
	})
}

func (m *model) spawnParticles() {
	symbols := []rune{'★', '✦', '✧', '◆', '◇', '●', '◉', '▲', '▼', '♦'}
	
	// Reduce from 3 to 2 particles per spawn
	for i := 0; i < 2; i++ {
		// Pulse particles with beat
		size := 1 + rand.Intn(2) // Smaller max size
		if m.beatPulse > 0.5 {
			size++
		}
		
		m.particles = append(m.particles, Particle{
			X:       rand.Float64() * float64(m.width),
			Y:       rand.Float64() * float64(m.height),
			VX:      (rand.Float64() - 0.5) * 2, // Slower movement
			VY:      (rand.Float64() - 0.5) * 2, // Slower movement
			Life:    80 + rand.Intn(120),        // Longer life
			MaxLife: 80 + rand.Intn(120),
			Symbol:  symbols[rand.Intn(len(symbols))],
			Color:   m.getMemberColor(),
			Size:    size,
			Spin:    rand.Float64() * math.Pi * 2,
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
			Color:   m.getMemberColor(),
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

	// Spawn tons of particles
	for i := 0; i < 300; i++ {
		angle := rand.Float64() * math.Pi * 2
		speed := rand.Float64()*10 + 2
		
		emoji := ""
		if i%3 == 0 {
			emoji = crazyEmojis[rand.Intn(len(crazyEmojis))]
		}
		
		m.particles = append(m.particles, Particle{
			X:       centerX,
			Y:       centerY,
			VX:      math.Cos(angle) * speed,
			VY:      math.Sin(angle) * speed,
			Life:    100 + rand.Intn(200),
			MaxLife: 100 + rand.Intn(200),
			Symbol:  []rune{'★', '✦', '◆', '●', '▲', '♦'}[rand.Intn(6)],
			Color:   getRainbowColor(float64(i) * 0.1),
			Size:    1 + rand.Intn(4),
			Emoji:   emoji,
		})
	}
	
	// Massive screen shake
	m.triggerScreenShake(50)

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

// Rainbow color generator
func getRainbowColor(phase float64) lipgloss.Color {
	r := int(math.Sin(phase)*127 + 128)
	g := int(math.Sin(phase+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(phase+4*math.Pi/3)*127 + 128)
	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))
}

// Get member theme color with rainbow override
func (m model) getMemberColor() lipgloss.Color {
	if m.rainbowMode {
		return getRainbowColor(float64(m.frame) * 0.05)
	}
	if m.memberMode && m.currentMember != "" {
		if color, ok := memberColors[m.currentMember]; ok {
			return color
		}
	}
	return lipgloss.Color("#FFD700")
}

// Trigger screen shake
func (m *model) triggerScreenShake(intensity int) {
	m.screenShake = intensity
	m.screenShakeX = rand.Intn(intensity*2) - intensity
	m.screenShakeY = rand.Intn(intensity*2) - intensity
}

// Spawn Wu-Tang logo
func (m *model) spawnWuTangLogo() {
	logoText := strings.Join(wuTangLogo, "\n")
	m.wuTangLogos = append(m.wuTangLogos, FloatingQuote{
		Text:    logoText,
		X:       rand.Float64() * float64(m.width-22),
		Y:       rand.Float64() * float64(m.height-10),
		VX:      (rand.Float64() - 0.5) * 0.3,
		VY:      (rand.Float64() - 0.5) * 0.2,
		Life:    600,
		MaxLife: 600,
		Color:   m.getMemberColor(),
		Phase:   rand.Float64() * math.Pi * 2,
	})
}

// Spawn emoji rain
func (m *model) spawnEmojiRain() {
	// Reduce from 5 to 2 emojis per spawn
	for i := 0; i < 2; i++ {
		emoji := crazyEmojis[rand.Intn(len(crazyEmojis))]
		m.particles = append(m.particles, Particle{
			X:       rand.Float64() * float64(m.width),
			Y:       -5,
			VX:      (rand.Float64() - 0.5) * 1, // Slower horizontal movement
			VY:      rand.Float64()*2 + 0.5,     // Slower fall speed
			Life:    150 + rand.Intn(100),       // Longer life
			MaxLife: 150 + rand.Intn(100),
			Emoji:   emoji,
			Color:   m.getMemberColor(),
			Size:    1 + rand.Intn(2), // Smaller max size
			Spin:    rand.Float64() * math.Pi * 2,
		})
	}
}

func (m *model) updateParticles() {
	for i := len(m.particles) - 1; i >= 0; i-- {
		p := &m.particles[i]
		p.X += p.VX
		p.Y += p.VY
		p.VY += 0.05 // Lighter gravity
		p.Life--
		p.Spin += 0.05 // Slower spin

		if p.Life <= 0 || p.X < 0 || p.X >= float64(m.width) || p.Y >= float64(m.height) {
			m.particles = append(m.particles[:i], m.particles[i+1:]...)
		}
	}
}

func (m *model) updateQuotes() {
	for i := len(m.quotes) - 1; i >= 0; i-- {
		q := &m.quotes[i]
		q.Phase += 0.02 // Slower phase change
		q.Wobble = math.Sin(q.Phase) * 1 // Less wobble
		q.X += q.VX + q.Wobble*0.05 // Less wobble influence
		q.Y += q.VY
		q.Life--

		// Bounce off walls with some damping
		if q.X <= 0 || q.X >= float64(m.width-len(q.Text)) {
			q.VX = -q.VX * 0.8
		}
		if q.Y <= 0 || q.Y >= float64(m.height-1) {
			q.VY = -q.VY * 0.8
		}

		if q.Life <= 0 {
			// Create smaller mini explosion when quote dies
			for j := 0; j < 2; j++ {
				m.particles = append(m.particles, Particle{
					X:       q.X + float64(len(q.Text)/2),
					Y:       q.Y,
					VX:      (rand.Float64() - 0.5) * 1.5, // Slower
					VY:      (rand.Float64() - 0.5) * 1.5, // Slower
					Life:    30 + rand.Intn(20),           // Longer life
					MaxLife: 30 + rand.Intn(20),
					Symbol:  '✧',
					Color:   q.Color,
				})
			}
			m.quotes = append(m.quotes[:i], m.quotes[i+1:]...)
		}
	}
}

func (m *model) updateWuTangLogos() {
	for i := len(m.wuTangLogos) - 1; i >= 0; i-- {
		logo := &m.wuTangLogos[i]
		logo.Phase += 0.02
		logo.X += logo.VX + math.Sin(logo.Phase)*0.5
		logo.Y += logo.VY
		logo.Life--

		// Bounce off walls
		if logo.X <= 0 || logo.X >= float64(m.width-22) {
			logo.VX = -logo.VX
		}
		if logo.Y <= 0 || logo.Y >= float64(m.height-8) {
			logo.VY = -logo.VY
		}

		if logo.Life <= 0 {
			m.wuTangLogos = append(m.wuTangLogos[:i], m.wuTangLogos[i+1:]...)
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

	// Apply screen shake
	if m.screenShake > 0 {
		s.WriteString(strings.Repeat("\n", m.screenShakeY))
		s.WriteString(strings.Repeat(" ", m.screenShakeX))
	}

	// Title with spinning effect
	title := "🐉 WU-TANG COUNTDOWN CHAMBER 🐉"
	if m.spinningText && m.frame%20 < 10 {
		title = "🐲 CHAMBER COUNTDOWN TANG-WU 🐲"
	}
	titleRendered := titleStyle.Render(title)
	s.WriteString(lipgloss.Place(m.width, 5, lipgloss.Center, lipgloss.Center, titleRendered))
	s.WriteString("\n\n")

	// Subtitle
	subtitle := lipgloss.NewStyle().
		Foreground(m.getMemberColor()).
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

		// Show custom time input
		if i == 3 && m.inputMode {
			item = "  🎯 ENTER MINUTES: " + m.customTimeInput + "_  "
		}

		rendered := style.Render(item)
		s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, rendered))
		s.WriteString("\n")
	}

	// Effects status
	s.WriteString("\n")
	effectsStatus := fmt.Sprintf("🌈 Rainbow: %v | 🎭 Emoji Rain: %v | 💫 Spinning: %v | ⚡ Strobe: %v",
		m.rainbowMode, m.emojiRain, m.spinningText, m.strobeEffect)
	effectsStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#666666"))
	s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, effectsStyle.Render(effectsStatus)))

	s.WriteString("\n\n")
	instructions := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Render("Use ↑↓ or j/k to navigate • Enter to select • q to quit")
	s.WriteString(lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, instructions))

	return s.String()
}

func (m model) countdownView() string {
	// Apply strobe effect
	if m.strobeEffect && m.frame%10 < 2 {
		return strings.Repeat(" ", m.width*m.height)
	}

	// Create a simple grid to track quote positions
	lines := make([]string, m.height)
	for i := range lines {
		lines[i] = strings.Repeat(" ", m.width)
	}

	// Apply screen shake offset
	offsetX := m.screenShakeX
	offsetY := m.screenShakeY

	// Place Wu-Tang logos
	for _, logo := range m.wuTangLogos {
		x := int(logo.X) + offsetX
		y := int(logo.Y) + offsetY
		
		logoLines := strings.Split(logo.Text, "\n")
		for lineIdx, logoLine := range logoLines {
			lineY := y + lineIdx
			if lineY >= 0 && lineY < len(lines) && x >= 0 && x < m.width-22 {
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
	for _, q := range m.quotes {
		x := int(q.X+q.Wobble*5) + offsetX
		y := int(q.Y) + offsetY

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
		if m.spinningText && rand.Float64() < 0.1 {
			// Reverse random words
			words := strings.Split(text, " ")
			if len(words) > 1 && rand.Float64() < 0.3 {
				idx := rand.Intn(len(words))
				words[idx] = reverseString(words[idx])
			}
			text = strings.Join(words, " ")
		}

		// Insert the styled quote into the line
		if y >= 0 && y < len(lines) {
			styledText := style.Render(text)
			if x < m.width {
				prefix := strings.Repeat(" ", x)
				if len(prefix)+len(text) < m.width {
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

	// Build timer with member theme
	minutes := m.timeRemaining / 60
	seconds := m.timeRemaining % 60
	timerText := fmt.Sprintf("⏰ %02d:%02d ⏰", minutes, seconds)
	
	if m.memberMode {
		timerText = fmt.Sprintf("🎤 %s | %02d:%02d 🎤", m.currentMember, minutes, seconds)
	}

	// Apply effects to timer
	if m.glitchEffect && m.frame%10 < 3 {
		timerText = glitchStyle.Render(timerText)
	} else {
		timerStyle := timerStyle.Foreground(m.getMemberColor())
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
			
			if p.Emoji != "" {
				// Emoji particles
				particleStr += p.Emoji + " "
			} else {
				// Regular particles with spin effect
				symbol := p.Symbol
				if m.spinningText && rand.Float64() < 0.3 {
					symbols := []rune{'/', '-', '\\', '|'}
					symbol = symbols[int(p.Spin)%4]
				}
				styled := lipgloss.NewStyle().Foreground(p.Color).Bold(true).Render(string(symbol))
				particleStr += styled + " "
			}
		}
		centeredParticles := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, particleStr)
		lines[timerY+2] = centeredParticles
	}

	// Add progress bar near bottom - CENTERED with pulse effect
	if len(lines) >= 3 {
		progressPercent := float64(m.totalTime-m.timeRemaining) / float64(m.totalTime)
		
		// Pulse the progress bar width
		pulseWidth := m.progress.Width
		if m.beatPulse > 0 {
			pulseWidth = int(float64(pulseWidth) * (1 + m.beatPulse*0.1))
		}
		
		m.progress.Width = pulseWidth
		progressBar := m.progress.ViewAs(progressPercent)
		centeredProgress := lipgloss.Place(m.width, 1, lipgloss.Center, lipgloss.Center, progressBar)
		lines[len(lines)-3] = centeredProgress
	}

	// Add status line at bottom - CENTERED
	if len(lines) >= 1 {
		statusLine := ""
		if m.glitchEffect {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true).Render("[GLITCH] ")
		}
		if m.fireworksMode {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Bold(true).Render("[FIREWORKS] ")
		}
		if m.rainbowMode {
			statusLine += lipgloss.NewStyle().Foreground(getRainbowColor(float64(m.frame)*0.1)).Bold(true).Render("[RAINBOW] ")
		}
		if m.emojiRain {
			statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#FF69B4")).Bold(true).Render("[EMOJI RAIN] ")
		}
		
		statusLine += "\n"
		statusLine += lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).
			Render("g:glitch f:fireworks r:rainbow e:emoji w:wu-logo s:spin t:strobe m:member space:explode")

		centeredStatus := lipgloss.Place(m.width, 2, lipgloss.Center, lipgloss.Center, statusLine)
		lines[len(lines)-2] = centeredStatus[:len(centeredStatus)-1] // Remove trailing newline
		lines[len(lines)-1] = centeredStatus[len(centeredStatus)-1:]
	}

	// Join all lines
	return strings.Join(lines, "\n")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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

	// Apply screen shake to finished message
	offsetX := m.screenShakeX
	offsetY := m.screenShakeY
	
	if offsetY > 0 {
		s.WriteString(strings.Repeat("\n", offsetY))
	}
	if offsetX > 0 {
		s.WriteString(strings.Repeat(" ", offsetX))
	}

	// Flashing completion message
	var message string
	if m.frame%20 < 10 {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Background(lipgloss.Color("#FFFF00")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("🔥🔥🔥 TIME'S UP! 🔥🔥🔥\nWU-TANG FOREVER!\n36 CHAMBERS COMPLETE\n" + m.currentMember + " APPROVES!")
	} else {
		message = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Background(lipgloss.Color("#FF0000")).
			Bold(true).
			Padding(2, 4).
			Border(lipgloss.ThickBorder()).
			Render("💀💀💀 SHAOLIN STYLE! 💀💀💀\nBRING DA RUCKUS!\nKILLA BEEZ ON ATTACK\n" + m.currentMember + " IN THE HOUSE!")
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

func (m model) renderMassiveExplosion() string {
	var s strings.Builder

	// Create epic explosion pattern with emojis and colors
	for i := 0; i < 15; i++ {
		line := ""
		for j := 0; j < m.width/5; j++ {
			if rand.Float64() < 0.6 {
				// Mix of emojis and symbols
				if rand.Float64() < 0.5 {
					emoji := crazyEmojis[rand.Intn(len(crazyEmojis))]
					line += emoji + " "
				} else {
					symbols := []string{"★", "✦", "◆", "●", "▲", "♦", "*", "+", "x"}
					symbol := symbols[rand.Intn(len(symbols))]
					color := getRainbowColor(float64(i*j) * 0.1)
					line += lipgloss.NewStyle().Foreground(color).Bold(true).Render(symbol + " ")
				}
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