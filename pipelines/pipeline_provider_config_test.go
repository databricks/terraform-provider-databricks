package pipelines_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func pipelineProviderConfigTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_pipeline" "this" {
		name = "test-pipeline"
		library {
			notebook {
				path = "/test"
			}
		}
		%s
	}
	`, providerConfig)
}

func TestAccPipeline_ProviderConfig_Invalid(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: pipelineProviderConfigTemplate(`
			provider_config {
				workspace_id = "invalid"
			}
		`),
		ExpectError: regexp.MustCompile(`workspace_id must be a positive integer without leading zeros`),
		PlanOnly:    true,
	})
}

func TestAccPipeline_ProviderConfig_Required(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: pipelineProviderConfigTemplate(`
			provider_config {
			}
		`),
		ExpectError: regexp.MustCompile(`The argument "workspace_id" is required, but no definition was found.`),
		PlanOnly:    true,
	})
}

func TestAccPipeline_ProviderConfig_EmptyID(t *testing.T) {
	acceptance.WorkspaceLevel(t, acceptance.Step{
		Template: pipelineProviderConfigTemplate(`
			provider_config {
				workspace_id = ""
			}
		`),
		ExpectError: regexp.MustCompile(`expected "provider_config.0.workspace_id" to not be an empty string`),
		PlanOnly:    true,
	})
}
