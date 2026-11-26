// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package cleanrooms

import (
	"context"
)

// Clean Room Asset Revisions denote new versions of uploaded assets (e.g.
// notebooks) in the clean room.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CleanRoomAssetRevisionsService interface {

	// Get a specific revision of an asset
	Get(ctx context.Context, request GetCleanRoomAssetRevisionRequest) (*CleanRoomAsset, error)

	// List revisions for an asset
	List(ctx context.Context, request ListCleanRoomAssetRevisionsRequest) (*ListCleanRoomAssetRevisionsResponse, error)
}

// Clean room assets are data and code objects — Tables, volumes, and
// notebooks that are shared with the clean room.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CleanRoomAssetsService interface {

	// Create a clean room asset —share an asset like a notebook or table into
	// the clean room. For each UC asset that is added through this method, the
	// clean room owner must also have enough privilege on the asset to consume
	// it. The privilege must be maintained indefinitely for the clean room to
	// be able to access the asset. Typically, you should use a group as the
	// clean room owner.
	Create(ctx context.Context, request CreateCleanRoomAssetRequest) (*CleanRoomAsset, error)

	// Submit an asset review
	CreateCleanRoomAssetReview(ctx context.Context, request CreateCleanRoomAssetReviewRequest) (*CreateCleanRoomAssetReviewResponse, error)

	// Delete a clean room asset - unshare/remove the asset from the clean room
	Delete(ctx context.Context, request DeleteCleanRoomAssetRequest) error

	// Get the details of a clean room asset by its type and full name.
	Get(ctx context.Context, request GetCleanRoomAssetRequest) (*CleanRoomAsset, error)

	// List the assets in a clean room. Listed assets do not include their
	// asset-type specific details, such as table columns or notebook content;
	// use the get method to retrieve them.
	List(ctx context.Context, request ListCleanRoomAssetsRequest) (*ListCleanRoomAssetsResponse, error)

	// Update a clean room asset. For example, updating the content of a
	// notebook; changing the shared partitions of a table; etc.
	Update(ctx context.Context, request UpdateCleanRoomAssetRequest) (*CleanRoomAsset, error)
}

// Clean room auto-approval rules automatically create an approval on your
// behalf when an asset (e.g. notebook) meeting specific criteria is shared in a
// clean room.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CleanRoomAutoApprovalRulesService interface {

	// Create an auto-approval rule
	Create(ctx context.Context, request CreateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error)

	// Delete a auto-approval rule by rule ID
	Delete(ctx context.Context, request DeleteCleanRoomAutoApprovalRuleRequest) error

	// Get a auto-approval rule by rule ID
	Get(ctx context.Context, request GetCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error)

	// List all auto-approval rules for the caller
	List(ctx context.Context, request ListCleanRoomAutoApprovalRulesRequest) (*ListCleanRoomAutoApprovalRulesResponse, error)

	// Update a auto-approval rule by rule ID
	Update(ctx context.Context, request UpdateCleanRoomAutoApprovalRuleRequest) (*CleanRoomAutoApprovalRule, error)
}

// Clean room task runs are the executions of notebooks and JAR analyses in a
// clean room.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CleanRoomTaskRunsService interface {

	// List all the historical notebook task runs in a clean room.
	List(ctx context.Context, request ListCleanRoomNotebookTaskRunsRequest) (*ListCleanRoomNotebookTaskRunsResponse, error)

	// List all the historical task runs in a clean room.
	ListCleanRoomTaskRunsHandler(ctx context.Context, request ListCleanRoomTaskRunsRequest) (*ListCleanRoomTaskRunsResponse, error)
}

// A clean room uses Delta Sharing and serverless compute to provide a secure
// and privacy-protecting environment where multiple parties can work together
// on sensitive enterprise data without direct access to each other's data.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type CleanRoomsService interface {

	// Create a new clean room with the specified collaborators. This method is
	// asynchronous; the name field of the returned clean room can be used to
	// poll the clean room status, using the [cleanrooms/get] method. When this
	// method returns, the clean room will be in a PROVISIONING state, with only
	// name, owner, comment, created_at and status populated. The clean room
	// will be usable once it enters an ACTIVE state.
	//
	// The caller must be a metastore admin or have the **CREATE_CLEAN_ROOM**
	// privilege on the metastore.
	//
	// [cleanrooms/get]: https://docs.databricks.com/api/workspace/cleanrooms/get
	Create(ctx context.Context, request CreateCleanRoomRequest) (*CleanRoom, error)

	// Create the output catalog of the clean room. The output catalog is a
	// catalog in the caller's metastore where the output tables of the clean
	// room runs that the caller triggers are made available. Each collaborator
	// can create one output catalog per clean room, and the clean room must be
	// ACTIVE.
	CreateOutputCatalog(ctx context.Context, request CreateCleanRoomOutputCatalogRequest) (*CreateCleanRoomOutputCatalogResponse, error)

	// Delete a clean room. After deletion, the clean room will be removed from
	// the metastore. If the other collaborators have not deleted the clean
	// room, they will still have the clean room in their metastore, but it will
	// be in a DELETED state and no operations other than deletion can be
	// performed on it.
	Delete(ctx context.Context, request DeleteCleanRoomRequest) error

	// Get the details of a clean room given its name.
	Get(ctx context.Context, request GetCleanRoomRequest) (*CleanRoom, error)

	// Get a list of all clean rooms of the metastore. Only clean rooms the
	// caller has access to are returned.
	List(ctx context.Context, request ListCleanRoomsRequest) (*ListCleanRoomsResponse, error)

	// Update a clean room. The caller must be the owner of the clean room, have
	// **MODIFY_CLEAN_ROOM** privilege, or be metastore admin.
	//
	// When the caller is a metastore admin, only the __owner__ field can be
	// updated.
	Update(ctx context.Context, request UpdateCleanRoomRequest) (*CleanRoom, error)
}
