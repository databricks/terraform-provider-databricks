package privateendpointrule

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// resourceSchema returns the Plugin Framework schema for
// databricks_mws_ncc_private_endpoint_rule. The attribute set matches the
// SDKv2 implementation in mws/resource_mws_ncc_private_endpoint_rule.go
// one-to-one; parity is enforced by `make diff-schema`.
//
// Notes on parity choices:
//   - Fields the SDK auto-generator marks Optional (because the Go field has
//     `omitempty`) and that the SDKv2 customisation does not override stay
//     Optional in the PF schema. For fields the SDKv2 calls `SetComputed()`
//     on, the PF schema is Optional+Computed (matching what the auto-generator
//     plus SetComputed produces in SDKv2).
//   - `gcp_endpoint` is rendered as a `ListNestedBlock` rather than a
//     `SingleNestedAttribute` so the serialised registry schema reproduces
//     the SDKv2 `Block` shape (existing HCL and state files continue to
//     parse without a state upgrader).
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"endpoint_service": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				},
			},
			"resource_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vpc_endpoint_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"connection_state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"creation_time": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"updated_time": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"deactivated": schema.BoolAttribute{
				Optional: true,
			},
			"deactivated_at": schema.Int64Attribute{
				Optional: true,
			},
			"error_message": schema.StringAttribute{
				Optional: true,
			},
		},
		Blocks: map[string]schema.Block{
			"gcp_endpoint": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"psc_endpoint_uri":   schema.StringAttribute{Optional: true},
						"service_attachment": schema.StringAttribute{Optional: true},
					},
				},
			},
		},
	}
}
