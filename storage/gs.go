package storage

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/terraform-provider-databricks/clusters"
	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// GSMount describes the object for a GS mount using google service account
type GSMount struct {
	BucketName     string `json:"bucket_name" tf:"force_new"`
	ServiceAccount string `json:"service_account,omitempty" tf:"force_new"`
}

// Source ...
func (m GSMount) Source(_ *common.DatabricksClient) string {
	return fmt.Sprintf("gs://%s", m.BucketName)
}

func (m GSMount) Name() string {
	return m.BucketName
}

func (m GSMount) ValidateAndApplyDefaults(d *schema.ResourceData, client *common.DatabricksClient) error {
	nm := d.Get("name").(string)
	if nm != "" {
		return nil
	}
	nm = m.Name()
	if nm != "" {
		d.Set("name", nm)
		return nil
	}
	return fmt.Errorf("'name' is not detected & it's impossible to infer it")
}

// Config ...
func (m GSMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

func preprocessGsMount(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m any) error {
	var gm GenericMount
	common.DataToStructPointer(d, s, &gm)
	if !(strings.HasPrefix(gm.URI, "gs://") || gm.Gs != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	serviceAccount := ""
	if gm.Gs != nil {
		serviceAccount = gm.Gs.ServiceAccount
	}
	return createOrValidateClusterForGoogleStorage(ctx, m, d, clusterID, serviceAccount)
}

func createOrValidateClusterForGoogleStorage(ctx context.Context, m any,
	d *schema.ResourceData, clusterID, serviceAccount string) error {
	clustersAPI := clusters.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if apierr.IsMissing(err) {
			cluster, err := GetOrCreateMountingClusterWithGcpServiceAccount(clustersAPI, serviceAccount)
			if err != nil {
				return fmt.Errorf("cannot re-create mounting cluster: %w", err)
			}
			return d.Set("cluster_id", cluster.ClusterID)
		}
		if err != nil {
			return fmt.Errorf("cannot get mounting cluster: %w", err)
		}
		if clusterInfo.GcpAttributes == nil || len(clusterInfo.GcpAttributes.GoogleServiceAccount) == 0 {
			return fmt.Errorf("cluster %s must have GCP service account attached", clusterID)
		}
		return nil
	} else if serviceAccount != "" {
		cluster, err := GetOrCreateMountingClusterWithGcpServiceAccount(clustersAPI, serviceAccount)
		if err != nil {
			return fmt.Errorf("cannot create mounting cluster: %w", err)
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return fmt.Errorf("either cluster_id or service_account must be specified to mount GCS bucket")
}

// GetOrCreateMountingClusterWithGcpServiceAccount ...
func GetOrCreateMountingClusterWithGcpServiceAccount(
	clustersAPI clusters.ClustersAPI, serviceAccount string) (i clusters.ClusterInfo, err error) {
	clusterName := fmt.Sprintf("terraform-mount-gcs-%x", md5.Sum([]byte(serviceAccount)))
	cluster := getCommonClusterObject(clustersAPI, clusterName)
	cluster.GcpAttributes = &clusters.GcpAttributes{GoogleServiceAccount: serviceAccount}
	return clustersAPI.GetOrCreateRunningCluster(clusterName, cluster)
}
