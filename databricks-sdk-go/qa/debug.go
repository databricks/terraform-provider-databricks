package qa

import (
	"os"
	"path"
	"strings"
)

// detects if test is run from "debug test" feature in VSCode
func IsInDebug() bool {
	ex, _ := os.Executable()
	return strings.HasPrefix(path.Base(ex), "__debug_bin")
}
