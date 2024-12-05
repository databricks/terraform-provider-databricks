package common

import (
	"context"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComplexFieldTypeProvider interface {
	GetComplexFieldTypes() map[string]reflect.Type
	ToAttrType(context.Context) types.ObjectType
}
