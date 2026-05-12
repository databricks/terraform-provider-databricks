package acceptance

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// ==========================================
// Environment Variables
// ==========================================
//
// Workspace-level tests require:
//   - THIS_WORKSPACE_ID: numeric ID of the workspace the provider points to
//
// Account-level tests require:
//   - TEST_WORKSPACE_ID:   primary workspace ID (account must have access)
//   - TEST_WORKSPACE_ID_2: second workspace ID (for ForceNew / change tests)
//   - TEST_WORKSPACE_URL: deploy URL of the primary workspace (for migration tests)
//
// TEST_WORKSPACE_URL must correspond to TEST_WORKSPACE_ID.
// Account credentials (OAuth) must work against TEST_WORKSPACE_URL directly.
//
// These tests use databricks_directory (SDKv2 resource using Go SDK via
// WorkspaceClientUnifiedProvider) to validate the unified provider plumbing.
// It has no dependencies and is fast to create.

// ==========================================
// State Check Helpers
// ==========================================

// checkProviderConfigWSIDFromEnv verifies provider_config.0.workspace_id matches the given env var.
// The env var lookup is deferred to check time (after LoadAccountEnv/LoadWorkspaceEnv runs).
func checkProviderConfigWSIDFromEnv(resourceAddr, envVar string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		expected := os.Getenv(envVar)
		if expected == "" {
			return fmt.Errorf("env var %s is not set", envVar)
		}
		return resource.TestCheckResourceAttr(resourceAddr, "provider_config.0.workspace_id", expected)(s)
	}
}

// ==========================================
// Template Generators
// ==========================================

const directoryResource = "databricks_directory.test"

// directoryTemplate generates HCL for a databricks_directory without a provider block.
func directoryTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_directory" "test" {
		path = "/Shared/dwsid-test-{var.STICKY_RANDOM}"
		%s
	}
	`, providerConfig)
}

// directoryWithProviderBlock generates HCL for a databricks_directory with an explicit provider block.
func directoryWithProviderBlock(providerAttrs, providerConfig string) string {
	return fmt.Sprintf(`
	provider "databricks" {
		%s
	}
	resource "databricks_directory" "test" {
		path = "/Shared/dwsid-test-{var.STICKY_RANDOM}"
		%s
	}
	`, providerAttrs, providerConfig)
}

// ==========================================
// Validation Tests
// ==========================================

// TestAccWorkspaceID_InvalidWorkspaceID tests that
// invalid workspace_id values in the provider block are rejected.
// The SDK's workspace_id field does not have schema-level validation,
// so invalid values are caught during workspace client creation.
func TestMwsAccWorkspaceID_InvalidWorkspaceID(t *testing.T) {
	AccountLevel(t, Step{
		Template:    directoryWithProviderBlock(`workspace_id = "invalid"`, ""),
		ExpectError: regexp.MustCompile(`failed to parse workspace_id`),
		PlanOnly:    true,
	})
}

// ==========================================
// Workspace-Level Lifecycle Tests
// ==========================================

// TestAccWorkspaceID_WS_AddProviderConfig tests adding provider_config to an existing resource.
// Step 1: Create without provider_config. Post-Read hook populates state with workspace ID.
// Step 2: Add provider_config with matching workspace_id -> Noop (same value already in state).
func TestAccWorkspaceID_WS_AddProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: directoryTemplate(""),
			Check:    checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: directoryTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceID_WS_RemoveProviderConfig tests removing provider_config from a resource.
// Step 1: Create with provider_config.
// Step 2: Remove provider_config -> Noop. Optional+Computed preserves state value,
// and effective workspace is unchanged (same workspace).
func TestAccWorkspaceID_WS_RemoveProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: directoryTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: directoryTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceID_WS_ChangeProviderConfig tests that changing workspace_id
// to a different workspace on a workspace-level provider produces an error.
// Step 1: Create with workspace_id matching current workspace.
// Step 2: Change to different workspace_id -> error (workspace-level provider
// cannot target a different workspace).
func TestAccWorkspaceID_WS_ChangeProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: directoryTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: directoryTemplate(`
				provider_config {
					workspace_id = "123"
				}
			`),
			ExpectError: regexp.MustCompile(`workspace_id mismatch`),
			PlanOnly:    true,
		},
	)
}

// ==========================================
// Account-Level: New Setup
// ==========================================
//
// Account provider + workspace_id → create workspace-level resource.
// Verifies: resource created, state has correct provider_config.0.workspace_id.

func TestMwsAccWorkspaceID_AccountNewSetup(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			"",
		),
		Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
	})
}

// TestAccWorkspaceID_AccountNewSetupWithOverride tests that provider_config.workspace_id
// takes precedence over the provider-level workspace_id.
func TestMwsAccWorkspaceID_AccountNewSetupWithOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			`provider_config {
				workspace_id = "{env.TEST_WORKSPACE_ID_2}"
			}`,
		),
		Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
	})
}

// ==========================================
// Implicit From Provider Default
// ==========================================
//
// Account provider with workspace_id. Resource has NO provider_config.
// Provider's workspace_id is used implicitly.
// Step 2 re-applies the same config to verify no perpetual diff (noop).

func TestMwsAccWorkspaceID_ImplicitFromProviderDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			// Same config — should be a noop (no perpetual diff).
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Migration: Same Workspace
// ==========================================
//
// Existing resource on workspace-level provider. Switch to account provider
// with workspace_id pointing to the SAME workspace.
// Expected: no changes (noop).
//
// Step 1: provider { host = TEST_WORKSPACE_URL } → workspace-level create.
// Step 2: provider { workspace_id = TEST_WORKSPACE_ID } → account-level, same ws.
//
// Requires: account-level OAuth credentials that also work against the workspace host.

func TestMwsAccWorkspaceID_MigrationSameWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Migration: Different Workspace
// ==========================================
//
// Same as MigrationSameWorkspace, but workspace_id points to a DIFFERENT workspace.
// Expected: ForceNew (destroy in old workspace, recreate in new).
//
// Step 1: provider { host = TEST_WORKSPACE_URL } → create in workspace 1.
// Step 2: provider { workspace_id = TEST_WORKSPACE_ID_2 } → ForceNew.

func TestMwsAccWorkspaceID_MigrationDiffWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		},
	)
}

// ==========================================
// Add provider_config Override (Same Workspace)
// ==========================================
//
// Resource was created using workspace_id. User adds explicit
// provider_config with the SAME workspace ID.
// Expected: Noop. State already has the same value from the default; no diff.

func TestMwsAccWorkspaceID_AddOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Add provider_config Override (Different Workspace)
// ==========================================
//
// Resource was created using workspace_id = TEST_WORKSPACE_ID.
// User adds provider_config { workspace_id = TEST_WORKSPACE_ID_2 }.
// Expected: ForceNew (effective workspace changes).

func TestMwsAccWorkspaceID_AddOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		},
	)
}

// ==========================================
// Change provider_config Override
// ==========================================
//
// Resource has provider_config { workspace_id = X }. User changes it to Y.
// Expected: ForceNew (effective workspace changes).

func TestMwsAccWorkspaceID_ChangeOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		},
	)
}

// ==========================================
// Remove provider_config Override (Falls Back to Same Default)
// ==========================================
//
// Resource has provider_config { workspace_id = X }. User removes it.
// workspace_id is also X.
// Expected: Noop. Optional+Computed preserves state value, and effective
// workspace is unchanged (default fallback is the same value).

func TestMwsAccWorkspaceID_RemoveOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Remove provider_config Override (Falls Back to Different Default)
// ==========================================
//
// Resource has provider_config { workspace_id = TEST_WORKSPACE_ID_2 }.
// workspace_id = TEST_WORKSPACE_ID (different).
// User removes provider_config → effective workspace changes → ForceNew.
//
// Step 1: Create in TEST_WORKSPACE_ID_2 (override wins over default).
// Step 2: Remove override → falls back to TEST_WORKSPACE_ID → ForceNew.

func TestMwsAccWorkspaceID_RemoveOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Change workspace_id (No provider_config)
// ==========================================
//
// User changes workspace_id from TEST_WORKSPACE_ID to TEST_WORKSPACE_ID_2.
// Resources have no explicit provider_config.
// Expected: ForceNew for all affected resources.

func TestMwsAccWorkspaceID_ChangeDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		},
	)
}

// ==========================================
// Change workspace_id (Resource Has Explicit Override)
// ==========================================
//
// User changes workspace_id from TEST_WORKSPACE_ID to TEST_WORKSPACE_ID_2.
// Resource has explicit provider_config { workspace_id = TEST_WORKSPACE_ID }.
// Expected: Noop. The explicit override is the same value in both steps;
// the default changed but the override shields the resource from any diff.

func TestMwsAccWorkspaceID_ChangeDefaultWithOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Set workspace_id on Workspace-Level Provider
// ==========================================
//
// workspace_id on a workspace-level provider is validated during CustomizeDiff.
// If it matches the host's workspace ID, no error. If it doesn't, workspace_id mismatch.

func TestAccWorkspaceID_DefaultOnWorkspaceProvider_Same(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: directoryWithProviderBlock(
			fmt.Sprintf(`workspace_id = "%s"`, os.Getenv("THIS_WORKSPACE_ID")),
			"",
		),
		Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
	})
}

func TestAccWorkspaceID_DefaultOnWorkspaceProvider_Diff(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template:    directoryWithProviderBlock(`workspace_id = "12345"`, ""),
		PlanOnly:    true,
		ExpectError: regexp.MustCompile(`workspace_id mismatch`),
	})
}

// ==========================================
// Account Provider Without workspace_id or provider_config
// ==========================================
//
// Account-level provider with no workspace_id. Resource has no provider_config.
// Expected: error during CRUD — no workspace_id available for routing.

func TestMwsAccWorkspaceID_NoDefaultNoOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock("", ""),
		ExpectError: regexp.MustCompile(
			`managing workspace-level resources requires a workspace_id, but none was found in the resource's provider_config block or the provider's workspace_id attribute`,
		),
		PlanOnly: true,
	})
}

// ==========================================
// Provider Upgrade (Existing Resources, No Config Changes)
// ==========================================
//
// Approximation: create a resource at workspace level without any unified
// provider features (no provider_config, no workspace_id). Then verify
// a second plan/apply with the same config produces no changes.
// This validates backward compatibility — the new provider_config schema
// field doesn't cause unexpected diffs on existing resources.

func TestAccWorkspaceID_ProviderUpgrade(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: directoryTemplate(""),
			Check:    checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: directoryTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(directoryResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Remove workspace_id from Provider Config
// ==========================================
//
// User had workspace_id and resources relying on it (no explicit provider_config).
// User removes workspace_id.
// Expected: error — provider_config.workspace_id is in state but no source for it
// in config and no workspace_id set.
//
// Step 1: Create with workspace_id.
// Step 2: Remove workspace_id → plan error.

func TestMwsAccWorkspaceID_RemoveDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock("", ""),
			ExpectError: regexp.MustCompile(
				`resource has provider_config.workspace_id = .* in state, but managing workspace-level resources requires a workspace_id`,
			),
			PlanOnly: true,
		},
	)
}

// ==========================================
// Remove provider_config Override (No Provider Fallback)
// ==========================================
//
// Resource has explicit provider_config { workspace_id = TEST_WORKSPACE_ID },
// but provider has NO workspace_id set.
// User removes provider_config → no fallback → plan error.
//
// Complements RemoveDefault (which tests removing the provider-level workspace_id
// when resources have no explicit provider_config).

func TestMwsAccWorkspaceID_RemoveOverrideNoFallback(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				"",
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: directoryWithProviderBlock("", ""),
			ExpectError: regexp.MustCompile(
				`resource has provider_config.workspace_id = .* in state, but managing workspace-level resources requires a workspace_id`,
			),
			PlanOnly: true,
		},
	)
}
