package config

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

// AuthDetails contains the details of the authentication configuration.
type AuthDetails struct {
	AuthType      string            `json:"auth_type"`
	Host          string            `json:"host,omitempty"`
	Configuration AuthConfiguration `json:"configuration"`
}

func (a *AuthDetails) String() string {
	s := ""
	if a.Host != "" {
		s += fmt.Sprintf(`host=%s`, a.Host)
	}

	if a.AuthType != "" {
		s += fmt.Sprintf(", auth_type=%s", a.AuthType)
	}

	conf := a.Configuration.String()
	if conf != "" {
		s += fmt.Sprintf(`
Configuration:
%s`, conf)
	}

	return s
}

// AuthConfiguration is a map of attribute name to its configuration.
type AuthConfiguration map[string]*AttrConfig

func (c AuthConfiguration) String() string {
	var conf []string
	for _, a := range ConfigAttributes {
		if _, ok := c[a.Name]; ok {
			k := a.Name
			v := c[k]
			conf = append(conf, fmt.Sprintf("- %s=%s", k, v))
		}
	}
	return strings.Join(conf, "\n")
}

type AttrConfig struct {
	Value            string `json:"value"`
	Source           Source `json:"source"`
	AuthTypeMismatch bool   `json:"auth_type_mismatch"`
}

func (a *AttrConfig) String() string {
	s := fmt.Sprintf(`%s (from %s)`, a.Value, a.Source.String())
	if a.AuthTypeMismatch {
		s += " (not used by the current auth type)"
	}
	return s
}

type AuthDetailsOptions int

const (
	ShowSensitive AuthDetailsOptions = iota
)

func (c *Config) GetAuthDetails(opts ...AuthDetailsOptions) AuthDetails {
	var showSensitive bool
	for _, opt := range opts {
		switch opt {
		case ShowSensitive:
			showSensitive = true
		}
	}

	attrSet := make(map[string]*AttrConfig, 0)
	for _, a := range ConfigAttributes {
		if a.IsZero(c) {
			continue
		}

		attrSet[a.Name] = &AttrConfig{
			Value:            getValue(c, &a, showSensitive),
			Source:           c.getSource(&a),
			AuthTypeMismatch: a.HasAuthAttribute() && c.AuthType != "" && !slices.Contains(a.AuthTypes, c.AuthType),
		}
	}

	return AuthDetails{
		AuthType:      c.AuthType,
		Host:          c.Host,
		Configuration: attrSet,
	}
}

func (c *Config) SetAttrSource(attr *ConfigAttribute, source Source) {
	if c.attrSource == nil {
		c.attrSource = make(map[string]Source)
	}
	c.attrSource[attr.Name] = source
}

func (c *Config) getSource(a *ConfigAttribute) Source {
	if c.attrSource == nil {
		return Source{
			Type: SourceDynamicConfig,
		}
	}

	attrSource, ok := c.attrSource[a.Name]
	if !ok {
		return Source{
			Type: SourceDynamicConfig,
		}
	}

	return attrSource
}

func getValue(cfg *Config, a *ConfigAttribute, showSensitive bool) string {
	if !showSensitive && a.Sensitive {
		return "********"
	}

	return a.GetString(cfg)
}
