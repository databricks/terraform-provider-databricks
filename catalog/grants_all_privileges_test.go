package catalog

import (
	"slices"
	"testing"

	"github.com/databricks/databricks-sdk-go/experimental/mocks"
	"github.com/databricks/databricks-sdk-go/service/catalog"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGrantDiffRevokingAllPrivileges(t *testing.T) {
	tests := []struct {
		name     string
		existing []catalog.Privilege
		desired  []catalog.Privilege
		add      []catalog.Privilege
		remove   []catalog.Privilege
	}{
		{"catalog retains privilege", []catalog.Privilege{"ALL_PRIVILEGES", "USE_CATALOG"}, []catalog.Privilege{"SELECT", "USE_CATALOG"}, []catalog.Privilege{"SELECT", "USE_CATALOG"}, []catalog.Privilege{"ALL_PRIVILEGES"}},
		{"schema retains privilege", []catalog.Privilege{"all privileges", "USE_SCHEMA"}, []catalog.Privilege{"SELECT", "USE_SCHEMA"}, []catalog.Privilege{"SELECT", "USE_SCHEMA"}, []catalog.Privilege{"ALL_PRIVILEGES"}},
		{"only remove all", []catalog.Privilege{"ALL_PRIVILEGES", "USE_CATALOG"}, []catalog.Privilege{"USE_CATALOG"}, []catalog.Privilege{"USE_CATALOG"}, []catalog.Privilege{"ALL_PRIVILEGES"}},
		{"retain all", []catalog.Privilege{"ALL_PRIVILEGES", "USE_CATALOG"}, []catalog.Privilege{"ALL_PRIVILEGES", "SELECT"}, []catalog.Privilege{"SELECT"}, []catalog.Privilege{"USE_CATALOG"}},
		{"retain unaffected privilege", []catalog.Privilege{"ALL_PRIVILEGES", "MANAGE"}, []catalog.Privilege{"MANAGE"}, []catalog.Privilege{"MANAGE"}, []catalog.Privilege{"ALL_PRIVILEGES"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			desired := []catalog.PrivilegeAssignment{{Principal: "target", Privileges: tt.desired}, {Principal: "other", Privileges: []catalog.Privilege{"SELECT"}}}
			existing := []catalog.PrivilegeAssignment{{Principal: "target", Privileges: tt.existing}, {Principal: "other", Privileges: []catalog.Privilege{"MODIFY"}}}
			for name, diff := range map[string][]catalog.PermissionsChange{
				"grants": diffPermissions(desired, existing),
				"grant":  diffPermissionsForPrincipal("target", desired, existing),
			} {
				t.Run(name, func(t *testing.T) {
					assert.Equal(t, "target", diff[len(diff)-1].Principal)
					assert.ElementsMatch(t, tt.add, diff[len(diff)-1].Add)
					assert.ElementsMatch(t, tt.remove, diff[len(diff)-1].Remove)
					if name == "grants" {
						assert.Len(t, diff, 2)
						assert.ElementsMatch(t, []catalog.Privilege{"SELECT"}, diff[0].Add)
						assert.ElementsMatch(t, []catalog.Privilege{"MODIFY"}, diff[0].Remove)
					} else {
						assert.Len(t, diff, 1)
					}
				})
			}
		})
	}
}

// mockRevokeAllPrivileges expects the update of catalog `sandbox` that replaces ALL_PRIVILEGES
// and USE_CATALOG of `account users` with SELECT and USE_CATALOG. The update must grant
// USE_CATALOG again, because revoking ALL_PRIVILEGES also revokes the privileges it implies.
func mockRevokeAllPrivileges(w *mocks.MockWorkspaceClient) {
	grants := func(privileges ...catalog.Privilege) *catalog.GetPermissionsResponse {
		return &catalog.GetPermissionsResponse{
			PrivilegeAssignments: []catalog.PrivilegeAssignment{
				{Principal: "account users", Privileges: privileges},
			},
		}
	}
	e := w.GetMockGrantsAPI().EXPECT()
	e.GetBySecurableTypeAndFullName(mock.Anything, "catalog", "sandbox").
		Return(grants("ALL_PRIVILEGES", "USE_CATALOG"), nil).Once()
	e.Update(mock.Anything, mock.MatchedBy(func(r catalog.UpdatePermissions) bool {
		if r.SecurableType != "catalog" || r.FullName != "sandbox" || len(r.Changes) != 1 {
			return false
		}
		c := r.Changes[0]
		slices.Sort(c.Add)
		return c.Principal == "account users" &&
			slices.Equal(c.Add, []catalog.Privilege{"SELECT", "USE_CATALOG"}) &&
			slices.Equal(c.Remove, []catalog.Privilege{"ALL_PRIVILEGES"})
	})).Return(nil, nil).Once()
	// read by the wait for the update and by the read after it
	e.GetBySecurableTypeAndFullName(mock.Anything, "catalog", "sandbox").
		Return(grants("SELECT", "USE_CATALOG"), nil)
}

func TestGrantsUpdateRevokingAllPrivileges(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: mockRevokeAllPrivileges,
		Resource:                ResourceGrants(),
		Update:                  true,
		ID:                      "catalog/sandbox",
		InstanceState: map[string]string{
			"catalog": "sandbox",
		},
		HCL: `
		catalog = "sandbox"

		grant {
			principal = "account users"
			privileges = ["SELECT", "USE_CATALOG"]
		}
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"catalog": "sandbox",
		"grant.#": 1,
	})
}

func TestResourceGrantUpdateRevokingAllPrivileges(t *testing.T) {
	qa.ResourceFixture{
		MockWorkspaceClientFunc: mockRevokeAllPrivileges,
		Resource:                ResourceGrant(),
		Update:                  true,
		ID:                      "catalog/sandbox/account users",
		InstanceState: map[string]string{
			"catalog":   "sandbox",
			"principal": "account users",
		},
		HCL: `
		catalog = "sandbox"

		principal = "account users"
		privileges = ["SELECT", "USE_CATALOG"]
		`,
	}.ApplyAndExpectData(t, map[string]any{
		"catalog":      "sandbox",
		"principal":    "account users",
		"privileges.#": 2,
	})
}
