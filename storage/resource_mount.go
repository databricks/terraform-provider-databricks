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

// GSMount describes the object for a GS mount using google service account
type GSMount struct {
	BucketName     string `json:"bucket_name" tf:"force_new"`
	ServiceAccount string `json:"service_account,omitempty" tf:"force_new"`
}

// Source ...
func (m GSMount) Source() string {
	return fmt.Sprintf("gs://%s", m.BucketName)
}

// Config ...
func (m GSMount) Config(client *common.DatabricksClient) map[string]string {
	return make(map[string]string) // return empty map so nil map does not marshal to null
}

// TODO: add support for encryption parameters in S3
type GenericMount struct {
	URI     string              `json:"uri,omitempty" tf:"force_new"`
	Options map[string]string   `json:"extra_configs,omitempty" tf:"force_new"`
	Abfs    *AzureADLSGen2Mount `json:"abfs,omitempty" tf:"force_new"`
	S3      *S3IamMount         `json:"s3,omitempty" tf:"force_new"`
	Adl     *AzureADLSGen1Mount `json:"adl,omitempty" tf:"force_new"`
	Wasb    *AzureBlobMount     `json:"wasb,omitempty" tf:"force_new"`
	Gs      *GSMount            `json:"gs,omitempty" tf:"force_new"`

	ClusterID string `json:"cluster_id,omitempty" tf:"computed,force_new"`
	MountName string `json:"mount_name" tf:"force_new"`
}

// Source returns URI backing the mount
func (m GenericMount) Source() string {
	if m.Abfs != nil {
		return m.Abfs.Source()
	} else if m.Adl != nil {
		return m.Adl.Source()
	} else if m.Wasb != nil {
		return m.Wasb.Source()
	} else if m.S3 != nil {
		return m.S3.Source()
	} else if m.Gs != nil {
		return m.Gs.Source()
	}
	return m.URI
}

// Config returns mount configurations
func (m GenericMount) Config(client *common.DatabricksClient) map[string]string {
	if m.Abfs != nil {
		return m.Abfs.Config(client)
	} else if m.Adl != nil {
		return m.Adl.Config(client)
	} else if m.Wasb != nil {
		return m.Wasb.Config(client)
	} else if m.S3 != nil {
		return m.S3.Config(client)
	} else if m.Gs != nil {
		return m.Gs.Config(client)
	}
	return m.Options
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

func preprocessGsMount(ctx context.Context, s map[string]*schema.Schema, d *schema.ResourceData, m interface{}) error {
	var gm GenericMount
	if err := common.DataToStructPointer(d, s, &gm); err != nil {
		return err
	}
	if !(strings.HasPrefix(gm.URI, "gs://") || gm.Gs != nil) {
		return nil
	}
	clusterID := gm.ClusterID
	serviceAccount := ""
	if gm.Gs != nil {
		serviceAccount = gm.Gs.ServiceAccount
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

		s["uri"].ConflictsWith = []string{"abfs", "wasb", "s3", "adl", "gs"}
		s["extra_configs"].ConflictsWith = []string{"abfs", "wasb", "s3", "adl", "gs"}
		s["abfs"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "adl", "gs"}
		s["wasb"].ConflictsWith = []string{"uri", "extra_configs", "abfs", "s3", "adl", "gs"}
		s["s3"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "abfs", "adl", "gs"}
		s["adl"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "abfs", "gs"}
		s["gs"].ConflictsWith = []string{"uri", "extra_configs", "wasb", "s3", "abfs", "adl"}
		blocks := []string{"abfs", "wasb", "s3", "adl", "gs"}
		for _, nm := range blocks {
			s[nm].DiffSuppressFunc = common.MakeEmptyBlockSuppressFunc(nm)
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
		if err := preprocessGsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountCreate(tpl, r)(ctx, d, m)
	}
	r.ReadContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountRead(tpl, r)(ctx, d, m)
	}
	r.DeleteContext = func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		if err := preprocessS3MountGeneric(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		if err := preprocessGsMount(ctx, scm, d, m); err != nil {
			return diag.FromErr(err)
		}
		return mountDelete(tpl, r)(ctx, d, m)
	}
	return r
}
