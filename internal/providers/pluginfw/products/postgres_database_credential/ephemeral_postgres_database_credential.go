package postgres_database_credential

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/databricks/databricks-sdk-go/common/types/duration"
	sdktime "github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/service/postgres"
	"github.com/databricks/terraform-provider-databricks/common"
	pluginfwcommon "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/common"
	pluginfwcontext "github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	ephemeralResourceName = "postgres_database_credential"
	defaultTTL            = time.Hour
	minimumTTL            = 5 * time.Minute
	maximumTTL            = time.Hour
)

var endpointPattern = regexp.MustCompile(`^projects/[^/]+/branches/[^/]+/endpoints/[^/]+$`)

var (
	_ ephemeral.EphemeralResource              = (*databaseCredentialResource)(nil)
	_ ephemeral.EphemeralResourceWithConfigure = (*databaseCredentialResource)(nil)
	_ validator.String                         = databaseCredentialTTLValidator{}
)

type databaseCredentialResource struct {
	client *common.DatabricksClient
}

type databaseCredentialModel struct {
	Endpoint   types.String `tfsdk:"endpoint"`
	TTL        types.String `tfsdk:"ttl"`
	Token      types.String `tfsdk:"token"`
	ExpireTime types.String `tfsdk:"expire_time"`
}

func New() ephemeral.EphemeralResource {
	return &databaseCredentialResource{}
}

func (r *databaseCredentialResource) Metadata(
	ctx context.Context,
	req ephemeral.MetadataRequest,
	resp *ephemeral.MetadataResponse,
) {
	resp.TypeName = pluginfwcommon.GetDatabricksProductionName(ephemeralResourceName)
}

func (r *databaseCredentialResource) Schema(
	ctx context.Context,
	req ephemeral.SchemaRequest,
	resp *ephemeral.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		Description: "Generates a short-lived OAuth credential for a Lakebase Postgres endpoint.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Required:    true,
				Description: "The Lakebase endpoint resource name in projects/{project_id}/branches/{branch_id}/endpoints/{endpoint_id} format.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(endpointPattern, "must be a Lakebase endpoint resource name"),
				},
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The credential lifetime as a Go duration. Must be between 5 minutes and 1 hour. Defaults to 1 hour.",
				Validators:  []validator.String{databaseCredentialTTLValidator{}},
			},
			"token": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "The OAuth token to use as the Postgres password.",
			},
			"expire_time": schema.StringAttribute{
				Computed:    true,
				Description: "The UTC timestamp when the credential expires.",
			},
		},
	}
}

func (r *databaseCredentialResource) Configure(
	ctx context.Context,
	req ephemeral.ConfigureRequest,
	resp *ephemeral.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*common.DatabricksClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Ephemeral Resource Configure Type",
			fmt.Sprintf("Expected *common.DatabricksClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *databaseCredentialResource) Open(
	ctx context.Context,
	req ephemeral.OpenRequest,
	resp *ephemeral.OpenResponse,
) {
	ctx = pluginfwcontext.SetUserAgentInEphemeralResourceContext(ctx, ephemeralResourceName)
	var config databaseCredentialModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ttl, diags := credentialTTL(config.TTL)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if config.TTL.IsNull() {
		config.TTL = types.StringValue(defaultTTL.String())
	}

	credential, diags := r.generateCredential(ctx, config.Endpoint.ValueString(), ttl)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.Token = types.StringValue(credential.Token)
	config.ExpireTime = types.StringPointerValue(timePointerToString(credential.ExpireTime))
	resp.Diagnostics.Append(resp.Result.Set(ctx, &config)...)
}

func (r *databaseCredentialResource) generateCredential(
	ctx context.Context,
	endpoint string,
	ttl time.Duration,
) (*postgres.DatabaseCredential, diag.Diagnostics) {
	var diags diag.Diagnostics
	if r.client == nil {
		diags.AddError("Provider not configured", "The Databricks provider must be configured before generating a database credential.")
		return nil, diags
	}
	workspaceClient, clientDiags := r.client.GetWorkspaceClientForUnifiedProviderWithDiagnostics(ctx, "")
	diags.Append(clientDiags...)
	if diags.HasError() {
		return nil, diags
	}
	credential, err := workspaceClient.Postgres.GenerateDatabaseCredential(ctx, postgres.GenerateDatabaseCredentialRequest{
		Endpoint: endpoint,
		Ttl:      duration.New(ttl),
	})
	if err != nil {
		diags.AddError("Unable to generate Lakebase database credential", err.Error())
		return nil, diags
	}
	return credential, diags
}

func timePointerToString(value *sdktime.Time) *string {
	if value == nil {
		return nil
	}
	formatted := value.AsTime().UTC().Format(time.RFC3339)
	return &formatted
}

type databaseCredentialTTLValidator struct{}

func (databaseCredentialTTLValidator) Description(context.Context) string {
	return "value must be a duration between 5 minutes and 1 hour"
}

func (v databaseCredentialTTLValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (databaseCredentialTTLValidator) ValidateString(
	ctx context.Context,
	req validator.StringRequest,
	resp *validator.StringResponse,
) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	_, diags := credentialTTL(req.ConfigValue)
	for _, diagnostic := range diags {
		resp.Diagnostics.AddAttributeError(req.Path, diagnostic.Summary(), diagnostic.Detail())
	}
}

func credentialTTL(value types.String) (time.Duration, diag.Diagnostics) {
	var diags diag.Diagnostics
	if value.IsNull() {
		return defaultTTL, diags
	}
	if value.IsUnknown() {
		diags.AddError("Unknown credential TTL", "The credential TTL must be known before the credential can be generated.")
		return 0, diags
	}
	ttl, err := time.ParseDuration(value.ValueString())
	if err != nil {
		diags.AddError("Invalid credential TTL", fmt.Sprintf("The credential TTL must be a valid Go duration: %s", err))
		return 0, diags
	}
	if ttl < minimumTTL || ttl > maximumTTL {
		diags.AddError("Invalid credential TTL", "The credential TTL must be between 5 minutes and 1 hour.")
		return 0, diags
	}
	return ttl, diags
}
