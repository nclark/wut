package effects

import "github.com/charmbracelet/lipgloss"

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