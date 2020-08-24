package qa

import (
	"fmt"
	"os"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func TestEnvironmentTemplate_other_vars(t *testing.T) {
	otherVar := map[string]string{"TEST": "value"}
	res := EnvironmentTemplate(t, `
	resource "user" "me" {
		name  = "{var.TEST}"
	}`, otherVar)
	assert.Equal(t, otherVar["TEST"], FirstKeyValue(t, res, "name"))
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

func TestResourceFixture(t *testing.T) {
	client, server, err := HttpFixtureClient(t, []HTTPFixture{
		{
			Method:   "POST",
			Resource: "/api/2.0/a/b/c",
			Response: HTTPFixture{
				Method: "SOME",
			},
			ExpectedRequest: map[string]bool{
				"check": true,
			},
		},
	})
	defer server.Close()
	assert.NoError(t, err)

	var a HTTPFixture
	err = client.Post("/a/b/c", map[string]bool{
		"check": true,
	}, &a)
	assert.NoError(t, err)
	assert.Equal(t, "SOME", a.Method)
}

func TestResourceFixture_Hint(t *testing.T) {
	t2 := testing.T{}
	client, server, err := HttpFixtureClient(&t2, []HTTPFixture{})
	defer server.Close()
	assert.NoError(t, err)

	var a HTTPFixture
	err = client.Post("/a/b/c", map[string]bool{
		"check": true,
	}, &a)
	assert.Error(t, err)
	assert.True(t, t2.Failed())
}

func TestResourceFixture_Apply(t *testing.T) {
	d, err := ResourceFixture{
		CommandMock: func(commandStr string) (string, error) {
			return "yes", nil
		},
		Resource: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dummy": {
					Type:     schema.TypeBool,
					Required: true,
				},
			},
			Read:   schema.Noop,
			Create: schema.Noop,
			Update: schema.Noop,
			Delete: schema.Noop,
		},
		ID:     "x",
		Read:   true,
		Update: true,
		Delete: true,
		HCL:    `dummy = true`,
	}.Apply(t)
	assert.NoError(t, err)
	assert.Equal(t, true, d.Get("dummy"))
}

func TestResourceFixture_Apply_Fail(t *testing.T) {
	_, err := ResourceFixture{
		CommandMock: func(commandStr string) (string, error) {
			return "yes", nil
		},
		Resource: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"dummy": {
					Type:     schema.TypeBool,
					Required: true,
				},
			},
			Read:   schema.Noop,
			Create: schema.Noop,
			Update: schema.Noop,
			Delete: schema.Noop,
		},
		Create:   true,
		State: map[string]interface{}{
			"check": false,
		},
	}.Apply(t)
	assert.Error(t, err)
}