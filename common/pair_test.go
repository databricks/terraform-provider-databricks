package common

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		schema               func(m map[string]*schema.Schema) map[string]*schema.Schema
	}
	tests := []bindResourceFixture{
		{
			read:        true,
			id:          "a",
			assertError: "invalid ID: a",
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
			assertError: "invalid ID: a",
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
			id:       "a|123",
			left:     "a",
			right:    "123",
			assertID: "a|123",
			schema: func(m map[string]*schema.Schema) map[string]*schema.Schema {
				m["right_id"].Type = schema.TypeInt
				return m
			},
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
			read: true,
			id:   "a|b",
			err: &apierr.APIError{
				ErrorCode:  "NOT_FOUND",
				StatusCode: 404,
				Message:    "nope",
			},
			left:     "a",
			right:    "b",
			assertID: "",
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
			var state map[string]string
			if tt.create {
				state = map[string]string{
					"left_id":  tt.left,
					"right_id": tt.right,
				}
			}
			p := NewPairID("left_id", "right_id")
			if tt.schema != nil {
				p.Schema(tt.schema)
			}
			resource := p.BindResource(BindResource{
				ReadContext: func(ctx context.Context, left, right string, c *DatabricksClient) error {
					return tt.err
				},
				CreateContext: func(ctx context.Context, left, right string, c *DatabricksClient) error {
					return tt.err
				},
				DeleteContext: func(ctx context.Context, left, right string, c *DatabricksClient) error {
					return tt.err
				},
			}).ToResource()
			ctx := context.Background()
			d := resource.Data(&terraform.InstanceState{
				Attributes: state,
				ID:         tt.id,
			})
			var err error
			var diags diag.Diagnostics
			client := &DatabricksClient{}
			switch {
			case tt.create:
				diags = resource.CreateContext(ctx, d, client)
			case tt.read:
				diags = resource.ReadContext(ctx, d, client)
			case tt.delete:
				diags = resource.DeleteContext(ctx, d, client)
			}
			if diags != nil {
				err = errors.New(diags[0].Summary)
			}
			if tt.assertError != "" {
				require.NotNilf(t, err, "Expected to have %s error", tt.assertError)
				require.True(t, strings.HasPrefix(err.Error(), tt.assertError), err)
			} else {
				require.NoError(t, err)
			}
			assert.Equal(t, tt.assertID, d.Id(), "ID does not match")
			assert.Equal(t, tt.left, d.Get("left_id"), "Invalid left")
			assert.Equal(t, tt.right, fmt.Sprintf("%v", d.Get("right_id")), "Invalid right")
		})
	}
}
