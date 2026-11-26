package marshal

import (
	"testing"
)

func TestStringDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str: "",
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestStringValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str: "val",
			},
			jsonString:              `{"string":"val","childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			unmarshalForceSendField: []string{"Str"},
		},
	)
}

func TestStringForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Str:             "",
				ForceSendFields: []string{"Str"},
			},
			jsonString:              `{"string":"","childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			unmarshalForceSendField: []string{"Str"},
		},
	)
}
