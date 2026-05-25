package privateendpointrule

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/databricks/terraform-provider-databricks/internal/retrier"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tightBackoff keeps the retrier's sleep between polling attempts negligible.
var tightBackoff = retrier.BackoffPolicy{
	Initial: time.Nanosecond,
	Maximum: time.Nanosecond,
	Factor:  1,
}

// fakeAPI satisfies apiClient with closure fields. Tests wire only the
// methods they exercise; nil closures panic if called unexpectedly, which
// pinpoints the unexpected call more directly than a mock framework would.
type fakeAPI struct {
	create func(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	get    func(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	update func(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
	del    func(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error)
}

func (f *fakeAPI) CreatePrivateEndpointRule(ctx context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
	return f.create(ctx, req)
}

func (f *fakeAPI) GetPrivateEndpointRule(ctx context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
	return f.get(ctx, req)
}

func (f *fakeAPI) UpdatePrivateEndpointRule(ctx context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
	return f.update(ctx, req)
}

func (f *fakeAPI) DeletePrivateEndpointRule(ctx context.Context, req settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
	return f.del(ctx, req)
}

// fillListDefaults sets list fields left as zero-value types.List from
// emptyModel(). The framework rejects a list without an element type; this
// lets tests use terse model literals.
func fillListDefaults(m *model) {
	base := emptyModel()
	if m.DomainNames.ElementType(context.Background()) == nil {
		m.DomainNames = base.DomainNames
	}
	if m.ResourceNames.ElementType(context.Background()) == nil {
		m.ResourceNames = base.ResourceNames
	}
}

// rawPlan builds a tfsdk.Plan whose Raw value reflects m, equivalent to what
// Terraform Core hands a resource's CRUD method during plan/apply.
func rawPlan(t *testing.T, ctx context.Context, m model) tfsdk.Plan {
	t.Helper()
	fillListDefaults(&m)
	p := tfsdk.Plan{Schema: resourceSchema()}
	if diags := p.Set(ctx, &m); diags.HasError() {
		t.Fatalf("rawPlan: %v", diags)
	}
	return p
}

// rawState is rawPlan's analogue for tfsdk.State.
func rawState(t *testing.T, ctx context.Context, m model) tfsdk.State {
	t.Helper()
	fillListDefaults(&m)
	s := tfsdk.State{Schema: resourceSchema()}
	if diags := s.Set(ctx, &m); diags.HasError() {
		t.Fatalf("rawState: %v", diags)
	}
	return s
}

// emptyState is what Terraform Core hands a resource on a fresh Create or
// ImportState: a Raw value that is null at the root.
func emptyState() tfsdk.State {
	return tfsdk.State{Schema: resourceSchema()}
}

// readModel extracts the framework state back into a model for inspection.
// It is the corollary of rawState/rawPlan on the response side.
func readModel(t *testing.T, ctx context.Context, s tfsdk.State) model {
	t.Helper()
	var m model
	if diags := s.Get(ctx, &m); diags.HasError() {
		t.Fatalf("readModel: %v", diags)
	}
	return m
}

// fatalIfDiag aborts the test if a PF diagnostics slice carries any error.
func fatalIfDiag(t *testing.T, diags diag.Diagnostics) {
	t.Helper()
	if diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
}

// ---------- Behavioural tests for the CRUD contract ----------

func TestCreate_PollsUntilNotCreating_ThenPersistsFinalState(t *testing.T) {
	ctx := context.Background()
	var getCalls int
	api := &fakeAPI{
		create: func(_ context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			if req.NetworkConnectivityConfigId != "ncc-1" || req.PrivateEndpointRule.EndpointService != "com.amazonaws.example" {
				t.Errorf("create request body: %+v", req)
			}
			// Echo the NCC ID like real servers do, so Create's polling Get
			// can address the new rule via pendingRule.NetworkConnectivityConfigId.
			return &settings.NccPrivateEndpointRule{
				RuleId:                      "rule-1",
				NetworkConnectivityConfigId: req.NetworkConnectivityConfigId,
				ConnectionState:             "CREATING",
			}, nil
		},
		get: func(_ context.Context, req settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			if req.NetworkConnectivityConfigId != "ncc-1" || req.PrivateEndpointRuleId != "rule-1" {
				t.Errorf("get key: %+v", req)
			}
			getCalls++
			if getCalls < 2 {
				return &settings.NccPrivateEndpointRule{RuleId: "rule-1", ConnectionState: "CREATING"}, nil
			}
			return &settings.NccPrivateEndpointRule{
				RuleId:          "rule-1",
				ConnectionState: "PENDING",
				VpcEndpointId:   "vpce-abc",
			}, nil
		},
	}
	r := &resourcePrivateEndpointRule{api: api, backoff: tightBackoff}

	req := resource.CreateRequest{Plan: rawPlan(t, ctx, model{
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
		EndpointService:             types.StringValue("com.amazonaws.example"),
	})}
	resp := resource.CreateResponse{State: emptyState()}
	r.Create(ctx, req, &resp)

	fatalIfDiag(t, resp.Diagnostics)
	if getCalls != 2 {
		t.Errorf("get called %d times, want 2 (CREATING then PENDING)", getCalls)
	}
	final := readModel(t, ctx, resp.State)
	if got, want := final.ID.ValueString(), "ncc-1/rule-1"; got != want {
		t.Errorf("ID: got %q, want %q", got, want)
	}
	if got, want := final.VpcEndpointId.ValueString(), "vpce-abc"; got != want {
		t.Errorf("VpcEndpointId: got %q, want %q", got, want)
	}
	if got, want := final.ConnectionState.ValueString(), "PENDING"; got != want {
		t.Errorf("ConnectionState: got %q, want %q", got, want)
	}
}

func TestCreate_CreateFailedSurfacesServerErrorMessage(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		create: func(_ context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return &settings.NccPrivateEndpointRule{
				RuleId:                      "rule-1",
				NetworkConnectivityConfigId: req.NetworkConnectivityConfigId,
				ConnectionState:             "CREATING",
			}, nil
		},
		get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return &settings.NccPrivateEndpointRule{
				RuleId:          "rule-1",
				ConnectionState: "CREATE_FAILED",
				ErrorMessage:    "quota exceeded",
			}, nil
		},
	}
	r := &resourcePrivateEndpointRule{api: api, backoff: tightBackoff}

	req := resource.CreateRequest{Plan: rawPlan(t, ctx, model{
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
	})}
	resp := resource.CreateResponse{State: emptyState()}
	r.Create(ctx, req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected Create to surface CREATE_FAILED as a diagnostic")
	}
	if got := resp.Diagnostics.Errors()[0].Detail(); !strings.Contains(got, "quota exceeded") {
		t.Errorf("error detail: got %q, want it to contain %q", got, "quota exceeded")
	}
}

// TestCreate_PersistsStateBeforePolling guards the invariant that a polling
// failure must still leave a destroyable resource. Without the intermediate
// State.Set in Create, this would regress silently.
func TestCreate_PersistsStateBeforePolling(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		create: func(_ context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return &settings.NccPrivateEndpointRule{
				RuleId:                      "rule-1",
				NetworkConnectivityConfigId: req.NetworkConnectivityConfigId,
				ConnectionState:             "CREATING",
			}, nil
		},
		get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return nil, errors.New("network unreachable")
		},
	}
	r := &resourcePrivateEndpointRule{api: api, backoff: tightBackoff}

	req := resource.CreateRequest{Plan: rawPlan(t, ctx, model{
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
	})}
	resp := resource.CreateResponse{State: emptyState()}
	r.Create(ctx, req, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("expected polling error")
	}
	final := readModel(t, ctx, resp.State)
	if got, want := final.ID.ValueString(), "ncc-1/rule-1"; got != want {
		t.Errorf("after polling failure, state.id: got %q, want %q (destroy could not clean up)", got, want)
	}
}

// TestCreate_SendsGcpEndpoint guards against silently dropping
// gcp_endpoint.service_attachment from the create request. SDKv2 copied the
// whole struct via DataToStructPointer; the PF translation must not regress.
func TestCreate_SendsGcpEndpoint(t *testing.T) {
	ctx := context.Background()
	var sentReq *settings.CreatePrivateEndpointRuleRequest
	api := &fakeAPI{
		create: func(_ context.Context, req settings.CreatePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			sentReq = &req
			return &settings.NccPrivateEndpointRule{
				RuleId:                      "rule-1",
				NetworkConnectivityConfigId: req.NetworkConnectivityConfigId,
				ConnectionState:             "PENDING",
			}, nil
		},
		get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return &settings.NccPrivateEndpointRule{RuleId: "rule-1", ConnectionState: "PENDING"}, nil
		},
	}
	r := &resourcePrivateEndpointRule{api: api, backoff: tightBackoff}

	req := resource.CreateRequest{Plan: rawPlan(t, ctx, model{
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
		GcpEndpoint: []gcpEndpointModel{{
			ServiceAttachment: types.StringValue("projects/p/regions/r/serviceAttachments/sa"),
		}},
	})}
	resp := resource.CreateResponse{State: emptyState()}
	r.Create(ctx, req, &resp)

	fatalIfDiag(t, resp.Diagnostics)
	if sentReq == nil {
		t.Fatal("Create never reached the API")
	}
	if sentReq.PrivateEndpointRule.GcpEndpoint == nil {
		t.Fatal("GcpEndpoint dropped from create request")
	}
	if got, want := sentReq.PrivateEndpointRule.GcpEndpoint.ServiceAttachment, "projects/p/regions/r/serviceAttachments/sa"; got != want {
		t.Errorf("ServiceAttachment: got %q, want %q", got, want)
	}
}

func TestRead_404RemovesFromState(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return nil, apierr.ErrNotFound
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, model{ID: types.StringValue("ncc-1/rule-1")})
	req := resource.ReadRequest{State: startState}
	resp := resource.ReadResponse{State: startState}
	r.Read(ctx, req, &resp)

	fatalIfDiag(t, resp.Diagnostics)
	if !resp.State.Raw.IsNull() {
		t.Errorf("state should be removed on 404, got Raw: %v", resp.State.Raw)
	}
}

func TestRead_NonMissingErrorSurfacesAsDiagnostic(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		get: func(context.Context, settings.GetPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return nil, errors.New("internal server error")
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, model{ID: types.StringValue("ncc-1/rule-1")})
	resp := resource.ReadResponse{State: startState}
	r.Read(ctx, resource.ReadRequest{State: startState}, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("non-404 Read errors should surface as diagnostics")
	}
}

func TestUpdate_SkipsAPICallWhenMaskEmpty(t *testing.T) {
	ctx := context.Background()
	unchanged := model{
		ID:                          types.StringValue("ncc-1/rule-1"),
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
		Enabled:                     types.BoolValue(true),
	}
	api := &fakeAPI{
		update: func(context.Context, settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			t.Error("Update was called for an unchanged plan")
			return nil, nil
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, unchanged)
	req := resource.UpdateRequest{Plan: rawPlan(t, ctx, unchanged), State: startState}
	resp := resource.UpdateResponse{State: startState}
	r.Update(ctx, req, &resp)

	fatalIfDiag(t, resp.Diagnostics)
}

func TestUpdate_SendsMaskedFieldsToServer(t *testing.T) {
	ctx := context.Background()
	prev := model{
		ID:                          types.StringValue("ncc-1/rule-1"),
		NetworkConnectivityConfigId: types.StringValue("ncc-1"),
		Enabled:                     types.BoolValue(true),
	}
	next := prev
	next.Enabled = types.BoolValue(false)

	var sentReq *settings.UpdateNccPrivateEndpointRuleRequest
	api := &fakeAPI{
		update: func(_ context.Context, req settings.UpdateNccPrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			sentReq = &req
			return &settings.NccPrivateEndpointRule{RuleId: "rule-1", Enabled: false, ConnectionState: "PENDING"}, nil
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, prev)
	req := resource.UpdateRequest{Plan: rawPlan(t, ctx, next), State: startState}
	resp := resource.UpdateResponse{State: startState}
	r.Update(ctx, req, &resp)

	fatalIfDiag(t, resp.Diagnostics)
	if sentReq == nil {
		t.Fatal("Update never reached the API")
	}
	if got, want := sentReq.UpdateMask, "enabled"; got != want {
		t.Errorf("UpdateMask: got %q, want %q", got, want)
	}
	if sentReq.PrivateEndpointRule.Enabled != false {
		t.Errorf("enabled in request: got %v, want false", sentReq.PrivateEndpointRule.Enabled)
	}
}

func TestDelete_404IsIdempotent(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		del: func(context.Context, settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return nil, apierr.ErrNotFound
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, model{ID: types.StringValue("ncc-1/rule-1")})
	resp := resource.DeleteResponse{State: startState}
	r.Delete(ctx, resource.DeleteRequest{State: startState}, &resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("Delete with 404 should be idempotent: %v", resp.Diagnostics)
	}
}

func TestDelete_SurfacesNon404Errors(t *testing.T) {
	ctx := context.Background()
	api := &fakeAPI{
		del: func(context.Context, settings.DeletePrivateEndpointRuleRequest) (*settings.NccPrivateEndpointRule, error) {
			return nil, errors.New("forbidden")
		},
	}
	r := &resourcePrivateEndpointRule{api: api}

	startState := rawState(t, ctx, model{ID: types.StringValue("ncc-1/rule-1")})
	resp := resource.DeleteResponse{State: startState}
	r.Delete(ctx, resource.DeleteRequest{State: startState}, &resp)

	if !resp.Diagnostics.HasError() {
		t.Fatal("non-404 Delete errors should surface as diagnostics")
	}
}

func TestImportState_ParsesCompositeID(t *testing.T) {
	ctx := context.Background()
	r := &resourcePrivateEndpointRule{}
	resp := resource.ImportStateResponse{State: emptyState()}
	r.ImportState(ctx, resource.ImportStateRequest{ID: packID("ncc-1", "rule-1")}, &resp)

	fatalIfDiag(t, resp.Diagnostics)
	// ImportState only seeds the ID and its two components; Read fills the
	// rest, so we only assert on what ImportState is contractually responsible
	// for.
	got := readModel(t, ctx, resp.State)
	if want := "ncc-1/rule-1"; got.ID.ValueString() != want {
		t.Errorf("ID: got %q, want %q", got.ID.ValueString(), want)
	}
	if want := "ncc-1"; got.NetworkConnectivityConfigId.ValueString() != want {
		t.Errorf("NetworkConnectivityConfigId: got %q, want %q", got.NetworkConnectivityConfigId.ValueString(), want)
	}
	if want := "rule-1"; got.RuleId.ValueString() != want {
		t.Errorf("RuleId: got %q, want %q", got.RuleId.ValueString(), want)
	}
}

func TestImportState_RejectsInvalidID(t *testing.T) {
	ctx := context.Background()
	r := &resourcePrivateEndpointRule{}
	for _, id := range []string{"", "/", "no-separator", "/leading", "trailing/"} {
		t.Run(id, func(t *testing.T) {
			resp := resource.ImportStateResponse{State: emptyState()}
			r.ImportState(ctx, resource.ImportStateRequest{ID: id}, &resp)
			if !resp.Diagnostics.HasError() {
				t.Errorf("ID %q should produce a diagnostic", id)
			}
		})
	}
}

// ---------- Translation logic ----------

// TestToUpdateRequest_ComputesMaskFromChangedFields is the one
// translation-layer test with non-trivial behaviour: deciding which fields
// belong in the update_mask based on plan-vs-state equality. The other
// translations (toCreateRequest, fromAPI) are exercised by the Create/Read
// behavioural tests above.
func TestToUpdateRequest_ComputesMaskFromChangedFields(t *testing.T) {
	ctx := context.Background()
	baseId := types.StringValue(packID("ncc-id", "rule-id"))
	resA, _ := types.ListValueFrom(ctx, types.StringType, []string{"bucket-1"})
	resB, _ := types.ListValueFrom(ctx, types.StringType, []string{"bucket-2"})
	domA, _ := types.ListValueFrom(ctx, types.StringType, []string{"a.example.com"})
	domB, _ := types.ListValueFrom(ctx, types.StringType, []string{"b.example.com"})

	prev := model{ID: baseId, Enabled: types.BoolValue(true), DomainNames: domA, ResourceNames: resA}

	tests := []struct {
		name     string
		plan     model
		wantMask string
	}{
		{"nothing changed", prev, ""},
		{"enabled only", model{ID: baseId, Enabled: types.BoolValue(false), DomainNames: domA, ResourceNames: resA}, "enabled"},
		{"domain_names only", model{ID: baseId, Enabled: types.BoolValue(true), DomainNames: domB, ResourceNames: resA}, "domain_names"},
		{"resource_names only", model{ID: baseId, Enabled: types.BoolValue(true), DomainNames: domA, ResourceNames: resB}, "resource_names"},
		{"all three", model{ID: baseId, Enabled: types.BoolValue(false), DomainNames: domB, ResourceNames: resB}, "enabled,domain_names,resource_names"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, diags := tt.plan.toUpdateRequest(ctx, prev)
			fatalIfDiag(t, diags)
			if req.UpdateMask != tt.wantMask {
				t.Errorf("UpdateMask: got %q, want %q", req.UpdateMask, tt.wantMask)
			}
		})
	}
}

// ---------- Polling predicate ----------

func TestIsStillCreating(t *testing.T) {
	transportErr := errors.New("transport failure")
	tests := []struct {
		name  string
		state settings.NccPrivateEndpointRulePrivateLinkConnectionState
		err   error
		want  bool
	}{
		{"CREATING → retry", settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreating, nil, true},
		{"PENDING → done", settings.NccPrivateEndpointRulePrivateLinkConnectionStatePending, nil, false},
		{"ESTABLISHED → done", settings.NccPrivateEndpointRulePrivateLinkConnectionStateEstablished, nil, false},
		{"CREATE_FAILED → done", settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreateFailed, nil, false},
		{"REJECTED → done", settings.NccPrivateEndpointRulePrivateLinkConnectionStateRejected, nil, false},
		{"any error halts retry", settings.NccPrivateEndpointRulePrivateLinkConnectionStateCreating, transportErr, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rule := &settings.NccPrivateEndpointRule{ConnectionState: tt.state}
			if got := isStillCreating(rule, tt.err); got != tt.want {
				t.Errorf("isStillCreating(state=%s, err=%v) = %v, want %v", tt.state, tt.err, got, tt.want)
			}
		})
	}
}

// ---------- ID round-trip ----------

func TestUnpackID_RejectsGarbage(t *testing.T) {
	for _, id := range []string{"", "/", "//", "no-separator", "/leading", "trailing/"} {
		t.Run(id, func(t *testing.T) {
			if _, _, err := unpackID(id); err == nil {
				t.Errorf("unpackID(%q) = nil err, want error", id)
			}
		})
	}
}

func FuzzPackUnpackID_Roundtrip(f *testing.F) {
	f.Add("ncc-1", "rule-1")
	f.Add("a", "b")
	f.Add("with-dash-and.dot", "uuid-like-12345678")
	f.Fuzz(func(t *testing.T, nccId, ruleId string) {
		// packID makes no promises for inputs containing "/" or empty parts;
		// callers normalize before calling, so skip those rather than assert.
		if nccId == "" || ruleId == "" || strings.Contains(nccId, "/") || strings.Contains(ruleId, "/") {
			t.Skip()
		}
		composed := packID(nccId, ruleId)
		gotNcc, gotRule, err := unpackID(composed)
		if err != nil {
			t.Fatalf("unpackID(%q) returned err: %v", composed, err)
		}
		if gotNcc != nccId || gotRule != ruleId {
			t.Errorf("roundtrip mismatch: pack(%q, %q) = %q, unpack → (%q, %q)", nccId, ruleId, composed, gotNcc, gotRule)
		}
	})
}
