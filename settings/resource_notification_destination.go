package settings

import (
	"context"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NDStruct struct {
	settings.NotificationDestination
}

func (NDStruct) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	// Required fields

	return s
}

var ndSchema = common.StructToSchema(NDStruct{}, nil)

func ResourceNotificationDestination() common.Resource {
	return common.Resource{
		Schema: ndSchema,
		Create: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var newNDrequest settings.CreateNotificationDestinationRequest
			common.DataToStructPointer(d, ndSchema, &newNDrequest)
			createdND, err := w.NotificationDestinations.Create(ctx, newNDrequest)
			if err != nil {
				return err
			}
			d.SetId(createdND.Id)
			return nil
		},
		Read: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			readND, err := w.NotificationDestinations.Get(ctx, settings.GetNotificationDestinationRequest{
				Id: d.Id(),
			})
			if err != nil {
				return err
			}
			return common.StructToData(readND, ndSchema, d)
		},
		Update: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			var updateNDRequest settings.UpdateNotificationDestinationRequest
			common.DataToStructPointer(d, ndSchema, &updateNDRequest)
			_, err = w.NotificationDestinations.Update(ctx, updateNDRequest)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, d *schema.ResourceData, c *common.DatabricksClient) error {
			w, err := c.WorkspaceClient()
			if err != nil {
				return err
			}
			return w.NotificationDestinations.Delete(ctx, settings.DeleteNotificationDestinationRequest{
				Id: d.Id(),
			})
		},
	}
}
