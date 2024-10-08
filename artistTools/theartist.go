package ascii

import (
	"fmt"
	"os"
	"strings"
)

// Artist is a function that takes an input string and read the standard file then pass both to PrintLineAsAscii.
// It reads the file containing the ASCII graph representation and puts the content in asciiGraph.
// If there is an error reading the file, it prints the error message and returns.
// It splits the input string into lines using the "\\n" delimiter.
func Artist(input, template, color, substring string) string {
	// check for errors too in case file no longer exist
	asciiGraph, err := ReadFile("banners/" + template + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	for idx, LINE := range asciiGraph {
		asciiGraph[idx] = strings.ReplaceAll(LINE, "\r", "")
	}
	lines := strings.Split(input, "\\n")
	// If the input string is empty after removing "\\n", it removes the first line.
	if strings.ReplaceAll(input, "\\n", "") == "" {
		lines = lines[1:]
	}
	finalAsciiArt := ""
	for _, line := range lines {
		// add each line to finalAsciiArt using the LineAsAscii function.
		finalAsciiArt += LineAsAscii(line, color, substring, asciiGraph)
	}
	return finalAsciiArt
}

func LineAsAscii(line, color, substring string, asciiGraph []string) string {
	var asciiChars []string
	finalAsciiArt := ""
	//substring := "X"
	substringIndex := 0
	if line != "" {
		if strings.Contains(line, substring) {
			substringIndex = strings.Index(line, substring)
		}
		for idx, char := range line {
			if idx >= substringIndex && idx < substringIndex+len(substring) {
				for i := 8; i > 0; i-- {
					asciiChars = append(asciiChars, color+string(asciiGraph[findLastLine(char)-i])+ColorMap["ColorSTOP"])
				}
			} else {
				for i := 8; i > 0; i-- {
					asciiChars = append(asciiChars, string(asciiGraph[findLastLine(char)-i]))
				}
			}
		}

		for i := 0; i < 8; i++ {
			for j := 0; j < len(asciiChars); j += 8 {
				finalAsciiArt += asciiChars[i+j]
			}
			finalAsciiArt += "\n"
		}
	} else {
		finalAsciiArt += "\n"
	}
	return finalAsciiArt
}

// find the last line after the char
func findLastLine(char rune) int {
	return int((char - 31) * (9))
}
