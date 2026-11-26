package marshal

import (
	"testing"
)

func TestSliceDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Slice: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestSliceValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Slice: []int{3, 4},
			},
			jsonString:   `{"slice":[3,4],"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestSliceEmptyValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Slice: []int{},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
			unmarshalResult: &customStruct{
				Slice: nil,
			},
		},
	)
}

func TestSliceForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				Slice:           nil,
				ForceSendFields: []string{"Slice"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
