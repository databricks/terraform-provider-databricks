package libraries

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// NewLibrariesAPI creates LibrariesAPI instance from provider meta
func NewLibrariesAPI(ctx context.Context, m any) LibrariesAPI {
	// TODO: context.WithValue
	return LibrariesAPI{
		client:  m.(*common.DatabricksClient),
		context: ctx,
	}
}

// LibrariesAPI exposes the Library API
type LibrariesAPI struct {
	client  *common.DatabricksClient
	context context.Context
}

// Install library list on cluster
func (a LibrariesAPI) Install(req ClusterLibraryList) error {
	return a.client.Post(a.context, "/libraries/install", req, nil)
}

// Uninstall library list from cluster
func (a LibrariesAPI) Uninstall(req ClusterLibraryList) error {
	return a.client.Post(a.context, "/libraries/uninstall", req, nil)
}

// ClusterStatus returns library status in cluster
func (a LibrariesAPI) ClusterStatus(clusterID string) (cls ClusterLibraryStatuses, err error) {
	err = a.client.Get(a.context, "/libraries/cluster-status", ClusterID{
		ClusterID: clusterID,
	}, &cls)
	return
}

type Wait struct {
	ClusterID string
	Timeout   time.Duration
	IsRunning bool
	IsRefresh bool
}

func (a LibrariesAPI) UpdateLibraries(clusterID string, add, remove ClusterLibraryList, timeout time.Duration) error {
	if len(remove.Libraries) > 0 {
		err := a.Uninstall(remove)
		if err != nil {
			return err
		}
	}
	if len(add.Libraries) > 0 {
		err := a.Install(add)
		if err != nil {
			return err
		}
	}
	_, err := a.WaitForLibrariesInstalled(Wait{
		ClusterID: clusterID,
		Timeout:   timeout,
		IsRunning: true,
		IsRefresh: false,
	})
	return err
}

// clusterID string, timeout time.Duration, isActive bool, refresh bool
func (a LibrariesAPI) WaitForLibrariesInstalled(wait Wait) (result *ClusterLibraryStatuses, err error) {
	err = resource.RetryContext(a.context, wait.Timeout, func() *resource.RetryError {
		libsClusterStatus, err := a.ClusterStatus(wait.ClusterID)
		if err != nil {
			apiErr, ok := err.(common.APIError)
			if !ok {
				return resource.NonRetryableError(err)
			}
			if apiErr.StatusCode != 404 && strings.Contains(apiErr.Message,
				fmt.Sprintf("Cluster %s does not exist", wait.ClusterID)) {
				apiErr.StatusCode = 404
			}
			return resource.NonRetryableError(apiErr)
		}
		if !wait.IsRunning {
			log.Printf("[INFO] Cluster %s is currently not running, so just returning list of %d libraries",
				wait.ClusterID, len(libsClusterStatus.LibraryStatuses))
			result = &libsClusterStatus
			return nil
		}
		retry, err := libsClusterStatus.IsRetryNeeded(wait.IsRefresh)
		if retry {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		result = &libsClusterStatus
		return nil
	})
	if err != nil {
		return
	}
	if wait.IsRunning {
		installed := []LibraryStatus{}
		cleanup := ClusterLibraryList{
			ClusterID: wait.ClusterID,
			Libraries: []Library{},
		}
		// cleanup libraries that failed to install
		for _, v := range result.LibraryStatuses {
			if v.Status == "FAILED" {
				log.Printf("[WARN] Removing failed library %s from %s", v.Library, wait.ClusterID)
				cleanup.Libraries = append(cleanup.Libraries, *v.Library)
				continue
			}
			installed = append(installed, v)
		}
		// and result contains only the libraries that were successfully installed
		result.LibraryStatuses = installed
		if len(cleanup.Libraries) > 0 {
			err = a.Uninstall(cleanup)
			if err != nil {
				err = fmt.Errorf("cannot cleanup libraries: %w", err)
			}
		}
	}
	return
}

// TODO: make package private
type ClusterID struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}

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

// NewLibraryFromInstanceState returns library from instance state for
// custom schema hash function. The thing is that for sets of types with
// optional subtypes resource.SerializeResourceForHash doesn't seem to
// predictably give consistent hashcode - that's why AWS launched
// TestAccClusterLibraryList is flaky. As an alternative, i tried reflect
// resource, but init of everything takes way more lines of code, than
// this thing. And this function is readable enough.
func NewLibraryFromInstanceState(i any) (lib Library) {
	raw := i.(map[string]any)
	// set field value to "default value", if there's no raw map value
	lib.Jar, _ = raw["jar"].(string)
	lib.Egg, _ = raw["egg"].(string)
	lib.Whl, _ = raw["whl"].(string)
	// remember - nested blocks are lists for terraform
	pypiList, ok := raw["pypi"].([]any)
	if ok && len(pypiList) == 1 {
		lib.Pypi = &PyPi{}
		pypi := pypiList[0].(map[string]any)
		lib.Pypi.Package, _ = pypi["package"].(string)
		lib.Pypi.Repo, _ = pypi["repo"].(string)
	}
	mavenList, ok := raw["maven"].([]any)
	if ok && len(mavenList) == 1 {
		lib.Maven = &Maven{}
		maven := mavenList[0].(map[string]any)
		lib.Maven.Coordinates, _ = maven["coordinates"].(string)
		lib.Maven.Repo, _ = maven["repo"].(string)
	}
	cranList, ok := raw["cran"].([]any)
	if ok && len(cranList) == 1 {
		lib.Cran = &Cran{}
		cran := cranList[0].(map[string]any)
		lib.Cran.Package, _ = cran["package"].(string)
		lib.Cran.Repo, _ = cran["repo"].(string)
	}
	return lib
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

func (library Library) String() string {
	if library.Whl != "" {
		return fmt.Sprintf("whl:%s", library.Whl)
	}
	if library.Jar != "" {
		return fmt.Sprintf("jar:%s", library.Jar)
	}
	if library.Pypi != nil && library.Pypi.Package != "" {
		return fmt.Sprintf("pypi:%s%s", library.Pypi.Repo, library.Pypi.Package)
	}
	if library.Maven != nil && library.Maven.Coordinates != "" {
		mvn := library.Maven
		return fmt.Sprintf("mvn:%s%s%s", mvn.Repo, mvn.Coordinates,
			strings.Join(mvn.Exclusions, ""))
	}
	if library.Egg != "" {
		return fmt.Sprintf("egg:%s", library.Egg)
	}
	if library.Cran != nil && library.Cran.Package != "" {
		return fmt.Sprintf("cran:%s%s", library.Cran.Repo, library.Cran.Package)
	}
	return "unknown"
}

// ClusterLibraryList is request body for install and uninstall
type ClusterLibraryList struct {
	ClusterID string    `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Libraries []Library `json:"libraries,omitempty" tf:"slice_set,alias:library"`
}

// Diff returns install/uninstall lists given a cluster lib status
func (cll *ClusterLibraryList) Diff(cls ClusterLibraryStatuses) (ClusterLibraryList, ClusterLibraryList) {
	inConfig := map[string]Library{}
	for _, lib := range cll.Libraries {
		inConfig[lib.String()] = lib
	}
	inState := map[string]Library{}
	for _, status := range cls.LibraryStatuses {
		lib := *status.Library
		inState[lib.String()] = lib
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
	toInstall.Sort()
	toUninstall.Sort()
	return toInstall, toUninstall
}

func (cll *ClusterLibraryList) Sort() {
	sort.Slice(cll.Libraries, func(i, j int) bool {
		return cll.Libraries[i].String() < cll.Libraries[j].String()
	})
}

func (cll *ClusterLibraryList) String() string {
	libs := make([]string, len(cll.Libraries))
	for i, lib := range cll.Libraries {
		libs[i] = lib.String()
	}
	return fmt.Sprintf("%s/%s", cll.ClusterID, strings.Join(libs, ","))
}

// LibraryStatus is the status on a given cluster when using the libraries status api
type LibraryStatus struct {
	Library  *Library `json:"library,omitempty"`
	Status   string   `json:"status,omitempty"`
	IsGlobal bool     `json:"is_library_for_all_clusters,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

// ClusterLibraryStatuses  A status will be available for all libraries installed on the cluster via the API or
// the libraries UI as well as libraries set to be installed on all clusters via the libraries UI. If a library
// has been set to be installed on all clusters, is_library_for_all_clusters will be true, even if the library
// was also installed on the cluster.
type ClusterLibraryStatuses struct {
	ClusterID       string          `json:"cluster_id,omitempty"`
	LibraryStatuses []LibraryStatus `json:"library_statuses,omitempty"`
}

// ToLibraryList convert to envity for convenient comparison
func (cls ClusterLibraryStatuses) ToLibraryList() ClusterLibraryList {
	cll := ClusterLibraryList{ClusterID: cls.ClusterID}
	for _, lib := range cls.LibraryStatuses {
		cll.Libraries = append(cll.Libraries, *lib.Library)
	}
	cll.Sort()
	return cll
}

// IsRetryNeeded returns first bool if there needs to be retry.
// If there needs to be retry, error message will explain why.
// If retry does not need to happen and error is not nil - it failed.
func (cls ClusterLibraryStatuses) IsRetryNeeded(refresh bool) (bool, error) {
	pending := 0
	ready := 0
	errors := []string{}
	for _, lib := range cls.LibraryStatuses {
		if lib.IsGlobal {
			continue
		}
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
			if refresh {
				// we're reading library list on a running cluster and some of the libs failed to install
				continue
			}
			errors = append(errors, fmt.Sprintf("%s failed: %s", lib.Library, strings.Join(lib.Messages, ", ")))
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
