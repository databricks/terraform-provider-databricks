package storage

import (
	"context"
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/databrickslabs/terraform-provider-databricks/common"
	"github.com/databrickslabs/terraform-provider-databricks/compute"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// AWSIamMount describes the object for a aws mount using iam role
type S3IamMount struct {
	BucketName      string `json:"bucket_name" tf:"force_new"`
	InstanceProfile string `json:"instance_profile,omitempty" tf:"force_new"`
}

// Source ...
func (m S3IamMount) Source() string {
	return fmt.Sprintf("s3a://%s", m.BucketName)
}

// Config ...
func (m S3IamMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

// GCSMount describes the object for a GCS mount using google service account
type GCSMount struct {
	BucketName     string `json:"bucket_name" tf:"force_new"`
	ServiceAccount string `json:"service_account,omitempty" tf:"force_new"`
}

// Source ...
func (m GCSMount) Source() string {
	return fmt.Sprintf("gs://%s", m.BucketName)
}

// Config ...
func (m GCSMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

// TODO: add support for encryption parameters in S3
type GenericMount struct {
	URI     string              `json:"uri,omitempty" tf:"force_new"`
	Options map[string]string   `json:"extra_configs,omitempty" tf:"force_new"`
	Abfss   *AzureADLSGen2Mount `json:"abfss,omitempty"`
	S3      *S3IamMount         `json:"s3,omitempty"`
	Adl     *AzureADLSGen1Mount `json:"adl,omitempty"`
	Wasbs   *AzureBlobMount     `json:"wasbs,omitempty"`
	Gcs     *GCSMount           `json:"gcs,omitempty"`

	ClusterID string `json:"cluster_id,omitempty" tf:"computed"`
	MountName string `json:"mount_name" tf:"force_new"`
}

// Source returns URI backing the mount
func (m GenericMount) Source() string {
	if m.Abfss != nil {
		return m.Abfss.Source()
	} else if m.Adl != nil {
		m.Adl.Source()
	} else if m.Wasbs != nil {
		return m.Wasbs.Source()
	} else if m.S3 != nil {
		return m.S3.Source()
	} else if m.Gcs != nil {
		return m.Gcs.Source()
	}
	return m.URI
}

// Config returns mount configurations
func (m GenericMount) Config(client *common.DatabricksClient) map[string]string {
	if m.Abfss != nil {
		return m.Abfss.Config(client)
	} else if m.Adl != nil {
		m.Adl.Config(client)
	} else if m.Wasbs != nil {
		return m.Wasbs.Config(client)
	} else if m.S3 != nil {
		return m.S3.Config(client)
	} else if m.Gcs != nil {
		return m.Gcs.Config(client)
	}
	return m.Options
}

func extractFieldNames(elem interface{}) []string {
	m := elem.(*schema.Resource)
	keys := make([]string, 0, len(m.Schema))
	for k := range m.Schema {
		keys = append(keys, k)
	}
	return keys
}

func preprocessS3MountGeneric(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, s, &gm); err != nil {
		return err
	}
	if !(strings.HasPrefix(gm.URI, "s3://") || strings.HasPrefix(gm.URI, "s3a://") || gm.S3 != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	instanceProfile := ""
	if gm.S3 != nil {
		instanceProfile = gm.S3.InstanceProfile
	}
	if clusterID == "" && instanceProfile == "" {
		return fmt.Errorf("either cluster_id or instance_profile must be specified to mount S3 bucket")
	}
	clustersAPI := compute.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.AwsAttributes == nil {
			return fmt.Errorf("cluster %s must have AWS attributes", clusterID)
		}
		if len(clusterInfo.AwsAttributes.InstanceProfileArn) == 0 {
			return fmt.Errorf("cluster %s must have EC2 instance profile attached", clusterID)
		}
	}
	if instanceProfile != "" {
		cluster, err := GetOrCreateMountingClusterWithInstanceProfile(clustersAPI, instanceProfile)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}

func preprocessGcsMount(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, s, &gm); err != nil {
		return err
	}
	if !(strings.HasPrefix(gm.URI, "gs://") || gm.Gcs != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	serviceAccount := ""
	if gm.Gcs != nil {
		serviceAccount = gm.Gcs.ServiceAccount
	}
	if clusterID == "" && serviceAccount == "" {
		return fmt.Errorf("either cluster_id or service_account must be specified to mount GCS bucket")
	}
	clustersAPI := compute.NewClustersAPI(ctx, m)
	if clusterID != "" {
		clusterInfo, err := clustersAPI.Get(clusterID)
		if err != nil {
			return err
		}
		if clusterInfo.GcpAttributes == nil {
			return fmt.Errorf("cluster %s must have GCP attributes", clusterID)
		}
		if len(clusterInfo.GcpAttributes.GoogleServiceAccount) == 0 {
			return fmt.Errorf("cluster %s must have GCP service account attached", clusterID)
		}
	}
	if serviceAccount != "" {
		cluster, err := GetOrCreateMountingClusterWithGcpServiceAccount(clustersAPI, serviceAccount)
		if err != nil {
			return err
		}
		return d.Set("cluster_id", cluster.ClusterID)
	}
	return nil
}

// GetOrCreateMountingClusterWithInstanceProfile ...
func GetOrCreateMountingClusterWithGcpServiceAccount(
	clustersAPI compute.ClustersAPI, serviceAccount string) (i compute.ClusterInfo, err error) {
	clusterName := fmt.Sprintf("terraform-mount-gcs-%x", md5.Sum([]byte(serviceAccount)))
	cluster := getCommonClusterObject(clustersAPI, clusterName)
	cluster.GcpAttributes = &compute.GcpAttributes{GoogleServiceAccount: serviceAccount}
	return clustersAPI.GetOrCreateRunningCluster(clusterName, cluster)
}

// ResourceDatabricksMount mounts using given configuration
func ResourceDatabricksMount() *schema.Resource {
	tpl := GenericMount{}
	scm := common.StructToSchema(tpl, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		s["source"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}

		s["uri"].ConflictsWith = []string{"abfss", "wasbs", "s3", "adl", "gcs"}
		s["extra_configs"].ConflictsWith = []string{"abfss", "wasbs", "s3", "adl", "gcs"}
		s["abfss"].ConflictsWith = []string{"uri", "extra_configs", "wasbs", "s3", "adl", "gcs"}
		s["wasbs"].ConflictsWith = []string{"uri", "extra_configs", "abfss", "s3", "adl", "gcs"}
		s["s3"].ConflictsWith = []string{"uri", "extra_configs", "wasbs", "abfss", "adl", "gcs"}
		s["adl"].ConflictsWith = []string{"uri", "extra_configs", "wasbs", "s3", "abfss", "gcs"}
		blocks := []string{"abfss", "wasbs", "s3", "adl", "gcs"}
		// TODO: remove this when depricating other mount types (will require to add `force_new` annotation to structs as well)
		for _, nm := range blocks {
			s[nm].DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc(nm)
			s[nm].ForceNew = true
			for _, field := range extractFieldNames(s[nm].Elem) {
				p, err := common.SchemaPath(s, nm, field)
				if err == nil {
					p.ForceNew = true
				}
			}
		}
		// TODO: We need to have a validation function that will check that source isn't empty if other blocks aren't specified

		return s
	})

	r := commonMountResource(tpl, scm)
	r.CreateContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		// TODO: convert data into struct here & pass instead of converting in function itself? it would be required for GS & others
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGcsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountCreate(tpl, r)(ctx, d, m)
	}
	r.ReadContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGcsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountRead(tpl, r)(ctx, d, m)
	}
	r.DeleteContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGcsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountDelete(tpl, r)(ctx, d, m)
	}
	return r
}
