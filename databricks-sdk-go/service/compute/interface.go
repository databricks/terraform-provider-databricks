// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"context"
)

// You can use cluster policies to control users' ability to configure clusters
// based on a set of rules. These rules specify which attributes or attribute
// values can be used during cluster creation. Cluster policies have ACLs that
// limit their use to specific users and groups.
//
// With cluster policies, you can: - Auto-install cluster libraries on the next
// restart by listing them in the policy's "libraries" field (Public Preview). -
// Limit users to creating clusters with the prescribed settings. - Simplify the
// user interface, enabling more users to create clusters, by fixing and hiding
// some fields. - Manage costs by setting limits on attributes that impact the
// hourly rate.
//
// Cluster policy permissions limit which policies a user can select in the
// Policy drop-down when the user creates a cluster: - A user who has
// unrestricted cluster create permission can select the Unrestricted policy and
// create fully-configurable clusters. - A user who has both unrestricted
// cluster create permission and access to cluster policies can select the
// Unrestricted policy and policies they have access to. - A user that has
// access to only cluster policies, can select the policies they have access to.
//
// If no policies exist in the workspace, the Policy drop-down doesn't appear.
// Only admin users can create, edit, and delete policies. Admin users also have
// access to all policies.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ClusterPoliciesService interface {

	// Creates a new policy with prescribed settings.
	Create(ctx context.Context, request CreatePolicy) (*CreatePolicyResponse, error)

	// Delete a policy for a cluster. Clusters governed by this policy can still
	// run, but cannot be edited.
	Delete(ctx context.Context, request DeletePolicy) error

	// Update an existing policy for cluster. This operation may make some
	// clusters governed by the previous policy invalid.
	Edit(ctx context.Context, request EditPolicy) error

	// Get a cluster policy entity. Creation and editing is available to admins
	// only.
	Get(ctx context.Context, request GetClusterPolicyRequest) (*Policy, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetClusterPolicyPermissionLevelsRequest) (*GetClusterPolicyPermissionLevelsResponse, error)

	// Gets the permissions of a cluster policy. Cluster policies can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error)

	// Returns a list of policies accessible by the requesting user.
	List(ctx context.Context, request ListClusterPoliciesRequest) (*ListPoliciesResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error)

	// Updates the permissions on a cluster policy. Cluster policies can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request ClusterPolicyPermissionsRequest) (*ClusterPolicyPermissions, error)
}

// The Clusters API allows you to create, start, edit, list, terminate, and
// delete clusters.
//
// Databricks maps cluster node instance types to compute units known as DBUs.
// See the instance type pricing page for a list of the supported instance types
// and their corresponding DBUs.
//
// A Databricks cluster is a set of computation resources and configurations on
// which you run data engineering, data science, and data analytics workloads,
// such as production ETL pipelines, streaming analytics, ad-hoc analytics, and
// machine learning.
//
// You run these workloads as a set of commands in a notebook or as an automated
// job. Databricks makes a distinction between all-purpose clusters and job
// clusters. You use all-purpose clusters to analyze data collaboratively using
// interactive notebooks. You use job clusters to run fast and robust automated
// jobs.
//
// You can create an all-purpose cluster using the UI, CLI, or REST API. You can
// manually terminate and restart an all-purpose cluster. Multiple users can
// share such clusters to do collaborative interactive analysis.
//
// IMPORTANT: Databricks retains cluster configuration information for
// terminated clusters for 30 days. To keep an all-purpose cluster configuration
// even after it has been terminated for more than 30 days, an administrator can
// pin a cluster to the cluster list.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ClustersService interface {

	// Change the owner of the cluster. You must be an admin and the cluster
	// must be terminated to perform this operation. The service principal
	// application ID can be supplied as an argument to `owner_username`.
	ChangeOwner(ctx context.Context, request ChangeClusterOwner) error

	// Creates a new Spark cluster. This method will acquire new instances from
	// the cloud provider if necessary. This method is asynchronous; the
	// returned ``cluster_id`` can be used to poll the cluster status. When this
	// method returns, the cluster will be in a ``PENDING`` state. The cluster
	// will be usable once it enters a ``RUNNING`` state. Note: Databricks may
	// not be able to acquire some of the requested nodes, due to cloud provider
	// limitations (account limits, spot price, etc.) or transient network
	// issues.
	//
	// If Databricks acquires at least 85% of the requested on-demand nodes,
	// cluster creation will succeed. Otherwise the cluster will terminate with
	// an informative error message.
	//
	// Rather than authoring the cluster's JSON definition from scratch,
	// Databricks recommends filling out the [create compute UI] and then
	// copying the generated JSON definition from the UI.
	//
	// [create compute UI]: https://docs.databricks.com/compute/configure.html
	Create(ctx context.Context, request CreateCluster) (*CreateClusterResponse, error)

	// Terminates the Spark cluster with the specified ID. The cluster is
	// removed asynchronously. Once the termination has completed, the cluster
	// will be in a `TERMINATED` state. If the cluster is already in a
	// `TERMINATING` or `TERMINATED` state, nothing will happen.
	Delete(ctx context.Context, request DeleteCluster) error

	// Updates the configuration of a cluster to match the provided attributes
	// and size. A cluster can be updated if it is in a `RUNNING` or
	// `TERMINATED` state.
	//
	// If a cluster is updated while in a `RUNNING` state, it will be restarted
	// so that the new attributes can take effect.
	//
	// If a cluster is updated while in a `TERMINATED` state, it will remain
	// `TERMINATED`. The next time it is started using the `clusters/start` API,
	// the new attributes will take effect. Any attempt to update a cluster in
	// any other state will be rejected with an `INVALID_STATE` error code.
	//
	// Clusters created by the Databricks Jobs service cannot be edited.
	Edit(ctx context.Context, request EditCluster) error

	// Retrieves a list of events about the activity of a cluster. This API is
	// paginated. If there are more events to read, the response includes all
	// the parameters necessary to request the next page of events.
	Events(ctx context.Context, request GetEvents) (*GetEventsResponse, error)

	// Retrieves the information for a cluster given its identifier. Clusters
	// can be described while they are running, or up to 60 days after they are
	// terminated.
	Get(ctx context.Context, request GetClusterRequest) (*ClusterDetails, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetClusterPermissionLevelsRequest) (*GetClusterPermissionLevelsResponse, error)

	// Gets the permissions of a cluster. Clusters can inherit permissions from
	// their root object.
	GetPermissions(ctx context.Context, request GetClusterPermissionsRequest) (*ClusterPermissions, error)

	// Return information about all pinned and active clusters, and all clusters
	// terminated within the last 30 days. Clusters terminated prior to this
	// period are not included.
	List(ctx context.Context, request ListClustersRequest) (*ListClustersResponse, error)

	// Returns a list of supported Spark node types. These node types can be
	// used to launch a cluster.
	ListNodeTypes(ctx context.Context) (*ListNodeTypesResponse, error)

	// Returns a list of availability zones where clusters can be created in
	// (For example, us-west-2a). These zones can be used to launch a cluster.
	ListZones(ctx context.Context) (*ListAvailableZonesResponse, error)

	// Permanently deletes a Spark cluster. This cluster is terminated and
	// resources are asynchronously removed.
	//
	// In addition, users will no longer see permanently deleted clusters in the
	// cluster list, and API users can no longer perform any action on
	// permanently deleted clusters.
	PermanentDelete(ctx context.Context, request PermanentDeleteCluster) error

	// Pinning a cluster ensures that the cluster will always be returned by the
	// ListClusters API. Pinning a cluster that is already pinned will have no
	// effect. This API can only be called by workspace admins.
	Pin(ctx context.Context, request PinCluster) error

	// Resizes a cluster to have a desired number of workers. This will fail
	// unless the cluster is in a `RUNNING` state.
	Resize(ctx context.Context, request ResizeCluster) error

	// Restarts a Spark cluster with the supplied ID. If the cluster is not
	// currently in a `RUNNING` state, nothing will happen.
	Restart(ctx context.Context, request RestartCluster) error

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error)

	// Returns the list of available Spark versions. These versions can be used
	// to launch a cluster.
	SparkVersions(ctx context.Context) (*GetSparkVersionsResponse, error)

	// Starts a terminated Spark cluster with the supplied ID. This works
	// similar to `createCluster` except: - The previous cluster id and
	// attributes are preserved. - The cluster starts with the last specified
	// cluster size. - If the previous cluster was an autoscaling cluster, the
	// current cluster starts with the minimum number of nodes. - If the cluster
	// is not currently in a ``TERMINATED`` state, nothing will happen. -
	// Clusters launched to run a job cannot be started.
	Start(ctx context.Context, request StartCluster) error

	// Unpinning a cluster will allow the cluster to eventually be removed from
	// the ListClusters API. Unpinning a cluster that is not pinned will have no
	// effect. This API can only be called by workspace admins.
	Unpin(ctx context.Context, request UnpinCluster) error

	// Updates the configuration of a cluster to match the partial set of
	// attributes and size. Denote which fields to update using the
	// `update_mask` field in the request body. A cluster can be updated if it
	// is in a `RUNNING` or `TERMINATED` state. If a cluster is updated while in
	// a `RUNNING` state, it will be restarted so that the new attributes can
	// take effect. If a cluster is updated while in a `TERMINATED` state, it
	// will remain `TERMINATED`. The updated attributes will take effect the
	// next time the cluster is started using the `clusters/start` API. Attempts
	// to update a cluster in any other state will be rejected with an
	// `INVALID_STATE` error code. Clusters created by the Databricks Jobs
	// service cannot be updated.
	Update(ctx context.Context, request UpdateCluster) error

	// Updates the permissions on a cluster. Clusters can inherit permissions
	// from their root object.
	UpdatePermissions(ctx context.Context, request ClusterPermissionsRequest) (*ClusterPermissions, error)
}

// This API allows execution of Python, Scala, SQL, or R commands on running
// Databricks Clusters. This API only supports (classic) all-purpose clusters.
// Serverless compute is not supported.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CommandExecutionService interface {

	// Cancels a currently running command within an execution context.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	// Cancellation is asynchronous: this method returns once the request is
	// accepted, and the command moves through the `Cancelling` status before
	// reaching a terminal one. Poll __commandStatus__ to observe the outcome; a
	// cancelled command ends in the `Finished` status with an `error` result
	// that describes the cancellation.
	Cancel(ctx context.Context, request CancelCommand) error

	// Gets the status of and, if available, the results from a currently
	// executing command.
	//
	// The command ID is obtained from a prior successful call to __execute__.
	// Poll this method until the status is one of the terminal values
	// `Finished`, `Error` or `Cancelled`; the `results` field is populated only
	// once the command has reached a terminal status.
	CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error)

	// Gets the status for an execution context.
	//
	// After __Create__, poll this method until the status is `Running` or
	// `Error`. Requests for a context that has been deleted fail with an error
	// instead of returning a status.
	ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error)

	// Creates an execution context for running cluster commands.
	//
	// If successful, this method returns the ID of the new execution context.
	// The context may still be initializing when this method returns, so poll
	// __contextStatus__ until the status is `Running` before running commands
	// in it. This method and the other methods of this API require the CAN
	// ATTACH TO permission on the cluster.
	Create(ctx context.Context, request CreateContext) (*Created, error)

	// Deletes an execution context.
	//
	// Delete contexts you no longer need, because a cluster can host only a
	// limited number of execution contexts at a time. Deleting a context that
	// no longer exists succeeds, so the call is safe to retry.
	Destroy(ctx context.Context, request DestroyContext) error

	// Runs a cluster command in the given execution context, using the provided
	// language.
	//
	// If successful, it returns an ID for tracking the status of the command's
	// execution. This method returns as soon as the command is accepted, so
	// poll __commandStatus__ with the returned ID to wait for the command to
	// finish and to read its results.
	Execute(ctx context.Context, request Command) (*Created, error)
}

// The Global Init Scripts API enables Workspace administrators to configure
// global initialization scripts for their workspace. These scripts run on every
// node in every cluster in the workspace.
//
// **Important:** Existing clusters must be restarted to pick up any changes
// made to global init scripts. Global init scripts are run in order. If the
// init script returns with a bad exit code, the Apache Spark container fails to
// launch and init scripts with later position are skipped. If enough containers
// fail, the entire cluster fails with a `GLOBAL_INIT_SCRIPT_FAILURE` error
// code.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type GlobalInitScriptsService interface {

	// Creates a new global init script in this workspace. The number of global
	// init scripts per workspace is limited.
	Create(ctx context.Context, request GlobalInitScriptCreateRequest) (*CreateResponse, error)

	// Deletes a global init script.
	Delete(ctx context.Context, request DeleteGlobalInitScriptRequest) error

	// Gets all the details of a script, including its Base64-encoded contents.
	Get(ctx context.Context, request GetGlobalInitScriptRequest) (*GlobalInitScriptDetailsWithContent, error)

	// Get a list of all global init scripts for this workspace. This returns
	// all properties for each script but **not** the script contents. To
	// retrieve the contents of a script, use the [get a global init
	// script](:method:globalinitscripts/get) operation.
	List(ctx context.Context) (*ListGlobalInitScriptsResponse, error)

	// Updates a global init script, specifying only the fields to change. All
	// fields are optional. Unspecified fields retain their current value.
	Update(ctx context.Context, request GlobalInitScriptUpdateRequest) error
}

// Instance Pools API are used to create, edit, delete and list instance pools
// by using ready-to-use cloud instances which reduces a cluster start and
// auto-scaling times.
//
// Databricks pools reduce cluster start and auto-scaling times by maintaining a
// set of idle, ready-to-use instances. When a cluster is attached to a pool,
// cluster nodes are created using the pool’s idle instances. If the pool has
// no idle instances, the pool expands by allocating a new instance from the
// instance provider in order to accommodate the cluster’s request. When a
// cluster releases an instance, it returns to the pool and is free for another
// cluster to use. Only clusters attached to a pool can use that pool’s idle
// instances.
//
// You can specify a different pool for the driver node and worker nodes, or use
// the same pool for both.
//
// Databricks does not charge DBUs while instances are idle in the pool.
// Instance provider billing does apply. See pricing.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type InstancePoolsService interface {

	// Creates a new instance pool using idle and ready-to-use cloud instances.
	Create(ctx context.Context, request CreateInstancePool) (*CreateInstancePoolResponse, error)

	// Deletes the instance pool permanently. The idle instances in the pool are
	// terminated asynchronously.
	Delete(ctx context.Context, request DeleteInstancePool) error

	// Modifies the configuration of an existing instance pool. Settings that
	// are fixed at creation, such as the node type, cannot be changed and the
	// request fails if it attempts to change one. Omitted `min_idle_instances`,
	// `max_capacity` and `idle_instance_autotermination_minutes` are reset to
	// their defaults; omitted `custom_tags` keep their current value. On Azure,
	// an omitted `azure_attributes.capacity_reservation_group` is removed.
	Edit(ctx context.Context, request EditInstancePool) error

	// Retrieve the information for an instance pool based on its identifier.
	Get(ctx context.Context, request GetInstancePoolRequest) (*GetInstancePool, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetInstancePoolPermissionLevelsRequest) (*GetInstancePoolPermissionLevelsResponse, error)

	// Gets the permissions of an instance pool. Instance pools can inherit
	// permissions from their root object.
	GetPermissions(ctx context.Context, request GetInstancePoolPermissionsRequest) (*InstancePoolPermissions, error)

	// Gets a list of instance pools with their statistics.
	List(ctx context.Context) (*ListInstancePools, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error)

	// Updates the permissions on an instance pool. Instance pools can inherit
	// permissions from their root object.
	UpdatePermissions(ctx context.Context, request InstancePoolPermissionsRequest) (*InstancePoolPermissions, error)
}

// The Instance Profiles API allows admins to add, list, and remove instance
// profiles that users can launch clusters with. Regular users can list the
// instance profiles available to them. See [Secure access to S3 buckets] using
// instance profiles for more information.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
//
// [Secure access to S3 buckets]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html
type InstanceProfilesService interface {

	// Registers an instance profile in Databricks. In the UI, you can then give
	// users the permission to use this instance profile when launching
	// clusters.
	//
	// This API is only available to admin users.
	Add(ctx context.Context, request AddInstanceProfile) error

	// The only supported field to change is the optional IAM role ARN
	// associated with the instance profile. It is required to specify the IAM
	// role ARN if both of the following are true:
	//
	// * Your role name and instance profile name do not match. The name is the
	// part after the last slash in each ARN. * You want to use the instance
	// profile with [Databricks SQL Serverless].
	//
	// To understand where these fields are in the AWS console, see [Enable
	// serverless SQL warehouses].
	//
	// This API is only available to admin users.
	//
	// [Databricks SQL Serverless]: https://docs.databricks.com/sql/admin/serverless.html
	// [Enable serverless SQL warehouses]: https://docs.databricks.com/sql/admin/serverless.html
	Edit(ctx context.Context, request InstanceProfile) error

	// List the instance profiles that the calling user can use to launch a
	// cluster.
	//
	// This API is available to all users.
	List(ctx context.Context) (*ListInstanceProfilesResponse, error)

	// Remove the instance profile with the provided ARN. Existing clusters with
	// this instance profile will continue to function.
	//
	// This API is only accessible to admin users.
	Remove(ctx context.Context, request RemoveInstanceProfile) error
}

// The Libraries API allows you to install and uninstall libraries and get the
// status of libraries on a cluster.
//
// To make third-party or custom code available to notebooks and jobs running on
// your clusters, you can install a library. Libraries can be written in Python,
// Java, Scala, and R. You can upload Python, Java, Scala and R libraries and
// point to external packages in PyPI, Maven, and CRAN repositories.
//
// Cluster libraries can be used by all notebooks running on a cluster. You can
// install a cluster library directly from a public repository such as PyPI or
// Maven, using a previously installed workspace library, or using an init
// script.
//
// When you uninstall a library from a cluster, the library is removed only when
// you restart the cluster. Until you restart the cluster, the status of the
// uninstalled library appears as Uninstall pending restart.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type LibrariesService interface {

	// Get the status of all libraries on all clusters in the workspace. A
	// status is returned for all libraries installed on each cluster via the
	// API or the libraries UI. If the workspace has more library statuses than
	// the response limit allows, the request fails; call the per-cluster status
	// method for each cluster instead.
	AllClusterStatuses(ctx context.Context) (*ListAllClusterLibraryStatusesResponse, error)

	// Get the status of libraries on a cluster. A status is returned for all
	// libraries installed on this cluster via the API or the libraries UI. The
	// order of returned libraries is as follows: 1. Libraries set to be
	// installed on this cluster, in the order that the libraries were added to
	// the cluster, are returned first. 2. Libraries that were previously
	// requested to be installed on this cluster but are now marked for removal,
	// in no particular order, are returned last.
	ClusterStatus(ctx context.Context, request ClusterStatus) (*ClusterLibraryStatuses, error)

	// Add libraries to install on a cluster. The installation is asynchronous;
	// it happens in the background after the completion of this request.
	Install(ctx context.Context, request InstallLibraries) error

	// Set libraries to uninstall from a cluster. The libraries won't be
	// uninstalled until the cluster is restarted. A request to uninstall a
	// library that is not currently installed is ignored.
	Uninstall(ctx context.Context, request UninstallLibraries) error
}

// The policy compliance APIs allow you to view and manage the policy compliance
// status of clusters in your workspace.
//
// A cluster is compliant with its policy if its configuration satisfies all its
// policy rules. Clusters could be out of compliance if their policy was updated
// after the cluster was last edited.
//
// The get and list compliance APIs allow you to view the policy compliance
// status of a cluster. The enforce compliance API allows you to update a
// cluster to be compliant with the current version of its policy.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PolicyComplianceForClustersService interface {

	// Cancels a pending enforcement on a cluster. After canceling the pending
	// enforcement, the cluster will no longer update on the next termination or
	// restart. Pending enforcements cannot be canceled when a cluster is in
	// `TERMINATING` state. Only workspace admins can cancel pending
	// enforcements.
	CancelPendingClusterEnforcement(ctx context.Context, request CancelPendingClusterEnforcementRequest) (*CancelPendingClusterEnforcementResponse, error)

	// Updates a cluster to be compliant with the current version of its policy.
	//
	// If a cluster is updated while in a `TERMINATED` state, it will remain
	// `TERMINATED`. The next time the cluster is started, the new attributes
	// will take effect.
	//
	// For clusters in other states, the behavior depends on the `enforce_mode`
	// used.
	//
	// Clusters created by the Databricks Jobs, SDP, or Models services cannot
	// be enforced by this API. Instead, use the "Enforce job policy compliance"
	// API to enforce policy compliance on jobs.
	EnforceCompliance(ctx context.Context, request EnforceClusterComplianceRequest) (*EnforceClusterComplianceResponse, error)

	// Returns the policy compliance status of a cluster. Clusters could be out
	// of compliance if their policy was updated after the cluster was last
	// edited.
	GetCompliance(ctx context.Context, request GetClusterComplianceRequest) (*GetClusterComplianceResponse, error)

	// Returns the policy compliance status of all clusters that use a given
	// policy. Clusters could be out of compliance if their policy was updated
	// after the cluster was last edited.
	ListCompliance(ctx context.Context, request ListClusterCompliancesRequest) (*ListClusterCompliancesResponse, error)
}

// View available policy families. A policy family contains a policy definition
// providing best practices for configuring clusters for a particular use case.
//
// Databricks manages and provides policy families for several common cluster
// use cases. You cannot create, edit, or delete policy families.
//
// Policy families cannot be used directly to create clusters. Instead, you
// create cluster policies using a policy family. Cluster policies created using
// a policy family inherit the policy family's policy definition.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type PolicyFamiliesService interface {

	// Retrieve the information for an policy family based on its identifier and
	// version
	Get(ctx context.Context, request GetPolicyFamilyRequest) (*PolicyFamily, error)

	// Returns the list of policy definition types available to use at their
	// latest version. This API is paginated.
	List(ctx context.Context, request ListPolicyFamiliesRequest) (*ListPolicyFamiliesResponse, error)
}

// ServerlessCompute Service API.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type ServerlessComputeServiceService interface {

	// Creates a custom compute policy in the workspace.
	CreateComputePolicy(ctx context.Context, request CreateComputePolicyRequest) (*ComputePolicy, error)

	// Deletes a custom compute policy in the workspace. Default
	// (system-provisioned) policies cannot be deleted.
	DeleteComputePolicy(ctx context.Context, request DeleteComputePolicyRequest) error

	// Retrieves a compute policy by its ID.
	GetComputePolicy(ctx context.Context, request GetComputePolicyRequest) (*ComputePolicy, error)

	// Lists Interactive serverless computes that the caller can use or manage.
	// Results are returned in an unspecified order.
	List(ctx context.Context, request ListServerlessComputesRequest) (*ListServerlessComputesResponse, error)

	// Lists compute policies in the workspace. Results are returned in an
	// unspecified order; callers should sort client-side if a specific order is
	// needed.
	ListComputePolicies(ctx context.Context, request ListComputePoliciesRequest) (*ListComputePoliciesResponse, error)

	// Updates a custom compute policy in the workspace. The display_name and
	// allowed_compute_products are mutable; use update_mask to control which
	// fields change. Default (system-provisioned) policies cannot be updated.
	UpdateComputePolicy(ctx context.Context, request UpdateComputePolicyRequest) (*ComputePolicy, error)
}
