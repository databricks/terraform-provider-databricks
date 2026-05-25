// Package provider gives resources and data sources access to the
// configured Databricks API clients.
package provider

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

const reportMessage = "please report this to the provider developers"

// providerConfig is the minimum surface the resolvers need from a configured
// provider. *common.DatabricksClient implements it (asserted below); tests
// substitute a fake.
type providerConfig interface {
	AccountClient() (*databricks.AccountClient, error)
	WorkspaceClient() (*databricks.WorkspaceClient, error)
}

var _ providerConfig = (*common.DatabricksClient)(nil)

// AccountClient resolves the account-level SDK client from ProviderData.
//
// Returns:
//   - (client, nil) on success.
//   - (nil, nil) when providerData is nil. The framework can call a resource's
//     Configure before the provider's own Configure has set ProviderData (during
//     terraform validate, schema generation, etc.). Callers treat this as a
//     no-op and return without setting state.
//   - (nil, error diagnostics) when ProviderData is the wrong type or the SDK
//     can't resolve account credentials.
func AccountClient(providerData any) (*databricks.AccountClient, diag.Diagnostics) {
	if providerData == nil {
		return nil, nil
	}

	client, ok := providerData.(providerConfig)
	if !ok {
		d := diag.Diagnostics{}
		d.AddError("unexpected provider data type", fmt.Sprintf("expected *common.DatabricksClient, got %T (%s)", providerData, reportMessage))
		return nil, d
	}

	acc, err := client.AccountClient()
	if err != nil {
		d := diag.Diagnostics{}
		d.AddError("failed to resolve account client", err.Error())
		return nil, d
	}

	return acc, nil
}

// WorkspaceClient resolves the workspace-level SDK client from ProviderData.
//
// Returns:
//   - (client, nil) on success.
//   - (nil, nil) when providerData is nil. The framework can call a resource's
//     Configure before the provider's own Configure has set ProviderData (during
//     terraform validate, schema generation, etc.). Callers treat this as a
//     no-op and return without setting state.
//   - (nil, error diagnostics) when ProviderData is the wrong type or the SDK
//     can't resolve workspace credentials.
func WorkspaceClient(providerData any) (*databricks.WorkspaceClient, diag.Diagnostics) {
	if providerData == nil {
		return nil, nil
	}

	client, ok := providerData.(providerConfig)
	if !ok {
		d := diag.Diagnostics{}
		d.AddError("unexpected provider data type", fmt.Sprintf("expected *common.DatabricksClient, got %T (%s)", providerData, reportMessage))
		return nil, d
	}

	wc, err := client.WorkspaceClient()
	if err != nil {
		d := diag.Diagnostics{}
		d.AddError("failed to resolve workspace client", err.Error())
		return nil, d
	}

	return wc, nil
}
