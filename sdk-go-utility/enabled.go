package sdkgoutility

import (
	"os"
	"slices"
	"strings"
	"unicode"
)

// enabledEnvName is the environment variable that lists the resources for
// which Enabled is true.
const enabledEnvName = "USE_SDK_GO_RESOURCES"

// Enabled reports whether the environment variable USE_SDK_GO_RESOURCES names
// resource, a Terraform resource type such as databricks_mlflow_model. A
// resource that has a classic SDK path and a modular SDK path calls it to pick
// one of them for an operation.
//
// The variable holds resource types separated by commas, white space (as
// unicode.IsSpace defines it) or both, for example:
//
//	USE_SDK_GO_RESOURCES="databricks_mlflow_model, databricks_volume"
//
// Empty entries are ignored. A name is compared with resource exactly: letter
// case counts, and a name doesn't stand for a longer or a shorter one. Enabled
// is false for every resource when the variable is unset or empty, and for a
// resource that is empty or contains a separator.
//
// Enabled reads the variable on every call and keeps no state, so a change of
// the variable takes effect on the next call, and a test switches paths with
// t.Setenv.
//
// Enabled deliberately doesn't:
//
//   - validate names. It doesn't know which resources have a modular SDK
//     path, so a name in the variable that no resource asks about, such as a
//     misspelled one, enables nothing and isn't reported. The caller passes
//     the resource type as Terraform spells it.
//   - pick the path or remember the answer. The caller branches on the
//     result, and calls Enabled again for every operation that is to follow a
//     change of the variable.
func Enabled(resource string) bool {
	names := strings.FieldsFunc(os.Getenv(enabledEnvName), enabledIsSeparator)
	return slices.Contains(names, resource)
}

// enabledIsSeparator reports whether r separates two names in the value of
// USE_SDK_GO_RESOURCES.
func enabledIsSeparator(r rune) bool {
	return r == ',' || unicode.IsSpace(r)
}
