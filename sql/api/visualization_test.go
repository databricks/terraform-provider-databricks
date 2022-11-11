package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisualizationMarshalUnmarshal(t *testing.T) {
	v := Visualization{
		ID:          "12345",
		QueryID:     "queryID",
		Type:        "type",
		Name:        "name",
		Description: "description",
		Options:     json.RawMessage("{}"),
		QueryPlan:   json.RawMessage(`{"foo":"bar"}`),
	}

	out, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var vp Visualization
	if err := json.Unmarshal(out, &vp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, v, vp)
}
