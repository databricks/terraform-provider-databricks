package sqlanalytics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VisualizationEntity defines the parameters that can be set in the resource.
type VisualizationEntity struct {
	QueryID     string `json:"query_id"`
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

	if data.Id() != "" {
		id, err := strconv.Atoi(data.Id())
		if err != nil {
			return nil, err
		}
		av.ID = id
	}

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
	v.Type = strings.ToLower(av.Type)
	v.Name = av.Name
	v.Description = av.Description
	v.Options = string(av.Options)

	// Transform to ResourceData.
	if err := common.StructToData(*v, schema, data); err != nil {
		return err
	}

	return nil
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

func (a VisualizationAPI) buildPath(path ...int) string {
	out := "/preview/sql/visualizations"
	if len(path) == 1 {
		out = fmt.Sprintf("%s/%d", out, path[0])
	}
	return out
}

// Create ...
func (a VisualizationAPI) Create(v *api.Visualization) (*api.Visualization, error) {
	var vp api.Visualization
	err := a.client.Post(a.context, a.buildPath(), v, &vp)
	if err != nil {
		return nil, err
	}

	// Set query ID on returned object.
	// It's not included in the POST response.
	vp.QueryID = v.QueryID

	return &vp, err
}

// Read ...
func (a VisualizationAPI) Read(v *api.Visualization) (*api.Visualization, error) {
	if v.QueryID == "" {
		return nil, fmt.Errorf("Cannot read visualization without query ID")
	}

	q, err := NewQueryAPI(a.context, a.client).Read(&api.Query{ID: v.QueryID})
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

		if vnew.ID == v.ID {
			// Include query ID in returned object.
			// It's not part of the API response.
			vnew.QueryID = v.QueryID
			return &vnew, nil
		}
	}

	return nil, fmt.Errorf("Cannot find visualization %d attached to query %s", v.ID, v.QueryID)
}

// Update ...
func (a VisualizationAPI) Update(v *api.Visualization) (*api.Visualization, error) {
	var vp api.Visualization
	err := a.client.Post(a.context, a.buildPath(v.ID), v, &vp)
	if err != nil {
		return nil, err
	}

	// Set query ID on returned object.
	// It's not included in the POST response.
	vp.QueryID = v.QueryID

	return &vp, nil
}

// Delete ...
func (a VisualizationAPI) Delete(v *api.Visualization) error {
	return a.client.Delete(a.context, a.buildPath(v.ID), nil)
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
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			avNew, err := NewVisualizationAPI(ctx, c).Create(av)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(strconv.Itoa(avNew.ID))
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			avNew, err := NewVisualizationAPI(ctx, c).Read(av)
			if err != nil {
				return err
			}

			return v.fromAPIObject(avNew, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			_, err = NewVisualizationAPI(ctx, c).Update(av)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			return nil
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var v VisualizationEntity
			av, err := v.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return NewVisualizationAPI(ctx, c).Delete(av)
		},
		Schema: s,
	}.ToResource()
}
