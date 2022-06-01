package catalog

import (
	"context"
	"net/url"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ExternalLocationsAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewExternalLocationsAPI(ctx context.Context, m interface{}) ExternalLocationsAPI {
	return ExternalLocationsAPI{m.(*common.DatabricksClient), ctx}
}

type ExternalLocationInfo struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	CredentialName string `json:"credential_name"`
	Comment        string `json:"comment,omitempty"`
	SkipValidation bool   `json:"skip_validation,omitempty"`
	Owner          string `json:"owner,omitempty" tf:"computed"`
	MetastoreID    string `json:"metastore_id,omitempty" tf:"computed"`
}

func (a ExternalLocationsAPI) create(el *ExternalLocationInfo) error {
	return a.client.Post(a.context, "/unity-catalog/external-locations", el, &el)
}

func (a ExternalLocationsAPI) get(name string) (el ExternalLocationInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/external-locations/"+url.PathEscape(name), nil, &el)
	return
}

func (a ExternalLocationsAPI) update(name string, update map[string]interface{}) error {
	return a.client.Patch(a.context, "/unity-catalog/external-locations/"+url.PathEscape(name), update)
}

func (a ExternalLocationsAPI) delete(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/external-locations/"+url.PathEscape(name), nil)
}

func ResourceExternalLocation() *schema.Resource {
	s := common.StructToSchema(ExternalLocationInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	update := updateFunctionFactory("external-location", []string{"owner", "name", "comment", "url", "credential_name", "skip_validation"})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var el ExternalLocationInfo
			common.DataToStructPointer(d, s, &el)
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
			return NewExternalLocationsAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
