package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWidgetMarshalUnmarshal(t *testing.T) {
	i678 := 678
	sAbc := "text"

	w := Widget{
		ID:          12345,
		DashboardID: "dashboardID",

		VisualizationID: &i678,
		Text:            &sAbc,

		Options: WidgetOptions{
			ParameterMapping: map[string]WidgetParameterMapping{
				"p1": {
					Name:  "p1",
					Type:  "text",
					MapTo: "mapTo",
					Title: "title",
				},
			},
			Position: &WidgetPosition{
				AutoHeight: false,
				SizeX:      1,
				SizeY:      2,
				PosX:       3,
				PosY:       4,
			},
		},
	}

	out, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}

	var wp Widget
	if err := json.Unmarshal(out, &wp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, w, wp)
}

func TestWidgetUnmarshalWithVisualization(t *testing.T) {
	w := Widget{
		Visualization: json.RawMessage(`
			{
				"id": 12345
			}
		`),
	}

	out, err := json.Marshal(w)
	if err != nil {
		t.Fatal(err)
	}

	var wp Widget
	if err := json.Unmarshal(out, &wp); err != nil {
		t.Fatal(err)
	}

	if assert.NotNil(t, wp.VisualizationID) {
		assert.Equal(t, 12345, *wp.VisualizationID)
	}
}
