package sqlanalytics

import (
	"context"
	"encoding/json"
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

	// Mutually exclusive.
	Value  string   `json:"value,omitempty"`
	Values []string `json:"values,omitempty"`
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

	if len(w.Parameter) > 0 {
		aw.Options.ParameterMapping = make(map[string]api.WidgetParameterMapping)
		for _, wp := range w.Parameter {
			wpm := api.WidgetParameterMapping{
				Name:  wp.Name,
				Type:  wp.Type,
				MapTo: wp.MapTo,
				Title: wp.Title,
			}

			if len(wp.Values) > 0 {
				wpm.Value = wp.Values
			} else {
				wpm.Value = wp.Value
			}

			aw.Options.ParameterMapping[wp.Name] = wpm
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

// NewWidgetAPI ...
func NewWidgetAPI(ctx context.Context, m interface{}) WidgetAPI {
	return WidgetAPI{m.(*common.DatabricksClient), ctx}
}

// WidgetAPI ...
type WidgetAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func (a WidgetAPI) buildPath(path ...int) string {
	out := "/preview/sql/widgets"
	if len(path) == 1 {
		out = fmt.Sprintf("%s/%d", out, path[0])
	}
	return out
}

// Create ...
func (a WidgetAPI) Create(w *api.Widget) (*api.Widget, error) {
	var wout api.Widget
	err := a.client.Post(a.context, a.buildPath(), w, &wout)
	if err != nil {
		return nil, err
	}

	return &wout, err
}

// Read ...
func (a WidgetAPI) Read(w *api.Widget) (*api.Widget, error) {
	if w.DashboardID == "" {
		return nil, fmt.Errorf("Cannot read widget without dashboard ID")
	}

	d, err := NewDashboardAPI(a.context, a.client).Read(&api.Dashboard{ID: w.DashboardID})
	if err != nil {
		return nil, err
	}

	// Look for matching widget ID.
	for _, wp := range d.Widgets {
		var wnew api.Widget
		err = json.Unmarshal(wp, &wnew)
		if err != nil {
			return nil, err
		}

		if wnew.ID == w.ID {
			// Include dashboard ID in returned object.
			// It's not part of the API response.
			wnew.DashboardID = w.DashboardID
			return &wnew, nil
		}
	}

	return nil, fmt.Errorf("Cannot find widget %d attached to dashboard %s", w.ID, w.DashboardID)
}

// Update ...
func (a WidgetAPI) Update(w *api.Widget) (*api.Widget, error) {
	var wout api.Widget
	err := a.client.Post(a.context, a.buildPath(w.ID), w, &wout)
	if err != nil {
		return nil, err
	}

	return &wout, nil
}

// Delete ...
func (a WidgetAPI) Delete(w *api.Widget) error {
	return a.client.Delete(a.context, a.buildPath(w.ID), nil)
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

			awp, err := NewWidgetAPI(ctx, c).Create(aw)
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

			awNew, err := NewWidgetAPI(ctx, c).Read(aw)
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

			_, err = NewWidgetAPI(ctx, c).Update(ad)
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

			return NewWidgetAPI(ctx, c).Delete(aw)
		},
		Schema: s,
	}.ToResource()
}
