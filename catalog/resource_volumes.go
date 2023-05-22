package catalog

import (
	"context"
	"net/url"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This describes an enum
type VolumeType string

type VolumesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

type CreateVolumeRequestContent struct {
	Name            string     `json:"name"`
	CatalogName     string     `json:"catalog_name"`
	VolumeType      VolumeType `json:"volume_type"`
	SchemaName      string     `json:"schema_name"`
	StorageLocation string     `json:"storage_location,omitempty"`
	Comment         string     `json:"comment,omitempty"`
}

type UpdateVolumeRequestContent struct {
	Comment string `json:"comment,omitempty"`
	Name    string `json:"name,omitempty"`
	Owner   string `json:"owner,omitempty"`
}

type ListVolumesResponseContent struct {
	Volumes *VolumeInfo `json:"volumes,omitempty"`
}

func NewVolumesAPI(ctx context.Context, m any) VolumesAPI {
	return VolumesAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

// TODO: check all for tf:"computed", for eg: it is probably required in created at/by, updated at/by
type VolumeInfo struct {
	Owner           string     `json:"owner,omitempty"`
	StorageLocation string     `json:"storage_location,omitempty"`
	VolumeId        string     `json:"volume_id,omitempty"`
	CatalogName     string     `json:"catalog_name,omitempty"`
	SchemaName      string     `json:"schema_name,omitempty"`
	VolumeType      VolumeType `json:"volume_type,omitempty"`
	MetaStoreId     string     `json:"metastore_id,omitempty"`
	Name            string     `json:"name,omitempty"`
	FullName        string     `json:"full_name,omitempty"`
	CreatedBy       string     `json:"created_by,omitempty"`
	CreatedAt       float64    `json:"created_at,omitempty"`
	UpdatedBy       string     `json:"updated_by,omitempty"`
	UpdatedAt       float64    `json:"updated_at,omitempty"`
	Comment         string     `json:"comment,omitempty"`
}

func (a VolumesAPI) create(v *VolumeInfo) error {
	return a.client.Post(a.context, "/unity-catalog/volumes", v, &v)
}

func (a VolumesAPI) get(name string) (v VolumeInfo, err error) {
	err = a.client.Get(a.context, "/unity-catalog/volumes/"+url.PathEscape(name), nil, &v)
	return
}

func (a VolumesAPI) delete(name string) error {
	return a.client.Delete(a.context, "/unity-catalog/volumes/"+url.PathEscape(name), nil)
}

func ResourceVolumes() *schema.Resource {
	s := common.StructToSchema(VolumeInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	update := updateFunctionFactory("/unity-catalog/volumes", []string{"owner", "name", "comment"})
	return common.Resource{
		Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var v VolumeInfo
			common.DataToStructPointer(d, s, &v)
			v.Owner = ""
			err := NewVolumesAPI(ctx, c).create(&v)
			if err != nil {
				return err
			}
			d.SetId(v.Name)
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			v, err := NewVolumesAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(v, s, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewVolumesAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
