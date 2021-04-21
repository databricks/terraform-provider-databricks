package sqlanalytics

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/sqlanalytics/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WidgetEntity defines the parameters that can be set in the resource.
type WidgetEntity struct {
	DashboardID string `json:"dashboard_id"`
	WidgetID    string `json:"widget_id,omitempty" tf:"computed"`

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

type sortWidgetParameter []WidgetParameter

func (a sortWidgetParameter) Len() int {
	return len(a)
}

func (a sortWidgetParameter) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sortWidgetParameter) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func (w *WidgetEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Widget, error) {
	var aw api.Widget

	// Extract from ResourceData.
	if err := common.DataToStructPointer(data, schema, w); err != nil {
		return nil, err
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
	w.WidgetID = strconv.Itoa(aw.ID)

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

	w.Parameter = make([]WidgetParameter, 0, len(aw.Options.ParameterMapping))
	for _, p := range aw.Options.ParameterMapping {
		wp := WidgetParameter{
			Name:  p.Name,
			Type:  p.Type,
			MapTo: p.MapTo,
			Title: p.Title,
		}

		// Re-marshal value so we can try to unmarshal different types.
		// We don't know about the type it holds, because it depends
		// on the parameter's type, which we don't have access to.
		b, err := json.Marshal(p.Value)
		if err != nil {
			return err
		}

		// Try unmarshalling `string`.
		{
			var v string
			err := json.Unmarshal(b, &v)
			if err == nil {
				wp.Value = v
				w.Parameter = append(w.Parameter, wp)
				continue
			}
		}

		// Try unmarshalling `[]string`.
		{
			var vs []string
			err := json.Unmarshal(b, &vs)
			if err == nil {
				wp.Values = vs
				w.Parameter = append(w.Parameter, wp)
				continue
			}
		}

		return fmt.Errorf("Unable to derive type from message: %v", string(b))
	}

	// Sort parameters by their name for deterministic order.
	sort.Sort(sortWidgetParameter(w.Parameter))

	// Pass to ResourceData.
	return common.StructToData(*w, schema, data)
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

// Create ...
func (a WidgetAPI) Create(w *api.Widget) error {
	return a.client.Post(a.context, "/preview/sql/widgets", w, &w)
}

// Read ...
func (a WidgetAPI) Read(dashboardID, widgetID string) (*api.Widget, error) {
	d, err := NewDashboardAPI(a.context, a.client).Read(&api.Dashboard{ID: dashboardID})
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

		if strconv.Itoa(wnew.ID) == widgetID {
			// Include dashboard ID in returned object.
			// It's not part of the API response.
			wnew.DashboardID = dashboardID
			return &wnew, nil
		}
	}

	return nil, fmt.Errorf("Cannot find widget %s attached to dashboard %s", widgetID, dashboardID)
}

// Update ...
func (a WidgetAPI) Update(widgetID string, w *api.Widget) error {
	return a.client.Post(a.context, fmt.Sprintf("/preview/sql/widgets/%s", widgetID), w, nil)
}

// Delete ...
func (a WidgetAPI) Delete(widgetID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/preview/sql/widgets/%s", widgetID), nil)
}

// ResourceWidget ...
func ResourceWidget() *schema.Resource {
	p := common.NewPairSeparatedID("dashboard_id", "widget_id", "/")
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

			err = NewWidgetAPI(ctx, c).Create(aw)
			if err != nil {
				return err
			}

			// Convert API object back to resource data.
			// This includes setting the `widget_id`, which is
			// needed to synthesize the composite resource identifier.
			err = w.fromAPIObject(aw, s, data)
			if err != nil {
				return err
			}

			// Set composite resource identifier.
			p.Pack(data)
			return nil
		},
		Read: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			dashboardID, widgetID, err := p.Unpack(data)
			if err != nil {
				return err
			}

			aw, err := NewWidgetAPI(ctx, c).Read(dashboardID, widgetID)
			if err != nil {
				return err
			}

			var w WidgetEntity
			return w.fromAPIObject(aw, s, data)
		},
		Update: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			_, widgetID, err := p.Unpack(data)
			if err != nil {
				return err
			}

			var w WidgetEntity
			aw, err := w.toAPIObject(s, data)
			if err != nil {
				return err
			}

			return NewWidgetAPI(ctx, c).Update(widgetID, aw)
		},
		Delete: func(ctx context.Context, data *schema.ResourceData, c *common.DatabricksClient) error {
			_, widgetID, err := p.Unpack(data)
			if err != nil {
				return err
			}
			return NewWidgetAPI(ctx, c).Delete(widgetID)
		},
		Schema: s,
	}.ToResource()
}
