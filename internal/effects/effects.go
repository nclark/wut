package effects

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/nclark/wut/internal/wutang"
)

// GenerateMatrixChars creates random matrix-style characters
func GenerateMatrixChars(length int) []rune {
	chars := []rune("ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇﾍ012345789ZXCVBNMASDFGHJKLQWERTYUIOP")
	result := make([]rune, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return result
}

// GetRainbowColor generates rainbow colors using sine wave calculations
func GetRainbowColor(phase float64) lipgloss.Color {
	r := int(math.Sin(phase)*127 + 128)
	g := int(math.Sin(phase+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(phase+4*math.Pi/3)*127 + 128)
	return lipgloss.Color(fmt.Sprintf("#%02x%02x%02x", r, g, b))
}

// SpawnQuote creates a floating Wu-Tang quote
func SpawnQuote(width, height, frame int) FloatingQuote {
	quote := wutang.Quotes[rand.Intn(len(wutang.Quotes))]
	return FloatingQuote{
		Text:    quote,
		X:       rand.Float64() * float64(width-len(quote)),
		Y:       rand.Float64() * float64(height),
		VX:      (rand.Float64() - 0.5) * 1.5, // Much faster movement - increased from 0.4 to 1.5
		VY:      (rand.Float64() - 0.5) * 1.5, // Much faster movement - increased from 0.4 to 1.5
		Life:    300 + rand.Intn(300),
		MaxLife: 300 + rand.Intn(300),
		Color:   GetRainbowColor(rand.Float64() * 10), // Random rainbow color for each quote
		Phase:   rand.Float64() * math.Pi * 2,
	}
}

// SpawnParticles creates animated particles
func SpawnParticles(width, height, frame int, count int) []Particle {
	var particles []Particle
	for i := 0; i < count; i++ {
		emoji := ""
		if i%10 == 0 {
			emoji = wutang.CrazyEmojis[rand.Intn(len(wutang.CrazyEmojis))]
		}

		particles = append(particles, Particle{
			X:       rand.Float64() * float64(width),
			Y:       rand.Float64() * float64(height),
			VX:      (rand.Float64() - 0.5) * 2,
			VY:      (rand.Float64() - 0.5) * 2,
			Life:    120 + rand.Intn(180),
			MaxLife: 120 + rand.Intn(180),
			Symbol:  []rune{'✦', '★', '◆', '●', '▲', '♦', '✧', '◇'}[rand.Intn(8)],
			Color:   GetRainbowColor(float64(frame) * 0.05),
			Size:    1 + rand.Intn(3),
			Emoji:   emoji,
			Spin:    rand.Float64() * math.Pi * 2,
		})
	}
	return particles
}

// SpawnExplosion creates an explosion effect
func SpawnExplosion(width, height, frame int) Explosion {
	explosion := Explosion{
		X:       rand.Float64() * float64(width),
		Y:       rand.Float64() * float64(height),
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
			Color:   GetRainbowColor(float64(frame) * 0.05),
		})
	}

	return explosion
}

// CreateMassiveExplosion generates the final massive explosion
func CreateMassiveExplosion(width, height int) []Particle {
	centerX := float64(width) / 2
	centerY := float64(height) / 2
	
	var particles []Particle

	// Spawn tons of particles
	for i := 0; i < 300; i++ {
		angle := rand.Float64() * math.Pi * 2
		speed := rand.Float64()*10 + 2
		
		emoji := ""
		if i%3 == 0 {
			emoji = wutang.CrazyEmojis[rand.Intn(len(wutang.CrazyEmojis))]
		}
		
		particles = append(particles, Particle{
			X:       centerX,
			Y:       centerY,
			VX:      math.Cos(angle) * speed,
			VY:      math.Sin(angle) * speed,
			Life:    100 + rand.Intn(200),
			MaxLife: 100 + rand.Intn(200),
			Symbol:  []rune{'★', '✦', '◆', '●', '▲', '♦'}[rand.Intn(6)],
			Color:   GetRainbowColor(float64(i) * 0.1),
			Size:    1 + rand.Intn(4),
			Emoji:   emoji,
		})
	}
	
	return particles
}

// SpawnWuTangLogo creates a floating Wu-Tang logo
func SpawnWuTangLogo(width, height, frame int) FloatingQuote {
	logoText := strings.Join(wutang.Logo, "\n")
	return FloatingQuote{
		Text:    logoText,
		X:       rand.Float64() * float64(width-22),
		Y:       rand.Float64() * float64(height-10),
		VX:      (rand.Float64() - 0.5) * 1.0, // Faster movement - increased from 0.3 to 1.0
		VY:      (rand.Float64() - 0.5) * 0.8, // Faster movement - increased from 0.2 to 0.8
		Life:    600,
		MaxLife: 600,
		Color:   GetRainbowColor(rand.Float64() * 10), // Random rainbow colors like quotes
		Phase:   rand.Float64() * math.Pi * 2,
	}
}

// SpawnEmojiRain creates falling emoji particles
func SpawnEmojiRain(width, frame int) []Particle {
	var particles []Particle
	// Reduce from 5 to 2 emojis per spawn
	for i := 0; i < 2; i++ {
		emoji := wutang.CrazyEmojis[rand.Intn(len(wutang.CrazyEmojis))]
		particles = append(particles, Particle{
			X:       rand.Float64() * float64(width),
			Y:       -5,
			VX:      (rand.Float64() - 0.5) * 1, // Slower horizontal movement
			VY:      rand.Float64()*2 + 0.5,     // Slower fall speed
			Life:    150 + rand.Intn(100),       // Longer life
			MaxLife: 150 + rand.Intn(100),
			Emoji:   emoji,
			Color:   GetRainbowColor(float64(frame) * 0.05),
			Size:    1 + rand.Intn(2), // Smaller max size
			Spin:    rand.Float64() * math.Pi * 2,
		})
	}
	return particles
}

// UpdateParticles updates particle physics
func UpdateParticles(particles []Particle, width, height int) []Particle {
	for i := len(particles) - 1; i >= 0; i-- {
		p := &particles[i]
		p.X += p.VX
		p.Y += p.VY
		p.VY += 0.05 // Lighter gravity
		p.Life--
		p.Spin += 0.05 // Slower spin

		if p.Life <= 0 || p.X < 0 || p.X >= float64(width) || p.Y >= float64(height) {
			particles = append(particles[:i], particles[i+1:]...)
		}
	}
	return particles
}

// UpdateQuotes updates floating quote positions and creates mini explosions
func UpdateQuotes(quotes []FloatingQuote, width, height int) ([]FloatingQuote, []Particle) {
	var newParticles []Particle
	
	for i := len(quotes) - 1; i >= 0; i-- {
		q := &quotes[i]
		q.Phase += 0.02 // Slower phase change
		q.Wobble = math.Sin(q.Phase) * 1 // Less wobble
		q.X += q.VX + q.Wobble*0.05 // Less wobble influence
		q.Y += q.VY
		q.Life--

		// Bounce off walls with some damping
		if q.X <= 0 || q.X >= float64(width-len(q.Text)) {
			q.VX = -q.VX * 0.8
		}
		if q.Y <= 0 || q.Y >= float64(height-1) {
			q.VY = -q.VY * 0.8
		}

		if q.Life <= 0 {
			// Create smaller mini explosion when quote dies
			for j := 0; j < 2; j++ {
				newParticles = append(newParticles, Particle{
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
			quotes = append(quotes[:i], quotes[i+1:]...)
		}
	}
	return quotes, newParticles
}

// UpdateWuTangLogos updates Wu-Tang logo floating animations
func UpdateWuTangLogos(logos []FloatingQuote, width, height int) []FloatingQuote {
	for i := len(logos) - 1; i >= 0; i-- {
		logo := &logos[i]
		logo.Phase += 0.02
		logo.X += logo.VX + math.Sin(logo.Phase)*0.5
		logo.Y += logo.VY
		logo.Life--

		// Bounce off walls
		if logo.X <= 0 || logo.X >= float64(width-22) {
			logo.VX = -logo.VX
		}
		if logo.Y <= 0 || logo.Y >= float64(height-8) {
			logo.VY = -logo.VY
		}

		if logo.Life <= 0 {
			logos = append(logos[:i], logos[i+1:]...)
		}
	}
	return logos
}

// UpdateMatrixRain updates matrix rain background effect
func UpdateMatrixRain(rains []MatrixRain, width, height int) []MatrixRain {
	for i := range rains {
		rain := &rains[i]
		// Randomly change characters occasionally
		if rand.Float64() < 0.1 {
			for j := range rain.Chars {
				if rand.Float64() < 0.3 {
					chars := []rune("ﾊﾐﾋｰｳｼﾅﾓﾆｻﾜﾂｵﾘｱﾎﾃﾏｹﾒｴｶｷﾑﾕﾗｾﾈｽﾀﾇﾍ012345689")
					rain.Chars[j] = chars[rand.Intn(len(chars))]
				}
			}
		}
	}
	return rains
}

// UpdateExplosions updates explosion animations
func UpdateExplosions(explosions []Explosion, width, height int) []Explosion {
	for i := len(explosions) - 1; i >= 0; i-- {
		explosion := &explosions[i]
		explosion.Life--
		
		// Update explosion particles
		explosion.Particles = UpdateParticles(explosion.Particles, width, height)
		
		if explosion.Life <= 0 || len(explosion.Particles) == 0 {
			explosions = append(explosions[:i], explosions[i+1:]...)
		}
	}
	return explosions
}