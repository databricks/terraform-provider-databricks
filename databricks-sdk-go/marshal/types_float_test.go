package marshal

import (
	"testing"
)

func TestFloatDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float: 0.0,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestFloatValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float: 2.3,
			},
			jsonString:              `{"float":2.3,"childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			unmarshalForceSendField: []string{"Float"},
		},
	)
}

func TestFloatForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Float:           0,
				ForceSendFields: []string{"Float"},
			},
			jsonString:              `{"float":0,"childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			unmarshalForceSendField: []string{"Float"},
		},
	)
}
