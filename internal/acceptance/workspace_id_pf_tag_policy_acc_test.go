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
// These tests use databricks_tag_policy (PF-only resource using types.Object Namespace via
// GetWorkspaceIDResource / PopulateProviderConfigInState) to validate the unified
// provider plumbing for Plugin Framework resources.

// ==========================================
// State Check Helpers
// ==========================================

// checkTagPolicyProviderConfigWSIDFromEnv verifies provider_config.workspace_id matches the given env var.
// The env var lookup is deferred to check time (after LoadAccountEnv/LoadWorkspaceEnv runs).
func checkTagPolicyProviderConfigWSIDFromEnv(resourceAddr, envVar string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		expected := os.Getenv(envVar)
		if expected == "" {
			return fmt.Errorf("env var %s is not set", envVar)
		}
		return resource.TestCheckResourceAttr(resourceAddr, "provider_config.workspace_id", expected)(s)
	}
}

// ==========================================
// Template Generators
// ==========================================

const tagPolicyResource = "databricks_tag_policy.test"

// tagPolicyTemplate generates HCL for a databricks_tag_policy without a provider block.
func tagPolicyTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_tag_policy" "test" {
		tag_key     = "dwsid-tp-{var.STICKY_RANDOM}"
		description = "workspace_id acceptance test"
		%s
	}
	`, providerConfig)
}

// tagPolicyWithProviderBlock generates HCL for a databricks_tag_policy with an explicit provider block.
func tagPolicyWithProviderBlock(providerAttrs, providerConfig string) string {
	return fmt.Sprintf(`
	provider "databricks" {
		%s
	}
	resource "databricks_tag_policy" "test" {
		tag_key     = "dwsid-tp-{var.STICKY_RANDOM}"
		description = "workspace_id acceptance test"
		%s
	}
	`, providerAttrs, providerConfig)
}

// ==========================================
// Validation Tests
// ==========================================

// TestMwsAccWorkspaceIDTagPolicy_InvalidWorkspaceID tests that
// invalid workspace_id values in the provider block are rejected.
func TestMwsAccWorkspaceIDTagPolicy_InvalidWorkspaceID(t *testing.T) {
	AccountLevel(t, Step{
		Template:    tagPolicyWithProviderBlock(`workspace_id = "invalid"`, ""),
		PlanOnly:    true,
		ExpectError: regexp.MustCompile(`failed to parse workspace_id`),
	})
}

// ==========================================
// Workspace-Level Lifecycle Tests
// ==========================================

// TestAccWorkspaceIDTagPolicy_WS_AddProviderConfig tests adding provider_config to an existing resource.
// Step 1: Create without provider_config. PopulateProviderConfigInState populates state with workspace ID.
// Step 2: Add provider_config with matching workspace_id -> Noop (same value already in state).
func TestAccWorkspaceIDTagPolicy_WS_AddProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tagPolicyTemplate(""),
			Check:    checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceIDTagPolicy_WS_RemoveProviderConfig tests removing provider_config from a resource.
// Step 1: Create with explicit provider_config.
// Step 2: Remove provider_config -> Noop (ProviderConfigPlanModifier preserves state value).
func TestAccWorkspaceIDTagPolicy_WS_RemoveProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tagPolicyTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceIDTagPolicy_WS_ChangeProviderConfig tests changing provider_config.workspace_id to a
// mismatched value on a workspace-level provider.
// Expected: error — workspace_id mismatch.
func TestAccWorkspaceIDTagPolicy_WS_ChangeProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tagPolicyTemplate(`
				provider_config = {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyTemplate(`
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

func TestMwsAccWorkspaceIDTagPolicy_AccountNewSetup(t *testing.T) {
	AccountLevel(t, Step{
		Template: tagPolicyWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			"",
		),
		Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
	})
}

// TestMwsAccWorkspaceIDTagPolicy_AccountNewSetupWithOverride tests that provider_config.workspace_id
// takes precedence over the provider-level workspace_id.
func TestMwsAccWorkspaceIDTagPolicy_AccountNewSetupWithOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: tagPolicyWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			`provider_config = {
				workspace_id = "{env.TEST_WORKSPACE_ID_2}"
			}`,
		),
		Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
	})
}

// ==========================================
// Implicit Default (No Perpetual Diff)
// ==========================================
//
// Account provider + workspace_id, resource has no explicit provider_config.
// Step 1: Create → PopulateProviderConfigInState resolves effective workspace ID.
// Step 2: Same config → Noop (ProviderConfigPlanModifier preserves state value).

func TestMwsAccWorkspaceIDTagPolicy_ImplicitFromProviderDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			// Same config — should be a noop (no perpetual diff).
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
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

func TestMwsAccWorkspaceIDTagPolicy_MigrationSameWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
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

func TestMwsAccWorkspaceIDTagPolicy_MigrationDiffWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
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

func TestMwsAccWorkspaceIDTagPolicy_AddOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
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

func TestMwsAccWorkspaceIDTagPolicy_AddOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
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

func TestMwsAccWorkspaceIDTagPolicy_ChangeOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
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

func TestMwsAccWorkspaceIDTagPolicy_RemoveOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Remove provider_config Override (Falls Back to Different Default)
// ==========================================
//
// Resource has provider_config = { workspace_id = TEST_WORKSPACE_ID_2 }.
// workspace_id = TEST_WORKSPACE_ID (different).
// User removes provider_config → effective workspace changes → ForceNew.
//
// Step 1: Create in TEST_WORKSPACE_ID_2 (override wins over default).
// Step 2: Remove override → falls back to TEST_WORKSPACE_ID → ForceNew.

func TestMwsAccWorkspaceIDTagPolicy_RemoveOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
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

func TestMwsAccWorkspaceIDTagPolicy_ChangeDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID_2"),
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

func TestMwsAccWorkspaceIDTagPolicy_ChangeDefaultWithOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
	)
}

// ==========================================
// Set workspace_id on Workspace-Level Provider
// ==========================================
//
// workspace_id on a workspace-level provider is validated during ModifyPlan.
// If it matches the host's workspace ID, no error. If it doesn't, workspace_id mismatch.

func TestAccWorkspaceIDTagPolicy_DefaultOnWorkspaceProvider_Same(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template: tagPolicyWithProviderBlock(`workspace_id = "{env.THIS_WORKSPACE_ID}"`, ""),
		Check:    checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
	})
}

func TestAccWorkspaceIDTagPolicy_DefaultOnWorkspaceProvider_Diff(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template:    tagPolicyWithProviderBlock(`workspace_id = "12345"`, ""),
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

func TestMwsAccWorkspaceIDTagPolicy_NoDefaultNoOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: tagPolicyWithProviderBlock("", ""),
		PlanOnly: true,
		ExpectError: regexp.MustCompile(
			`(?s)failed to get workspace client`,
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

func TestAccWorkspaceIDTagPolicy_ProviderUpgrade(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tagPolicyTemplate(""),
			Check:    checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tagPolicyResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "THIS_WORKSPACE_ID"),
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

func TestMwsAccWorkspaceIDTagPolicy_RemoveDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock("", ""),
			PlanOnly: true,
			ExpectError: regexp.MustCompile(
				`(?s)provider_config\.workspace_id = \d+ in state but no\s+workspace_id is configured`,
			),
		},
	)
}

// TestMwsAccWorkspaceIDTagPolicy_RemoveOverrideNoFallback tests removing an explicit
// provider_config when there is no provider-level workspace_id to fall back to.
// Step 1: Create with explicit provider_config.workspace_id (no provider-level workspace_id).
// Step 2: Remove provider_config → plan error (workspace_id in state but no source).
func TestMwsAccWorkspaceIDTagPolicy_RemoveOverrideNoFallback(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tagPolicyWithProviderBlock(
				"",
				`provider_config = {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check: checkTagPolicyProviderConfigWSIDFromEnv(tagPolicyResource, "TEST_WORKSPACE_ID"),
		},
		Step{
			Template: tagPolicyWithProviderBlock("", ""),
			PlanOnly: true,
			ExpectError: regexp.MustCompile(
				`(?s)provider_config\.workspace_id = \d+ in state but no\s+workspace_id is configured`,
			),
		},
	)
}
