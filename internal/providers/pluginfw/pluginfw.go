// Package pluginfw contains the changes specific to the plugin framework
//
// Note: This shouldn't depend on internal/providers/sdkv2 or internal/providers
package pluginfw

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/commands"
	"github.com/databricks/terraform-provider-databricks/common"
	providercommon "github.com/databricks/terraform-provider-databricks/internal/providers/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/cluster"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/library"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/notificationdestinations"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/qualitymonitor"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/registered_model"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/volume"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func GetDatabricksProviderPluginFramework() provider.Provider {
	p := &DatabricksProviderPluginFramework{}
	return p
}

type DatabricksProviderPluginFramework struct {
}

var _ provider.Provider = (*DatabricksProviderPluginFramework)(nil)

func (p *DatabricksProviderPluginFramework) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		qualitymonitor.ResourceQualityMonitor,
		library.ResourceLibrary,
		sharing.ResourceShare,
	}
}

func (p *DatabricksProviderPluginFramework) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		cluster.DataSourceCluster,
		volume.DataSourceVolumes,
		registered_model.DataSourceRegisteredModel,
		notificationdestinations.DataSourceNotificationDestinations,
	}
}

func (p *DatabricksProviderPluginFramework) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerSchemaPluginFramework()
}

func (p *DatabricksProviderPluginFramework) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = providercommon.ProviderName
	resp.Version = common.Version()
}

func (p *DatabricksProviderPluginFramework) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	client := configureDatabricksClient_PluginFramework(ctx, req, resp)
	resp.DataSourceData = client
	resp.ResourceData = client
}

// Function returns a schema.Schema based on config attributes where each attribute is mapped to the appropriate
// schema type (BoolAttribute, StringAttribute, Int64Attribute).
func providerSchemaPluginFramework() schema.Schema {
	ps := map[string]schema.Attribute{}
	for _, attr := range config.ConfigAttributes {
		switch attr.Kind {
		case reflect.Bool:
			ps[attr.Name] = schema.BoolAttribute{
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		case reflect.String:
			ps[attr.Name] = schema.StringAttribute{
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		case reflect.Int:
			ps[attr.Name] = schema.Int64Attribute{
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		}
	}
	return schema.Schema{
		Attributes: ps,
	}
}

func configureDatabricksClient_PluginFramework(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) any {
	cfg := &config.Config{}
	attrsUsed := []string{}
	for _, attr := range config.ConfigAttributes {
		switch attr.Kind {
		case reflect.Bool:
			var attrValue types.Bool
			diags := req.Config.GetAttribute(ctx, path.Root(attr.Name), &attrValue)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return nil
			}
			err := attr.Set(cfg, attrValue.ValueBool())
			if err != nil {
				resp.Diagnostics.Append(diag.NewErrorDiagnostic("Failed to set attribute", err.Error()))
				return nil
			}
		case reflect.Int:
			var attrValue types.Int64
			diags := req.Config.GetAttribute(ctx, path.Root(attr.Name), &attrValue)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return nil
			}
			err := attr.Set(cfg, int(attrValue.ValueInt64()))
			if err != nil {
				resp.Diagnostics.Append(diag.NewErrorDiagnostic("Failed to set attribute", err.Error()))
				return nil
			}
		case reflect.String:
			var attrValue types.String
			diags := req.Config.GetAttribute(ctx, path.Root(attr.Name), &attrValue)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return nil
			}
			err := attr.Set(cfg, attrValue.ValueString())
			if err != nil {
				resp.Diagnostics.AddError(fmt.Sprintf("Failed to set attribute: %s", attr.Name), err.Error())
				return nil
			}
		}
		if attr.Kind == reflect.String {
			attrsUsed = append(attrsUsed, attr.Name)
		}
	}
	sort.Strings(attrsUsed)
	tflog.Info(ctx, fmt.Sprintf("Explicit and implicit attributes: %s", strings.Join(attrsUsed, ", ")))
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
			log.Printf("[INFO] Changing required auth_type from %s to %s", cfg.AuthType, newer)
			cfg.AuthType = newer
		}
	}
	client, err := client.New(cfg)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic(err.Error(), ""))
		return nil
	}
	pc := &common.DatabricksClient{
		DatabricksClient: client,
	}
	pc.WithCommandExecutor(func(ctx context.Context, client *common.DatabricksClient) common.CommandExecutor {
		return commands.NewCommandsAPI(ctx, client)
	})
	return pc
}
