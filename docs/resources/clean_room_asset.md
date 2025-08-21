---
subcategory: "Clean Rooms"
---
# databricks_clean_room_asset Resource
# Clean Room Asset Resource

Clean room assets are data and code objects -- tables, volumes, and notebooks that are shared with the clean room.

## Example Usage
# Example: Clean Room Asset Resource

### Example for sharing a table
This is an example for sharing an asset (table) in a clean room:
```hcl
resource "databricks_clean_rooms_asset" "this" {
      clean_room_name = "existing_clean_room"
      name = "creator.default.myasset"
      asset_type = "TABLE"
      table_local_details = {
          local_name = "some_creator.default.myasset"
      }
}
```


## Arguments
The following arguments are supported:
* `asset_type` (string, required) - The type of the asset. Possible values are: `FOREIGN_TABLE`, `NOTEBOOK_FILE`, `TABLE`, `VIEW`, `VOLUME`
* `name` (string, required) - A fully qualified name that uniquely identifies the asset within the clean room.
  This is also the name displayed in the clean room UI.
  
  For UC securable assets (tables, volumes, etc.), the format is *shared_catalog*.*shared_schema*.*asset_name*
  
  For notebooks, the name is the notebook file name.
  For jar analyses, the name is the jar analysis name
* `clean_room_name` (string, optional) - The name of the clean room this asset belongs to.
  This field is required for create operations and populated by the server for responses
* `foreign_table` (CleanRoomAssetForeignTable, optional) - Foreign table details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **FOREIGN_TABLE**
* `foreign_table_local_details` (CleanRoomAssetForeignTableLocalDetails, optional) - Local details for a foreign that are only available to its owner.
  Present if and only if **asset_type** is **FOREIGN_TABLE**
* `notebook` (CleanRoomAssetNotebook, optional) - Notebook details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **NOTEBOOK_FILE**
* `table` (CleanRoomAssetTable, optional) - Table details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **TABLE**
* `table_local_details` (CleanRoomAssetTableLocalDetails, optional) - Local details for a table that are only available to its owner.
  Present if and only if **asset_type** is **TABLE**
* `view` (CleanRoomAssetView, optional) - View details available to all collaborators of the clean room.
  Present if and only if **asset_type** is **VIEW**
* `view_local_details` (CleanRoomAssetViewLocalDetails, optional) - Local details for a view that are only available to its owner.
  Present if and only if **asset_type** is **VIEW**
* `volume_local_details` (CleanRoomAssetVolumeLocalDetails, optional) - Local details for a volume that are only available to its owner.
  Present if and only if **asset_type** is **VOLUME**

### CleanRoomAssetForeignTableLocalDetails
* `local_name` (string, required) - The fully qualified name of the foreign table in its owner's local metastore,
  in the format of *catalog*.*schema*.*foreign_table_name*

### CleanRoomAssetNotebook
* `notebook_content` (string, required) - Base 64 representation of the notebook contents.
  This is the same format as returned by :method:workspace/export with the format of **HTML**
* `runner_collaborator_aliases` (list of string, optional) - Aliases of collaborators that can run the notebook

### CleanRoomAssetTableLocalDetails
* `local_name` (string, required) - The fully qualified name of the table in its owner's local metastore,
  in the format of *catalog*.*schema*.*table_name*
* `partitions` (list of Partition, optional) - Partition filtering specification for a shared table

### CleanRoomAssetViewLocalDetails
* `local_name` (string, required) - The fully qualified name of the view in its owner's local metastore,
  in the format of *catalog*.*schema*.*view_name*

### CleanRoomAssetVolumeLocalDetails
* `local_name` (string, required) - The fully qualified name of the volume in its owner's local metastore,
  in the format of *catalog*.*schema*.*volume_name*

### CleanRoomNotebookReview
* `comment` (string, optional) - Review comment
* `created_at_millis` (integer, optional) - When the review was submitted, in epoch milliseconds
* `review_state` (string, optional) - Review outcome. Possible values are: `APPROVED`, `PENDING`, `REJECTED`
* `review_sub_reason` (string, optional) - Specified when the review was not explicitly made by a user. Possible values are: `AUTO_APPROVED`, `BACKFILLED`
* `reviewer_collaborator_alias` (string, optional) - Collaborator alias of the reviewer

### ColumnInfo
* `comment` (string, optional) - User-provided free-form text description
* `mask` (ColumnMask, optional)
* `name` (string, optional) - Name of Column
* `nullable` (boolean, optional) - Whether field may be Null (default: true)
* `partition_index` (integer, optional) - Partition index for column
* `position` (integer, optional) - Ordinal position of column (starting at position 0)
* `type_interval_type` (string, optional) - Format of IntervalType
* `type_json` (string, optional) - Full data type specification, JSON-serialized
* `type_name` (string, optional) - . Possible values are: `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `GEOGRAPHY`, `GEOMETRY`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TABLE_TYPE`, `TIMESTAMP`, `TIMESTAMP_NTZ`, `USER_DEFINED_TYPE`, `VARIANT`
* `type_precision` (integer, optional) - Digits of precision; required for DecimalTypes
* `type_scale` (integer, optional) - Digits to right of decimal; Required for DecimalTypes
* `type_text` (string, optional) - Full data type specification as SQL/catalogString text

### ColumnMask
* `function_name` (string, optional) - The full name of the column mask SQL UDF
* `using_column_names` (list of string, optional) - The list of additional table columns to be passed as input to the column mask function. The
  first arg of the mask function should be of the type of the column being masked and the
  types of the rest of the args should match the types of columns in 'using_column_names'

### Partition
* `values` (list of PartitionValue, optional) - An array of partition values

### PartitionValue
* `name` (string, optional) - The name of the partition column
* `op` (string, optional) - The operator to apply for the value. Possible values are: `EQUAL`, `LIKE`
* `recipient_property_key` (string, optional) - The key of a Delta Sharing recipient's property. For example "databricks-account-id".
  When this field is set, field `value` can not be set
* `value` (string, optional) - The value of the partition column. When this value is not set, it means `null` value.
  When this field is set, field `recipient_property_key` can not be set

## Attributes
In addition to the above arguments, the following attributes are exported:
* `added_at` (integer) - When the asset is added to the clean room, in epoch milliseconds
* `owner_collaborator_alias` (string) - The alias of the collaborator who owns this asset
* `status` (string) - Status of the asset. Possible values are: `ACTIVE`, `PENDING`, `PERMISSION_DENIED`

### CleanRoomAssetForeignTable
* `columns` (list of ColumnInfo) - The metadata information of the columns in the foreign table

### CleanRoomAssetNotebook
* `etag` (string) - Server generated etag that represents the notebook version
* `review_state` (string) - Top-level status derived from all reviews. Possible values are: `APPROVED`, `PENDING`, `REJECTED`
* `reviews` (list of CleanRoomNotebookReview) - All existing approvals or rejections

### CleanRoomAssetTable
* `columns` (list of ColumnInfo) - The metadata information of the columns in the table

### CleanRoomAssetView
* `columns` (list of ColumnInfo) - The metadata information of the columns in the view

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = clean_room_name,name,asset_type
  to = databricks_clean_room_asset.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_clean_room_asset clean_room_name,name,asset_type
```