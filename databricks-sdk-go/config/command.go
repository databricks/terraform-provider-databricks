package config

import (
	"context"
	"os/exec"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
)

// Run a command and return the output.
func runCommand(ctx context.Context, cmd string, args []string) ([]byte, error) {
	logger.Debugf(ctx, "Running command: %s %v", cmd, strings.Join(args, " "))
	return exec.Command(cmd, args...).Output()
}
