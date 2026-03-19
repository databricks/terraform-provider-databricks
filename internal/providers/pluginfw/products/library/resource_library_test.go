package library

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResourceLibrary_ImplementsResourceWithModifyPlan(t *testing.T) {
	r := ResourceLibrary()
	_, ok := r.(resource.ResourceWithModifyPlan)
	assert.True(t, ok, "LibraryResource must implement ResourceWithModifyPlan")
}

func TestResourceLibrary_SchemaPreserved(t *testing.T) {
	r := ResourceLibrary()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	// Verify cluster_id exists and is required
	clusterAttr, ok := s.Attributes["cluster_id"]
	require.True(t, ok, "cluster_id attribute must exist")
	strAttr, ok := clusterAttr.(schema.StringAttribute)
	require.True(t, ok, "cluster_id must be a string attribute")
	assert.True(t, strAttr.Required, "cluster_id should be required")
	assert.Len(t, strAttr.PlanModifiers, 1, "cluster_id should have RequiresReplace plan modifier")

	// Verify all string attributes have RequiresReplace
	for _, field := range []string{"jar", "whl", "requirements"} {
		attr, ok := s.Attributes[field]
		if !ok {
			continue // some may be nested differently
		}
		strA, ok := attr.(schema.StringAttribute)
		if ok {
			assert.Len(t, strA.PlanModifiers, 1, "%s should have RequiresReplace plan modifier", field)
		}
	}

	// Verify id is computed and optional
	idAttr, ok := s.Attributes["id"]
	require.True(t, ok, "id attribute must exist")
	idStr, ok := idAttr.(schema.StringAttribute)
	require.True(t, ok, "id must be a string attribute")
	assert.True(t, idStr.Computed, "id should be computed")
	assert.True(t, idStr.Optional, "id should be optional")

	// Verify provider_config block exists (SdkV2 compatible = list nested block)
	pcBlock, ok := s.Blocks["provider_config"]
	require.True(t, ok, "provider_config block must exist")
	pcList, ok := pcBlock.(schema.ListNestedBlock)
	require.True(t, ok, "provider_config must be a list nested block (SdkV2 compatible)")
	assert.Len(t, pcList.Validators, 1, "provider_config should have SizeAtMost(1) validator")

	// Verify workspace_id inside provider_config
	wsAttr, ok := pcList.NestedObject.Attributes["workspace_id"]
	require.True(t, ok, "workspace_id must exist in provider_config")
	wsStr, ok := wsAttr.(schema.StringAttribute)
	require.True(t, ok, "workspace_id must be a string attribute")
	assert.True(t, wsStr.Required, "workspace_id should be required")
	assert.Len(t, wsStr.PlanModifiers, 1, "workspace_id should have RequiresReplaceIf plan modifier")
	assert.Len(t, wsStr.Validators, 2, "workspace_id should have 2 validators")
}

func TestResourceLibrary_ModifyPlan_SkipsDestroyAndNilClient(t *testing.T) {
	r := &LibraryResource{Client: nil}
	req := resource.ModifyPlanRequest{}
	resp := &resource.ModifyPlanResponse{}
	r.ModifyPlan(context.Background(), req, resp)
	assert.False(t, resp.Diagnostics.HasError(), "should not error on null plan with nil client")
}
