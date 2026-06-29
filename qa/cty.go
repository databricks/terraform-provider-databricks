package qa

import (
	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// https://github.com/hashicorp/terraform-plugin-sdk/blob/866d0b19a878fe2241fa8e008bee8c6cb8b2c32b/internal/configs/hcl2shim/values.go#L130-L194
func hcl2ValueFromConfigValue(v interface{}) cty.Value {
	if v == nil {
		return cty.NullVal(cty.DynamicPseudoType)
	}

	switch tv := v.(type) {
	case bool:
		return cty.BoolVal(tv)
	case string:
		return cty.StringVal(tv)
	case int:
		return cty.NumberIntVal(int64(tv))
	case float64:
		return cty.NumberFloatVal(tv)
	case []interface{}:
		vals := make([]cty.Value, len(tv))
		for i, ev := range tv {
			vals[i] = hcl2ValueFromConfigValue(ev)
		}
		return cty.TupleVal(vals)
	case map[string]interface{}:
		vals := map[string]cty.Value{}
		for k, ev := range tv {
			vals[k] = hcl2ValueFromConfigValue(ev)
		}
		return cty.ObjectVal(vals)
	default:
		return cty.NullVal(cty.DynamicPseudoType)
	}
}

func makeResourceRawConfig(config *terraform.ResourceConfig, resource *schema.Resource) cty.Value {
	original := hcl2ValueFromConfigValue(config.Raw)
	coerced, err := resource.CoreConfigSchema().CoerceValue(original)
	if err != nil {
		return cty.NullVal(cty.DynamicPseudoType)
	}
	return coerced
}
