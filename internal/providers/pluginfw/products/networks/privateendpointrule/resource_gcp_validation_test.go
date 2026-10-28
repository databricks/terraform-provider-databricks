package privateendpointrule

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/terraform-provider-databricks/mws"
	ctymsgpack "github.com/hashicorp/go-cty/cty/msgpack"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	providerschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

// Only the account API is mocked; validation and planning use the real resource.
type privateEndpointRuleTestProvider struct {
	api apiClient
}

func TestGcpTargetValidationDefersUnknownValues(t *testing.T) {
	for _, tt := range []struct {
		name string
		gcp  gcpEndpointModel
	}{
		{"attachment", gcpEndpointModel{ServiceAttachment: types.StringUnknown()}},
		{"all services", gcpEndpointModel{AllVpcScServices: types.BoolUnknown()}},
		{"API list", gcpEndpointModel{GoogleApiEndpoints: []googleApiEndpointsModel{{Endpoints: types.ListUnknown(types.StringType)}}}},
		{"API hostname", gcpEndpointModel{GoogleApiEndpoints: []googleApiEndpointsModel{{
			Endpoints: types.ListValueMust(types.StringType, []attr.Value{types.StringUnknown()}),
		}}}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			config := rawPlan(t, ctx, model{
				NetworkConnectivityConfigId: types.StringValue("ncc_id"),
				GcpEndpoint:                 []gcpEndpointModel{tt.gcp},
			})
			dynamic, err := tfprotov6.NewDynamicValue(config.Raw.Type(), config.Raw)
			require.NoError(t, err)
			server := providerserver.NewProtocol6(&privateEndpointRuleTestProvider{})()
			resp, err := server.ValidateResourceConfig(ctx, &tfprotov6.ValidateResourceConfigRequest{
				TypeName: "databricks_mws_ncc_private_endpoint_rule", Config: &dynamic,
			})
			require.NoError(t, err)
			require.Empty(t, resp.Diagnostics)

			sdkResource := mws.ResourceMwsNccPrivateEndpointRule().ToResource()
			block := sdkResource.CoreConfigSchema()
			value, err := ctymsgpack.Unmarshal(dynamic.MsgPack, block.ImpliedType())
			require.NoError(t, err)
			diags := sdkResource.Validate(terraform.NewResourceConfigShimmed(value, block))
			require.False(t, diags.HasError(), "%v", diags)
		})
	}
}

func (p *privateEndpointRuleTestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "databricks"
}

func (p *privateEndpointRuleTestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = providerschema.Schema{}
}

func (p *privateEndpointRuleTestProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *privateEndpointRuleTestProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func (p *privateEndpointRuleTestProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		func() resource.Resource { return &resourcePrivateEndpointRule{api: p.api, backoff: tightBackoff} },
	}
}

func TestGcpTargetValidationMatchesSDKv2(t *testing.T) {
	googleAPIs := []any{map[string]any{"endpoints": []any{"storage.googleapis.com"}}}
	attachment := "projects/p/regions/r/serviceAttachments/a"
	tests := []struct {
		name  string
		gcp   map[string]any
		valid bool
	}{
		{"non-GCP rule", nil, true},
		{"service attachment", map[string]any{"service_attachment": attachment}, true},
		{"Google APIs", map[string]any{"google_api_endpoints": googleAPIs}, true},
		{"all VPC-SC services", map[string]any{"all_vpc_sc_services": true}, true},
		{"empty target", map[string]any{}, false},
		{"explicit false", map[string]any{"all_vpc_sc_services": false}, false},
		{"empty attachment", map[string]any{"service_attachment": ""}, false},
		{"missing endpoints", map[string]any{"google_api_endpoints": []any{map[string]any{}}}, false},
		{"empty endpoints", map[string]any{"google_api_endpoints": []any{map[string]any{"endpoints": []any{}}}}, false},
		{"attachment and APIs", map[string]any{"service_attachment": attachment, "google_api_endpoints": googleAPIs}, false},
		{"attachment and all services", map[string]any{"service_attachment": attachment, "all_vpc_sc_services": true}, false},
		{"APIs and all services", map[string]any{"google_api_endpoints": googleAPIs, "all_vpc_sc_services": true}, false},
		{"APIs and explicit false", map[string]any{"google_api_endpoints": googleAPIs, "all_vpc_sc_services": false}, false},
		{"APIs and empty attachment", map[string]any{"google_api_endpoints": googleAPIs, "service_attachment": ""}, false},
		{"all services and empty APIs block", map[string]any{"all_vpc_sc_services": true, "google_api_endpoints": []any{map[string]any{}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := map[string]any{"network_connectivity_config_id": "ncc_id"}
			if tt.gcp != nil {
				config["gcp_endpoint"] = []any{tt.gcp}
			} else {
				config["resource_id"] = "resource_id"
				config["group_id"] = "blob"
			}
			diags := mws.ResourceMwsNccPrivateEndpointRule().ToResource().Validate(terraform.NewResourceConfigRaw(config))
			require.Equal(t, tt.valid, !diags.HasError(), "SDKv2 diagnostics: %v", diags)

			// Terraform represents an omitted nested block as an empty list.
			if tt.gcp != nil && tt.gcp["google_api_endpoints"] == nil {
				tt.gcp["google_api_endpoints"] = []any{}
			}
			data, err := json.Marshal(config)
			require.NoError(t, err)
			server := providerserver.NewProtocol6(&privateEndpointRuleTestProvider{})()
			resp, err := server.ValidateResourceConfig(context.Background(), &tfprotov6.ValidateResourceConfigRequest{
				TypeName: "databricks_mws_ncc_private_endpoint_rule",
				Config:   &tfprotov6.DynamicValue{JSON: data},
			})
			require.NoError(t, err)
			hasError := false
			for _, d := range resp.Diagnostics {
				hasError = hasError || d.Severity == tfprotov6.DiagnosticSeverityError
			}
			require.Equal(t, tt.valid, !hasError, "Plugin Framework diagnostics: %v", resp.Diagnostics)
		})
	}
}
