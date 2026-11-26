package marshal

import (
	"testing"
)

func TestPBoolDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PBool: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPBoolValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PBool: Ptr(false),
			},
			jsonString:   `{"pbool":false,"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPBoolForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PBool:           nil,
				ForceSendFields: []string{"PBool"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
