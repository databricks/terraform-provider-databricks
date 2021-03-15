package sqlanalytics

import (
	"context"
	"fmt"
	"strconv"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WidgetEntity defines the parameters that can be set in the resource.
type WidgetEntity struct {
	DashboardID     string `json:"dashboard_id"`
	Text            string `json:"text,omitempty"`
	VisualizationID string `json:"visualization_id,omitempty"`

	Position  *WidgetPosition   `json:"position,omitempty"`
	Parameter []WidgetParameter `json:"parameter,omitempty"`
}

// WidgetPosition ...
type WidgetPosition struct {
	SizeX int `json:"size_x"`
	SizeY int `json:"size_y"`
	PosX  int `json:"pos_x"`
	PosY  int `json:"pos_y"`
}

// WidgetParameter ...
type WidgetParameter struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	MapTo string `json:"map_to,omitempty"`
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
}

func (w *WidgetEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Widget, error) {
	var aw api.Widget

	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, w); err != nil {
		return nil, err
	}

	// Only copy over the ID if this is an existing resource.
	if data.Id() != "" {
		id, err := strconv.Atoi(data.Id())
		if err != nil {
			return nil, err
		}
		aw.ID = id
	}

	aw.DashboardID = w.DashboardID

	// The visualization ID is a string for the Terraform resource and an integer in the API.
	if w.VisualizationID != "" {
		visualizationID, err := strconv.Atoi(w.VisualizationID)
		if err != nil {
			return nil, err
		}
		aw.VisualizationID = &visualizationID
	}

	if w.Text != "" {
		aw.Text = &w.Text
	}

	if w.Position != nil {
		aw.Options.Position = &api.WidgetPosition{
			AutoHeight: false,
			SizeX:      w.Position.SizeX,
			SizeY:      w.Position.SizeY,
			PosX:       w.Position.PosX,
			PosY:       w.Position.PosY,
		}
	}

	aw.Options.ParameterMapping = make(map[string]api.WidgetParameterMapping)
	for _, wp := range w.Parameter {
		aw.Options.ParameterMapping[wp.Name] = api.WidgetParameterMapping{
			Name:  wp.Name,
			Type:  wp.Type,
			MapTo: wp.MapTo,
			Title: wp.Title,
			Value: wp.Value,
		}
	}

	return &aw, nil
}

func (w *WidgetEntity) fromAPIObject(aw *api.Widget, schema map[string]*schema.Schema, data *schema.ResourceData) error {
	// Copy from API object.
	w.DashboardID = aw.DashboardID

	if aw.VisualizationID != nil {
		w.VisualizationID = fmt.Sprint(*aw.VisualizationID)
	}

	if aw.Text != nil {
		w.Text = *aw.Text
	}

	if pos := aw.Options.Position; pos != nil {
		w.Position = &WidgetPosition{
			SizeX: pos.SizeX,
			SizeY: pos.SizeY,
			PosX:  pos.PosX,
			PosY:  pos.PosY,
		}
	}

	// Pass to ResourceData.
	if err := common.StructToData(*w, schema, data); err != nil {
		return err
	}

	return nil
}

// ResourceWidget ...
func ResourceWidget() *schema.Resource {
	s := common.StructToSchema(
		WidgetEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["text"].ConflictsWith = []string{"visualization_id"}
			return m
		})

	return common.Resource{
		Create: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var w WidgetEntity
			aw, err := w.toAPIObject(s, data)
			if err != nil {
				return err
			}

			awp, err := api.NewWrapper(ctx, c).CreateWidget(aw)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			data.SetId(fmt.Sprint(awp.ID))
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var w WidgetEntity
			aw, err := w.toAPIObject(s, data)
			if err != nil {
				return err
			}

			awNew, err := api.NewWrapper(ctx, c).ReadWidget(aw)
			if err != nil {
				return err
			}

			return w.fromAPIObject(awNew, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var d WidgetEntity
			ad, err := d.toAPIObject(s, data)
			if err != nil {
				return err
			}

			_, err = api.NewWrapper(ctx, c).UpdateWidget(ad)
			if err != nil {
				return err
			}

			// No need to set anything because the resource is going to be
			// read immediately after being created.
			return nil
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			var w WidgetEntity
			aw, err := w.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return api.NewWrapper(ctx, c).DeleteWidget(aw)
		},
		Schema: s,
	}.ToResource()
}
