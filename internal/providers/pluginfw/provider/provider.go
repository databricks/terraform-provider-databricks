// Package provider exposes the configured-state of the Databricks provider
// to resources and data sources, hiding the underlying common.DatabricksClient
// type behind dependency-getter functions. A resource that only needs the
// account client calls provider.AccountClient and never imports common.
package provider

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

const reportMessage = "please report this to the provider developers"

// providerConfig is an interface that provides the account and workspace clients
// for the provider. The interface should be such that it is implemented by the
// common.DatabricksClient type. This is used to ease testing by building a mock
// implementation of the interface.
type providerConfig interface {
	AccountClient() (*databricks.AccountClient, error)
	WorkspaceClient() (*databricks.WorkspaceClient, error)
}

var _ providerConfig = (*common.DatabricksClient)(nil)

// AccountClient resolves the account-level SDK client from the ProviderData
// passed to a Resource or DataSource Configure method.
//
// The contract:
//
//   - (client, nil-diags) on success.
//   - (nil, nil-diags) when providerData is nil. PF may call a resource or
//     data source Configure before the provider's own Configure has populated
//     ProviderData (during terraform validate, schema generation, and other
//     pre-wire phases). Callers should treat this as a no-op and return
//     without setting state.
//   - (nil, diags-with-error) on a real failure (wrong ProviderData type, or
//     the SDK client could not resolve account credentials).
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

// WorkspaceClient resolves the workspace-level SDK client from the ProviderData
// passed to a Resource or DataSource Configure method.
//
// The contract:
//
//   - (client, nil-diags) on success.
//   - (nil, nil-diags) when providerData is nil. PF may call a resource or
//     data source Configure before the provider's own Configure has populated
//     ProviderData (during terraform validate, schema generation, and other
//     pre-wire phases). Callers should treat this as a no-op and return
//     without setting state.
//   - (nil, diags-with-error) on a real failure (wrong ProviderData type, or
//     the SDK client could not resolve workspace credentials).
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
