package sqlanalytics

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// VisualizationEntity defines the parameters that can be set in the resource.
type VisualizationEntity struct {
	ID          string `json:"id,omitempty" tf:"computed"`
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

	if v.ID != "" {
		id, err := strconv.Atoi(v.ID)
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
	v.ID = strconv.Itoa(av.ID)
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

			avNew, err := api.NewWrapper(ctx, c).CreateVisualization(av)
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

			avNew, err := api.NewWrapper(ctx, c).ReadVisualization(av)
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

			_, err = api.NewWrapper(ctx, c).UpdateVisualization(av)
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

			return api.NewWrapper(ctx, c).DeleteVisualization(av)
		},
		Schema: s,
	}.ToResource()
}
