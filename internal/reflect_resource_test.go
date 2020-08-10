package internal

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

type Address struct {
	Line      string `json:"line" tf:"group:v"`
	Lijn      string `json:"lijn" tf:"group:v"`
	IsPrimary bool   `json:"primary"`
}

type Dummy struct {
	Enabled     bool              `json:"enabled" tf:"conflicts:workers"`
	Workers     int               `json:"workers,omitempty"`
	Description string            `json:"description,omitempty"`
	Addresses   []Address         `json:"addresses,omitempty" tf:"max_items:10"`
	Unique      []Address         `json:"unique,omitempty" tf:"slice_set"`
	Things      []string          `json:"things,omitempty" tf:"slice_set"`
	Tags        map[string]string `json:"tags,omitempty" tf:"max_items:5"`
	Home        *Address          `json:"home,omitempty" tf:"group:v"`
	House       *Address          `json:"house,omitempty" tf:"group:v"`
}

func TestStructToData(t *testing.T) {
	s := StructToSchema(Dummy{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	assert.NotNil(t, s)
	assert.Equal(t, 5, s["tags"].MaxItems)
	assert.Equal(t, 10, s["addresses"].MaxItems)

	sp, err := SchemaPath(s, "addresses", "line")
	assert.NoError(t, err)
	assert.Equal(t, schema.TypeString, sp.Type)

	dummy := Dummy{
		Enabled:     false,
		Workers:     1004,
		Description: "something",
		Addresses: []Address{
			{
				Line:      "abc",
				IsPrimary: false,
			},
			{
				Line:      "def",
				IsPrimary: true,
			},
		},
		Unique: []Address{
			{
				Line:      "oop",
				IsPrimary: false,
			},
		},
		Things: []string{"one", "two", "two"},
		Tags: map[string]string{
			"Foo": "Bar",
		},
		Home: &Address{
			Line:      "bcd",
			IsPrimary: true,
		},
	}

	d := schema.TestResourceDataRaw(t, s, map[string]interface{}{})
	d.MarkNewResource()
	err = StructToData(dummy, s, d)
	assert.NoError(t, err)

	assert.Equal(t, "something", d.Get("description"))
	assert.Equal(t, false, d.Get("enabled"))
	assert.Equal(t, 2, d.Get("addresses.#"))

	var dummyCopy Dummy
	err = DataToStructPointer(d, s, &dummyCopy)
	assert.NoError(t, err)

	assert.Equal(t, len(dummyCopy.Addresses), len(dummy.Addresses))
	assert.Len(t, dummyCopy.Things, 2)
	assert.Len(t, dummy.Things, 3)

	err = d.Set("addresses", []interface{}{
		map[string]string{
			"line": "ABC",
			"lijn": "CBA",
		},
	})
	assert.NoError(t, err)

	err = DataToStructPointer(d, s, &dummyCopy)
	assert.EqualError(t, err, "addresses: validation conflicts: line and lijn,lijn and line")
}
