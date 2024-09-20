package main

import (
	"fmt"
	"os"
	"strings"

	ascii "ascii/artistTools"
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

// main is the entry point of the program.
// it validates the input argument, prints any validation errors,
// and then calls the Artist function to process the input.
func main() {
	args := len(os.Args[1:])
	template := "standard"
	color := ColorMap["white"]
	Result := ""
	fileName := ""
	substring := ""
	switch {
	case args == 0, args > 4:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return // exit the program
	case args == 1:
		if !ascii.IsValidASCII(os.Args[1]) {
			fmt.Println("Invalid ASCII characters in input")
		}
		input := os.Args[1]
		Result = ascii.Artist(input, template, color, substring)
		ascii.OutputFinal(Result, fileName)
	case args == 2:
		if ascii.IsOutputFlag(os.Args[1]) {
			if len(os.Args[1]) < 14 || !strings.HasSuffix(os.Args[1], ".txt") {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
			}
			input := os.Args[2]
			fileName = os.Args[1][9:]
			if !ascii.IsValidOutputFileName(fileName) {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			Result = ascii.Artist(input, template, color, substring)
			ascii.OutputFinal(Result, fileName)
		} else if ascii.IsColorFlag(os.Args[1]) {
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			input := os.Args[2]
			color := ColorMap[ascii.IsColor(os.Args[1][8:])]
			Result = ascii.Artist(input, template, color, substring)
			Result = ascii.ApplyColor(Result, color)
			ascii.OutputFinal(Result, fileName)

		} else {
			if !ascii.IsValidASCII(os.Args[1]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[2]) {
				fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			template = Template(os.Args[2])
			input := os.Args[1]
			Result = ascii.Artist(input, template, color, substring)
			ascii.OutputFinal(Result, fileName)
		}
	case args == 3:
		if ascii.IsOutputFlag(os.Args[1]) {
			if len(os.Args[1]) < 14 || !strings.HasSuffix(os.Args[1], ".txt") {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			if !ascii.IsValidASCII(os.Args[2]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[3]) {
				fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			fileName = os.Args[1][9:]
			if !ascii.IsValidOutputFileName(fileName) {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
			input := os.Args[2]
			template = Template(os.Args[3])
			Result = ascii.Artist(input, template, color, substring)
			ascii.OutputFinal(Result, fileName)
		} else if ascii.IsColorFlag(os.Args[1]) {
			if ascii.IsValidSubString(os.Args[3], os.Args[2]) && !ascii.IsValidBanner(os.Args[3]) {
				if !ascii.IsValidASCII(os.Args[3]) {
					fmt.Println("Invalid ASCII characters in input")
					return
				}
				input := os.Args[3]
				substring = os.Args[2]
				color := ColorMap[ascii.IsColor(os.Args[1][8:])]
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			} else if ascii.IsValidBanner(os.Args[3]) {
				if !ascii.IsValidASCII(os.Args[2]) {
					fmt.Println("Invalid ASCII characters in input")
					return
				}
				input := os.Args[2]
				template = Template(os.Args[3])
				Result = ascii.Artist(input, template, color, substring)
				color := ColorMap[ascii.IsColor(os.Args[1][8:])]
				Result = ascii.ApplyColor(Result, color)
				ascii.OutputFinal(Result, fileName)
			} else {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
				return
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	case args == 4:
		if ascii.IsColorFlag(os.Args[1]) {
			color := ColorMap[ascii.IsColor(os.Args[1][8:])]
			if !ascii.IsValidASCII(os.Args[3]) {
				fmt.Println("Invalid ASCII characters in input")
				return
			}
			if !ascii.IsValidBanner(os.Args[4]) {
				fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\" standard")
				return
			}
			if ascii.IsValidSubString(os.Args[3], os.Args[2]) {
				input := os.Args[3]
				template = Template(os.Args[4])
				substring = os.Args[2]
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			} else {
				input := os.Args[3]
				template = Template(os.Args[4])
				Result = ascii.Artist(input, template, color, substring)
				ascii.OutputFinal(Result, fileName)
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	}

}

func Template(template string) string {
	banner := ""
	if strings.HasSuffix(template, ".txt") {
		banner = template[0 : len(os.Args[2])-4]
	} else {
		banner = template
	}
	return banner
}
