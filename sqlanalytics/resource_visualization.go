package sqlanalytics

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Visualization ...
type Visualization struct {
	ID          string `json:"id,omitempty" tf:"computed"`
	QueryID     string `json:"query_id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Options     string `json:"options"`
}

type visualizationResource struct {
	schema map[string]*schema.Schema
}

func (r *visualizationResource) toAPIObject(d *schema.ResourceData) (*api.Visualization, error) {
	var v Visualization

	// Transform from ResourceData.
	if err := common.DataToStructPointer(d, r.schema, &v); err != nil {
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

func (r *visualizationResource) fromAPIObject(av *api.Visualization, d *schema.ResourceData) error {
	var v Visualization

	// Transform from API object.
	v.ID = strconv.Itoa(av.ID)
	v.QueryID = av.QueryID
	v.Type = strings.ToLower(av.Type)
	v.Name = av.Name
	v.Description = av.Description
	v.Options = string(av.Options)

	// Transform to ResourceData.
	if err := common.StructToData(v, r.schema, d); err != nil {
		return err
	}

	return nil
}

func (r *visualizationResource) create(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	av, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	avNew, err := w.CreateVisualization(av)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(avNew, d)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(avNew.ID))
	return nil
}

func (r *visualizationResource) read(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	av, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	avNew, err := w.ReadVisualization(av)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(avNew, d)
	if err != nil {
		return err
	}

	return nil
}

func (r *visualizationResource) update(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	av, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	avNew, err := w.UpdateVisualization(av)
	if err != nil {
		return err
	}

	err = r.fromAPIObject(avNew, d)
	if err != nil {
		return err
	}

	return nil
}

func (r *visualizationResource) delete(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
	av, err := r.toAPIObject(d)
	if err != nil {
		return err
	}

	var w = api.NewWrapper(ctx, c)
	err = w.DeleteVisualization(av)
	if err != nil {
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
	r := visualizationResource{
		common.StructToSchema(
			Visualization{},
			func(m map[string]*schema.Schema) map[string]*schema.Schema {
				// We care only about logical changes to the JSON payload in `options`.
				m["options"].DiffSuppressFunc = func(_, old, new string, d *schema.ResourceData) bool {
					oldp, err := jsonRemarshal([]byte(old))
					if err != nil {
						panic(err)
					}
					newp, err := jsonRemarshal([]byte(new))
					if err != nil {
						panic(err)
					}
					return bytes.Compare(oldp, newp) == 0
				}
				return m
			}),
	}
	return common.Resource{
		Schema: r.schema,
		Create: r.create,
		Read:   r.read,
		Update: r.update,
		Delete: r.delete,
	}.ToResource()
}
