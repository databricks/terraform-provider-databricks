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
// These tests use databricks_directory (Go SDK resource via WorkspaceClientUnifiedProvider)
// to validate the unified provider plumbing. It has no dependencies and is fast to create.

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
func TestAccWorkspaceID_InvalidWorkspaceID(t *testing.T) {
	AccountLevel(t, Step{
		Template:    directoryWithProviderBlock(`workspace_id = "invalid"`, ""),
		ExpectError: regexp.MustCompile(`failed to parse workspace_id`),
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
		},
	)
}

// ==========================================
// CUJ 1: New unified provider setup
// ==========================================
//
// Account provider + workspace_id → create workspace-level resource.
// Verifies: resource created, state has correct provider_config.0.workspace_id.

func TestAccWorkspaceID_CUJ1_NewSetup(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			"",
		),
		Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
	})
}

// TestAccWorkspaceID_CUJ1_NewSetupWithOverride tests the second half of CUJ 1:
// one resource uses the default, another overrides with provider_config.
// We test the override path: provider_config.workspace_id takes precedence.
func TestAccWorkspaceID_CUJ1_NewSetupWithOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			`provider_config {
				workspace_id = "{env.TEST_WORKSPACE_ID_2}"
			}`,
		),
		Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
	})
}

// ==========================================
// CUJ 2: Migration — same workspace
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

func TestAccWorkspaceID_CUJ2_MigrationSameWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 3: Migration — different workspace
// ==========================================
//
// Same as CUJ 2, but workspace_id points to a DIFFERENT workspace.
// Expected: ForceNew (destroy in old workspace, recreate in new).
//
// Step 1: provider { host = TEST_WORKSPACE_URL } → create in workspace 1.
// Step 2: provider { workspace_id = TEST_WORKSPACE_ID_2 } → ForceNew.

func TestAccWorkspaceID_CUJ3_MigrationDiffWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 4: Add provider_config override (same workspace)
// ==========================================
//
// Resource was created using workspace_id. User adds explicit
// provider_config with the SAME workspace ID.
// Expected: Noop. State already has the same value from the default; no diff.

func TestAccWorkspaceID_CUJ4_AddOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 5: Add provider_config override (different workspace)
// ==========================================
//
// Resource was created using workspace_id = TEST_WORKSPACE_ID.
// User adds provider_config { workspace_id = TEST_WORKSPACE_ID_2 }.
// Expected: ForceNew (effective workspace changes).

func TestAccWorkspaceID_CUJ5_AddOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 6: Remove provider_config override (falls back to same default)
// ==========================================
//
// Resource has provider_config { workspace_id = X }. User removes it.
// workspace_id is also X.
// Expected: Noop. Optional+Computed preserves state value, and effective
// workspace is unchanged (default fallback is the same value).

func TestAccWorkspaceID_CUJ6_RemoveOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 7: Remove provider_config override (falls back to different default)
// ==========================================
//
// Resource has provider_config { workspace_id = TEST_WORKSPACE_ID_2 }.
// workspace_id = TEST_WORKSPACE_ID (different).
// User removes provider_config → effective workspace changes → ForceNew.
//
// Step 1: Create in TEST_WORKSPACE_ID_2 (override wins over default).
// Step 2: Remove override → falls back to TEST_WORKSPACE_ID → ForceNew.

func TestAccWorkspaceID_CUJ7_RemoveOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 8: Change workspace_id (no provider_config)
// ==========================================
//
// User changes workspace_id from TEST_WORKSPACE_ID to TEST_WORKSPACE_ID_2.
// Resources have no explicit provider_config.
// Expected: ForceNew for all affected resources.

func TestAccWorkspaceID_CUJ8_ChangeDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 9: Change workspace_id (resource has explicit override)
// ==========================================
//
// User changes workspace_id from TEST_WORKSPACE_ID to TEST_WORKSPACE_ID_2.
// Resource has explicit provider_config { workspace_id = TEST_WORKSPACE_ID }.
// Expected: Noop. The explicit override is the same value in both steps;
// the default changed but the override shields the resource from any diff.

func TestAccWorkspaceID_CUJ9_ChangeDefaultWithOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 10: Set workspace_id on workspace-level provider
// ==========================================
//
// User accidentally sets workspace_id on a workspace-level provider.
// Expected: configuration error at provider initialization.

func TestAccWorkspaceID_CUJ10_DefaultOnWorkspaceProvider(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template:    directoryWithProviderBlock(`workspace_id = "12345"`, ""),
		ExpectError: regexp.MustCompile(`workspace_id cannot be used with a workspace-level provider; it is only supported when the provider is configured at the account level`),
		PlanOnly:    true,
	})
}

// ==========================================
// CUJ 11: Account provider without workspace_id or provider_config
// ==========================================
//
// Account-level provider with no workspace_id. Resource has no provider_config.
// Expected: error during CRUD — no workspace_id available for routing.

func TestAccWorkspaceID_CUJ11_NoDefaultNoOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: directoryWithProviderBlock("", ""),
		ExpectError: regexp.MustCompile(
			`managing workspace-level resources requires a workspace_id, but none was found in provider_config or the provider configuration`,
		),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
	})
}

// ==========================================
// CUJ 13: Provider upgrade (existing resources, no config changes)
// ==========================================
//
// Approximation: create a resource at workspace level without any unified
// provider features (no provider_config, no workspace_id). Then verify
// a second plan/apply with the same config produces no changes.
// This validates backward compatibility — the new provider_config schema
// field doesn't cause unexpected diffs on existing resources.

func TestAccWorkspaceID_CUJ13_ProviderUpgrade(t *testing.T) {
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
// CUJ 14: Remove workspace_id from provider config
// ==========================================
//
// User had workspace_id and resources relying on it (no explicit provider_config).
// User removes workspace_id.
// Expected: error — provider_config.workspace_id is in state but no source for it
// in config and no workspace_id set.
//
// Step 1: Create with workspace_id.
// Step 2: Remove workspace_id → plan error.

func TestAccWorkspaceID_CUJ14_RemoveDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: directoryWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkProviderConfigWSIDFromEnv(directoryResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: directoryWithProviderBlock("", ""),
			ExpectError: regexp.MustCompile(
				`resource has provider_config.workspace_id = .* in state, but managing workspace-level resources requires a workspace_id`,
			),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// CUJ 15: Provider downgrade
// ==========================================
//
// Cannot be tested in acceptance tests. Provider downgrade to a version without
// workspace_id support relies on standard Terraform behavior: unknown
// attributes in state are silently dropped when the schema no longer includes them.
// This is Terraform core behavior, not provider logic.
