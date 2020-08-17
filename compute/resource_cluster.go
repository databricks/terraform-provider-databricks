package compute

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/databrickslabs/databricks-terraform/common"
	"github.com/databrickslabs/databricks-terraform/internal"
)

var clusterSchema = resourceClusterSchema()

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 2,
		Create:        resourceClusterCreate,
		Read:          resourceClusterRead,
		Update:        resourceClusterUpdate,
		Delete:        resourceClusterDelete,
		Schema:        clusterSchema,
		// see usage at https://github.com/terraform-providers/terraform-provider-aws/blob/9900515b28b413505e5dd7c8d8ea258c59515f32/aws/resource_aws_vpc_peering_connection.go#L288
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func addMavenExclusions(scm *schema.Schema) {
	resource, ok := scm.Elem.(*schema.Resource)
	if !ok {
		log.Printf("[DEBUG] invalid elem not a resource, unable to wrap maven exclusions to resource")
		return
	}
	resource.Schema["exclusions"] = &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func librarySchema(dims ...string) *schema.Schema {
	fields := map[string]*schema.Schema{
		"messages": {
			// consider removing it...
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
	for _, dim := range dims {
		fields[dim] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		}
	}
	return &schema.Schema{
		Deprecated: "`library_*` blocks are deprecated and will be removed in v0.3. Please use more generic `library` block",
		Type:       schema.TypeSet,
		Optional:   true,
		ConfigMode: schema.SchemaConfigModeAttr,
		Elem: &schema.Resource{
			Schema: fields,
		},
	}
}

func resourceClusterSchema() map[string]*schema.Schema {
	return internal.StructToSchema(Cluster{}, func(s map[string]*schema.Schema) map[string]*schema.Schema {
		// adds `libraries` configuration block
		s["libraries"] = internal.StructToSchema(ClusterLibraryList{},
			func(ss map[string]*schema.Schema) map[string]*schema.Schema {
				return ss
			})["libraries"]

		p, err := internal.SchemaPath(s, "docker_image", "basic_auth", "password")
		if err == nil {
			p.Sensitive = true
		}
		s["autotermination_minutes"].Default = 60
		s["idempotency_token"].ForceNew = true
		s["cluster_id"] = &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		}
		s["state"] = &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		}
		s["default_tags"] = &schema.Schema{
			Type:     schema.TypeMap,
			Computed: true,
		}
		// legacy library configuration blocks
		s["library_jar"] = librarySchema("path")
		s["library_egg"] = librarySchema("path")
		s["library_whl"] = librarySchema("path")
		s["library_pypi"] = librarySchema("package", "repo")
		s["library_cran"] = librarySchema("package", "repo")
		s["library_maven"] = librarySchema("coordinates", "repo")
		addMavenExclusions(s["library_maven"])
		return s
	})
}

func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusters := NewClustersAPI(client)
	var cluster Cluster
	err := internal.DataToStructPointer(d, clusterSchema, &cluster)
	if err != nil {
		return err
	}
	modifyClusterRequest(&cluster)
	clusterInfo, err := clusters.Create(cluster)
	if err != nil {
		return err
	}
	d.SetId(clusterInfo.ClusterID)
	err = d.Set("cluster_id", clusterInfo.ClusterID)
	if err != nil {
		return err
	}
	var libraryList ClusterLibraryList
	err = internal.DataToStructPointer(d, clusterSchema, &libraryList)
	if err != nil {
		return err
	}
	if len(libraryList.Libraries) == 0 {
		// LEGACY support
		libraryList = legacyReadLibraryListFromData(d)
	}
	if len(libraryList.Libraries) > 0 {
		err = NewLibrariesAPI(client).Install(libraryList)
		if err != nil {
			return err
		}
		_, err := waitForLibrariesInstalled(NewLibrariesAPI(client), d.Id())
		if err != nil {
			// do we consider failed installation as missing cluster?...
			return err
		}
	}
	return resourceClusterRead(d, m)
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusterInfo, err := NewClustersAPI(client).Get(d.Id())
	if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
		log.Printf("Missing cluster with id: %s.", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}
	err = internal.StructToData(clusterInfo, clusterSchema, d)
	if err != nil {
		return err
	}
	libsClusterStatus, err := waitForLibrariesInstalled(NewLibrariesAPI(client), d.Id())
	if err != nil {
		// do we consider failed installation as missing cluster?...
		return err
	}
	libList := libsClusterStatus.ToLibraryList()
	return internal.StructToData(libList, clusterSchema, d)
}

func waitForLibrariesInstalled(
	libraries LibrariesAPI, clusterID string) (result *ClusterLibraryStatuses, err error) {
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		libsClusterStatus, err := libraries.ClusterStatus(clusterID)
		if ae, ok := err.(common.APIError); ok && ae.IsMissing() {
			// eventual consistency error
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		retry, err := libsClusterStatus.IsRetryNeeded()
		if retry {
			return resource.RetryableError(err)
		}
		if err != nil {
			return resource.NonRetryableError(err)
		}
		// TODO: let this be reviewed by someone. looks okay, though... KTEST!
		result = &libsClusterStatus
		return nil
	})
	return
}

func legacyReadLibraryListFromData(d *schema.ResourceData) (cll ClusterLibraryList) {
	for _, n := range []string{"library_jar", "library_egg",
		"library_whl", "library_pypi", "library_maven",
		"library_cran"} {
		libs, ok := d.GetOk(n)
		if !ok {
			continue
		}
		for _, l := range libs.(*schema.Set).List() {
			cll.AddLibraryFromMap(n, l.(map[string]interface{}))
		}
	}
	return
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.DatabricksClient)
	clusters := NewClustersAPI(client)
	clusterID := d.Id()
	cluster := Cluster{ClusterID: clusterID}
	err := internal.DataToStructPointer(d, clusterSchema, &cluster)
	if err != nil {
		return err
	}
	modifyClusterRequest(&cluster)
	clusterInfo, err := clusters.Edit(cluster)
	if err != nil {
		return err
	}
	var libraryList ClusterLibraryList
	err = internal.DataToStructPointer(d, clusterSchema, &libraryList)
	if err != nil {
		return err
	}
	if len(libraryList.Libraries) == 0 {
		// LEGACY support
		libraryList = legacyReadLibraryListFromData(d)
	}
	libraries := NewLibrariesAPI(client)
	libsClusterStatus, err := libraries.ClusterStatus(clusterID)
	if err != nil {
		return err
	}
	libraryList.ClusterID = clusterID
	libsToInstall, libsToUninstall := libraryList.Diff(libsClusterStatus)
	if len(libsToUninstall.Libraries) > 0 || len(libsToInstall.Libraries) > 0 {
		if !clusterInfo.IsRunningOrResizing() {
			err = clusters.Start(clusterID)
			if err != nil {
				return err
			}
		}
		err := updateLibraries(libraries, libsToInstall, libsToUninstall)
		if err != nil {
			return err
		}
		if clusterInfo.State == ClusterStateTerminated {
			log.Printf("[INFO] %s was in TERMINATED state, so terminating it again", clusterID)
			err = clusters.Terminate(clusterID)
		}
		if err != nil {
			return err
		}
	}
	return resourceClusterRead(d, m)
}

// modifyClusterRequest helps remove all request fields that should not be submitted when instance pool is selected.
func modifyClusterRequest(clusterModel *Cluster) {
	// Instance profile id does not exist or not set
	if clusterModel.InstancePoolID == "" {
		return
	}
	if clusterModel.AwsAttributes != nil {
		// Reset AwsAttributes
		awsAttributes := AwsAttributes{
			InstanceProfileArn: clusterModel.AwsAttributes.InstanceProfileArn,
		}
		clusterModel.AwsAttributes = &awsAttributes
	}
	clusterModel.EnableElasticDisk = false
	clusterModel.NodeTypeID = ""
	clusterModel.DriverNodeTypeID = ""
}

func updateLibraries(libraries LibrariesAPI, libsToInstall, libsToUninstall ClusterLibraryList) error {
	if len(libsToUninstall.Libraries) > 0 {
		err := libraries.Uninstall(libsToUninstall)
		if err != nil {
			return err
		}
	}
	if len(libsToInstall.Libraries) > 0 {
		err := libraries.Install(libsToInstall)
		if err != nil {
			return err
		}
	}
	_, err := waitForLibrariesInstalled(libraries, libsToUninstall.ClusterID)
	return err
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	return ClustersAPI{m.(*common.DatabricksClient)}.PermanentDelete(d.Id())
}
