package marshal

import (
	"testing"
)

func TestPChildFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: nil,
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildFSValueEmpty(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{},
			},
			jsonString:   `{"pchildfs":{},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{
					Int: 3,
				},
			},
			jsonString:   `{"pchildfs":{"childint":3},"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS:        nil,
				ForceSendFields: []string{"PChildFS"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestPChildFSChildForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				PChildFS: &structFS{
					ForceSendFields: []string{"Int"},
				},
				ForceSendFields: []string{"Int"},
			},
			jsonString:              `{"childfs":{},"pchildfs":{"childint":0},"childnofs":{}, "int":0}`,
			matchClassic:            false,
			unmarshalForceSendField: []string{"Int"},
		},
	)
}
