package marshal

import (
	"testing"
)

func TestMapToArrayDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToArray: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestMapToArrayValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToArray: map[string][]interface{}{
					"key": {structNoFS{
						Int: 3,
					}},
				},
			},
			jsonString:   `{"maptoarray":{"key":[{"childint":3}]},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
			unmarshalResult: &customStruct{
				MapToArray: map[string][]interface{}{
					"key": {map[string]interface{}{
						"childint": 3.0,
					}},
				},
			},
		},
	)
}

func TestMapToArrayForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToArray:      nil,
				ForceSendFields: []string{"MapToArray"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
