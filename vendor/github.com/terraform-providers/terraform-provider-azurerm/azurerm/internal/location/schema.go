package location

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azuread/azuread/helpers/validate"
)

// Schema returns the default Schema which should be used for Location fields
// where these are Required and Cannot be Changed
func Schema() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		ValidateFunc:     validate.NoEmptyStrings,
		StateFunc:        StateFunc,
		DiffSuppressFunc: DiffSuppressFunc,
	}
}

// SchemaOptional returns the Schema for a Location field where this can be optionally specified
func SchemaOptional() *schema.Schema {
	return &schema.Schema{
		Type:             schema.TypeString,
		Optional:         true,
		ForceNew:         true,
		StateFunc:        StateFunc,
		DiffSuppressFunc: DiffSuppressFunc,
	}
}

func SchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}
}

func DiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	return Normalize(old) == Normalize(new)
}

func HashCode(location interface{}) int {
	loc := location.(string)
	return hashcode.String(Normalize(loc))
}

func StateFunc(location interface{}) string {
	input := location.(string)
	return Normalize(input)
}
