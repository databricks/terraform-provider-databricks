package catalog

import (
	"context"
	"net/url"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExternalLocationsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewExternalLocationsAPI(ctx context.Context, m any) ExternalLocationsAPI {
	return ExternalLocationsAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

type ExternalLocationInfo struct {
	Name           string `json:"name" tf:"force_new"`
	URL            string `json:"url"`
	CredentialName string `json:"credential_name"`
	Comment        string `json:"comment,omitempty"`
	SkipValidation bool   `json:"skip_validation,omitempty"`
	Owner          string `json:"owner,omitempty" tf:"computed"`
	MetastoreID    string `json:"metastore_id,omitempty" tf:"computed"`
	ReadOnly       bool   `json:"read_only,omitempty"`
}

func (a ExternalLocationsAPI) create(el *ExternalLocationInfo) error {
	return a.client.Post(a.context, "/unity-catalog/external-locations", el, &el)
}

func (a ExternalLocationsAPI) get(name string) (el ExternalLocationInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/external-locations/"+url.PathEscape(name), nil, &el)
	return
}

func (a ExternalLocationsAPI) delete(name string, force bool) error {
	return a.client.Delete(a.context, "/unity-catalog/external-locations/"+url.PathEscape(name), map[string]any{
		"force": force,
	})
}

func ResourceExternalLocation() *schema.Resource {
	s := common.StructToSchema(ExternalLocationInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			m["force_destroy"] = &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			}
			m["skip_validation"].DiffSuppressFunc = func(k, old, new string, d *schema.ResourceData) bool {
				return old == "false" && new == "true"
			}
			m["url"].DiffSuppressFunc = ucDirectoryPathSlashOnlySuppressDiff
			return m
		})
	update := updateFunctionFactory("/unity-catalog/external-locations", []string{"owner", "comment", "url", "credential_name"})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var el ExternalLocationInfo
			common.DataToStructPointer(d, s, &el)
			el.Owner = ""
			err := NewExternalLocationsAPI(ctx, c).create(&el)
			if err != nil {
				return err
			}
			d.SetId(el.Name)
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			el, err := NewExternalLocationsAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(el, s, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			force := d.Get("force_destroy").(bool)
			return NewExternalLocationsAPI(ctx, c).delete(d.Id(), force)
		},
	}.ToResource()
}
