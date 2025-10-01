package sql

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/sql/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// WidgetEntity defines the parameters that can be set in the resource.
type WidgetEntity struct {
	DashboardID string `json:"dashboard_id" tf:"force_new"`
	WidgetID    string `json:"widget_id,omitempty" tf:"computed,force_new"`

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	Text            string `json:"text,omitempty"`
	VisualizationID string `json:"visualization_id,omitempty" tf:"force_new"`

	Position  *WidgetPosition   `json:"position,omitempty"`
	Parameter []WidgetParameter `json:"parameter,omitempty" tf:"slice_set"`

	common.ProviderConfig
}

// WidgetPosition ...
type WidgetPosition struct {
	SizeX      int  `json:"size_x"`
	SizeY      int  `json:"size_y"`
	PosX       int  `json:"pos_x" tf:"optional,default:0"`
	PosY       int  `json:"pos_y" tf:"optional,default:0"`
	AutoHeight bool `json:"auto_height,omitempty"`
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

// Use second part of ID if it's a composite ID, or the verbatim value.
// This allows setting the visualization ID as the backend visualization ID
// or as the visualization's resource composite ID. Both will work.
func extractVisualizationID(id string) string {
	parts := strings.SplitN(id, "/", 2)
	return parts[len(parts)-1]
}

func (w *WidgetEntity) toAPIObject(schema map[string]*schema.Schema, data *schema.ResourceData) (*api.Widget, error) {
	var aw api.Widget

	// Extract from ResourceData.
	common.DataToStructPointer(data, schema, w)

	aw.DashboardID = w.DashboardID
	aw.Options.Title = w.Title
	aw.Options.Description = w.Description

	// The visualization ID is a string for the Terraform resource and an integer in the API.
	if w.VisualizationID != "" {
		visualizationID := api.NewStringOrInt(extractVisualizationID(w.VisualizationID))
		aw.VisualizationID = &visualizationID
	}

	if w.Text != "" {
		aw.Text = &w.Text
	}

	if w.Position != nil {
		aw.Options.Position = &api.WidgetPosition{
			AutoHeight: w.Position.AutoHeight,
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
	w.WidgetID = aw.ID.String()
	w.Title = aw.Options.Title
	w.Description = aw.Options.Description

	if aw.VisualizationID != nil {
		w.VisualizationID = fmt.Sprint(*aw.VisualizationID)
	}

	if aw.Text != nil {
		w.Text = *aw.Text
	}

	if pos := aw.Options.Position; pos != nil {
		w.Position = &WidgetPosition{
			AutoHeight: pos.AutoHeight,
			SizeX:      pos.SizeX,
			SizeY:      pos.SizeY,
			PosX:       pos.PosX,
			PosY:       pos.PosY,
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

		return fmt.Errorf("unable to derive type from message: %v", string(b))
	}

	// Sort parameters by their name for deterministic order.
	sort.Sort(sortWidgetParameter(w.Parameter))

	// Pass to ResourceData.
	return common.StructToData(*w, schema, data)
}

// NewWidgetAPI ...
func NewWidgetAPI(ctx context.Context, m any) WidgetAPI {
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
	d, err := NewDashboardAPI(a.context, a.client).Read(dashboardID)
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

		if wnew.ID.String() == widgetID {
			// Include dashboard ID in returned object.
			// It's not part of the API response.
			wnew.DashboardID = dashboardID
			return &wnew, nil
		}
	}

	return nil, &apierr.APIError{
		ErrorCode:  "NOT_FOUND",
		StatusCode: 404,
		Message:    fmt.Sprintf("Cannot find widget %s attached to dashboard %s", widgetID, dashboardID),
	}
}

// Update ...
func (a WidgetAPI) Update(widgetID string, w *api.Widget) error {
	return a.client.Post(a.context, fmt.Sprintf("/preview/sql/widgets/%s", widgetID), w, nil)
}

// Delete ...
func (a WidgetAPI) Delete(widgetID string) error {
	return a.client.Delete(a.context, fmt.Sprintf("/preview/sql/widgets/%s", widgetID), nil)
}

func ResourceSqlWidget() common.Resource {
	p := common.NewPairSeparatedID("dashboard_id", "widget_id", "/")
	s := common.StructToSchema(
		WidgetEntity{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["text"].ConflictsWith = []string{"visualization_id"}

			// Ignore the query ID part in composite visualization ID.
			// It is present in this field if users refer to a visualization by the native
			// Terraform resource ID (e.g. `databricks_sql_visualization.name.id`)
			m["visualization_id"].DiffSuppressFunc = func(_, old, new string, d *schema.ResourceData) bool {
				return extractVisualizationID(old) == extractVisualizationID(new)
			}

			// Add provider_config customizations
			common.CustomizeSchemaPath(m, "provider_config").SetOptional()
			common.CustomizeSchemaPath(m, "provider_config", "workspace_id").SetRequired()

			return m
		})

	return common.Resource{
		DeprecationMessage: "This resource is deprecated and will be removed in future. Please switch to databricks_dashboard to author new AI/BI dashboards using the latest tooling.",
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
	}
}
