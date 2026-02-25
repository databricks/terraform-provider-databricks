package gcp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcpUnityCatalogPermissions_ContainsExpectedPermissions(t *testing.T) {
	assert.Contains(t, gcpUnityCatalogPermissions, "pubsub.subscriptions.consume")
	assert.Contains(t, gcpUnityCatalogPermissions, "pubsub.topics.create")
	assert.Contains(t, gcpUnityCatalogPermissions, "storage.buckets.update")
}

func TestGcpUnityCatalogPermissions_Count(t *testing.T) {
	assert.Len(t, gcpUnityCatalogPermissions, 14)
}
