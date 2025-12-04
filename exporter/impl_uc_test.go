package exporter

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/tags"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/qa"
	"github.com/stretchr/testify/assert"
)

func TestEmitUserSpOrGroup(t *testing.T) {
	ic := importContextForTest()
	ic.enableServices("users,groups")
	emitUserSpOrGroup(ic, "user@example.com")
	assert.Equal(t, 1, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_user[<unknown>] (user_name: user@example.com)")

	emitUserSpOrGroup(ic, "users")
	assert.Equal(t, 2, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users)")

	emitUserSpOrGroup(ic, "abcd1234-ab12-cd34-ef56-abcdef123456")
	assert.Equal(t, 3, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_service_principal[<unknown>] (application_id: abcd1234-ab12-cd34-ef56-abcdef123456)")

	emitUserSpOrGroup(ic, "users @ test.com")
	assert.Equal(t, 4, len(ic.testEmits))
	assert.Contains(t, ic.testEmits, "databricks_group[<unknown>] (display_name: users @ test.com)")

}

func TestTagPolicyExport(t *testing.T) {
	qa.HTTPFixturesApply(t, []qa.HTTPFixture{
		meAdminFixture,
		noCurrentMetastoreAttached,
		{
			Method:   "GET",
			Resource: "/api/2.1/tag-policies?",
			Response: tags.ListTagPoliciesResponse{
				TagPolicies: []tags.TagPolicy{
					{
						TagKey:      "environment",
						Description: "Environment tag policy",
						Values: []tags.Value{
							{Name: "dev"},
							{Name: "staging"},
							{Name: "production"},
						},
					},
					{
						TagKey:      "team",
						Description: "Team tag policy",
						Values: []tags.Value{
							{Name: "engineering"},
							{Name: "data"},
						},
					},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/tag-policies/environment?",
			ReuseRequest: true,
			Response: tags.TagPolicy{
				TagKey:      "environment",
				Description: "Environment tag policy",
				Values: []tags.Value{
					{Name: "dev"},
					{Name: "staging"},
					{Name: "production"},
				},
			},
		},
		{
			Method:       "GET",
			Resource:     "/api/2.1/tag-policies/team?",
			ReuseRequest: true,
			Response: tags.TagPolicy{
				TagKey:      "team",
				Description: "Team tag policy",
				Values: []tags.Value{
					{Name: "engineering"},
					{Name: "data"},
				},
			},
		},
	}, func(ctx context.Context, client *common.DatabricksClient) {
		tmpDir := fmt.Sprintf("/tmp/tf-%s", qa.RandomName())
		defer os.RemoveAll(tmpDir)

		ic := newImportContext(client)
		ic.noFormat = true
		ic.Directory = tmpDir
		ic.enableListing("uc-tags")
		ic.enableServices("uc-tags")

		err := ic.Run()
		assert.NoError(t, err)

		content, err := os.ReadFile(tmpDir + "/uc-tags.tf")
		assert.NoError(t, err)
		contentStr := normalizeWhitespace(string(content))
		assert.Contains(t, contentStr, `resource "databricks_tag_policy" "environment"`)
		assert.Contains(t, contentStr, `resource "databricks_tag_policy" "team"`)
		assert.Contains(t, contentStr, `tag_key = "environment"`)
		assert.Contains(t, contentStr, `description = "Environment tag policy"`)
		assert.Contains(t, contentStr, `name = "dev"`)
		assert.Contains(t, contentStr, `name = "staging"`)
		assert.Contains(t, contentStr, `name = "production"`)
	})
}
