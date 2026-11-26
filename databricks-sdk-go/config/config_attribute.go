package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// getenv is the function used to read environment variables.
// It defaults to os.Getenv but can be overwritten in tests.
var getenv = os.Getenv

// getUserHomeDir is the function used to get user home directory.
// It defaults to os.UserHomeDir but can be overwritten in tests.
var getUserHomeDir = os.UserHomeDir

type Source struct {
	Type SourceType `json:"type"`
	Name string     `json:"name,omitempty"`
}

func (s *Source) String() string {
	if s.Name == "" {
		return string(s.Type)
	}
	return fmt.Sprintf("%s %s", s.Name, s.Type)
}

type SourceType string

const (
	SourceEnv           SourceType = "environment variable"
	SourceFile          SourceType = "config file"
	SourceDynamicConfig SourceType = "dynamic configuration"
)

// ConfigAttribute provides generic way to work with Config configuration
// attributes and parses `name`, `env`, and `auth` field tags.
//
// Internal: this field can become unexported in the future
type ConfigAttribute struct {
	Name      string
	Kind      reflect.Kind
	EnvVars   []string
	Auth      string
	AuthTypes []string
	Sensitive bool
	Internal  bool
	num       int
}

func (a *ConfigAttribute) ReadEnv() (string, string) {
	for _, envName := range a.EnvVars {
		v := getenv(envName)
		if v == "" {
			continue
		}
		return v, envName
	}
	return "", ""
}

func (a *ConfigAttribute) SetS(cfg *Config, v string) error {
	switch a.Kind {
	case reflect.String:
		return a.Set(cfg, v)
	case reflect.Bool:
		vv, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	case reflect.Int:
		vv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	case reflect.Slice:
		rawParts := strings.Split(v, ",")
		var parts []string
		for _, part := range rawParts {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				parts = append(parts, trimmed)
			}
		}
		return a.Set(cfg, parts)
	default:
		return fmt.Errorf("cannot set %s of unknown type %s",
			a.Name, reflectKind(a.Kind))
	}
}

func (a *ConfigAttribute) Set(cfg *Config, i interface{}) error {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	switch a.Kind {
	case reflect.String:
		field.SetString(i.(string))
	case reflect.Bool:
		field.SetBool(i.(bool))
	case reflect.Int:
		field.SetInt(int64(i.(int)))
	case reflect.Slice:
		field.Set(reflect.ValueOf(i.([]string)))
	default:
		// must extensively test with providerFixture to avoid this one
		return fmt.Errorf("cannot set %s of unknown type %s", a.Name, reflectKind(a.Kind))
	}
	return nil
}

func (a *ConfigAttribute) IsZero(cfg *Config) bool {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	return field.IsZero()
}

func (a *ConfigAttribute) GetString(cfg *Config) string {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	switch a.Kind {
	case reflect.Slice:
		parts := field.Interface().([]string)
		return strings.Join(parts, ",")
	default:
		return fmt.Sprintf("%v", field.Interface())
	}
}

func (a *ConfigAttribute) HasAuthAttribute() bool {
	return a.Auth != ""
}
