package marshal

import (
	"testing"
)

func TestBasicMapDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				BasicMap: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestBasicMapValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				BasicMap: map[string]string{
					"key": "value",
				},
			},
			jsonString:   `{"basicmap":{"key":"value"},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestBasicMapForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				BasicMap:        nil,
				ForceSendFields: []string{"BasicMap"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
