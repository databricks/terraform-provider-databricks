package sqlanalytics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VisualizationEntity defines the parameters that can be set in the resource.
type VisualizationEntity struct {
	QueryID         string `json:"query_id"`
	VisualizationID string `json:"visualization_id,omitempty" tf:"computed"`

	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Options     string `json:"options"`
}

func (v *VisualizationEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Visualization, error) {
	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, v); err != nil {
		return nil, err
	}

	// Transform to API object.
	var av api.Visualization

	av.QueryID = v.QueryID
	av.Type = strings.ToUpper(v.Type)
	av.Name = v.Name
	av.Description = v.Description
	av.Options = json.RawMessage(v.Options)
	return &av, nil
}

func (v *VisualizationEntity) fromAPIObject(av *api.Visualization, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	v.QueryID = av.QueryID
	v.VisualizationID = av.ID.String()
	v.Type = strings.ToLower(av.Type)
	v.Name = av.Name
	v.Description = av.Description
	v.Options = string(av.Options)

	// Transform to ResourceData.
	return common.StructToData(*v, schema, data)
}

// NewVisualizationAPI ...
func NewVisualizationAPI(ctx context.Context, m interface{}) VisualizationAPI {
	return VisualizationAPI{m.(*common.DatabricksClient), ctx}
}

// VisualizationAPI ...
type VisualizationAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Create ...
func (a VisualizationAPI) Create(v *api.Visualization) error {
	return a.client.Post(a.context, "/preview/sql/visualizations", v, &v)
}

// Read ...
func (a VisualizationAPI) Read(queryID, visualizationID string) (*api.Visualization, error) {
	q, err := NewQueryAPI(a.context, a.client).Read(queryID)
	if err != nil {
		return nil, err
	}

	// Look for matching visualization ID.
	for _, vp := range q.Visualizations {
		var vnew api.Visualization
		err = json.Unmarshal(vp, &vnew)
		if err != nil {
			return nil, err
		}

		if vnew.ID.String() == visualizationID {
			// Include query ID in returned object.
			// It's not part of the API response.
			vnew.QueryID = queryID
			return &vnew, nil
		}
	}

	return nil, common.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: http.StatusNotFound,
		Message:    fmt.Sprintf("Cannot find visualization %s attached to query %s", visualizationID, queryID),
	}
}

// Update ...
func (a VisualizationAPI) Update(visualizationID string, v *api.Visualization) error {
	return a.client.Post(a.context, fmt.Sprintf("/preview/sql/visualizations/%s", visualizationID), &v, nil)
}

// Delete ...
func (a VisualizationAPI) Delete(visualizationID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/preview/sql/visualizations/%s", visualizationID), nil)
}

func jsonRemarshal(in []byte) ([]byte, error) {
	var v interface{}
	if len(in) == 0 {
		return in, nil
	}
	err := json.Unmarshal(in, &v)
	if err != nil {
		return nil, err
	}
	out, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceVisualization ...
func ResourceVisualization() *schema.Resource {
	p := common.NewPairSeparatedID("query_id", "visualization_id", "/")
	s := common.StructToSchema(
		VisualizationEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			// We care only about logical changes to the JSON payload in `options`.
			m["options"].DiffSuppressFunc = func(_, old, new string, d *schema.ResourceData) bool {
				oldp, err := jsonRemarshal([]byte(old))
				if err != nil {
					log.Printf("[WARN] Unable to remarshal value %#v", old)
					return false
				}
				newp, err := jsonRemarshal([]byte(new))
				if err != nil {
					log.Printf("[WARN] Unable to remarshal value %#v", new)
					return false
				}
				return bytes.Equal(oldp, newp)
			}

			// Changing part of the composite ID forces recreate.
			m["query_id"].ForceNew = true
			m["visualization_id"].ForceNew = true

			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			err = NewVisualizationAPI(ctx, c).Create(av)
			if err != nil {
				return err
			}

			// Convert API object back to resource data.
			// This includes setting the `visualization_id`, which is
			// needed to synthesize the composite resource identifier.
			err = v.fromAPIObject(av, s, data)
			if err != nil {
				return err
			}

			// Set composite resource identifier.
			p.Pack(data)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			queryID, visualizationID, err := p.Unpack(data)
			if err != nil {
				return err
			}

			av, err := NewVisualizationAPI(ctx, c).Read(queryID, visualizationID)
			if err != nil {
				return err
			}

			var v VisualizationEntity
			return v.fromAPIObject(av, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			_, visualizationID, err := p.Unpack(data)
			if err != nil {
				return err
			}

			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return NewVisualizationAPI(ctx, c).Update(visualizationID, av)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			_, visualizationID, err := p.Unpack(data)
			if err != nil {
				return err
			}
			return NewVisualizationAPI(ctx, c).Delete(visualizationID)
		},
		Schema: s,
	}.ToResource()
}
