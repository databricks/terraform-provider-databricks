package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeGcpVpcPermissions_Base(t *testing.T) {
	perms := computeGcpVpcPermissions(false, false, false)
	assert.Equal(t, gcpVpcBasePermissions, perms)
	assert.Len(t, perms, 12)
	assert.NotContains(t, perms, "compute.forwardingRules.get")
	assert.NotContains(t, perms, "compute.subnetworks.get")
	assert.NotContains(t, perms, "cloudkms.cryptoKeys.getIamPolicy")
}

func TestComputeGcpVpcPermissions_WithByovpc(t *testing.T) {
	perms := computeGcpVpcPermissions(true, false, false)
	assert.Len(t, perms, len(gcpVpcBasePermissions)+len(gcpByovpcPermissions))
	assert.Contains(t, perms, "compute.subnetworks.get")
	assert.Contains(t, perms, "compute.subnetworks.getIamPolicy")
	assert.Contains(t, perms, "compute.subnetworks.setIamPolicy")
}

func TestComputeGcpVpcPermissions_WithCmk(t *testing.T) {
	perms := computeGcpVpcPermissions(false, true, false)
	assert.Len(t, perms, len(gcpVpcBasePermissions)+len(gcpCmkPermissions))
	assert.Contains(t, perms, "cloudkms.cryptoKeys.getIamPolicy")
	assert.Contains(t, perms, "cloudkms.cryptoKeys.setIamPolicy")
}

func TestComputeGcpVpcPermissions_WithPsc(t *testing.T) {
	perms := computeGcpVpcPermissions(false, false, true)
	assert.Len(t, perms, len(gcpVpcBasePermissions)+len(gcpPscPermissions))
	assert.Contains(t, perms, "compute.forwardingRules.get")
	assert.Contains(t, perms, "compute.forwardingRules.list")
}

func TestComputeGcpVpcPermissions_AllFlags(t *testing.T) {
	perms := computeGcpVpcPermissions(true, true, true)
	expected := len(gcpVpcBasePermissions) + len(gcpByovpcPermissions) + len(gcpCmkPermissions) + len(gcpPscPermissions)
	assert.Len(t, perms, expected)
	assert.Contains(t, perms, "compute.subnetworks.get")
	assert.Contains(t, perms, "cloudkms.cryptoKeys.getIamPolicy")
	assert.Contains(t, perms, "compute.forwardingRules.get")
}

func TestComputeGcpVpcPermissions_DoesNotMutateBase(t *testing.T) {
	originalLen := len(gcpVpcBasePermissions)
	computeGcpVpcPermissions(true, true, true)
	assert.Len(t, gcpVpcBasePermissions, originalLen)
}

func TestComputeGcpVpcPermissions_ContainsComputePermissions(t *testing.T) {
	perms := computeGcpVpcPermissions(false, false, false)
	assert.Contains(t, perms, "compute.networks.get")
	assert.Contains(t, perms, "compute.firewalls.create")
}

func TestComputeGcpVpcPermissions_NoIamRolesDelete(t *testing.T) {
	perms := computeGcpVpcPermissions(false, false, false)
	assert.NotContains(t, perms, "iam.roles.delete", "iam.roles.delete is workspace project only")
}
