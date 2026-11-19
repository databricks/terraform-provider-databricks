package terraformtest

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
)

// journeyAuth chooses the provider's credentials in place of the fixture PAT. Key
// type is the SDK auth type; the other keys are the SDK's credential attributes,
// named as in the provider block.
type journeyAuth map[string]string

// replayAuthTypes are the auth types that make every request, token requests
// included, through the provider's transport, so fixtures can serve all of them.
var replayAuthTypes = []string{"pat", "basic", "oauth-m2m", "azure-client-secret",
	"azure-msi"}

func (a journeyAuth) validate() error {
	if a == nil {
		return nil
	}
	if !slices.Contains(replayAuthTypes, a["type"]) {
		return fmt.Errorf("auth.type: expected one of %s, got %q",
			strings.Join(replayAuthTypes, ", "), a["type"])
	}
	for _, name := range slices.Sorted(maps.Keys(a)) {
		if name == "type" {
			continue
		}
		attr := credentialAttribute(name)
		if attr == nil {
			return fmt.Errorf("auth.%s: not a credential attribute of the provider", name)
		}
		if a[name] == "" {
			return fmt.Errorf("auth.%s: value must not be empty", name)
		}
		if err := attr.SetS(&config.Config{}, a[name]); err != nil {
			return fmt.Errorf("auth.%s: %w", name, err)
		}
	}
	return nil
}

// withCredentials replaces the fixture PAT that customize sets with the journey's
// credentials. It clears the credential attributes of the provider block first, so
// that auth is their only source.
func (a journeyAuth) withCredentials(customize func(*config.Config) error) func(*config.Config) error {
	if a == nil {
		return customize
	}
	return func(cfg *config.Config) error {
		if err := customize(cfg); err != nil {
			return err
		}
		for i := range config.ConfigAttributes {
			if err := clearCredential(cfg, &config.ConfigAttributes[i]); err != nil {
				return err
			}
		}
		for name, value := range a {
			if name == "type" {
				continue
			}
			if err := credentialAttribute(name).SetS(cfg, value); err != nil {
				return err
			}
		}
		// With an auth type and no strategy, the SDK configures only that type, as it
		// does for auth_type in a provider block.
		cfg.AuthType, cfg.Credentials = a["type"], nil
		return nil
	}
}

func credentialAttribute(name string) *config.ConfigAttribute {
	for i := range config.ConfigAttributes {
		if attr := &config.ConfigAttributes[i]; attr.Name == name && attr.HasAuthAttribute() {
			return attr
		}
	}
	return nil
}

func clearCredential(cfg *config.Config, attr *config.ConfigAttribute) error {
	switch {
	case !attr.HasAuthAttribute():
		return nil
	case attr.Kind == reflect.String:
		return attr.Set(cfg, "")
	case attr.Kind == reflect.Bool:
		return attr.Set(cfg, false)
	}
	return fmt.Errorf("cannot clear credential attribute %s", attr.Name)
}
