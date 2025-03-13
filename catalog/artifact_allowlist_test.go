package catalog_test

import (
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
)

func TestUcAccArtifactAllowlistResourceFullLifecycle(t *testing.T) {
	acceptance.UnityWorkspaceLevel(t, acceptance.Step{
		Template: `
		resource "databricks_artifact_allowlist" "init" {
			artifact_type = "INIT_SCRIPT"
			artifact_matcher {
				artifact = "/Volumes/inits"
				match_type = "PREFIX_MATCH"
			}
		}
		resource "databricks_artifact_allowlist" "maven" {
			artifact_type = "LIBRARY_MAVEN"
			artifact_matcher {
				artifact = "com.databricks:spark-xml"
				match_type = "PREFIX_MATCH"
			}
		}
		resource "databricks_artifact_allowlist" "jar" {
			artifact_type = "LIBRARY_JAR"
			artifact_matcher {
				artifact = "/Volumes/inits"
				match_type = "PREFIX_MATCH"
			}
		}`,
	})
}
