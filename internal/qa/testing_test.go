package qa

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentTemplate(t *testing.T) {
	defer common.CleanupEnvironment()
	err := os.Setenv("USER", RandomName())
	assert.NoError(t, err)

	res := EnvironmentTemplate(t, `
	resource "user" "me" {
		name  = "{env.USER}"
		email = "{env.USER}+{var.RANDOM}@example.com"
	}`)
	assert.Equal(t, os.Getenv("USER"), FirstKeyValue(t, res, "name"))
}

func TestEnvironmentTemplate_unset_env(t *testing.T) {
	res, err := environmentTemplate(t, `
	resource "user" "me" {
		name  = "{env.USER}"
		email = "{env.USER}+{var.RANDOM}@example.com"
	}`)
	assert.Equal(t, "", res)
	assert.Errorf(t, err, fmt.Sprintf("please set %d variables and restart.", 2))
}
