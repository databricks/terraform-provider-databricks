package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashboardMarshalUnmarshal(t *testing.T) {
	d := Dashboard{
		ID:   "id",
		Name: "name",
		Tags: []string{"tag1", "tag2"},
	}

	out, err := json.Marshal(d)
	if err != nil {
		t.Fatal(err)
	}

	var dp Dashboard
	if err := json.Unmarshal(out, &dp); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, d, dp)
}
