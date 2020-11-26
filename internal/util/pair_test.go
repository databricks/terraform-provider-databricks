package util

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPairIDResource(t *testing.T) {
	type bindResourceFixture struct {
		create, read, delete bool
		left, right          string
		id                   string
		assertID             string
		err                  error
		assertError          string
	}
	tests := []bindResourceFixture{
		{
			read:        true,
			id:          "a",
			assertError: "Invalid ID: a",
		},
		{
			read:        true,
			id:          "a|",
			assertError: "right_id cannot be empty",
		},
		{
			read:        true,
			id:          "|b",
			assertError: "left_id cannot be empty",
		},
		{
			delete:      true,
			id:          "a",
			assertError: "Invalid ID: a",
		},
		{
			read:     true,
			id:       "a|b",
			left:     "a",
			right:    "b",
			assertID: "a|b",
		},
		{
			read:     true,
			id:       "a|b|c|d",
			left:     "a",
			right:    "b|c|d",
			assertID: "a|b|c|d",
		},
		{
			delete:   true,
			id:       "a|b|c|d",
			left:     "a",
			right:    "b|c|d",
			assertID: "a|b|c|d",
		},
		{
			read:  true,
			id:    "a|b",
			err:   common.NotFound("Nope"),
			left:  "a",
			right: "b",
		},
		{
			read:        true,
			id:          "a|b",
			err:         fmt.Errorf("Nope"),
			left:        "a",
			right:       "b",
			assertID:    "a|b",
			assertError: "Nope",
		},
		{
			create:      true,
			left:        "a",
			assertError: "right_id cannot be empty",
		},
		{
			create:      true,
			right:       "a",
			assertError: "left_id cannot be empty",
		},
		{
			create:   true,
			left:     "a",
			right:    "b",
			assertID: "a|b",
		},
		{
			create:      true,
			left:        "a",
			right:       "b",
			err:         fmt.Errorf("Nope"),
			assertError: "Nope",
			// ID is not set on error for create
			assertID: "",
		},
		{
			delete:      true,
			id:          "a|b",
			assertID:    "a|b",
			left:        "a",
			right:       "b",
			err:         fmt.Errorf("Nope"),
			assertError: "Nope",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt), func(t *testing.T) {
			var state map[string]interface{}
			if tt.create {
				state = map[string]interface{}{
					"left_id":  tt.left,
					"right_id": tt.right,
				}
			}
			d, err := qa.ResourceFixture{
				Resource: NewPairID("left_id", "right_id").BindResource(BindResource{
					ReadContext: func(ctx context.Context, left, right string, c *common.DatabricksClient) error {
						return tt.err
					},
					CreateContext: func(ctx context.Context, left, right string, c *common.DatabricksClient) error {
						return tt.err
					},
					DeleteContext: func(ctx context.Context, left, right string, c *common.DatabricksClient) error {
						return tt.err
					},
				}),
				Create: tt.create,
				Read:   tt.read,
				Delete: tt.delete,
				ID:     tt.id,
				State:  state,
			}.Apply(t)
			if tt.assertError != "" {
				require.NotNilf(t, err, "Expected to have %s error", tt.assertError)
				require.True(t, strings.HasPrefix(err.Error(), tt.assertError), err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.assertID, d.Id(), "ID does not match")
			assert.Equal(t, tt.left, d.Get("left_id"), "Invalid left")
			assert.Equal(t, tt.right, d.Get("right_id"), "Invalid right")
		})
	}
}
