package u2m

import (
	"context"
	"strings"
	"testing"
)

func TestDiscoveryOAuthArgument_GetCacheKey(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("my-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	got := arg.GetCacheKey()
	if got != "my-profile" {
		t.Errorf("GetCacheKey(): want %q, got %q", "my-profile", got)
	}
}

func TestDiscoveryOAuthArgument_SetAndGetDiscoveredHost(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("test-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	arg.SetDiscoveredHost("https://adb-123.azuredatabricks.net")
	got := arg.GetDiscoveredHost()
	if got != "https://adb-123.azuredatabricks.net" {
		t.Errorf("GetDiscoveredHost(): want %q, got %q", "https://adb-123.azuredatabricks.net", got)
	}
}

func TestDiscoveryOAuthArgument_ZeroValueHasEmptyDiscoveredHost(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("test-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	got := arg.GetDiscoveredHost()
	if got != "" {
		t.Errorf("GetDiscoveredHost(): want empty string, got %q", got)
	}
}

func TestDiscoveryOAuthArgument_ConstructorRejectsEmptyProfile(t *testing.T) {
	_, err := NewBasicDiscoveryOAuthArgument("")
	if err == nil {
		t.Fatal("NewBasicDiscoveryOAuthArgument(): want error for empty profile, got nil")
	}
}

func TestDiscoveryOAuthArgument_NewPersistentAuthRequiresDiscoveryLogin(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("discovery-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	_, err = NewPersistentAuth(context.Background(), WithOAuthArgument(arg))
	if err == nil {
		t.Fatal("NewPersistentAuth(): want error for DiscoveryOAuthArgument without WithDiscoveryLogin, got nil")
	}
	if !strings.Contains(err.Error(), "requires WithDiscoveryLogin") {
		t.Fatalf("NewPersistentAuth(): want discovery login requirement error, got %v", err)
	}
}

func TestDiscoveryOAuthArgument_NewPersistentAuthAcceptsItWithDiscoveryLogin(t *testing.T) {
	arg, err := NewBasicDiscoveryOAuthArgument("discovery-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		context.Background(),
		WithOAuthArgument(arg),
		WithDiscoveryLogin(),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
}

func TestDiscoveryOAuthArgument_DiscoveryLoginRejectsNonDiscoveryArgument(t *testing.T) {
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	_, err = NewPersistentAuth(
		context.Background(),
		WithOAuthArgument(arg),
		WithDiscoveryLogin(),
	)
	if err == nil {
		t.Fatal("NewPersistentAuth(): want error for discovery login without DiscoveryOAuthArgument, got nil")
	}
	if !strings.Contains(err.Error(), "discovery login requires DiscoveryOAuthArgument") {
		t.Fatalf("NewPersistentAuth(): want discovery login validation error, got %v", err)
	}
}

type unsupportedOAuthArgument struct{}

func (u unsupportedOAuthArgument) GetCacheKey() string { return "unsupported" }

func TestDiscoveryOAuthArgument_ValidateArgRejectsUnknownTypes(t *testing.T) {
	_, err := NewPersistentAuth(context.Background(), WithOAuthArgument(unsupportedOAuthArgument{}))
	if err == nil {
		t.Fatal("NewPersistentAuth(): want error for unsupported OAuthArgument type, got nil")
	}
}
