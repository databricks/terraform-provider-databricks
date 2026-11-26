package mlflow

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

// mlflowModelAttribute is the part of an attribute's schema that Terraform plans with.
type mlflowModelAttribute struct {
	Type     schema.ValueType
	Required bool
	Optional bool
	Computed bool
	ForceNew bool
	// MaxItems tells a block that may appear once (1) from a repeatable one (0).
	MaxItems int
}

// flattenMlflowModelSchema keys every attribute by its dotted path and descends into
// nested blocks, so an attribute that is added, removed or turned into another kind of
// collection changes the result.
func flattenMlflowModelSchema(prefix string, s map[string]*schema.Schema, flat map[string]mlflowModelAttribute) {
	for name, attribute := range s {
		flat[prefix+name] = mlflowModelAttribute{
			Type:     attribute.Type,
			Required: attribute.Required,
			Optional: attribute.Optional,
			Computed: attribute.Computed,
			ForceNew: attribute.ForceNew,
			MaxItems: attribute.MaxItems,
		}
		if block, ok := attribute.Elem.(*schema.Resource); ok {
			flattenMlflowModelSchema(prefix+name+".", block.Schema, flat)
		}
	}
}

// TestResourceMlflowModelSchema pins the schema that common.StructToSchema derives from
// ml.CreateModelRequest and common.Namespace. The request struct is the only source of
// name, description and tags, so replacing it, for example with another SDK's request
// type, could otherwise change an attribute or its flags unnoticed.
func TestResourceMlflowModelSchema(t *testing.T) {
	mlflowModelOnBothPaths(t, testResourceMlflowModelSchema)
}

func testResourceMlflowModelSchema(t *testing.T) {
	flat := map[string]mlflowModelAttribute{}
	flattenMlflowModelSchema("", ResourceMlflowModel().ToResource().Schema, flat)

	assert.Equal(t, map[string]mlflowModelAttribute{
		"name":                         {Type: schema.TypeString, Required: true, ForceNew: true},
		"description":                  {Type: schema.TypeString, Optional: true},
		"tags":                         {Type: schema.TypeList, Optional: true},
		"tags.key":                     {Type: schema.TypeString, Optional: true},
		"tags.value":                   {Type: schema.TypeString, Optional: true},
		"registered_model_id":          {Type: schema.TypeString, Computed: true},
		"provider_config":              {Type: schema.TypeList, Optional: true, Computed: true, MaxItems: 1},
		"provider_config.workspace_id": {Type: schema.TypeString, Optional: true, Computed: true},
	}, flat)
}
