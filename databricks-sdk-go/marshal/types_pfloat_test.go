package marshal

import (
	"testing"
)

func TestPFloatDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPFloatValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat: Ptr(0.0),
			},
			jsonString:   `{"pfloat":0,"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPFloatForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PFloat:          nil,
				ForceSendFields: []string{"PFloat"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
