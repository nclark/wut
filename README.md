*note: i did not write any of this*

# 🐉 WUT - Wu-Tang Ultimate Timer 🐉

<div align="center">
  <img src="https://raw.githubusercontent.com/nclark/wut/refs/heads/master/demo.gif" alt="WUT - Wu-Tang Ultimate Timer Demo" width="800">
</div>

## WHAT YALL THOUGH YA WASN'T GONNA SEE ME

*Enter the 36 Chambers of Time Management*

```
docker run --rm -it ghcr.io/nclark/wut
```

## 🔥 Features 🔥

### Wu-Tang Everything
- **9 Member Modes**: Each Wu-Tang member has their own color theme
- **70+ Classic Wu-Tang Quotes**: From Triumph, C.R.E.A.M, ODB classics, and more
- **ASCII Wu-Tang Logo**: Bouncing around like the DVD screensaver
- **Member-Specific Countdown**: Shows which member is blessing your timer

### Visual Chaos
- **Rainbow Mode**: Everything cycles through rainbow colors (ALWAYS ON!)
- **Emoji Rain**: 40+ different emojis falling from the sky (ON by default!)
- **Fireworks Mode**: Continuous particle explosions (ON by default!)
- **Screen Shake**: Automatic shakes on explosions and member switches
- **Strobe Effects**: For maximum chaos (press 't' to toggle)
- **Spinning Text**: Words randomly reverse for that glitch aesthetic
- **Particle Explosions**: Press SPACE to trigger manual explosions
- **Beat Pulse System**: Particles pulse to an invisible beat
- **Massive Final Explosion**: 300+ particles when time's up
- **Quote Storm**: 20+ Wu-Tang quotes immediately flood the screen
- **High-Speed Movement**: All text hurls around at maximum velocity

### Timer Options
- 🔥 15 MINUTE WU-TANG COUNTDOWN
- ⚡ 5 MINUTE SHAOLIN SPECIAL
- 💀 1 MINUTE DEATH CHAMBER
- 🛡️ 30 SECOND FOR THE CHILDREN
- ⚔️ 15 SECOND PROTECT YA NECK
- 🎯 CUSTOM TIME (Enter any number of minutes)

## 🎮 Controls

### Menu Controls
- `↑/↓` or `j/k`: Navigate menu
- `Enter`: Select option
- `q`: Quit

### Countdown Controls
- `g`: Toggle GLITCH effect
- `f`: Toggle FIREWORKS mode
- `r`: Toggle RAINBOW mode
- `e`: Toggle EMOJI RAIN
- `w`: Spawn Wu-Tang logo
- `s`: Toggle SPINNING text
- `t`: Toggle STROBE effect
- `m`: Switch Wu-Tang MEMBER
- `SPACE`: Trigger EXPLOSION
- `ESC`: Return to menu
- `q`: Quit

## 💥 Bring da Ruckus

### **Docker One-Liner (No Install Needed):**
```bash
# Run instantly without installing anything
docker run --rm -it ghcr.io/nclark/wut
```

### **Easy Install (Go Users):**
```bash
# Install directly from GitHub
go install github.com/nclark/wut@latest
```

### **Manual Build:**
```bash
# Clone and build yourself
git clone https://github.com/nclark/wut.git
cd wut
make build
make install
```

### **Download Binary:**
Go to [Releases](https://github.com/nclark/wut/releases) and download the binary for your platform.

## 🎯 Quick Start

```bash
# Just run it!
wut
```

## 🌈 Enhanced Features Breakdown

### Member Modes
Each Wu-Tang member brings their own vibe:
- **RZA**: Gold theme - The Abbott blessing your time
- **GZA**: Cyan theme - Liquid swords of productivity
- **Method Man**: Red theme - Bring da ruckus to deadlines
- **Raekwon**: Purple theme - Only built 4 productive linx
- **Ghostface Killah**: Orange theme - Supreme time management
- **Inspectah Deck**: Green theme - Above average timer
- **U-God**: Blue theme - Raw timer power
- **Masta Killa**: Yellow theme - No said date for your tasks
- **Ol' Dirty Bastard**: Hot Pink theme - Shimmy shimmy ya productivity

### Effect Combinations
Mix and match for maximum chaos:
- Rainbow + Strobe + Emoji Rain = ULTIMATE CHAOS MODE
- Member Mode + Glitch = PERSONALIZED GLITCH
- Fireworks + Space Bar = EXPLOSION OVERLOAD
- All Effects On = YOUR TERMINAL MIGHT EXPLODE

## 🏗️ Building & Development

```bash
# Standard build
make build

# Build and run immediately
make run

# Build for all platforms
make build-all

# Create optimized release
make release

# Development with hot reload (requires air)
make dev
```

### 🧱 Project Structure

The codebase has been completely refactored from a 1200+ line monolith into a clean, modular architecture:

```
cmd/wut/main.go           # Simple 26-line entry point
internal/
├── app/                  # Core application logic
│   ├── model.go         # Data structures and state management
│   ├── update.go        # Bubble Tea Update logic and input handling
│   └── view.go          # UI rendering and visual effects
├── effects/             # Visual effects system
│   ├── types.go         # Particle, explosion, and animation types
│   └── effects.go       # All effect spawn/update functions
├── ui/                  # UI styling
│   └── styles.go        # Lipgloss styles and themes
├── utils/               # Helper utilities
│   └── utils.go         # String manipulation and math helpers
└── wutang/              # Wu-Tang specific data
    └── data.go          # 70+ quotes, emojis, ASCII art, member data
```

This modular structure makes the codebase:
- **Maintainable**: Clear separation of concerns
- **Extensible**: Easy to add new effects or features
- **Professional**: Follows Go project conventions
- **Reusable**: Components can be imported independently

## 🎪 Pro Tips

1. **Instant Chaos**: The timer now starts with maximum visual chaos enabled by default!
2. **Wu-Tang Logo Party**: Spam 'w' to fill screen with even more logos
3. **Explosion Mania**: Hold spacebar for continuous explosions on top of the default fireworks
4. **Member Hopping**: Press 'm' repeatedly for rainbow member switching
5. **Custom Time**: Use 420 minutes for the ultimate session
6. **Quote Overload**: The screen starts with 20+ quotes flying around at high speed

## ⚠️ Warnings

- May cause uncontrollable head nodding
- Screen shake may induce Wu-Tang Forever syndrome
- Strobe effect not recommended for extended use
- Your productivity might become TOO supreme

## 🙏 Credits

Wu-Tang Clan ain't nuthin' ta f' wit!

Built with:
- Go
- Bubble Tea (Terminal UI framework)
- Lipgloss (Styling)
- Pure Wu-Tang Energy

## 📜 License

This timer is for the children. Use freely.

---

*"I bomb atomically, Socrates' philosophies and hypotheses
Can't define how I be droppin' these mockeries"*

SUUUUUUUUUUUUUUUUUUUUUUUUUUUUUUU! 🐝
