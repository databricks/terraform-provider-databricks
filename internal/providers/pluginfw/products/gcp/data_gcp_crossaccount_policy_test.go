package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeGcpCrossaccountPermissions_Base(t *testing.T) {
	perms := computeGcpCrossaccountPermissions()
	assert.Equal(t, gcpWorkspaceCreatorBasePermissions, perms)
	assert.Len(t, perms, 14)
}

func TestComputeGcpCrossaccountPermissions_DoesNotMutateBase(t *testing.T) {
	originalLen := len(gcpWorkspaceCreatorBasePermissions)
	computeGcpCrossaccountPermissions()
	assert.Len(t, gcpWorkspaceCreatorBasePermissions, originalLen)
}

func TestComputeGcpCrossaccountPermissions_NoComputePermissions(t *testing.T) {
	perms := computeGcpCrossaccountPermissions()
	for _, p := range perms {
		assert.NotContains(t, p, "compute.", "crossaccount policy must not contain compute permissions (use databricks_gcp_vpc_policy instead)")
	}
}

func TestComputeGcpCrossaccountPermissions_NoCmkPermissions(t *testing.T) {
	perms := computeGcpCrossaccountPermissions()
	assert.NotContains(t, perms, "cloudkms.cryptoKeys.getIamPolicy", "CMK permissions belong to the VPC project (use databricks_gcp_vpc_policy instead)")
	assert.NotContains(t, perms, "cloudkms.cryptoKeys.setIamPolicy")
}

func TestComputeGcpCrossaccountPermissions_ContainsIamRoleDelete(t *testing.T) {
	perms := computeGcpCrossaccountPermissions()
	assert.Contains(t, perms, "iam.roles.delete", "iam.roles.delete is workspace project only")
}
