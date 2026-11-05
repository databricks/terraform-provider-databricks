# Pending acceptance scenarios

**7 original functions remain pending** in the 545-function baseline. See
[MIGRATION.md](MIGRATION.md) for counts and [FOLLOW_UPS.md](FOLLOW_UPS.md) for
blocker details. Original tests remain unchanged.

A row can describe a reproduced failure, a limitation of the current runner, or
a source-inspection risk. Inspection risks are not claims that every listed
scenario was run and failed. P17 is an inventory exception, not missing coverage.
Provider races are tracked independently of conversion; a P1 warning alone
does not require different replay fixtures or prevent writing the test.
P4 now tracks 16 cases ready for conversion after source review, not a missing
engine feature. Its SQL-warehouse import case is tracked under P5 instead.
The P4 cases still count as pending until their replay copies pass.
All five P5 import cases now have replay copies.

| Source | Original function | Blockers | Reason |
| --- | --- | --- | --- |
| [catalog/credential_test.go](../../catalog/credential_test.go) | `TestAccCredentialOwner` | P13 | Workspace credential configured without workspace routing. |
| [internal/acceptance/tf_provider_test.go](../../internal/acceptance/tf_provider_test.go) | `TestAccProviderPlanShouldSucceedWithIncompleteConfiguration` | P6 | Additional noop provider; already hermetic. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Invalid` | P8 | Current diagnostic differs from the original expectation. |
| [internal/providers/pluginfw/products/sharing/resource_share_acc_test.go](../../internal/providers/pluginfw/products/sharing/resource_share_acc_test.go) | `TestAccShare_ProviderConfig_Recreate` | P9 | Skipped for migration: planning rejects the changed workspace ID before the expected replacement plan. Original test remains unchanged. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesProvisioningToRunning` | P14 | Provisioning Read loses the input-only expected status. |
| [mws/mws_workspaces_test.go](../../mws/mws_workspaces_test.go) | `TestMwsAccGcpWorkspacesUnsetExpectedState` | P14 | Provisioning Read loses the input-only expected status. |
| [permissions/resource_permissions_test.go](../../permissions/resource_permissions_test.go) | `TestAccessControlHashFunction` | P17 | Already hermetic; inventory classification. |
