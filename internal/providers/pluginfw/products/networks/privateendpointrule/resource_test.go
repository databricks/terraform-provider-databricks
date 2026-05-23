package privateendpointrule

import (
	"context"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	frameworkschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestResource_SchemaPreserved is a guard for the SDKv2-to-PluginFramework
// schema parity. `make diff-schema` is the authoritative check; this test
// catches the most common regressions cheaply.
func TestResource_SchemaPreserved(t *testing.T) {
	r := ResourcePrivateEndpointRule()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	s := resp.Schema

	// The framework does not surface "this is a RequiresReplace modifier" as a
	// typed predicate, so the cheap check is the description string emitted by
	// stringplanmodifier.RequiresReplace() ("...destroy and recreate...").
	requireReplaceFields := []string{"network_connectivity_config_id", "endpoint_service", "group_id", "resource_id"}
	for _, name := range requireReplaceFields {
		attr, ok := s.Attributes[name]
		require.True(t, ok, "%s attribute must exist", name)
		str, ok := attr.(frameworkschema.StringAttribute)
		require.True(t, ok, "%s must be a StringAttribute", name)
		var hasRequiresReplace bool
		for _, pm := range str.PlanModifiers {
			if containsSubstring(pm.Description(context.Background()), "destroy and recreate") {
				hasRequiresReplace = true
				break
			}
		}
		assert.True(t, hasRequiresReplace, "%s must carry a RequiresReplace plan modifier", name)
	}

	computedOnly := []string{"id", "rule_id", "account_id", "endpoint_name", "vpc_endpoint_id", "connection_state", "creation_time", "updated_time", "deactivated", "deactivated_at", "error_message", "gcp_endpoint"}
	for _, name := range computedOnly {
		attr, ok := s.Attributes[name]
		require.True(t, ok, "%s attribute must exist", name)
		assert.True(t, attr.IsComputed(), "%s must be Computed", name)
		assert.False(t, attr.IsRequired(), "%s must not be Required", name)
	}

	for _, name := range []string{"enabled", "domain_names", "resource_names"} {
		attr, ok := s.Attributes[name]
		require.True(t, ok, "%s attribute must exist", name)
		assert.True(t, attr.IsOptional(), "%s must be Optional", name)
		assert.True(t, attr.IsComputed(), "%s must be Computed", name)
	}
}

func containsSubstring(haystack, needle string) bool {
	for i := 0; i+len(needle) <= len(haystack); i++ {
		if haystack[i:i+len(needle)] == needle {
			return true
		}
	}
	return false
}

func TestModel_ToCreateRequest_AllUserFields(t *testing.T) {
	ctx := context.Background()
	domains, _ := types.ListValueFrom(ctx, types.StringType, []string{"a.example.com", "b.example.com"})
	resources, _ := types.ListValueFrom(ctx, types.StringType, []string{"bucket-1", "bucket-2"})
	m := model{
		NetworkConnectivityConfigId: types.StringValue("ncc-id"),
		EndpointService:             types.StringValue("com.amazonaws.example"),
		GroupId:                     types.StringValue("blob"),
		ResourceId:                  types.StringValue("/subscriptions/.../resourceId"),
		DomainNames:                 domains,
		ResourceNames:               resources,
	}
	got, diags := m.toCreateRequest(ctx)
	require.False(t, diags.HasError(), "diagnostics: %s", diags)
	want := &settings.CreatePrivateEndpointRuleRequest{
		NetworkConnectivityConfigId: "ncc-id",
		PrivateEndpointRule: settings.CreatePrivateEndpointRule{
			EndpointService: "com.amazonaws.example",
			GroupId:         "blob",
			ResourceId:      "/subscriptions/.../resourceId",
			DomainNames:     []string{"a.example.com", "b.example.com"},
			ResourceNames:   []string{"bucket-1", "bucket-2"},
		},
	}
	if diff := cmp.Diff(want, got, cmpopts.IgnoreFields(settings.CreatePrivateEndpointRule{}, "ForceSendFields")); diff != "" {
		t.Fatalf("CreatePrivateEndpointRuleRequest mismatch (-want +got):\n%s", diff)
	}
}

func TestModel_ToCreateRequest_OmitsNullLists(t *testing.T) {
	ctx := context.Background()
	m := model{
		NetworkConnectivityConfigId: types.StringValue("ncc-id"),
		GroupId:                     types.StringValue("blob"),
		ResourceId:                  types.StringValue("/path"),
		DomainNames:                 types.ListNull(types.StringType),
		ResourceNames:               types.ListNull(types.StringType),
	}
	got, diags := m.toCreateRequest(ctx)
	require.False(t, diags.HasError())
	assert.Nil(t, got.PrivateEndpointRule.DomainNames)
	assert.Nil(t, got.PrivateEndpointRule.ResourceNames)
}

func TestModel_ToUpdateRequest_OnlyChangedFields(t *testing.T) {
	ctx := context.Background()
	baseId := types.StringValue("ncc-id/rule-id")
	resourcesA, _ := types.ListValueFrom(ctx, types.StringType, []string{"bucket-1"})
	resourcesB, _ := types.ListValueFrom(ctx, types.StringType, []string{"bucket-2"})
	domainsA, _ := types.ListValueFrom(ctx, types.StringType, []string{"a.example.com"})
	domainsB, _ := types.ListValueFrom(ctx, types.StringType, []string{"b.example.com"})

	prev := model{
		ID:            baseId,
		Enabled:       types.BoolValue(true),
		DomainNames:   domainsA,
		ResourceNames: resourcesA,
	}

	tests := []struct {
		name     string
		plan     model
		wantMask []string
	}{
		{
			name: "enabled only",
			plan: model{
				ID:            baseId,
				Enabled:       types.BoolValue(false),
				DomainNames:   domainsA,
				ResourceNames: resourcesA,
			},
			wantMask: []string{"enabled"},
		},
		{
			name: "domain_names only",
			plan: model{
				ID:            baseId,
				Enabled:       types.BoolValue(true),
				DomainNames:   domainsB,
				ResourceNames: resourcesA,
			},
			wantMask: []string{"domain_names"},
		},
		{
			name: "resource_names only",
			plan: model{
				ID:            baseId,
				Enabled:       types.BoolValue(true),
				DomainNames:   domainsA,
				ResourceNames: resourcesB,
			},
			wantMask: []string{"resource_names"},
		},
		{
			name: "all three",
			plan: model{
				ID:            baseId,
				Enabled:       types.BoolValue(false),
				DomainNames:   domainsB,
				ResourceNames: resourcesB,
			},
			wantMask: []string{"enabled", "domain_names", "resource_names"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, mask, diags := tt.plan.toUpdateRequest(ctx, prev)
			require.False(t, diags.HasError())
			assert.Equal(t, tt.wantMask, mask)
		})
	}
}

func TestModel_FromAPI(t *testing.T) {
	ctx := context.Background()
	rule := &settings.NccPrivateEndpointRule{
		AccountId:                   "acct",
		ConnectionState:             "PENDING",
		CreationTime:                111,
		Deactivated:                 false,
		DeactivatedAt:               0,
		DomainNames:                 []string{"a.example.com"},
		Enabled:                     true,
		EndpointName:                "vpce-name",
		EndpointService:             "com.amazonaws.example",
		ErrorMessage:                "",
		GcpEndpoint:                 &settings.GcpEndpoint{PscEndpointUri: "uri", ServiceAttachment: "attach"},
		GroupId:                     "blob",
		NetworkConnectivityConfigId: "ncc",
		ResourceId:                  "/path",
		ResourceNames:               []string{"bucket-1"},
		RuleId:                      "rule",
		UpdatedTime:                 222,
		VpcEndpointId:               "vpce-abc",
	}
	var m model
	diags := m.fromAPI(ctx, rule)
	require.False(t, diags.HasError())
	assert.Equal(t, "rule", m.RuleId.ValueString())
	assert.Equal(t, "acct", m.AccountId.ValueString())
	assert.Equal(t, "blob", m.GroupId.ValueString())
	assert.Equal(t, "PENDING", m.ConnectionState.ValueString())
	assert.Equal(t, "vpce-abc", m.VpcEndpointId.ValueString())
	assert.Equal(t, "vpce-name", m.EndpointName.ValueString())
	assert.Equal(t, int64(111), m.CreationTime.ValueInt64())
	assert.Equal(t, int64(222), m.UpdatedTime.ValueInt64())
	require.NotNil(t, m.GcpEndpoint)
	assert.Equal(t, "uri", m.GcpEndpoint.PscEndpointUri.ValueString())
	assert.Equal(t, "attach", m.GcpEndpoint.ServiceAttachment.ValueString())

	var domainNames []string
	d := m.DomainNames.ElementsAs(ctx, &domainNames, false)
	require.False(t, d.HasError())
	assert.Equal(t, []string{"a.example.com"}, domainNames)
}

func TestWaitForPrivateEndpointRuleCreate_AsyncPolling(t *testing.T) {
	creating := &settings.NccPrivateEndpointRule{
		RuleId:                      "rule",
		NetworkConnectivityConfigId: "ncc",
		ConnectionState:             "CREATING",
	}
	pending := &settings.NccPrivateEndpointRule{
		RuleId:                      "rule",
		NetworkConnectivityConfigId: "ncc",
		ConnectionState:             "PENDING",
		VpcEndpointId:               "vpce-abc",
	}
	m := mocks.NewMockAccountClient(t)
	e := m.GetMockNetworkConnectivityAPI().EXPECT()
	e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc", "rule").Return(creating, nil).Once()
	e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc", "rule").Return(pending, nil).Once()
	got, err := waitForPrivateEndpointRuleCreate(context.Background(), m.AccountClient, "ncc", "rule", 30*time.Second)
	require.NoError(t, err)
	assert.Equal(t, "PENDING", string(got.ConnectionState))
	assert.Equal(t, "vpce-abc", got.VpcEndpointId)
}

func TestWaitForPrivateEndpointRuleCreate_CreateFailed(t *testing.T) {
	failed := &settings.NccPrivateEndpointRule{
		RuleId:          "rule",
		ConnectionState: "CREATE_FAILED",
		ErrorMessage:    "quota exceeded",
	}
	m := mocks.NewMockAccountClient(t)
	e := m.GetMockNetworkConnectivityAPI().EXPECT()
	e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc", "rule").Return(failed, nil)
	_, err := waitForPrivateEndpointRuleCreate(context.Background(), m.AccountClient, "ncc", "rule", 30*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "quota exceeded")
}

func TestWaitForPrivateEndpointRuleCreate_TerminalRejected(t *testing.T) {
	rejected := &settings.NccPrivateEndpointRule{
		RuleId:          "rule",
		ConnectionState: "REJECTED",
	}
	m := mocks.NewMockAccountClient(t)
	e := m.GetMockNetworkConnectivityAPI().EXPECT()
	e.GetPrivateEndpointRuleByNetworkConnectivityConfigIdAndPrivateEndpointRuleId(mock.Anything, "ncc", "rule").Return(rejected, nil)
	_, err := waitForPrivateEndpointRuleCreate(context.Background(), m.AccountClient, "ncc", "rule", 30*time.Second)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "REJECTED")
}
