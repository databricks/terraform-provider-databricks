package sharing

import (
	"context"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

type ProvidersAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

func NewProvidersAPI(ctx context.Context, m interface{}) ProvidersAPI {
	return ProvidersAPI{m.(*common.DatabricksClient), ctx}
}

type ProviderInfo struct {
	common.Namespace
	Name                string `json:"name" tf:"force_new"`
	Comment             string `json:"comment,omitempty"`
	AuthenticationType  string `json:"authentication_type"`
	RecipientProfileStr string `json:"recipient_profile_str" tf:"sensitive"`
}

type Providers struct {
	Providers []ProviderInfo `json:"providers"`
}

func (a ProvidersAPI) createProvider(ci *ProviderInfo) error {
	return a.client.Post(a.context, "/unity-catalog/providers", ci, ci)
}

func (a ProvidersAPI) getProvider(name string) (ci ProviderInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/providers/"+name, nil, &ci)
	return
}

func (a ProvidersAPI) deleteProvider(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/providers/"+name, nil)
}

func (a ProvidersAPI) updateProvider(ci *ProviderInfo) error {
	patch := struct {
		Comment string `json:"comment"`
	}{
		Comment: ci.Comment,
	}
	return a.client.Patch(a.context, "/unity-catalog/providers/"+ci.Name, patch)
}

func ResourceProvider() common.Resource {
	providerSchema := common.StructToSchema(ProviderInfo{}, func(m map[string]*schema.Schema) map[string]*schema.Schema {
		m["authentication_type"].ValidateFunc = validation.StringInSlice([]string{"TOKEN"}, false)
		common.NamespaceCustomizeSchemaMap(m)
		return m
	})

	providerSchemaForRead := map[string]*schema.Schema{
		"name":                providerSchema["name"],
		"comment":             providerSchema["comment"],
		"authentication_type": providerSchema["authentication_type"],
	}
	return common.Resource{
		Schema: providerSchema,
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff) error {
			return common.NamespaceCustomizeDiff(d)
		},
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ri ProviderInfo
			common.DataToStructPointer(d, providerSchema, &ri)
			if err := NewProvidersAPI(ctx, c).createProvider(&ri); err != nil {
				return err
			}
			d.SetId(ri.Name)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			ri, err := NewProvidersAPI(ctx, c).getProvider(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(ri, providerSchemaForRead, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var ri ProviderInfo
			common.DataToStructPointer(d, providerSchema, &ri)
			return NewProvidersAPI(ctx, c).updateProvider(&ri)
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewProvidersAPI(ctx, c).deleteProvider(d.Id())
		},
	}
}
