package wutang

import "github.com/charmbracelet/lipgloss"

// Wu-Tang ASCII Art
var Logo = []string{
	"    ╔═══╗   ╔═══╗    ",
	"   ╔╝   ╚╗ ╔╝   ╚╗   ",
	"  ╔╝  ╔╗ ╚═╝ ╔╗  ╚╗  ",
	" ╔╝  ╔╝╚╗   ╔╝╚╗  ╚╗ ",
	"╔╝  ╔╝  ╚═══╝  ╚╗  ╚╗",
	"╚══╝     WU      ╚══╝",
}

// Wu-Tang members for special modes
var Members = []string{
	"RZA", "GZA", "Method Man", "Raekwon", "Ghostface Killah",
	"Inspectah Deck", "U-God", "Masta Killa", "Ol' Dirty Bastard",
}

// Member-specific colors
var MemberColors = map[string]lipgloss.Color{
	"RZA":               lipgloss.Color("#FFD700"), // Gold
	"GZA":               lipgloss.Color("#00FFFF"), // Cyan
	"Method Man":        lipgloss.Color("#FF0000"), // Red
	"Raekwon":           lipgloss.Color("#800080"), // Purple
	"Ghostface Killah":  lipgloss.Color("#FFA500"), // Orange
	"Inspectah Deck":    lipgloss.Color("#00FF00"), // Green
	"U-God":             lipgloss.Color("#0000FF"), // Blue
	"Masta Killa":       lipgloss.Color("#FFFF00"), // Yellow
	"Ol' Dirty Bastard": lipgloss.Color("#FF1493"), // Deep Pink
}

// Wu-Tang quotes
var Quotes = []string{
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
	"Suuuuuu",
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
	// More classic Wu quotes and ODB madness
	"I smoke on the mic like smoking Joe Frazier",
	"First things first man you're crazy",
	"I got your money",
	"Shimmy shimmy ya",
	"Ooh baby I like it raw",
	"There ain't no father to his style",
	"Big Baby Jesus",
	"Dirt Dog",
	"I don't have no trouble with you",
	"Wu-Tang is here forever",
	"Protect your neck kid",
	"Swarm like killer bees",
	"36 styles of danger",
	"The RZA the GZA",
	"Straight from the slums of Shaolin",
	"Enter the 36 chambers",
	"Check ya self",
	"Keep it real",
	"Raw talent",
	"Microphone check",
	"Bring that beat back",
	"Wu wear the gear",
	"Staten Island in the house",
	"Peace to the Gods",
	"Mathematics",
	"Knowledge wisdom understanding",
	"Yo yo yo",
	"Check it out",
	"Word is bond",
	"No doubt",
	"For real though",
}

// Emojis for maximum ridiculousness
var CrazyEmojis = []string{
	"🔥", "💯", "🐉", "⚡", "🎯", "💀", "👹", "🤯", "🎆", "✨",
	"🌟", "💫", "🔮", "💎", "🏆", "🎸", "🎤", "🎧", "📢", "🔊",
	"🚀", "💣", "🌈", "🦾", "👑", "🗡️", "⚔️", "🛡️", "🏴‍☠️", "🎭",
	"🌪️", "🌊", "🌋", "⚡", "🔥", "💥", "✨", "🎇", "🎆", "🎉",
}

// Flame ASCII characters
var FlameChars = []string{
	"🔥", "火", "炎", "燃", "焔", "灬", "㊋", "◢◤", "▲", "△",
}