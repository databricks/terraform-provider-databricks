package exporter

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestNoServicesSkipsRun(t *testing.T) {
	assert.EqualError(t, (&importContext{}).Run(), "no services to import")
}

func TestMatchesName(t *testing.T) {
	assert.False(t, (&importContext{match: "x"}).MatchesName("y"))
}

func TestImportContextFindSkips(t *testing.T) {
	assert.Nil(t, (&importContext{
		State: stateApproximation{
			Resources: []resourceApproximation{
				{
					Type: "a",
					Instances: []instanceApproximation{
						{
							Attributes: map[string]interface{}{
								"b": nil,
							},
						},
					},
				},
			},
		},
	}).Find(&resource{
		Resource:  "a",
		Attribute: "b",
		Name:      "c",
	}, "x"))
}

func TestImportContextHas(t *testing.T) {
	assert.True(t, (&importContext{
		State: stateApproximation{
			Resources: []resourceApproximation{
				{
					Type: "a",
					Instances: []instanceApproximation{
						{
							Attributes: map[string]interface{}{
								"b": "d",
							},
						},
					},
				},
			},
		},
	}).Has(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	}))
}

func TestEmitNaResource(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoImportable(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoSearchAvail(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
			},
		},
		services: "e",
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoSearchFails(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					return fmt.Errorf("just fails")
				},
			},
		},
		services: "e",
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoSearchNoId(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					return nil
				},
			},
		},
		services: "e",
	}).Emit(&resource{
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}

func TestEmitNoSearchSucceedsImportFails(t *testing.T) {
	(&importContext{
		importing: map[string]bool{},
		Resources: map[string]*schema.Resource{
			"a": {},
		},
		Importables: map[string]importable{
			"a": {
				Service: "e",
				Search: func(ic *importContext, r *resource) error {
					r.ID = "some"
					return nil
				},
				Import: func(ic *importContext, r *resource) error {
					return fmt.Errorf("fails")
				},
			},
		},
		services: "e",
	}).Emit(&resource{
		Data:      &schema.ResourceData{},
		Resource:  "a",
		Attribute: "b",
		Value:     "d",
		Name:      "c",
	})
}
