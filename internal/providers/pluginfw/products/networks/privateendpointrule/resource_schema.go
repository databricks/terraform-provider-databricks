package privateendpointrule

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceSchema returns the Plugin Framework schema for
// databricks_mws_ncc_private_endpoint_rule. The attribute set matches the
// SDKv2 implementation in mws/resource_mws_ncc_private_endpoint_rule.go so
// opting in needs no config or state change. `make diff-schema` does NOT
// guard this parity: it dumps the default (SDKv2) provider, and this resource
// is opt-in, so the PF schema below is never introspected by that gate.
// TestSchema_MatchesSDKv2 asserts the parity instead. The one intentional
// divergence is gcp_endpoint.service_attachment: SDKv2 leaves it updatable
// (ForceNew=false), but the Update path never sends gcp_endpoint, so an edit
// is silently dropped into a perpetual diff; PF marks it RequiresReplace.
//
// Attribute descriptions are sourced from the SDK struct comments
// (databricks-sdk-go service/settings) and surfaced through
// `terraform providers schema`, editor hovers, and validation diagnostics.
func resourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The composite resource identifier, in the form `<network_connectivity_config_id>/<rule_id>`. Used for import.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_connectivity_config_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Canonical unique identifier of the Network Connectivity Configuration that owns this private endpoint rule. Changing this forces a new resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"rule_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the private endpoint rule, assigned by Databricks.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The Databricks account ID that owns this private endpoint rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_service": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "(AWS only) The full target AWS endpoint service name that connects to the destination resources of the private endpoint, e.g. `com.amazonaws.us-east-1.s3`. Changing this forces a new resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				// fromAPI collapses a server "" to null, so an explicit "" would
				// fail the post-apply consistency check. Reject it at plan time;
				// "" is never a valid value for these create-only inputs. This is
				// the scalar twin of the domain_names/resource_names SizeAtLeast.
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"group_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "(Azure only) The sub-resource type (group ID) of the target resource, e.g. `blob`, `dfs`, or `sqlServer`. To connect to workspace root storage (root DBFS) you need two rules, one for `blob` and one for `dfs`. Changing this forces a new resource. Conflicts with `domain_names` and `resource_names`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.ConflictsWith(
						path.MatchRoot("domain_names"),
						path.MatchRoot("resource_names"),
					),
					stringvalidator.LengthAtLeast(1),
				},
			},
			"resource_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "(Azure only) The Azure resource ID of the target resource. Changing this forces a new resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"domain_names": schema.ListAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Domain names of the target private link service (Azure) or target resource FQDNs reachable via the VPC endpoint service (AWS). When updating, the full desired list must be specified. Conflicts with `group_id` and `resource_names`.",
				Validators: []validator.List{
					listvalidator.ConflictsWith(
						path.MatchRoot("group_id"),
						path.MatchRoot("resource_names"),
					),
					// fromAPI collapses an empty server list to null, so an
					// explicit `domain_names = []` (a known empty list) would
					// fail Terraform's post-apply consistency check (plan []
					// vs state null). An empty list is not a meaningful value
					// for this attribute; reject it at plan time instead.
					listvalidator.SizeAtLeast(1),
				},
			},
			"resource_names": schema.ListAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "(AWS S3 only) Globally unique S3 bucket names accessed via the VPC endpoint, in the same region as the NCC. When updating, the full desired list must be specified. Conflicts with `group_id` and `domain_names`.",
				Validators: []validator.List{
					listvalidator.ConflictsWith(
						path.MatchRoot("group_id"),
						path.MatchRoot("domain_names"),
					),
					listvalidator.SizeAtLeast(1),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "(AWS S3 only) Whether the private endpoint rule is active. Toggle to activate or deactivate egress access from serverless compute. Can only be set after the rule is successfully created (rules are created disabled), and is not supported on Azure.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The name of the Azure private endpoint resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vpc_endpoint_id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "(AWS) The VPC endpoint ID created by Databricks for this rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_state": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The current connection status of the private endpoint. The rule is effective only when `ESTABLISHED`; new endpoints must be approved in the cloud console before they take effect. One of `PENDING`, `ESTABLISHED`, `REJECTED`, `DISCONNECTED`, `EXPIRED`, `CREATING`, or `CREATE_FAILED`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"creation_time": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Time in epoch milliseconds when this rule was created.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"updated_time": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Time in epoch milliseconds when this rule was last updated.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"deactivated": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether this private endpoint rule is deactivated.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"deactivated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "Time in epoch milliseconds when this rule was deactivated.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"error_message": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Error message describing why the rule is in a `CREATE_FAILED` or otherwise failed state, if any.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Blocks: map[string]schema.Block{
			// gcp_endpoint is logically a single object, but we use
			// ListNestedBlock + SizeAtMost(1) for backward compatibility
			// with SDKv2. Changing the on-disk JSON shape would need a
			// state upgrader, and Terraform can't downgrade state if a
			// customer rolls back to an older provider. The idiomatic PF
			// shape for a 0-or-1 object is SingleNestedAttribute; adopt it
			// only if the SDKv2 resource is retired.
			"gcp_endpoint": schema.ListNestedBlock{
				MarkdownDescription: "(GCP only) Private Service Connect endpoint to a target service attachment.",
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"psc_endpoint_uri": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: "The URI of the created Private Service Connect endpoint.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						// service_attachment is create-only: toUpdateRequest never
						// adds gcp_endpoint to the update mask, so an edit would be
						// silently dropped and leave a perpetual diff. RequiresReplace
						// matches its create-only siblings (endpoint_service, etc.).
						"service_attachment": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "(GCP only) The full URL of the target service attachment, e.g. `projects/my-project/regions/us-east4/serviceAttachments/my-attachment`. Changing this forces a new resource.",
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
		},
	}
}
