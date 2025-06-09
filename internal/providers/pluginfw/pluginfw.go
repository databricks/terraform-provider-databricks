// Package pluginfw contains the changes specific to the plugin framework
//
// Note: This shouldn't depend on internal/providers/sdkv2 or internal/providers
package pluginfw

import (
	"context"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/client"
	providercommon "github.com/databricks/terraform-provider-databricks/internal/providers/common"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func GetDatabricksProviderPluginFramework(opts ...PluginFrameworkOption) provider.Provider {
	providerOptions := &pluginFrameworkOptions{}
	for _, opt := range opts {
		opt.Apply(providerOptions)
	}
	p := &DatabricksProviderPluginFramework{
		sdkV2ResourceFallbacks:   providerOptions.resourceFallbacks,
		sdkV2DataSourceFallbacks: providerOptions.dataSourceFallbacks,
		configCustomizer:         providerOptions.configCustomizer,
	}
	return p
}

type DatabricksProviderPluginFramework struct {
	sdkV2ResourceFallbacks   []string
	sdkV2DataSourceFallbacks []string
	configCustomizer         func(*config.Config) error
}

var _ provider.Provider = (*DatabricksProviderPluginFramework)(nil)

func (p *DatabricksProviderPluginFramework) Resources(ctx context.Context) []func() resource.Resource {
	return getPluginFrameworkResourcesToRegister(p.sdkV2ResourceFallbacks)
}

func (p *DatabricksProviderPluginFramework) DataSources(ctx context.Context) []func() datasource.DataSource {
	return getPluginFrameworkDataSourcesToRegister(p.sdkV2DataSourceFallbacks)
}

func (p *DatabricksProviderPluginFramework) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerSchemaPluginFramework()
}

func (p *DatabricksProviderPluginFramework) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = providercommon.ProviderName
	resp.Version = common.Version()
}

func (p *DatabricksProviderPluginFramework) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	client := p.configureDatabricksClient(ctx, req, resp)
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

// setAttribute sets the attribute value in the SDK config corresponding to the attribute name in the provider configuration.
// It returns true if the attribute was set, false if it was not set (because it was unknown or null), and a diag.Diagnostics object in case of error.
func (p *DatabricksProviderPluginFramework) setAttribute(
	ctx context.Context,
	providerConfig tfsdk.Config,
	attr config.ConfigAttribute,
	cfg *config.Config,
) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	switch attr.Kind {
	case reflect.Bool:
		var attrValue types.Bool
		diags.Append(providerConfig.GetAttribute(ctx, path.Root(attr.Name), &attrValue)...)
		if diags.HasError() {
			return false, diags
		}
		if attrValue.IsNull() || attrValue.IsUnknown() {
			return false, diags
		}
		err := attr.Set(cfg, attrValue.ValueBool())
		if err != nil {
			diags.Append(diag.NewErrorDiagnostic("Failed to set attribute", err.Error()))
			return false, diags
		}
	case reflect.Int:
		var attrValue types.Int64
		diags.Append(providerConfig.GetAttribute(ctx, path.Root(attr.Name), &attrValue)...)
		if diags.HasError() {
			return false, diags
		}
		if attrValue.IsNull() || attrValue.IsUnknown() {
			return false, diags
		}
		err := attr.Set(cfg, int(attrValue.ValueInt64()))
		if err != nil {
			diags.Append(diag.NewErrorDiagnostic("Failed to set attribute", err.Error()))
			return false, diags
		}
	case reflect.String:
		var attrValue types.String
		diags.Append(providerConfig.GetAttribute(ctx, path.Root(attr.Name), &attrValue)...)
		if diags.HasError() {
			return false, diags
		}
		if attrValue.IsNull() || attrValue.IsUnknown() {
			return false, diags
		}
		err := attr.Set(cfg, attrValue.ValueString())
		if err != nil {
			diags.Append(diag.NewErrorDiagnostic(fmt.Sprintf("Failed to set attribute: %s", attr.Name), err.Error()))
			return false, diags
		}
	}
	return true, diags
}

func (p *DatabricksProviderPluginFramework) configureDatabricksClient(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) any {
	cfg := &config.Config{}
	attrsUsed := []string{}
	for _, attr := range config.ConfigAttributes {
		ok, diags := p.setAttribute(ctx, req.Config, attr, cfg)
		resp.Diagnostics.Append(diags...)
		if ok {
			attrsUsed = append(attrsUsed, attr.Name)
		}
	}
	if len(attrsUsed) > 0 {
		sort.Strings(attrsUsed)
		tflog.Info(ctx, fmt.Sprintf("(plugin framework) Attributes specified in provider configuration: %s", strings.Join(attrsUsed, ", ")))
	} else {
		tflog.Info(ctx, "(plugin framework) No attributes specified in provider configuration")
	}
	databricksClient, err := client.PrepareDatabricksClient(ctx, cfg, p.configCustomizer)
	if err != nil {
		resp.Diagnostics.AddError("Failed to configure Databricks client", err.Error())
		return nil
	}
	return databricksClient
}
