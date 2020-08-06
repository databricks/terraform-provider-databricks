package internal

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

type Address struct {
	Line      string `json:"line"`
	IsPrimary bool   `json:"primary"`
}

type Dummy struct {
	Enabled     bool              `json:"enabled"`
	Workers     int               `json:"workers,omitempty"`
	Description string            `json:"description,omitempty"`
	Addresses   []Address         `json:"addresses,omitempty" tf:"max_items:10"`
	Unique      []Address         `json:"unique,omitempty" tf:"slice_set"`
	Things      []string          `json:"things,omitempty" tf:"slice_set"`
	Tags        map[string]string `json:"tags,omitempty"`
	Home        *Address          `json:"home,omitempty"`
}

func TestStructToData(t *testing.T) {
	s := StructToSchema(Dummy{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		return s
	})
	assert.NotNil(t, s)

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
	DataToStructPointer(d, s, &dummyCopy)

	assert.Equal(t, len(dummyCopy.Addresses), len(dummy.Addresses))
	assert.Len(t, dummyCopy.Things, 2)
	assert.Len(t, dummy.Things, 3)

}
