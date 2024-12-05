package common

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ObjectTypable defines the Terraform type to be used for a given object. This allows
// the framework to dynamically instantiate instances of this type, either as an object,
// or as an element in a collection of objects, such as a list or map.
//
// Note that this interface is different from ObjectTypable in the plugin framework itself.
// That interface allows you to do custom serde when interacting with Terraform state.
// The TF SDK in this provider uses types from the plugin framework, and the custom handling
// of serde is mainly done at the boundary between the TFSDK and Databricks Go SDK types.
type ObjectTypable interface {
	// ToObjectType returns the types.ObjectType that describes the object's type in the
	// Terraform value type system.
	ToObjectType(context.Context) types.ObjectType
}
