package service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	dc := DatabricksClient{
		Username: "foo",
		Password: "bar",
	}
	err := dc.Configure("dev")

	assertErrorStartsWith(t, err, "Host is empty, but is required by basic_auth")
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc := DatabricksClient{
		Host:     "https://localhost:443",
		Username: "foo",
		Password: "bar",
	}
	err := dc.Configure("dev")

	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc := DatabricksClient{
		Host:  "localhost:443",
		Token: "...",
	}
	err := dc.Configure("dev")

	assert.Equal(t, "...", dc.Token)
	assert.Equal(t, "https://localhost:443", dc.Host)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	dc := DatabricksClient{
		Token: "dapi345678",
	}
	err := dc.Configure("dev")

	assertErrorStartsWith(t, err, "Host is empty, but is required by token")
	assert.Equal(t, "dapi345678", dc.Token)
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	dc := DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
	}
	err := dc.Configure("dev")
	assert.Nil(t, err)
}

func TestDatabricksClientConfigure_BasicAuthTakePrecedence(t *testing.T) {
	dc := DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		Username:   "foo",
		Password:   "bar",
		ConfigFile: "testdata/.databrickscfg",
	}
	err := dc.Configure("dev")
	assert.Nil(t, err)
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc := DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
	}
	err := dc.Configure("dev")
	assert.Nil(t, err)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	dc := DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "nohost",
	}
	err := dc.Configure("dev")
	assert.NotNil(t, err)
}

func TestDatabricksClientConfigure_NoTokenGivesError(t *testing.T) {
	dc := DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	}
	err := dc.Configure("dev")
	assert.NotNil(t, err)
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	dc := DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "invalidhost",
	}
	err := dc.Configure("dev")
	assert.NotNil(t, err)
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	dc := DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.invalid file",
		Profile:    "invalidhost",
	}
	err := dc.Configure("dev")
	assert.NotNil(t, err)
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	dc := DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/policy01.json",
		Profile:    "invalidhost",
	}
	err := dc.Configure("dev")
	assert.NotNil(t, err)
}
