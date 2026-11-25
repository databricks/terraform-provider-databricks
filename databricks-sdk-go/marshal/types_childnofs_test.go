package marshal

import (
	"testing"
)

func TestChildNoFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildNoFS: structNoFS{},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestChildNoFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildNoFS: structNoFS{
					Int: 3,
				},
			},
			jsonString:   `{"childnofs":{"childint":3},"childfs":{}}`,
			matchClassic: true,
		},
	)
}

func TestChildNoFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildNoFS:       structNoFS{},
				ForceSendFields: []string{"ChildNoFS"},
			},
			jsonString:   `{"childnofs":{},"childfs":{}}`,
			matchClassic: true,
		},
	)
}
