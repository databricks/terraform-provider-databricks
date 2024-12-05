package common

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComplexFieldTypeProvider interface {
	GetComplexFieldTypes(context.Context) map[string]reflect.Type
}

type ObjectTypable interface {
	ToObjectType(context.Context) types.ObjectType
}
