package libraries

import (
	"sort"

	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type LibraryList struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId string `json:"cluster_id"`
	// The libraries to install.
	Libraries []compute.Library `json:"libraries,omitempty"`
}

func (l *LibraryList) Sort() {
	sort.Slice(l.Libraries, func(i, j int) bool {
		return l.Libraries[i].String() < l.Libraries[j].String()
	})
}

func (l LibraryList) Aliases() map[string]map[string]string {
	return map[string]map[string]string{
		"libraries.LibraryList": {
			"libraries": "library",
		},
	}
}

func (l LibraryList) CustomizeSchema(m map[string]*schema.Schema) map[string]*schema.Schema {
	common.CustomizeSchemaPath(m, "library").Schema.Set = func(i any) int {
		lib := NewLibraryFromInstanceState(i)
		return schema.HashString(lib.String())
	}
	common.CustomizeSchemaPath(m, "library").SetSliceSet()
	return m
}

// NewLibraryFromInstanceState returns library from instance state for
// custom schema hash function. The thing is that for sets of types with
// optional subtypes resource.SerializeResourceForHash doesn't seem to
// predictably give consistent hashcode - that's why AWS launched
// TestAccClusterLibraryList is flaky. As an alternative, i tried reflect
// resource, but init of everything takes way more lines of code, than
// this thing. And this function is readable enough.
func NewLibraryFromInstanceState(i any) (lib compute.Library) {
	raw := i.(map[string]any)
	// set field value to "default value", if there's no raw map value
	lib.Jar, _ = raw["jar"].(string)
	lib.Egg, _ = raw["egg"].(string)
	lib.Whl, _ = raw["whl"].(string)
	// remember - nested blocks are lists for terraform
	pypiList, ok := raw["pypi"].([]any)
	if ok && len(pypiList) == 1 {
		lib.Pypi = &compute.PythonPyPiLibrary{}
		pypi := pypiList[0].(map[string]any)
		lib.Pypi.Package, _ = pypi["package"].(string)
		lib.Pypi.Repo, _ = pypi["repo"].(string)
	}
	mavenList, ok := raw["maven"].([]any)
	if ok && len(mavenList) == 1 {
		lib.Maven = &compute.MavenLibrary{}
		maven := mavenList[0].(map[string]any)
		lib.Maven.Coordinates, _ = maven["coordinates"].(string)
		lib.Maven.Repo, _ = maven["repo"].(string)
	}
	cranList, ok := raw["cran"].([]any)
	if ok && len(cranList) == 1 {
		lib.Cran = &compute.RCranLibrary{}
		cran := cranList[0].(map[string]any)
		lib.Cran.Package, _ = cran["package"].(string)
		lib.Cran.Repo, _ = cran["repo"].(string)
	}
	return lib
}

// Diff returns install/uninstall lists given a cluster lib status
func GetLibrariesToInstallAndUninstall(cll LibraryList, cls *compute.ClusterLibraryStatuses) (LibraryList, LibraryList) {
	inConfig := map[string]compute.Library{}
	for _, lib := range cll.Libraries {
		inConfig[lib.String()] = lib
	}
	inState := map[string]compute.Library{}
	for _, status := range cls.LibraryStatuses {
		lib := *status.Library
		inState[lib.String()] = lib
	}
	toInstall := LibraryList{ClusterId: cll.ClusterId}
	toUninstall := LibraryList{ClusterId: cll.ClusterId}
	for key, lib := range inConfig {
		_, exists := inState[key]
		if exists {
			continue
		}
		toInstall.Libraries = append(toInstall.Libraries, lib)
	}
	for key, lib := range inState {
		_, exists := inConfig[key]
		if exists {
			continue
		}
		toUninstall.Libraries = append(toUninstall.Libraries, lib)
	}
	toInstall.Sort()
	toUninstall.Sort()
	return toInstall, toUninstall
}
