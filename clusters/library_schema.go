package clusters

import (
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
)

// EggDeprecationWarning is referenced by the cluster resource to deprecate
// the `library.egg` field in databricks_cluster.
const EggDeprecationWarning = "The `egg` library type is deprecated. Please use `whl` or `pypi` instead."

// LibraryResource is a schema-only type used to register schema customizations
// for nested compute.Library blocks (e.g. the `library` block inside
// databricks_cluster). The standalone databricks_library resource is served by
// the Plugin Framework and no longer lives in this package, but the nested-block
// schema is still derived via this registration so existing cluster state files
// continue to round-trip.
type LibraryResource struct {
	compute.Library
	common.Namespace
}

func (LibraryResource) CustomizeSchema(s *common.CustomizableSchema) *common.CustomizableSchema {
	common.NamespaceCustomizeSchema(s)
	s.SchemaPath("egg").SetDeprecated(EggDeprecationWarning)
	return s
}
