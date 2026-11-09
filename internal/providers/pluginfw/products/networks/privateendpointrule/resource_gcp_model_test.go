package privateendpointrule

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/require"
)

func TestGcpCreateAndReadNormalizeUnusedTargets(t *testing.T) {
	ctx := context.Background()
	const pscURI = "projects/p/regions/r/forwardingRules/endpoint"
	for _, tt := range []struct {
		name   string
		config gcpEndpointModel
		api    settings.GcpEndpoint
	}{
		{
			name: "Google APIs",
			config: gcpEndpointModel{GoogleApiEndpoints: []googleApiEndpointsModel{{
				Endpoints: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("storage.googleapis.com")}),
			}}},
			api: settings.GcpEndpoint{GoogleApiEndpoints: &settings.GoogleApiEndpoints{Endpoints: []string{"storage.googleapis.com"}}},
		},
		{
			name:   "all VPC-SC services",
			config: gcpEndpointModel{AllVpcScServices: types.BoolValue(true)},
			api:    settings.GcpEndpoint{AllVpcScServices: true},
		},
		{
			name:   "service attachment",
			config: gcpEndpointModel{ServiceAttachment: types.StringValue("projects/p/regions/r/serviceAttachments/a")},
			api:    settings.GcpEndpoint{ServiceAttachment: "projects/p/regions/r/serviceAttachments/a"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			apiTarget := tt.api
			apiTarget.PscEndpointUri = pscURI
			rule := &settings.NccPrivateEndpointRule{RuleId: "rule_id", GcpEndpoint: &apiTarget, ConnectionState: "ESTABLISHED"}
			r := &resourcePrivateEndpointRule{backoff: tightBackoff, api: &fakeAPI{
				create: func(_ context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
					require.Equal(t, &tt.api, req.PrivateEndpointRule.GcpEndpoint)
					return rule, nil
				},
				get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
					return rule, nil
				},
			}}
			plan := model{NetworkConnectivityConfigId: types.StringValue("ncc_id"), GcpEndpoint: []gcpEndpointModel{tt.config}}
			created := resource.CreateResponse{State: emptyState()}
			r.Create(ctx, resource.CreateRequest{Plan: rawPlan(t, ctx, plan)}, &created)
			fatalIfDiag(t, created.Diagnostics)
			want := tt.config
			want.PscEndpointUri = types.StringValue(pscURI)
			state := readModel(t, ctx, created.State)
			require.Equal(t, []gcpEndpointModel{want}, state.GcpEndpoint)

			// SDKv2 stores zero values for unused scalars, unlike a fresh PF plan.
			if state.GcpEndpoint[0].AllVpcScServices.IsNull() {
				state.GcpEndpoint[0].AllVpcScServices = types.BoolValue(false)
			}
			if state.GcpEndpoint[0].ServiceAttachment.IsNull() {
				state.GcpEndpoint[0].ServiceAttachment = types.StringValue("")
			}
			read := resource.ReadResponse{State: emptyState()}
			r.Read(ctx, resource.ReadRequest{State: rawState(t, ctx, state)}, &read)
			fatalIfDiag(t, read.Diagnostics)
			require.Equal(t, []gcpEndpointModel{want}, readModel(t, ctx, read.State).GcpEndpoint)
		})
	}
}

func TestGcpUpdateRequestsExcludeReadOnlyFields(t *testing.T) {
	ctx := context.Background()
	googleAPIs := gcpEndpointModel{GoogleApiEndpoints: []googleApiEndpointsModel{{
		Endpoints: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("storage.googleapis.com")}),
	}}}
	allServices := gcpEndpointModel{AllVpcScServices: types.BoolValue(true)}
	for _, tt := range []struct {
		name        string
		before      gcpEndpointModel
		after       gcpEndpointModel
		wantPayload string
	}{
		{"APIs to all services", googleAPIs, allServices, `{"gcp_endpoint":{"all_vpc_sc_services":true}}`},
		{"all services to APIs", allServices, googleAPIs, `{"gcp_endpoint":{"google_api_endpoints":{"endpoints":["storage.googleapis.com"]}}}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			previous := model{ID: types.StringValue("ncc_id/rule_id"), GcpEndpoint: []gcpEndpointModel{tt.before}}
			fillListDefaults(&previous)
			plan := previous
			tt.after.PscEndpointUri = types.StringValue("projects/p/regions/r/forwardingRules/endpoint")
			plan.GcpEndpoint = []gcpEndpointModel{tt.after}
			req, diags := plan.toUpdateRequest(ctx, previous)
			fatalIfDiag(t, diags)
			require.Equal(t, "gcp_endpoint", req.UpdateMask)
			payload, err := json.Marshal(req.PrivateEndpointRule)
			require.NoError(t, err)
			require.JSONEq(t, tt.wantPayload, string(payload))
		})
	}
}

func TestGcpReadOnlyChangeDoesNotUpdateTarget(t *testing.T) {
	previous := model{
		ID:          types.StringValue("ncc_id/rule_id"),
		GcpEndpoint: []gcpEndpointModel{{AllVpcScServices: types.BoolValue(true), PscEndpointUri: types.StringValue("old")}},
	}
	fillListDefaults(&previous)
	plan := previous
	plan.GcpEndpoint = []gcpEndpointModel{{AllVpcScServices: types.BoolValue(true), PscEndpointUri: types.StringValue("new")}}
	req, diags := plan.toUpdateRequest(context.Background(), previous)
	fatalIfDiag(t, diags)
	require.Empty(t, req.UpdateMask)
	require.Nil(t, req.PrivateEndpointRule.GcpEndpoint)
}
