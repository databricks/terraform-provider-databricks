package config

import "fmt"

// unsupportedGroupRoleAssumption rejects authentication strategies that can
// only obtain normal-access credentials when a group role was requested.
func unsupportedGroupRoleAssumption(cfg *Config, authType string) error {
	if cfg.GroupID == "" {
		return nil
	}
	return fmt.Errorf("auth type %q does not support group role assumption. Use Databricks OAuth authentication", authType)
}
