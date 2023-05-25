package storage

import (
	"context"
	"fmt"
	"github.com/databricks/terraform-provider-databricks/common"
)

// This describes an enum
type FileType string

type FileAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

type FileInfo struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Owner   string `json:"owner,omitempty"`
}

type CreateFileRequestContent struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Owner   string `json:"owner,omitempty"`
}

type UpdateFileRequestContent struct {
	Name    string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Owner   string `json:"owner,omitempty"`
}

type ListFileResponseContent struct {
	Volumes *FileInfo `json:"files,omitempty"`
}

func NewFileAPI(ctx context.Context, m any) FileAPI {
	return FileAPI{m.(*common.DatabricksClient), context.WithValue(ctx, common.Api, common.API_2_1)}
}

func (a FileAPI) get(name string) (v FileInfo, err error) {
	err = a.client.Get(a.context, "/files"+url.PathEscape(name), nil, &v)
	return
}

func (a FileAPI) put(v *FileInfo) error {
	return a.client.Put(a.context, "/files", v, &v)
}

func (a FileAPI) create(v *FileInfo) error {
	return a.client.Post(a.context, "/files", v, &v)
}

func (a FileAPI) delete(name string) error {
	return a.client.Delete(a.context, "/files"+url.PathEscape(name), nil)
}

func ResourceFiles() *schema.Resource {
	s := common.StructToSchema(FileInfo{},
		func(m map[string]*schema.Schema) map[string]*schema.Schema {
			return m
		})
	update := updateFunctionFactory("/files", []string{"owner", "name", "comment"})
	return common.Resource{
	s := common.StructToSchema(FileInfo{},	Schema: s,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			var f FileInfo
			common.DataToStructPointer(d, s, &f)
			err := NewFileAPI(ctx, c).create(&f)
			if err != nil {
				return err
			}
			d.SetId(f.Name)
			return update(ctx, d, c)
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			f, err := NewFileAPI(ctx, c).get(d.Id())
			if err != nil {
				return err
			}
			return common.StructToData(f, s, d)
		},
		Update: update,
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			return NewFileAPI(ctx, c).delete(d.Id())
		},
	}.ToResource()
}
