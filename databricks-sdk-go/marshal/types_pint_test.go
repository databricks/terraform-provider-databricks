package marshal

import (
	"testing"
)

func TestPIntDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PInt: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPIntValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PInt: Ptr(0),
			},
			jsonString:   `{"pint":0,"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPIntForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PInt:            nil,
				ForceSendFields: []string{"PInt"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
