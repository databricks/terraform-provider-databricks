package model

import (
	"fmt"
	"strings"
)

// PyPi is a python library hosted on PYPI
type PyPi struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`
}

// Maven is a jar library hosted on Maven
type Maven struct {
	Coordinates string   `json:"coordinates"`
	Repo        string   `json:"repo,omitempty"`
	Exclusions  []string `json:"exclusions,omitempty"`
}

// Cran is a R library hosted on Maven
type Cran struct {
	Package string `json:"package"`
	Repo    string `json:"repo,omitempty"`
}

// Library is a construct that contains information of the location of the library and how to download it
type Library struct { // TODO: discuss if we can make a dedicated entity just for terraform...
	Jar   string `json:"jar,omitempty" tf:"group:lib"`
	Egg   string `json:"egg,omitempty" tf:"group:lib"`
	Whl   string `json:"whl,omitempty" tf:"group:lib"`
	Pypi  *PyPi  `json:"pypi,omitempty" tf:"group:lib"`
	Maven *Maven `json:"maven,omitempty" tf:"group:lib"`
	Cran  *Cran  `json:"cran,omitempty" tf:"group:lib"`
}

// TypeAndKey can be used for computing differences
func (library Library) TypeAndKey() (string, string) {
	switch {
	case len(library.Whl) > 0:
		return "library_whl", library.Whl
	case len(library.Egg) > 0:
		return "library_egg", library.Egg
	case len(library.Jar) > 0:
		return "library_jar", library.Jar
	case library.Pypi != nil && len(library.Pypi.Package) > 0:
		return "library_pypi", library.Pypi.Package + library.Pypi.Repo
	case library.Maven != nil && len(library.Maven.Coordinates) > 0:
		return "library_maven", library.Maven.Coordinates + library.Maven.Repo + strings.Join(library.Maven.Exclusions, "")
	case library.Cran != nil && len(library.Cran.Package) > 0:
		return "library_cran", library.Cran.Package + library.Cran.Repo
	}
	return "", ""
}

// ClusterLibraryList is request body for install and uninstall
type ClusterLibraryList struct {
	ClusterID string    `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Libraries []Library `json:"libraries,omitempty" url:"libraries,omitempty" tf:"slice_set,alias:library"`
}

// Diff returns install/uninstall lists given a cluster lib status
func (cll *ClusterLibraryList) Diff(cls ClusterLibraryStatuses) (ClusterLibraryList, ClusterLibraryList) {
	inConfig := map[string]Library{}
	for _, lib := range cll.Libraries {
		_, key := lib.TypeAndKey()
		inConfig[key] = lib
	}
	inState := map[string]Library{}
	for _, status := range cls.LibraryStatuses {
		lib := *status.Library
		_, key := lib.TypeAndKey()
		inState[key] = lib
	}
	toInstall := ClusterLibraryList{ClusterID: cll.ClusterID}
	toUninstall := ClusterLibraryList{ClusterID: cll.ClusterID}
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
	return toInstall, toUninstall
}

// Apply runs a function cb when number of libraries is non-zero
func (cll *ClusterLibraryList) Apply(cb func(this ClusterLibraryList) error) error {
	if len(cll.Libraries) == 0 {
		return nil
	}
	return cb(*cll)
}

// AddLibraryFromMap is convenience method
func (cll *ClusterLibraryList) AddLibraryFromMap(libraryType string, m map[string]interface{}) {
	lib := Library{}
	cll.Libraries = append(cll.Libraries, lib)
	switch libraryType {
	case "library_whl":
		if v, ok := m["path"].(string); ok {
			lib.Whl = v
		}
	case "library_egg":
		if v, ok := m["path"].(string); ok {
			lib.Egg = v
		}
	case "library_jar":
		if v, ok := m["path"].(string); ok {
			lib.Jar = v
		}
	case "library_pypi":
		lib.Pypi = &PyPi{}
		if v, ok := m["package"].(string); ok {
			lib.Pypi.Package = v
		}
		if v, ok := m["repo"].(string); ok {
			lib.Pypi.Package = v
		}
	case "library_maven":
		lib.Maven = &Maven{}
		if v, ok := m["coordinates"].(string); ok {
			lib.Maven.Coordinates = v
		}
		if v, ok := m["repo"].(string); ok {
			lib.Maven.Repo = v
		}
		// exclusions won't be added for now.
	case "library_cran":
		lib.Cran = &Cran{}
		if v, ok := m["package"].(string); ok {
			lib.Cran.Package = v
		}
		if v, ok := m["repo"].(string); ok {
			lib.Cran.Package = v
		}
	}
}

// LibraryStatus is the status on a given cluster when using the libraries status api
type LibraryStatus struct {
	Library                         *Library `json:"library,omitempty"`
	Status                          string   `json:"status,omitempty"`
	IsLibraryInstalledOnAllClusters bool     `json:"is_library_for_all_clusters,omitempty"`
	Messages                        []string `json:"messages,omitempty"`
}

func (lib LibraryStatus) toMap() map[string]interface{} {
	// TODO: make map[string]string, as messages should not be in the state
	m := map[string]interface{}{}
	//m["messages"] = lib.Messages
	m["status"] = lib.Status
	switch {
	case len(lib.Library.Jar) > 0:
		m["path"] = lib.Library.Jar
		return m
	case len(lib.Library.Egg) > 0:
		m["path"] = lib.Library.Egg
		return m
	case len(lib.Library.Whl) > 0:
		m["path"] = lib.Library.Whl
		return m
	case lib.Library.Maven != nil && len(lib.Library.Maven.Coordinates) > 0:
		m["coordinates"] = lib.Library.Maven.Coordinates
		if len(lib.Library.Maven.Repo) > 0 {
			m["repo"] = lib.Library.Maven.Repo
		}
		// if len(lib.Library.Maven.Exclusions) > 0 {
		// 	m["exclusions"] = lib.Library.Maven.Exclusions
		// }
		return m
	case lib.Library.Pypi != nil && len(lib.Library.Pypi.Package) > 0:
		m["package"] = lib.Library.Pypi.Package
		if len(lib.Library.Pypi.Repo) > 0 {
			m["repo"] = lib.Library.Pypi.Repo
		}
		return m
	case lib.Library.Cran != nil && len(lib.Library.Cran.Package) > 0:
		m["package"] = lib.Library.Cran.Package
		if len(lib.Library.Cran.Repo) > 0 {
			m["repo"] = lib.Library.Cran.Repo
		}
		return m
	}
	return m
}

// ClusterLibraryStatuses  A status will be available for all libraries installed on the cluster via the API or
// the libraries UI as well as libraries set to be installed on all clusters via the libraries UI. If a library
// has been set to be installed on all clusters, is_library_for_all_clusters will be true, even if the library
// was also installed on the cluster.
type ClusterLibraryStatuses struct {
	ClusterID       string          `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	LibraryStatuses []LibraryStatus `json:"library_statuses,omitempty" url:"libraries,omitempty"`
}

// ToLibraryList convert to envity for convenient comparison
func (cls ClusterLibraryStatuses) ToLibraryList() ClusterLibraryList {
	cll := ClusterLibraryList{ClusterID: cls.ClusterID}
	for _, lib := range cls.LibraryStatuses {
		cll.Libraries = append(cll.Libraries, *lib.Library)
	}
	return cll
}

// Apply calls a function with type of library list and library list
func (cls ClusterLibraryStatuses) Apply(
	cb func(key string, value interface{}) error) error {
	x := make(map[string][]map[string]interface{})
	errors := []string{}
	for _, lib := range cls.LibraryStatuses {
		libraryType, key := lib.Library.TypeAndKey()
		if lib.Status == "FAILED" {
			errors = append(errors, fmt.Sprintf("%s[%s] failed: %s", libraryType, key, strings.Join(lib.Messages, ", ")))
			continue
		}
		if len(libraryType) < 1 {
			continue
		}
		m := lib.toMap()
		// Some step in installation failed. More information can be found in the messages field.
		x[libraryType] = append(x[libraryType], m)
	}
	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	for key, value := range x {
		err := cb(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsRetryNeeded returns first bool if there needs to be retry.
// If there needs to be retry, error message will explain why.
// If retry does not need to happen and error is not nil - it failed.
func (cls ClusterLibraryStatuses) IsRetryNeeded() (bool, error) {
	pending := 0
	ready := 0
	errors := []string{}
	for _, lib := range cls.LibraryStatuses {
		switch lib.Status {
		// No action has yet been taken to install the library. This state should be very short lived.
		case "PENDING":
			pending++
		// Metadata necessary to install the library is being retrieved from the provided repository.
		case "RESOLVING":
			pending++
		// The library is actively being installed, either by adding resources to Spark
		// or executing system commands inside the Spark nodes.
		case "INSTALLING":
			pending++
		// The library has been successfully installed.
		case "INSTALLED":
			ready++
		// Installation on a Databricks Runtime 7.0 or above cluster was skipped due to Scala version incompatibility.
		case "SKIPPED":
			ready++
		// The library has been marked for removal. Libraries can be removed only when clusters are restarted.
		case "UNINSTALL_ON_RESTART":
			ready++
			//Some step in installation failed. More information can be found in the messages field.
		case "FAILED":
			libraryType, key := lib.Library.TypeAndKey()
			errors = append(errors, fmt.Sprintf("%s[%s] failed: %s", libraryType, key, strings.Join(lib.Messages, ", ")))
			continue
		}
	}
	if pending > 0 {
		return true, fmt.Errorf("%d libraries are ready, but there are still %d pending", ready, pending)
	}
	if len(errors) > 0 {
		return false, fmt.Errorf("%s", strings.Join(errors, "\n"))
	}
	return false, nil
}
