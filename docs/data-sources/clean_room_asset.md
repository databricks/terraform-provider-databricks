---
subcategory: "Clean Rooms"
---
# databricks_clean_room_asset Data Source
# Datasource (Singular) Artifact

This data source can be used to get a single clean room asset.

## Example Usage
# Example: Clean Room Asset Datasource

```hcl
resource "databricks_clean_rooms_asset" "this" {
      name = "example-cleanroom-asset"
}
``` 

## Arguments
The following arguments are supported:
* `asset_type` (string, required) - The type of the asset. Possible values are: `FOREIGN_TABLE`, `NOTEBOOK_FILE`, `TABLE`, `VIEW`, `VOLUME`
* `clean_room_name` (string, required) - The name of the clean room this asset belongs to.
  This field is required for create operations and populated by the server for responses
* `name` (string, required) - A fully qualified name that uniquely identifies the asset within the clean room.
  This is also the name displayed in the clean room UI.
  
  For UC securable assets (tables, volumes, etc.), the format is *shared_catalog*.*shared_schema*.*asset_name*
  
  For notebooks, the name is the notebook file name.
  For jar analyses, the name is the jar analysis name

## Attributes
The following attributes are exported:
* `added_at` (integer) - When the asset is added to the clean room, in epoch milliseconds
* `asset_type` (string) - The type of the asset. Possible values are: `FOREIGN_TABLE`, `NOTEBOOK_FILE`, `TABLE`, `VIEW`, `VOLUME`
* `clean_room_name` (string) - The name of the clean room this asset belongs to.
  This field is required for create operations and populated by the server for responses
* `foreign_table` (CleanRoomAssetForeignTable) - Foreign table details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **FOREIGN_TABLE**
* `foreign_table_local_details` (CleanRoomAssetForeignTableLocalDetails) - Local details for a foreign that are only available to its owner.
  Present if and only if **asset_type** is **FOREIGN_TABLE**
* `name` (string) - A fully qualified name that uniquely identifies the asset within the clean room.
  This is also the name displayed in the clean room UI.
  
  For UC securable assets (tables, volumes, etc.), the format is *shared_catalog*.*shared_schema*.*asset_name*
  
  For notebooks, the name is the notebook file name.
  For jar analyses, the name is the jar analysis name
* `notebook` (CleanRoomAssetNotebook) - Notebook details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **NOTEBOOK_FILE**
* `owner_collaborator_alias` (string) - The alias of the collaborator who owns this asset
* `status` (string) - Status of the asset. Possible values are: `ACTIVE`, `PENDING`, `PERMISSION_DENIED`
* `table` (CleanRoomAssetTable) - Table details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **TABLE**
* `table_local_details` (CleanRoomAssetTableLocalDetails) - Local details for a table that are only available to its owner.
  Present if and only if **asset_type** is **TABLE**
* `view` (CleanRoomAssetView) - View details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **VIEW**
* `view_local_details` (CleanRoomAssetViewLocalDetails) - Local details for a view that are only available to its owner.
  Present if and only if **asset_type** is **VIEW**
* `volume_local_details` (CleanRoomAssetVolumeLocalDetails) - Local details for a volume that are only available to its owner.
  Present if and only if **asset_type** is **VOLUME**

### CleanRoomAssetForeignTable
* `columns` (list of ColumnInfo) - The metadata information of the columns in the foreign table

### CleanRoomAssetForeignTableLocalDetails
* `local_name` (string) - The fully qualified name of the foreign table in its owner's local metastore,
  in the format of *catalog*.*schema*.*foreign_table_name*

### CleanRoomAssetNotebook
* `etag` (string) - Server generated etag that represents the notebook version
* `notebook_content` (string) - Base 64 representation of the notebook contents.
  This is the same format as returned by :method:workspace/export with the format of **HTML**
* `review_state` (string) - Top-level status derived from all reviews. Possible values are: `APPROVED`, `PENDING`, `REJECTED`
* `reviews` (list of CleanRoomNotebookReview) - All existing approvals or rejections
* `runner_collaborator_aliases` (list of string) - collaborators that can run the notebook

### CleanRoomAssetTable
* `columns` (list of ColumnInfo) - The metadata information of the columns in the table

### CleanRoomAssetTableLocalDetails
* `local_name` (string) - The fully qualified name of the table in its owner's local metastore,
  in the format of *catalog*.*schema*.*table_name*
* `partitions` (list of Partition) - Partition filtering specification for a shared table

### CleanRoomAssetView
* `columns` (list of ColumnInfo) - The metadata information of the columns in the view

### CleanRoomAssetViewLocalDetails
* `local_name` (string) - The fully qualified name of the view in its owner's local metastore,
  in the format of *catalog*.*schema*.*view_name*

### CleanRoomAssetVolumeLocalDetails
* `local_name` (string) - The fully qualified name of the volume in its owner's local metastore,
  in the format of *catalog*.*schema*.*volume_name*

### CleanRoomNotebookReview
* `comment` (string) - Review comment
* `created_at_millis` (integer) - When the review was submitted, in epoch milliseconds
* `review_state` (string) - Review outcome. Possible values are: `APPROVED`, `PENDING`, `REJECTED`
* `review_sub_reason` (string) - Specified when the review was not explicitly made by a user. Possible values are: `AUTO_APPROVED`, `BACKFILLED`
* `reviewer_collaborator_alias` (string) - Collaborator alias of the reviewer

### ColumnInfo
* `comment` (string) - User-provided free-form text description
* `mask` (ColumnMask)
* `name` (string) - Name of Column
* `nullable` (boolean) - Whether field may be Null (default: true)
* `partition_index` (integer) - Partition index for column
* `position` (integer) - Ordinal position of column (starting at position 0)
* `type_interval_type` (string) - Format of IntervalType
* `type_json` (string) - Full data type specification, JSON-serialized
* `type_name` (string) - . Possible values are: `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `GEOGRAPHY`, `GEOMETRY`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`
* `type_precision` (integer) - Digits of precision; required for DecimalTypes
* `type_scale` (integer) - Digits to right of decimal; Required for DecimalTypes
* `type_text` (string) - Full data type specification as SQL/catalogString text

### ColumnMask
* `function_name` (string) - The full name of the column mask SQL UDF
* `using_column_names` (list of string) - The list of additional table columns to be passed as input to the column mask function. The
  first arg of the mask function should be of the type of the column being masked and the
  types of the rest of the args should match the types of columns in 'using_column_names'

### Partition
* `values` (list of PartitionValue) - An array of partition values

### PartitionValue
* `name` (string) - The name of the partition column
* `op` (string) - The operator to apply for the value. Possible values are: `EQUAL`, `LIKE`
* `recipient_property_key` (string) - The key of a Delta Sharing recipient's property. For example "databricks-account-id".
  When this field is set, field `value` can not be set
* `value` (string) - The value of the partition column. When this value is not set, it means `null` value.
  When this field is set, field `recipient_property_key` can not be set