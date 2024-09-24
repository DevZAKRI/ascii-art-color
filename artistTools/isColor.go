package ascii

import (
	"strings"
)

func IsColorFlag(flag string) bool {
	return strings.HasPrefix(flag, "--color=")
}

func IsColor(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "black", "red", "green", "yellow", "blue", "magenta", "cyan", "white", "orange":
		return color
	default:
		return "white"
	}
}

func IsValidSubString(str, subString string) bool {
	return strings.Contains(str, subString)
}
