package common

import "fmt"

func OwnerRollbackError(oldOwner string, newOwner string) string {
	return fmt.Sprintf("[WARN] Owner of this resource was updated but other fields couldn't be updated and owner couldn't be rolled back. \n As a result, the owner of this resource is updated to %s but other attributes aren't. To revert the owner change, please manually change the owner to %s. \n\n You can also use the databricks cli (https://docs.databricks.com/en/dev-tools/cli/install.html) to update the owner.", newOwner, oldOwner)
}
