package ascii

// ValidateArgument validates the command-line arguments for the ASCII art tool.
// It checks if the number of arguments is correct and if the input string contains valid ASCII characters.
// If the arguments are not valid, it returns an error message.
// If the arguments are valid, it returns an empty string.
func ValidateArgument(args []string) string {
	if len(args) < 2 {
		return "Error: You need to enter the STRING you want in a graphic representation using ASCII"
	} else if len(args) > 2 {
		return "Error: Only one STRING at a time is allowed!"
	}

	if !IsValidASCII(args[1]) {
		return "Error: Input contains invalid characters"
	}
	return ""
}
