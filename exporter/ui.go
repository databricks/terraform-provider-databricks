package exporter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Would you like to update? [Y/n]

var input = os.Stdin

func askFor(prompt string) string {
	var s string
	r := bufio.NewReader(input)
	for {
		fmt.Fprint(os.Stdout, prompt+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func askFlag(prompt string) bool {
	res := askFor(fmt.Sprintf("%s [Y/n]", prompt))
	if res == "" {
		return true
	}
	if strings.ToLower(res) == "y" {
		return true
	}
	return false
}
