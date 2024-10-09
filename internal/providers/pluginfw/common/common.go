package common

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// ConfigureDataSource is a helper function for configuring a general data source.
// It returns the DatabricksClient if it can be successfully fetched from the ProviderData in the request;
// otherwise, the error is appended to the diagnostics of the response.
func ConfigureDataSource(req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) *common.DatabricksClient {
	// Nil case for acceptance tests.
	if req.ProviderData == nil {
		return nil
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	return client
}

// ConfigureDataSource is a helper function for configuring a general resource.
// It returns the DatabricksClient if it can be successfully fetched from the ProviderData in the request;
// otherwise, the error is appended to the diagnostics of the response.
func ConfigureResource(req resource.ConfigureRequest, resp *resource.ConfigureResponse) *common.DatabricksClient {
	// Nil case for acceptance tests.
	if req.ProviderData == nil {
		return nil
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	return client
}

// GetDatabricksStagingName returns the resource name for a given resource with _pluginframework suffix.
// Once a migrated resource is ready to be used as default, the Metadata method for that resource should be updated to use GetDatabricksProductionName.
func GetDatabricksStagingName(name string) string {
	return fmt.Sprintf("databricks_%s_pluginframework", name)
}

func GetDatabricksProductionName(name string) string {
	return fmt.Sprintf("databricks_%s", name)
}

func SetResourceNameInContext(ctx context.Context, resourceName string) context.Context {
	return useragent.InContext(ctx, "resource", resourceName)
}

func SetDataSourceNameInContext(ctx context.Context, dataSourceName string) context.Context {
	return useragent.InContext(ctx, "data", dataSourceName)
}
