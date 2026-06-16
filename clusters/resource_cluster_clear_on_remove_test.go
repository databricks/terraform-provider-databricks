package clusters

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"

	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/mock"
)

// These tests cover the ES-1927346 behavior and the clear_cloud_attributes_on_remove
// fix for databricks_cluster cloud-attribute blocks. The platform echoes default
// cloud attributes the user never set; removing a block was silently suppressed,
// so a value could not be cleared through Terraform.
//
// Coverage maps to the behavior tables in the design doc, at three altitudes:
//   - TestClusterAwsAttributesClearBehavior: the diff altitude (given state +
//     config, what does plan produce) across both flag values, plus the
//     explicit-empty and empty-block edges.
//   - TestClusterReadDropsUnconfiguredServerDefaults: the read altitude (server
//     defaults for an unconfigured block are dropped, never stored).
//   - TestClusterUpdateRemoveAwsAttributes: the update altitude (whether the diff
//     flows into a clearing Edit).
//
// These tests are deliberately written without the qa fixture harness, driving
// the resource directly with a mocked SDK workspace client. The SDK mock client
// uses testify/mock for its matchers (mock.Anything, mock.MatchedBy); assertions
// are plain t.Fatalf/t.Errorf.

const (
	testInstanceProfileArn  = "arn:aws:iam::123:instance-profile/Foo"
	otherInstanceProfileArn = "arn:aws:iam::123:instance-profile/Bar"
)

// clusterStateWithAws builds the terraform state of an existing cluster with the
// given aws_attributes nested fields. An empty arn or avail is omitted from state.
func clusterStateWithAws(arn, avail string) *terraform.InstanceState {
	attrs := map[string]string{
		"cluster_name":            "test",
		"spark_version":           "13.3.x-scala2.12",
		"node_type_id":            "i3.xlarge",
		"num_workers":             "2",
		"autotermination_minutes": "60", // schema default; present in real state, must match to avoid a spurious diff
		"aws_attributes.#":        "1",
	}
	if arn != "" {
		attrs["aws_attributes.0.instance_profile_arn"] = arn
	}
	if avail != "" {
		attrs["aws_attributes.0.availability"] = avail
	}
	return &terraform.InstanceState{ID: "abc-123", Attributes: attrs}
}

// priorStateWithInstanceProfile is a cluster with only an instance profile set.
func priorStateWithInstanceProfile() *terraform.InstanceState {
	return clusterStateWithAws(testInstanceProfileArn, "")
}

// clusterConfig builds a cluster config map with the given aws_attributes value.
// aws is the raw value for the block: an empty list ([]any{}) means the block is
// removed, which is how Terraform encodes a removed TypeList block in raw config
// (an omitted key would decode to cty null and mask block-deletion bugs).
func clusterConfig(clearOnRemove bool, aws []any) map[string]any {
	c := map[string]any{
		"cluster_name":   "test",
		"spark_version":  "13.3.x-scala2.12",
		"node_type_id":   "i3.xlarge",
		"num_workers":    2,
		"aws_attributes": aws,
	}
	if clearOnRemove {
		c["clear_cloud_attributes_on_remove"] = true
	}
	return c
}

// clusterConfigWithoutAws is the cluster config with the aws_attributes block removed.
func clusterConfigWithoutAws(clearOnRemove bool) map[string]any {
	return clusterConfig(clearOnRemove, []any{})
}

// testClient returns a DatabricksClient pointed at a workspace host with the
// given mocked workspace client injected as the cached client.
func testClient(w *mocks.MockWorkspaceClient) *common.DatabricksClient {
	c := &common.DatabricksClient{
		DatabricksClient: &client.DatabricksClient{
			Config: &config.Config{
				Host:  "https://test.cloud.databricks.com",
				Token: "x",
			},
		},
	}
	c.SetWorkspaceClient(w.WorkspaceClient)
	return c
}

// rawConfigDiff drives the resource diff with a cty-valued RawConfig so the
// clear_cloud_attributes_on_remove suppressor (which reads GetRawConfigAt) sees the configured
// value; terraform.NewResourceConfigRaw alone does not populate cty values.
// Missing object attributes decode to null, so a partial config is sufficient.
func rawConfigDiff(t *testing.T, r *schema.Resource, is *terraform.InstanceState, config map[string]any, c *common.DatabricksClient) (*terraform.InstanceDiff, error) {
	t.Helper()
	block := (&schema.Resource{Schema: r.Schema}).CoreConfigSchema()
	b, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("marshal config: %v", err)
	}
	ctyVal, err := ctyjson.Unmarshal(b, block.ImpliedType())
	if err != nil {
		t.Fatalf("decode config into cty: %v", err)
	}
	is.RawConfig = ctyVal
	rc := terraform.NewResourceConfigShimmed(ctyVal, block)
	return r.Diff(context.Background(), is, rc, c)
}

// attrChange is an expected diff on a single attribute: from old value to new.
type attrChange struct{ from, to string }

// assertAttrChange checks the diff for key against want (nil = expect no diff).
func assertAttrChange(t *testing.T, diff *terraform.InstanceDiff, key string, want *attrChange) {
	t.Helper()
	var attr *terraform.ResourceAttrDiff
	if diff != nil {
		attr = diff.Attributes[key]
	}
	if want == nil {
		if attr != nil {
			t.Errorf("%s: expected no diff, got %#v", key, attr)
		}
		return
	}
	if attr == nil {
		t.Fatalf("%s: expected a diff, got none (full diff: %#v)", key, diff)
	}
	if attr.Old != want.from {
		t.Errorf("%s: old = %q, want %q", key, attr.Old, want.from)
	}
	if attr.New != want.to {
		t.Errorf("%s: new = %q, want %q", key, attr.New, want.to)
	}
}

// TestClusterAwsAttributesClearBehavior covers the diff-altitude rows of the
// behavior tables in the design doc, across both flag values, plus the
// explicit-empty and empty-block edges. Each case asserts the expected change
// (or its absence) on the user-set field (instance_profile_arn) and on a
// server-default sibling (availability).
func TestClusterAwsAttributesClearBehavior(t *testing.T) {
	r := ResourceCluster().ToResource()
	withArn := func(v string) []any { return []any{map[string]any{"instance_profile_arn": v}} }
	removed := []any{}
	emptyBlock := []any{map[string]any{}}
	const spot = "SPOT_WITH_FALLBACK"

	for _, tc := range []struct {
		name       string
		clear      bool
		config     []any
		stateArn   string
		stateAvail string
		wantArn    *attrChange
		wantAvail  *attrChange
	}{
		// clear_cloud_attributes_on_remove = false (historical behavior, unchanged)
		{"F1 unchanged kept block", false, withArn(testInstanceProfileArn), testInstanceProfileArn, "", nil, nil},
		{"F2 change value", false, withArn(otherInstanceProfileArn), testInstanceProfileArn, "", &attrChange{testInstanceProfileArn, otherInstanceProfileArn}, nil},
		{"F3 remove block keeps value (the bug)", false, removed, testInstanceProfileArn, "", nil, nil},
		{"F4 server default suppressed", false, withArn(testInstanceProfileArn), testInstanceProfileArn, spot, nil, nil},
		// clear_cloud_attributes_on_remove = true
		{"T1 change value", true, withArn(otherInstanceProfileArn), testInstanceProfileArn, "", &attrChange{testInstanceProfileArn, otherInstanceProfileArn}, nil},
		{"T2 remove block clears block", true, removed, testInstanceProfileArn, spot, &attrChange{testInstanceProfileArn, ""}, &attrChange{spot, ""}},
		{"T3 kept block, server default not cleared", true, withArn(testInstanceProfileArn), testInstanceProfileArn, spot, nil, nil},
		// edges not in the doc tables
		{"E1 explicit empty clears without flag", false, withArn(""), testInstanceProfileArn, "", &attrChange{testInstanceProfileArn, ""}, nil},
		{"E2 empty block is kept, not deleted", true, emptyBlock, testInstanceProfileArn, "", nil, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := mocks.NewMockWorkspaceClient(t)
			diff, err := rawConfigDiff(t, r, clusterStateWithAws(tc.stateArn, tc.stateAvail), clusterConfig(tc.clear, tc.config), testClient(w))
			if err != nil {
				t.Fatalf("diff: %v", err)
			}
			assertAttrChange(t, diff, "aws_attributes.0.instance_profile_arn", tc.wantArn)
			assertAttrChange(t, diff, "aws_attributes.0.availability", tc.wantAvail)
		})
	}
}

// TestClusterReadDropsUnconfiguredServerDefaults covers the "default never stored"
// rows (the availability-only rows under both flag values): on read of an
// existing cluster, server defaults for a block the user never configured are
// discarded and never enter state, so they cannot produce a diff. This is
// StructToData read behavior, independent of clear_cloud_attributes_on_remove.
func TestClusterReadDropsUnconfiguredServerDefaults(t *testing.T) {
	is := &terraform.InstanceState{
		ID: "abc-123",
		Attributes: map[string]string{
			"cluster_name":            "test",
			"spark_version":           "13.3.x-scala2.12",
			"node_type_id":            "i3.xlarge",
			"num_workers":             "2",
			"autotermination_minutes": "60",
			// aws_attributes intentionally absent: the user never configured it.
		},
	}
	d, err := schema.InternalMap(clusterSchema).Data(is, nil)
	if err != nil {
		t.Fatalf("build resource data: %v", err)
	}
	if d.IsNewResource() {
		t.Fatal("precondition: this models a read of an existing resource, but IsNewResource() is true")
	}

	// The platform returns aws_attributes populated with defaults the user never set.
	info := &compute.ClusterDetails{
		ClusterId:     "abc-123",
		ClusterName:   "test",
		SparkVersion:  "13.3.x-scala2.12",
		NodeTypeId:    "i3.xlarge",
		NumWorkers:    2,
		AwsAttributes: &compute.AwsAttributes{Availability: "SPOT_WITH_FALLBACK", ZoneId: "us-east-1b"},
	}
	if err := common.StructToData(info, clusterSchema, d); err != nil {
		t.Fatalf("StructToData: %v", err)
	}

	if n := d.Get("aws_attributes.#").(int); n != 0 {
		t.Errorf("server defaults for an unconfigured block must be dropped on read; aws_attributes.# = %d, want 0", n)
	}
}

// TestClusterUpdateRemoveAwsAttributes checks the update altitude: the
// suppressed diff sends no Edit, while the clear_cloud_attributes_on_remove diff sends an Edit that
// drops aws_attributes (full-replacement semantics then clear the value). The
// Edit expectation firing (or its absence) is the proof.
func TestClusterUpdateRemoveAwsAttributes(t *testing.T) {
	ctx := context.Background()
	r := ResourceCluster().ToResource()
	for _, tc := range []struct {
		name     string
		clear    bool
		wantEdit bool
	}{
		{"no edit when suppressed", false, false},
		{"clearing edit when clear_cloud_attributes_on_remove set", true, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			w := mocks.NewMockWorkspaceClient(t)
			c := testClient(w)
			w.GetMockClustersAPI().EXPECT().
				GetByClusterId(mock.Anything, "abc-123").
				Return(&compute.ClusterDetails{
					ClusterId:     "abc-123",
					ClusterName:   "test",
					SparkVersion:  "13.3.x-scala2.12",
					NodeTypeId:    "i3.xlarge",
					NumWorkers:    2,
					State:         compute.StateRunning,
					AwsAttributes: &compute.AwsAttributes{InstanceProfileArn: testInstanceProfileArn},
				}, nil).Maybe()

			// Only the clearing row sets an Edit expectation. The suppressed row
			// sets none, so gomock fails if the update path calls Edit at all.
			if tc.wantEdit {
				w.GetMockClustersAPI().EXPECT().
					Edit(mock.Anything, mock.MatchedBy(func(e compute.EditCluster) bool {
						return e.AwsAttributes == nil
					})).
					Return(nil, nil).
					Once()
			}

			is := priorStateWithInstanceProfile()
			diff, err := rawConfigDiff(t, r, is, clusterConfigWithoutAws(tc.clear), c)
			if err != nil {
				t.Fatalf("diff: %v", err)
			}

			d, err := schema.InternalMap(r.Schema).Data(is, diff)
			if err != nil {
				t.Fatalf("build resource data: %v", err)
			}

			// Call the update function directly (not r.UpdateContext) to isolate
			// the update behavior from the post-update Read common.Resource runs.
			if err := resourceClusterUpdate(ctx, d, c); err != nil {
				t.Fatalf("update: %v", err)
			}

			if !tc.wantEdit {
				// Nothing was cleared: the instance profile remains in state.
				if got := d.Get("aws_attributes.0.instance_profile_arn"); got != testInstanceProfileArn {
					t.Errorf("instance profile should remain in state; got %q, want %q", got, testInstanceProfileArn)
				}
			}
		})
	}
}

// TestHasClusterConfigChanged verifies that only changes to genuine cluster
// configuration fields trigger an update; changes confined to the ignore list
// (library, is_pinned, no_wait, clear_cloud_attributes_on_remove) do not. Driving the check off an
// explicit diff keeps each case to exactly the fields under test, sidestepping
// schema defaults.
func TestHasClusterConfigChanged(t *testing.T) {
	r := ResourceCluster().ToResource()
	for _, tc := range []struct {
		name    string
		changed map[string]*terraform.ResourceAttrDiff
		want    bool
	}{
		{
			name:    "no changes",
			changed: map[string]*terraform.ResourceAttrDiff{},
			want:    false,
		},
		{
			name: "real config field changed",
			changed: map[string]*terraform.ResourceAttrDiff{
				"spark_version": {Old: "13.3.x-scala2.12", New: "14.0.x-scala2.12"},
			},
			want: true,
		},
		{
			name: "only ignored fields changed",
			changed: map[string]*terraform.ResourceAttrDiff{
				"is_pinned":                        {Old: "false", New: "true"},
				"clear_cloud_attributes_on_remove": {Old: "false", New: "true"},
			},
			want: false,
		},
		{
			name: "ignored and real field changed",
			changed: map[string]*terraform.ResourceAttrDiff{
				"clear_cloud_attributes_on_remove": {Old: "false", New: "true"},
				"node_type_id":                     {Old: "i3.xlarge", New: "i3.2xlarge"},
			},
			want: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			is := &terraform.InstanceState{ID: "abc-123"}
			d, err := schema.InternalMap(r.Schema).Data(is, &terraform.InstanceDiff{Attributes: tc.changed})
			if err != nil {
				t.Fatalf("build resource data: %v", err)
			}
			if got := hasClusterConfigChanged(d); got != tc.want {
				t.Errorf("hasClusterConfigChanged() = %v, want %v", got, tc.want)
			}
		})
	}
}
