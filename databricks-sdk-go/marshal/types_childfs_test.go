package marshal

import (
	"testing"
)

func TestChildFSDefault(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildFS: structFS{},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestChildFSValue(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildFS: structFS{
					Int: 3,
				},
			},
			jsonString:   `{"childfs":{"childint":3},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestChildFSForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildFS:         structFS{},
				ForceSendFields: []string{"ChildFS"},
			},
			jsonString:   `{"childfs":{},"childnofs":{}}`,
			matchClassic: true,
		},
	)
}

func TestChildFSChildForce(t *testing.T) {
	executeBasicMarshalTest(t,
		basicMarshalTest{
			st: customStruct{
				ChildFS: structFS{
					ForceSendFields: []string{"Int"},
				},
				ForceSendFields: []string{"Int"},
			},
			jsonString:              `{"childfs":{"childint":0},"childnofs":{}, "int":0}`,
			matchClassic:            false,
			unmarshalForceSendField: []string{"Int"},
		},
	)
}
