package common

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertErrorStartsWith(t *testing.T, err error, message string) bool {
	return assert.True(t, strings.HasPrefix(err.Error(), message), err.Error())
}

func configureAndAuthenticate(dc *DatabricksClient) (*DatabricksClient, error) {
	err := dc.Configure()
	if err != nil {
		return dc, err
	}
	return dc, dc.Authenticate()
}

func TestDatabricksClientConfigure_Nothing(t *testing.T) {
	defer CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")

	_, err := configureAndAuthenticate(&DatabricksClient{})
	AssertErrorStartsWith(t, err, "authentication is not configured for provider")
}

func TestDatabricksClientConfigure_BasicAuth_NoHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Username: "foo",
		Password: "bar",
	})

	AssertErrorStartsWith(t, err, "cannot configure auth: host is empty, but is required by basic_auth")
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_BasicAuth(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:     "https://localhost:443",
		Username: "foo",
		Password: "bar",
	})

	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_HostWithoutScheme(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:  "localhost:443",
		Token: "...",
	})

	assert.Equal(t, "...", dc.Token)
	assert.Equal(t, "https://localhost:443", dc.Host)
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_Token_NoHost(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Token: "dapi345678",
	})

	AssertErrorStartsWith(t, err, "cannot configure auth: host is empty, but is required by token")
	assert.Equal(t, "dapi345678", dc.Token)
}

func TestDatabricksClientConfigure_HostTokensTakePrecedence(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
}

func TestDatabricksClientConfigure_BasicAuthTakePrecedence(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		Host:       "foo",
		Token:      "connfigured",
		Username:   "foo",
		Password:   "bar",
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Zm9vOmJhcg==", dc.Token)
}

func TestDatabricksClientConfigure_ConfigRead(t *testing.T) {
	dc, err := configureAndAuthenticate(&DatabricksClient{
		ConfigFile: "testdata/.databrickscfg",
	})
	assert.NoError(t, err)
	assert.Equal(t, "PT0+IC9kZXYvdXJhbmRvbSA8PT0KYFZ", dc.Token)
}

func TestDatabricksClientConfigure_NoHostGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "nohost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_NoTokenGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "notoken",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_InvalidProfileGivesError(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.databrickscfg",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_MissingFile(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/.invalid file",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

func TestDatabricksClientConfigure_InvalidConfigFilePath(t *testing.T) {
	_, err := configureAndAuthenticate(&DatabricksClient{
		Token:      "connfigured",
		ConfigFile: "testdata/policy01.json",
		Profile:    "invalidhost",
	})
	assert.Error(t, err)
}

// func TestDatabricksClientConfigure_InvalidHome(t *testing.T) {
// 	defer CleanupEnvironment()()
// 	os.Setenv("HOME", "whatever")
// 	_, err := configureAndAuthenticate(&DatabricksClient{
// 		Profile: "invalidhost",
// 	})
// 	assert.EqualError(t, err, ".")
// }

func TestDatabricksClient_FormatURL(t *testing.T) {
	client := DatabricksClient{Host: "https://some.host"}
	assert.Equal(t, "https://some.host/#job/123", client.FormatURL("#job/123"))
}
