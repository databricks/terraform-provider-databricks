package common

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getField(v interface{}, fieldName string) interface{} {
	return reflect.ValueOf(v).Elem().FieldByName(fieldName).Interface()
}

func setField(v interface{}, fieldName string, value interface{}) {
	reflect.ValueOf(v).Elem().FieldByName(fieldName).Set(reflect.ValueOf(value))
}

func getNewStructOfType(v interface{}) interface{} {
	return reflect.New(reflect.TypeOf(v).Elem()).Interface()
}

func UpdateWithOwnerInSeparateRequest(
	ctx context.Context,
	d *schema.ResourceData,
	updateMethod func(ctx context.Context, updateStruct interface{}) (*interface{}, error),
	updateStruct interface{}) (interface{}, error) {

	if d.HasChange("owner") {
		ownerUpdateStruct := getNewStructOfType(updateStruct)
		setField(ownerUpdateStruct, "Name", getField(updateStruct, "Name"))
		setField(ownerUpdateStruct, "Owner", getField(updateStruct, "Owner"))
		_, err := updateMethod(ctx, ownerUpdateStruct)
		if err != nil {
			return nil, err
		}
	}

	setField(updateStruct, "Owner", "")
	response, err := updateMethod(ctx, updateStruct)

	if err != nil {
		if d.HasChange("owner") {
			// Rollback
			oldOwner, newOwner := d.GetChange("owner")
			ownerRollbackStruct := reflect.New(reflect.TypeOf(updateStruct).Elem()).Interface()
			setField(ownerRollbackStruct, "Name", getField(updateStruct, "Name"))
			setField(ownerRollbackStruct, "Owner", oldOwner.(string))
			_, rollbackErr := updateMethod(ctx, ownerRollbackStruct)
			if rollbackErr != nil {
				log.Printf("[WARN] Owner of this resource was updated but other fields couldn't be updated and owner couldn't be rolled back. \n As a result, the owner of this resource is updated to %s but other attributes aren't. To revert the owner change, please manually change the owner to %s. \n\n You can also use the databricks cli (https://docs.databricks.com/en/dev-tools/cli/install.html) to update the owner.", newOwner.(string), oldOwner.(string))
				return nil, fmt.Errorf("%w. Owner rollback also failed: %w", err, rollbackErr)
			}
		}
		return nil, err
	}
	return response, nil
}
