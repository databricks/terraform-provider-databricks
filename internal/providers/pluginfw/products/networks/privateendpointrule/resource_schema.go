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
func resourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_connectivity_config_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"rule_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_service": schema.StringAttribute{
				Optional: true,
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
				Optional: true,
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
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"domain_names": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
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
				Optional:    true,
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.ConflictsWith(
						path.MatchRoot("group_id"),
						path.MatchRoot("domain_names"),
					),
					listvalidator.SizeAtLeast(1),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"endpoint_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vpc_endpoint_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_state": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"creation_time": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"updated_time": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"deactivated": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"deactivated_at": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"error_message": schema.StringAttribute{
				Computed: true,
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
			// customer rolls back to an older provider.
			"gcp_endpoint": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"psc_endpoint_uri": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						// service_attachment is create-only: toUpdateRequest never
						// adds gcp_endpoint to the update mask, so an edit would be
						// silently dropped and leave a perpetual diff. RequiresReplace
						// matches its create-only siblings (endpoint_service, etc.).
						"service_attachment": schema.StringAttribute{
							Optional: true,
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
