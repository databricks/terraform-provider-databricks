package exporter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	cliInput  io.Reader = os.Stdin
	cliOutput io.Writer = os.Stdout
)

func askFor(prompt string) string {
	var s string
	r := bufio.NewReader(cliInput)
	for {
		fmt.Fprint(cliOutput, prompt+" ")
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
