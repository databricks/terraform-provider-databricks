package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestEnvironmentTemplate(t *testing.T) {
	defer common.CleanupEnvironment()
	err := os.Setenv("USER", qa.RandomName())
	assert.NoError(t, err)

	res := EnvironmentTemplate(t, `
	resource "user" "me" {
		name  = "{env.USER}"
		email = "{env.USER}+{var.RANDOM}@example.com"
	}`)
	assert.Equal(t, os.Getenv("USER"), qa.FirstKeyValue(t, res, "name"))
}

func TestEnvironmentTemplate_other_vars(t *testing.T) {
	otherVar := map[string]string{"TEST": "value"}
	res := EnvironmentTemplate(t, `
	resource "user" "me" {
		name  = "{var.TEST}"
	}`, otherVar)
	assert.Equal(t, otherVar["TEST"], qa.FirstKeyValue(t, res, "name"))
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
