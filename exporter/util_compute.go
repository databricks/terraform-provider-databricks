package exporter

import (
	"fmt"
	"log"
	"regexp"

	"github.com/databricks/terraform-provider-databricks/common"

	"github.com/databricks/databricks-sdk-go/service/compute"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func (ic *importContext) emitInitScripts(initScripts []compute.InitScriptInfo) {
	for _, is := range initScripts {
		if is.Dbfs != nil {
			ic.Emit(&resource{
				Resource: "databricks_dbfs_file",
				ID:       is.Dbfs.Destination,
			})
		}
		if is.Workspace != nil {
			ic.emitWorkspaceFileOrRepo(is.Workspace.Destination)
		}
		if is.Volumes != nil {
			// TODO: we should emit allow list for init scripts as well
			ic.emitIfVolumeFile(is.Volumes.Destination)
		}
	}
}

func (ic *importContext) emitSecretsFromSecretPathString(v string) {
	if res := secretPathRegex.FindStringSubmatch(v); res != nil {
		ic.Emit(&resource{
			Resource: "databricks_secret_scope",
			ID:       res[1],
		})
	}
}

func (ic *importContext) emitSecretsFromSecretsPathMap(m map[string]string) {
	for _, v := range m {
		ic.emitSecretsFromSecretPathString(v)
	}
}

func (ic *importContext) emitLibraries(libs []compute.Library) {
	for _, lib := range libs {
		ic.emitLibrary(&lib)
	}
}

func (ic *importContext) emitLibrary(lib *compute.Library) {
	// Files on DBFS
	ic.emitIfDbfsFile(lib.Whl)
	ic.emitIfDbfsFile(lib.Jar)
	ic.emitIfDbfsFile(lib.Egg)
	// Files on WSFS
	ic.emitIfWsfsFile(lib.Whl)
	ic.emitIfWsfsFile(lib.Jar)
	ic.emitIfWsfsFile(lib.Egg)
	ic.emitIfWsfsFile(lib.Requirements)
	// Files on UC Volumes
	ic.emitIfVolumeFile(lib.Whl)
	// TODO: we should emit UC allow list as well
	ic.emitIfVolumeFile(lib.Jar)
	ic.emitIfVolumeFile(lib.Requirements)
}

func (ic *importContext) importLibraries(d *schema.ResourceData, s map[string]*schema.Schema) error {
	var cll compute.InstallLibraries
	common.DataToStructPointer(d, s, &cll)
	ic.emitLibraries(cll.Libraries)
	return nil
}

func (ic *importContext) importClusterLibraries(d *schema.ResourceData) error {
	libraries := ic.workspaceClient.Libraries
	cll, err := libraries.ClusterStatusByClusterId(ic.Context, d.Id())
	if err != nil {
		return err
	}
	for _, lib := range cll.LibraryStatuses {
		if lib.Library != nil {
			ic.emitLibrary(lib.Library)
		}
	}
	return nil
}

func (ic *importContext) getBuiltinPolicyFamilies() map[string]compute.PolicyFamily {
	ic.builtInPoliciesMutex.Lock()
	defer ic.builtInPoliciesMutex.Unlock()
	if ic.builtInPolicies == nil {
		if !ic.accountLevel {
			log.Printf("[DEBUG] Going to initialize ic.builtInPolicies. Getting policy families...")
			families, err := ic.workspaceClient.PolicyFamilies.ListAll(ic.Context, compute.ListPolicyFamiliesRequest{})
			log.Printf("[DEBUG] Going to initialize ic.builtInPolicies. Getting policy families...")
			if err == nil {
				ic.builtInPolicies = make(map[string]compute.PolicyFamily, len(families))
				for _, f := range families {
					f2 := f
					ic.builtInPolicies[f2.PolicyFamilyId] = f2
				}
			} else {
				log.Printf("[ERROR] Can't fetch cluster policy families: %v", err)
				ic.builtInPolicies = map[string]compute.PolicyFamily{}
			}
		} else {
			log.Print("[WARN] Can't list cluster policy families on account level")
			ic.builtInPolicies = map[string]compute.PolicyFamily{}
		}
	}
	return ic.builtInPolicies
}

func makeShouldOmitFieldForCluster(regex *regexp.Regexp) func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
	return func(ic *importContext, pathString string, as *schema.Schema, d *schema.ResourceData) bool {
		prefix := ""
		if regex != nil {
			if res := regex.FindStringSubmatch(pathString); res != nil {
				prefix = res[0]
			} else {
				return false
			}
		}
		raw := d.Get(pathString)
		workerInstPoolID := d.Get(prefix + "instance_pool_id").(string)
		switch pathString {
		case prefix + "node_type_id":
			return workerInstPoolID != ""
		case prefix + "driver_node_type_id":
			driverInstPoolID := d.Get(prefix + "driver_instance_pool_id").(string)
			nodeTypeID := d.Get(prefix + "node_type_id").(string)
			return workerInstPoolID != "" || driverInstPoolID != "" || raw.(string) == nodeTypeID
		case prefix + "driver_instance_pool_id":
			return raw.(string) == workerInstPoolID
		case prefix + "enable_elastic_disk", prefix + "aws_attributes", prefix + "azure_attributes", prefix + "gcp_attributes":
			return workerInstPoolID != ""
		case prefix + "enable_local_disk_encryption":
			return false
		case prefix + "spark_conf":
			return fmt.Sprintf("%v", d.Get(prefix+"spark_conf")) == "map[spark.databricks.delta.preview.enabled:true]"
		case prefix + "spark_env_vars":
			return fmt.Sprintf("%v", d.Get(prefix+"spark_env_vars")) == "map[PYSPARK_PYTHON:/databricks/python3/bin/python3]"
		}

		return defaultShouldOmitFieldFunc(ic, pathString, as, d)
	}
}
