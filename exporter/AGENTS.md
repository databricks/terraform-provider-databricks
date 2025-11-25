# Terraform Exporter

The exporter (`exporter/` directory) generates Terraform configuration (`.tf` files) and import scripts from existing Databricks resources.

## How it works under the hood

Let's look at how it works:

- When you start the Exporter, it allows you to select interactively which resources you want to export. You can also specify the resource types by using `-listing` and `-services` options that accept a comma-separated list of values (if nothing is specified, all resources are exported). The first option is used to list objects of specific resource classes, and the second option is used to handle transitive dependencies, such as permissions, notebooks referred to in jobs, etc.
- Next, the Exporter looks to determine what objects it can represent as Terraform. It first goes through all supported resource types and, for each of them, lists available objects and imports these objects' current state into corresponding [Terraform objects](https://developer.hashicorp.com/terraform/language/expressions/type-constraints#object). If necessary, the exporter also handles dependencies between resources, for example, importing permissions, triggering imports of the notebook(s) referred to in the job definition, etc. It also downloads external resources, such as libraries, init scripts, notebooks, etc., and stores them as files on the disk. external resources - libraries, init scripts, notebooks, …
- Finally, the Exporter generates the Terraform code and writes it to individual files corresponding to service types.

The logic of importing the individual resource type is isolated from the generation of the HCL (HashiCorp configuration language) code that is used by Terraform.  Each resource type is defined as an instance of the `importable` structure that is related to a specific Databricks Terraform resource. This structure consists of the following fields (usually we need to define only what is strictly required for each of the resource types - for most of the resources these are: `Service`, `List`, `Name`, `Import`, `WorkspaceLevel`/`AccountLevel`, and `Depends`):

- `Service` defines to which service this resource belongs (`compute`, `notebooks`, `access`, etc.). These values could be used with the `-services` command-line option to select only specific service types to export.
- `Name` is a function that generates a Terraform resource name for a specific instance of the resource.
- `List` is a function responsible for listing the objects of a given resource type.  This function just emits the object ID, and the actual import happens later.
- `Search` is a function that is used to find the actual ID of the object referred to by some attribute. For example, if the user is referred to by its display name.
- `Import` is a function that may emit additional objects after the object's state is read (read happens automatically).  Most often it's used to emit references to permissions, or, for example, to generate files on the disk for init scripts, etc.
- `Body` is a function that could be used to customize the HCL structure that will be put in the file.  But it's rarely required, only for a few complex cases, such as `dataricks_mount,` where we need to generate nested blocks based on the storage URL.
- `Ignore` is a function that checks if the current object needs to be ignored or not.
- `ShouldOmitField` is a function that checks if the specific field of the resource should be omitted from the generated HCL code. Primarily, it’s used for ignoring computed attributes or attributes with default values. For example, it’s used to omit the automatically generated `application_id` attribute for the `databricks_service_principal` resource on AWS & GCP.
- `Depends` defines how to extract references to other objects, guaranteeing the referential integrity of the generated Terraform code, as explained below.
- `ApiVersion` is an optional REST API version that should be used for the given resource. For example, we use REST API version 2.1 for the `databricks_job` resource. This attribute will not be necessary when we fully switch to using the Go SDK everywhere.
- `AccountLevel` defines if the specific service is an account-level resource.
- `WorkspaceLevel` defines if the specific service is a workspace-level resource.
- `PluginFramework` defines if the specific service is written using Terraform Plugin Framework.

## Adding a new resource to Exporter

When adding a new resource to Terraform Exporter we need to perform next steps:

1. Define a new `importable` instance in the `importables.go`.
2. Specify if it's account-level or workspace-level resource, or both.
3. Specify a service to which resource belongs to. Either use one of the existing, if it fits, or define a new one (ask user for confirmation).
4. Implement the `List` function that will be discover and emit instances of the specific resource.  When implementing it, prefer to use `List` method of Go SDK instead of `ListAll`.
5. (Optional) Implement the `Name` function that will extract TF resource name from an instance of a specific resource.
6. (Recommended) Implement the `Import` function that is responsible for emitting of dependencies for this resource - permissions/grants, etc.
7. (Optional) Implement the `ShouldOmitField` if some fields should be conditionally omitted.
8. (Recommended) Add `Depends` that describes relationships between fields of the current resource and its dependencies.
9. (Recommended) Add unit test that will validate the generated code, similar to `TestImportingLakeviewDashboards` or `TestNotificationDestinationExport` tests in `exporter_test.go` file.  For resources that use Go SDK, use `MockWorkspaceClientFunc` from Databricks Go SDK instead of `HTTPFixture`.
10. Update support matrix in `docs/guides/experimental-exporter.md` to indicate support for the new resource.  Keep list of supported resources sorted.
11. If new service name was introduced, add it to the corresponding section of the documentation. Keep it sorted alphabetically.

Recommendations:

- Use Go SDK as much as possible in `List` implementation and other places.
- Existing resources may emit newly implemented resource, i.e., `databricks_sql_table` should emit `databricks_quality_monitor` when it's added to Exporter.
- In some cases, references to dependencies could be ambiguous, i.e., there could be tables or schemas with the same name in different catalogs/schema. In this case we may need to add `IsValidApproximation` implementation.
- When there is a need to access a one or a few attributes in the Terraform resource data/state, use `.Get`, but if there is a need to access fields in nested structures, or have access to multiple fields, convert resource data/state into a Go SDK struct and use it.
- Refer [Databricks REST API documentation](https://docs.databricks.com/api/llms.txt) to understand a payload used in specific API.

**When Adding Exporter Support for resource implemented with Terraform plugin framework**:

1. Define resource in `exporter/importables.go` with `PluginFramework: true` for Plugin Framework resources
2. Use `convertPluginFrameworkToGoSdk` helper for Plugin Framework
3. Use unified callbacks (`ShouldOmitFieldUnified`, `ShouldGenerateFieldUnified`) for custom field logic


## Unified HCL Code Generation

The exporter uses a **unified code generation approach** for both SDKv2 and Plugin Framework resources:

**Entry Point**: `unifiedDataToHcl()` in `exporter/codegen.go`
- Works with both SDKv2 and Plugin Framework resources through abstraction layers
- Uses `ResourceDataWrapper` and `SchemaWrapper` interfaces for unified data access
- Extracts ~70% of common logic into shared helper functions

**Architecture**:
```
unifiedDataToHcl()
├── extractFieldsForGeneration() ← Shared: field iteration, omission, variable refs, skip logic
├── generateSdkv2Field() ← SDKv2-specific: generates blocks
├── generatePluginFrameworkField() ← Plugin Framework-specific: generates attributes
└── generateDependsOnAttribute() ← Shared: depends_on generation
```

**Resource Definition Callbacks** (`importable` struct in `exporter/model.go`):
- `ShouldOmitFieldUnified` - Unified callback for field omission (works with both frameworks)
- `ShouldGenerateFieldUnified` - Unified callback for field generation (works with both frameworks)
- `ShouldOmitField` - Legacy SDKv2-only callback (deprecated, use Unified version)
- `ShouldGenerateField` - Legacy SDKv2-only callback (deprecated, use Unified version)

**Key Differences**:
- SDKv2 generates nested structures as **blocks**: `evaluation { ... }`
- Plugin Framework generates nested structures as **attributes**: `evaluation = { ... }`

## Cloud Attribute Conversion

The exporter supports converting cloud-specific attributes when exporting for a different target cloud using the `-targetCloud` flag.

**Supported Resources:**
- `databricks_cluster`
- `databricks_job` (new_cluster in tasks and job_clusters)
- `databricks_pipeline`
- `databricks_instance_pool`

**Compatible Attributes:**
- `availability` - Automatically converted to target cloud format
- `first_on_demand` - Preserved as-is
- `zone_id` - Only "auto" value is portable between AWS and GCP
- `ebs_volume_count` ↔ `local_ssd_count` - Only with GENERAL_PURPOSE_SSD
- `disk_spec.disk_type` (instance pools only):
  - AWS `ebs_volume_type` ↔ Azure `azure_disk_volume_type`
    - `GENERAL_PURPOSE_SSD` ↔ `PREMIUM_LRS`
    - `THROUGHPUT_OPTIMIZED_HDD` ↔ `STANDARD_LRS`
  - GCP: `disk_type` not supported (removed when converting to GCP)
  - `disk_count` is preserved across all clouds

**Incompatible Attributes:**
Attributes without equivalents in the target cloud (e.g., `instance_profile_arn` on AWS) are omitted from the converted output.

**Implementation:**
Cloud conversion logic is in `exporter/cloud_utils.go` and integrates into the resource import pipeline before HCL generation.

**Policy for Incompatible Attributes:**
When converting cloud attributes, source blocks are ALWAYS cleared. If no compatible attributes exist for conversion, the resource will have no cloud attributes block in the generated HCL. Users should review the generated code and manually add cloud-specific attributes as needed.

**Important Limitations:**
- Only the attributes block is converted; instance types, zones, and other resource-specific configurations must be updated manually
- Instance profile ARNs and similar cloud-specific identifiers cannot be converted and will be omitted
- The exporter will log warnings when attributes cannot be converted

## Node Type Mapping

The exporter supports converting `node_type_id` and `driver_node_type_id` fields using an external JSON mapping file specified via the `-nodeTypeMappingFile` flag.

**Usage:**
```bash
databricks exporter -targetCloud azure -nodeTypeMappingFile node_mappings.json
```

**Mapping File Format:**
```json
{
  "version": "1.0",
  "mappings": [
    {
      "azure": "Standard_F4s",
      "aws": "i3.xlarge",
      "gcp": "n1-standard-4"
    }
  ]
}
```

**Features:**
- Bidirectional mappings: Single file works for all cloud combinations
- Converts both `node_type_id` and `driver_node_type_id` in clusters, jobs, and pipelines
- Preserves unmapped node types with warnings
- Can only be used together with `-targetCloud` flag

**Implementation:**
- Node type mapping logic is in `exporter/node_type_mapping.go`
- Integration happens in `cloud_utils.go` alongside cloud attribute conversion
- Conversion is applied at both top-level and nested cluster configurations

### Generating Node Type Mappings

The `exporter/generate_node_mappings.py` script generates node type mapping files from Databricks API responses.

**Obtaining Node Type Data:**

First, extract node type information from each cloud using the Databricks CLI or API:

```bash
# AWS workspace
databricks clusters list-node-types --profile aws-workspace > node-types-aws.json

# Azure workspace
databricks clusters list-node-types --profile azure-workspace > node-types-azure.json

# GCP workspace
databricks clusters list-node-types --profile gcp-workspace > node-types-gcp.json
```

**Running the Generator:**

```bash
cd exporter
python3 generate_node_mappings.py \
  --aws node-types-aws.json \
  --azure node-types-azure.json \
  --gcp node-types-gcp.json \
  --output node_type_mapping.json
```

**Algorithm:**

The script uses a two-phase approach to maximize mapping coverage:

1. **Exact Matching Phase** (minimum similarity score: 150)
   - Core count must match exactly (required)
   - Memory within 90% tolerance (weighted: 50 points)
   - Category matching: memory/highmem, compute, storage, general (weighted: 30 points)
   - Disk configuration similarity (weighted: 20 points)

2. **Approximate Matching Phase**
   - For mappings missing one or two clouds, finds approximate matches
   - Uses existing cloud mapping characteristics as reference
   - Requires same core count, accepts lower similarity scores
   - Results in ~99.5% complete three-way mappings (AWS + Azure + GCP)

**Output:**

The script generates a bidirectional JSON mapping file with statistics:
- Total mappings: typically 800-850
- Complete three-way mappings: ~99.5%
- Incomplete mappings: ~0.5% (usually high-core instances without cross-cloud equivalents)

**Filters Applied:**

The script automatically filters out:
- Deprecated node types (`is_deprecated: true`)
- Hidden node types (`is_hidden: true`)
- GPU-enabled instances (handled separately)
