package internal

import (
	"strings"
)

// TrimLeadingWhitespace removes leading whitespace
func TrimLeadingWhitespace(commandStr string) (newCommand string) {
	lines := strings.Split(strings.ReplaceAll(commandStr, "\t", "    "), "\n")
	leadingWhitespace := 1<<31 - 1
	for _, line := range lines {
		for pos, char := range line {
			if char == ' ' || char == '\t' {
				continue
			}
			// first non-whitespace character
			if pos < leadingWhitespace {
				leadingWhitespace = pos
			}
			// is not needed further
			break
		}
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || strings.Trim(lines[i], " \t") == "" {
			continue
		}
		if len(lines[i]) < leadingWhitespace {
			newCommand += lines[i] + "\n" // or not..
		} else {
			newCommand += lines[i][leadingWhitespace:] + "\n"
		}
	}
	return
}
