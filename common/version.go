package common

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

var (
	version    = ""
	onceGitVer sync.Once
)

func versionFromGit() string {
	out, err := exec.Command("git", "describe",
		"--always", "--long").Output()
	ee, ok := err.(*exec.ExitError)
	if ok && !ee.Success() || err != nil {
		return "dev"
	}
	return strings.TrimSpace(strings.ReplaceAll(string(out), "v", ""))
}

// Version returns version of provider
func Version() string {
	if version == "" {
		onceGitVer.Do(func() {
			// calling git once per process to know
			// the version that integration tests are running
			version = versionFromGit()
		})
	}
	return version
}

// UserAgent returns provider's user agent with the correct version
func UserAgent() string {
	return fmt.Sprintf("databricks-tf-provider/%s", Version())
}
