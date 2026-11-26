package marshal

import (
	"testing"
)

func TestPChildNoFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildNoFS: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildNoFSValueEmpty(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildNoFS: &structNoFS{},
			},
			jsonString:   `{"pchildnofs":{},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildNoFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildNoFS: &structNoFS{
					Int: 3,
				},
			},
			jsonString:   `{"pchildnofs":{"childint":3},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildNoFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildNoFS:      nil,
				ForceSendFields: []string{"PChildNoFS"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}
