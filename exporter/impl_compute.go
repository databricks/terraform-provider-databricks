package exporter

import (
	"log"
	"strings"

	sdk_compute "github.com/databricks/databricks-sdk-go/service/compute"
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
