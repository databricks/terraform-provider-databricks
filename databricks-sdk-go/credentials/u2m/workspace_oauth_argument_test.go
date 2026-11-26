package u2m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicWorkspaceOAuthArgument_GetCacheKey(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		profile     string
		wantKey     string
		wantHostKey string
	}{
		{
			name:        "without profile returns host-based key",
			host:        "https://myworkspace.cloud.databricks.com",
			wantKey:     "https://myworkspace.cloud.databricks.com",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name:        "with profile returns profile name",
			host:        "https://myworkspace.cloud.databricks.com",
			profile:     "my-profile",
			wantKey:     "my-profile",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name:        "empty profile returns host-based key",
			host:        "https://myworkspace.cloud.databricks.com",
			profile:     "",
			wantKey:     "https://myworkspace.cloud.databricks.com",
			wantHostKey: "https://myworkspace.cloud.databricks.com",
		},
		{
			name:        "localhost without profile",
			host:        "http://127.0.0.1:5656",
			wantKey:     "http://127.0.0.1:5656",
			wantHostKey: "http://127.0.0.1:5656",
		},
		{
			name:        "localhost with profile",
			host:        "http://127.0.0.1:5656",
			profile:     "local-profile",
			wantKey:     "local-profile",
			wantHostKey: "http://127.0.0.1:5656",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arg BasicWorkspaceOAuthArgument
			var err error
			if tt.profile != "" {
				arg, err = NewProfileWorkspaceOAuthArgument(tt.host, tt.profile)
			} else {
				arg, err = NewBasicWorkspaceOAuthArgument(tt.host)
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, arg.GetCacheKey())
			assert.Equal(t, tt.wantHostKey, arg.GetHostCacheKey())
		})
	}
}

func TestNewProfileWorkspaceOAuthArgument_ValidatesHost(t *testing.T) {
	_, err := NewProfileWorkspaceOAuthArgument("http://some-host.com", "my-profile")
	assert.EqualError(t, err, "host must start with 'https://': http://some-host.com")

	_, err = NewProfileWorkspaceOAuthArgument("https://some-host.com/", "my-profile")
	assert.EqualError(t, err, "host must not have a trailing slash: https://some-host.com/")
}

func TestBasicWorkspaceOAuthArgument_ImplementsHostCacheKeyProvider(t *testing.T) {
	arg, err := NewBasicWorkspaceOAuthArgument("https://myworkspace.cloud.databricks.com")
	assert.NoError(t, err)
	var _ HostCacheKeyProvider = arg
}

func TestValidateHost(t *testing.T) {
	tests := []struct {
		host string
		want string
	}{
		// Valid hosts
		{"https://some-host.com", ""},
		{"http://127.0.0.1", ""},
		{"http://127.0.0.1:5656", ""},

		// Invalid hosts
		{"http://some-host.com", "host must start with 'https://': http://some-host.com"},
		{"https://some-host.com/", "host must not have a trailing slash: https://some-host.com/"},
	}

	for _, test := range tests {
		err := validateHost(test.host)
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}

		_, err = NewBasicWorkspaceOAuthArgument(test.host)
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}

		_, err = NewBasicAccountOAuthArgument(test.host, "123")
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}
	}
}
