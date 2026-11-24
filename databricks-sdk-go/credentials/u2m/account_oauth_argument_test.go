package u2m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProfileAccountOAuthArgument_ValidatesHost(t *testing.T) {
	_, err := NewProfileAccountOAuthArgument("http://insecure.com", "abc", "my-profile")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "host must start with 'https://'")

	_, err = NewProfileAccountOAuthArgument("https://accounts.cloud.databricks.com/", "abc", "my-profile")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "host must not have a trailing slash")
}

func TestBasicAccountOAuthArgument_ProfileCacheKeys(t *testing.T) {
	tests := []struct {
		name        string
		host        string
		accountID   string
		profile     string
		wantKey     string
		wantHostKey string
	}{
		{
			name:        "without profile returns host-based key",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "abc",
			wantKey:     "https://accounts.cloud.databricks.com/oidc/accounts/abc",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name:        "with profile returns profile name",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "abc",
			profile:     "my-profile",
			wantKey:     "my-profile",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name:        "empty profile returns host-based key",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "abc",
			profile:     "",
			wantKey:     "https://accounts.cloud.databricks.com/oidc/accounts/abc",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
		{
			name:        "different profiles on same host get different keys",
			host:        "https://accounts.cloud.databricks.com",
			accountID:   "abc",
			profile:     "profile-a",
			wantKey:     "profile-a",
			wantHostKey: "https://accounts.cloud.databricks.com/oidc/accounts/abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var arg BasicAccountOAuthArgument
			var err error
			if tt.profile != "" {
				arg, err = NewProfileAccountOAuthArgument(tt.host, tt.accountID, tt.profile)
			} else {
				arg, err = NewBasicAccountOAuthArgument(tt.host, tt.accountID)
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantKey, arg.GetCacheKey())
			assert.Equal(t, tt.wantHostKey, arg.GetHostCacheKey())
		})
	}
}

func TestBasicAccountOAuthArgument_ImplementsHostCacheKeyProvider(t *testing.T) {
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "abc")
	assert.NoError(t, err)
	var _ HostCacheKeyProvider = arg
}
