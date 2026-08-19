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

// capturedHostMeta holds the user-supplied and host-advertised workspace_id
// observed during the SDK's single /.well-known/databricks-config resolution. The
// user value is captured before EnsureResolved back-fills it from the host, so a
// mismatch between the two can still be detected afterwards.
type capturedHostMeta struct {
	userWorkspaceID string
	hostWorkspaceID string
}

// installWorkspaceIDCapture wraps cfg.HostMetadataResolver so the provider can
// observe both the user-supplied and host-advertised workspace_id from the single
// metadata fetch the SDK already performs during EnsureResolved. It does NOT add a
// /.well-known/databricks-config request: the wrapper delegates to whatever
// resolver the SDK would otherwise have used, so the total number of metadata
// fetches is unchanged.
//
// Must be installed after any configCustomizer (which may replace the config or
// set its own resolver) and before EnsureResolved. Returns the capture struct the
// wrapper writes into; the caller reads it only after EnsureResolved returns.
func installWorkspaceIDCapture(cfg *config.Config) *capturedHostMeta {
	captured := &capturedHostMeta{}
	// Captured before the wrapper overwrites cfg.HostMetadataResolver below.
	prev := cfg.HostMetadataResolver
	cfg.HostMetadataResolver = func(ctx context.Context, host string) (*config.HostMetadata, error) {
		// Observed post-loaders, pre-back-fill: the effective user value from
		// config/env/profile (including any host query param promoted by the SDK).
		captured.userWorkspaceID = cfg.WorkspaceID
		// Delegate with the SAME precedence the SDK's resolveHostMetadata uses when
		// HostMetadataResolver is nil: a pre-existing resolver, then the
		// DefaultHostMetadataResolverFactory global, then the built-in HTTP fetch.
		// Without the factory tier, installing this wrapper (which always sets
		// HostMetadataResolver) would silently bypass a factory a caller installed.
		delegate := prev
		if delegate == nil {
			if factory := config.DefaultHostMetadataResolverFactory; factory != nil {
				delegate = factory(cfg)
			}
		}
		if delegate == nil {
			delegate = cfg.DefaultHostMetadataResolver()
		}
		meta, err := delegate(ctx, host)
		if err == nil && meta != nil {
			captured.hostWorkspaceID = meta.WorkspaceID
		}
		// Return verbatim: account/unified hosts still need meta's other fields
		// back-filled downstream, and returning the SDK's error unchanged preserves
		// its non-fatal metadata handling.
		return meta, err
	}
	return captured
}

// PrepareDatabricksClient makes some common adjustments to the config that apply in all cases
// and returns a ready-to-use Databricks client. This includes:
//   - mapping deprecated auth types to their newer counterparts
//   - ensuring the config is resolved
//   - setting a default retry timeout if not set
//   - setting a default HTTP timeout if not set
//   - for workspace hosts, reconciling and seeding the workspace_id from host metadata,
//     resolving it eagerly via /Me when the metadata omits it — fatal on failure
//     (see ReconcileWorkspaceIDFromHostMetadata)
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
	// Installed after the customizer (which may replace the config or set its own
	// resolver) and before EnsureResolved, which invokes the resolver exactly once.
	captured := installWorkspaceIDCapture(cfg)
	cfg.EnsureResolved()
	// "none" is a sentinel the Databricks CLI writes to ~/.databrickscfg for
	// account-level profiles with no workspace bound. The SDK passes it through
	// verbatim; treat it as "no workspace_id" so downstream parseWorkspaceID and
	// fallback logic don't choke on it.
	if cfg.WorkspaceID == "none" {
		cfg.WorkspaceID = ""
	}
	client, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	pc := &common.DatabricksClient{
		DatabricksClient: client,
	}
	// For workspace hosts, reconcile the user-supplied workspace_id against the
	// host's discovery metadata and seed the cache so downstream SCIM /Me calls are
	// avoided; when the metadata omits it, resolve eagerly via /Me here (fatal on failure).
	if err := pc.ReconcileWorkspaceIDFromHostMetadata(ctx, pc.HostTypeForTerraform(), captured.userWorkspaceID, captured.hostWorkspaceID); err != nil {
		return nil, err
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return pc, nil
}
