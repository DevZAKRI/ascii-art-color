package ascii

import (
	"fmt"
	"strings"
)

var (
	ColorMap = map[string]string{
		"black":     "\033[30m",
		"red":       "\033[31m",
		"green":     "\033[32m",
		"yellow":    "\033[33m",
		"blue":      "\033[34m",
		"magenta":   "\033[35m",
		"cyan":      "\033[36m",
		"white":     "\033[37m",
		"ColorSTOP": "\033[0m",
	}
)

func IsColorFlag(flag string) bool {
	return strings.HasPrefix(flag, "--color=")
}

func IsColor(color string) string {
	switch color {
	case "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white":
		return color
	default:

		return "white"
	}
}

func IsValidSubString(str, subString string) bool {
	return strings.Contains(str, subString)
}

func ApplyColor(asciiArt, color string) string {
	return fmt.Sprintf("%s%s%s", color, asciiArt, ColorMap["ColorSTOP"])
}

// colorSubstring colors the substring in the ASCII art.
func ColorSubstring(asciiArt, input, substring, color string) string {
	coloredArt := asciiArt
	substringIndex := strings.Index(input, substring)
	if substringIndex != -1 {
		asciiLines := strings.Split(asciiArt, "\n")
		substringLength := len(substring)
		for i := 0; i < 8; i++ {
			start := substringIndex * 8
			end := start + substringLength*8
			asciiLines[i] = asciiLines[i][:start] + fmt.Sprintf("%s%s%s", color, asciiLines[i][start:end], ColorMap["ColorSTOP"]) + asciiLines[i][end:]
		}
		coloredArt = strings.Join(asciiLines, "\n")
	}
	return coloredArt
}
