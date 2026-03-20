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
// Environment Variables (same as workspace_id_sdkv2_gosdk_acc_test.go)
// ==========================================
//
// Workspace-level tests require:
//   - THIS_WORKSPACE_ID: numeric ID of the workspace the provider points to
//
// Account-level tests require:
//   - TEST_WORKSPACE_ID:   primary workspace ID (account must have access)
//   - TEST_WORKSPACE_ID_2: second workspace ID (for ForceNew / change tests)
//   - TEST_WORKSPACE_URL:  deploy URL of the primary workspace (for migration tests)
//
// These tests use databricks_token (SDKv2 resource using HTTP paths via
// DatabricksClientForUnifiedProvider, NOT Go SDK) to validate the unified
// provider plumbing for non-Go-SDK resources.
// It has no dependencies and is fast to create.

// ==========================================
// State Check Helpers
// ==========================================

// checkTokenProviderConfigWSIDFromEnv verifies provider_config.0.workspace_id matches the given env var.
func checkTokenProviderConfigWSIDFromEnv(resourceAddr, envVar string) func(*terraform.State) error {
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

const tokenResource = "databricks_token.test"

// tokenTemplate generates HCL for a databricks_token without a provider block.
func tokenTemplate(providerConfig string) string {
	return fmt.Sprintf(`
	resource "databricks_token" "test" {
		comment          = "dwsid-test-{var.STICKY_RANDOM}"
		lifetime_seconds = 3600
		%s
	}
	`, providerConfig)
}

// tokenWithProviderBlock generates HCL for a databricks_token with an explicit provider block.
func tokenWithProviderBlock(providerAttrs, providerConfig string) string {
	return fmt.Sprintf(`
	provider "databricks" {
		%s
	}
	resource "databricks_token" "test" {
		comment          = "dwsid-test-{var.STICKY_RANDOM}"
		lifetime_seconds = 3600
		%s
	}
	`, providerAttrs, providerConfig)
}

// ==========================================
// Validation Tests
// ==========================================

// TestAccWorkspaceIDHttp_InvalidWorkspaceID tests that
// invalid workspace_id values in the provider block are rejected.
func TestAccWorkspaceIDHttp_InvalidWorkspaceID(t *testing.T) {
	AccountLevel(t, Step{
		Template:    tokenWithProviderBlock(`workspace_id = "invalid"`, ""),
		ExpectError: regexp.MustCompile(`failed to parse workspace_id`),
	})
}

// ==========================================
// Workspace-Level Lifecycle Tests
// ==========================================

// TestAccWorkspaceIDHttp_WS_AddProviderConfig tests adding provider_config to an existing resource.
// Step 1: Create without provider_config. Post-Read hook populates state with workspace ID.
// Step 2: Add provider_config with matching workspace_id -> Noop (same value already in state).
func TestAccWorkspaceIDHttp_WS_AddProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tokenTemplate(""),
			Check:    checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tokenTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceIDHttp_WS_RemoveProviderConfig tests removing provider_config from a resource.
// Step 1: Create with provider_config.
// Step 2: Remove provider_config -> Noop. Optional+Computed preserves state value,
// and effective workspace is unchanged (same workspace).
func TestAccWorkspaceIDHttp_WS_RemoveProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tokenTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tokenTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
	)
}

// TestAccWorkspaceIDHttp_WS_ChangeProviderConfig tests that changing workspace_id
// to a different workspace on a workspace-level provider produces an error.
// Step 1: Create with workspace_id matching current workspace.
// Step 2: Change to different workspace_id -> error (workspace-level provider
// cannot target a different workspace).
func TestAccWorkspaceIDHttp_WS_ChangeProviderConfig(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tokenTemplate(`
				provider_config {
					workspace_id = "{env.THIS_WORKSPACE_ID}"
				}
			`),
			Check: checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tokenTemplate(`
				provider_config {
					workspace_id = "123"
				}
			`),
			ExpectError: regexp.MustCompile(`workspace_id mismatch`),
		},
	)
}

// ==========================================
// Account-Level: New Setup
// ==========================================
//
// Account provider + workspace_id → create workspace-level resource.
// Verifies: resource created, state has correct provider_config.0.workspace_id.

func TestAccWorkspaceIDHttp_AccountNewSetup(t *testing.T) {
	AccountLevel(t, Step{
		Template: tokenWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			"",
		),
		Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
	})
}

// TestAccWorkspaceIDHttp_AccountNewSetupWithOverride tests that provider_config.workspace_id
// takes precedence over the provider-level workspace_id.
func TestAccWorkspaceIDHttp_AccountNewSetupWithOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: tokenWithProviderBlock(
			`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
			`provider_config {
				workspace_id = "{env.TEST_WORKSPACE_ID_2}"
			}`,
		),
		Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID_2"),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
	})
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

func TestAccWorkspaceIDHttp_MigrationSameWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_MigrationDiffWorkspace(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`host = "{env.TEST_WORKSPACE_URL}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_AddOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_AddOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_RemoveOverrideSame(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_RemoveOverrideDiff(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID_2}"
				}`,
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_ChangeDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				"",
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionDestroyBeforeCreate),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID_2"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_ChangeDefaultWithOverride(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID_2}"`,
				`provider_config {
					workspace_id = "{env.TEST_WORKSPACE_ID}"
				}`,
			),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}

// ==========================================
// Set workspace_id on Workspace-Level Provider
// ==========================================
//
// User accidentally sets workspace_id on a workspace-level provider.
// Expected: configuration error at provider initialization.

func TestAccWorkspaceIDHttp_DefaultOnWorkspaceProvider(t *testing.T) {
	WorkspaceLevel(t, Step{
		Template:    tokenWithProviderBlock(`workspace_id = "12345"`, ""),
		ExpectError: regexp.MustCompile(`workspace_id cannot be used with a workspace-level provider; it is only supported when the provider is configured at the account level`),
		PlanOnly:    true,
	})
}

// ==========================================
// Account Provider Without workspace_id or provider_config
// ==========================================
//
// Account-level provider with no workspace_id. Resource has no provider_config.
// Expected: error during CRUD — no workspace_id available for routing.

func TestAccWorkspaceIDHttp_NoDefaultNoOverride(t *testing.T) {
	AccountLevel(t, Step{
		Template: tokenWithProviderBlock("", ""),
		ExpectError: regexp.MustCompile(
			`managing workspace-level resources requires a workspace_id, but none was found in provider_config or the provider configuration`,
		),
		ProtoV6ProviderFactories: noOidcProviderFactories(),
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

func TestAccWorkspaceIDHttp_ProviderUpgrade(t *testing.T) {
	WorkspaceLevel(t,
		Step{
			Template: tokenTemplate(""),
			Check:    checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
		},
		Step{
			Template: tokenTemplate(""),
			ConfigPlanChecks: resource.ConfigPlanChecks{
				PreApply: []plancheck.PlanCheck{
					plancheck.ExpectResourceAction(tokenResource, plancheck.ResourceActionNoop),
				},
			},
			Check: checkTokenProviderConfigWSIDFromEnv(tokenResource, "THIS_WORKSPACE_ID"),
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

func TestAccWorkspaceIDHttp_RemoveDefault(t *testing.T) {
	AccountLevel(t,
		Step{
			Template: tokenWithProviderBlock(
				`workspace_id = "{env.TEST_WORKSPACE_ID}"`,
				"",
			),
			Check:                    checkTokenProviderConfigWSIDFromEnv(tokenResource, "TEST_WORKSPACE_ID"),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
		Step{
			Template: tokenWithProviderBlock("", ""),
			ExpectError: regexp.MustCompile(
				`resource has provider_config.workspace_id = .* in state, but managing workspace-level resources requires a workspace_id`,
			),
			ProtoV6ProviderFactories: noOidcProviderFactories(),
		},
	)
}
