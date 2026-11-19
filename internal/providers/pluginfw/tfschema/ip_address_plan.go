// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
package tfschema

import (
	"context"
	"net/netip"
	"slices"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/iptypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// withIPAddressPlanModifier prepends ipAddressPlanModifier to IP address
// attributes, so that later modifiers such as RequiresReplace see the prior spelling.
func withIPAddressPlanModifier(attribute schema.Attribute) schema.Attribute {
	switch a := attribute.(type) {
	case schema.StringAttribute:
		if _, ok := a.CustomType.(iptypes.IPAddressType); ok {
			a.PlanModifiers = append([]planmodifier.String{ipAddressPlanModifier{}}, a.PlanModifiers...)
			return a
		}
	case schema.ListAttribute:
		if _, ok := a.ElementType.(iptypes.IPAddressType); ok {
			a.PlanModifiers = append([]planmodifier.List{ipAddressPlanModifier{}}, a.PlanModifiers...)
			return a
		}
	case schema.SetAttribute:
		if _, ok := a.ElementType.(iptypes.IPAddressType); ok {
			a.PlanModifiers = append([]planmodifier.Set{ipAddressPlanModifier{}}, a.PlanModifiers...)
			return a
		}
	case schema.MapAttribute:
		if _, ok := a.ElementType.(iptypes.IPAddressType); ok {
			a.PlanModifiers = append([]planmodifier.Map{ipAddressPlanModifier{}}, a.PlanModifiers...)
			return a
		}
	}
	return attribute
}

// ipAddressPlanModifier plans the prior value of an IP address attribute when
// the configuration only respells it, so that Terraform plans no change. A
// collection keeps its prior value only as a whole: Terraform rejects a plan
// that mixes prior and configured spellings within one collection.
type ipAddressPlanModifier struct{}

func (ipAddressPlanModifier) Description(context.Context) string {
	return "Keeps the prior spelling of an equivalent IP address."
}

func (m ipAddressPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (ipAddressPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// The framework passes scalar values as plain strings.
	config, state := iptypes.IPAddress{StringValue: req.ConfigValue}, iptypes.IPAddress{StringValue: req.StateValue}
	if respellsState(req.ConfigValue, req.PlanValue, req.StateValue) && sameAddress(ctx, config, state) {
		resp.PlanValue = req.StateValue
	}
}

func (ipAddressPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	config, state := req.ConfigValue.Elements(), req.StateValue.Elements()
	if !respellsState(req.ConfigValue, req.PlanValue, req.StateValue) || len(config) != len(state) {
		return
	}
	for i := range config {
		if !sameAddress(ctx, config[i], state[i]) {
			return
		}
	}
	resp.PlanValue = req.StateValue
}

func (ipAddressPlanModifier) PlanModifySet(_ context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	config, okConfig := parseIPAddresses(req.ConfigValue.Elements())
	state, okState := parseIPAddresses(req.StateValue.Elements())
	if !respellsState(req.ConfigValue, req.PlanValue, req.StateValue) || !okConfig || !okState {
		return
	}
	// Each address must occur equally often, in any order.
	slices.SortFunc(config, netip.Addr.Compare)
	slices.SortFunc(state, netip.Addr.Compare)
	if slices.Equal(config, state) {
		resp.PlanValue = req.StateValue
	}
}

func (ipAddressPlanModifier) PlanModifyMap(ctx context.Context, req planmodifier.MapRequest, resp *planmodifier.MapResponse) {
	config, state := req.ConfigValue.Elements(), req.StateValue.Elements()
	if !respellsState(req.ConfigValue, req.PlanValue, req.StateValue) || len(config) != len(state) {
		return
	}
	for key, v := range config {
		if !sameAddress(ctx, v, state[key]) {
			return
		}
	}
	resp.PlanValue = req.StateValue
}

// respellsState reports whether a known state may replace the plan: the plan
// is the known configuration, so no other modifier changed it.
func respellsState(config, plan, state attr.Value) bool {
	return !config.IsNull() && !config.IsUnknown() && !state.IsNull() && !state.IsUnknown() && plan.Equal(config) && !plan.Equal(state)
}

// sameAddress reports whether a and b are IP addresses that spell the same
// address. Collection elements are IPAddress values; a missing map key is nil.
func sameAddress(ctx context.Context, a, b attr.Value) bool {
	x, okX := a.(iptypes.IPAddress)
	y, okY := b.(iptypes.IPAddress)
	if !okX || !okY {
		return false
	}
	equal, diags := x.StringSemanticEquals(ctx, y)
	return equal && !diags.HasError()
}

func parseIPAddresses(values []attr.Value) ([]netip.Addr, bool) {
	addresses := make([]netip.Addr, len(values))
	for i, v := range values {
		ip, ok := v.(iptypes.IPAddress)
		if !ok {
			return nil, false
		}
		address, err := ip.ValueIPAddress()
		if err != nil {
			return nil, false
		}
		addresses[i] = address
	}
	return addresses, true
}

// PreserveEquivalentIPAddressPlan plans the prior state when the configuration
// only respells IP addresses, because the framework has already marked
// unconfigured computed attributes unknown. Plans with errors or a deferral are
// left alone.
func PreserveEquivalentIPAddressPlan(_ context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if resp.Diagnostics.HasError() || resp.Deferred != nil {
		return
	}
	if req.State.Raw.IsNull() || resp.Plan.Raw.IsNull() || !req.Config.Raw.IsFullyKnown() {
		return
	}
	root := planPosition{typ: resp.Plan.Schema.Type(), config: req.Config.Raw, plan: resp.Plan.Raw, state: req.State.Raw}
	if matches, respelled := planMatchesState(root); matches && respelled {
		resp.Plan.Raw = req.State.Raw.Copy()
	}
}

// planPosition is one value of the configuration, plan, and state, with its type.
type planPosition struct {
	typ                 attr.Type
	config, plan, state tftypes.Value
}

// planMatchesState compares the plan with the state under the configuration.
// matches reports whether the plan equals the state at every configured
// position and is unknown or equal to it elsewhere. respelled reports whether
// some configured IP address position differs from the state while the plan
// equals it, which only the IP address modifier causes.
//
// Objects, lists, and maps are compared child by child, because objects can
// hold unconfigured computed attributes. Sets are compared whole, because
// their elements cannot be paired.
func planMatchesState(at planPosition) (matches, respelled bool) {
	config, plan, state := at.config, at.plan, at.state
	// For a computed value without configuration, Terraform proposes the prior
	// value; the framework marks it wholly unknown only because the plan
	// differs elsewhere. Unknown or prior is therefore the no-change plan.
	if config.IsNull() {
		return !plan.IsKnown() || plan.Equal(state), false
	}
	var children []planPosition
	if plan.IsKnown() && !plan.IsNull() && !state.IsNull() {
		switch config.Type().(type) {
		case tftypes.Object, tftypes.Map:
			var c, p, s map[string]tftypes.Value
			if config.As(&c) == nil && plan.As(&p) == nil && state.As(&s) == nil && len(p) == len(c) && len(s) == len(c) {
				// A missing key yields a zero value, which matches only another missing key.
				for key := range c {
					children = append(children, planPosition{childType(at.typ, key), c[key], p[key], s[key]})
				}
			}
		case tftypes.List:
			var c, p, s []tftypes.Value
			if config.As(&c) == nil && plan.As(&p) == nil && state.As(&s) == nil && len(p) == len(c) && len(s) == len(c) {
				for i := range c {
					children = append(children, planPosition{childType(at.typ, ""), c[i], p[i], s[i]})
				}
			}
		}
	}
	if len(children) == 0 {
		elementType := childType(at.typ, "")
		if elementType == nil {
			elementType = at.typ
		}
		_, holdsIPAddresses := elementType.(iptypes.IPAddressType)
		matches = plan.Equal(state)
		return matches, matches && holdsIPAddresses && !config.Equal(state)
	}
	for _, child := range children {
		childMatches, childRespelled := planMatchesState(child)
		if !childMatches {
			return false, false
		}
		respelled = respelled || childRespelled
	}
	return true, respelled
}

// childType returns the type of an object attribute, or of each collection
// element, and nil for other types.
func childType(typ attr.Type, key string) attr.Type {
	switch t := typ.(type) {
	case attr.TypeWithAttributeTypes:
		return t.AttributeTypes()[key]
	case attr.TypeWithElementType:
		return t.ElementType()
	}
	return nil
}
