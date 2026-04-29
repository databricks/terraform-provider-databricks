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
// UC workspace-level tests require:
//   - THIS_WORKSPACE_ID: numeric ID of the workspace the provider points to
//
// UC account-level tests require:
//   - TEST_WORKSPACE_ID:   primary workspace ID (account must have access, has UC metastore)
//   - TEST_WORKSPACE_ID_2: second workspace ID (for ForceNew / change tests; may lack UC)
//   - TEST_WORKSPACE_URL: deploy URL of the primary workspace (for migration tests)
//
// TEST_WORKSPACE_URL must correspond to TEST_WORKSPACE_ID.
// Account credentials (OAuth) must work against TEST_WORKSPACE_URL directly.
//
// These tests use databricks_share (PF-only resource using types.Object Namespace via
// GetWorkspaceIDResource / PopulateProviderConfigInState) to validate the unified
// provider plumbing for Plugin Framework resources that have UC dependencies.
//
// Two template families are used:
//   - shareTemplate / shareWithProviderBlock:
//     Include managed SDKv2 prereqs (catalog + schema). Use when the provider-level
//     workspace_id stays on TEST_WORKSPACE_ID across ALL steps.
//   - shareWithProviderBlockNoPrereqs:
//     References the pre-existing main.default schema. Use when the provider-level
//     workspace_id changes between steps (SDKv2 resources would fail to refresh on
//     a workspace without UC metastore).

// ==========================================
// State Check Helpers
// ==========================================

// checkShareProviderConfigWSIDFromEnv verifies provider_config.workspace_id matches the given env var.
// The env var lookup is deferred to check time (after LoadUcwsEnv/LoadUcacctEnv runs).
// PF uses provider_config.workspace_id (SingleNestedAttribute / types.Object).
func checkShareProviderConfigWSIDFromEnv(resourceAddr, envVar string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		expected := os.Getenv(envVar)
		if expected == "" {
			return fmt.Errorf("env var %s is not set", envVar)
		}
		return resource.TestCheckResourceAttr(resourceAddr, "provider_config.workspace_id", expected)(s)
	}
}

// ==========================================
// Prerequisites and Template Generators
// ==========================================

// --- WITH managed SDKv2 prereqs ---

// sharePrereqs creates UC resources needed for the share's object block.
// A share must have at least one object to avoid "UpdateShare no fields to update" on Create.
const sharePrereqs = `
	resource "databricks_catalog" "wsid_test" {
		name          = "wsid{var.STICKY_RANDOM}"
		comment       = "catalog for workspace_id acceptance tests"
		force_destroy = true
	}

	resource "databricks_schema" "wsid_test" {
		catalog_name = databricks_catalog.wsid_test.id
		name         = "wsid_schema{var.STICKY_RANDOM}"
	}
`

const shareResource = "databricks_share.test"

// shareTemplate generates HCL for a databricks_share without a provider block.
// Uses managed SDKv2 prereqs — only safe when provider workspace_id == TEST_WORKSPACE_ID.
func shareTemplate(providerConfig string) string {
	return sharePrereqs + fmt.Sprintf(`
	resource "databricks_share" "test" {
		name = "dwsid-share-{var.STICKY_RANDOM}"
		object {
			name             = databricks_schema.wsid_test.id
			data_object_type = "SCHEMA"
		}
		%s
	}
	`, providerConfig)
}

// shareWithProviderBlock generates HCL for a databricks_share with an explicit provider block.
// Uses managed SDKv2 prereqs — only safe when provider workspace_id == TEST_WORKSPACE_ID.
func shareWithProviderBlock(providerAttrs, providerConfig string) string {
	return sharePrereqs + fmt.Sprintf(`
	provider "databricks" {
		%s
	}
	resource "databricks_share" "test" {
		name = "dwsid-share-{var.STICKY_RANDOM}"
		object {
			name             = databricks_schema.wsid_test.id
			data_object_type = "SCHEMA"
		}
		%s
	}
	`, providerAttrs, providerConfig)
}

// --- WITHOUT managed SDKv2 prereqs ---
//
// Reference the pre-existing main.default schema. No SDKv2 resources are created,
// so plan refresh succeeds even when the provider targets a workspace without UC.

// shareWithProviderBlockNoPrereqs generates HCL for a databricks_share with a provider block
// but without managed SDKv2 prereqs.
func shareWithProviderBlockNoPrereqs(providerAttrs, providerConfig string) string {
	return fmt.Sprintf(`
	provider "databricks" {
		%s
	}
	resource "databricks_share" "test" {
		name = "dwsid-share-{var.STICKY_RANDOM}"
		object {
			name             = "main.default"
			data_object_type = "SCHEMA"
		}
		%s
	}
	`, providerAttrs, providerConfig)
}

// ==========================================
// Validation Tests
// ==========================================

// TestUcAccWorkspaceIDShare_InvalidWorkspaceID tests that
// invalid workspace_id values in the provider block are rejected.
func TestUcAccWorkspaceIDShare_InvalidWorkspaceID(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template:    shareWithProviderBlock(`workspace_id = "invalid"`, ""),
		PlanOnly:    true,
		ExpectError: regexp.MustCompile(`failed to parse workspace_id`),
	})
}

// ==========================================
// Workspace-Level Lifecycle Tests
// ==========================================

// TestUcAccWorkspaceIDShare_WS_AddProviderConfig tests adding provider_config to an existing resource.
// Step 1: Create without provider_config. PopulateProviderConfigInState populates state with workspace ID.
// Step 2: Add provider_config with matching workspace_id -> Noop (same value already in state).
func TestUcAccWorkspaceIDShare_WS_AddProviderConfig(t *testing.T) {
	UnityWorkspaceLevel(t,
		Step{
			Template: shareTemplate(""),
			Check:    checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: shareTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestUcAccWorkspaceIDShare_WS_RemoveProviderConfig tests removing provider_config from a resource.
// Step 1: Create with explicit provider_config.
// Step 2: Remove provider_config -> Noop (ProviderConfigPlanModifier preserves state value).
func TestUcAccWorkspaceIDShare_WS_RemoveProviderConfig(t *testing.T) {
	UnityWorkspaceLevel(t,
		Step{
			Template: shareTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: shareTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestUcAccWorkspaceIDShare_WS_ChangeProviderConfig tests changing provider_config.workspace_id to a
// mismatched value on a workspace-level provider.
// Expected: error — workspace_id mismatch.
func TestUcAccWorkspaceIDShare_WS_ChangeProviderConfig(t *testing.T) {
	UnityWorkspaceLevel(t,
		Step{
			Template: shareTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: shareTemplate(`
				provider_config = {
					workspace_id = "123"
				}
			`),
			PlanOnly:    true,
			ExpectError: regexp.MustCompile(`workspace_id mismatch`),
		},
	)
}

// ==========================================
// Account-Level: New Setup
// ==========================================
//
// Account provider + workspace_id → create workspace-level resource.
// Verifies: resource created, state has correct provider_config.workspace_id.

func TestUcAccWorkspaceIDShare_AccountNewSetup(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template: shareWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			"",
		),
		Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
	})
}

// TestUcAccWorkspaceIDShare_AccountNewSetupWithOverride tests that provider_config.workspace_id
// takes precedence over the provider-level workspace_id.
//
// Uses no-prereqs template because the provider targets TEST_WORKSPACE_ID_2 (may lack UC).
// Direction is inverted from the app test (provider WS2, override WS1) so the share
// is created on WS1 (which has UC). The precedence logic is the same: override wins.
func TestUcAccWorkspaceIDShare_AccountNewSetupWithOverride(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template: shareWithProviderBlockNoPrereqs(
			`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
			`provider_config = {
				workspace_id = "{env.TEST_WORKSPACE_ID}"
			}`,
		),
		Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
	})
}

// ==========================================
// Implicit Default (No Perpetual Diff)
// ==========================================
//
// Account provider + workspace_id, resource has no explicit provider_config.
// Step 1: Create → PopulateProviderConfigInState resolves effective workspace ID.
// Step 2: Same config → Noop (ProviderConfigPlanModifier preserves state value).

func TestUcAccWorkspaceIDShare_ImplicitFromProviderDefault(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			// Same config — should be a noop (no perpetual diff).
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
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

func TestUcAccWorkspaceIDShare_MigrationSameWorkspace(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
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
//
// Uses no-prereqs because step 2's provider targets TEST_WORKSPACE_ID_2 (may lack UC),
// causing SDKv2 prereqs to fail refresh. Step 2 is PlanOnly to avoid creating on WS2.

func TestUcAccWorkspaceIDShare_MigrationDiffWorkspace(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
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

func TestUcAccWorkspaceIDShare_AddOverrideSame(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Add provider_config Override (Different Workspace)
// ==========================================
//
// Resource was created using workspace_id = TEST_WORKSPACE_ID.
// User adds provider_config = { workspace_id = TEST_WORKSPACE_ID_2 }.
// Expected: ForceNew (effective workspace changes).
//
// Provider stays WS1 in both steps (SDKv2 prereqs refresh on WS1 → OK).
// Step 2 is PlanOnly because ForceNew would create on WS2 (no UC).

func TestUcAccWorkspaceIDShare_AddOverrideDiff(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
		},
	)
}

// ==========================================
// Change provider_config Override
// ==========================================
//
// Resource has explicit provider_config = { workspace_id = TEST_WORKSPACE_ID }.
// User changes it to TEST_WORKSPACE_ID_2.
// Expected: ForceNew (effective workspace changes).
//
// Provider stays WS1 in both steps (SDKv2 prereqs refresh on WS1 → OK).
// Step 2 is PlanOnly because ForceNew would create on WS2 (no UC).

func TestUcAccWorkspaceIDShare_ChangeOverride(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
		},
	)
}

// ==========================================
// Remove provider_config Override (Falls Back to Same Default)
// ==========================================
//
// Resource has provider_config = { workspace_id = X }. User removes it.
// workspace_id is also X.
// Expected: Noop. ProviderConfigPlanModifier preserves state value, and effective
// workspace is unchanged (default fallback is the same value).

func TestUcAccWorkspaceIDShare_RemoveOverrideSame(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Remove provider_config Override (Falls Back to Different Default)
// ==========================================
//
// Resource has provider_config = { workspace_id = TEST_WORKSPACE_ID }.
// workspace_id = TEST_WORKSPACE_ID_2 (different).
// User removes provider_config → effective workspace changes → ForceNew.
//
// Direction is inverted from the app test: provider WS2, override WS1 → creates on WS1.
// This avoids creating on WS2 (no UC). Step 2 removes override → falls back to WS2 → ForceNew.
//
// Uses no-prereqs because provider targets WS2 (may lack UC for SDKv2 refresh).
// Step 2 is PlanOnly to avoid creating on WS2.

func TestUcAccWorkspaceIDShare_RemoveOverrideDiff(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
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
//
// Uses no-prereqs because step 2's provider targets WS2 (SDKv2 refresh would fail).
// Step 2 is PlanOnly to avoid creating on WS2 (no UC).

func TestUcAccWorkspaceIDShare_ChangeDefault(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			PlanOnly:           true,
			ExpectNonEmptyPlan: true,
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PostApplyPreRefresh: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
		},
	)
}

// ==========================================
// Change workspace_id (Resource Has Explicit Override)
// ==========================================
//
// User changes workspace_id from TEST_WORKSPACE_ID to TEST_WORKSPACE_ID_2.
// Resource has explicit provider_config = { workspace_id = TEST_WORKSPACE_ID }.
// Expected: Noop. The explicit override is the same value in both steps;
// the default changed but the override shields the resource from any diff.
//
// Uses no-prereqs because step 2's provider targets WS2 (SDKv2 refresh would fail).

func TestUcAccWorkspaceIDShare_ChangeDefaultWithOverride(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Set workspace_id on Workspace-Level Provider
// ==========================================
//
// workspace_id on a workspace-level provider is validated during ModifyPlan.
// If it matches the host's workspace ID, no error. If it doesn't, workspace_id mismatch.

func TestUcAccWorkspaceIDShare_DefaultOnWorkspaceProvider_Same(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
		Template: shareWithProviderBlock(`workspace_id = "{env.THIS_WORKSPACE_ID}"`, ""),
		Check:    checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
	})
}

func TestUcAccWorkspaceIDShare_DefaultOnWorkspaceProvider_Diff(t *testing.T) {
	UnityWorkspaceLevel(t, Step{
		Template:    shareWithProviderBlock(`workspace_id = "12345"`, ""),
		PlanOnly:    true,
		ExpectError: regexp.MustCompile(`workspace_id mismatch`),
	})
}

// ==========================================
// Account Provider Without workspace_id or provider_config
// ==========================================
//
// Account-level provider with no workspace_id. Resource has no provider_config.
// Expected: error during plan — no workspace_id available for routing.

func TestUcAccWorkspaceIDShare_NoDefaultNoOverride(t *testing.T) {
	UnityAccountLevel(t, Step{
		Template: shareWithProviderBlock("", ""),
		PlanOnly: true,
		// shareWithProviderBlock includes SDKv2 prereqs (catalog + schema)
		// whose CustomizeDiff (NamespaceCustomizeDiffNoForceNew) now validates
		// workspace_id during plan and surfaces the SDKv2 error before the PF
		// share's ValidateWorkspaceID would. The SDKv2 error is returned plain
		// (not wrapped in the PF "failed to get workspace client" diagnostic).
		ExpectError: regexp.MustCompile(
			`(?s)managing workspace-level resources requires a workspace_id`,
		),
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

func TestUcAccWorkspaceIDShare_ProviderUpgrade(t *testing.T) {
	UnityWorkspaceLevel(t,
		Step{
			Template: shareTemplate(""),
			Check:    checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: shareTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(shareResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "THIS_WORKSPACE_ID"),
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
//
// Uses no-prereqs because step 2's provider has no workspace_id (SDKv2 resources
// in state would fail to refresh without a workspace client).

func TestUcAccWorkspaceIDShare_RemoveDefault(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs("", ""),
			PlanOnly: true,
			ExpectError: regexp.MustCompile(
				`(?s)provider_config\.workspace_id = \d+ in state but no\s+workspace_id is configured`,
			),
		},
	)
}

// ==========================================
// Remove provider_config Override (No Fallback)
// ==========================================
//
// Step 1: Create with explicit provider_config.workspace_id (no provider-level workspace_id).
// Step 2: Remove provider_config → plan error (workspace_id in state but no source).
//
// Uses no-prereqs because step 2's provider has no workspace_id (SDKv2 resources
// in state would fail to refresh without a workspace client).

func TestUcAccWorkspaceIDShare_RemoveOverrideNoFallback(t *testing.T) {
	UnityAccountLevel(t,
		Step{
			Template: shareWithProviderBlockNoPrereqs(
				"",
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkShareProviderConfigWSIDFromEnv(shareResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: shareWithProviderBlockNoPrereqs("", ""),
			PlanOnly: true,
			ExpectError: regexp.MustCompile(
				`(?s)provider_config\.workspace_id = \d+ in state but no\s+workspace_id is configured`,
			),
		},
	)
}
