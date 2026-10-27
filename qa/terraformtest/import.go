package terraformtest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/databricks/terraform-provider-databricks/qa/httpfixture"
)

// importStep exercises the provider's importer without replacing the lifecycle's
// state. Verification compares imported values with the previously applied state.
type importStep struct {
	Resource            string   `yaml:"resource"`
	ID                  string   `yaml:"id,omitempty"`
	Verify              bool     `yaml:"verify,omitempty"`
	Ignore              []string `yaml:"ignore,omitempty"`
	IdentifierAttribute string   `yaml:"identifier_attribute,omitempty"`
}

func (s step) validateImport(hasPriorState bool) error {
	if s.Import == nil {
		return nil
	}
	if strings.TrimSpace(s.Import.Resource) == "" {
		return fmt.Errorf("import.resource is required")
	}
	if !hasPriorState {
		return fmt.Errorf("import requires a prior successful apply step or initial_state")
	}
	if s.PlanOnly || s.Plan != nil || s.State != nil {
		return fmt.Errorf("import cannot be combined with plan_only, plan or state; use import.verify for state comparison")
	}
	if !s.Import.Verify && (len(s.Import.Ignore) != 0 || s.Import.IdentifierAttribute != "") {
		return fmt.Errorf("import.ignore and import.identifier_attribute require import.verify: true")
	}
	if s.Import.Verify && s.ExpectError != nil {
		return fmt.Errorf("import.verify cannot be combined with expect_error")
	}
	for _, prefix := range s.Import.Ignore {
		if strings.TrimSpace(prefix) == "" {
			return fmt.Errorf("import.ignore must contain nonempty attribute prefixes")
		}
	}
	return nil
}

// configureImport delegates CLI execution and state comparison to HashiCorp.
// ImportStatePersist stays false: importing into a separate working directory
// must not discard other managed resources or replace state needed for cleanup.
func configureImport(native *resource.TestStep, definition importStep, handler *httpfixture.Handler) error {
	// HashiCorp initializes a fresh directory from this raw configuration; its
	// normal apply-time provider-requirements injection does not run there.
	hcl, err := checkpointConfig(native.Config)
	if err != nil {
		return err
	}
	native.Config = hcl
	native.ImportState = true
	native.ResourceName = definition.Resource
	native.ImportStateId = definition.ID
	native.ImportStateVerify = definition.Verify
	native.ImportStateVerifyIgnore = definition.Ignore
	native.ImportStateVerifyIdentifierAttribute = definition.IdentifierAttribute
	// Import steps do not invoke PostApplyFunc. Verify before HashiCorp checks
	// imported state so unused fixtures fail within the import step itself.
	native.PostApplyFunc = nil
	native.ImportStateCheck = func(_ []*terraform.InstanceState) error {
		return handler.Verify()
	}
	return nil
}
