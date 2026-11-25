package marshal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type jsonTags struct {
	NoJsonTag    *string
	EmptyJsonTag *string `json:""`
	OnlyOmit     *string `json:",omitempty"`
	JsonIgnore   *string `json:"-"`
}

func (s *jsonTags) UnmarshalJSON(b []byte) error {
	return Unmarshal(b, s)
}

func (s jsonTags) MarshalJSON() ([]byte, error) {
	return Marshal(s)
}

func TestEmpty(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st:           customStruct{},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestNoJsonFields(t *testing.T) {
	testString := "value"
	st := jsonTags{
		NoJsonTag:    &testString,
		EmptyJsonTag: &testString,
		JsonIgnore:   &testString,
		OnlyOmit:     &testString,
	}
	jsonString, err := json.Marshal(st)
	type C jsonTags
	original, _ := json.Marshal((C)(st))
	assert.NoError(t, err)
	compareJSON(t, string(jsonString), string(original))
	var res jsonTags
	err = json.Unmarshal(jsonString, &res)
	expected := jsonTags{
		NoJsonTag:    &testString,
		EmptyJsonTag: &testString,
		OnlyOmit:     &testString,
	}
	assert.NoError(t, err)
	assert.Equal(t, res, expected)
}

func TestDefaults(t *testing.T) {
	executeBasicMarshalTest(t, basicMarshalTest{
		st: customStruct{
			Int:   0,
			Float: 0.0,
			Str:   "",
			Bool:  false,

			PInt:   nil,
			PFloat: nil,
			PStr:   nil,
			PBool:  nil,

			Slice:      nil,
			Interface:  nil,
			BasicMap:   nil,
			MapToArray: nil,

			ChildFS:   structFS{},
			ChildNoFS: structNoFS{},

			PChildFS:   nil,
			PChildNoFS: nil,

			MapToCustomTypeFS: nil,
		},
		jsonString:   `{"childfs":{},"childnofs":{}}`,
		matchClassic: true,
	})
}

func TestForce(t *testing.T) {
	executeBasicMarshalTest(t, basicMarshalTest{
		st: customStruct{
			Int:   0,
			Float: 0.0,
			Str:   "",
			Bool:  false,

			PInt:   nil,
			PFloat: nil,
			PStr:   nil,
			PBool:  nil,

			Slice:      nil,
			Interface:  nil,
			BasicMap:   nil,
			MapToArray: nil,

			ChildFS:   structFS{},
			ChildNoFS: structNoFS{},

			PChildFS:   nil,
			PChildNoFS: nil,

			MapToCustomTypeFS: nil,

			ForceSendFields: []string{"Int", "Float", "Str", "Bool"},
		},
		jsonString:              `{"int":0, "float":0, "string":"", "bool":false, "childfs":{},"childnofs":{}}`,
		matchClassic:            false,
		unmarshalForceSendField: []string{"Int", "Float", "Str", "Bool"},
	})
}

func TestNotDefault(t *testing.T) {
	st := customStruct{
		Int:   1,
		Float: 1.0,
		Str:   "1",
		Bool:  true,

		PInt:   Ptr(0),
		PFloat: Ptr(0.0),
		PStr:   Ptr(""),
		PBool:  Ptr(false),

		Slice: []int{9},
		BasicMap: map[string]string{
			"key": "value",
		},
		MapToArray: map[string][]interface{}{
			"key": {2.0, "field"},
		},

		ChildFS: structFS{
			Int: 1,
		},
		ChildNoFS: structNoFS{
			Int: 2,
		},

		PChildFS: &structFS{
			Int:             3,
			ForceSendFields: []string{"Int"},
		},
		PChildNoFS: &structNoFS{
			Int: 4,
		},

		MapToCustomTypeFS: map[string]structFS{
			"struct1": {
				Int:             1,
				ForceSendFields: []string{"Int"},
			},
		},
	}

	rc := executeBasicMarshalTest(t, basicMarshalTest{
		st: st,
		jsonString: `{"int":1, "float":1, "string":"1", "bool":true,
		"pint":0, "pfloat":0,"pbool":false,"pstring":"",
		"slice":[9], "childfs":{"childint":1},"childnofs":{"childint":2},
		"basicmap":{"key":"value"},"maptoarray":{"key":[2,"field"]},
		"pchildfs":{"childint":3},"pchildnofs":{"childint":4},
		"maptocustomtypefs":{"struct1":{"childint":1}}}`,
		matchClassic:            true,
		unmarshalForceSendField: []string{"Int", "Bool", "Str", "Float"},
	})
	rc.ForceSendFields = nil
	assert.Equal(t, rc, st)

}

type SimpleStruct struct {
	ID string `json:"id"`
}

func TestIntToString(t *testing.T) {
	jsonString := `{"id":1}`
	var res SimpleStruct
	err := Unmarshal([]byte(jsonString), &res)
	assert.NoError(t, err)
	assert.Equal(t, "1", res.ID)
}

func TestNoOmit(t *testing.T) {
	in := SimpleStruct{
		ID: "",
	}
	jsonString := `{"id":""}`
	res, err := Marshal(in)
	assert.NoError(t, err)
	assert.Equal(t, jsonString, string(res))
}

func TestWrongField(t *testing.T) {
	st := customStruct{
		ForceSendFields: []string{"notAField"},
	}
	_, err := json.Marshal(st)
	assert.ErrorContains(t, err, "field notAField cannot be found in struct customStruct")
}
