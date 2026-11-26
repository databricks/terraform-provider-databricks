package marshal

import (
	"testing"
)

func TestPStrDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PStr: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPStrValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PStr: Ptr(""),
			},
			jsonString:   `{"pstring":"","childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPStrForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PStr:            nil,
				ForceSendFields: []string{"PStr"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
