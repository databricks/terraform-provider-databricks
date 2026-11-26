// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package disasterrecovery

import (
	"context"
)

// Manage disaster recovery configurations and execute failover operations.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type DisasterRecoveryService interface {

	// Create a failover group. The group starts in the CREATING state and moves
	// to INITIAL_REPLICATION while the first replication cycle copies all
	// in-scope data to the secondary region; failover becomes available once
	// the group is ACTIVE. Set validate_only to check the request without
	// creating anything.
	CreateFailoverGroup(ctx context.Context, request CreateFailoverGroupRequest) (*FailoverGroup, error)

	// Create a new stable URL.
	CreateStableUrl(ctx context.Context, request CreateStableUrlRequest) (*StableUrl, error)

	// Delete a failover group. Deleting stops replication for the group; stable
	// URLs linked to it are detached but not deleted.
	DeleteFailoverGroup(ctx context.Context, request DeleteFailoverGroupRequest) error

	// Delete a stable URL.
	DeleteStableUrl(ctx context.Context, request DeleteStableUrlRequest) error

	// Initiate a failover that makes target_primary_region the primary region.
	// Stable URLs linked to the group are pointed at the new primary workspace,
	// replication is reversed so that the former primary becomes the secondary,
	// and the group moves through FAILING_OVER to INITIAL_REPLICATION until it
	// has a replication point in the new direction. The group must be ACTIVE;
	// repeating the request while it is FAILING_OVER is accepted. To fail back,
	// repeat the request with the former primary region as the target once the
	// group is ACTIVE again. In an unplanned failover, data written after the
	// replication point can be lost.
	FailoverFailoverGroup(ctx context.Context, request FailoverFailoverGroupRequest) (*FailoverGroup, error)

	// Get a failover group, including its current state and replication point.
	GetFailoverGroup(ctx context.Context, request GetFailoverGroupRequest) (*FailoverGroup, error)

	// Get a stable URL.
	GetStableUrl(ctx context.Context, request GetStableUrlRequest) (*StableUrl, error)

	// List the failover groups in an account.
	//
	// List entries are abbreviated: `state`, `replication_point` and
	// `unity_catalog_assets.location_mappings` are not populated. Call
	// GetFailoverGroup to retrieve the full resource.
	ListFailoverGroups(ctx context.Context, request ListFailoverGroupsRequest) (*ListFailoverGroupsResponse, error)

	// List stable URLs for an account.
	ListStableUrls(ctx context.Context, request ListStableUrlsRequest) (*ListStableUrlsResponse, error)

	// Update a failover group. Only the fields named in update_mask change. You
	// can turn workspace asset replication on or off, attach or detach a stable
	// URL, and change the Unity Catalog catalogs and location mappings; the
	// regions, the workspace set name and the workspace IDs cannot be changed.
	// Updates are rejected while a create, failover or delete is in progress,
	// and after a create or delete has failed.
	UpdateFailoverGroup(ctx context.Context, request UpdateFailoverGroupRequest) (*FailoverGroup, error)
}
