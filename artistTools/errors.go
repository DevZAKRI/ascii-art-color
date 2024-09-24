package ascii

import (
	"fmt"
)

func ErrInvalidChar() {
	fmt.Println("Invalid STRING in input\" Contain Invalid Characters\"")
}

func ErrArgs() {
	fmt.Println("Invalid number of arguments\n\n Usage: go run . [OPTION] [SUBSTRING] [STRING] [BANNER]\n\n [OPTION]: --output=<fileName.txt> or --color=<color>\n [SUBSTRING]: In Case Want to color a certain substring in string, Only Work With Color Flag \n [STRING]: The string to be converted to ASCII\n [BANNER]: The ASCII banner to be used\n\n EX: go run . --output=fileName.txt something standard")
}

func UsageGlobal() {
	fmt.Println("Usage: go run . [OPTION] [SUBSTRING] [STRING] [BANNER]\n\n [OPTION]: --output=<fileName.txt> or --color=<color>\n [SUBSTRING]: In Case Want to color a certain substring in string, Only Work With Color Flag \n [STRING]: The string to be converted to ASCII\n [BANNER]: The ASCII banner to be used\n\n EX: go run . --output=fileName.txt something standard")
}

func ErrFsOnly() {
	fmt.Println("Usage: go run . [STRING] [BANNER]\n [STRING]: The string to be converted to ASCII\n [BANNER]: The ASCII banner to be used\n\t -standard\n\t -shadow\n\t -thinkertoy\n\n EX: go run . something standard")
}

func ErrOutput() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func ErrColor() {
	fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
}

// func ErrReverse() {
// 	fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
// }
