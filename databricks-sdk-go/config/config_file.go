package config

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"gopkg.in/ini.v1"
)

// File represents the contents of a databrickscfg file.
type File struct {
	*ini.File

	path string
}

// Path returns the path of the loaded databrickscfg file.
func (f *File) Path() string {
	return f.path
}

func defaultConfigFilePath(path string) string {
	if path == "" {
		path = "~/.databrickscfg"
	}

	return path
}

// LoadFile loads the databrickscfg file at the specified path.
// The function loads ~/.databrickscfg if the specified path is an empty string.
// The function expands ~ to the user's home directory.
func LoadFile(path string) (*File, error) {
	path = defaultConfigFilePath(path)

	// Expand ~ to home directory.
	if strings.HasPrefix(path, "~") {
		homedir, err := getUserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("cannot find homedir: %w", err)
		}
		path = fmt.Sprintf("%s%s", homedir, path[1:])
	}

	iniFile, err := ini.LoadSources(ini.LoadOptions{
		// Passwords may contain '#'. Require a space before '#' for inline comments.
		// See https://github.com/databricks/databricks-sdk-go/issues/594
		SpaceBeforeInlineComment: true,
	}, path)
	if err != nil {
		return nil, err
	}

	return &File{iniFile, path}, err
}

var ConfigFile = configFileLoader{}

type configFileLoader struct{}

func (l configFileLoader) Name() string {
	return "config-file"
}

func (l configFileLoader) Configure(cfg *Config) error {
	// Skip loading config file if some authentication is already explicitly
	// configured directly in the config by a user.
	// See: https://github.com/databricks/databricks-sdk-go/issues/304
	if cfg.Profile == "" && (l.isAnyAuthConfigured(cfg) || cfg.Host != "" || cfg.AzureResourceID != "") {
		return nil
	}

	configFilePath := defaultConfigFilePath(cfg.ConfigFile)
	configFile, err := LoadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// early return for non-configured machines
			logger.Debugf(context.Background(), "%s not found on current host", configFilePath)
			return nil
		}
		return fmt.Errorf("cannot parse config file: %w", err)
	}
	profile, isFallback, err := resolveProfile(cfg.Profile, configFile)
	if err != nil {
		return fmt.Errorf("%s: %w", configFile.Path(), err)
	}
	profileValues := configFile.Section(profile)
	if len(profileValues.Keys()) == 0 {
		if isFallback {
			logger.Debugf(context.Background(), "%s has no %s profile configured", configFile.Path(), profile)
			return nil
		}
		return fmt.Errorf("%s has no %s profile configured", configFile.Path(), profile)
	}
	logger.Debugf(context.Background(), "Loading %s profile from %s", profile, configFile.Path())
	err = ConfigAttributes.ResolveFromStringMapWithSource(cfg, profileValues.KeysHash(), Source{Type: SourceFile, Name: configFile.Path()})
	if err != nil {
		return fmt.Errorf("%s %s profile: %w", configFile.Path(), profile, err)
	}
	// Pin the resolved profile name on cfg, including the legacy fallback to
	// [DEFAULT]. Callers (notably the CLI's U2M OAuth cache, which derives a
	// per-profile keyring entry from cfg.Profile) need to see the same name
	// that was used to load values, otherwise the cache key drifts between
	// a `databricks auth login --profile DEFAULT` write and a later read.
	cfg.Profile = profile
	return nil
}

// settingsSection is the reserved section name for tool settings in
// .databrickscfg. It is not a profile. The CLI writes to this section
// (e.g. via "databricks auth switch") to store the user's default profile.
const settingsSection = "__settings__"

// resolveProfile returns the profile name to use and whether it is a
// legacy fallback to the DEFAULT section:
//   - If requestedProfile is set, returns it with isFallback=false.
//   - If [__settings__].default_profile is set, returns it with isFallback=false.
//   - Otherwise, returns "DEFAULT" with isFallback=true.
//
// Returns an error if the resolved profile is the reserved __settings__
// section name.
func resolveProfile(requestedProfile string, f *File) (string, bool, error) {
	if requestedProfile != "" {
		if requestedProfile == settingsSection {
			return "", false, fmt.Errorf(
				"%s is a reserved section name and cannot be used as a profile", settingsSection)
		}
		return requestedProfile, false, nil
	}
	section, err := f.GetSection(settingsSection)
	if err == nil {
		key, err := section.GetKey("default_profile")
		if err == nil {
			v := strings.TrimSpace(key.String())
			if v != "" {
				if v == settingsSection {
					return "", false, fmt.Errorf(
						"%s is a reserved section name and cannot be used as a profile", settingsSection)
				}
				return v, false, nil
			}
		}
	}
	return "DEFAULT", true, nil
}

func (l configFileLoader) isAnyAuthConfigured(cfg *Config) bool {
	for _, a := range ConfigAttributes {
		if !a.HasAuthAttribute() {
			continue
		}
		if !a.IsZero(cfg) {
			return true
		}
	}
	return false
}
