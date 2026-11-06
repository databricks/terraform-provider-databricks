package terraformtest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// importStep makes state ownership explicit: persistent imports adopt resources
// into the journey; nonpersistent imports use separate, discarded state.
type importStep struct {
	Resource string `yaml:"resource"`
	// A pointer distinguishes an explicit false from an omitted or null choice.
	Persist             *bool    `yaml:"persist"`
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
	if s.Import.Persist == nil {
		return fmt.Errorf("import.persist is required; set true to retain imported state or false to discard it")
	}
	if strings.TrimSpace(s.Import.ID) == "" && (*s.Import.Persist || !hasPriorState) {
		return fmt.Errorf("import.id is required for persistent imports or when no prior state exists")
	}
	if s.Import.Verify && (*s.Import.Persist || !hasPriorState) {
		return fmt.Errorf("import.verify requires persist: false and prior state to compare against; check persistent imports in a following step")
	}
	if s.PlanOnly || s.Plan != nil || s.State != nil {
		return fmt.Errorf("import cannot be combined with plan_only, plan or state; check persistent imports in a following step or use import.verify with persist: false")
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

// configureImport delegates CLI execution to HashiCorp. Persist selects the
// journey's working directory; false selects a separate directory whose imported
// state is discarded without replacing the journey's state.
func configureImport(native *resource.TestStep, definition importStep, handler stepHTTP) error {
	// Native import installs raw configuration without the normal apply-time
	// provider-requirements injection, so supply the canonical provider address.
	hcl, err := checkpointConfig(native.Config)
	if err != nil {
		return err
	}
	native.Config = hcl
	native.ImportState = true
	native.ImportStatePersist = *definition.Persist
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
