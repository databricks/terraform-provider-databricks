package sharing_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// shareServer serves a fixed share off the Shares API endpoints. The reported
// objects are deliberately ordered differently from the config to mimic the API
// (which gives no ordering guarantee) and carry populated computed fields.
func shareServer(t *testing.T) *httptest.Server {
	share := sharing.ShareInfo{
		Name:            "my_share",
		Owner:           "me@example.com",
		StorageLocation: "s3://loc",
		UpdatedAt:       123,
		UpdatedBy:       "me@example.com",
		Objects: []sharing.SharedDataObject{
			{Name: "cat.sch.t2", DataObjectType: "TABLE", SharedAs: "sch.t2", HistoryDataSharingStatus: "ENABLED", Status: "ACTIVE", AddedAt: 5, AddedBy: "me@example.com", StartVersion: 1},
			{Name: "cat.sch.t1", DataObjectType: "TABLE", SharedAs: "sch.t1", HistoryDataSharingStatus: "ENABLED", Status: "ACTIVE", AddedAt: 5, AddedBy: "me@example.com", StartVersion: 1},
		},
	}
	writeShare := func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(share)
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/.well-known/databricks-config":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{}`))
		case r.Method == http.MethodPost && r.URL.Path == "/api/2.1/unity-catalog/shares":
			writeShare(w)
		case r.Method == http.MethodPatch && r.URL.Path == "/api/2.1/unity-catalog/shares/my_share":
			writeShare(w)
		case r.Method == http.MethodGet && r.URL.Path == "/api/2.1/unity-catalog/shares/my_share":
			writeShare(w)
		case r.Method == http.MethodDelete && r.URL.Path == "/api/2.1/unity-catalog/shares/my_share":
			w.WriteHeader(http.StatusOK)
		default:
			t.Logf("unhandled %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

func mockHostProviderFactories(host string) map[string]func() (tfprotov6.ProviderServer, error) {
	customizer := func(cfg *config.Config) error {
		cfg.Host = host
		cfg.Token = "token"
		cfg.Credentials = nil
		return nil
	}
	return map[string]func() (tfprotov6.ProviderServer, error){
		"databricks": func() (tfprotov6.ProviderServer, error) {
			ctx := context.Background()
			return providers.GetProviderServer(ctx,
				providers.WithSdkV2Provider(sdkv2.DatabricksProvider(sdkv2.WithConfigCustomizer(customizer))),
				providers.WithPluginFrameworkProvider(pluginfw.GetDatabricksProviderPluginFramework(pluginfw.WithConfigCustomizer(customizer))),
			)
		},
	}
}

// TestShareApplyThenPlanIsEmpty guards against the perpetual-diff regression:
// after apply, a refresh must not drop the synthetic id to null. When it did,
// the next plan reported an out-of-band change (id -> null) and any output
// referencing databricks_share.<name>.id read back null. With PlanOnly and
// ExpectNonEmptyPlan left false, the harness fails the second step if the
// post-apply plan is not empty.
func TestShareApplyThenPlanIsEmpty(t *testing.T) {
	server := shareServer(t)
	defer server.Close()

	hcl := `
resource "databricks_share" "dbx_shares" {
  name = "my_share"

  object {
    name                        = "cat.sch.t1"
    data_object_type            = "TABLE"
    shared_as                   = "sch.t1"
    history_data_sharing_status = "ENABLED"
  }
  object {
    name                        = "cat.sch.t2"
    data_object_type            = "TABLE"
    shared_as                   = "sch.t2"
    history_data_sharing_status = "ENABLED"
  }
}

output "share_id" {
  value = databricks_share.dbx_shares.id
}
`

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: mockHostProviderFactories(server.URL),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("databricks_share.dbx_shares", "id", "my_share"),
					resource.TestCheckResourceAttr("databricks_share.dbx_shares", "name", "my_share"),
					resource.TestCheckOutput("share_id", "my_share"),
				),
			},
			{
				Config:   hcl,
				PlanOnly: true,
			},
		},
	})
}
