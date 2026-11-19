package exporter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var cliInput io.Reader = os.Stdin
var cliOutput io.Writer = os.Stdout

func askFor(prompt string) string {
	fmt.Fprint(cliOutput, prompt+" ")
	// ReadString returns an empty string only on error (e.g., EOF), so there is no need to retry
	s, _ := bufio.NewReader(cliInput).ReadString('\n')
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
