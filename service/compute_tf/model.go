// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.
/*
These generated types are for terraform plugin framework to interact with the terraform state conveniently.

These types follow the same structure as the types in go-sdk.
The only difference is that the primitive types are no longer using the go-native types, but with tfsdk types.
Plus the json tags get converted into tfsdk tags.
We use go-native types for lists and maps intentionally for the ease for converting these types into the go-sdk types.
*/

package compute_tf

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AddInstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn types.String `tfsdk:"iam_role_arn" tf:"optional"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile" tf:"optional"`
	// By default, Databricks validates that it has sufficient permissions to
	// launch instances with the instance profile. This validation uses AWS
	// dry-run mode for the RunInstances API. If validation fails with an error
	// message that does not indicate an IAM related permission issue, (e.g.
	// “Your requested instance type is not supported in your requested
	// availability zone”), you can pass this flag to skip the validation and
	// forcibly add the instance profile.
	SkipValidation types.Bool `tfsdk:"skip_validation" tf:"optional"`
}

type AddResponse struct {
}

type Adlsgen2Info struct {
	// abfss destination, e.g.
	// `abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>`.
	Destination types.String `tfsdk:"destination" tf:""`
}

type AutoScale struct {
	// The maximum number of workers to which the cluster can scale up when
	// overloaded. Note that `max_workers` must be strictly greater than
	// `min_workers`.
	MaxWorkers types.Int64 `tfsdk:"max_workers" tf:"optional"`
	// The minimum number of workers to which the cluster can scale down when
	// underutilized. It is also the initial number of workers the cluster will
	// have after creation.
	MinWorkers types.Int64 `tfsdk:"min_workers" tf:"optional"`
}

type AwsAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones.
	//
	// Note: If `first_on_demand` is zero, this availability type will be used
	// for the entire cluster.
	Availability AwsAvailability `tfsdk:"availability" tf:"optional"`
	// The number of volumes launched for each instance. Users can choose up to
	// 10 volumes. This feature is only enabled for supported node types. Legacy
	// node types cannot specify custom EBS volumes. For node types with no
	// instance store, at least one EBS volume needs to be specified; otherwise,
	// cluster creation will fail.
	//
	// These EBS volumes will be mounted at `/ebs0`, `/ebs1`, and etc. Instance
	// store volumes will be mounted at `/local_disk0`, `/local_disk1`, and etc.
	//
	// If EBS volumes are attached, Databricks will configure Spark to use only
	// the EBS volumes for scratch storage because heterogenously sized scratch
	// devices can lead to inefficient disk utilization. If no EBS volumes are
	// attached, Databricks will configure Spark to use instance store volumes.
	//
	// Please note that if EBS volumes are specified, then the Spark
	// configuration `spark.local.dir` will be overridden.
	EbsVolumeCount types.Int64 `tfsdk:"ebs_volume_count" tf:"optional"`
	// If using gp3 volumes, what IOPS to use for the disk. If this is not set,
	// the maximum performance of a gp2 volume with the same volume size will be
	// used.
	EbsVolumeIops types.Int64 `tfsdk:"ebs_volume_iops" tf:"optional"`
	// The size of each EBS volume (in GiB) launched for each instance. For
	// general purpose SSD, this value must be within the range 100 - 4096. For
	// throughput optimized HDD, this value must be within the range 500 - 4096.
	EbsVolumeSize types.Int64 `tfsdk:"ebs_volume_size" tf:"optional"`
	// If using gp3 volumes, what throughput to use for the disk. If this is not
	// set, the maximum performance of a gp2 volume with the same volume size
	// will be used.
	EbsVolumeThroughput types.Int64 `tfsdk:"ebs_volume_throughput" tf:"optional"`
	// The type of EBS volumes that will be launched with this cluster.
	EbsVolumeType EbsVolumeType `tfsdk:"ebs_volume_type" tf:"optional"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. If this value is greater than 0, the cluster driver
	// node in particular will be placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand" tf:"optional"`
	// Nodes for this cluster will only be placed on AWS instances with this
	// instance profile. If ommitted, nodes will be placed on instances without
	// an IAM instance profile. The instance profile must have previously been
	// added to the Databricks environment by an account administrator.
	//
	// This feature may only be available to certain customer plans.
	//
	// If this field is ommitted, we will pull in the default from the conf if
	// it exists.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:"optional"`
	// The bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	//
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidPricePercent and
	// CommonConf.maxSpotBidPricePercent.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent" tf:"optional"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. If the zone specified is "auto", will try to place cluster
	// in a zone with high availability, and will retry placement in a different
	// AZ if there is not enough capacity. The list of available zones as well
	// as the default value can be found by using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones.
//
// Note: If `first_on_demand` is zero, this availability type will be used for
// the entire cluster.
type AwsAvailability string

const AwsAvailabilityOnDemand AwsAvailability = `ON_DEMAND`

const AwsAvailabilitySpot AwsAvailability = `SPOT`

const AwsAvailabilitySpotWithFallback AwsAvailability = `SPOT_WITH_FALLBACK`

// String representation for [fmt.Print]
func (f *AwsAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AwsAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND`, `SPOT`, `SPOT_WITH_FALLBACK`:
		*f = AwsAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND", "SPOT", "SPOT_WITH_FALLBACK"`, v)
	}
}

// Type always returns AwsAvailability to satisfy [pflag.Value] interface
func (f *AwsAvailability) Type() string {
	return "AwsAvailability"
}

type AzureAttributes struct {
	// Availability type used for all subsequent nodes past the
	// `first_on_demand` ones. Note: If `first_on_demand` is zero (which only
	// happens on pool clusters), this availability type will be used for the
	// entire cluster.
	Availability AzureAvailability `tfsdk:"availability" tf:"optional"`
	// The first `first_on_demand` nodes of the cluster will be placed on
	// on-demand instances. This value should be greater than 0, to make sure
	// the cluster driver node is placed on an on-demand instance. If this value
	// is greater than or equal to the current cluster size, all nodes will be
	// placed on on-demand instances. If this value is less than the current
	// cluster size, `first_on_demand` nodes will be placed on on-demand
	// instances and the remainder will be placed on `availability` instances.
	// Note that this value does not affect cluster size and cannot currently be
	// mutated over the lifetime of a cluster.
	FirstOnDemand types.Int64 `tfsdk:"first_on_demand" tf:"optional"`
	// Defines values necessary to configure and run Azure Log Analytics agent
	LogAnalyticsInfo *LogAnalyticsInfo `tfsdk:"log_analytics_info" tf:"optional"`
	// The max bid price to be used for Azure spot instances. The Max price for
	// the bid cannot be higher than the on-demand price of the instance. If not
	// specified, the default value is -1, which specifies that the instance
	// cannot be evicted on the basis of price, and only on the basis of
	// availability. Further, the value should > 0 or -1.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price" tf:"optional"`
}

// Availability type used for all subsequent nodes past the `first_on_demand`
// ones. Note: If `first_on_demand` is zero (which only happens on pool
// clusters), this availability type will be used for the entire cluster.
type AzureAvailability string

const AzureAvailabilityOnDemandAzure AzureAvailability = `ON_DEMAND_AZURE`

const AzureAvailabilitySpotAzure AzureAvailability = `SPOT_AZURE`

const AzureAvailabilitySpotWithFallbackAzure AzureAvailability = `SPOT_WITH_FALLBACK_AZURE`

// String representation for [fmt.Print]
func (f *AzureAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AzureAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_AZURE`, `SPOT_AZURE`, `SPOT_WITH_FALLBACK_AZURE`:
		*f = AzureAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_AZURE", "SPOT_AZURE", "SPOT_WITH_FALLBACK_AZURE"`, v)
	}
}

// Type always returns AzureAvailability to satisfy [pflag.Value] interface
func (f *AzureAvailability) Type() string {
	return "AzureAvailability"
}

type CancelCommand struct {
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`

	CommandId types.String `tfsdk:"commandId" tf:"optional"`

	ContextId types.String `tfsdk:"contextId" tf:"optional"`
}

type CancelResponse struct {
}

type ChangeClusterOwner struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// New owner of the cluster_id after this RPC.
	OwnerUsername types.String `tfsdk:"owner_username" tf:""`
}

type ChangeClusterOwnerResponse struct {
}

type ClientsTypes struct {
	// With jobs set, the cluster can be used for jobs
	Jobs types.Bool `tfsdk:"jobs" tf:"optional"`
	// With notebooks set, this cluster can be used for notebooks
	Notebooks types.Bool `tfsdk:"notebooks" tf:"optional"`
}

type CloneCluster struct {
	// The cluster that is being cloned.
	SourceClusterId types.String `tfsdk:"source_cluster_id" tf:""`
}

type CloudProviderNodeInfo struct {
	Status []CloudProviderNodeStatus `tfsdk:"status" tf:"optional"`
}

type CloudProviderNodeStatus string

const CloudProviderNodeStatusNotAvailableInRegion CloudProviderNodeStatus = `NotAvailableInRegion`

const CloudProviderNodeStatusNotEnabledOnSubscription CloudProviderNodeStatus = `NotEnabledOnSubscription`

// String representation for [fmt.Print]
func (f *CloudProviderNodeStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CloudProviderNodeStatus) Set(v string) error {
	switch v {
	case `NotAvailableInRegion`, `NotEnabledOnSubscription`:
		*f = CloudProviderNodeStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NotAvailableInRegion", "NotEnabledOnSubscription"`, v)
	}
}

// Type always returns CloudProviderNodeStatus to satisfy [pflag.Value] interface
func (f *CloudProviderNodeStatus) Type() string {
	return "CloudProviderNodeStatus"
}

type ClusterAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type ClusterAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type ClusterAttributes struct {
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode DataSecurityMode `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage *DockerImage `tfsdk:"docker_image" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
	// unspecified, the runtime engine is inferred from spark_version.
	RuntimeEngine RuntimeEngine `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType *WorkloadType `tfsdk:"workload_type" tf:"optional"`
}

type ClusterDetails struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// Number of CPU cores available for this cluster. Note that this can be
	// fractional, e.g. 7.5 cores, since certain node types are configured to
	// share cores between Spark nodes on the same instance.
	ClusterCores types.Float64 `tfsdk:"cluster_cores" tf:"optional"`
	// Canonical identifier for the cluster. This id is retained during cluster
	// restarts and resizes, while each new cluster has a globally unique id.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Cluster log delivery status.
	ClusterLogStatus *LogSyncStatus `tfsdk:"cluster_log_status" tf:"optional"`
	// Total amount of cluster memory, in megabytes
	ClusterMemoryMb types.Int64 `tfsdk:"cluster_memory_mb" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Determines whether the cluster was created by a user through the UI,
	// created by the Databricks Jobs Scheduler, or through an API request. This
	// is the same as cluster_creator, but read only.
	ClusterSource ClusterSource `tfsdk:"cluster_source" tf:"optional"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode DataSecurityMode `tfsdk:"data_security_mode" tf:"optional"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - Creator: <username_of_creator>
	//
	// - ClusterName: <name_of_cluster>
	//
	// - ClusterId: <id_of_cluster>
	//
	// - Name: <Databricks internal use>
	DefaultTags map[string]types.String `tfsdk:"default_tags" tf:"optional"`

	DockerImage *DockerImage `tfsdk:"docker_image" tf:"optional"`
	// Node on which the Spark driver resides. The driver node contains the
	// Spark master and the Databricks application that manages the per-notebook
	// Spark REPLs.
	Driver *SparkNode `tfsdk:"driver" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Nodes on which the Spark executors reside.
	Executors []SparkNode `tfsdk:"executors" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// Port on which Spark JDBC server is listening, in the driver nod. No
	// service will be listeningon on this port in executor nodes.
	JdbcPort types.Int64 `tfsdk:"jdbc_port" tf:"optional"`
	// the timestamp that the cluster was started/restarted
	LastRestartedTime types.Int64 `tfsdk:"last_restarted_time" tf:"optional"`
	// Time when the cluster driver last lost its state (due to a restart or
	// driver failure).
	LastStateLossTime types.Int64 `tfsdk:"last_state_loss_time" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
	// unspecified, the runtime engine is inferred from spark_version.
	RuntimeEngine RuntimeEngine `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// A canonical SparkContext identifier. This value *does* change when the
	// Spark driver restarts. The pair `(cluster_id, spark_context_id)` is a
	// globally unique identifier over all Spark contexts.
	SparkContextId types.Int64 `tfsdk:"spark_context_id" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:"optional"`
	// `spec` contains a snapshot of the field values that were used to create
	// or edit this cluster. The contents of `spec` can be used in the body of a
	// create cluster request. This field might not be populated for older
	// clusters. Note: not included in the response of the ListClusters API.
	Spec *ClusterSpec `tfsdk:"spec" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`
	// Time (in epoch milliseconds) when the cluster creation request was
	// received (when the cluster entered a `PENDING` state).
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
	// Current state of the cluster.
	State State `tfsdk:"state" tf:"optional"`
	// A message associated with the most recent state transition (e.g., the
	// reason why the cluster entered a `TERMINATED` state).
	StateMessage types.String `tfsdk:"state_message" tf:"optional"`
	// Time (in epoch milliseconds) when the cluster was terminated, if
	// applicable.
	TerminatedTime types.Int64 `tfsdk:"terminated_time" tf:"optional"`
	// Information about why the cluster was terminated. This field only appears
	// when the cluster is in a `TERMINATING` or `TERMINATED` state.
	TerminationReason *TerminationReason `tfsdk:"termination_reason" tf:"optional"`

	WorkloadType *WorkloadType `tfsdk:"workload_type" tf:"optional"`
}

type ClusterEvent struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// <needs content added>
	DataPlaneEventDetails *DataPlaneEventDetails `tfsdk:"data_plane_event_details" tf:"optional"`
	// <needs content added>
	Details *EventDetails `tfsdk:"details" tf:"optional"`
	// The timestamp when the event occurred, stored as the number of
	// milliseconds since the Unix epoch. If not provided, this will be assigned
	// by the Timeline service.
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`

	Type EventType `tfsdk:"type" tf:"optional"`
}

type ClusterLibraryStatuses struct {
	// Unique identifier for the cluster.
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
	// Status of all libraries on the cluster.
	LibraryStatuses []LibraryFullStatus `tfsdk:"library_statuses" tf:"optional"`
}

type ClusterLogConf struct {
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo `tfsdk:"dbfs" tf:"optional"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 *S3StorageInfo `tfsdk:"s3" tf:"optional"`
}

type ClusterPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type ClusterPermissionLevel string

const ClusterPermissionLevelCanAttachTo ClusterPermissionLevel = `CAN_ATTACH_TO`

const ClusterPermissionLevelCanManage ClusterPermissionLevel = `CAN_MANAGE`

const ClusterPermissionLevelCanRestart ClusterPermissionLevel = `CAN_RESTART`

// String representation for [fmt.Print]
func (f *ClusterPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_MANAGE`, `CAN_RESTART`:
		*f = ClusterPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_MANAGE", "CAN_RESTART"`, v)
	}
}

// Type always returns ClusterPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPermissionLevel) Type() string {
	return "ClusterPermissionLevel"
}

type ClusterPermissions struct {
	AccessControlList []ClusterAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type ClusterPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type ClusterPermissionsRequest struct {
	AccessControlList []ClusterAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-" url:"-"`
}

type ClusterPolicyAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type ClusterPolicyAccessControlResponse struct {
	// All permissions.
	AllPermissions []ClusterPolicyPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type ClusterPolicyPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type ClusterPolicyPermissionLevel string

const ClusterPolicyPermissionLevelCanUse ClusterPolicyPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *ClusterPolicyPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterPolicyPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = ClusterPolicyPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns ClusterPolicyPermissionLevel to satisfy [pflag.Value] interface
func (f *ClusterPolicyPermissionLevel) Type() string {
	return "ClusterPolicyPermissionLevel"
}

type ClusterPolicyPermissions struct {
	AccessControlList []ClusterPolicyAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type ClusterPolicyPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel ClusterPolicyPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type ClusterPolicyPermissionsRequest struct {
	AccessControlList []ClusterPolicyAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-" url:"-"`
}

type ClusterSize struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
}

// Determines whether the cluster was created by a user through the UI, created
// by the Databricks Jobs Scheduler, or through an API request. This is the same
// as cluster_creator, but read only.
type ClusterSource string

const ClusterSourceApi ClusterSource = `API`

const ClusterSourceJob ClusterSource = `JOB`

const ClusterSourceModels ClusterSource = `MODELS`

const ClusterSourcePipeline ClusterSource = `PIPELINE`

const ClusterSourcePipelineMaintenance ClusterSource = `PIPELINE_MAINTENANCE`

const ClusterSourceSql ClusterSource = `SQL`

const ClusterSourceUi ClusterSource = `UI`

// String representation for [fmt.Print]
func (f *ClusterSource) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterSource) Set(v string) error {
	switch v {
	case `API`, `JOB`, `MODELS`, `PIPELINE`, `PIPELINE_MAINTENANCE`, `SQL`, `UI`:
		*f = ClusterSource(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "API", "JOB", "MODELS", "PIPELINE", "PIPELINE_MAINTENANCE", "SQL", "UI"`, v)
	}
}

// Type always returns ClusterSource to satisfy [pflag.Value] interface
func (f *ClusterSource) Type() string {
	return "ClusterSource"
}

type ClusterSpec struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode DataSecurityMode `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage *DockerImage `tfsdk:"docker_image" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
	// unspecified, the runtime engine is inferred from spark_version.
	RuntimeEngine RuntimeEngine `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:"optional"`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType *WorkloadType `tfsdk:"workload_type" tf:"optional"`
}

// Get status
type ClusterStatus struct {
	// Unique identifier of the cluster whose status should be retrieved.
	ClusterId types.String `tfsdk:"-" url:"cluster_id"`
}

type Command struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`
	// Executable code
	Command types.String `tfsdk:"command" tf:"optional"`
	// Running context id
	ContextId types.String `tfsdk:"contextId" tf:"optional"`

	Language Language `tfsdk:"language" tf:"optional"`
}

type CommandStatus string

const CommandStatusCancelled CommandStatus = `Cancelled`

const CommandStatusCancelling CommandStatus = `Cancelling`

const CommandStatusError CommandStatus = `Error`

const CommandStatusFinished CommandStatus = `Finished`

const CommandStatusQueued CommandStatus = `Queued`

const CommandStatusRunning CommandStatus = `Running`

// String representation for [fmt.Print]
func (f *CommandStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommandStatus) Set(v string) error {
	switch v {
	case `Cancelled`, `Cancelling`, `Error`, `Finished`, `Queued`, `Running`:
		*f = CommandStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Cancelled", "Cancelling", "Error", "Finished", "Queued", "Running"`, v)
	}
}

// Type always returns CommandStatus to satisfy [pflag.Value] interface
func (f *CommandStatus) Type() string {
	return "CommandStatus"
}

// Get command info
type CommandStatusRequest struct {
	ClusterId types.String `tfsdk:"-" url:"clusterId"`

	CommandId types.String `tfsdk:"-" url:"commandId"`

	ContextId types.String `tfsdk:"-" url:"contextId"`
}

type CommandStatusResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`

	Results *Results `tfsdk:"results" tf:"optional"`

	Status CommandStatus `tfsdk:"status" tf:"optional"`
}

type ContextStatus string

const ContextStatusError ContextStatus = `Error`

const ContextStatusPending ContextStatus = `Pending`

const ContextStatusRunning ContextStatus = `Running`

// String representation for [fmt.Print]
func (f *ContextStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ContextStatus) Set(v string) error {
	switch v {
	case `Error`, `Pending`, `Running`:
		*f = ContextStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Error", "Pending", "Running"`, v)
	}
}

// Type always returns ContextStatus to satisfy [pflag.Value] interface
func (f *ContextStatus) Type() string {
	return "ContextStatus"
}

// Get status
type ContextStatusRequest struct {
	ClusterId types.String `tfsdk:"-" url:"clusterId"`

	ContextId types.String `tfsdk:"-" url:"contextId"`
}

type ContextStatusResponse struct {
	Id types.String `tfsdk:"id" tf:"optional"`

	Status ContextStatus `tfsdk:"status" tf:"optional"`
}

type CreateCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// When specified, this clones libraries from a source cluster during the
	// creation of a new cluster.
	CloneFrom *CloneCluster `tfsdk:"clone_from" tf:"optional"`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode DataSecurityMode `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage *DockerImage `tfsdk:"docker_image" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
	// unspecified, the runtime engine is inferred from spark_version.
	RuntimeEngine RuntimeEngine `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType *WorkloadType `tfsdk:"workload_type" tf:"optional"`
}

type CreateClusterResponse struct {
	ClusterId types.String `tfsdk:"cluster_id" tf:"optional"`
}

type CreateContext struct {
	// Running cluster id
	ClusterId types.String `tfsdk:"clusterId" tf:"optional"`

	Language Language `tfsdk:"language" tf:"optional"`
}

type CreateInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec `tfsdk:"disk_spec" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:""`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []types.String `tfsdk:"preloaded_spark_versions" tf:"optional"`
}

type CreateInstancePoolResponse struct {
	// The ID of the created instance pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
}

type CreatePolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:""`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
}

type CreatePolicyResponse struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
}

type CreateResponse struct {
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
}

type Created struct {
	Id types.String `tfsdk:"id" tf:"optional"`
}

type DataPlaneEventDetails struct {
	// <needs content added>
	EventType DataPlaneEventDetailsEventType `tfsdk:"event_type" tf:"optional"`
	// <needs content added>
	ExecutorFailures types.Int64 `tfsdk:"executor_failures" tf:"optional"`
	// <needs content added>
	HostId types.String `tfsdk:"host_id" tf:"optional"`
	// <needs content added>
	Timestamp types.Int64 `tfsdk:"timestamp" tf:"optional"`
}

// <needs content added>
type DataPlaneEventDetailsEventType string

const DataPlaneEventDetailsEventTypeNodeBlacklisted DataPlaneEventDetailsEventType = `NODE_BLACKLISTED`

const DataPlaneEventDetailsEventTypeNodeExcludedDecommissioned DataPlaneEventDetailsEventType = `NODE_EXCLUDED_DECOMMISSIONED`

// String representation for [fmt.Print]
func (f *DataPlaneEventDetailsEventType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataPlaneEventDetailsEventType) Set(v string) error {
	switch v {
	case `NODE_BLACKLISTED`, `NODE_EXCLUDED_DECOMMISSIONED`:
		*f = DataPlaneEventDetailsEventType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NODE_BLACKLISTED", "NODE_EXCLUDED_DECOMMISSIONED"`, v)
	}
}

// Type always returns DataPlaneEventDetailsEventType to satisfy [pflag.Value] interface
func (f *DataPlaneEventDetailsEventType) Type() string {
	return "DataPlaneEventDetailsEventType"
}

// Data security mode decides what data governance model to use when accessing
// data from a cluster.
//
// * `NONE`: No security isolation for multiple users sharing the cluster. Data
// governance features are not available in this mode. * `SINGLE_USER`: A secure
// cluster that can only be exclusively used by a single user specified in
// `single_user_name`. Most programming languages, cluster features and data
// governance features are available in this mode. * `USER_ISOLATION`: A secure
// cluster that can be shared by multiple users. Cluster users are fully
// isolated so that they cannot see each other's data and credentials. Most data
// governance features are supported in this mode. But programming languages and
// cluster features might be limited.
//
// The following modes are deprecated starting with Databricks Runtime 15.0 and
// will be removed for future Databricks Runtime versions:
//
// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table ACL
// clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating from
// legacy Passthrough on high concurrency clusters. * `LEGACY_SINGLE_USER`: This
// mode is for users migrating from legacy Passthrough on standard clusters. *
// `LEGACY_SINGLE_USER_STANDARD`: This mode provides a way that doesn’t have
// UC nor passthrough enabled.
type DataSecurityMode string

// This mode is for users migrating from legacy Passthrough on high concurrency
// clusters.
const DataSecurityModeLegacyPassthrough DataSecurityMode = `LEGACY_PASSTHROUGH`

// This mode is for users migrating from legacy Passthrough on standard
// clusters.
const DataSecurityModeLegacySingleUser DataSecurityMode = `LEGACY_SINGLE_USER`

// This mode provides a way that doesn’t have UC nor passthrough enabled.
const DataSecurityModeLegacySingleUserStandard DataSecurityMode = `LEGACY_SINGLE_USER_STANDARD`

// This mode is for users migrating from legacy Table ACL clusters.
const DataSecurityModeLegacyTableAcl DataSecurityMode = `LEGACY_TABLE_ACL`

// No security isolation for multiple users sharing the cluster. Data governance
// features are not available in this mode.
const DataSecurityModeNone DataSecurityMode = `NONE`

// A secure cluster that can only be exclusively used by a single user specified
// in `single_user_name`. Most programming languages, cluster features and data
// governance features are available in this mode.
const DataSecurityModeSingleUser DataSecurityMode = `SINGLE_USER`

// A secure cluster that can be shared by multiple users. Cluster users are
// fully isolated so that they cannot see each other's data and credentials.
// Most data governance features are supported in this mode. But programming
// languages and cluster features might be limited.
const DataSecurityModeUserIsolation DataSecurityMode = `USER_ISOLATION`

// String representation for [fmt.Print]
func (f *DataSecurityMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DataSecurityMode) Set(v string) error {
	switch v {
	case `LEGACY_PASSTHROUGH`, `LEGACY_SINGLE_USER`, `LEGACY_SINGLE_USER_STANDARD`, `LEGACY_TABLE_ACL`, `NONE`, `SINGLE_USER`, `USER_ISOLATION`:
		*f = DataSecurityMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LEGACY_PASSTHROUGH", "LEGACY_SINGLE_USER", "LEGACY_SINGLE_USER_STANDARD", "LEGACY_TABLE_ACL", "NONE", "SINGLE_USER", "USER_ISOLATION"`, v)
	}
}

// Type always returns DataSecurityMode to satisfy [pflag.Value] interface
func (f *DataSecurityMode) Type() string {
	return "DataSecurityMode"
}

type DbfsStorageInfo struct {
	// dbfs destination, e.g. `dbfs:/my/path`
	Destination types.String `tfsdk:"destination" tf:""`
}

type DeleteCluster struct {
	// The cluster to be terminated.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

type DeleteClusterResponse struct {
}

// Delete init script
type DeleteGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-" url:"-"`
}

type DeleteInstancePool struct {
	// The instance pool to be terminated.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
}

type DeleteInstancePoolResponse struct {
}

type DeletePolicy struct {
	// The ID of the policy to delete.
	PolicyId types.String `tfsdk:"policy_id" tf:""`
}

type DeletePolicyResponse struct {
}

type DeleteResponse struct {
}

type DestroyContext struct {
	ClusterId types.String `tfsdk:"clusterId" tf:""`

	ContextId types.String `tfsdk:"contextId" tf:""`
}

type DestroyResponse struct {
}

type DiskSpec struct {
	// The number of disks launched for each instance: - This feature is only
	// enabled for supported node types. - Users can choose up to the limit of
	// the disks supported by the node type. - For node types with no OS disk,
	// at least one disk must be specified; otherwise, cluster creation will
	// fail.
	//
	// If disks are attached, Databricks will configure Spark to use only the
	// disks for scratch storage, because heterogenously sized scratch devices
	// can lead to inefficient disk utilization. If no disks are attached,
	// Databricks will configure Spark to use instance store disks.
	//
	// Note: If disks are specified, then the Spark configuration
	// `spark.local.dir` will be overridden.
	//
	// Disks will be mounted at: - For AWS: `/ebs0`, `/ebs1`, and etc. - For
	// Azure: `/remote_volume0`, `/remote_volume1`, and etc.
	DiskCount types.Int64 `tfsdk:"disk_count" tf:"optional"`

	DiskIops types.Int64 `tfsdk:"disk_iops" tf:"optional"`
	// The size of each disk (in GiB) launched for each instance. Values must
	// fall into the supported range for a particular instance type.
	//
	// For AWS: - General Purpose SSD: 100 - 4096 GiB - Throughput Optimized
	// HDD: 500 - 4096 GiB
	//
	// For Azure: - Premium LRS (SSD): 1 - 1023 GiB - Standard LRS (HDD): 1-
	// 1023 GiB
	DiskSize types.Int64 `tfsdk:"disk_size" tf:"optional"`

	DiskThroughput types.Int64 `tfsdk:"disk_throughput" tf:"optional"`
	// The type of disks that will be launched with this cluster.
	DiskType *DiskType `tfsdk:"disk_type" tf:"optional"`
}

type DiskType struct {
	AzureDiskVolumeType DiskTypeAzureDiskVolumeType `tfsdk:"azure_disk_volume_type" tf:"optional"`

	EbsVolumeType DiskTypeEbsVolumeType `tfsdk:"ebs_volume_type" tf:"optional"`
}

type DiskTypeAzureDiskVolumeType string

const DiskTypeAzureDiskVolumeTypePremiumLrs DiskTypeAzureDiskVolumeType = `PREMIUM_LRS`

const DiskTypeAzureDiskVolumeTypeStandardLrs DiskTypeAzureDiskVolumeType = `STANDARD_LRS`

// String representation for [fmt.Print]
func (f *DiskTypeAzureDiskVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DiskTypeAzureDiskVolumeType) Set(v string) error {
	switch v {
	case `PREMIUM_LRS`, `STANDARD_LRS`:
		*f = DiskTypeAzureDiskVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PREMIUM_LRS", "STANDARD_LRS"`, v)
	}
}

// Type always returns DiskTypeAzureDiskVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeAzureDiskVolumeType) Type() string {
	return "DiskTypeAzureDiskVolumeType"
}

type DiskTypeEbsVolumeType string

const DiskTypeEbsVolumeTypeGeneralPurposeSsd DiskTypeEbsVolumeType = `GENERAL_PURPOSE_SSD`

const DiskTypeEbsVolumeTypeThroughputOptimizedHdd DiskTypeEbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

// String representation for [fmt.Print]
func (f *DiskTypeEbsVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DiskTypeEbsVolumeType) Set(v string) error {
	switch v {
	case `GENERAL_PURPOSE_SSD`, `THROUGHPUT_OPTIMIZED_HDD`:
		*f = DiskTypeEbsVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENERAL_PURPOSE_SSD", "THROUGHPUT_OPTIMIZED_HDD"`, v)
	}
}

// Type always returns DiskTypeEbsVolumeType to satisfy [pflag.Value] interface
func (f *DiskTypeEbsVolumeType) Type() string {
	return "DiskTypeEbsVolumeType"
}

type DockerBasicAuth struct {
	// Password of the user
	Password types.String `tfsdk:"password" tf:"optional"`
	// Name of the user
	Username types.String `tfsdk:"username" tf:"optional"`
}

type DockerImage struct {
	BasicAuth *DockerBasicAuth `tfsdk:"basic_auth" tf:"optional"`
	// URL of the docker image.
	Url types.String `tfsdk:"url" tf:"optional"`
}

// The type of EBS volumes that will be launched with this cluster.
type EbsVolumeType string

const EbsVolumeTypeGeneralPurposeSsd EbsVolumeType = `GENERAL_PURPOSE_SSD`

const EbsVolumeTypeThroughputOptimizedHdd EbsVolumeType = `THROUGHPUT_OPTIMIZED_HDD`

// String representation for [fmt.Print]
func (f *EbsVolumeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EbsVolumeType) Set(v string) error {
	switch v {
	case `GENERAL_PURPOSE_SSD`, `THROUGHPUT_OPTIMIZED_HDD`:
		*f = EbsVolumeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GENERAL_PURPOSE_SSD", "THROUGHPUT_OPTIMIZED_HDD"`, v)
	}
}

// Type always returns EbsVolumeType to satisfy [pflag.Value] interface
func (f *EbsVolumeType) Type() string {
	return "EbsVolumeType"
}

type EditCluster struct {
	// When set to true, fixed and default values from the policy will be used
	// for fields that are omitted. When set to false, only fixed values from
	// the policy will be applied.
	ApplyPolicyDefaultValues types.Bool `tfsdk:"apply_policy_default_values" tf:"optional"`
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// Automatically terminates the cluster after it is inactive for this time
	// in minutes. If not set, this cluster will not be automatically
	// terminated. If specified, the threshold must be between 10 and 10000
	// minutes. Users can also set this value to 0 to explicitly disable
	// automatic termination.
	AutoterminationMinutes types.Int64 `tfsdk:"autotermination_minutes" tf:"optional"`
	// Attributes related to clusters running on Amazon Web Services. If not
	// specified at cluster creation, a set of default values will be used.
	AwsAttributes *AwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to clusters running on Microsoft Azure. If not
	// specified at cluster creation, a set of default values will be used.
	AzureAttributes *AzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// ID of the cluser
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The configuration for delivering spark logs to a long-term storage
	// destination. Two kinds of destinations (dbfs and s3) are supported. Only
	// one destination can be specified for one cluster. If the conf is given,
	// the logs will be delivered to the destination every `5 mins`. The
	// destination of driver logs is `$destination/$clusterId/driver`, while the
	// destination of executor logs is `$destination/$clusterId/executor`.
	ClusterLogConf *ClusterLogConf `tfsdk:"cluster_log_conf" tf:"optional"`
	// Cluster name requested by the user. This doesn't have to be unique. If
	// not specified at creation, the cluster name will be an empty string.
	ClusterName types.String `tfsdk:"cluster_name" tf:"optional"`
	// Additional tags for cluster resources. Databricks will tag all cluster
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	//
	// - Clusters can only reuse cloud resources if the resources' tags are a
	// subset of the cluster tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Data security mode decides what data governance model to use when
	// accessing data from a cluster.
	//
	// * `NONE`: No security isolation for multiple users sharing the cluster.
	// Data governance features are not available in this mode. * `SINGLE_USER`:
	// A secure cluster that can only be exclusively used by a single user
	// specified in `single_user_name`. Most programming languages, cluster
	// features and data governance features are available in this mode. *
	// `USER_ISOLATION`: A secure cluster that can be shared by multiple users.
	// Cluster users are fully isolated so that they cannot see each other's
	// data and credentials. Most data governance features are supported in this
	// mode. But programming languages and cluster features might be limited.
	//
	// The following modes are deprecated starting with Databricks Runtime 15.0
	// and will be removed for future Databricks Runtime versions:
	//
	// * `LEGACY_TABLE_ACL`: This mode is for users migrating from legacy Table
	// ACL clusters. * `LEGACY_PASSTHROUGH`: This mode is for users migrating
	// from legacy Passthrough on high concurrency clusters. *
	// `LEGACY_SINGLE_USER`: This mode is for users migrating from legacy
	// Passthrough on standard clusters. * `LEGACY_SINGLE_USER_STANDARD`: This
	// mode provides a way that doesn’t have UC nor passthrough enabled.
	DataSecurityMode DataSecurityMode `tfsdk:"data_security_mode" tf:"optional"`

	DockerImage *DockerImage `tfsdk:"docker_image" tf:"optional"`
	// The optional ID of the instance pool for the driver of the cluster
	// belongs. The pool cluster uses the instance pool with id
	// (instance_pool_id) if the driver pool is not assigned.
	DriverInstancePoolId types.String `tfsdk:"driver_instance_pool_id" tf:"optional"`
	// The node type of the Spark driver. Note that this field is optional; if
	// unset, the driver node type will be set as the same value as
	// `node_type_id` defined above.
	DriverNodeTypeId types.String `tfsdk:"driver_node_type_id" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this cluster will dynamically
	// acquire additional disk space when its Spark workers are running low on
	// disk space. This feature requires specific AWS permissions to function
	// correctly - refer to the User Guide for more details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Whether to enable LUKS on cluster VMs' local disks
	EnableLocalDiskEncryption types.Bool `tfsdk:"enable_local_disk_encryption" tf:"optional"`
	// Attributes related to clusters running on Google Cloud Platform. If not
	// specified at cluster creation, a set of default values will be used.
	GcpAttributes *GcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// The configuration for storing init scripts. Any number of destinations
	// can be specified. The scripts are executed sequentially in the order
	// provided. If `cluster_log_conf` is specified, init script logs are sent
	// to `<destination>/<cluster-ID>/init_scripts`.
	InitScripts []InitScriptInfo `tfsdk:"init_scripts" tf:"optional"`
	// The optional ID of the instance pool to which the cluster belongs.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
	// The ID of the cluster policy used to create the cluster if applicable.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
	// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
	// unspecified, the runtime engine is inferred from spark_version.
	RuntimeEngine RuntimeEngine `tfsdk:"runtime_engine" tf:"optional"`
	// Single user name if data_security_mode is `SINGLE_USER`
	SingleUserName types.String `tfsdk:"single_user_name" tf:"optional"`
	// An object containing a set of optional, user-specified Spark
	// configuration key-value pairs. Users can also pass in a string of extra
	// JVM options to the driver and the executors via
	// `spark.driver.extraJavaOptions` and `spark.executor.extraJavaOptions`
	// respectively.
	SparkConf map[string]types.String `tfsdk:"spark_conf" tf:"optional"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs. Please note that key-value pair of the form
	// (X,Y) will be exported as is (i.e., `export X='Y'`) while launching the
	// driver and workers.
	//
	// In order to specify an additional set of `SPARK_DAEMON_JAVA_OPTS`, we
	// recommend appending them to `$SPARK_DAEMON_JAVA_OPTS` as shown in the
	// example below. This ensures that all default databricks managed
	// environmental variables are included as well.
	//
	// Example Spark environment variables: `{"SPARK_WORKER_MEMORY": "28000m",
	// "SPARK_LOCAL_DIRS": "/local_disk0"}` or `{"SPARK_DAEMON_JAVA_OPTS":
	// "$SPARK_DAEMON_JAVA_OPTS -Dspark.shuffle.service.enabled=true"}`
	SparkEnvVars map[string]types.String `tfsdk:"spark_env_vars" tf:"optional"`
	// The Spark version of the cluster, e.g. `3.3.x-scala2.11`. A list of
	// available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	SparkVersion types.String `tfsdk:"spark_version" tf:""`
	// SSH public key contents that will be added to each Spark node in this
	// cluster. The corresponding private keys can be used to login with the
	// user name `ubuntu` on port `2200`. Up to 10 keys can be specified.
	SshPublicKeys []types.String `tfsdk:"ssh_public_keys" tf:"optional"`

	WorkloadType *WorkloadType `tfsdk:"workload_type" tf:"optional"`
}

type EditClusterResponse struct {
}

type EditInstancePool struct {
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Instance pool ID
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:""`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
}

type EditInstancePoolResponse struct {
}

type EditPolicy struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:""`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family. The cluster policy's policy definition inherits
	// the policy family's policy definition.
	//
	// Cannot be used with `definition`. Use
	// `policy_family_definition_overrides` instead to customize the policy
	// definition.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
	// The ID of the policy to update.
	PolicyId types.String `tfsdk:"policy_id" tf:""`
}

type EditPolicyResponse struct {
}

type EditResponse struct {
}

// The environment entity used to preserve serverless environment side panel and
// jobs' environment for non-notebook task. In this minimal environment spec,
// only pip dependencies are supported.
type Environment struct {
	// Client version used by the environment The client is the user-facing
	// environment of the runtime. Each client comes with a specific set of
	// pre-installed libraries. The version is a string, consisting of the major
	// client version.
	Client types.String `tfsdk:"client" tf:""`
	// List of pip dependencies, as supported by the version of pip in this
	// environment. Each dependency is a pip requirement file line
	// https://pip.pypa.io/en/stable/reference/requirements-file-format/ Allowed
	// dependency could be <requirement specifier>, <archive url/path>, <local
	// project path>(WSFS or Volumes in Databricks), <vcs project url> E.g.
	// dependencies: ["foo==0.0.1", "-r /Workspace/test/requirements.txt"]
	Dependencies []types.String `tfsdk:"dependencies" tf:"optional"`
}

type EventDetails struct {
	// * For created clusters, the attributes of the cluster. * For edited
	// clusters, the new attributes of the cluster.
	Attributes *ClusterAttributes `tfsdk:"attributes" tf:"optional"`
	// The cause of a change in target size.
	Cause EventDetailsCause `tfsdk:"cause" tf:"optional"`
	// The actual cluster size that was set in the cluster creation or edit.
	ClusterSize *ClusterSize `tfsdk:"cluster_size" tf:"optional"`
	// The current number of vCPUs in the cluster.
	CurrentNumVcpus types.Int64 `tfsdk:"current_num_vcpus" tf:"optional"`
	// The current number of nodes in the cluster.
	CurrentNumWorkers types.Int64 `tfsdk:"current_num_workers" tf:"optional"`
	// <needs content added>
	DidNotExpandReason types.String `tfsdk:"did_not_expand_reason" tf:"optional"`
	// Current disk size in bytes
	DiskSize types.Int64 `tfsdk:"disk_size" tf:"optional"`
	// More details about the change in driver's state
	DriverStateMessage types.String `tfsdk:"driver_state_message" tf:"optional"`
	// Whether or not a blocklisted node should be terminated. For
	// ClusterEventType NODE_BLACKLISTED.
	EnableTerminationForNodeBlocklisted types.Bool `tfsdk:"enable_termination_for_node_blocklisted" tf:"optional"`
	// <needs content added>
	FreeSpace types.Int64 `tfsdk:"free_space" tf:"optional"`
	// List of global and cluster init scripts associated with this cluster
	// event.
	InitScripts *InitScriptEventDetails `tfsdk:"init_scripts" tf:"optional"`
	// Instance Id where the event originated from
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`
	// Unique identifier of the specific job run associated with this cluster
	// event * For clusters created for jobs, this will be the same as the
	// cluster name
	JobRunName types.String `tfsdk:"job_run_name" tf:"optional"`
	// The cluster attributes before a cluster was edited.
	PreviousAttributes *ClusterAttributes `tfsdk:"previous_attributes" tf:"optional"`
	// The size of the cluster before an edit or resize.
	PreviousClusterSize *ClusterSize `tfsdk:"previous_cluster_size" tf:"optional"`
	// Previous disk size in bytes
	PreviousDiskSize types.Int64 `tfsdk:"previous_disk_size" tf:"optional"`
	// A termination reason: * On a TERMINATED event, this is the reason of the
	// termination. * On a RESIZE_COMPLETE event, this indicates the reason that
	// we failed to acquire some nodes.
	Reason *TerminationReason `tfsdk:"reason" tf:"optional"`
	// The targeted number of vCPUs in the cluster.
	TargetNumVcpus types.Int64 `tfsdk:"target_num_vcpus" tf:"optional"`
	// The targeted number of nodes in the cluster.
	TargetNumWorkers types.Int64 `tfsdk:"target_num_workers" tf:"optional"`
	// The user that caused the event to occur. (Empty if it was done by the
	// control plane.)
	User types.String `tfsdk:"user" tf:"optional"`
}

// The cause of a change in target size.
type EventDetailsCause string

const EventDetailsCauseAutorecovery EventDetailsCause = `AUTORECOVERY`

const EventDetailsCauseAutoscale EventDetailsCause = `AUTOSCALE`

const EventDetailsCauseReplaceBadNodes EventDetailsCause = `REPLACE_BAD_NODES`

const EventDetailsCauseUserRequest EventDetailsCause = `USER_REQUEST`

// String representation for [fmt.Print]
func (f *EventDetailsCause) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventDetailsCause) Set(v string) error {
	switch v {
	case `AUTORECOVERY`, `AUTOSCALE`, `REPLACE_BAD_NODES`, `USER_REQUEST`:
		*f = EventDetailsCause(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTORECOVERY", "AUTOSCALE", "REPLACE_BAD_NODES", "USER_REQUEST"`, v)
	}
}

// Type always returns EventDetailsCause to satisfy [pflag.Value] interface
func (f *EventDetailsCause) Type() string {
	return "EventDetailsCause"
}

type EventType string

const EventTypeAutoscalingStatsReport EventType = `AUTOSCALING_STATS_REPORT`

const EventTypeCreating EventType = `CREATING`

const EventTypeDbfsDown EventType = `DBFS_DOWN`

const EventTypeDidNotExpandDisk EventType = `DID_NOT_EXPAND_DISK`

const EventTypeDriverHealthy EventType = `DRIVER_HEALTHY`

const EventTypeDriverNotResponding EventType = `DRIVER_NOT_RESPONDING`

const EventTypeDriverUnavailable EventType = `DRIVER_UNAVAILABLE`

const EventTypeEdited EventType = `EDITED`

const EventTypeExpandedDisk EventType = `EXPANDED_DISK`

const EventTypeFailedToExpandDisk EventType = `FAILED_TO_EXPAND_DISK`

const EventTypeInitScriptsFinished EventType = `INIT_SCRIPTS_FINISHED`

const EventTypeInitScriptsStarted EventType = `INIT_SCRIPTS_STARTED`

const EventTypeMetastoreDown EventType = `METASTORE_DOWN`

const EventTypeNodesLost EventType = `NODES_LOST`

const EventTypeNodeBlacklisted EventType = `NODE_BLACKLISTED`

const EventTypeNodeExcludedDecommissioned EventType = `NODE_EXCLUDED_DECOMMISSIONED`

const EventTypePinned EventType = `PINNED`

const EventTypeResizing EventType = `RESIZING`

const EventTypeRestarting EventType = `RESTARTING`

const EventTypeRunning EventType = `RUNNING`

const EventTypeSparkException EventType = `SPARK_EXCEPTION`

const EventTypeStarting EventType = `STARTING`

const EventTypeTerminating EventType = `TERMINATING`

const EventTypeUnpinned EventType = `UNPINNED`

const EventTypeUpsizeCompleted EventType = `UPSIZE_COMPLETED`

// String representation for [fmt.Print]
func (f *EventType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EventType) Set(v string) error {
	switch v {
	case `AUTOSCALING_STATS_REPORT`, `CREATING`, `DBFS_DOWN`, `DID_NOT_EXPAND_DISK`, `DRIVER_HEALTHY`, `DRIVER_NOT_RESPONDING`, `DRIVER_UNAVAILABLE`, `EDITED`, `EXPANDED_DISK`, `FAILED_TO_EXPAND_DISK`, `INIT_SCRIPTS_FINISHED`, `INIT_SCRIPTS_STARTED`, `METASTORE_DOWN`, `NODES_LOST`, `NODE_BLACKLISTED`, `NODE_EXCLUDED_DECOMMISSIONED`, `PINNED`, `RESIZING`, `RESTARTING`, `RUNNING`, `SPARK_EXCEPTION`, `STARTING`, `TERMINATING`, `UNPINNED`, `UPSIZE_COMPLETED`:
		*f = EventType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUTOSCALING_STATS_REPORT", "CREATING", "DBFS_DOWN", "DID_NOT_EXPAND_DISK", "DRIVER_HEALTHY", "DRIVER_NOT_RESPONDING", "DRIVER_UNAVAILABLE", "EDITED", "EXPANDED_DISK", "FAILED_TO_EXPAND_DISK", "INIT_SCRIPTS_FINISHED", "INIT_SCRIPTS_STARTED", "METASTORE_DOWN", "NODES_LOST", "NODE_BLACKLISTED", "NODE_EXCLUDED_DECOMMISSIONED", "PINNED", "RESIZING", "RESTARTING", "RUNNING", "SPARK_EXCEPTION", "STARTING", "TERMINATING", "UNPINNED", "UPSIZE_COMPLETED"`, v)
	}
}

// Type always returns EventType to satisfy [pflag.Value] interface
func (f *EventType) Type() string {
	return "EventType"
}

type GcpAttributes struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	Availability GcpAvailability `tfsdk:"availability" tf:"optional"`
	// boot disk size in GB
	BootDiskSize types.Int64 `tfsdk:"boot_disk_size" tf:"optional"`
	// If provided, the cluster will impersonate the google service account when
	// accessing gcloud services (like GCS). The google service account must
	// have previously been added to the Databricks environment by an account
	// administrator.
	GoogleServiceAccount types.String `tfsdk:"google_service_account" tf:"optional"`
	// If provided, each node (workers and driver) in the cluster will have this
	// number of local SSDs attached. Each local SSD is 375GB in size. Refer to
	// [GCP documentation] for the supported number of local SSDs for each
	// instance type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count" tf:"optional"`
	// This field determines whether the spark executors will be scheduled to
	// run on preemptible VMs (when set to true) versus standard compute engine
	// VMs (when set to false; default). Note: Soon to be deprecated, use the
	// availability field instead.
	UsePreemptibleExecutors types.Bool `tfsdk:"use_preemptible_executors" tf:"optional"`
	// Identifier for the availability zone in which the cluster resides. This
	// can be one of the following: - "HA" => High availability, spread nodes
	// across availability zones for a Databricks deployment region [default] -
	// "AUTO" => Databricks picks an availability zone to schedule the cluster
	// on. - A GCP availability zone => Pick One of the available zones for
	// (machine type + region) from
	// https://cloud.google.com/compute/docs/regions-zones.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

// This field determines whether the instance pool will contain preemptible VMs,
// on-demand VMs, or preemptible VMs with a fallback to on-demand VMs if the
// former is unavailable.
type GcpAvailability string

const GcpAvailabilityOnDemandGcp GcpAvailability = `ON_DEMAND_GCP`

const GcpAvailabilityPreemptibleGcp GcpAvailability = `PREEMPTIBLE_GCP`

const GcpAvailabilityPreemptibleWithFallbackGcp GcpAvailability = `PREEMPTIBLE_WITH_FALLBACK_GCP`

// String representation for [fmt.Print]
func (f *GcpAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GcpAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_GCP`, `PREEMPTIBLE_GCP`, `PREEMPTIBLE_WITH_FALLBACK_GCP`:
		*f = GcpAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_GCP", "PREEMPTIBLE_GCP", "PREEMPTIBLE_WITH_FALLBACK_GCP"`, v)
	}
}

// Type always returns GcpAvailability to satisfy [pflag.Value] interface
func (f *GcpAvailability) Type() string {
	return "GcpAvailability"
}

type GcsStorageInfo struct {
	// GCS destination/URI, e.g. `gs://my-bucket/some-prefix`
	Destination types.String `tfsdk:"destination" tf:""`
}

// Get cluster permission levels
type GetClusterPermissionLevelsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-" url:"-"`
}

type GetClusterPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get cluster permissions
type GetClusterPermissionsRequest struct {
	// The cluster for which to get or manage permissions.
	ClusterId types.String `tfsdk:"-" url:"-"`
}

// Get cluster policy permission levels
type GetClusterPolicyPermissionLevelsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-" url:"-"`
}

type GetClusterPolicyPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ClusterPolicyPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get cluster policy permissions
type GetClusterPolicyPermissionsRequest struct {
	// The cluster policy for which to get or manage permissions.
	ClusterPolicyId types.String `tfsdk:"-" url:"-"`
}

// Get a cluster policy
type GetClusterPolicyRequest struct {
	// Canonical unique identifier for the cluster policy.
	PolicyId types.String `tfsdk:"-" url:"policy_id"`
}

// Get cluster info
type GetClusterRequest struct {
	// The cluster about which to retrieve information.
	ClusterId types.String `tfsdk:"-" url:"cluster_id"`
}

type GetEvents struct {
	// The ID of the cluster to retrieve events about.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The end time in epoch milliseconds. If empty, returns events up to the
	// current time.
	EndTime types.Int64 `tfsdk:"end_time" tf:"optional"`
	// An optional set of event types to filter on. If empty, all event types
	// are returned.
	EventTypes []EventType `tfsdk:"event_types" tf:"optional"`
	// The maximum number of events to include in a page of events. Defaults to
	// 50, and maximum allowed value is 500.
	Limit types.Int64 `tfsdk:"limit" tf:"optional"`
	// The offset in the result set. Defaults to 0 (no offset). When an offset
	// is specified and the results are requested in descending order, the
	// end_time field is required.
	Offset types.Int64 `tfsdk:"offset" tf:"optional"`
	// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
	Order GetEventsOrder `tfsdk:"order" tf:"optional"`
	// The start time in epoch milliseconds. If empty, returns events starting
	// from the beginning of time.
	StartTime types.Int64 `tfsdk:"start_time" tf:"optional"`
}

// The order to list events in; either "ASC" or "DESC". Defaults to "DESC".
type GetEventsOrder string

const GetEventsOrderAsc GetEventsOrder = `ASC`

const GetEventsOrderDesc GetEventsOrder = `DESC`

// String representation for [fmt.Print]
func (f *GetEventsOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetEventsOrder) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*f = GetEventsOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Type always returns GetEventsOrder to satisfy [pflag.Value] interface
func (f *GetEventsOrder) Type() string {
	return "GetEventsOrder"
}

type GetEventsResponse struct {
	// <content needs to be added>
	Events []ClusterEvent `tfsdk:"events" tf:"optional"`
	// The parameters required to retrieve the next page of events. Omitted if
	// there are no more events to read.
	NextPage *GetEvents `tfsdk:"next_page" tf:"optional"`
	// The total number of events filtered by the start_time, end_time, and
	// event_types.
	TotalCount types.Int64 `tfsdk:"total_count" tf:"optional"`
}

// Get an init script
type GetGlobalInitScriptRequest struct {
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-" url:"-"`
}

type GetInstancePool struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags map[string]types.String `tfsdk:"default_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec `tfsdk:"disk_spec" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:""`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:"optional"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []types.String `tfsdk:"preloaded_spark_versions" tf:"optional"`
	// Current state of the instance pool.
	State InstancePoolState `tfsdk:"state" tf:"optional"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `tfsdk:"stats" tf:"optional"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `tfsdk:"status" tf:"optional"`
}

// Get instance pool permission levels
type GetInstancePoolPermissionLevelsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-" url:"-"`
}

type GetInstancePoolPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []InstancePoolPermissionsDescription `tfsdk:"permission_levels" tf:"optional"`
}

// Get instance pool permissions
type GetInstancePoolPermissionsRequest struct {
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-" url:"-"`
}

// Get instance pool information
type GetInstancePoolRequest struct {
	// The canonical unique identifier for the instance pool.
	InstancePoolId types.String `tfsdk:"-" url:"instance_pool_id"`
}

// Get policy family information
type GetPolicyFamilyRequest struct {
	PolicyFamilyId types.String `tfsdk:"-" url:"-"`
}

type GetSparkVersionsResponse struct {
	// All the available Spark versions.
	Versions []SparkVersion `tfsdk:"versions" tf:"optional"`
}

type GlobalInitScriptCreateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:""`
	// The position of a global init script, where 0 represents the first script
	// to run, 1 is the second script to run, in ascending order.
	//
	// If you omit the numeric position for a new global init script, it
	// defaults to last position. It will run after all current scripts. Setting
	// any value greater than the position of the last script is equivalent to
	// the last position. Example: Take three existing scripts with positions 0,
	// 1, and 2. Any position of (3) or greater puts the script in the last
	// position. If an explicit position value conflicts with an existing script
	// value, your request succeeds, but the original script at that position
	// and all later scripts have their positions incremented by 1.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:""`
}

type GlobalInitScriptDetails struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:"optional"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type GlobalInitScriptDetailsWithContent struct {
	// Time when the script was created, represented as a Unix timestamp in
	// milliseconds.
	CreatedAt types.Int64 `tfsdk:"created_at" tf:"optional"`
	// The username of the user who created the script.
	CreatedBy types.String `tfsdk:"created_by" tf:"optional"`
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:"optional"`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:"optional"`
	// The global init script ID.
	ScriptId types.String `tfsdk:"script_id" tf:"optional"`
	// Time when the script was updated, represented as a Unix timestamp in
	// milliseconds.
	UpdatedAt types.Int64 `tfsdk:"updated_at" tf:"optional"`
	// The username of the user who last updated the script
	UpdatedBy types.String `tfsdk:"updated_by" tf:"optional"`
}

type GlobalInitScriptUpdateRequest struct {
	// Specifies whether the script is enabled. The script runs only if enabled.
	Enabled types.Bool `tfsdk:"enabled" tf:"optional"`
	// The name of the script
	Name types.String `tfsdk:"name" tf:""`
	// The position of a script, where 0 represents the first script to run, 1
	// is the second script to run, in ascending order. To move the script to
	// run first, set its position to 0.
	//
	// To move the script to the end, set its position to any value greater or
	// equal to the position of the last script. Example, three existing scripts
	// with positions 0, 1, and 2. Any position value of 2 or greater puts the
	// script in the last position (2).
	//
	// If an explicit position value conflicts with an existing script, your
	// request succeeds, but the original script at that position and all later
	// scripts have their positions incremented by 1.
	Position types.Int64 `tfsdk:"position" tf:"optional"`
	// The Base64-encoded content of the script.
	Script types.String `tfsdk:"script" tf:""`
	// The ID of the global init script.
	ScriptId types.String `tfsdk:"-" url:"-"`
}

type InitScriptEventDetails struct {
	// The cluster scoped init scripts associated with this cluster event
	Cluster []InitScriptInfoAndExecutionDetails `tfsdk:"cluster" tf:"optional"`
	// The global init scripts associated with this cluster event
	Global []InitScriptInfoAndExecutionDetails `tfsdk:"global" tf:"optional"`
	// The private ip address of the node where the init scripts were run.
	ReportedForNode types.String `tfsdk:"reported_for_node" tf:"optional"`
}

type InitScriptExecutionDetails struct {
	// Addition details regarding errors.
	ErrorMessage types.String `tfsdk:"error_message" tf:"optional"`
	// The duration of the script execution in seconds.
	ExecutionDurationSeconds types.Int64 `tfsdk:"execution_duration_seconds" tf:"optional"`
	// The current status of the script
	Status InitScriptExecutionDetailsStatus `tfsdk:"status" tf:"optional"`
}

// The current status of the script
type InitScriptExecutionDetailsStatus string

const InitScriptExecutionDetailsStatusFailedExecution InitScriptExecutionDetailsStatus = `FAILED_EXECUTION`

const InitScriptExecutionDetailsStatusFailedFetch InitScriptExecutionDetailsStatus = `FAILED_FETCH`

const InitScriptExecutionDetailsStatusNotExecuted InitScriptExecutionDetailsStatus = `NOT_EXECUTED`

const InitScriptExecutionDetailsStatusSkipped InitScriptExecutionDetailsStatus = `SKIPPED`

const InitScriptExecutionDetailsStatusSucceeded InitScriptExecutionDetailsStatus = `SUCCEEDED`

const InitScriptExecutionDetailsStatusUnknown InitScriptExecutionDetailsStatus = `UNKNOWN`

// String representation for [fmt.Print]
func (f *InitScriptExecutionDetailsStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InitScriptExecutionDetailsStatus) Set(v string) error {
	switch v {
	case `FAILED_EXECUTION`, `FAILED_FETCH`, `NOT_EXECUTED`, `SKIPPED`, `SUCCEEDED`, `UNKNOWN`:
		*f = InitScriptExecutionDetailsStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_EXECUTION", "FAILED_FETCH", "NOT_EXECUTED", "SKIPPED", "SUCCEEDED", "UNKNOWN"`, v)
	}
}

// Type always returns InitScriptExecutionDetailsStatus to satisfy [pflag.Value] interface
func (f *InitScriptExecutionDetailsStatus) Type() string {
	return "InitScriptExecutionDetailsStatus"
}

type InitScriptInfo struct {
	// destination needs to be provided. e.g. `{ "abfss" : { "destination" :
	// "abfss://<container-name>@<storage-account-name>.dfs.core.windows.net/<directory-name>"
	// } }
	Abfss *Adlsgen2Info `tfsdk:"abfss" tf:"optional"`
	// destination needs to be provided. e.g. `{ "dbfs" : { "destination" :
	// "dbfs:/home/cluster_log" } }`
	Dbfs *DbfsStorageInfo `tfsdk:"dbfs" tf:"optional"`
	// destination needs to be provided. e.g. `{ "file" : { "destination" :
	// "file:/my/local/file.sh" } }`
	File *LocalFileInfo `tfsdk:"file" tf:"optional"`
	// destination needs to be provided. e.g. `{ "gcs": { "destination":
	// "gs://my-bucket/file.sh" } }`
	Gcs *GcsStorageInfo `tfsdk:"gcs" tf:"optional"`
	// destination and either the region or endpoint need to be provided. e.g.
	// `{ "s3": { "destination" : "s3://cluster_log_bucket/prefix", "region" :
	// "us-west-2" } }` Cluster iam role is used to access s3, please make sure
	// the cluster iam role in `instance_profile_arn` has permission to write
	// data to the s3 destination.
	S3 *S3StorageInfo `tfsdk:"s3" tf:"optional"`
	// destination needs to be provided. e.g. `{ "volumes" : { "destination" :
	// "/Volumes/my-init.sh" } }`
	Volumes *VolumesStorageInfo `tfsdk:"volumes" tf:"optional"`
	// destination needs to be provided. e.g. `{ "workspace" : { "destination" :
	// "/Users/user1@databricks.com/my-init.sh" } }`
	Workspace *WorkspaceStorageInfo `tfsdk:"workspace" tf:"optional"`
}

type InitScriptInfoAndExecutionDetails struct {
	// Details about the script
	ExecutionDetails *InitScriptExecutionDetails `tfsdk:"execution_details" tf:"optional"`
	// The script
	Script *InitScriptInfo `tfsdk:"script" tf:"optional"`
}

type InstallLibraries struct {
	// Unique identifier for the cluster on which to install these libraries.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The libraries to install.
	Libraries []Library `tfsdk:"libraries" tf:""`
}

type InstallLibrariesResponse struct {
}

type InstancePoolAccessControlRequest struct {
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `tfsdk:"permission_level" tf:"optional"`
	// application ID of a service principal
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type InstancePoolAccessControlResponse struct {
	// All permissions.
	AllPermissions []InstancePoolPermission `tfsdk:"all_permissions" tf:"optional"`
	// Display name of the user or service principal.
	DisplayName types.String `tfsdk:"display_name" tf:"optional"`
	// name of the group
	GroupName types.String `tfsdk:"group_name" tf:"optional"`
	// Name of the service principal.
	ServicePrincipalName types.String `tfsdk:"service_principal_name" tf:"optional"`
	// name of the user
	UserName types.String `tfsdk:"user_name" tf:"optional"`
}

type InstancePoolAndStats struct {
	// Attributes related to instance pools running on Amazon Web Services. If
	// not specified at pool creation, a set of default values will be used.
	AwsAttributes *InstancePoolAwsAttributes `tfsdk:"aws_attributes" tf:"optional"`
	// Attributes related to instance pools running on Azure. If not specified
	// at pool creation, a set of default values will be used.
	AzureAttributes *InstancePoolAzureAttributes `tfsdk:"azure_attributes" tf:"optional"`
	// Additional tags for pool resources. Databricks will tag all pool
	// resources (e.g., AWS instances and EBS volumes) with these tags in
	// addition to `default_tags`. Notes:
	//
	// - Currently, Databricks allows at most 45 custom tags
	CustomTags map[string]types.String `tfsdk:"custom_tags" tf:"optional"`
	// Tags that are added by Databricks regardless of any `custom_tags`,
	// including:
	//
	// - Vendor: Databricks
	//
	// - InstancePoolCreator: <user_id_of_creator>
	//
	// - InstancePoolName: <name_of_pool>
	//
	// - InstancePoolId: <id_of_pool>
	DefaultTags map[string]types.String `tfsdk:"default_tags" tf:"optional"`
	// Defines the specification of the disks that will be attached to all spark
	// containers.
	DiskSpec *DiskSpec `tfsdk:"disk_spec" tf:"optional"`
	// Autoscaling Local Storage: when enabled, this instances in this pool will
	// dynamically acquire additional disk space when its Spark workers are
	// running low on disk space. In AWS, this feature requires specific AWS
	// permissions to function correctly - refer to the User Guide for more
	// details.
	EnableElasticDisk types.Bool `tfsdk:"enable_elastic_disk" tf:"optional"`
	// Attributes related to instance pools running on Google Cloud Platform. If
	// not specified at pool creation, a set of default values will be used.
	GcpAttributes *InstancePoolGcpAttributes `tfsdk:"gcp_attributes" tf:"optional"`
	// Automatically terminates the extra instances in the pool cache after they
	// are inactive for this time in minutes if min_idle_instances requirement
	// is already met. If not set, the extra pool instances will be
	// automatically terminated after a default timeout. If specified, the
	// threshold must be between 0 and 10000 minutes. Users can also set this
	// value to 0 to instantly remove idle instances from the cache if min cache
	// size could still hold.
	IdleInstanceAutoterminationMinutes types.Int64 `tfsdk:"idle_instance_autotermination_minutes" tf:"optional"`
	// Canonical unique identifier for the pool.
	InstancePoolId types.String `tfsdk:"instance_pool_id" tf:"optional"`
	// Pool name requested by the user. Pool name must be unique. Length must be
	// between 1 and 100 characters.
	InstancePoolName types.String `tfsdk:"instance_pool_name" tf:"optional"`
	// Maximum number of outstanding instances to keep in the pool, including
	// both instances used by clusters and idle instances. Clusters that require
	// further instance provisioning will fail during upsize requests.
	MaxCapacity types.Int64 `tfsdk:"max_capacity" tf:"optional"`
	// Minimum number of idle instances to keep in the instance pool
	MinIdleInstances types.Int64 `tfsdk:"min_idle_instances" tf:"optional"`
	// This field encodes, through a single value, the resources available to
	// each of the Spark nodes in this cluster. For example, the Spark nodes can
	// be provisioned and optimized for memory or compute intensive workloads. A
	// list of available node types can be retrieved by using the
	// :method:clusters/listNodeTypes API call.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:"optional"`
	// Custom Docker Image BYOC
	PreloadedDockerImages []DockerImage `tfsdk:"preloaded_docker_images" tf:"optional"`
	// A list containing at most one preloaded Spark image version for the pool.
	// Pool-backed clusters started with the preloaded Spark version will start
	// faster. A list of available Spark versions can be retrieved by using the
	// :method:clusters/sparkVersions API call.
	PreloadedSparkVersions []types.String `tfsdk:"preloaded_spark_versions" tf:"optional"`
	// Current state of the instance pool.
	State InstancePoolState `tfsdk:"state" tf:"optional"`
	// Usage statistics about the instance pool.
	Stats *InstancePoolStats `tfsdk:"stats" tf:"optional"`
	// Status of failed pending instances in the pool.
	Status *InstancePoolStatus `tfsdk:"status" tf:"optional"`
}

type InstancePoolAwsAttributes struct {
	// Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAwsAvailability
	Availability InstancePoolAwsAttributesAvailability `tfsdk:"availability" tf:"optional"`
	// Calculates the bid price for AWS spot instances, as a percentage of the
	// corresponding instance type's on-demand price. For example, if this field
	// is set to 50, and the cluster needs a new `r3.xlarge` spot instance, then
	// the bid price is half of the price of on-demand `r3.xlarge` instances.
	// Similarly, if this field is set to 200, the bid price is twice the price
	// of on-demand `r3.xlarge` instances. If not specified, the default value
	// is 100. When spot instances are requested for this cluster, only spot
	// instances whose bid price percentage matches this field will be
	// considered. Note that, for safety, we enforce this field to be no more
	// than 10000.
	//
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidPricePercent and
	// CommonConf.maxSpotBidPricePercent.
	SpotBidPricePercent types.Int64 `tfsdk:"spot_bid_price_percent" tf:"optional"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west-2a". The provided
	// availability zone must be in the same region as the Databricks
	// deployment. For example, "us-west-2a" is not a valid zone id if the
	// Databricks deployment resides in the "us-east-1" region. This is an
	// optional field at cluster creation, and if not specified, a default zone
	// will be used. The list of available zones as well as the default value
	// can be found by using the `List Zones` method.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

// Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAwsAvailability
type InstancePoolAwsAttributesAvailability string

const InstancePoolAwsAttributesAvailabilityOnDemand InstancePoolAwsAttributesAvailability = `ON_DEMAND`

const InstancePoolAwsAttributesAvailabilitySpot InstancePoolAwsAttributesAvailability = `SPOT`

// String representation for [fmt.Print]
func (f *InstancePoolAwsAttributesAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolAwsAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND`, `SPOT`:
		*f = InstancePoolAwsAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND", "SPOT"`, v)
	}
}

// Type always returns InstancePoolAwsAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAwsAttributesAvailability) Type() string {
	return "InstancePoolAwsAttributesAvailability"
}

type InstancePoolAzureAttributes struct {
	// Shows the Availability type used for the spot nodes.
	//
	// The default value is defined by
	// InstancePoolConf.instancePoolDefaultAzureAvailability
	Availability InstancePoolAzureAttributesAvailability `tfsdk:"availability" tf:"optional"`
	// The default value and documentation here should be kept consistent with
	// CommonConf.defaultSpotBidMaxPrice.
	SpotBidMaxPrice types.Float64 `tfsdk:"spot_bid_max_price" tf:"optional"`
}

// Shows the Availability type used for the spot nodes.
//
// The default value is defined by
// InstancePoolConf.instancePoolDefaultAzureAvailability
type InstancePoolAzureAttributesAvailability string

const InstancePoolAzureAttributesAvailabilityOnDemandAzure InstancePoolAzureAttributesAvailability = `ON_DEMAND_AZURE`

const InstancePoolAzureAttributesAvailabilitySpotAzure InstancePoolAzureAttributesAvailability = `SPOT_AZURE`

// String representation for [fmt.Print]
func (f *InstancePoolAzureAttributesAvailability) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolAzureAttributesAvailability) Set(v string) error {
	switch v {
	case `ON_DEMAND_AZURE`, `SPOT_AZURE`:
		*f = InstancePoolAzureAttributesAvailability(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ON_DEMAND_AZURE", "SPOT_AZURE"`, v)
	}
}

// Type always returns InstancePoolAzureAttributesAvailability to satisfy [pflag.Value] interface
func (f *InstancePoolAzureAttributesAvailability) Type() string {
	return "InstancePoolAzureAttributesAvailability"
}

type InstancePoolGcpAttributes struct {
	// This field determines whether the instance pool will contain preemptible
	// VMs, on-demand VMs, or preemptible VMs with a fallback to on-demand VMs
	// if the former is unavailable.
	GcpAvailability GcpAvailability `tfsdk:"gcp_availability" tf:"optional"`
	// If provided, each node in the instance pool will have this number of
	// local SSDs attached. Each local SSD is 375GB in size. Refer to [GCP
	// documentation] for the supported number of local SSDs for each instance
	// type.
	//
	// [GCP documentation]: https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	LocalSsdCount types.Int64 `tfsdk:"local_ssd_count" tf:"optional"`
	// Identifier for the availability zone/datacenter in which the cluster
	// resides. This string will be of a form like "us-west1-a". The provided
	// availability zone must be in the same region as the Databricks workspace.
	// For example, "us-west1-a" is not a valid zone id if the Databricks
	// workspace resides in the "us-east1" region. This is an optional field at
	// instance pool creation, and if not specified, a default zone will be
	// used.
	//
	// This field can be one of the following: - "HA" => High availability,
	// spread nodes across availability zones for a Databricks deployment region
	// - A GCP availability zone => Pick One of the available zones for (machine
	// type + region) from https://cloud.google.com/compute/docs/regions-zones
	// (e.g. "us-west1-a").
	//
	// If empty, Databricks picks an availability zone to schedule the cluster
	// on.
	ZoneId types.String `tfsdk:"zone_id" tf:"optional"`
}

type InstancePoolPermission struct {
	Inherited types.Bool `tfsdk:"inherited" tf:"optional"`

	InheritedFromObject []types.String `tfsdk:"inherited_from_object" tf:"optional"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

// Permission level
type InstancePoolPermissionLevel string

const InstancePoolPermissionLevelCanAttachTo InstancePoolPermissionLevel = `CAN_ATTACH_TO`

const InstancePoolPermissionLevelCanManage InstancePoolPermissionLevel = `CAN_MANAGE`

// String representation for [fmt.Print]
func (f *InstancePoolPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_ATTACH_TO`, `CAN_MANAGE`:
		*f = InstancePoolPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_ATTACH_TO", "CAN_MANAGE"`, v)
	}
}

// Type always returns InstancePoolPermissionLevel to satisfy [pflag.Value] interface
func (f *InstancePoolPermissionLevel) Type() string {
	return "InstancePoolPermissionLevel"
}

type InstancePoolPermissions struct {
	AccessControlList []InstancePoolAccessControlResponse `tfsdk:"access_control_list" tf:"optional"`

	ObjectId types.String `tfsdk:"object_id" tf:"optional"`

	ObjectType types.String `tfsdk:"object_type" tf:"optional"`
}

type InstancePoolPermissionsDescription struct {
	Description types.String `tfsdk:"description" tf:"optional"`
	// Permission level
	PermissionLevel InstancePoolPermissionLevel `tfsdk:"permission_level" tf:"optional"`
}

type InstancePoolPermissionsRequest struct {
	AccessControlList []InstancePoolAccessControlRequest `tfsdk:"access_control_list" tf:"optional"`
	// The instance pool for which to get or manage permissions.
	InstancePoolId types.String `tfsdk:"-" url:"-"`
}

// Current state of the instance pool.
type InstancePoolState string

const InstancePoolStateActive InstancePoolState = `ACTIVE`

const InstancePoolStateDeleted InstancePoolState = `DELETED`

const InstancePoolStateStopped InstancePoolState = `STOPPED`

// String representation for [fmt.Print]
func (f *InstancePoolState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *InstancePoolState) Set(v string) error {
	switch v {
	case `ACTIVE`, `DELETED`, `STOPPED`:
		*f = InstancePoolState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DELETED", "STOPPED"`, v)
	}
}

// Type always returns InstancePoolState to satisfy [pflag.Value] interface
func (f *InstancePoolState) Type() string {
	return "InstancePoolState"
}

type InstancePoolStats struct {
	// Number of active instances in the pool that are NOT part of a cluster.
	IdleCount types.Int64 `tfsdk:"idle_count" tf:"optional"`
	// Number of pending instances in the pool that are NOT part of a cluster.
	PendingIdleCount types.Int64 `tfsdk:"pending_idle_count" tf:"optional"`
	// Number of pending instances in the pool that are part of a cluster.
	PendingUsedCount types.Int64 `tfsdk:"pending_used_count" tf:"optional"`
	// Number of active instances in the pool that are part of a cluster.
	UsedCount types.Int64 `tfsdk:"used_count" tf:"optional"`
}

type InstancePoolStatus struct {
	// List of error messages for the failed pending instances. The
	// pending_instance_errors follows FIFO with maximum length of the min_idle
	// of the pool. The pending_instance_errors is emptied once the number of
	// exiting available instances reaches the min_idle of the pool.
	PendingInstanceErrors []PendingInstanceError `tfsdk:"pending_instance_errors" tf:"optional"`
}

type InstanceProfile struct {
	// The AWS IAM role ARN of the role associated with the instance profile.
	// This field is required if your role name and instance profile name do not
	// match and you want to use the instance profile with [Databricks SQL
	// Serverless].
	//
	// Otherwise, this field is optional.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	IamRoleArn types.String `tfsdk:"iam_role_arn" tf:"optional"`
	// The AWS ARN of the instance profile to register with Databricks. This
	// field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
	// Boolean flag indicating whether the instance profile should only be used
	// in credential passthrough scenarios. If true, it means the instance
	// profile contains an meta IAM role which could assume a wide range of
	// roles. Therefore it should always be used with authorization. This field
	// is optional, the default value is `false`.
	IsMetaInstanceProfile types.Bool `tfsdk:"is_meta_instance_profile" tf:"optional"`
}

type Language string

const LanguagePython Language = `python`

const LanguageScala Language = `scala`

const LanguageSql Language = `sql`

// String representation for [fmt.Print]
func (f *Language) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Language) Set(v string) error {
	switch v {
	case `python`, `scala`, `sql`:
		*f = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "python", "scala", "sql"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (f *Language) Type() string {
	return "Language"
}

type Library struct {
	// Specification of a CRAN library to be installed as part of the library
	Cran *RCranLibrary `tfsdk:"cran" tf:"optional"`
	// URI of the egg library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "egg":
	// "/Workspace/path/to/library.egg" }`, `{ "egg" :
	// "/Volumes/path/to/library.egg" }` or `{ "egg":
	// "s3://my-bucket/library.egg" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Egg types.String `tfsdk:"egg" tf:"optional"`
	// URI of the JAR library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "jar":
	// "/Workspace/path/to/library.jar" }`, `{ "jar" :
	// "/Volumes/path/to/library.jar" }` or `{ "jar":
	// "s3://my-bucket/library.jar" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Jar types.String `tfsdk:"jar" tf:"optional"`
	// Specification of a maven library to be installed. For example: `{
	// "coordinates": "org.jsoup:jsoup:1.7.2" }`
	Maven *MavenLibrary `tfsdk:"maven" tf:"optional"`
	// Specification of a PyPi library to be installed. For example: `{
	// "package": "simplejson" }`
	Pypi *PythonPyPiLibrary `tfsdk:"pypi" tf:"optional"`
	// URI of the requirements.txt file to install. Only Workspace paths and
	// Unity Catalog Volumes paths are supported. For example: `{
	// "requirements": "/Workspace/path/to/requirements.txt" }` or `{
	// "requirements" : "/Volumes/path/to/requirements.txt" }`
	Requirements types.String `tfsdk:"requirements" tf:"optional"`
	// URI of the wheel library to install. Supported URIs include Workspace
	// paths, Unity Catalog Volumes paths, and S3 URIs. For example: `{ "whl":
	// "/Workspace/path/to/library.whl" }`, `{ "whl" :
	// "/Volumes/path/to/library.whl" }` or `{ "whl":
	// "s3://my-bucket/library.whl" }`. If S3 is used, please make sure the
	// cluster has read access on the library. You may need to launch the
	// cluster with an IAM role to access the S3 URI.
	Whl types.String `tfsdk:"whl" tf:"optional"`
}

// The status of the library on a specific cluster.
type LibraryFullStatus struct {
	// Whether the library was set to be installed on all clusters via the
	// libraries UI.
	IsLibraryForAllClusters types.Bool `tfsdk:"is_library_for_all_clusters" tf:"optional"`
	// Unique identifier for the library.
	Library *Library `tfsdk:"library" tf:"optional"`
	// All the info and warning messages that have occurred so far for this
	// library.
	Messages []types.String `tfsdk:"messages" tf:"optional"`
	// Status of installing the library on the cluster.
	Status LibraryInstallStatus `tfsdk:"status" tf:"optional"`
}

// The status of a library on a specific cluster.
type LibraryInstallStatus string

const LibraryInstallStatusFailed LibraryInstallStatus = `FAILED`

const LibraryInstallStatusInstalled LibraryInstallStatus = `INSTALLED`

const LibraryInstallStatusInstalling LibraryInstallStatus = `INSTALLING`

const LibraryInstallStatusPending LibraryInstallStatus = `PENDING`

const LibraryInstallStatusResolving LibraryInstallStatus = `RESOLVING`

const LibraryInstallStatusRestored LibraryInstallStatus = `RESTORED`

const LibraryInstallStatusSkipped LibraryInstallStatus = `SKIPPED`

const LibraryInstallStatusUninstallOnRestart LibraryInstallStatus = `UNINSTALL_ON_RESTART`

// String representation for [fmt.Print]
func (f *LibraryInstallStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LibraryInstallStatus) Set(v string) error {
	switch v {
	case `FAILED`, `INSTALLED`, `INSTALLING`, `PENDING`, `RESOLVING`, `RESTORED`, `SKIPPED`, `UNINSTALL_ON_RESTART`:
		*f = LibraryInstallStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "INSTALLED", "INSTALLING", "PENDING", "RESOLVING", "RESTORED", "SKIPPED", "UNINSTALL_ON_RESTART"`, v)
	}
}

// Type always returns LibraryInstallStatus to satisfy [pflag.Value] interface
func (f *LibraryInstallStatus) Type() string {
	return "LibraryInstallStatus"
}

type ListAllClusterLibraryStatusesResponse struct {
	// A list of cluster statuses.
	Statuses []ClusterLibraryStatuses `tfsdk:"statuses" tf:"optional"`
}

type ListAvailableZonesResponse struct {
	// The availability zone if no `zone_id` is provided in the cluster creation
	// request.
	DefaultZone types.String `tfsdk:"default_zone" tf:"optional"`
	// The list of available zones (e.g., ['us-west-2c', 'us-east-2']).
	Zones []types.String `tfsdk:"zones" tf:"optional"`
}

// List cluster policies
type ListClusterPoliciesRequest struct {
	// The cluster policy attribute to sort by. * `POLICY_CREATION_TIME` - Sort
	// result list by policy creation time. * `POLICY_NAME` - Sort result list
	// by policy name.
	SortColumn ListSortColumn `tfsdk:"-" url:"sort_column,omitempty"`
	// The order in which the policies get listed. * `DESC` - Sort result list
	// in descending order. * `ASC` - Sort result list in ascending order.
	SortOrder ListSortOrder `tfsdk:"-" url:"sort_order,omitempty"`
}

// List all clusters
type ListClustersRequest struct {
	// Filter clusters based on what type of client it can be used for. Could be
	// either NOTEBOOKS or JOBS. No input for this field will get all clusters
	// in the workspace without filtering on its supported client
	CanUseClient types.String `tfsdk:"-" url:"can_use_client,omitempty"`
}

type ListClustersResponse struct {
	// <needs content added>
	Clusters []ClusterDetails `tfsdk:"clusters" tf:"optional"`
}

type ListGlobalInitScriptsResponse struct {
	Scripts []GlobalInitScriptDetails `tfsdk:"scripts" tf:"optional"`
}

type ListInstancePools struct {
	InstancePools []InstancePoolAndStats `tfsdk:"instance_pools" tf:"optional"`
}

type ListInstanceProfilesResponse struct {
	// A list of instance profiles that the user can access.
	InstanceProfiles []InstanceProfile `tfsdk:"instance_profiles" tf:"optional"`
}

type ListNodeTypesResponse struct {
	// The list of available Spark node types.
	NodeTypes []NodeType `tfsdk:"node_types" tf:"optional"`
}

type ListPoliciesResponse struct {
	// List of policies.
	Policies []Policy `tfsdk:"policies" tf:"optional"`
}

// List policy families
type ListPolicyFamiliesRequest struct {
	// The max number of policy families to return.
	MaxResults types.Int64 `tfsdk:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken types.String `tfsdk:"-" url:"page_token,omitempty"`
}

type ListPolicyFamiliesResponse struct {
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken types.String `tfsdk:"next_page_token" tf:"optional"`
	// List of policy families.
	PolicyFamilies []PolicyFamily `tfsdk:"policy_families" tf:""`
}

type ListSortColumn string

const ListSortColumnPolicyCreationTime ListSortColumn = `POLICY_CREATION_TIME`

const ListSortColumnPolicyName ListSortColumn = `POLICY_NAME`

// String representation for [fmt.Print]
func (f *ListSortColumn) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortColumn) Set(v string) error {
	switch v {
	case `POLICY_CREATION_TIME`, `POLICY_NAME`:
		*f = ListSortColumn(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_CREATION_TIME", "POLICY_NAME"`, v)
	}
}

// Type always returns ListSortColumn to satisfy [pflag.Value] interface
func (f *ListSortColumn) Type() string {
	return "ListSortColumn"
}

type ListSortOrder string

const ListSortOrderAsc ListSortOrder = `ASC`

const ListSortOrderDesc ListSortOrder = `DESC`

// String representation for [fmt.Print]
func (f *ListSortOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListSortOrder) Set(v string) error {
	switch v {
	case `ASC`, `DESC`:
		*f = ListSortOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ASC", "DESC"`, v)
	}
}

// Type always returns ListSortOrder to satisfy [pflag.Value] interface
func (f *ListSortOrder) Type() string {
	return "ListSortOrder"
}

type LocalFileInfo struct {
	// local file destination, e.g. `file:/my/local/file.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}

type LogAnalyticsInfo struct {
	// <needs content added>
	LogAnalyticsPrimaryKey types.String `tfsdk:"log_analytics_primary_key" tf:"optional"`
	// <needs content added>
	LogAnalyticsWorkspaceId types.String `tfsdk:"log_analytics_workspace_id" tf:"optional"`
}

type LogSyncStatus struct {
	// The timestamp of last attempt. If the last attempt fails,
	// `last_exception` will contain the exception in the last attempt.
	LastAttempted types.Int64 `tfsdk:"last_attempted" tf:"optional"`
	// The exception thrown in the last attempt, it would be null (omitted in
	// the response) if there is no exception in last attempted.
	LastException types.String `tfsdk:"last_exception" tf:"optional"`
}

type MavenLibrary struct {
	// Gradle-style maven coordinates. For example: "org.jsoup:jsoup:1.7.2".
	Coordinates types.String `tfsdk:"coordinates" tf:""`
	// List of dependences to exclude. For example: `["slf4j:slf4j",
	// "*:hadoop-client"]`.
	//
	// Maven dependency exclusions:
	// https://maven.apache.org/guides/introduction/introduction-to-optional-and-excludes-dependencies.html.
	Exclusions []types.String `tfsdk:"exclusions" tf:"optional"`
	// Maven repo to install the Maven package from. If omitted, both Maven
	// Central Repository and Spark Packages are searched.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

type NodeInstanceType struct {
	InstanceTypeId types.String `tfsdk:"instance_type_id" tf:"optional"`

	LocalDiskSizeGb types.Int64 `tfsdk:"local_disk_size_gb" tf:"optional"`

	LocalDisks types.Int64 `tfsdk:"local_disks" tf:"optional"`

	LocalNvmeDiskSizeGb types.Int64 `tfsdk:"local_nvme_disk_size_gb" tf:"optional"`

	LocalNvmeDisks types.Int64 `tfsdk:"local_nvme_disks" tf:"optional"`
}

type NodeType struct {
	Category types.String `tfsdk:"category" tf:"optional"`
	// A string description associated with this node type, e.g., "r3.xlarge".
	Description types.String `tfsdk:"description" tf:""`

	DisplayOrder types.Int64 `tfsdk:"display_order" tf:"optional"`
	// An identifier for the type of hardware that this node runs on, e.g.,
	// "r3.2xlarge" in AWS.
	InstanceTypeId types.String `tfsdk:"instance_type_id" tf:""`
	// Whether the node type is deprecated. Non-deprecated node types offer
	// greater performance.
	IsDeprecated types.Bool `tfsdk:"is_deprecated" tf:"optional"`
	// AWS specific, whether this instance supports encryption in transit, used
	// for hipaa and pci workloads.
	IsEncryptedInTransit types.Bool `tfsdk:"is_encrypted_in_transit" tf:"optional"`

	IsGraviton types.Bool `tfsdk:"is_graviton" tf:"optional"`

	IsHidden types.Bool `tfsdk:"is_hidden" tf:"optional"`

	IsIoCacheEnabled types.Bool `tfsdk:"is_io_cache_enabled" tf:"optional"`
	// Memory (in MB) available for this node type.
	MemoryMb types.Int64 `tfsdk:"memory_mb" tf:""`

	NodeInfo *CloudProviderNodeInfo `tfsdk:"node_info" tf:"optional"`

	NodeInstanceType *NodeInstanceType `tfsdk:"node_instance_type" tf:"optional"`
	// Unique identifier for this node type.
	NodeTypeId types.String `tfsdk:"node_type_id" tf:""`
	// Number of CPU cores available for this node type. Note that this can be
	// fractional, e.g., 2.5 cores, if the the number of cores on a machine
	// instance is not divisible by the number of Spark nodes on that machine.
	NumCores types.Float64 `tfsdk:"num_cores" tf:""`

	NumGpus types.Int64 `tfsdk:"num_gpus" tf:"optional"`

	PhotonDriverCapable types.Bool `tfsdk:"photon_driver_capable" tf:"optional"`

	PhotonWorkerCapable types.Bool `tfsdk:"photon_worker_capable" tf:"optional"`

	SupportClusterTags types.Bool `tfsdk:"support_cluster_tags" tf:"optional"`

	SupportEbsVolumes types.Bool `tfsdk:"support_ebs_volumes" tf:"optional"`

	SupportPortForwarding types.Bool `tfsdk:"support_port_forwarding" tf:"optional"`
	// Indicates if this node type can be used for an instance pool or cluster
	// with elastic disk enabled. This is true for most node types.
	SupportsElasticDisk types.Bool `tfsdk:"supports_elastic_disk" tf:"optional"`
}

type PendingInstanceError struct {
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`

	Message types.String `tfsdk:"message" tf:"optional"`
}

type PermanentDeleteCluster struct {
	// The cluster to be deleted.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

type PermanentDeleteClusterResponse struct {
}

type PinCluster struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

type PinClusterResponse struct {
}

type Policy struct {
	// Creation time. The timestamp (in millisecond) when this Cluster Policy
	// was created.
	CreatedAtTimestamp types.Int64 `tfsdk:"created_at_timestamp" tf:"optional"`
	// Creator user name. The field won't be included in the response if the
	// user has already been deleted.
	CreatorUserName types.String `tfsdk:"creator_user_name" tf:"optional"`
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:"optional"`
	// Additional human-readable description of the cluster policy.
	Description types.String `tfsdk:"description" tf:"optional"`
	// If true, policy is a default policy created and managed by Databricks.
	// Default policies cannot be deleted, and their policy families cannot be
	// changed.
	IsDefault types.Bool `tfsdk:"is_default" tf:"optional"`
	// A list of libraries to be installed on the next cluster restart that uses
	// this policy. The maximum number of libraries is 500.
	Libraries []Library `tfsdk:"libraries" tf:"optional"`
	// Max number of clusters per user that can be active using this policy. If
	// not present, there is no max limit.
	MaxClustersPerUser types.Int64 `tfsdk:"max_clusters_per_user" tf:"optional"`
	// Cluster Policy name requested by the user. This has to be unique. Length
	// must be between 1 and 100 characters.
	Name types.String `tfsdk:"name" tf:"optional"`
	// Policy definition JSON document expressed in [Databricks Policy
	// Definition Language]. The JSON document must be passed as a string and
	// cannot be embedded in the requests.
	//
	// You can use this to customize the policy definition inherited from the
	// policy family. Policy rules specified here are merged into the inherited
	// policy definition.
	//
	// [Databricks Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	PolicyFamilyDefinitionOverrides types.String `tfsdk:"policy_family_definition_overrides" tf:"optional"`
	// ID of the policy family.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:"optional"`
	// Canonical unique identifier for the Cluster Policy.
	PolicyId types.String `tfsdk:"policy_id" tf:"optional"`
}

type PolicyFamily struct {
	// Policy definition document expressed in [Databricks Cluster Policy
	// Definition Language].
	//
	// [Databricks Cluster Policy Definition Language]: https://docs.databricks.com/administration-guide/clusters/policy-definition.html
	Definition types.String `tfsdk:"definition" tf:""`
	// Human-readable description of the purpose of the policy family.
	Description types.String `tfsdk:"description" tf:""`
	// Name of the policy family.
	Name types.String `tfsdk:"name" tf:""`
	// ID of the policy family.
	PolicyFamilyId types.String `tfsdk:"policy_family_id" tf:""`
}

type PythonPyPiLibrary struct {
	// The name of the pypi package to install. An optional exact version
	// specification is also supported. Examples: "simplejson" and
	// "simplejson==3.8.0".
	Package types.String `tfsdk:"package" tf:""`
	// The repository where the package can be found. If not specified, the
	// default pip index is used.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

type RCranLibrary struct {
	// The name of the CRAN package to install.
	Package types.String `tfsdk:"package" tf:""`
	// The repository where the package can be found. If not specified, the
	// default CRAN repo is used.
	Repo types.String `tfsdk:"repo" tf:"optional"`
}

type RemoveInstanceProfile struct {
	// The ARN of the instance profile to remove. This field is required.
	InstanceProfileArn types.String `tfsdk:"instance_profile_arn" tf:""`
}

type RemoveResponse struct {
}

type ResizeCluster struct {
	// Parameters needed in order to automatically scale clusters up and down
	// based on load. Note: autoscaling works best with DB runtime versions 3.0
	// or later.
	Autoscale *AutoScale `tfsdk:"autoscale" tf:"optional"`
	// The cluster to be resized.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// Number of worker nodes that this cluster should have. A cluster has one
	// Spark Driver and `num_workers` Executors for a total of `num_workers` + 1
	// Spark nodes.
	//
	// Note: When reading the properties of a cluster, this field reflects the
	// desired number of workers rather than the actual current number of
	// workers. For instance, if a cluster is resized from 5 to 10 workers, this
	// field will immediately be updated to reflect the target size of 10
	// workers, whereas the workers listed in `spark_info` will gradually
	// increase from 5 to 10 as the new nodes are provisioned.
	NumWorkers types.Int64 `tfsdk:"num_workers" tf:"optional"`
}

type ResizeClusterResponse struct {
}

type RestartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// <needs content added>
	RestartUser types.String `tfsdk:"restart_user" tf:"optional"`
}

type RestartClusterResponse struct {
}

type ResultType string

const ResultTypeError ResultType = `error`

const ResultTypeImage ResultType = `image`

const ResultTypeImages ResultType = `images`

const ResultTypeTable ResultType = `table`

const ResultTypeText ResultType = `text`

// String representation for [fmt.Print]
func (f *ResultType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ResultType) Set(v string) error {
	switch v {
	case `error`, `image`, `images`, `table`, `text`:
		*f = ResultType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "error", "image", "images", "table", "text"`, v)
	}
}

// Type always returns ResultType to satisfy [pflag.Value] interface
func (f *ResultType) Type() string {
	return "ResultType"
}

type Results struct {
	// The cause of the error
	Cause types.String `tfsdk:"cause" tf:"optional"`

	Data any `tfsdk:"data" tf:"optional"`
	// The image filename
	FileName types.String `tfsdk:"fileName" tf:"optional"`

	FileNames []types.String `tfsdk:"fileNames" tf:"optional"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema types.Bool `tfsdk:"isJsonSchema" tf:"optional"`
	// internal field used by SDK
	Pos types.Int64 `tfsdk:"pos" tf:"optional"`

	ResultType ResultType `tfsdk:"resultType" tf:"optional"`
	// The table schema
	Schema []map[string]any `tfsdk:"schema" tf:"optional"`
	// The summary of the error
	Summary types.String `tfsdk:"summary" tf:"optional"`
	// true if partial results are returned.
	Truncated types.Bool `tfsdk:"truncated" tf:"optional"`
}

// Decides which runtime engine to be use, e.g. Standard vs. Photon. If
// unspecified, the runtime engine is inferred from spark_version.
type RuntimeEngine string

const RuntimeEngineNull RuntimeEngine = `NULL`

const RuntimeEnginePhoton RuntimeEngine = `PHOTON`

const RuntimeEngineStandard RuntimeEngine = `STANDARD`

// String representation for [fmt.Print]
func (f *RuntimeEngine) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RuntimeEngine) Set(v string) error {
	switch v {
	case `NULL`, `PHOTON`, `STANDARD`:
		*f = RuntimeEngine(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NULL", "PHOTON", "STANDARD"`, v)
	}
}

// Type always returns RuntimeEngine to satisfy [pflag.Value] interface
func (f *RuntimeEngine) Type() string {
	return "RuntimeEngine"
}

type S3StorageInfo struct {
	// (Optional) Set canned access control list for the logs, e.g.
	// `bucket-owner-full-control`. If `canned_cal` is set, please make sure the
	// cluster iam role has `s3:PutObjectAcl` permission on the destination
	// bucket and prefix. The full list of possible canned acl can be found at
	// http://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.
	// Please also note that by default only the object owner gets full
	// controls. If you are using cross account role for writing data, you may
	// want to set `bucket-owner-full-control` to make bucket owner able to read
	// the logs.
	CannedAcl types.String `tfsdk:"canned_acl" tf:"optional"`
	// S3 destination, e.g. `s3://my-bucket/some-prefix` Note that logs will be
	// delivered using cluster iam role, please make sure you set cluster iam
	// role and the role has write access to the destination. Please also note
	// that you cannot use AWS keys to deliver logs.
	Destination types.String `tfsdk:"destination" tf:""`
	// (Optional) Flag to enable server side encryption, `false` by default.
	EnableEncryption types.Bool `tfsdk:"enable_encryption" tf:"optional"`
	// (Optional) The encryption type, it could be `sse-s3` or `sse-kms`. It
	// will be used only when encryption is enabled and the default type is
	// `sse-s3`.
	EncryptionType types.String `tfsdk:"encryption_type" tf:"optional"`
	// S3 endpoint, e.g. `https://s3-us-west-2.amazonaws.com`. Either region or
	// endpoint needs to be set. If both are set, endpoint will be used.
	Endpoint types.String `tfsdk:"endpoint" tf:"optional"`
	// (Optional) Kms key which will be used if encryption is enabled and
	// encryption type is set to `sse-kms`.
	KmsKey types.String `tfsdk:"kms_key" tf:"optional"`
	// S3 region, e.g. `us-west-2`. Either region or endpoint needs to be set.
	// If both are set, endpoint will be used.
	Region types.String `tfsdk:"region" tf:"optional"`
}

type SparkNode struct {
	// The private IP address of the host instance.
	HostPrivateIp types.String `tfsdk:"host_private_ip" tf:"optional"`
	// Globally unique identifier for the host instance from the cloud provider.
	InstanceId types.String `tfsdk:"instance_id" tf:"optional"`
	// Attributes specific to AWS for a Spark node.
	NodeAwsAttributes *SparkNodeAwsAttributes `tfsdk:"node_aws_attributes" tf:"optional"`
	// Globally unique identifier for this node.
	NodeId types.String `tfsdk:"node_id" tf:"optional"`
	// Private IP address (typically a 10.x.x.x address) of the Spark node. Note
	// that this is different from the private IP address of the host instance.
	PrivateIp types.String `tfsdk:"private_ip" tf:"optional"`
	// Public DNS address of this node. This address can be used to access the
	// Spark JDBC server on the driver node. To communicate with the JDBC
	// server, traffic must be manually authorized by adding security group
	// rules to the "worker-unmanaged" security group via the AWS console.
	//
	// Actually it's the public DNS address of the host instance.
	PublicDns types.String `tfsdk:"public_dns" tf:"optional"`
	// The timestamp (in millisecond) when the Spark node is launched.
	//
	// The start_timestamp is set right before the container is being launched.
	// The timestamp when the container is placed on the ResourceManager, before
	// its launch and setup by the NodeDaemon. This timestamp is the same as the
	// creation timestamp in the database.
	StartTimestamp types.Int64 `tfsdk:"start_timestamp" tf:"optional"`
}

type SparkNodeAwsAttributes struct {
	// Whether this node is on an Amazon spot instance.
	IsSpot types.Bool `tfsdk:"is_spot" tf:"optional"`
}

type SparkVersion struct {
	// Spark version key, for example "2.1.x-scala2.11". This is the value which
	// should be provided as the "spark_version" when creating a new cluster.
	// Note that the exact Spark version may change over time for a "wildcard"
	// version (i.e., "2.1.x-scala2.11" is a "wildcard" version) with minor bug
	// fixes.
	Key types.String `tfsdk:"key" tf:"optional"`
	// A descriptive name for this Spark version, for example "Spark 2.1".
	Name types.String `tfsdk:"name" tf:"optional"`
}

type StartCluster struct {
	// The cluster to be started.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

type StartClusterResponse struct {
}

// Current state of the cluster.
type State string

const StateError State = `ERROR`

const StatePending State = `PENDING`

const StateResizing State = `RESIZING`

const StateRestarting State = `RESTARTING`

const StateRunning State = `RUNNING`

const StateTerminated State = `TERMINATED`

const StateTerminating State = `TERMINATING`

const StateUnknown State = `UNKNOWN`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `ERROR`, `PENDING`, `RESIZING`, `RESTARTING`, `RUNNING`, `TERMINATED`, `TERMINATING`, `UNKNOWN`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "PENDING", "RESIZING", "RESTARTING", "RUNNING", "TERMINATED", "TERMINATING", "UNKNOWN"`, v)
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	Code TerminationReasonCode `tfsdk:"code" tf:"optional"`
	// list of parameters that provide additional information about why the
	// cluster was terminated
	Parameters map[string]types.String `tfsdk:"parameters" tf:"optional"`
	// type of the termination
	Type TerminationReasonType `tfsdk:"type" tf:"optional"`
}

// status code indicating why the cluster was terminated
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

// String representation for [fmt.Print]
func (f *TerminationReasonCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_SHUTDOWN`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `DATABASE_CONNECTION_FAILURE`, `DBFS_COMPONENT_UNHEALTHY`, `DOCKER_IMAGE_PULL_FAILURE`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `EXECUTION_COMPONENT_UNHEALTHY`, `GCP_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_UNREACHABLE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_SPARK_IMAGE`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `STORAGE_DOWNLOAD_FAILURE`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_SHUTDOWN", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "DATABASE_CONNECTION_FAILURE", "DBFS_COMPONENT_UNHEALTHY", "DOCKER_IMAGE_PULL_FAILURE", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "EXECUTION_COMPONENT_UNHEALTHY", "GCP_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_DELETED", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_UNREACHABLE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_SPARK_IMAGE", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "STORAGE_DOWNLOAD_FAILURE", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR"`, v)
	}
}

// Type always returns TerminationReasonCode to satisfy [pflag.Value] interface
func (f *TerminationReasonCode) Type() string {
	return "TerminationReasonCode"
}

// type of the termination
type TerminationReasonType string

const TerminationReasonTypeClientError TerminationReasonType = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonType = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonType = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationReasonType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `SERVICE_FAULT`, `SUCCESS`:
		*f = TerminationReasonType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "SERVICE_FAULT", "SUCCESS"`, v)
	}
}

// Type always returns TerminationReasonType to satisfy [pflag.Value] interface
func (f *TerminationReasonType) Type() string {
	return "TerminationReasonType"
}

type UninstallLibraries struct {
	// Unique identifier for the cluster on which to uninstall these libraries.
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
	// The libraries to uninstall.
	Libraries []Library `tfsdk:"libraries" tf:""`
}

type UninstallLibrariesResponse struct {
}

type UnpinCluster struct {
	// <needs content added>
	ClusterId types.String `tfsdk:"cluster_id" tf:""`
}

type UnpinClusterResponse struct {
}

type UpdateResponse struct {
}

type VolumesStorageInfo struct {
	// Unity Catalog Volumes file destination, e.g. `/Volumes/my-init.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}

type WorkloadType struct {
	// defined what type of clients can use the cluster. E.g. Notebooks, Jobs
	Clients ClientsTypes `tfsdk:"clients" tf:""`
}

type WorkspaceStorageInfo struct {
	// workspace files destination, e.g.
	// `/Users/user1@databricks.com/my-init.sh`
	Destination types.String `tfsdk:"destination" tf:""`
}
