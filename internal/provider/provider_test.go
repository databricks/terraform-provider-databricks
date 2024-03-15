package provider

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestProviderSchema_SuppressDiffAddedForAllNeededFields(t *testing.T) {
	p := provider.DatabricksProvider()
	for _, v := range p.ResourcesMap {
		traverseResource(v, checkResource(t))
	}
}

// For every resource:
//
//	For every field:
//	  if the field is TypeList or TypeSet and the Elem is a schema.Resource:
//	    if all fields of the schema.Resource are primitive and optional:
//	      suppress_diff must be annotated.
func checkResource(t *testing.T) func(r *schema.Resource) {
	return func(r *schema.Resource) {
		for k, v := range r.Schema {
			if v.Type == schema.TypeList || v.Type == schema.TypeSet {
				if nested, ok := v.Elem.(*schema.Resource); ok && nested != nil {
					needsSuppressDiff := true
					for _, vv := range nested.Schema {
						if vv.Type == schema.TypeList || vv.Type == schema.TypeSet || !vv.Optional {
							needsSuppressDiff = false
							break
						}
					}
					if needsSuppressDiff {
						assert.NotNil(t, v.DiffSuppressFunc, "DiffSuppressFunc must be annotated for %s", k)
					}
				}
			}
		}
	}
}

func traverseResource(r *schema.Resource, f func(*schema.Resource)) {
	f(r)

	for _, v := range r.Schema {
		if v.Type == schema.TypeList || v.Type == schema.TypeSet {
			if nested, ok := v.Elem.(*schema.Resource); ok && nested != nil {
				traverseResource(v.Elem.(*schema.Resource), f)
			}
		}
	}
}
