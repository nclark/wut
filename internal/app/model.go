package app

import (
	"math/rand"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/nclark/wut/internal/effects"
	"github.com/nclark/wut/internal/wutang"
)

// States
type State int

const (
	MenuState State = iota
	CountdownState
	FinishedState
)

type Model struct {
	State             State
	Width             int
	Height            int
	TimeRemaining     int
	TotalTime         int
	Progress          progress.Model
	Particles         []effects.Particle
	Quotes            []effects.FloatingQuote
	MatrixRains       []effects.MatrixRain
	Explosions        []effects.Explosion
	Frame             int
	MenuSelection     int
	GlitchEffect      bool
	FireworksMode     bool
	MenuItems         []string
	ScreenShake       int
	ScreenShakeX      int
	ScreenShakeY      int
	BeatPulse         float64
	CurrentMember     string
	MemberMode        bool
	SpinningText      bool
	StrobeEffect      bool
	WuTangLogos       []effects.FloatingQuote // Reuse FloatingQuote for logos
	FlameAnimations   []effects.Particle
	EmojiRain         bool
	CustomTimeInput   string
	InputMode         bool
	AutoWuLogos       bool
}

func InitialModel() Model {
	p := progress.New(progress.WithDefaultGradient())
	return Model{
		State:           MenuState,
		TotalTime:       900, // 15 minutes
		TimeRemaining:   900,
		Progress:        p,
		Particles:       make([]effects.Particle, 0),
		Quotes:          make([]effects.FloatingQuote, 0),
		MatrixRains:     make([]effects.MatrixRain, 0),
		Explosions:      make([]effects.Explosion, 0),
		WuTangLogos:     make([]effects.FloatingQuote, 0),
		FlameAnimations: make([]effects.Particle, 0),
		MenuSelection:   0,
		FireworksMode:   true,  // Start with fireworks blazing
		EmojiRain:       true,  // Start with emoji rain falling
		AutoWuLogos:     true,  // Start with auto logos enabled
		BeatPulse:       0,
		MenuItems: []string{
			"🔥 15 MINUTE WU-TANG COUNTDOWN 🔥",
			"⚡ 5 MINUTE SHAOLIN SPECIAL ⚡",
			"💀 1 MINUTE DEATH CHAMBER 💀",
			"🛡️ 30 SECOND FOR THE CHILDREN 🛡️",
			"⚔️ 15 SECOND PROTECT YA NECK ⚔️",
			"🎯 CUSTOM TIME (ENTER MINUTES)",
			"🐉 MEMBER MODE: " + wutang.Members[0] + " 🐉",
			"🌈 TOGGLE EFFECTS MENU 🌈",
			"❌ EXIT THE CHAMBER",
		},
		CurrentMember: wutang.Members[0],
	}
}

// TriggerScreenShake initiates screen shake effects
func (m *Model) TriggerScreenShake(intensity int) {
	m.ScreenShake = intensity
	if intensity > 0 {
		m.ScreenShakeX = rand.Intn(intensity*2) - intensity
		m.ScreenShakeY = rand.Intn(intensity*2) - intensity
	} else {
		m.ScreenShakeX = 0
		m.ScreenShakeY = 0
	}
}