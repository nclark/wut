# CLAUDE.md - Wu-Tang Ultimate Timer Project State

## 🐉 PROJECT OVERVIEW

**WUT (Wu-Tang Ultimate Timer)** is a terminal-based countdown timer that delivers maximum Wu-Tang visual chaos for productivity sessions. What started as a simple timer has evolved into a professional, modular Go application that provides an immediate visual assault of Wu-Tang culture.

## 🎯 CURRENT STATE (v2.0.1)

### 🏗️ ARCHITECTURE

**Complete Modular Refactoring Achieved:**
- **Before:** 1228-line monolithic `main.go`
- **After:** Clean 12-file modular architecture

```
cmd/wut/main.go           # 26-line entry point
internal/
├── app/                  # Core application logic (807 lines)
│   ├── model.go         # Data structures and state management
│   ├── update.go        # Bubble Tea Update logic and input handling  
│   └── view.go          # UI rendering and visual effects
├── effects/             # Visual effects system (327 lines)
│   ├── types.go         # Particle, explosion, and animation types
│   └── effects.go       # All effect spawn/update functions
├── ui/                  # UI styling (47 lines)
│   └── styles.go        # Lipgloss styles and themes
├── utils/               # Helper utilities (26 lines)
│   └── utils.go         # String manipulation and math helpers
└── wutang/              # Wu-Tang specific data (122 lines)
    └── data.go          # 70+ quotes, emojis, ASCII art, member data
```

### 🎪 WU-TANG EXPERIENCE

**Timer Options:**
- 🔥 15 MINUTE WU-TANG COUNTDOWN
- ⚡ 5 MINUTE SHAOLIN SPECIAL  
- 💀 1 MINUTE DEATH CHAMBER
- 🛡️ 30 SECOND FOR THE CHILDREN (NEW in v2.0.1)
- ⚔️ 15 SECOND PROTECT YA NECK (NEW in v2.0.1)
- 🎯 CUSTOM TIME (Enter any number of minutes)

**Default Experience (Maximum Chaos Out of Box):**
- **20+ Wu-Tang quotes** flood screen immediately on timer start
- **Fireworks mode ON** by default - continuous particle explosions
- **Emoji rain ON** by default - 40+ emojis falling constantly
- **Rainbow colors ALWAYS active** - each quote gets random vibrant color
- **4x faster movement speed** - all text hurls around at maximum velocity
- **Auto Wu-Tang logos** spawning and bouncing around
- **Screen shake, beat pulse, strobe effects** available via hotkeys

**Wu-Tang Content:**
- **70+ authentic quotes** from Triumph, C.R.E.A.M, ODB classics, and more
- **9 member modes** with individual color themes
- **ASCII Wu-Tang logo** animations
- **40+ crazy emojis** for maximum visual chaos
- All content vetted to remove problematic language while keeping energy

### 🛠️ TECHNICAL ACHIEVEMENTS

**Code Quality:**
- Professional Go project structure following conventions
- Clean separation of concerns across packages
- Zero circular dependencies
- Maintainable and extensible architecture
- Reusable components that can be imported independently

**Build System:**
- **Makefile** updated for new cmd/wut structure
- **Dockerfile** builds from modular source
- **GitHub Actions** creates multi-platform releases
- **Cross-platform binaries** (macOS M1/Intel, Linux, Windows)

**Deployment:**
- **GitHub Container Registry** hosting at `ghcr.io/nclark/wut`
- **Go module** published and installable via `go install`
- **Zero-install Docker experience** with one-liner
- **Automated releases** triggered by git tags

### 🎮 USER EXPERIENCE

**Immediate Impact:**
- No setup required - maximum chaos enabled by default
- Screen floods with visual effects the moment timer starts
- Ultra-short options (15s, 30s) for micro-productivity sessions
- All timer lengths deliver the same intense Wu-Tang experience

**Interactive Controls:**
- `g` - Toggle GLITCH effect
- `f` - Toggle FIREWORKS mode (on by default)
- `e` - Toggle EMOJI RAIN (on by default) 
- `w` - Spawn Wu-Tang logo manually
- `s` - Toggle SPINNING text
- `t` - Toggle STROBE effect
- `m` - Switch Wu-Tang MEMBER
- `SPACE` - Trigger manual EXPLOSION
- Full navigation and input controls in menu

## 🚀 DEVELOPMENT WORKFLOW

**Building:**
```bash
make build          # Standard build
make run           # Build and run immediately  
make dev           # Development with hot reload
make build-all     # Cross-platform builds
```

**Project Management:**
- All changes committed with detailed Wu-Tang themed messages
- README updated to reflect current architecture and features
- Version tags trigger automated Docker builds
- Clean git history documenting the transformation

## 🔮 TECHNICAL NOTES FOR FUTURE DEVELOPMENT

**Architecture Benefits:**
- **effects/**: Self-contained visual effects system, easy to extend
- **wutang/**: All Wu-Tang content centralized, easy to add quotes/members
- **app/**: Core Bubble Tea logic separated for clarity
- **ui/**: Styling system ready for themes/customization

**Performance Optimizations Applied:**
- Particle spawn rates tuned to prevent overwhelming
- Movement speeds calibrated for maximum chaos without lag
- Quote lifecycle management prevents memory leaks
- Screen update frequency optimized for smooth animation

**Extension Points:**
- New visual effects can be added to effects/ package
- Additional Wu-Tang content easily added to wutang/data.go
- New timer modes simple to implement in app/update.go
- UI themes can be added to ui/styles.go

## 📈 PROJECT EVOLUTION

**Phase 1:** Basic Wu-Tang timer (single file)
**Phase 2:** Enhanced with visual effects and more content
**Phase 3:** Complete architectural refactoring for maintainability  
**Phase 4:** Maximum chaos defaults and ultra-short timer options
**Phase 5:** Professional deployment and documentation

## 🎊 FINAL STATE

WUT is now a production-ready, professionally structured Wu-Tang chaos delivery system that:
- Starts with immediate visual assault requiring zero setup
- Maintains authentic Wu-Tang culture and energy
- Follows Go best practices for long-term maintainability
- Supports both micro-sessions (15s) and extended focus (15min)
- Deploys seamlessly across all platforms
- Delivers maximum productivity through Wu-Tang philosophy

**Wu-Tang Clan ain't nuthin' ta f' wit - including our codebase!** 🐉

---

*This project demonstrates the successful transformation of a creative idea into a professional, maintainable software application while preserving its unique cultural character and user experience.*