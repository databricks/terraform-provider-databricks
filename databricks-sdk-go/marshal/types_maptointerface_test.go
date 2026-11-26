package marshal

import (
	"testing"
)

func TestMapToInterfaceDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestMapToInterfaceValue(t *testing.T) {
	element := map[string]interface{}{
		"key": "3",
	}
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt: element,
			},
			jsonString:   `{"maptoint":{"key":"3"},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestMapToInterfaceForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				MapToInt:        nil,
				ForceSendFields: []string{"MapToInt"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
