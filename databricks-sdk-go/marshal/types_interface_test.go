package marshal

import (
	"testing"
)

func TestInterfaceDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestInterfaceValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface: structNoFS{
					Int: 3,
				},
			},
			jsonString:   `{"interface":{"childint":3},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
			unmarshalResult: &customStruct{ //Reflect can't determine the original type, so unmarshal returns this instead.
				Interface: map[string]interface{}{
					"childint": 3.0,
				},
			},
		},
	)
}

func TestInterfaceForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Interface:       nil,
				ForceSendFields: []string{"Interface"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
