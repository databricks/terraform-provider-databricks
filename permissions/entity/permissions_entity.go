package entity

import (
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/databricks/terraform-provider-databricks/common"
)

// PermissionsEntity is the one used for resource metadata
type PermissionsEntity struct {
	ObjectType        string                     `json:"object_type,omitempty" tf:"computed"`
	AccessControlList []iam.AccessControlRequest `json:"access_control" tf:"slice_set"`
	common.ProviderConfig
}

func (p PermissionsEntity) ContainsUserOrServicePrincipal(name string) bool {
	for _, ac := range p.AccessControlList {
		if ac.UserName == name || ac.ServicePrincipalName == name {
			return true
		}
	}
	return false
}
