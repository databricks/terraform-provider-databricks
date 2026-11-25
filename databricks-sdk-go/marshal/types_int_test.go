package marshal

import (
	"testing"
)

func TestIntDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Int: 0,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestIntValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Int: 2,
			},
			jsonString:              `{"int":2,"childfs":{},"childnofs":{}}`,
			matchClassic:            true,
			unmarshalForceSendField: []string{"Int"},
		},
	)
}

func TestIntForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Int:             0,
				ForceSendFields: []string{"Int"},
			},
			jsonString:              `{"int":0,"childfs":{},"childnofs":{}}`,
			matchClassic:            false,
			unmarshalForceSendField: []string{"Int"},
		},
	)
}
