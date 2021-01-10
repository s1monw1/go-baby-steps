package fancystrings

import (
	"bytes"
)

var emojiMap = map[string]string{
	"a": "😎",
	"b": "🙃",
	"c": "😍",
	"d": "😆",
	"e": "🤣",
	"f": "☺️",
	"g": "😚",
	"h": "🤩",
	"i": "🤗",
	"j": "😨",
	"k": "😤",
	"l": "🤓",
	"m": "🤨",
	"n": "😓",
	"o": "😰",
	"p": "🤭",
	"q": "🥰",
	"r": "😇",
	"s": "😁",
	"t": "🤥",
	"u": "🤧",
	"v": "😵",
	"w": "🤖",
	"x": "💀",
	"y": "😷",
	"z": "😵",
}

func Emojify(original string) string {
	runes := []rune(original)
	var b bytes.Buffer
	for _, char := range runes {
		match, ok := emojiMap[string(char)]
		if ok {
			b.WriteString(match)
		} else {
			b.WriteString(string(char))
		}
	}
	return b.String()
}
