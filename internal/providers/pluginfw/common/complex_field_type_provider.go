package common

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ComplexFieldTypeProvider interface {
	GetComplexFieldTypes() map[string]reflect.Type
	ToAttrType() types.ObjectType
}
