package exporter

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	sdk_compute "github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/pipelines"
	"github.com/databricks/terraform-provider-databricks/common"
	tf_dlt "github.com/databricks/terraform-provider-databricks/pipelines"
	tf_workspace "github.com/databricks/terraform-provider-databricks/workspace"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func listClusters(ic *importContext) error {
	lastActiveMs := ic.getLastActiveMs()
	interactiveClusters := []sdk_compute.ClusterSource{sdk_compute.ClusterSourceUi, sdk_compute.ClusterSourceApi}

	it := ic.workspaceClient.Clusters.List(ic.Context, sdk_compute.ListClustersRequest{
		FilterBy: &sdk_compute.ListClustersFilterBy{
			ClusterSources: interactiveClusters,
		},
		PageSize: 100,
	})
	i := 0
	for it.HasNext(ic.Context) {
		c, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++

		if strings.HasPrefix(c.ClusterName, "terraform-") {
			log.Printf("[INFO] Skipping terraform-specific cluster %s", c.ClusterName)
			continue
		}
		if !ic.MatchesName(c.ClusterName) {
			log.Printf("[INFO] Skipping %s because it doesn't match %s", c.ClusterName, ic.match)
			continue
		}
		if c.LastRestartedTime > 0 && c.LastRestartedTime < lastActiveMs {
			log.Printf("[INFO] Old inactive cluster %s", c.ClusterName)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_cluster",
			ID:       c.ClusterId,
		})
		if i%50 == 0 {
			log.Printf("[INFO] Scanned %d clusters", i)
		}
	}
	return nil
}

func (ic *importContext) importCluster(c *sdk_compute.ClusterSpec) {
	if c == nil {
		return
	}
	if c.AwsAttributes != nil && c.AwsAttributes.InstanceProfileArn != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_profile",
			ID:       c.AwsAttributes.InstanceProfileArn,
		})
	}
	if c.InstancePoolId != "" {
		// set enable_elastic_disk to false, and remove aws/gcp/azure_attributes
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.InstancePoolId,
		})
	}
	if c.DriverInstancePoolId != "" {
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       c.DriverInstancePoolId,
		})
	}
	if c.PolicyId != "" {
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       c.PolicyId,
		})
	}
	if c.DockerImage != nil && c.DockerImage.BasicAuth != nil {
		ic.emitSecretsFromSecretPathString(c.DockerImage.BasicAuth.Password)
		ic.emitSecretsFromSecretPathString(c.DockerImage.BasicAuth.Username)
	}
	ic.emitInitScripts(c.InitScripts)
	ic.emitSecretsFromSecretsPathMap(c.SparkConf)
	ic.emitSecretsFromSecretsPathMap(c.SparkEnvVars)
	ic.emitUserOrServicePrincipal(c.SingleUserName)
	if c.Kind.String() != "" && c.SingleUserName != "" {
		ic.Emit(&resource{
			Resource:  "databricks_group",
			Attribute: "display_name",
			Value:     c.SingleUserName,
		})
	}
}

func listInstancePools(ic *importContext) error {
	it := ic.workspaceClient.InstancePools.List(ic.Context)
	i := 0
	for it.HasNext(ic.Context) {
		pool, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(pool.InstancePoolName) {
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_instance_pool",
			ID:       pool.InstancePoolId,
		})
		if i%50 == 0 {
			log.Printf("[INFO] Imported %d instance pools", i)
		}
	}
	return nil
}

func instancePoolName(ic *importContext, d *schema.ResourceData) string {
	raw, ok := d.GetOk("instance_pool_name")
	if !ok || raw.(string) == "" {
		return strings.Split(d.Id(), "-")[2]
	}
	return raw.(string)
}

func importInstancePool(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/instance-pools/%s", r.ID),
		"inst_pool_"+ic.Importables["databricks_instance_pool"].Name(ic, r.Data))
	return nil
}

func importCluster(ic *importContext, r *resource) error {
	var c sdk_compute.ClusterSpec
	s := ic.Resources["databricks_cluster"].Schema
	common.DataToStructPointer(r.Data, s, &c)
	ic.importCluster(&c)
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/clusters/%s", r.ID),
		"cluster_"+ic.Importables["databricks_cluster"].Name(ic, r.Data))
	return ic.importClusterLibraries(r.Data)
}

func listClusterPolicies(ic *importContext) error {
	builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
	it := ic.workspaceClient.ClusterPolicies.List(ic.Context, sdk_compute.ListClusterPoliciesRequest{})
	i := 0
	for it.HasNext(ic.Context) {
		policy, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		family, isBuiltin := builtInClusterPolicies[policy.PolicyFamilyId]
		if policy.PolicyFamilyId != "" && isBuiltin && family.Name == policy.Name &&
			policy.PolicyFamilyDefinitionOverrides == "" {
			log.Printf("[DEBUG] Skipping builtin cluster policy '%s' without overrides", policy.Name)
			continue
		}
		if !ic.MatchesName(policy.Name) {
			log.Printf("[DEBUG] Policy %s doesn't match %s filter", policy.Name, ic.match)
			continue
		}
		ic.Emit(&resource{
			Resource: "databricks_cluster_policy",
			ID:       policy.PolicyId,
		})
		if i%10 == 0 {
			log.Printf("[INFO] Scanned %d cluster policies", i)
		}
	}
	return nil
}

func importClusterPolicy(ic *importContext, r *resource) error {
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/cluster-policies/%s", r.ID),
		"cluster_policy_"+ic.Importables["databricks_cluster_policy"].Name(ic, r.Data))

	var clusterPolicy sdk_compute.Policy
	s := ic.Resources["databricks_cluster_policy"].Schema
	common.DataToStructPointer(r.Data, s, &clusterPolicy)

	var definition map[string]map[string]any
	err := json.Unmarshal([]byte(clusterPolicy.Definition), &definition)
	if err != nil {
		return err
	}
	for k, policy := range definition {
		value, vok := policy["value"]
		defaultValue, dok := policy["defaultValue"]
		typ := policy["type"]
		if !vok && !dok {
			log.Printf("[DEBUG] Skipping policy element as it doesn't have both value and defaultValue. k='%v', policy='%v'",
				k, policy)
			continue
		}
		if k == "aws_attributes.instance_profile_arn" {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       eitherString(value, defaultValue),
			})
		}
		if k == "instance_pool_id" || k == "driver_instance_pool_id" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       eitherString(value, defaultValue),
			})
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".dbfs.destination") {
			ic.emitIfDbfsFile(eitherString(value, defaultValue))
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".volumes.destination") {
			ic.emitIfVolumeFile(eitherString(value, defaultValue))
		}
		if typ == "fixed" && strings.HasPrefix(k, "init_scripts.") &&
			strings.HasSuffix(k, ".workspace.destination") {
			ic.emitWorkspaceFileOrRepo(eitherString(value, defaultValue))
		}
		if typ == "fixed" && (strings.HasPrefix(k, "spark_conf.") || strings.HasPrefix(k, "spark_env_vars.")) {
			ic.emitSecretsFromSecretPathString(eitherString(value, defaultValue))
		}
	}

	for _, lib := range clusterPolicy.Libraries {
		ic.emitIfDbfsFile(lib.Whl)
		ic.emitIfDbfsFile(lib.Jar)
		ic.emitIfDbfsFile(lib.Egg)
		ic.emitIfWsfsFile(lib.Whl)
		ic.emitIfWsfsFile(lib.Jar)
		ic.emitIfWsfsFile(lib.Egg)
		ic.emitIfVolumeFile(lib.Whl)
		ic.emitIfVolumeFile(lib.Jar)
	}

	policyFamilyId := clusterPolicy.PolicyFamilyId
	if policyFamilyId != "" && clusterPolicy.Definition != "" {
		// we need to set definition to empty value because otherwise it will be put into
		// generated HCL code for data source, and it only supports the `name` attribute
		r.Data.Set("definition", "")
		builtInClusterPolicies := ic.getBuiltinPolicyFamilies()
		v, isBuiltin := builtInClusterPolicies[policyFamilyId]
		if isBuiltin && clusterPolicy.PolicyFamilyDefinitionOverrides == "" && v.Name == clusterPolicy.Name {
			r.Mode = "data"
		}
	}

	return nil
}

func listPipelines(ic *importContext) error {
	it := ic.workspaceClient.Pipelines.ListPipelines(ic.Context, pipelines.ListPipelinesRequest{
		MaxResults: 100,
	})
	i := 0
	for it.HasNext(ic.Context) {
		q, err := it.Next(ic.Context)
		if err != nil {
			return err
		}
		i++
		if !ic.MatchesName(q.Name) {
			continue
		}
		var modifiedAt int64
		if ic.incremental {
			pipeline, err := ic.workspaceClient.Pipelines.Get(ic.Context, pipelines.GetPipelineRequest{
				PipelineId: q.PipelineId,
			})
			if err != nil {
				return err
			}
			modifiedAt = pipeline.LastModified
		}
		ic.EmitIfUpdatedAfterMillis(&resource{
			Resource: "databricks_pipeline",
			ID:       q.PipelineId,
		}, modifiedAt, fmt.Sprintf("DLT Pipeline '%s'", q.Name))
		if i%100 == 0 {
			log.Printf("[INFO] Imported %d DLT Pipelines", i)
		}
	}
	log.Printf("[INFO] Listed %d DLT pipelines", i)
	return nil
}

func importPipeline(ic *importContext, r *resource) error {
	var pipeline tf_dlt.Pipeline
	s := ic.Resources["databricks_pipeline"].Schema
	common.DataToStructPointer(r.Data, s, &pipeline)
	if pipeline.Deployment != nil && pipeline.Deployment.Kind == "BUNDLE" {
		log.Printf("[INFO] Skipping processing of DLT Pipeline with ID %s (%s) as deployed with DABs",
			r.ID, pipeline.Name)
		return nil
	}
	if pipeline.Catalog != "" {
		var schemaName string
		if pipeline.Target != "" {
			schemaName = pipeline.Target
		} else if pipeline.Schema != "" {
			schemaName = pipeline.Schema
		}
		if schemaName != "" {
			ic.Emit(&resource{
				Resource: "databricks_schema",
				ID:       pipeline.Catalog + "." + pipeline.Target,
			})
		}
	}
	if pipeline.Deployment == nil || pipeline.Deployment.Kind != "BUNDLE" {
		nbAPI := tf_workspace.NewNotebooksAPI(ic.Context, ic.Client)
		for _, lib := range pipeline.Libraries {
			if lib.Notebook != nil {
				ic.emitNotebookOrRepo(lib.Notebook.Path)
			}
			if lib.File != nil {
				ic.emitWorkspaceFileOrRepo(lib.File.Path)
			}
			if lib.Glob != nil {
				if strings.HasSuffix(lib.Glob.Include, "/**") {
					// Emit all files and notebooks under the directory
					dirPath := strings.TrimSuffix(lib.Glob.Include, "/**")
					ic.emitDirectoryOrRepo(dirPath)
					objects, err := nbAPI.List(dirPath, false, true)
					if err == nil {
						for _, object := range objects {
							switch object.ObjectType {
							case tf_workspace.File:
								ic.emitWorkspaceFileOrRepo(object.Path)
							case tf_workspace.Notebook:
								ic.emitNotebookOrRepo(object.Path)
							}
						}
					} else {
						log.Printf("[WARN] Can't list directory %s for DLT pipeline %s", lib.Glob.Include, pipeline.Name)
					}
				} else {
					// Perform get-status call to check what is the object type
					object, err := nbAPI.GetStatus(lib.Glob.Include, false)
					if err == nil {
						switch object.ObjectType {
						case tf_workspace.File:
							ic.emitWorkspaceFileOrRepo(lib.Glob.Include)
						case tf_workspace.Notebook:
							ic.emitNotebookOrRepo(lib.Glob.Include)
						}
					} else {
						log.Printf("[WARN] Can't get status of %s for DLT pipeline %s", lib.Glob.Include, pipeline.Name)
					}
				}
			}
			ic.emitIfDbfsFile(lib.Jar)
			ic.emitIfDbfsFile(lib.Whl)
		}
		if pipeline.RootPath != "" {
			ic.emitDirectoryOrRepo(pipeline.RootPath)
		}
	}
	// Emit clusters
	for _, cluster := range pipeline.Clusters {
		if cluster.AwsAttributes != nil && cluster.AwsAttributes.InstanceProfileArn != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_profile",
				ID:       cluster.AwsAttributes.InstanceProfileArn,
			})
		}
		if cluster.InstancePoolId != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       cluster.InstancePoolId,
			})
		}
		if cluster.DriverInstancePoolId != "" {
			ic.Emit(&resource{
				Resource: "databricks_instance_pool",
				ID:       cluster.DriverInstancePoolId,
			})
		}
		if cluster.PolicyId != "" {
			ic.Emit(&resource{
				Resource: "databricks_cluster_policy",
				ID:       cluster.PolicyId,
			})
		}
		ic.emitInitScripts(cluster.InitScripts)
		ic.emitSecretsFromSecretsPathMap(cluster.SparkConf)
		ic.emitSecretsFromSecretsPathMap(cluster.SparkEnvVars)
	}
	ic.emitFilesFromMap(pipeline.Configuration)
	ic.emitSecretsFromSecretsPathMap(pipeline.Configuration)
	ic.emitPermissionsIfNotIgnored(r, fmt.Sprintf("/pipelines/%s", r.ID),
		"pipeline_"+ic.Importables["databricks_pipeline"].Name(ic, r.Data))
	if pipeline.Notifications != nil {
		for _, n := range pipeline.Notifications {
			ic.emitListOfUsers(n.EmailRecipients)
		}
	}
	if pipeline.EventLog != nil {
		var catalogName, schemaName string
		if pipeline.EventLog.Catalog != "" {
			catalogName = pipeline.EventLog.Catalog
		} else {
			catalogName = pipeline.Catalog
		}
		if pipeline.EventLog.Schema != "" {
			schemaName = pipeline.EventLog.Schema
		} else {
			schemaName = pipeline.Schema
		}
		if catalogName != "" && schemaName != "" && pipeline.EventLog.Name != "" {
			ic.Emit(&resource{
				Resource: "databricks_sql_table",
				ID:       catalogName + "." + schemaName + "." + pipeline.EventLog.Name,
			})
		}
	}
	if pipeline.RunAs != nil {
		if pipeline.RunAs.UserName != "" {
			ic.Emit(&resource{
				Resource:  "databricks_user",
				Attribute: "user_name",
				Value:     pipeline.RunAs.UserName,
			})
		}
		if pipeline.RunAs.ServicePrincipalName != "" {
			ic.Emit(&resource{
				Resource:  "databricks_service_principal",
				Attribute: "application_id",
				Value:     pipeline.RunAs.ServicePrincipalName,
			})
		}
	}
	if pipeline.Environment != nil {
		for _, dep := range pipeline.Environment.Dependencies {
			emitEnvironmentDependency(ic, dep)
		}
	}
	return nil
}
