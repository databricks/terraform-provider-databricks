package client

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PrepareDatabricksClient makes some common adjustments to the config that apply in all cases
// and returns a ready-to-use Databricks client. This includes:
// - mapping deprecated auth types to their newer counterparts
// - ensuring the config is resolved
// - setting a default retry timeout if not set
// - setting a default HTTP timeout if not set
//
// TODO: this should be colocated with the definition of DatabricksClient in common/client.go, but
// this isn't possible without introducing a circular dependency. Fixing this will require refactoring
// DatabricksClient out of the common package.
func PrepareDatabricksClient(ctx context.Context, cfg *config.Config, configCustomizer func(*config.Config) error) (*common.DatabricksClient, error) {
	if cfg.AuthType != "" {
		// mapping from previous Google authentication types
		// and current authentication types from Databricks Go SDK
		oldToNewerAuthType := map[string]string{
			"google-creds":     "google-credentials",
			"google-accounts":  "google-id",
			"google-workspace": "google-id",
		}
		newer, ok := oldToNewerAuthType[cfg.AuthType]
		if ok {
			tflog.Info(ctx, fmt.Sprintf("Changing required auth_type from %s to %s", cfg.AuthType, newer))
			cfg.AuthType = newer
		}
	}
	cfg.EnsureResolved()
	// Unless set explicitly, the provider will retry indefinitely until context is cancelled
	// by either a timeout or interrupt.
	if cfg.RetryTimeoutSeconds == 0 {
		cfg.RetryTimeoutSeconds = -1
	}
	// If not set, the default provider timeout is 65 seconds. Most APIs have a server-side timeout of 60 seconds.
	// The additional 5 seconds is to account for network latency.
	if cfg.HTTPTimeoutSeconds == 0 {
		cfg.HTTPTimeoutSeconds = 65
	}
	if configCustomizer != nil {
		err := configCustomizer(cfg)
		if err != nil {
			return nil, err
		}
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	pc := &common.DatabricksClient{
		DatabricksClient: client,
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return pc, nil
}
