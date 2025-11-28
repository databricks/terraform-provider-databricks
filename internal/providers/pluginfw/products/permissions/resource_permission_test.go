package permissions

import (
	"context"
	"testing"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPermissionResource_Schema(t *testing.T) {
	r := ResourcePermission()
	require.NotNil(t, r)

	var resp resource.SchemaResponse
	r.Schema(context.Background(), resource.SchemaRequest{}, &resp)

	// Verify schema has required fields
	require.NotNil(t, resp.Schema.Attributes)

	// Check principal fields
	_, ok := resp.Schema.Attributes["user_name"]
	assert.True(t, ok, "user_name should be in schema")

	_, ok = resp.Schema.Attributes["group_name"]
	assert.True(t, ok, "group_name should be in schema")

	_, ok = resp.Schema.Attributes["service_principal_name"]
	assert.True(t, ok, "service_principal_name should be in schema")

	// Check permission level
	_, ok = resp.Schema.Attributes["permission_level"]
	assert.True(t, ok, "permission_level should be in schema")

	// Check object identifier fields
	_, ok = resp.Schema.Attributes["object_type"]
	assert.True(t, ok, "object_type should be in schema")

	_, ok = resp.Schema.Attributes["object_id"]
	assert.True(t, ok, "object_id should be in schema")

	// Check computed fields
	idAttr, ok := resp.Schema.Attributes["id"]
	assert.True(t, ok, "id should be in schema")
	if stringAttr, ok := idAttr.(schema.StringAttribute); ok {
		assert.True(t, stringAttr.Computed, "id should be computed")
	}

	// Verify object_type is required
	objectTypeAttr, ok := resp.Schema.Attributes["object_type"]
	assert.True(t, ok, "object_type should be in schema")
	if stringAttr, ok := objectTypeAttr.(schema.StringAttribute); ok {
		assert.True(t, stringAttr.Required, "object_type should be required")
	}

	// Verify object_id is required
	objectIDAttr, ok := resp.Schema.Attributes["object_id"]
	assert.True(t, ok, "object_id should be in schema")
	if stringAttr, ok := objectIDAttr.(schema.StringAttribute); ok {
		assert.True(t, stringAttr.Required, "object_id should be required")
	}
}

func TestPermissionResource_Metadata(t *testing.T) {
	r := ResourcePermission()
	var resp resource.MetadataResponse
	r.Metadata(context.Background(), resource.MetadataRequest{
		ProviderTypeName: "databricks",
	}, &resp)

	assert.Equal(t, "databricks_permission", resp.TypeName)
}

func TestPermissionResource_ParseID(t *testing.T) {
	r := &PermissionResource{}

	tests := []struct {
		name           string
		id             string
		wantObjectType string
		wantObjectID   string
		wantPrincipal  string
		wantError      bool
	}{
		{
			name:           "cluster permission",
			id:             "clusters/test-cluster-id/user@example.com",
			wantObjectType: "clusters",
			wantObjectID:   "test-cluster-id",
			wantPrincipal:  "user@example.com",
			wantError:      false,
		},
		{
			name:           "job permission",
			id:             "jobs/123/test-group",
			wantObjectType: "jobs",
			wantObjectID:   "123",
			wantPrincipal:  "test-group",
			wantError:      false,
		},
		{
			name:           "authorization tokens",
			id:             "authorization/tokens/developers",
			wantObjectType: "authorization",
			wantObjectID:   "tokens",
			wantPrincipal:  "developers",
			wantError:      false,
		},
		{
			name:           "notebook with path containing slashes",
			id:             "notebooks/Shared/Team/notebook.py/user@example.com",
			wantObjectType: "notebooks",
			wantObjectID:   "Shared/Team/notebook.py",
			wantPrincipal:  "user@example.com",
			wantError:      false,
		},
		{
			name:      "invalid format - too few parts",
			id:        "clusters/test-cluster-id",
			wantError: true,
		},
		{
			name:      "invalid format - no slashes",
			id:        "test-id",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotObjectType, gotObjectID, gotPrincipal, err := r.parseID(tt.id)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantObjectType, gotObjectType)
				assert.Equal(t, tt.wantObjectID, gotObjectID)
				assert.Equal(t, tt.wantPrincipal, gotPrincipal)
			}
		})
	}
}

func TestPermissionResource_Configure(t *testing.T) {
	r := &PermissionResource{}
	client := &common.DatabricksClient{}

	var resp resource.ConfigureResponse
	r.Configure(context.Background(), resource.ConfigureRequest{
		ProviderData: client,
	}, &resp)

	assert.False(t, resp.Diagnostics.HasError())
	assert.Equal(t, client, r.Client)
}

func TestPermissionResource_ImportState(t *testing.T) {
	// Verify the resource implements the ImportState interface
	var _ resource.ResourceWithImportState = &PermissionResource{}

	// Test that parseID correctly handles import ID format
	r := &PermissionResource{}

	tests := []struct {
		name            string
		importID        string
		expectedObjType string
		expectedObjID   string
		expectedPrinc   string
		expectError     bool
	}{
		{
			name:            "valid cluster import",
			importID:        "clusters/cluster-123/user@example.com",
			expectedObjType: "clusters",
			expectedObjID:   "cluster-123",
			expectedPrinc:   "user@example.com",
			expectError:     false,
		},
		{
			name:            "valid job import",
			importID:        "jobs/job-456/data-engineers",
			expectedObjType: "jobs",
			expectedObjID:   "job-456",
			expectedPrinc:   "data-engineers",
			expectError:     false,
		},
		{
			name:            "valid authorization import",
			importID:        "authorization/tokens/team-a",
			expectedObjType: "authorization",
			expectedObjID:   "tokens",
			expectedPrinc:   "team-a",
			expectError:     false,
		},
		{
			name:        "invalid format - too few parts",
			importID:    "clusters/cluster-123",
			expectError: true,
		},
		{
			name:        "invalid format - no slashes",
			importID:    "cluster-123-user",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			objectType, objectID, principal, err := r.parseID(tt.importID)

			if tt.expectError {
				assert.Error(t, err, "Expected error for invalid import ID")
			} else {
				assert.NoError(t, err, "Expected no error for valid import ID")
				assert.Equal(t, tt.expectedObjType, objectType, "Object type should match")
				assert.Equal(t, tt.expectedObjID, objectID, "Object ID should match")
				assert.Equal(t, tt.expectedPrinc, principal, "Principal should match")
			}
		})
	}
}
