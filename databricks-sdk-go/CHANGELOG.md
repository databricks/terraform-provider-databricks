# Version changelog

## Release v0.204.0 (2026-11-25)

### Bug Fixes
* Callers waiting for a cached OAuth token can cancel their wait through the request context. A refresh already in progress can continue in the background and populate the cache. Blocking and proactive refresh attempts have a one-minute deadline.

### API Changes
* Add [geniemigrate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/geniemigrate) package.
* Add [w.GenieMigrate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/geniemigrate#GenieMigrateAPI) workspace-level service.
* Add `Resources` field for [genieagents.GenieAgentDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents#GenieAgentDefinition).
* Add `TextInstruction` field for [genieagents.GenieAgentInstructions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents#GenieAgentInstructions).
* Add `AssetTypeMetricView` and `AssetTypeAgentSkill` enum values for [marketplace.AssetType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#AssetType).
* [Breaking] Remove `DataSources` field for [genieagents.GenieAgentDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents#GenieAgentDefinition).
* [Breaking] Remove `SqlFunctions` and `TextInstructions` fields for [genieagents.GenieAgentInstructions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents#GenieAgentInstructions).
* [Breaking] Remove `Notifications` field for [ml.MaterializationSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializationSet).

## Release v0.203.0 (2026-11-24)

### API Changes
* Add `CreateComputePolicy`, `DeleteComputePolicy`, `GetComputePolicy` and `UpdateComputePolicy` methods for [w.ServerlessComputeService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ServerlessComputeServiceAPI) workspace-level service.
* Add `RescanCatalog` method for [w.DataClassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#DataClassificationAPI) workspace-level service.
* Add `UrlConfiguration` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations).
* Add `Streaming` field for [lakewatch.Rule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#Rule).
* Add `Lookup` and `SecondaryKeys` fields for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `ComputePolicyId` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `Network` field for [sandbox.SandboxSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxSpec).
* Add `Finalizing` enum value for [aisearch.EndpointScalingInfoPhase](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#EndpointScalingInfoPhase).
* Add `GoogleAds` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `Preemptable` enum value for [jobs.AiRuntimeTaskPriorityClass](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskPriorityClass).
* Add `Finalizing` enum value for [vectorsearch.EndpointScalingInfoScalingPhase](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfoScalingPhase).
* Change `Schedule` field for [lakewatch.Rule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#Rule) to no longer be required.
* [Breaking] Change `Schedule` field for [lakewatch.Rule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#Rule) to no longer be required.

## Release v0.202.0 (2026-11-19)

### API Changes
* Add `DeleteSoftwareArtifact` and `DeleteSoftwareArtifactVersion` methods for [w.SoftwareArtifacts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/softwareartifacts#SoftwareArtifactsAPI) workspace-level service.
* Add `High` enum value for [jobs.AiRuntimeTaskPriorityClass](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskPriorityClass).
* [Breaking] Remove `BestEffort` enum value for [jobs.AiRuntimeTaskPriorityClass](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskPriorityClass).

## Release v0.201.0 (2026-11-18)

### API Changes
* Add `CreateParser`, `CreateParserVersion`, `DeleteParser`, `DeleteParserVersion`, `GetBuiltinParser`, `GetBuiltinParserVersion`, `GetParser`, `GetParserVersion`, `ListBuiltinParserVersions`, `ListBuiltinParsers`, `ListParserVersions` and `ListParsers` methods for [w.Lakewatch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#LakewatchAPI) workspace-level service.
* Add `DetailedState` field for [aisearch.IndexStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#IndexStatus).
* Add `Operations` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `SkillMaxResults` field for [sharing.ListProviderShareAssetsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderShareAssetsRequest).
* Add `Skills` field for [sharing.ListProviderShareAssetsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderShareAssetsResponse).
* Add `DetailedState` and `FailedStatus` fields for [vectorsearch.VectorIndexStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndexStatus).
* Add `GoogleAds` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Skill` enum value for [sharing.SharedDataObjectDataObjectType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObjectDataObjectType).
* Change `SqlQueryId` field for [jobs.SqlConditionConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlConditionConfiguration) to no longer be required.
* [Breaking] Change `SqlQueryId` field for [jobs.SqlConditionConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlConditionConfiguration) to no longer be required.
* [Breaking] Remove `MigrationTypeOnline` enum value for [postgres.MigrationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#MigrationType).

## Release v0.200.0 (2026-11-16)

### API Changes
* Add [statementexecutionv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/statementexecutionv2) package.
* Add [w.StatementExecutionV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/statementexecutionv2#StatementExecutionV2API) workspace-level service.
* Add `GetAppDeploymentScanFinding`, `ListAppDeploymentScanFindings` and `UpdateAppDeploymentScanFinding` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service.
* Add `ScanOmittedCountIsLowerBound`, `ScanOmittedFindingsCount` and `ScanStatus` fields for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment).
* Add `ImageUri` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Add `IcebergTableSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `Testing` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* Add `Testing` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `Testing` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `Testing` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Add `TestOnly` and `TestOptions` fields for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `TestOnly` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* Add `Overwrite` field for [sandbox.UploadFileRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#UploadFileRequest).
* Add `Day` enum value for [billing.AlertConfigurationTimePeriod](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfigurationTimePeriod).

## Release v0.199.0 (2026-11-12)

### API Changes
* Add [w.AssumeGroup](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AssumeGroupAPI) workspace-level service.
* Add `EffectiveResources` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* [Breaking] Change `OutputRef` field for [aifunctions.Batch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#Batch) to be required.
* Change `OutputRef` field for [aifunctions.Batch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#Batch) to be required.
* [Breaking] Remove `Continuous` field for [lakewatch.RuleSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#RuleSchedule).
* [Breaking] Remove `Expired` enum value for [aifunctions.BatchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#BatchStatus).

## Release v0.198.0 (2026-11-11)

### API Changes
* Add `PtTerm` field for [serving.PtServedModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PtServedModel).
* [Breaking] Remove `CustomContentPackPath` and `CustomPresetsPath` fields for [lakewatch.CommonConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#CommonConfig).

## Release v0.197.0 (2026-11-10)

### API Changes
* Add `OutputTypeText` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `ServicePrincipal` field for [postgres.CdfConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CdfConfig).
* Add `GroupName` field for [postgres.GenerateDatabaseCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#GenerateDatabaseCredentialRequest).
* [Breaking] Change `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to no longer be required.
* Change `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to no longer be required.

## Release v0.196.0 (2026-11-09)

### API Changes
* Add [softwareartifacts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/softwareartifacts) package.
* Add [w.ServerlessComputeService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ServerlessComputeServiceAPI) workspace-level service.
* Add [w.SoftwareArtifacts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/softwareartifacts#SoftwareArtifactsAPI) workspace-level service.
* Add `TriggerMaterializationSet` method for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `AcknowledgeAssessmentWarnings`, `AssessMigration`, `CreateMigration`, `DeleteMigration`, `GetAssessment`, `GetMigration`, `ListAssessments`, `ListMigrations`, `StartMigration`, `UndeleteMigration` and `UpdateMigration` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `DelayedReason`, `EstimatedCompletionTime`, `IndexInitializationPercentage`, `Phase` and `TargetQpsChangeAllowedAfterTime` fields for [aisearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#EndpointScalingInfo).
* Add `DeltaTable` field for [jobs.ResolvedValuesAiFunctionTaskResolvedValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedValuesAiFunctionTaskResolvedValues).
* Add `Paused` field for [ml.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CronSchedule).
* Add `ExternalTrigger` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `Paused` field for [ml.TableTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TableTrigger).
* Add `EffectiveWorkspaceDashboardTheme` and `WorkspaceDashboardTheme` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Add `DelayedReason`, `EstimatedCompletionTimestamp`, `IndexInitializationPercentage`, `Phase` and `TargetQpsChangeAllowedAfterTimestamp` fields for [vectorsearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfo).
* [Breaking] Change `WarehouseId` field for [lakewatch.Datasource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#Datasource) to no longer be required.
* Change `WarehouseId` field for [lakewatch.Datasource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#Datasource) to no longer be required.

## Release v0.195.0 (2026-11-05)

### API Changes
* Add `UndropMetastoreById` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.

## Release v0.194.0 (2026-11-04)

### API Changes
* Add [lakewatch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch) package.
* Add [w.Lakewatch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/lakewatch#LakewatchAPI) workspace-level service.
* Add `CancelBatch`, `CreateBatch`, `GetBatch` and `ListBatches` methods for [w.AiFunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#AiFunctionsAPI) workspace-level service.
* Add `CreateFeatureSpec`, `DeleteFeatureSpec` and `GetFeatureSpec` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* [Breaking] Remove `Critical` enum value for [jobs.AiRuntimeTaskPriorityClass](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskPriorityClass).

## Release v0.193.0 (2026-11-03)

### API Changes
* Add `ViewName` field for [ml.IngestionDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#IngestionDestination).
* Add `Notifications` field for [ml.MaterializationSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializationSet).

## Release v0.192.0 (2026-11-02)

### API Changes
* Add [computemigrations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/computemigrations) and [genieagents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents) packages.
* Add [w.ComputeMigrations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/computemigrations#ComputeMigrationsAPI) workspace-level service.
* Add [w.GenieAgents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/genieagents#GenieAgentsAPI) workspace-level service.
* Add `DownloadFile` and `UploadFile` methods for [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `UpdateNetworkConnectivityConfig` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* [Breaking] Add `Parent` field for [iamv2.CreateDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateDirectGroupMemberProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.CreateDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateDirectGroupMemberRequest).
* Add `Parent` field for [iamv2.CreateGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateGroupProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.CreateGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateGroupRequest).
* Add `Parent` field for [iamv2.CreateServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateServicePrincipalProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.CreateServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateServicePrincipalRequest).
* Add `Parent` field for [iamv2.CreateUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateUserProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.CreateUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateUserRequest).
* Add `Parent` field for [iamv2.CreateWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateWorkspaceAssignmentProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.CreateWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateWorkspaceAssignmentRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteDirectGroupMemberProxyRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteDirectGroupMemberRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteGroupProxyRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteGroupRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteServicePrincipalProxyRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteServicePrincipalRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteUserProxyRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteUserRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteWorkspaceAssignmentProxyRequest).
* [Breaking] Add `Name` field for [iamv2.DeleteWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteWorkspaceAssignmentRequest).
* Add `Name` and `Parent` fields for [iamv2.DirectGroupMember](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DirectGroupMember).
* [Breaking] Add `Name` field for [iamv2.GetDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetDirectGroupMemberProxyRequest).
* [Breaking] Add `Name` field for [iamv2.GetDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetDirectGroupMemberRequest).
* [Breaking] Add `Name` field for [iamv2.GetGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetGroupProxyRequest).
* [Breaking] Add `Name` field for [iamv2.GetGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetGroupRequest).
* [Breaking] Add `Name` field for [iamv2.GetServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetServicePrincipalProxyRequest).
* [Breaking] Add `Name` field for [iamv2.GetServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetServicePrincipalRequest).
* [Breaking] Add `Name` field for [iamv2.GetUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetUserProxyRequest).
* [Breaking] Add `Name` field for [iamv2.GetUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetUserRequest).
* [Breaking] Add `Name` field for [iamv2.GetWorkspaceAccessDetailLocalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAccessDetailLocalRequest).
* [Breaking] Add `Name` field for [iamv2.GetWorkspaceAccessDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAccessDetailRequest).
* [Breaking] Add `Name` field for [iamv2.GetWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAssignmentProxyRequest).
* [Breaking] Add `Name` field for [iamv2.GetWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAssignmentRequest).
* [Breaking] Add `Name` field for [iamv2.GetWorkspaceIdentityDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceIdentityDetailRequest).
* Add `Name` and `Parent` fields for [iamv2.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#Group).
* [Breaking] Add `Parent` field for [iamv2.ListDirectGroupMembersProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListDirectGroupMembersProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListDirectGroupMembersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListDirectGroupMembersRequest).
* Add `Parent` field for [iamv2.ListGroupsProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListGroupsProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListGroupsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListGroupsRequest).
* Add `Parent` field for [iamv2.ListServicePrincipalsProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListServicePrincipalsProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListServicePrincipalsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListServicePrincipalsRequest).
* [Breaking] Add `Parent` field for [iamv2.ListTransitiveParentGroupsProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListTransitiveParentGroupsProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListTransitiveParentGroupsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListTransitiveParentGroupsRequest).
* Add `Parent` field for [iamv2.ListUsersProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListUsersProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListUsersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListUsersRequest).
* Add `Parent` field for [iamv2.ListWorkspaceAssignmentsProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListWorkspaceAssignmentsProxyRequest).
* [Breaking] Add `Parent` field for [iamv2.ListWorkspaceAssignmentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListWorkspaceAssignmentsRequest).
* Add `Name` and `Parent` fields for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal).
* Add `Name` field for [iamv2.TransitiveParentGroup](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#TransitiveParentGroup).
* [Breaking] Add `Name` field for [iamv2.UpdateGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateGroupProxyRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateGroupRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateServicePrincipalProxyRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateServicePrincipalRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateUserProxyRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateUserRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceAssignmentProxyRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceAssignmentRequest).
* [Breaking] Add `Name` field for [iamv2.UpdateWorkspaceIdentityDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceIdentityDetailRequest).
* Add `Name` and `Parent` fields for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).
* Add `Name` field for [iamv2.WorkspaceAccessDetail](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceAccessDetail).
* Add `Name` and `Parent` fields for [iamv2.WorkspaceAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceAssignment).
* Add `Name` field for [iamv2.WorkspaceIdentityDetail](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIdentityDetail).
* Add `ServerlessCompute` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster).
* Add `JobCompletion` field for [jobs.PerTriggerState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerTriggerState).
* Add `AiFunctionTask` field for [jobs.ResolvedValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedValues).
* Add `AiFunctionTaskOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Add `AiFunctionTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `AiFunctionTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `AiFunctionTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `JobCompletion` field for [jobs.TriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerConfiguration).
* Add `JobCompletion` field for [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
* Add `JobCompletion` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* Add `JobCompletion` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Add `BillingWorkspaceId` field for [settings.NetworkConnectivityConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityConfiguration).
* Add `CreateFeature` and `ReadFeature` enum values for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Add `Feature` enum value for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Add `JobCompletion` enum value for [jobs.TriggerType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerType).
* [Breaking] Change `CreateDirectGroupMember`, `CreateGroup`, `CreateServicePrincipal`, `CreateUser`, `CreateWorkspaceAssignment`, `DeleteDirectGroupMember`, `DeleteGroup`, `DeleteServicePrincipal`, `DeleteUser`, `DeleteWorkspaceAssignment`, `GetDirectGroupMember`, `GetGroup`, `GetServicePrincipal`, `GetUser`, `GetWorkspaceAccessDetail`, `GetWorkspaceAssignment`, `ListDirectGroupMembers`, `ListGroups`, `ListServicePrincipals`, `ListTransitiveParentGroups`, `ListUsers`, `ListWorkspaceAssignments`, `UpdateGroup`, `UpdateServicePrincipal`, `UpdateUser` and `UpdateWorkspaceAssignment` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service with new required argument order.
* [Breaking] Change `CreateDirectGroupMember`, `CreateGroup`, `CreateServicePrincipal`, `CreateUser`, `CreateWorkspaceAssignment`, `DeleteDirectGroupMember`, `DeleteGroup`, `DeleteServicePrincipal`, `DeleteUser`, `DeleteWorkspaceAssignment`, `GetDirectGroupMember`, `GetGroup`, `GetServicePrincipal`, `GetUser`, `GetWorkspaceAccessDetail`, `GetWorkspaceAssignment`, `ListDirectGroupMembers`, `ListGroups`, `ListServicePrincipals`, `ListTransitiveParentGroups`, `ListUsers`, `ListWorkspaceAssignments`, `UpdateGroup`, `UpdateServicePrincipal`, `UpdateUser` and `UpdateWorkspaceAssignment` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service. Method path has changed.
* [Breaking] Change `CreateDirectGroupMemberProxy`, `DeleteDirectGroupMemberProxy`, `DeleteGroupProxy`, `DeleteServicePrincipalProxy`, `DeleteUserProxy`, `DeleteWorkspaceAssignmentProxy`, `GetDirectGroupMemberProxy`, `GetGroupProxy`, `GetServicePrincipalProxy`, `GetUserProxy`, `GetWorkspaceAccessDetailLocal`, `GetWorkspaceAssignmentProxy`, `GetWorkspaceIdentityDetail`, `ListDirectGroupMembersProxy`, `ListTransitiveParentGroupsProxy`, `UpdateGroupProxy`, `UpdateServicePrincipalProxy`, `UpdateUserProxy`, `UpdateWorkspaceAssignmentProxy` and `UpdateWorkspaceIdentityDetail` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service with new required argument order.
* [Breaking] Change `CreateDirectGroupMemberProxy`, `DeleteDirectGroupMemberProxy`, `DeleteGroupProxy`, `DeleteServicePrincipalProxy`, `DeleteUserProxy`, `DeleteWorkspaceAssignmentProxy`, `GetDirectGroupMemberProxy`, `GetGroupProxy`, `GetServicePrincipalProxy`, `GetUserProxy`, `GetWorkspaceAccessDetailLocal`, `GetWorkspaceAssignmentProxy`, `GetWorkspaceIdentityDetail`, `ListDirectGroupMembersProxy`, `ListTransitiveParentGroupsProxy`, `UpdateGroupProxy`, `UpdateServicePrincipalProxy`, `UpdateUserProxy`, `UpdateWorkspaceAssignmentProxy` and `UpdateWorkspaceIdentityDetail` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service. Method path has changed.
* [Breaking] Remove `GroupId` field for [iamv2.CreateDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateDirectGroupMemberProxyRequest).
* [Breaking] Remove `AccountId` and `GroupId` fields for [iamv2.CreateDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateDirectGroupMemberRequest).
* [Breaking] Remove `AccountId` field for [iamv2.CreateGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateGroupRequest).
* [Breaking] Remove `AccountId` field for [iamv2.CreateServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateServicePrincipalRequest).
* [Breaking] Remove `AccountId` field for [iamv2.CreateUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateUserRequest).
* [Breaking] Remove `AccountId` and `WorkspaceId` fields for [iamv2.CreateWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#CreateWorkspaceAssignmentRequest).
* [Breaking] Remove `GroupId` and `PrincipalId` fields for [iamv2.DeleteDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteDirectGroupMemberProxyRequest).
* [Breaking] Remove `AccountId`, `GroupId` and `PrincipalId` fields for [iamv2.DeleteDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteDirectGroupMemberRequest).
* [Breaking] Remove `GroupId` field for [iamv2.DeleteGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteGroupProxyRequest).
* [Breaking] Remove `AccountId` and `GroupId` fields for [iamv2.DeleteGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteGroupRequest).
* [Breaking] Remove `ServicePrincipalId` field for [iamv2.DeleteServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteServicePrincipalProxyRequest).
* [Breaking] Remove `AccountId` and `ServicePrincipalId` fields for [iamv2.DeleteServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteServicePrincipalRequest).
* [Breaking] Remove `UserId` field for [iamv2.DeleteUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteUserProxyRequest).
* [Breaking] Remove `AccountId` and `UserId` fields for [iamv2.DeleteUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteUserRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.DeleteWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteWorkspaceAssignmentProxyRequest).
* [Breaking] Remove `AccountId`, `WorkspaceId` and `PrincipalId` fields for [iamv2.DeleteWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DeleteWorkspaceAssignmentRequest).
* [Breaking] Remove `GroupId` and `PrincipalId` fields for [iamv2.GetDirectGroupMemberProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetDirectGroupMemberProxyRequest).
* [Breaking] Remove `AccountId`, `GroupId` and `PrincipalId` fields for [iamv2.GetDirectGroupMemberRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetDirectGroupMemberRequest).
* [Breaking] Remove `GroupId` field for [iamv2.GetGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetGroupProxyRequest).
* [Breaking] Remove `AccountId` and `GroupId` fields for [iamv2.GetGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetGroupRequest).
* [Breaking] Remove `ServicePrincipalId` field for [iamv2.GetServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetServicePrincipalProxyRequest).
* [Breaking] Remove `AccountId` and `ServicePrincipalId` fields for [iamv2.GetServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetServicePrincipalRequest).
* [Breaking] Remove `UserId` field for [iamv2.GetUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetUserProxyRequest).
* [Breaking] Remove `AccountId` and `UserId` fields for [iamv2.GetUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetUserRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.GetWorkspaceAccessDetailLocalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAccessDetailLocalRequest).
* [Breaking] Remove `AccountId`, `WorkspaceId` and `PrincipalId` fields for [iamv2.GetWorkspaceAccessDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAccessDetailRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.GetWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAssignmentProxyRequest).
* [Breaking] Remove `AccountId`, `WorkspaceId` and `PrincipalId` fields for [iamv2.GetWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceAssignmentRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.GetWorkspaceIdentityDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#GetWorkspaceIdentityDetailRequest).
* [Breaking] Remove `GroupId` field for [iamv2.ListDirectGroupMembersProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListDirectGroupMembersProxyRequest).
* [Breaking] Remove `AccountId` and `GroupId` fields for [iamv2.ListDirectGroupMembersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListDirectGroupMembersRequest).
* [Breaking] Remove `AccountId` field for [iamv2.ListGroupsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListGroupsRequest).
* [Breaking] Remove `AccountId` field for [iamv2.ListServicePrincipalsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListServicePrincipalsRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.ListTransitiveParentGroupsProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListTransitiveParentGroupsProxyRequest).
* [Breaking] Remove `AccountId` and `PrincipalId` fields for [iamv2.ListTransitiveParentGroupsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListTransitiveParentGroupsRequest).
* [Breaking] Remove `AccountId` field for [iamv2.ListUsersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListUsersRequest).
* [Breaking] Remove `AccountId` and `WorkspaceId` fields for [iamv2.ListWorkspaceAssignmentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ListWorkspaceAssignmentsRequest).
* [Breaking] Remove `GroupId` field for [iamv2.UpdateGroupProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateGroupProxyRequest).
* [Breaking] Remove `AccountId` and `GroupId` fields for [iamv2.UpdateGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateGroupRequest).
* [Breaking] Remove `ServicePrincipalId` field for [iamv2.UpdateServicePrincipalProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateServicePrincipalProxyRequest).
* [Breaking] Remove `AccountId` and `ServicePrincipalId` fields for [iamv2.UpdateServicePrincipalRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateServicePrincipalRequest).
* [Breaking] Remove `UserId` field for [iamv2.UpdateUserProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateUserProxyRequest).
* [Breaking] Remove `AccountId` and `UserId` fields for [iamv2.UpdateUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateUserRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.UpdateWorkspaceAssignmentProxyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceAssignmentProxyRequest).
* [Breaking] Remove `AccountId`, `WorkspaceId` and `PrincipalId` fields for [iamv2.UpdateWorkspaceAssignmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceAssignmentRequest).
* [Breaking] Remove `PrincipalId` field for [iamv2.UpdateWorkspaceIdentityDetailRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#UpdateWorkspaceIdentityDetailRequest).

## Release v0.191.1 (2026-10-28)

### Bug Fixes
* Fix Databricks CLI authentication rejecting small wrappers around the CLI (e.g. launcher scripts or version manager shims) as the legacy Python CLI. Small files found in `PATH` are now only rejected if they are the entry point pip installs for the legacy Python CLI, and a path set in `DATABRICKS_CLI_PATH` (or `databricks_cli_path`) is used without this check.

## Release v0.191.0 (2026-10-27)

### API Changes
* Add `GetMemoryStorePermission`, `GetSessionStorePermission`, `GrantMemoryStorePermission`, `GrantSessionStorePermission`, `ListMemoryStorePermissions`, `ListSessionStorePermissions`, `RevokeMemoryStorePermission` and `RevokeSessionStorePermission` methods for [w.AgentKit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#AgentKitAPI) workspace-level service.
* Add `GetBudgetSpendStatus` and `ListBudgetSpendStatuses` methods for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service.
* Add `CreateMaterializationSet`, `DeleteMaterializationSet`, `GetMaterializationSet` and `UpdateMaterializationSet` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `StartEndpoint` and `StopEndpoint` methods for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `IncludeExternalSpend` field for [billing.BudgetConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfiguration).
* Add `IncludeExternalSpend` field for [billing.CreateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudget).
* Add `IncludeExternalSpend` field for [billing.UpdateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationBudget).
* Add `NotificationSettings` field for [dataquality.AnomalyDetectionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#AnomalyDetectionConfig).
* Add `DestinationIds` and `DestinationUserIds` fields for [dataquality.NotificationDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#NotificationDestination).
* Add `OnUnhealthyTable` field for [dataquality.NotificationSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#NotificationSettings).
* Add `Regex` field for [iamv2.IdentityVisibilityFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#IdentityVisibilityFilter).
* Add `FileOutput` field for [jobs.SqlOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlOutput).
* Add `RelativeError` field for [ml.ApproxCountDistinctFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ApproxCountDistinctFunction).
* Add `RelativeError` field for [ml.ApproxPercentileFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ApproxPercentileFunction).
* Add `MaterializationSetId` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `DeleteRows` field for [ml.PurgeFeatureEntitiesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PurgeFeatureEntitiesRequest).
* Add `Condition`, `MinTimeBetweenTriggers` and `WaitAfterLastChange` fields for [ml.TableTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TableTrigger).
* Add `OnedriveOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `FullRefreshSelectionByPath` and `RefreshSelectionByPath` fields for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `FullRefreshSelectionByPath` and `RefreshSelectionByPath` fields for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* Add `Metadata` field for [sandbox.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#Sandbox).
* Add `Stopping`, `Stopped` and `Starting` enum values for [aisearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#EndpointStatusState).
* Add `OperationActionTypeUnbind` enum value for [bundledeployments.OperationActionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#OperationActionType).
* Add `Onedrive` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Stopping`, `Stopped` and `Starting` enum values for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).

## Release v0.190.0 (2026-10-21)

### API Changes
* Add `SdpSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `CustomerVersionLabel`, `MaxDbu` and `ReydenClusterSizeMultiplier` fields for [sql.CreateWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWarehouseRequest).
* Add `CustomerVersionLabel`, `MaxDbu` and `ReydenClusterSizeMultiplier` fields for [sql.EditWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EditWarehouseRequest).
* Add `CustomerVersionLabel`, `MaxDbu` and `ReydenClusterSizeMultiplier` fields for [sql.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EndpointInfo).
* Add `CustomerVersionLabel`, `MaxDbu` and `ReydenClusterSizeMultiplier` fields for [sql.GetWarehouseResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehouseResponse).
* [Breaking] Remove `ReplicationConfig` field for [sharing.CreateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateShare).
* [Breaking] Remove `ReplicationConfig` field for [sharing.ShareInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ShareInfo).

## Release v0.189.0 (2026-10-19)

### API Changes
* Add `CreateIssue`, `GetIssue`, `SearchIssues` and `UpdateIssue` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* Add `RevealRolePassword` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `CreateDirectory`, `DeletePath`, `GetMetadata`, `ListDirectoryContents` and `RenamePath` methods for [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `View` field for [agentkit.GetManagedMemoryEntryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#GetManagedMemoryEntryRequest).
* Add `View` field for [agentkit.ListManagedMemoryEntriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#ListManagedMemoryEntriesRequest).
* Add `View` field for [agentkit.SearchManagedMemoryEntriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#SearchManagedMemoryEntriesRequest).
* Add `ShowDeleted` field for [bundledeployments.ListDeploymentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#ListDeploymentsRequest).
* Add `Secrets` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* Add `Secrets` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Add `GovernedTags` field for [catalog.PolicyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PolicyInfo).
* Add `Secrets` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* Add `HasReplicationPointBlockingErrors` and `InitialReplicationProgress` fields for [disasterrecovery.FailoverGroup](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#FailoverGroup).
* Add `DurationTarget` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Add `DurationTarget` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Add `DurationTarget` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Add `DurationTarget` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
* Add `DurationTarget` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Add `DurationTarget` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
* Add `DurationTarget` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Add `Attachment` field for [jobs.Subscription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Subscription).
* Add `Scopes` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest).
* Add `Scopes` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse).
* Add `Scopes` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo).
* Add `ReplicationConfig` field for [sharing.CreateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateShare).
* Add `ReplicationConfig` field for [sharing.ShareInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ShareInfo).
* Add `Realtime` enum value for [sql.CreateWarehouseRequestWarehouseType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWarehouseRequestWarehouseType).
* Add `Realtime` enum value for [sql.EditWarehouseRequestWarehouseType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EditWarehouseRequestWarehouseType).
* Add `Realtime` enum value for [sql.EndpointInfoWarehouseType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EndpointInfoWarehouseType).
* Add `Realtime` enum value for [sql.GetWarehouseResponseWarehouseType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehouseResponseWarehouseType).
* Add `Realtime` enum value for [sql.WarehouseTypePairWarehouseType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseTypePairWarehouseType).
* [Breaking] Remove `ReadMask` field for [agentkit.GetManagedMemoryEntryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#GetManagedMemoryEntryRequest).
* [Breaking] Remove `ReadMask` field for [agentkit.ListManagedMemoryEntriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#ListManagedMemoryEntriesRequest).
* [Breaking] Remove `ReadMask` field for [agentkit.SearchManagedMemoryEntriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#SearchManagedMemoryEntriesRequest).
* [Breaking] Remove `EffectiveServerlessComputeId` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).

## Release v0.188.0 (2026-10-12)

### API Changes
* Add `BedrockMantle` and `Pricing` fields for [catalog.ModelProviderServiceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfig).
* Add `DatabricksAppSlack` and `DatabricksAppTeams` fields for [settings.Config](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#Config).
* Add `ModelService` enum value for [apps.AppManifestAppResourceUcSecurableSpecUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurableType).
* Add `ModelService` enum value for [apps.AppResourceUcSecurableUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurableType).
* Add `ExternalModelProviderTypeBedrockMantle` enum value for [catalog.ModelProviderServiceConfigExternalModelProviderType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigExternalModelProviderType).
* Add `DatabricksAppSlack` and `DatabricksAppTeams` enum values for [settings.DestinationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DestinationType).

## Release v0.187.0 (2026-10-08)

### API Changes
* Add `CreateIdentityVisibilityFilter`, `DeleteIdentityVisibilityFilter`, `GetIdentityVisibilityFilter` and `ListIdentityVisibilityFilters` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `EndpointRoute` field for [catalog.ModelProviderServiceConfigModelTargetConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigModelTargetConfig).
* Add `Notifications` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `ShufflePartitions` field for [ml.StreamingMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamingMode).
* Add `ParentPath` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* Add `ParentPath` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `ParentPath` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `ParentPath` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Add `CustomTemplateFormat` field for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2).
* Add `StatementTimeout` field for [sql.CreateWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWarehouseRequest).
* Add `StatementTimeout` field for [sql.EditWarehouseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EditWarehouseRequest).
* Add `StatementTimeout` field for [sql.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EndpointInfo).
* Add `StatementTimeout` field for [sql.GetWarehouseResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehouseResponse).


## Release v0.186.0 (2026-10-07)

### API Changes
* Add [agentkit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit) package.
* Add [w.AgentKit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentkit#AgentKitAPI) workspace-level service.
* Add `HealthCheck` field for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment).
* [Breaking] Remove [mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason) package.
* [Breaking] Remove [w.Mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason#MasonAPI) workspace-level service.


## Release v0.185.0 (2026-10-05)

### API Changes
* Add [a.PrivateNetworkGateways](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking#PrivateNetworkGatewaysAPI) account-level service.
* Add `DataframeSchema` field for [ml.RequestSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RequestSource).


## Release v0.184.0 (2026-10-02)

### API Changes
* Add [mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason) package.
* Add [w.Mason](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/mason#MasonAPI) workspace-level service.
* Add `AiDecide` method for [w.AiFunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#AiFunctionsAPI) workspace-level service.
* Add `CreateSkill`, `DeleteSkill`, `FinalizeSkill`, `GetSkill`, `ListSkills` and `UpdateSkill` methods for [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `ListCommands` method for [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `ServiceCredential` field for [catalog.ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig).
* Add `SecretReference` field for [catalog.ModelProviderServiceConfigProviderSecret](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigProviderSecret).
* Add `ProjectEnvironment` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Add `EnvironmentVariables` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Add `EnvironmentVariables` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
* Add `EnvironmentVariables` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Add `EnvironmentVariables` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Add `EnvironmentVariablesKey` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `EnvironmentVariables` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Add `EnvironmentVariablesKey` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `EnvironmentVariablesKey` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
* Add `OnMaintenanceComplete` and `OnMaintenanceStart` fields for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
* Add `BudgetPolicyId` and `Tags` fields for [ml.BackfillFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#BackfillFeaturesRequest).
* Add `JobId` and `PipelineId` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `BudgetPolicyId` and `Tags` fields for [ml.PurgeFeatureEntitiesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PurgeFeatureEntitiesRequest).
* Add `Environment` field for [sandbox.SandboxSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxSpec).
* Add `TableForeignDeltaDeltasharing` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* [Breaking] Change `CommandPath` field for [jobs.DeploymentSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DeploymentSpec) to no longer be required.
* Change `CommandPath` field for [jobs.DeploymentSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DeploymentSpec) to no longer be required.
* [Breaking] Change `ApiSecretRef` field for [ml.SchemaRegistryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaRegistryConfig) to no longer be required.
* Change `ApiSecretRef` field for [ml.SchemaRegistryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaRegistryConfig) to no longer be required.


## Release v0.183.0 (2026-09-30)

### New Features and Improvements

- Add `Config.Headers`, a hook called on every request made by the client to set custom HTTP headers, on top of normal authentication. Use the `StaticHeaders` helper to set a fixed set of headers.


## Release v0.182.0 (2026-09-21)

### API Changes
* Add `CreateMcpServiceUserMappedCredential`, `DeleteMcpServiceUserMappedCredential` and `GetMcpServiceUserMappedCredential` methods for [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `TelemetryExportDestinations` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `Options` field for [catalog.McpServiceConfigSourceConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpServiceConfigSourceConnection).
* Add `HeaderAuth` field for [catalog.ModelProviderServiceConfigCustomProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigCustomProviderDirectConfig).
* Add `TimezoneId` field for [ml.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CronSchedule).
* Add `AvroOptions` and `ProtobufOptions` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `TiktokAds` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `AgentService` and `Skill` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Add `TiktokAds` and `Smartsheet` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Avro` and `Protobuf` enum values for [pipelines.TransformerFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TransformerFormat).


## Release v0.181.0 (2026-09-17)

### Internal Changes

- Reduce integration test cluster usage by replacing waiter coverage with HTTP fixtures and removing redundant live tests.


## Release v0.180.0 (2026-09-16)

### API Changes
* Add `GroupName` field for [pipelines.RunAs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#RunAs).


## Release v0.179.0 (2026-09-14)

### API Changes
* Add `UnityCatalogImagePath` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `FeatureViewSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `BudgetPolicyId` and `Tags` fields for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec).
* Add `Development` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `UcServiceCredentialName` field for [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig).
* Add `ExcludedScopes` field for [settings.CustomerFacingIngressNetworkPolicyWorkspaceApiDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyWorkspaceApiDestination).
* Add `EffectiveWorkspaceLabel` and `WorkspaceLabel` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Add `PriorityClass` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `Gpu8xB300` enum value for [jobs.ComputeSpecAcceleratorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeSpecAcceleratorType).


## Release v0.178.0 (2026-09-09)

### Internal Changes

- Test changes.

### API Changes
* Add [aifunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions), [domains](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/domains) and [sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox) packages.
* Add [w.AiFunctions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aifunctions#AiFunctionsAPI) workspace-level service.
* Add [w.Domains](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/domains#DomainsAPI) workspace-level service.
* Add [w.Sandbox](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sandbox#SandboxAPI) workspace-level service.
* Add `BackfillFeatures`, `CancelOperation`, `GetOperation` and `PurgeFeatureEntities` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `ProvisionedCapacityId` field for [jobs.ComputeSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeSpec).
* Add `Mode` field for [ml.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CronSchedule).
* Add `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* Add `FilterCondition`, `Inputs` and `TimeWindow` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `ExtraParameters` and `FunctionType` fields for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `BudgetPolicyId` and `Tags` fields for [ml.IngestionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#IngestionConfig).
* Add `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource).
* Add `BudgetPolicyId`, `CronSchedule`, `LatestBackfillOperation` and `Tags` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `FullFeatureName` field for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec).
* Add `ExcludedColumns` and `RecordTypeFilter` fields for [ml.Stream](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Stream).
* Add `Continuous` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `DeploymentResourceTypeSecret` and `DeploymentResourceTypeClusterPolicy` enum values for [bundledeployments.DeploymentResourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#DeploymentResourceType).
* Add `AwsSecretsManager` and `AzureKeyVault` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `ExternalUseLocation` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Add `Gpu8xB300` enum value for [compute.HardwareAcceleratorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#HardwareAcceleratorType).
* [Breaking] Remove `Disabled` field for [catalog.InferenceTableConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#InferenceTableConfig).
* [Breaking] Remove `Owner` field for [catalog.McpService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpService).
* [Breaking] Remove `Owner` field for [catalog.ModelProviderService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderService).
* [Breaking] Remove `PlanType` field for [catalog.ModelProviderServiceConfigAnthropicProviderRelayedConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAnthropicProviderRelayedConfig).
* [Breaking] Remove `Owner` field for [catalog.ModelService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelService).
* [Breaking] Remove `FirstTokenTimeout` and `TrafficSplitting` fields for [catalog.ModelServiceConfigRoutingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelServiceConfigRoutingConfig).
* [Breaking] Remove `RequestTagKey` and `RequestTagValue` fields for [catalog.RateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RateLimit).
* [Breaking] Remove `RateLimitKeyRequestTag` enum value for [catalog.RateLimitRateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RateLimitRateLimitKey).
* Add `Deny` field for [catalog.PolicyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PolicyInfo).
* Add `PolicyTypeDeny` enum value for [catalog.PolicyType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PolicyType).


## Release v0.177.0 (2026-08-31)

### API Changes
* Add `UpdateDeployment` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `GenieCancelResponse` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `GetExternalGroup`, `GetExternalServicePrincipal` and `GetExternalUser` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `GetExternalGroupProxy`, `GetExternalServicePrincipalProxy` and `GetExternalUserProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `CreateSnapshot`, `DeleteSnapshot`, `GetSnapshot`, `GetSnapshotSchedule`, `ListSnapshots` and `UpdateSnapshotSchedule` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `AssumeGroupId` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space).
* Add `Operations` field for [bundledeployments.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Version).
* Add `FunctionArgExpression` field for [catalog.FunctionArgument](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionArgument).
* Add `AwsContextId` field for [compute.NodeTypeFlexibility](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeTypeFlexibility).
* Add `Parameters` field for [jobs.AlertTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AlertTask).
* Add `MaintenanceWindow` field for [jobs.Continuous](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Continuous).
* Add `MaintenanceWindow` field for [jobs.ContinuousTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ContinuousTriggerConfiguration).
* Add `Lateness` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `Delay` and `Offset` fields for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow).
* Add `StartTime` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Delay` and `Offset` fields for [ml.TumblingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TumblingWindow).
* Add `RabbitmqOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `SourceSnapshot` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `SourceSnapshot` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `DeploymentResourceTypeVectorSearchEndpoint`, `DeploymentResourceTypeVectorSearchIndex`, `DeploymentResourceTypeJobRun`, `DeploymentResourceTypePostgresCatalog`, `DeploymentResourceTypePostgresSyncedTable`, `DeploymentResourceTypeGenieSpace`, `DeploymentResourceTypeInstancePool`, `DeploymentResourceTypePostgresDatabase` and `DeploymentResourceTypePostgresRole` enum values for [bundledeployments.DeploymentResourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#DeploymentResourceType).
* Add `OperationStatusPending` enum value for [bundledeployments.OperationStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#OperationStatus).
* Add `Gpu1xH100` enum value for [compute.HardwareAcceleratorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#HardwareAcceleratorType).
* Add `Rabbitmq` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Change `SourceSchema` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to no longer be required.
* Change `SourceSchema` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to no longer be required.
* [Breaking] Change `SourceTable` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to no longer be required.
* Change `SourceTable` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to no longer be required.
* [Breaking] Remove `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* [Breaking] Remove `FilterCondition`, `Inputs` and `TimeWindow` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Remove `ExtraParameters` and `FunctionType` fields for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* [Breaking] Remove `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource).
* [Breaking] Remove `CronSchedule` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* [Breaking] Remove `Continuous` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Model`, `ModelService`, `McpService` and `ModelProviderService` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).


## Release v0.176.0 (2026-08-27)

### New Features and Improvements

- Added group role assumption to built-in SDK OAuth M2M and workload identity
  federation authentication, with safeguards that prevent other built-in
  credential strategies, including Databricks CLI authentication, from
  silently obtaining normal-access credentials.

### Internal Changes

- Added group ID configuration and U2M OAuth support for group role assumption.
  Account authentication is rejected, while unified authentication passes the
  group ID to the OAuth server for compatibility with future support. The SDK's
  Databricks CLI authentication rejects group role assumption until the CLI
  behavior is supported.


## Release v0.175.0 (2026-08-19)

### API Changes
* Add `EnabledTelemetryFeatures` field for [serving.TelemetryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#TelemetryConfig).


## Release v0.174.0 (2026-08-18)

### Bug Fixes

- Recover from concurrent OAuth token cache writes by using the fresh token stored by another process.
- Fix Spark runtime selection for major-only runtime versions.
- Accept JSON numbers and decimal strings when unmarshalling `int64` API response fields.

### API Changes

* Add `AgentType` field for [dashboards.GenieConversationSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversationSummary).
* Add `GroupId` field for [iamv2.DirectGroupMember](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DirectGroupMember).
* Add `EffectiveServerlessComputeId` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `PipelineChannel` field for [postgres.NewPipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#NewPipelineSpec).
* Add `TextAttachmentPurposeAnswer` enum value for [dashboards.TextAttachmentPurpose](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachmentPurpose).
* [Breaking] Add pagination for [AccountIamV2API.ListDirectGroupMembers](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListDirectGroupMembers).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignmentDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignmentDetails).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListDirectGroupMembersProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListDirectGroupMembersProxy).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [database.SyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpecPgSpecificType).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [postgres.SyncedTableSyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpecPgSpecificType).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignments).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentsProxy).
* Add `DefaultGitSource`, `GitSource` and `SourceCodePath` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `AutoDeploy` and `CallerCredentialId` fields for [apps.GitRepository](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#GitRepository).
* Add `CustomUdf` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `EffectiveServerlessComputeId` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).


## Release v0.173.0 (2026-08-17)

### Breaking Changes

 * Raise the minimum supported Go version from 1.24 to 1.25

### Bug Fixes

- Recover from concurrent OAuth token cache writes by using the fresh token stored by another process.
- Fix Spark runtime selection for major-only runtime versions.

### API Changes
* Add `AgentType` field for [dashboards.GenieConversationSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversationSummary).
* Add `GroupId` field for [iamv2.DirectGroupMember](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#DirectGroupMember).
* Add `EffectiveServerlessComputeId` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `PipelineChannel` field for [postgres.NewPipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#NewPipelineSpec).
* Add `TextAttachmentPurposeAnswer` enum value for [dashboards.TextAttachmentPurpose](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachmentPurpose).
* [Breaking] Add pagination for [AccountIamV2API.ListDirectGroupMembers](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListDirectGroupMembers).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignmentDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignmentDetails).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListDirectGroupMembersProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListDirectGroupMembersProxy).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentDetailsProxy).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [database.SyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpecPgSpecificType).
* Add `PgSpecificTypeHalfvec` and `PgSpecificTypeVarchar` enum values for [postgres.SyncedTableSyncedTableSpecPgSpecificType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpecPgSpecificType).
* [Breaking] Add pagination for [AccountIamV2API.ListWorkspaceAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API.ListWorkspaceAssignments).
* [Breaking] Add pagination for [WorkspaceIamV2API.ListWorkspaceAssignmentsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API.ListWorkspaceAssignmentsProxy).
* Add `LinkedinAdsOptions` and `MarketoOptions` fields for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* [Breaking] Change `ActionType` and `Status` fields for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* Change `ActionType` and `Status` fields for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* [Breaking] Remove `CreateOperation` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.


## Release v0.172.0 (2026-08-13)

### API Changes
* Add `CreateDirectGroupMember`, `CreateGroup`, `CreateServicePrincipal`, `CreateUser`, `CreateWorkspaceAssignment`, `DeleteDirectGroupMember`, `DeleteGroup`, `DeleteServicePrincipal`, `DeleteUser`, `DeleteWorkspaceAssignment`, `GetDirectGroupMember`, `GetGroup`, `GetServicePrincipal`, `GetUser`, `GetWorkspaceAssignment`, `ListDirectGroupMembers`, `ListGroups`, `ListServicePrincipals`, `ListTransitiveParentGroups`, `ListUsers`, `ListWorkspaceAssignments`, `UpdateGroup`, `UpdateServicePrincipal`, `UpdateUser` and `UpdateWorkspaceAssignment` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `CreateDirectGroupMemberProxy`, `CreateGroupProxy`, `CreateServicePrincipalProxy`, `CreateUserProxy`, `CreateWorkspaceAssignmentProxy`, `DeleteDirectGroupMemberProxy`, `DeleteGroupProxy`, `DeleteServicePrincipalProxy`, `DeleteUserProxy`, `DeleteWorkspaceAssignmentProxy`, `GetDirectGroupMemberProxy`, `GetGroupProxy`, `GetServicePrincipalProxy`, `GetUserProxy`, `GetWorkspaceAssignmentProxy`, `GetWorkspaceIdentityDetail`, `ListDirectGroupMembersProxy`, `ListGroupsProxy`, `ListServicePrincipalsProxy`, `ListTransitiveParentGroupsProxy`, `ListUsersProxy`, `ListWorkspaceAssignmentsProxy`, `UpdateGroupProxy`, `UpdateServicePrincipalProxy`, `UpdateUserProxy`, `UpdateWorkspaceAssignmentProxy` and `UpdateWorkspaceIdentityDetail` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `CreateTime` and `UpdateTime` fields for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Change `NewCluster` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster) to no longer be required.
* [Breaking] Change `NewCluster` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster) to no longer be required.


## Release v0.171.0 (2026-08-12)

### API Changes
* Add `ForwardUserAccessToken` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ForwardUserAccessToken` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `IncludeValue` field for [catalog.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSecretRequest).
* Add `DockerImageUrl` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `InputColumn` and `OutputColumn` fields for [pipelines.Transformer](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Transformer).
* Add `Netsuite` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `GpuXlarge8` enum value for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `GpuXlarge8` enum value for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
* Add `LegacyMode` enum value for [settings.CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyCrossWorkspaceAccessRestrictionMode).
* Change `AccountSpStatus` and `DisplayName` fields for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal) to be required.
* Change `AccountUserStatus`, `FullName` and `Username` fields for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User) to be required.
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetMcpServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMcpServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelProviderServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelProviderServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetModelServiceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelServiceRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListMcpServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMcpServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelProviderServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelProviderServicesRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListModelServicesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelServicesRequest).
* [Breaking] Remove `BrowseOnly` field for [catalog.McpService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#McpService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelProviderService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderService).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelService](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelService).
* Add `MlflowArtifactLocation` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `TriggerDetails` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
* Add `Triggers` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Add `TriggerDetails` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
* Add `Triggers` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Add `ApiSourceConnectorOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `ApiSourceConnectorConfig` field for [pipelines.SourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SourceConfig).
* Add `AllVpcScServices` and `GoogleApiEndpoints` fields for [settings.GcpEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GcpEndpoint).


## Release v0.170.0 (2026-08-05)

### API Changes
* Add `PrincipalOverrides` and `ScopeType` fields for [billing.AlertConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfiguration).
* Add `ResourceType` field for [billing.BudgetConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfiguration).
* Add `ResourceType` field for [billing.CreateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudget).
* Add `PrincipalOverrides` and `ScopeType` fields for [billing.CreateBudgetConfigurationBudgetAlertConfigurations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudgetAlertConfigurations).
* Add `ResourceType` field for [billing.UpdateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationBudget).
* Add `Grant` field for [catalog.PolicyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PolicyInfo).
* Add `BlockUsage` enum value for [billing.ActionConfigurationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ActionConfigurationType).
* Add `PolicyTypeGrant` enum value for [catalog.PolicyType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PolicyType).


## Release v0.169.0 (2026-08-04)

### New Features and Improvements

* Add the Pi coding agent (`PI_CODING_AGENT`) to AI agent detection in the User-Agent header, so CLI usage driven by Pi is attributed to `pi`.

### API Changes
* Add `KinesisStreamConfig` field for [ml.StreamSourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamSourceConfig).
* Add `Mode` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* [Breaking] Change `State` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to type `string`.
* [Breaking] Change `State` field for [bundledeployments.Resource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Resource) to type `string`.
* [Breaking] Remove `AwsAccessKeyId` and `AwsSecretAccessKey` fields for [catalog.ModelProviderServiceConfigAmazonBedrockProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAmazonBedrockProviderDirectConfig).
* [Breaking] Remove `ClientId`, `ClientSecret` and `TenantId` fields for [catalog.ModelProviderServiceConfigAzureOpenAiProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAzureOpenAiProviderDirectConfig).
* [Breaking] Remove `ClientId`, `ClientSecret` and `TenantId` fields for [catalog.ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig).


## Release v0.168.0 (2026-08-03)

### API Changes
* Add `PauseStatus` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Add `DirectoryInfo` field for [workspace.ObjectInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ObjectInfo).
* Add `GroupId` field for [iamv2.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#Group).
* Add `ServicePrincipalId` field for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal).
* Add `UserId` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).
* [Breaking] Remove `InternalId` field for [iamv2.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#Group).
* [Breaking] Remove `InternalId` field for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal).
* [Breaking] Remove `InternalId` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).


## Release v0.167.0 (2026-07-30)

### API Changes
* Add `AwsAccessKey` field for [catalog.ModelProviderServiceConfigAmazonBedrockProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAmazonBedrockProviderDirectConfig).
* Add `EntraServicePrincipal` field for [catalog.ModelProviderServiceConfigAzureOpenAiProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigAzureOpenAiProviderDirectConfig).
* Add `EntraServicePrincipal` field for [catalog.ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig).


## Release v0.166.0 (2026-07-29)

### API Changes
* Add [w.AiGateway](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AiGatewayAPI) workspace-level service.
* Add `UpdateOperation` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `LastSuccessfulVersionId` and `UpdatedBy` fields for [bundledeployments.Deployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Deployment).
* Add `DashboardMetadata`, `SequenceId` and `UpdateTime` fields for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation).
* Add `DashboardMetadata` field for [bundledeployments.Resource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Resource).
* Add `PreviousVersionId` field for [bundledeployments.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Version).


## Release v0.165.0 (2026-07-28)

### API Changes
* Add `PatchTelemetryConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Add `FanoutOptions` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec).
* Add `Parameters` field for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2).


## Release v0.164.0 (2026-07-27)

### API Changes
* Add `SpecialDestinationSchemaOwner`, `SpecialDestinationTableOwner`, `SpecialDestinationVolumeOwner`, `SpecialDestinationFunctionOwner` and `SpecialDestinationRegisteredModelOwner` enum values for [catalog.SpecialDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SpecialDestination).
* [Breaking] Change `CreateDeployment` method for [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service with new required argument order.
* [Breaking] Remove `DeploymentId` field for [bundledeployments.CreateDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#CreateDeploymentRequest).


## Release v0.163.0 (2026-07-23)

### API Changes
* Add `DependencyMode` field for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
* Add `DependencyMode` field for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
* Add `DependencyMode` field for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
* Add `DependencyMode` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
* Add `DependencyMode` field for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
* Add `DependencyMode` field for [compute.EnforcePolicyComplianceForClusterResponseClusterSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforcePolicyComplianceForClusterResponseClusterSettings).
* Add `DependencyMode` field for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
* Add `CodeSourcePath` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).
* Add `SchemaRegistryConfig` field for [ml.StreamSchemaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamSchemaConfig).
* Add `ExtraColumns` field for [postgres.SyncedTableSyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpec).
* [Breaking] Remove `Lifetime` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).


## Release v0.162.0 (2026-07-22)

### API Changes
* Add `EffectiveEntitlements` field for [iamv2.WorkspaceAssignmentDetail](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceAssignmentDetail).
* Add `ServerlessComputeId` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster).


## Release v0.161.0 (2026-07-21)

### API Changes
* Add `PerformanceTarget` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Add `DataframeSchema` and `TransformationSql` fields for [ml.StreamSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamSource).
* Add `Sawtooth` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `EffectiveTablePrefix` field for [ml.UcTraceLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UcTraceLocation).
* Add `AwsVpcEndpointInfo` and `GcpPscEndpointInfo` fields for [networking.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking#Endpoint).
* Add `RedditAdsOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `TableNames` and `TelemetryProfileId` fields for [serving.TelemetryConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#TelemetryConfig).
* Add `GitCliEnabled` field for [workspace.GetRepoResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoResponse).
* Add `KsaEccCccDcc` enum value for [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
* [Breaking] Change `CreateWorkspaceAssignmentDetail`, `DeleteWorkspaceAssignmentDetail`, `GetWorkspaceAccessDetail`, `GetWorkspaceAssignmentDetail`, `ListWorkspaceAssignmentDetails`, `ResolveGroup`, `ResolveServicePrincipal`, `ResolveUser` and `UpdateWorkspaceAssignmentDetail` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service. Method path has changed.
* [Breaking] Change `CreateWorkspaceAssignmentDetailProxy`, `DeleteWorkspaceAssignmentDetailProxy`, `GetWorkspaceAccessDetailLocal`, `GetWorkspaceAssignmentDetailProxy`, `ListWorkspaceAssignmentDetailsProxy`, `ResolveGroupProxy`, `ResolveServicePrincipalProxy`, `ResolveUserProxy` and `UpdateWorkspaceAssignmentDetailProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service. Method path has changed.
* [Breaking] Change `WindowDuration` field for [ml.RollingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RollingWindow) to no longer be required.
* Change `WindowDuration` field for [ml.RollingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RollingWindow) to no longer be required.
* Change `WindowDuration` field for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow) to no longer be required.
* [Breaking] Change `WindowDuration` field for [ml.SlidingWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SlidingWindow) to no longer be required.
* [Breaking] Remove `LongRolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).


## Release v0.160.0 (2026-07-16)

### API Changes
* Add `ListCleanRoomTaskRunsHandler` method for [w.CleanRoomTaskRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomTaskRunsAPI) workspace-level service.
* Add `InitialParentPath` field for [bundledeployments.Deployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Deployment).
* Add `JarAnalysis` field for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset).
* Add `JarAnalysisReview` field for [cleanrooms.CreateCleanRoomAssetReviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewRequest).
* Add `JarAnalysisReviewState` and `JarAnalysisReviews` fields for [cleanrooms.CreateCleanRoomAssetReviewResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewResponse).
* Add `CapacityReservationGroup` field for [compute.InstancePoolAzureAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAzureAttributes).
* Add `Lifetime` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `Jdbc` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `JarAnalysis` enum value for [cleanrooms.CleanRoomAssetAssetType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetAssetType).


## Release v0.159.0 (2026-07-15)

### API Changes
* Add `EffectiveWorkspaceId` field for [disasterrecovery.StableUrl](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#StableUrl).
* Add `SourceMetadataColumn` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* [Breaking] Remove `CodeSourcePath` field for [jobs.AiRuntimeTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTask).


## Release v0.158.0 (2026-07-14)

### API Changes
* Add `EnableSharedOutput` field for [cleanrooms.CleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoom).
* Add `Description` and `EnvironmentVersion` fields for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook).
* Add `SharedOutputSchemaExpirationTime` and `SharedOutputSchemaName` fields for [cleanrooms.CleanRoomNotebookTaskRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomNotebookTaskRun).
* Add `EnableSharedOutput` and `PackageProviderCollaboratorAlias` fields for [cleanrooms.CleanRoomRemoteDetail](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomRemoteDetail).
* Add `CustomReportOptions` field for [pipelines.GoogleAdsOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GoogleAdsOptions).
* Add `CustomReportOptions` field for [pipelines.MetaMarketingOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#MetaMarketingOptions).
* Add `CustomReportOptions` field for [pipelines.TikTokAdsOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TikTokAdsOptions).


## Release v0.157.0 (2026-07-13)

### API Changes
* Add `CreateCdfConfig`, `DeleteCdfConfig`, `GetCdfConfig`, `GetCdfStatus`, `ListCdfConfigs` and `ListCdfStatuses` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Parent` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Add `Parent` field for [catalog.ListConnectionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListConnectionsRequest).
* Add `Minutes` enum value for [jobs.PeriodicTriggerConfigurationTimeUnit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PeriodicTriggerConfigurationTimeUnit).
* [Breaking] Change `InternalId` field for [iamv2.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#Group) to type `string`.
* [Breaking] Change `InternalId` field for [iamv2.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#ServicePrincipal) to type `string`.
* [Breaking] Change `InternalId` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User) to type `string`.


## Release v0.156.0 (2026-07-09)

### API Changes
* Add `StableWorkspaceId` field for [disasterrecovery.StableUrl](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#StableUrl).
* Add `DeploymentId` and `VersionId` fields for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Add `DeploymentId` and `VersionId` fields for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Add `TraceLocation` field for [ml.CreateExperiment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateExperiment).
* Add `TraceLocation` field for [ml.Experiment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Experiment).


## Release v0.155.0 (2026-07-08)

### API Changes
* Add `UpdateTime` field for [bundledeployments.Resource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Resource).
* [Breaking] Change `Role` field for [postgres.DatabaseDatabaseSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DatabaseDatabaseSpec) to be required.
* Change `Role` field for [postgres.DatabaseDatabaseSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DatabaseDatabaseSpec) to be required.


## Release v0.154.0 (2026-07-07)

### API Changes
* Add `ParentPath` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Add `ParentPath` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Add `GitCredentialId` field for [workspace.CreateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateRepoRequest).
* Add `GitCredentialId` field for [workspace.UpdateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateRepoRequest).
* Add `CpuLarge` and `CpuMedium` enum values for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `CpuLarge` and `CpuMedium` enum values for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).


## Release v0.153.0 (2026-07-06)

### API Changes
* Add `List` and `ListEffective` methods for [w.Grants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GrantsAPI) workspace-level service.
* Add `OmitPermissionsInResponse` field for [catalog.UpdatePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissions).
* Add `ExcludedSchemas` field for [dataclassification.CatalogConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#CatalogConfig).
* Add `FullName` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).
* Add `StatusMessage` field for [jobs.AiRuntimeTaskOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#AiRuntimeTaskOutput).
* Add `SqlCondition` field for [jobs.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CronSchedule).
* Add `SqlCondition` field for [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
* Add `SqlCondition` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* Add `SqlCondition` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Add `FirstDistinct`, `FirstN`, `LastDistinct` and `LastN` fields for [ml.AggregationFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#AggregationFunction).
* Add `AvroSchema` and `ProtoSchema` fields for [ml.SchemaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SchemaConfig).
* Add `FreshnessTarget` field for [ml.StreamingMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#StreamingMode).
* Add `LongRolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `ReadOnlyPooledHost` and `ReadWritePooledHost` fields for [postgres.EndpointHosts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointHosts).
* Add `LastActiveTime` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `ExpireTime` and `Ttl` fields for [postgres.GenerateDatabaseCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#GenerateDatabaseCredentialRequest).
* Add `ReadMetadata` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* [Breaking] Remove `IncludeBrowse` field for [catalog.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSecretRequest).
* [Breaking] Remove `IncludeBrowse` field for [catalog.ListSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecretsRequest).
* [Breaking] Remove `BrowseOnly` and `ExternalSecretId` fields for [catalog.Secret](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Secret).
* [Breaking] Remove `Name` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User).


## Release v0.152.0 (2026-06-29)

### Breaking Changes

* Query parameters listed in `ForceSendFields` are now sent on the wire when they hold a zero value. Previously `ForceSendFields` had no effect on query parameters (only JSON body fields were honored), so such a parameter was silently omitted; it is now serialized with its explicit value (for example `cascade=false`). Callers that unknowingly relied on the previous no-op behavior will now send the parameter.

### Bug Fixes

* Honor `ForceSendFields` for query parameters, including fields nested inside request sub-structs. A zero-valued scalar field (for example a `false` bool such as `cascade` on the pipelines Delete request) listed in `ForceSendFields` is now sent on the wire instead of being dropped by the `url` tag's `omitempty`, matching the existing behavior for JSON request bodies.

### API Changes
* Add `Etag` field for [disasterrecovery.UpdateFailoverGroupRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#UpdateFailoverGroupRequest).
* Add `DownloadMessageAttachmentVisualization` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `Viz` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Add `EnableVisualization` field for [dashboards.GenieCreateConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieCreateConversationMessageRequest).
* Add `EnableVisualization` field for [dashboards.GenieStartConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieStartConversationMessageRequest).
* Add `ReplaceExisting` field for [postgres.CreateDatabaseRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateDatabaseRequest).
* Add `ReplaceExisting` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest).
* Add `AutoscalingLimitMaxCu`, `AutoscalingLimitMinCu`, `NoSuspension` and `SuspendTimeoutDuration` fields for [postgres.InitialEndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#InitialEndpointSpec).
* Add `InitialBranchSpec` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `ComputeLastActiveTime` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `TelemetryConfig` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Add `TelemetryConfig` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Add `TelemetryConfig` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).


## Release v0.151.0 (2026-06-25)

### API Changes
* Add `Spec` field for [environments.WorkspaceBaseEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments#WorkspaceBaseEnvironment).


## Release v0.150.0 (2026-06-24)

### Bug Fixes

* Deduplicate identical key/value pairs in the user agent builder ([useragent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/useragent)), so repeatedly injecting the same dimension onto a reused `context.Context` no longer grows the `User-Agent` header without bound. Distinct values for the same key are still preserved.

### API Changes
* Add `CancelPendingClusterEnforcement` method for [w.PolicyComplianceForClusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#PolicyComplianceForClustersAPI) workspace-level service.
* Add `PendingEnforcement` field for [compute.ClusterCompliance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterCompliance).
* Add `EnforceMode` field for [compute.EnforceClusterComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceRequest).
* Add `EnforceResult` field for [compute.EnforceClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceResponse).
* Add `PendingEnforcement` field for [compute.GetClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterComplianceResponse).
* Add `DeferredPolicyEnforcementScheduled` and `DeferredPolicyEnforcementFailed` enum values for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Change `ReplicateWorkspaceAssets` field for [disasterrecovery.WorkspaceSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#WorkspaceSet) to no longer be required.
* [Breaking] Change `ReplicateWorkspaceAssets` field for [disasterrecovery.WorkspaceSet](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#WorkspaceSet) to no longer be required.


## Release v0.149.0 (2026-06-23)

### API Changes
* Add `BundleRootPath` field for [bundledeployments.WorkspaceInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#WorkspaceInfo).
* Add `AiRuntimeTask` field for [jobs.ResolvedValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedValues).
* Add `AiRuntimeTaskOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Add `AiRuntimeTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `AiRuntimeTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `AiRuntimeTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).


## Release v0.148.0 (2026-06-22)

### New Features and Improvements

* Added a `meta-harness` user-agent dimension that reports the omnigent meta-harness (detected via the `OMNIGENT` environment variable) independently of agent detection.

### API Changes
* Add `Xlarge` enum value for [apps.ComputeSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeSize).


## Release v0.147.0 (2026-06-17)

### API Changes
* Add `ServerlessComputeId` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* Add `ServerlessComputeId` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `ServerlessComputeId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `ServerlessComputeId` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).


## Release v0.146.0 (2026-06-16)

### API Changes
* Add `CreateDataApi`, `DeleteDataApi`, `GetDataApi` and `UpdateDataApi` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `CapacityReservationGroup` field for [compute.AzureAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#AzureAttributes).


## Release v0.145.0 (2026-06-15)

### Bug Fixes

* Encode the governed-tag key as a single path segment when it is sent as a URL path parameter, so hierarchical keys containing `/` route correctly instead of being split into extra path segments and resolving to no endpoint (`404` / `ENDPOINT_NOT_FOUND`). Affects `GetTagPolicy`/`DeleteTagPolicy`/`UpdateTagPolicy` ([tags.TagPoliciesAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPoliciesAPI)), `GetTagAssignment`/`DeleteTagAssignment`/`UpdateTagAssignment` ([tags.WorkspaceEntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#WorkspaceEntityTagAssignmentsAPI)), and `Get`/`Delete`/`Update` ([catalog.EntityTagAssignmentsAPI](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignmentsAPI)).

### API Changes
* Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* [Breaking] Change `ResourceId` field for [bundledeployments.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#Operation) to no longer be required.
* Add `Dynamics365` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `EndpointId` field for [vectorsearch.MiniVectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#MiniVectorIndex).
* Add `EndpointId` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).


## Release v0.144.0 (2026-06-11)

### API Changes
* Add `TypeOverrides` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* Add `TypeOverrides` field for [postgres.SyncedTableSyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpec).


## Release v0.143.0 (2026-06-10)

### API Changes
* Add `AcceleratedSync` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* Add `AcceleratedSync` field for [postgres.SyncedTableSyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpec).


## Release v0.142.0 (2026-06-09)

### API Changes
* Add [aisearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch) and [bundledeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments) packages.
* Add [w.AiSearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aisearch#AiSearchAPI) workspace-level service.
* Add [w.BundleDeployments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundledeployments#BundleDeploymentsAPI) workspace-level service.
* Add `RunningInstances` field for [apps.ApplicationStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ApplicationStatus).
* Add `CustomMaxRetentionHours` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
* Add `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* Add `CustomMaxRetentionHours` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
* Add `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Add `CustomMaxRetentionHours` field for [catalog.CreateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateSchema).
* Add `CustomMaxRetentionHours` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
* Add `CustomMaxRetentionHours` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
* Add `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* Add `CustomMaxRetentionHours` field for [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
* Add `StreamSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `IngestionConfig` field for [ml.KafkaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaConfig).
* Add `ClusteringColumns`, `EnableAutoClustering` and `TableProperties` fields for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `BranchId` field for [postgres.Branch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Branch).
* Add `CatalogId` field for [postgres.Catalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Catalog).
* Add `DatabaseId` field for [postgres.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Database).
* Add `EndpointId` field for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
* Add `ProjectId` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `RoleId` field for [postgres.Role](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Role).
* Add `SyncedTableId` field for [postgres.SyncedTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTable).
* Add `AllowedDatabricksDestinations` field for [settings.EgressNetworkPolicyNetworkAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyNetworkAccessPolicy).
* Add `Facets`, `QueryColumns` and `SortColumns` fields for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* Add `FacetResult` field for [vectorsearch.QueryVectorIndexResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexResponse).
* Add `FacetColumnCount` and `FacetColumns` fields for [vectorsearch.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ResultManifest).
* Add `DangerouslyForceDiscardAll` field for [workspace.UpdateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateRepoRequest).
* [Breaking] Remove [bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle) package.
* [Breaking] Remove [w.Bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#BundleAPI) workspace-level service.


## Release v0.141.0 (2026-06-02)

### API Changes
* Add `DeploymentMode` field for [bundle.Deployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Deployment).
* Add `DeploymentMode` field for [bundle.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Version).
* Add `CollaborationPlatformConnectivity` and `EffectiveCollaborationPlatformConnectivity` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).


## Release v0.140.0 (2026-06-01)

### New Features and Improvements

* Add `u2m.WithDiscoveryAccountTarget()` option that sets `target=ACCOUNT` on the login.databricks.com authorize URL, so the discovery flow lands the user on the account selector instead of the workspace selector.

### Internal Changes

* Switch workspace addressing header on workspace-scoped API calls from `X-Databricks-Org-Id` to `X-Databricks-Workspace-Id`. The value continues to come from `Config.WorkspaceID` (`DATABRICKS_WORKSPACE_ID`), and now accepts either a classic numeric workspace ID or a CPDR connection ID (server disambiguates). Part of the DECO + Unified Workspace Addressing initiative.

### API Changes
* Add `UpdateTokenManagement` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
* Add `DeploymentId` and `VersionId` fields for [jobs.JobDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeployment).
* Add `DeploymentId` and `VersionId` fields for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment).
* Add `AutoscopeEnabled` field for [settings.CreateOboTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateOboTokenRequest).
* Add `AutoscopeEnabled` field for [settings.CreateTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateTokenRequest).
* Add `AutoscopeState`, `BackfillScopes`, `InferredScopes` and `Scopes` fields for [settings.PublicTokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PublicTokenInfo).
* Add `AutoscopeState`, `BackfillScopes`, `InferredScopes` and `Scopes` fields for [settings.TokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenInfo).
* Add `ResourceType` field for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation).


## Release v0.139.0 (2026-05-28)

### API Changes
* Add `CreateStream`, `DeleteStream`, `GetStream`, `ListStreams` and `UpdateStream` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `Parameters` field for [jobs.PipelineTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PipelineTask).
* Add `PipelineTask` field for [jobs.ResolvedValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedValues).
* Add `Parameters` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `Parameters` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `Parameters` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* [Breaking] Remove `CatalogId` field for [postgres.CatalogCatalogStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CatalogCatalogStatus).
* [Breaking] Remove `SyncedTableId` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).


## Release v0.138.0 (2026-05-27)

### Bug Fixes

* `config`: The config-file loader now sets `cfg.Profile` to `"DEFAULT"` on the legacy fallback (no profile requested, no `[__settings__].default_profile` set, `[DEFAULT]` section loaded). Previously the loader silently used `[DEFAULT]`'s values but left `cfg.Profile` empty. Consumers that derive a per-profile identifier from `cfg.Profile` (e.g. the Databricks CLI's U2M OAuth cache key) need the resolved name to match across the login and read flows; without this fix a write under `"DEFAULT"` could not be found by a no-flag read.


## Release v0.137.0 (2026-05-25)

### New Features and Improvements

* Honor the Vercel `AI_AGENT=<name>` env var as a secondary fallback for
  AI agent detection in the User-Agent header (after the agents.md
  `AGENT=<name>` standard). Unrecognized fallback values now pass through
  the User-Agent sanitized and length-capped at 64 chars instead of being
  coerced to `agent/unknown`, so versioned variants such as
  `claude-code_2-1-141_agent` surface as-is.

### Internal Changes

* Pass `excludedAttributes=entitlements` on the SCIM `/Me` request made by `WorkspaceClient.CurrentWorkspaceID` ([#1681](https://github.com/databricks/databricks-sdk-go/pull/1681)).

  `CurrentWorkspaceID` only reads the `X-Databricks-Org-Id` response header and discards the body, so it has no use for the `User.Entitlements` field. Skipping that attribute avoids an expensive `getEffectivePermissions` scan on the SCIM backend, which has caused incidents on workspaces with large grant counts.

### API Changes
* Add `Revert` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
* Add `ParentPath` field for [dashboards.GenieUpdateSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieUpdateSpaceRequest).
* Add `ComputeMaxInstances` and `ComputeMinInstances` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ComputeMaxInstances` and `ComputeMinInstances` fields for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `CronScheduleTrigger`, `StreamingMode` and `TableTrigger` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `SyncedTableId` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).


## Release v0.136.0 (2026-05-21)

### API Changes
* Add `UndeleteBranch` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Attributes` and `ExcludedAttributes` fields for [iam.MeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MeRequest).
* Add `IncludeTriggerState` field for [jobs.GetJobRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobRequest).
* Add `DeleteTime` and `PurgeTime` fields for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `Purge` field for [postgres.DeleteBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DeleteBranchRequest).
* Add `ShowDeleted` field for [postgres.ListBranchesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ListBranchesRequest).
* Add `Deleted` enum value for [postgres.BranchStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatusState).
* [Breaking] Change `ActionType` and `ResourceId` fields for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation) to be required.
* Change `ActionType` and `ResourceId` fields for [bundle.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Operation) to be required.
* [Breaking] Change `CliVersion` field for [bundle.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Version) to be required.
* Change `CliVersion` field for [bundle.Version](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#Version) to be required.
* [Breaking] Change `Tags` field for [marketplace.ListListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListListingsRequest) to type [marketplace.ListingTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListingTag).
* [Breaking] Change pagination for [ClustersAPI.Events](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI.Events).


## Release v0.135.0 (2026-05-20)

### API Changes
* Add [bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle) package.
* Add [w.Bundle](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/bundle#BundleAPI) workspace-level service.
* Add `MtlsConfig` field for [ml.AuthConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#AuthConfig).


## Release v0.134.0 (2026-05-19)

### API Changes
* Add `FullRefreshSelection`, `RefreshFlowSelection`, `RefreshSelection` and `ResetCheckpointSelection` fields for [jobs.PipelineParams](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PipelineParams).
* Add `FullRefreshSelection`, `RefreshFlowSelection`, `RefreshSelection` and `ResetCheckpointSelection` fields for [jobs.PipelineTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PipelineTask).
* Add `EffectiveOperationalEmailCustomRecipient` and `OperationalEmailCustomRecipient` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).


## Release v0.133.0 (2026-05-18)

### API Changes
* Add `CreateWorkspaceAssignmentDetail`, `DeleteWorkspaceAssignmentDetail`, `GetWorkspaceAssignmentDetail`, `ListWorkspaceAssignmentDetails` and `UpdateWorkspaceAssignmentDetail` methods for [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service.
* Add `CreateWorkspaceAssignmentDetailProxy`, `DeleteWorkspaceAssignmentDetailProxy`, `GetWorkspaceAssignmentDetailProxy`, `ListWorkspaceAssignmentDetailsProxy` and `UpdateWorkspaceAssignmentDetailProxy` methods for [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Add `FailoverGroupName` field for [disasterrecovery.StableUrl](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#StableUrl).
* Add `Disabled` and `PythonOperatorTask` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `Disabled` and `PythonOperatorTask` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `PythonOperatorTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `CatalogName`, `CreatedAt`, `CreatedBy`, `Name` and `SchemaName` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Add `CatalogName` and `SchemaName` fields for [ml.ListFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListFeaturesRequest).
* Add `Rolling` field for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `KafkaOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `CrossWorkspaceAccess` field for [settings.CustomerFacingIngressNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicy).
* Add `AllowedAppsUserApiScopes` and `EffectiveAllowedAppsUserApiScopes` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Add `Hubspot`, `Github`, `Outlook` and `Smartsheet` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `GpuXlarge` enum value for [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
* Add `GpuXlarge` enum value for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
* [Breaking] Change `ListFeatures` method for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service with new required argument order.
* [Breaking] Remove `UnspecifiedResourceName` field for [postgres.RequestedResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RequestedResource).


## Release v0.132.0 (2026-05-07)

### API Changes
* Add `ConfluenceOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `Confluence` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Change `Description` field for [supervisoragents.SupervisorAgent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgent) to no longer be required.
* [Breaking] Change `Description` field for [supervisoragents.SupervisorAgent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgent) to no longer be required.
* Add `CreateExample`, `DeleteExample`, `GetExample`, `GetPermissionLevels`, `GetPermissions`, `ListExamples`, `SetPermissions`, `UpdateExample` and `UpdatePermissions` methods for [w.SupervisorAgents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgentsAPI) workspace-level service.
* Add `MetaAdsOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `MetaMarketing` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `MetaMarketing` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Change `Guidelines` field for [knowledgeassistants.Example](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#Example) to no longer be required.
* [Breaking] Change `Guidelines` field for [knowledgeassistants.Example](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#Example) to no longer be required.
* Change `Description` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool) to no longer be required.
* [Breaking] Change `Description` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool) to no longer be required.
* Add `Zendesk` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Add `R2TempCredentials` field for [catalog.TemporaryCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryCredentials).
* Add `ZendeskSupportOptions` field for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `AzureKeyInfo` field for [provisioning.CreateCustomerManagedKeyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateCustomerManagedKeyRequest).
* Add `TargetQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Add `ColumnsToIndex` field for [vectorsearch.DeltaSyncVectorIndexSpecRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecRequest).
* Add `ColumnsToIndex` field for [vectorsearch.DeltaSyncVectorIndexSpecResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecResponse).
* Add `RequestedTargetQps` field for [vectorsearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfo).
* Add `TargetQps` field for [vectorsearch.PatchEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#PatchEndpointRequest).
* Add `Jira` and `Zendesk` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Remove `MinQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* [Breaking] Remove `RequestedMinQps` field for [vectorsearch.EndpointScalingInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointScalingInfo).
* [Breaking] Remove `MinQps` field for [vectorsearch.PatchEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#PatchEndpointRequest).
* Add `GetPermissionLevels`, `GetPermissions`, `SetPermissions` and `UpdatePermissions` methods for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `PrivateAccess` field for [settings.CustomerFacingIngressNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicy).
* Add `AccountApi`, `AccountDatabricksOne` and `AccountUi` fields for [settings.CustomerFacingIngressNetworkPolicyRequestDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyRequestDestination).
* Add `ScopeQualifier` field for [settings.CustomerFacingIngressNetworkPolicyWorkspaceApiDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyWorkspaceApiDestination).
* Add `DisplayName` and `PreviewPhase` fields for [settingsv2.SettingsMetadata](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#SettingsMetadata).


## Release v0.131.0 (2026-04-28)

### API Changes
* Add `UndeleteProject` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `ConfidentialComputeType` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
* Add `Purge` field for [postgres.DeleteProjectRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DeleteProjectRequest).
* Add `ShowDeleted` field for [postgres.ListProjectsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ListProjectsRequest).
* Add `DeleteTime` and `PurgeTime` fields for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `UcConnection` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool).
* Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.
* [Breaking] Change `Name` field for [supervisoragents.Connection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Connection) to no longer be required.
* Add [disasterrecovery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery) package.
* Add [a.DisasterRecovery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/disasterrecovery#DisasterRecoveryAPI) account-level service.
* Add `CreateExample`, `DeleteExample`, `GetExample`, `ListExamples` and `UpdateExample` methods for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `DeltaTableName` field for [ml.BackfillSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#BackfillSource).
* Add `Confluence` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* [Breaking] Remove `Connection` field for [supervisoragents.Tool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#Tool).


## Release v0.130.0 (2026-04-23)

### Breaking Changes

* Remove the `Experimental_IsUnifiedHost` field (and the `DATABRICKS_EXPERIMENTAL_IS_UNIFIED_HOST` environment variable) from `Config`. Unified host detection is now automatic via the `/.well-known/databricks-config` endpoint.
* Remove the unused `ErrWorkspaceIDInAccountClient` exported variable. It was never returned from any production path, and its message contradicted the unified host workflow where a single profile with both `AccountID` and `WorkspaceID` produces both clients.
* Remove the file-based OAuth token cache from `credentials/u2m/cache`. The removed symbols are `cache.NewFileTokenCache`, `cache.FileTokenCacheOption`, `cache.WithFileLocation`, and the private `tokenCacheFile` struct. The `TokenCache` interface, `ErrNotFound` sentinel, `HostCacheKeyProvider`, and `DiscoveryOAuthArgument` remain exported. `NewPersistentAuth` now defaults to a new in-memory cache (`cache.NewInMemoryTokenCache`) when no `WithTokenCache` option is passed; consumers that relied on the previous file-backed default must supply their own persistent cache. See databricks/cli#5056 for the companion CLI change that moves the file cache into the CLI.

### New Features and Improvements

* Add `u2m.WithDiscoveryHost` option to override the default `https://login.databricks.com` host used by the discovery login flow. Intended for testing and development against non-production environments.
* Add support for unified hosts. A single configuration profile can now be used for both account-level and workspace-level operations when the host supports it and both `AccountID` and `WorkspaceID` are available.

### Bug Fixes

 * Fix CLI token source `--profile` fallback: `--profile` is a global Cobra flag that old CLIs accept silently instead of reporting "unknown flag", making the previous error-based detection dead code. Now uses `databricks version` to detect CLI capabilities at init time ([#1605](https://github.com/databricks/databricks-sdk-go/pull/1605)).

### Internal Changes

 * Pass `--force-refresh` to Databricks CLI `auth token` command to bypass the CLI's internal token cache ([#1628](https://github.com/databricks/databricks-sdk-go/pull/1628)).

### API Changes
* Add [w.TemporaryVolumeCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryVolumeCredentialsAPI) workspace-level service.
* Add `GetPermissionLevels`, `GetPermissions`, `SetPermissions` and `UpdatePermissions` methods for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `ThumbnailUrl` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `JiraOptions`, `OutlookOptions` and `SmartsheetOptions` fields for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `GoogleAdsConfig` field for [pipelines.SourceConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SourceConfig).
* Add `ReplaceExisting` field for [postgres.CreateBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateBranchRequest).
* Add `ReplaceExisting` field for [postgres.CreateEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateEndpointRequest).


## Release v0.129.0 (2026-04-22)

### New Features and Improvements

 * Add `u2m.WithDiscoveryHost` option to override the default `https://login.databricks.com` host used by the discovery login flow. Intended for testing and development against non-production environments.

### Internal Changes

 * Remove `Experimental_IsUnifiedHost` flag from `HostType()` resolution. The method is deprecated and the Terraform provider will use its own host type detection.
 * Expanded AI agent detection: added Goose, Amp, Augment, Copilot (VS Code), Kiro, Windsurf. Honors the `AGENT=<name>` standard and falls back to `unknown` for unrecognized values. When multiple agent env vars are present (e.g. a Cursor CLI subagent invoked from Claude Code), the user-agent reports `agent/multiple`.
 * Use resolved host type from host metadata in `HostType()` method, falling back to URL pattern matching when metadata is unavailable.

### API Changes
* Add [supervisoragents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents) package.
* Add [w.SecretsUc](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecretsUcAPI) workspace-level service.
* Add [w.SupervisorAgents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/supervisoragents#SupervisorAgentsAPI) workspace-level service.
* Add `Update` method for [w.Tokens](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokensAPI) workspace-level service.
* Add `Etag` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Add `Etag` field for [dashboards.GenieUpdateSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieUpdateSpaceRequest).
* Add `BranchId` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `CatalogId` field for [postgres.CatalogCatalogStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CatalogCatalogStatus).
* Add `DatabaseId` field for [postgres.DatabaseDatabaseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#DatabaseDatabaseStatus).
* Add `EndpointId` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `ProjectId` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `RoleId` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `Project` field for [postgres.SyncedTableSyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableStatus).
* Add `Manual` field for [provisioning.CreateGcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateGcpKeyInfo).
* Add `Manual` field for [provisioning.GcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#GcpKeyInfo).
* Add `AppsRuntime` and `LakebaseRuntime` fields for [settings.CustomerFacingIngressNetworkPolicyRequestDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyRequestDestination).
* Add `BlockedInternetDestinations` field for [settings.EgressNetworkPolicyNetworkAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyNetworkAccessPolicy).
* Add `ColumnsToSync` field for [vectorsearch.DeltaSyncVectorIndexSpecResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecResponse).
* Add `BreakingChange` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* [Breaking] Change `UpdateCatalogConfig` method for [w.DataClassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#DataClassificationAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateDefaultWorkspaceBaseEnvironment` method for [w.Environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments#EnvironmentsAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateKnowledgeAssistant` method for [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateBranch`, `UpdateDatabase`, `UpdateEndpoint`, `UpdateProject` and `UpdateRole` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service. Method path has changed.
* [Breaking] Change `UpdateDefaultWarehouseOverride` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service. Method path has changed.


## Release v0.128.0 (2026-04-20)

### New Features and Improvements

 * Add `config.DefaultHostMetadataResolverFactory`: a package-level variable consulted when `Config.HostMetadataResolver` is unset. Lets programs install a shared resolver (e.g. a caching one) once from an `init()` block (typically in a blank-imported package) instead of wiring per-Config. Experimental.

### Bug Fixes

 * Add `X-Databricks-Org-Id` header to `Workspace.Download()` and `Workspace.Upload()` for SPOG host compatibility.
 * `WorkspaceClient.CurrentWorkspaceID()` now returns `Config.WorkspaceID` directly when set, instead of calling `/api/2.0/preview/scim/v2/Me`. This removes an API round-trip on every call where the workspace ID is already known (profile, `?o=` query param, env var, or host metadata) and fixes a failure on SPOG hosts where the unauthenticated probe request was rejected with `Unable to load OAuth Config`.
 * Add `X-Databricks-Org-Id` header to `SharesAPI.internalList()` (`service/sharing/ext_api.go`) for SPOG host compatibility. Same class of bug as #1634 - a hand-written extension method was calling `a.client.Do()` without the header that the generated code paths set per-call.


## Release v0.127.0 (2026-04-15)

### Breaking Changes

 * Raise minimum Go version from 1.18 to 1.24 ([#1558](https://github.com/databricks/databricks-sdk-go/pull/1558)).

### New Features and Improvements

* Added `HostMetadataResolver` hook to allow callers to customize host metadata resolution, e.g. with caching ([#1572](https://github.com/databricks/databricks-sdk-go/pull/1572)).
* Added `NewLimitIterator` to `listing` package for lazy iteration with a cap on output items ([#1555](https://github.com/databricks/databricks-sdk-go/pull/1555)).
* Resolve `TokenAudience` from `token_federation_default_oidc_audiences` in host metadata discovery endpoint.

### Bug Fixes

 * Add `X-Databricks-Org-Id` header to deprecated workspace SCIM APIs (Groups, ServicePrincipals, Users) for SPOG host compatibility.
 * Add retry logic to token acquisition for OIDC, M2M, and Azure client secret credentials ([#1398](https://github.com/databricks/databricks-sdk-go/issues/1398), [#1072](https://github.com/databricks/databricks-sdk-go/issues/1072)).
 * Fix double-caching of OAuth tokens in Azure client secret credentials ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).
 * Disable async token refresh for GCP credential providers to avoid wasted refresh attempts caused by double-caching with Google's internal `oauth2.ReuseTokenSource` ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).
 * Fixed double-caching in M2M OAuth that prevented the proactive async token refresh from reaching the HTTP endpoint until ~10s before expiry, causing bursts of 401 errors at token rotation boundaries ([#1549](https://github.com/databricks/databricks-sdk-go/issues/1549)).
 * Fix data race in `authenticateIfNeeded` when lazily initializing `credentialsProvider` ([#1310](https://github.com/databricks/databricks-sdk-go/issues/1310)).

### Internal Changes

 * Normalize internal token sources on `auth.TokenSource` for proper context propagation ([#1577](https://github.com/databricks/databricks-sdk-go/pull/1577)).
 * Fix `TestAzureGithubOIDCCredentials` hang caused by missing `HTTPTransport` stub: `EnsureResolved` now calls `resolveHostMetadata`, which makes a real network request when no transport is set ([#1550](https://github.com/databricks/databricks-sdk-go/pull/1550)).
 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/slog ([#1566](https://github.com/databricks/databricks-sdk-go/pull/1566)).
 * Bump golang.org/x/net from 0.23.0 to 0.33.0 in /examples/slog ([#1127](https://github.com/databricks/databricks-sdk-go/pull/1127)).
 * Bump golang.org/x/oauth2 from 0.20.0 to 0.27.0 ([#1563](https://github.com/databricks/databricks-sdk-go/pull/1563)).
 * Bump golang.org/x/crypto from 0.21.0 to 0.45.0 in /examples/zerolog ([#1568](https://github.com/databricks/databricks-sdk-go/pull/1568)).
 * Bump golang.org/x/crypto from 0.31.0 to 0.45.0 ([#1564](https://github.com/databricks/databricks-sdk-go/pull/1564)).
 * Bump google.golang.org/grpc from 1.58.3 to 1.79.3 in /examples/zerolog ([#1565](https://github.com/databricks/databricks-sdk-go/pull/1565)).
 * Bump google.golang.org/grpc from 1.62.0 to 1.79.3 in /examples/slog ([#1569](https://github.com/databricks/databricks-sdk-go/pull/1569)).
 * Bump google.golang.org/grpc from 1.64.1 to 1.79.3 ([#1557](https://github.com/databricks/databricks-sdk-go/pull/1557)).

### API Changes
* Add `CreateCatalog`, `CreateSyncedTable`, `DeleteCatalog`, `DeleteSyncedTable`, `GetCatalog` and `GetSyncedTable` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `EffectiveFileEventQueue` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* Add `EffectiveFileEventQueue` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* Add `EffectiveFileEventQueue` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
* Add `ColumnSelection` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `Cascade` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `DefaultBranch` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `DefaultBranch` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `ApplyEnvironment` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
* Add `Name` and `Permission` fields for [apps.AppResourceApp](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceApp).
* Add `Ingress` and `IngressDryRun` fields for [settings.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicy).
* Add `GcpEndpoint` field for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule).
* Add `GcpEndpoint` field for [settings.NccPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRule).
* Add `GcpEndpoint` field for [settings.UpdatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePrivateEndpointRule).
* Add `TableDeltaUniformIcebergExternalDeltasharing` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* Add `ConnectorOptions` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec).
* Add `ConnectorOptions` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec).
* Add `GoogleDrive` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Remove `Project` field for [postgres.SyncedTableSyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#SyncedTableSyncedTableSpec).
* Add `DeleteAppThumbnail` and `UpdateAppThumbnail` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service.
* Add `CreateMessageComment`, `ListConversationComments` and `ListMessageComments` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `ManagedEncryptionSettings` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
* Add `ManagedEncryptionSettings` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
* Add `ManagedEncryptionSettings` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
* Add `Comment` field for [dashboards.GenieFeedback](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieFeedback).
* Add `Thoughts` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Add `Comment` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* Add `RequestSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `IsOnline` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `GoogleAdsOptions` and `TiktokAdsOptions` fields for [pipelines.ConnectorOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ConnectorOptions).
* Add `Scopes` field for [settings.CreateOboTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateOboTokenRequest).
* Add `Scopes` field for [settings.CreateTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateTokenRequest).
* Add `SqlState` field for [sql.StatementStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementStatus).
* Add `UsagePolicyId` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Add `IndexSubtype` field for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest).
* Add `BudgetPolicyId` field for [vectorsearch.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointInfo).
* Add `IndexSubtype` field for [vectorsearch.MiniVectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#MiniVectorIndex).
* Add `BudgetPolicyId` field for [vectorsearch.PatchEndpointBudgetPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#PatchEndpointBudgetPolicyResponse).
* Add `IndexSubtype` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
* Add `StorageOptimized` enum value for [vectorsearch.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointType).
* [Breaking] Remove `Apps` and `Lakebase` fields for [settings.CustomerFacingIngressNetworkPolicyRequestDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingIngressNetworkPolicyRequestDestination).


## Release v0.126.0 (2026-03-20)

### New Features and Improvements
* Add `PersistentAuth.ForceRefreshToken()` to force-refresh cached U2M OAuth tokens and return an error instead of falling back to an existing token when refresh fails.
* Add `ErrMissingRefreshToken`, returned when a token refresh is requested but the cached token has no refresh token (i.e. `Token()` or `ForceRefreshToken()`).


## Release v0.125.0 (2026-03-19)

### New Features and Improvements

* Add automatic detection of AI coding agents (Claude Code, Cursor, Cline, Codex, Copilot CLI, Gemini CLI, OpenCode, Antigravity, Openclaw) in the user-agent string. The SDK now appends `agent/<name>` to HTTP request headers when running inside a known AI agent environment.

### API Changes
* Add `DisableGovTagCreation` field for [settings.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessage).
* Add `DisableGovTagCreation` field for [settingsv2.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#RestrictWorkspaceAdminsMessage).


## Release v0.124.0 (2026-03-18)

### Internal Changes
* Proactively attempt to refresh OAuth tokens expiring within 5 minutes in `PersistentAuth.Token()` to reduce the chance of callers receiving near-expired tokens. If the refresh fails and the token is still valid, the existing token is returned.

### API Changes
* Add `ParentPath` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).


## Release v0.123.0 (2026-03-17)

### New Features and Improvements

* Support `default_profile` in `[__settings__]` section of `.databrickscfg` ([#1534](https://github.com/databricks/databricks-sdk-go/pull/1534)).


## Release v0.122.0 (2026-03-16)

### API Changes
* Add `AlertOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Add `AlertTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `AlertTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `AlertTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add [environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments) package.
* Add [w.Environments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/environments#EnvironmentsAPI) workspace-level service.
* Add `CanCreateApp` enum value for [iam.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionLevel).


## Release v0.121.0 (2026-03-12)

### API Changes
* Add `UpdateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Entities` and `TimeseriesColumn` fields for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `AggregationFunction` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function).
* Add `FilterCondition` field for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource).
* Add `ConnectorType` and `DataStagingOptions` fields for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `IngestionSourceCatalogName`, `IngestionSourceConnectionName`, `IngestionSourceSchemaName`, `IngestionSourceTableName` and `IngestionSourceTableVersion` fields for [pipelines.Origin](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Origin).
* Add `Attributes` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `SubDomain` field for [serving.ExternalFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalFunctionRequest).
* [Breaking] Change `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource) to no longer be required.
* Change `EntityColumns` and `TimeseriesColumn` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource) to no longer be required.
* [Breaking] Change `Inputs` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* Change `Inputs` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* [Breaking] Change `FunctionType` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function) to no longer be required.
* Change `FunctionType` field for [ml.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Function) to no longer be required.
* Change `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource) to no longer be required.
* [Breaking] Change `EntityColumnIdentifiers` and `TimeseriesColumnIdentifier` fields for [ml.KafkaSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaSource) to no longer be required.


## Release v0.120.0 (2026-03-10)

### New Features and Improvements

* Added discovery login flow for `login.databricks.com` to the U2M OAuth infrastructure. This enables the CLI to authenticate users without requiring a workspace URL upfront.

### Internal Changes
- **Auth cache**: Simplify experimental auth token refresh scheduling by replacing stale-duration/token-state tracking with an `asyncRefreshAllowedAfter` timestamp. Async refreshes now start based on a per-token dynamic lead time and, after a failed async refresh, are deferred by 1 minute before retrying while expiry still forces a blocking refresh.


## Release v0.119.0 (2026-03-09)

### API Changes
* Add [dataclassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification) and [knowledgeassistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants) packages.
* Add [w.DataClassification](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataclassification#DataClassificationAPI) workspace-level service.
* Add [w.KnowledgeAssistants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/knowledgeassistants#KnowledgeAssistantsAPI) workspace-level service.
* Add `GenieCreateEvalRun`, `GenieGetEvalResultDetails`, `GenieGetEvalRun`, `GenieListEvalResults` and `GenieListEvalRuns` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `TelemetryExportDestinations` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).


## Release v0.118.0 (2026-03-05)

### Internal Changes

* Fix `GetWorkspaceClient` for unified account hosts ([#1517](https://github.com/databricks/databricks-sdk-go/pull/1517)).
* Add test for `GetWorkspaceClient` with SPOG host ([#1518](https://github.com/databricks/databricks-sdk-go/pull/1518)).


## Release v0.117.0 (2026-03-04)

### Internal Changes
* Move cloud-based credential filtering from individual strategies into `DefaultCredentials`. Azure strategies are skipped on GCP/AWS hosts in auto-detect mode; GCP strategies are skipped on Azure/AWS hosts. When `auth_type` is explicitly set (e.g. `azure-cli`), cloud filtering is bypassed so the named strategy is always attempted regardless of host cloud.


## Release v0.116.0 (2026-03-03)

### Bug Fixes

* Fix `HostType()` misclassifying account hosts when the scheme is missing ([#1510](https://github.com/databricks/databricks-sdk-go/pull/1510)).


## Release v0.115.0 (2026-03-02)

### Internal Changes
* Implement dynamic auth token stale period based on initial token lifetime. Increased up to 20 mins for standard OAuth with proportionally shorter periods for short-lived tokens.

### API Changes
* Add `CreateDatabase`, `DeleteDatabase`, `GetDatabase`, `ListDatabases` and `UpdateDatabase` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `Postgres` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `DataframeSchema`, `FilterCondition` and `TransformationSql` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* Add `EnvironmentVersion` field for [pipelines.PipelinesEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesEnvironment).
* Add `ResetCheckpointSelection` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* [Breaking] Remove `Oauth2AppClientId` and `Oauth2AppIntegrationId` fields for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space).
* Add `EnablePgNativeLogin` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `EnablePgNativeLogin` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* [Breaking] Remove `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).


## Release v0.114.0 (2026-02-26)

### API Changes
* Add `EffectivePublishingMode` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* Add `DbrAutoscale` enum value for [compute.EventDetailsCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetailsCause).
* Change `OutputCatalog` field for [cleanrooms.CreateCleanRoomOutputCatalogResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomOutputCatalogResponse) to be required.
* [Breaking] Remove `InternalAttributes` field for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).
* [Breaking] Remove `InternalAttributes` field for [sharing.Volume](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Volume).


## Release v0.113.0 (2026-02-25)

### New Features and Improvements
* Add `DisableAsyncTokenRefresh` field to `Config` to allow disabling
  asynchronous token refresh for short-lived processes like the Databricks CLI.


## Release v0.112.0 (2026-02-24)

### Bug Fixes

* Pass `--profile` to CLI token source when profile is set, and add read-fallback to migrate legacy host-keyed tokens to profile keys ([#1497](https://github.com/databricks/databricks-sdk-go/pull/1497)).

### API Changes
* Add `Parameters` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `Parameters` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
* [Breaking] Change `GetDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service with new required argument order.
* [Breaking] Change `Name` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space) to be required.
* Change `Name` field for [apps.Space](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#Space) to be required.
* [Breaking] Change `Id` and `UserId` fields for [dashboards.GenieConversation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversation) to no longer be required.
* [Breaking] Change `CreatedTimestamp` and `Title` fields for [dashboards.GenieConversationSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversationSummary) to no longer be required.
* [Breaking] Change `DownloadIdSignature` field for [dashboards.GenieGetDownloadFullQueryResultRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetDownloadFullQueryResultRequest) to be required.
* [Breaking] Change `Id` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage) to no longer be required.


## Release v0.111.0 (2026-02-24)

### New Features and Improvements
* Add `DefaultCredentialStrategyProvider` package-level variable to allow 
  overriding the default authentication strategy for all `Config` instances.

### Documentation
* Added "Retries" section to README.

### Internal Changes

* Add `NewCredentialsChain` to allow internal tools to configure a custom
  credentials chain that differs from the SDK's default chain. This is not
  intended for end-user consumption.
* Export OIDC credential types (`GitHubOIDCCredentials`, `AzureDevOpsOIDCCredentials`,
  `EnvOIDCCredentials`, `FileOIDCCredentials`) so they can be used as building
  blocks in custom credentials chains.
* Make clusters creation in AWS not depend on cloud parsed from host. Changed default AWS availability for auto-created utility clusters from SPOT to SPOT_WITH_FALLBACK (API default).

### API Changes
* Add `PatchEndpoint` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `App` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `SecurableKind` field for [apps.AppResourceUcSecurable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurable).
* Add `ReplaceWhereOverrides` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `ReadOnlyHost` field for [postgres.EndpointHosts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointHosts).
* Add `Group` field for [postgres.EndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSpec).
* Add `Group` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `InitialEndpointSpec` field for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `MinQps` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Add `ScalingInfo` field for [vectorsearch.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointInfo).
* Add `Modify` enum value for [apps.AppResourceUcSecurableUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurablePermission).
* Add `HivemetastoreConnectivityFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `CouldNotGetDashboardSchemaException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `Degraded` enum value for [postgres.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatusState).
* Add `HivemetastoreConnectivityFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).


## Release v0.110.0 (2026-02-19)

### API Changes
* Add `UsingArguments` field for [catalog.ColumnMask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ColumnMask).
* Add `EffectiveEnableFileEvents` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* Add `EffectiveEnableFileEvents` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* Add `InputArguments` field for [catalog.TableRowFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableRowFilter).
* Add `EffectiveEnableFileEvents` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
* Add `Attributes` and `MembershipRoles` fields for [postgres.RoleRoleSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleSpec).
* Add `MembershipRoles` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `GeneralAccess` enum value for [provisioning.EndpointUseCase](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#EndpointUseCase).


## Release v0.109.0 (2026-02-18)

### New Features and Improvements
* Use profile name as OAuth token cache key instead of host URL, so two profiles on the same host store separate tokens.


## Release v0.108.0 (2026-02-16)

### API Changes
* Add [networking](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking) package.
* Add [a.Endpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/networking#EndpointsAPI) account-level service.
* Add `FilterCondition` and `Transformations` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).
* Add `BudgetPolicyId` and `CustomTags` fields for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* Add `BudgetPolicyId` and `CustomTags` fields for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `CreateSpace`, `DeleteSpace`, `GetSpace`, `GetSpaceOperation`, `ListSpaces` and `UpdateSpace` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service.
* Add `Space` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `Space` field for [apps.ListAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ListAppsRequest).
* [Breaking] Remove `FilterCondition` and `Transformations` fields for [ml.DeltaTableSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeltaTableSource).


## Release v0.107.0 (2026-02-12)

### API Changes
* Add `SourceType`, `UpdateTime` and `UpdatedBy` fields for [catalog.EntityTagAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignment).
* Add `SkipNotify` field for [dashboards.Subscription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Subscription).
* Add `SparseCheckout` field for [jobs.GitSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GitSource).
* Add `DisableAutoOptimization`, `MaxRetries`, `MinRetryIntervalMillis` and `RetryOnTimeout` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `DisableAutoOptimization`, `MaxRetries`, `MinRetryIntervalMillis` and `RetryOnTimeout` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `EdgegridAkamai` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).


## Release v0.106.0 (2026-02-10)

### API Changes
* Add `BaseEnvironment` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Add `Filters` field for [jobs.DashboardTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DashboardTask).
* Add `Id` field for [sharing.CreateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateRecipient).
* Add `Id` field for [sharing.RecipientInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientInfo).
* Add `Id` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient).
* Add `QueryTags` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
* Add `QueryTags` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Add `SswsToken` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Add `UcVolumeMisconfigured` enum value for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* [Breaking] Change `Username` field for [iamv2.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#User) to no longer be required.


## Release v0.105.0 (2026-02-09)

### Breaking Changes
* `WorkspaceId` in Config is changed to `WorkspaceID`


## Release v0.104.0 (2026-02-03)

### New Features and Improvements
* Add support for single Profile for Account and Workspace operations in Unified Mode.

### API Changes
* Add `BackfillSource` field for [ml.KafkaConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#KafkaConfig).
* Add `BurstScalingEnabled` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Add `BurstScalingEnabled` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Add `BurstScalingEnabled` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Add `BurstScalingEnabled` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Add `GenerateDownloadFullQueryResult` and `GetDownloadFullQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `ActiveInstances` field for [apps.ComputeStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeStatus).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service. HTTP method/verb has changed.
* Add `Compute` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Add `Compute` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Add `Compute` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Add `MtlsPortConnectivityFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `MtlsPortConnectivityFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).


## Release v0.103.0 (2026-01-29)

### Internal Changes
* Add integration test support for unified host infrastructure ([#1445](https://github.com/databricks/databricks-sdk-go/pull/1445))


## Release v0.102.0 (2026-01-28)

### API Changes
* Add `CloneMode` field for [pipelines.ClonePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ClonePipelineRequest).
* [Breaking] Change `CreateRole` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* Change `RoleId` field for [postgres.CreateRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateRoleRequest) to no longer be required.


## Release v0.101.0 (2026-01-27)

### API Changes
* Add `Force` field for [pipelines.DeletePipelineRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeletePipelineRequest).
* Add `PostgresRole` field for [postgres.RoleRoleSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleSpec).
* Add `PostgresRole` field for [postgres.RoleRoleStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#RoleRoleStatus).
* Add `Purge` field for [sql.TrashAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashAlertV2Request).


## Release v0.100.0 (2026-01-22)

### API Changes
* Add `Outputs` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
* Add `SessionId` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Add `NoSuspension` field for [postgres.EndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSpec).
* Add `NoSuspension` field for [postgres.ProjectDefaultEndpointSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectDefaultEndpointSettings).


## Release v0.99.0 (2026-01-20)

### Breaking Changes

* **[workspace]** Updated Git Credentials List methods to require `ListCredentialsRequest` parameter. The following methods now require a `workspace.ListCredentialsRequest` parameter instead of no parameters:
  - `GitCredentialsAPI.List(ctx, request)`
  - `GitCredentialsAPI.ListAll(ctx, request)`
  - `GitCredentialsAPI.CredentialInfoGitProviderToCredentialIdMap(ctx, request)`

  **Migration Guide:**
  ```go
  // Before
  credentials, err := w.GitCredentials.ListAll(ctx)

  // After
  credentials, err := w.GitCredentials.ListAll(ctx, workspace.ListCredentialsRequest{})
  ```

### API Changes
* Add `PrincipalId` field for [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
* Add `PrincipalId` field for [workspace.DeleteCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#DeleteCredentialsRequest).
* Add `PrincipalId` field for [workspace.GetCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsRequest).
* Add `PrincipalId` field for [workspace.ListCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ListCredentialsRequest).
* Add `PrincipalId` field for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).


## Release v0.98.0 (2026-01-20)

### API Changes
* Add `CreateDefaultWarehouseOverride`, `DeleteDefaultWarehouseOverride`, `GetDefaultWarehouseOverride`, `ListDefaultWarehouseOverrides` and `UpdateDefaultWarehouseOverride` methods for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
* Add `NodeTypeFlexibility` field for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
* Add `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Add `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Add `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Add `DriverNodeTypeFlexibility` and `WorkerNodeTypeFlexibility` fields for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
* Add `FullRefreshWindow` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `AutoFullRefreshPolicy` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `ExpireTime` and `Ttl` fields for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `ExpireTime` field for [postgres.BranchStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchStatus).
* Add `Hosts` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* Add `EndpointTypeReadWrite` and `EndpointTypeReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).
* Add `Deleted` enum value for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* [Breaking] Remove `Default` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* [Breaking] Remove `Host`, `LastActiveTime`, `StartTime` and `SuspendTime` fields for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* [Breaking] Remove `Settings` field for [postgres.ProjectSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectSpec).
* [Breaking] Remove `ComputeLastActiveTime` and `Settings` fields for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* [Breaking] Remove `ReadWrite` and `ReadOnly` enum values for [postgres.EndpointType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointType).
* Add `NoExpiry` field for [postgres.BranchSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#BranchSpec).
* Add `Stderr` field for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* Add `GenerateDatabaseCredential` method for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `GetPublicAccountUserPreference`, `ListAccountUserPreferencesMetadata` and `PatchPublicAccountUserPreference` methods for [a.AccountSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#AccountSettingsV2API) account-level service.


## Release v0.97.0 (2026-01-15)

### New Features and Improvements

* Increase async cache stale period from 3 to 5 minutes to cover the maximum monthly downtime of a 99.99% uptime SLA.

### API Changes
* Add `ErrorMessage` field for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule).
* Add `ErrorMessage` field for [settings.NccPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRule).
* Add `ErrorMessage` field for [settings.UpdatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePrivateEndpointRule).
* Add `RateLimited` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `Creating` and `CreateFailed` enum values for [settings.NccPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRulePrivateLinkConnectionState).
* Add `RateLimited` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteBranch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteBranch).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteEndpoint).
* [Breaking] Add long-running operation for [PostgresAPI.DeleteProject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteProject).
* [Breaking] Change `DeleteBranch`, `DeleteEndpoint` and `DeleteProject` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service to return [postgres.Operation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Operation).
* [Breaking] Remove `PgbouncerSettings` field for [postgres.EndpointSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSettings).
* [Breaking] Remove `PoolerMode` field for [postgres.EndpointSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointSpec).
* [Breaking] Remove `PoolerMode` field for [postgres.EndpointStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#EndpointStatus).
* [Breaking] Remove `PgbouncerSettings` field for [postgres.ProjectDefaultEndpointSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectDefaultEndpointSettings).
* Add `Owner` field for [postgres.ProjectStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#ProjectStatus).
* Add `ValidityCheckConfigurations` field for [qualitymonitorv2.QualityMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#QualityMonitor).
* Add `BurstScalingEnabled` field for [serving.PtServedModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PtServedModel).
* Add `SystemManaged` enum value for [jobs.JobDeploymentKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeploymentKind).
* [Breaking] Change `CreateBranch`, `CreateEndpoint` and `CreateProject` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service with new required argument order.
* [Breaking] Change `BranchId` field for [postgres.CreateBranchRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateBranchRequest) to be required.
* [Breaking] Change `EndpointId` field for [postgres.CreateEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateEndpointRequest) to be required.
* [Breaking] Change `ProjectId` field for [postgres.CreateProjectRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#CreateProjectRequest) to be required.


## Release v0.96.0 (2026-01-06)

### API Changes
* Add `GitRepository` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `GitSource` field for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment).
* Add `ExperimentSpec` field for [apps.AppManifestAppResourceSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceSpec).
* Add `Experiment` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `GitRepository` field for [apps.AppUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppUpdate).
* Add `ExcludedTableFullNames` field for [dataquality.AnomalyDetectionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#AnomalyDetectionConfig).
* Add `GroupName` field for [jobs.JobRunAs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobRunAs).
* Add `RowFilter` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Add `ExcludedTableFullNames` field for [qualitymonitorv2.AnomalyDetectionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#AnomalyDetectionConfig).
* Add `Execute` and `UseConnection` enum values for [apps.AppManifestAppResourceUcSecurableSpecUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurablePermission).
* Add `Function` and `Connection` enum values for [apps.AppManifestAppResourceUcSecurableSpecUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurableType).
* Add `Select`, `Execute` and `UseConnection` enum values for [apps.AppResourceUcSecurableUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurablePermission).
* Add `Table`, `Function` and `Connection` enum values for [apps.AppResourceUcSecurableUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResourceUcSecurableUcSecurableType).
* Add `Spec` and `Status` fields for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
* [Breaking] Remove `AutoscalingLimitMaxCu`, `AutoscalingLimitMinCu`, `CurrentState`, `Disabled`, `EffectiveAutoscalingLimitMaxCu`, `EffectiveAutoscalingLimitMinCu`, `EffectiveDisabled`, `EffectivePoolerMode`, `EffectiveSettings`, `EffectiveSuspendTimeoutDuration`, `EndpointType`, `Host`, `LastActiveTime`, `PendingState`, `PoolerMode`, `Settings`, `StartTime`, `SuspendTime` and `SuspendTimeoutDuration` fields for [postgres.Endpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Endpoint).
* Add `CreateRole`, `DeleteRole`, `GetRole` and `ListRoles` methods for [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `DestinationSourceSecurable` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations).
* Add `Spec` and `Status` fields for [postgres.Branch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Branch).
* Add `Spec` and `Status` fields for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `AccessModes` and `StorageLocation` fields for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).
* [Breaking] Remove `CurrentState`, `Default`, `EffectiveDefault`, `EffectiveIsProtected`, `EffectiveSourceBranch`, `EffectiveSourceBranchLsn`, `EffectiveSourceBranchTime`, `IsProtected`, `LogicalSizeBytes`, `PendingState`, `SourceBranch`, `SourceBranchLsn`, `SourceBranchTime` and `StateChangeTime` fields for [postgres.Branch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Branch).
* [Breaking] Remove `BranchLogicalSizeLimitBytes`, `ComputeLastActiveTime`, `DefaultEndpointSettings`, `DisplayName`, `EffectiveDefaultEndpointSettings`, `EffectiveDisplayName`, `EffectiveHistoryRetentionDuration`, `EffectivePgVersion`, `EffectiveSettings`, `HistoryRetentionDuration`, `PgVersion`, `Settings` and `SyntheticStorageSizeBytes` fields for [postgres.Project](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#Project).
* Add `Command` and `EnvVars` fields for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment).
* Add `FullName` and `SecurableType` fields for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations).
* [Breaking] Change `DeleteKafkaConfig` method for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service . Method path has changed.
* Add `TableDeltasharingOpenDirBased` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* [Breaking] Change long-running operation for [PostgresAPI.DeleteRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI.DeleteRole). Long running operation response type changed to empty response.
* [Breaking] Remove `AccessModes` and `StorageLocation` fields for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).


## Release v0.95.0 (2025-12-17)

### API Changes
* Add [postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres) package.
* Add [w.Postgres](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/postgres#PostgresAPI) workspace-level service.
* Add `EffectiveUsagePolicyId` and `UsagePolicyId` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `ExternalAccessEnabled` field for [catalog.CreateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateAccountsMetastore).
* Add `ExternalAccessEnabled` field for [catalog.CreateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastore).
* Add `ExternalAccessEnabled` field for [catalog.UpdateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAccountsMetastore).
* Add `ExternalAccessEnabled` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore).
* Add `UsagePolicyId` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore).
* Add `ErrorMessage` field for [settings.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule).
* Add `ErrorMessage` field for [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule).
* Add `ControlPlaneConnectionFailure` and `ControlPlaneConnectionFailureDueToMisconfig` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `ControlPlaneConnectionFailure` and `ControlPlaneConnectionFailureDueToMisconfig` enum values for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).


## Release v0.94.0 (2025-12-11)

### Bug Fixes

* Fixed U2M OAuth token caching issue that caused tokens to be read from disk 
  on every request instead of being cached in memory.

### API Changes
* Add `Clone` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
* Add `CronSchedule` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature).
* Add `GcpServiceAccount` field for [provisioning.CreateGcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateGcpKeyInfo).
* Add `GcpServiceAccount` field for [provisioning.GcpKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#GcpKeyInfo).
* Add `HasDeltaUniformIceberg` field for [sharing.TableInternalAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributes).
* Change `TimeWindow` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.
* [Breaking] Change `TimeWindow` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature) to no longer be required.


## Release v0.93.0 (2025-12-03)

### Bug Fixes

* Support for Spark versions with multiple Scala versions ([#1331](https://github.com/databricks/databricks-sdk-go/pull/1331)).

### API Changes
* Add [w.WorkspaceEntityTagAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#WorkspaceEntityTagAssignmentsAPI) workspace-level service.
* Add `DatasetCatalog` and `DatasetSchema` fields for [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
* Add `DatasetCatalog` and `DatasetSchema` fields for [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
* Add `PurgeData` field for [database.DeleteSyncedDatabaseTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DeleteSyncedDatabaseTableRequest).
* Add `Truncation` field for [pipelines.PipelineEvent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineEvent).
* Add `ForeignTable` and `Volume` enum values for [sharing.SharedDataObjectDataObjectType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObjectDataObjectType).


## Release v0.92.0 (2025-11-20)

### API Changes
* Add `CreateSpace` and `UpdateSpace` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Add `BatchCreateMaterializedFeatures`, `CreateKafkaConfig`, `DeleteKafkaConfig`, `GetKafkaConfig`, `ListKafkaConfigs` and `UpdateKafkaConfig` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `DeleteOnlineTable` method for [w.FeatureStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureStoreAPI) workspace-level service.
* Add `RetrieveUserVisibleMetrics` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Add `MajorVersion` field for [billing.CreateBillingUsageDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBillingUsageDashboardRequest).
* Add `IncludeSerializedSpace` field for [dashboards.GenieGetSpaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetSpaceRequest).
* Add `SerializedSpace` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Add `Purpose` field for [dashboards.TextAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachment).
* Add `BudgetPolicyId` field for [database.NewPipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#NewPipelineSpec).
* Add `Model` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* Add `KafkaSource` field for [ml.DataSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DataSource).
* Add `LineageContext` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* Add `ConnectionParameters` field for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition).
* Add `IngestFromUcForeignCatalog` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Add `RewindSpec` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
* Add `TypeText` field for [vectorsearch.ColumnInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ColumnInfo).
* Add `AutoscaleV2` enum value for [compute.EventDetailsCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetailsCause).
* Add `UnsupportedConversationTypeException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `ForeignCatalog` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Add `Creating` and `CreateFailed` enum values for [settings.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState).
* Add `Creating` and `CreateFailed` enum values for [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState).
* Add `RedState` and `YellowState` enum values for [vectorsearch.EndpointStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointStatusState).
* Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* [Breaking] Change `Destinations` field for [catalog.AccessRequestDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccessRequestDestinations) to no longer be required.
* [Breaking] Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* Change `TableNames` field for [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration) to be required.
* [Breaking] Change `OnlineStoreConfig` field for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to type [ml.OnlineStoreConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStoreConfig).


## Release v0.91.0 (2025-11-13)

### Internal Changes

* Add experimental core framework to replace existing API call execution.


## Release v0.90.0 (2025-11-03)

### New Features and Improvements
* Add support for unified hosts, i.e. hosts that support both workspace-level and account-level operations
* Deprecate Config.IsAccountClient, which will not work for unified hosts, and replace it with Config.HostType and Config.ConfigType methods.

### API Changes
* Add `GoogleAds`, `TiktokAds`, `SalesforceMarketingCloud`, `Hubspot`, `WorkdayHcm`, `Guidewire` and `Zendesk` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).


## Release v0.89.0 (2025-10-30)

### API Changes
* Add `InstanceProfileArn` field for [compute.InstancePoolAwsAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAwsAttributes).
* Add `Continuous`, `Sliding` and `Tumbling` fields for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).
* Add `UsagePolicyId` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Add `UsagePolicyId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Add `UsagePolicyId` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Add `ReadFilesBytes` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Add `Select` enum value for [apps.AppManifestAppResourceUcSecurableSpecUcSecurablePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurablePermission).
* Add `Table` enum value for [apps.AppManifestAppResourceUcSecurableSpecUcSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppManifestAppResourceUcSecurableSpecUcSecurableType).
* Add `DecommissionStarted` and `DecommissionEnded` enum values for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Add `DbrImageResolutionFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `DbrImageResolutionFailure` enum value for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* Change `OfflineStoreConfig` and `OnlineStoreConfig` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to no longer be required.
* [Breaking] Change `OfflineStoreConfig` and `OnlineStoreConfig` fields for [ml.MaterializedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeature) to no longer be required.
* [Breaking] Change `LifecycleState` field for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to type [sql.AlertLifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertLifecycleState).
* [Breaking] Remove `Table` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
* [Breaking] Remove `Duration` and `Offset` fields for [ml.TimeWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TimeWindow).


## Release v0.88.0 (2025-10-23)

### API Changes
* Add `CreateMaterializedFeature`, `DeleteMaterializedFeature`, `GetMaterializedFeature`, `ListMaterializedFeatures` and `UpdateMaterializedFeature` methods for [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Add `FilterCondition` field for [ml.Feature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Feature).
* [Breaking] Change `DisplayName`, `Evaluation`, `QueryText`, `Schedule` and `WarehouseId` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to be required.
* Change `DisplayName`, `Evaluation`, `QueryText`, `Schedule` and `WarehouseId` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2) to be required.
* [Breaking] Change `ComparisonOperator` and `Source` fields for [sql.AlertV2Evaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2Evaluation) to be required.
* Change `ComparisonOperator` and `Source` fields for [sql.AlertV2Evaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2Evaluation) to be required.
* [Breaking] Change `Name` field for [sql.AlertV2OperandColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2OperandColumn) to be required.
* Change `Name` field for [sql.AlertV2OperandColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2OperandColumn) to be required.
* [Breaking] Change `QuartzCronSchedule` and `TimezoneId` fields for [sql.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CronSchedule) to be required.
* Change `QuartzCronSchedule` and `TimezoneId` fields for [sql.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CronSchedule) to be required.
* [Breaking] Remove `Results` field for [sql.ListAlertsV2Response](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsV2Response).


## Release v0.87.0 (2025-10-20)

### API Changes
* Add `AbsoluteSessionLifetimeInMinutes` and `EnableSingleUseRefreshTokens` fields for [oauth2.TokenAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#TokenAccessPolicy).
* Add `NetworkConnectivityConfigId` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
* Add `OauthMtls` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Add `NetworkCheckNicFailureDueToMisconfig`, `NetworkCheckDnsServerFailureDueToMisconfig`, `NetworkCheckStorageFailureDueToMisconfig`, `NetworkCheckMetadataEndpointFailureDueToMisconfig`, `NetworkCheckControlPlaneFailureDueToMisconfig` and `NetworkCheckMultipleComponentsFailureDueToMisconfig` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `Creating` and `CreateFailed` enum values for [settings.NccPrivateEndpointRulePrivateLinkConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRulePrivateLinkConnectionState).
* Add `NetworkCheckNicFailureDueToMisconfig`, `NetworkCheckDnsServerFailureDueToMisconfig`, `NetworkCheckStorageFailureDueToMisconfig`, `NetworkCheckMetadataEndpointFailureDueToMisconfig`, `NetworkCheckControlPlaneFailureDueToMisconfig` and `NetworkCheckMultipleComponentsFailureDueToMisconfig` enum values for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* [Breaking] Remove `Update` method for [w.RecipientFederationPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientFederationPoliciesAPI) workspace-level service.


## Release v0.86.0 (2025-10-13)

### API Changes
* Add [dataquality](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality) package.
* Add [w.DataQuality](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#DataQualityAPI) workspace-level service.
* Add `CreateUpdate` and `GetUpdate` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service.
* Add `UpdateNotifications` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Add `ComputeSize` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Add `GenieSpace` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Add `SkipValidation` field for [catalog.AccountsCreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateStorageCredential).
* Add `SkipValidation` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential).
* Add `Aliases`, `BrowseOnly`, `CreatedAt`, `CreatedBy`, `FullName`, `MetastoreId`, `Owner`, `UpdatedAt` and `UpdatedBy` fields for [catalog.CreateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateRegisteredModelRequest).
* Add `CatalogName`, `Id`, `ModelName` and `SchemaName` fields for [catalog.RegisteredModelAlias](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelAlias).
* Add `Aliases`, `CatalogName`, `CreatedAt`, `CreatedBy`, `Id`, `MetastoreId`, `ModelName`, `ModelVersionDependencies`, `RunId`, `RunWorkspaceId`, `SchemaName`, `Source`, `Status`, `StorageLocation`, `UpdatedAt` and `UpdatedBy` fields for [catalog.UpdateModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateModelVersionRequest).
* Add `Aliases`, `BrowseOnly`, `CatalogName`, `CreatedAt`, `CreatedBy`, `MetastoreId`, `Name`, `SchemaName`, `StorageLocation`, `UpdatedAt` and `UpdatedBy` fields for [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
* Add `Parameters` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Add `DatabaseInstanceName` field for [database.CreateDatabaseInstanceRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#CreateDatabaseInstanceRoleRequest).
* Add `CustomTags`, `EffectiveCustomTags`, `EffectiveUsagePolicyId` and `UsagePolicyId` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Add `EffectiveAttributes` and `InstanceName` fields for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole).
* Add `KeyRegion` field for [provisioning.CreateAwsKeyInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateAwsKeyInfo).
* Add `RoleArn` field for [provisioning.CreateStorageConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateStorageConfigurationRequest).
* Add `AzureKeyInfo` field for [provisioning.CustomerManagedKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomerManagedKey).
* [Breaking] Add `CustomerFacingPrivateAccessSettings` field for [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* Add `RoleArn` field for [provisioning.StorageConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageConfiguration).
* [Breaking] Add `CustomerFacingWorkspace` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
* Add `UpdateMask` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
* Add `ComputeMode`, `Network`, `NetworkConnectivityConfigId` and `StorageMode` fields for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* Add `EnableServerlessCompute` field for [sql.GetWorkspaceWarehouseConfigResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWorkspaceWarehouseConfigResponse).
* Add `PageSize` and `PageToken` fields for [sql.ListWarehousesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListWarehousesRequest).
* Add `NextPageToken` field for [sql.ListWarehousesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListWarehousesResponse).
* Add `EnableServerlessCompute` field for [sql.SetWorkspaceWarehouseConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#SetWorkspaceWarehouseConfigRequest).
* Add `ModelVersionStatusUnknown` enum value for [catalog.ModelVersionInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfoStatus).
* Add `StreamNative` enum value for [catalog.SystemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemType).
* Add `K8sActivePodQuotaExceeded` and `CloudAccountPodQuotaExceeded` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Add `ExceededMaxTokenLengthException`, `InternalCatalogAssetCreationOngoingException`, `InternalCatalogAssetCreationFailedException` and `InternalCatalogAssetCreationUnsupportedException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Add `SshBootstrapFailure`, `AwsInaccessibleKmsKeyFailure`, `InitContainerNotFinished`, `SparkImageDownloadThrottled`, `SparkImageNotFound`, `ClusterOperationThrottled`, `ClusterOperationTimeout`, `ServerlessLongRunningTerminated`, `AzurePackedDeploymentPartialFailure`, `InvalidWorkerImageFailure`, `WorkspaceUpdate`, `InvalidAwsParameter`, `DriverOutOfDisk`, `DriverOutOfMemory`, `DriverLaunchTimeout`, `DriverUnexpectedFailure`, `UnexpectedPodRecreation`, `GcpInaccessibleKmsKeyFailure`, `GcpKmsKeyPermissionDenied`, `DriverEviction`, `UserInitiatedVmTermination`, `GcpIamTimeout`, `AwsResourceQuotaExceeded`, `CloudAccountSetupFailure`, `AwsInvalidKeyPair`, `DriverPodCreationFailure`, `MaintenanceMode`, `InternalCapacityFailure`, `ExecutorPodUnscheduled`, `StorageDownloadFailureSlow`, `StorageDownloadFailureThrottled`, `DynamicSparkConfSizeExceeded`, `AwsInstanceProfileUpdateFailure`, `InstancePoolNotFound`, `InstancePoolMaxCapacityReached`, `AwsInvalidKmsKeyState`, `GcpInsufficientCapacity`, `GcpApiRateQuotaExceeded`, `GcpResourceQuotaExceeded`, `GcpIpSpaceExhausted`, `GcpServiceAccountAccessDenied`, `GcpServiceAccountNotFound`, `GcpForbidden`, `GcpNotFound`, `ResourceUsageBlocked`, `DataAccessConfigChanged`, `AccessTokenFailure`, `InvalidInstancePlacementProtocol`, `BudgetPolicyResolutionFailure`, `InPenaltyBox`, `DisasterRecoveryReplication`, `BootstrapTimeoutDueToMisconfig`, `InstanceUnreachableDueToMisconfig`, `StorageDownloadFailureDueToMisconfig`, `ControlPlaneRequestFailureDueToMisconfig`, `CloudProviderLaunchFailureDueToMisconfig`, `GcpSubnetNotReady`, `CloudOperationCancelled`, `CloudProviderInstanceNotLaunched`, `GcpTrustedImageProjectsViolated`, `BudgetPolicyLimitEnforcementActivated`, `EosSparkImage`, `NoMatchedK8s`, `LazyAllocationTimeout`, `DriverNodeUnreachable`, `SecretCreationFailure`, `PodSchedulingFailure`, `PodAssignmentFailure`, `AllocationTimeout`, `AllocationTimeoutNoUnallocatedClusters`, `AllocationTimeoutNoMatchedClusters`, `AllocationTimeoutNoReadyClusters`, `AllocationTimeoutNoWarmedUpClusters`, `AllocationTimeoutNodeDaemonNotReady`, `AllocationTimeoutNoHealthyClusters`, `NetvisorSetupTimeout`, `NoMatchedK8sTestingTag`, `CloudProviderResourceStockoutDueToMisconfig`, `GkeBasedClusterTermination`, `AllocationTimeoutNoHealthyAndWarmedUpClusters`, `DockerInvalidOsException`, `DockerContainerCreationException`, `DockerImageTooLargeForInstanceException`, `DnsResolutionError`, `GcpDeniedByOrgPolicy`, `SecretPermissionDenied`, `NetworkCheckNicFailure`, `NetworkCheckDnsServerFailure`, `NetworkCheckStorageFailure`, `NetworkCheckMetadataEndpointFailure`, `NetworkCheckControlPlaneFailure`, `NetworkCheckMultipleComponentsFailure`, `DriverUnhealthy`, `SecurityAgentsFailedInitialVerification`, `DriverDnsResolutionFailure`, `NoActivatedK8s`, `UsagePolicyEntitlementDenied`, `NoActivatedK8sTestingTag`, `K8sActivePodQuotaExceeded` and `CloudAccountPodQuotaExceeded` enum values for [sql.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TerminationReasonCode).
* [Breaking] Change `Create` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to start returning [catalog.AccountsCreateMetastoreAssignmentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateMetastoreAssignmentResponse).
* [Breaking] Change `Delete` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to start returning [catalog.AccountsDeleteMetastoreAssignmentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsDeleteMetastoreAssignmentResponse).
* [Breaking] Change `Update` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to start returning [catalog.AccountsUpdateMetastoreAssignmentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateMetastoreAssignmentResponse).
* [Breaking] Change `Create` method for [a.AccountMetastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoresAPI) account-level service to return [catalog.AccountsCreateMetastoreResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateMetastoreResponse).
* [Breaking] Change `Delete` method for [a.AccountMetastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoresAPI) account-level service to start returning [catalog.AccountsDeleteMetastoreResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsDeleteMetastoreResponse).
* [Breaking] Change `Get` method for [a.AccountMetastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoresAPI) account-level service to return [catalog.AccountsGetMetastoreResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsGetMetastoreResponse).
* [Breaking] Change `List` method for [a.AccountMetastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoresAPI) account-level service to return [catalog.AccountsListMetastoresResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsListMetastoresResponse).
* [Breaking] Change `Update` method for [a.AccountMetastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoresAPI) account-level service to return [catalog.AccountsUpdateMetastoreResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateMetastoreResponse).
* [Breaking] Change `Create` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsCreateStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateStorageCredentialInfo).
* [Breaking] Change `Delete` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to start returning [catalog.AccountsDeleteStorageCredentialResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsDeleteStorageCredentialResponse).
* [Breaking] Change `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsUpdateStorageCredentialResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredentialResponse).
* [Breaking] Change `Create` method for [w.RegisteredModels](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelsAPI) workspace-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CredentialsAPI) account-level service to start returning [provisioning.Credential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Credential).
* [Breaking] Change `Delete` method for [a.EncryptionKeys](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#EncryptionKeysAPI) account-level service to start returning [provisioning.CustomerManagedKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomerManagedKey).
* [Breaking] Change `Create` method for [a.Networks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworksAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.Networks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworksAPI) account-level service to start returning [provisioning.Network](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Network).
* [Breaking] Change `Create` and `Replace` methods for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` and `Replace` methods for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service to start returning [provisioning.PrivateAccessSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessSettings).
* [Breaking] Change `Delete` method for [a.Storage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageAPI) account-level service to start returning [provisioning.StorageConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StorageConfiguration).
* [Breaking] Change `Create` method for [a.VpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpointsAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` method for [a.VpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpointsAPI) account-level service to start returning [provisioning.VpcEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#VpcEndpoint).
* [Breaking] Change `Create` and `Update` methods for [a.Workspaces](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI) account-level service with new required argument order.
* [Breaking] Change `Delete` and `Update` methods for [a.Workspaces](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI) account-level service to start returning [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* [Breaking] Change `ExecuteStatement` method for [w.StatementExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementExecutionAPI) workspace-level service . Method path has changed.
* [Breaking] Change `MetastoreInfo` field for [catalog.AccountsCreateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateMetastore) to type [catalog.CreateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateAccountsMetastore).
* [Breaking] Change `CredentialInfo` field for [catalog.AccountsCreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsCreateStorageCredential) to type [catalog.CreateAccountsStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateAccountsStorageCredential).
* [Breaking] Change `MetastoreInfo` field for [catalog.AccountsUpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateMetastore) to type [catalog.UpdateAccountsMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAccountsMetastore).
* [Breaking] Change `CredentialInfo` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential) to type [catalog.UpdateAccountsStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAccountsStorageCredential).
* Change `CatalogName`, `Name` and `SchemaName` fields for [catalog.CreateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateRegisteredModelRequest) to no longer be required.
* [Breaking] Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.
* Change `Name` field for [database.DatabaseInstanceRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstanceRole) to be required.
* Change `NetworkName` field for [provisioning.CreateNetworkRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateNetworkRequest) to no longer be required.
* Change `PrivateAccessSettingsName` and `Region` fields for [provisioning.CreatePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreatePrivateAccessSettingsRequest) to no longer be required.
* Change `VpcEndpointName` field for [provisioning.CreateVpcEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateVpcEndpointRequest) to no longer be required.
* Change `WorkspaceName` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest) to no longer be required.
* [Breaking] Change `DataplaneRelay` and `RestApi` fields for [provisioning.NetworkVpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworkVpcEndpoints) to no longer be required.
* Change `DataplaneRelay` and `RestApi` fields for [provisioning.NetworkVpcEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#NetworkVpcEndpoints) to no longer be required.
* [Breaking] Change waiter for [WorkspacesAPI.Update](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#WorkspacesAPI.Update).
* [Breaking] Remove `BrowseOnly` field for [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
* [Breaking] Remove `JarDependencies` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* [Breaking] Remove `IsNoPublicIpEnabled` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
* [Breaking] Remove `AllowedVpcEndpointIds`, `PrivateAccessLevel`, `PrivateAccessSettingsName`, `PublicAccessEnabled` and `Region` fields for [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* [Breaking] Remove `ExternalId` field for [provisioning.StsRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#StsRole).
* [Breaking] Remove `AwsRegion`, `CredentialsId`, `CustomTags`, `ManagedServicesCustomerManagedKeyId`, `NetworkConnectivityConfigId`, `NetworkId`, `PrivateAccessSettingsId`, `StorageConfigurationId` and `StorageCustomerManagedKeyId` fields for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
* [Breaking] Remove `ExternalCustomerInfo` and `IsNoPublicIpEnabled` fields for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* [Breaking] Remove `StatusUnspecified` enum value for [sql.Status](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Status).
* Add `IncludeUnbound` field for [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
* Add `IncludeUnbound` field for [catalog.ListCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCredentialsRequest).
* Add `IncludeUnbound` field for [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest).
* Add `IncludeUnbound` field for [catalog.ListStorageCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsRequest).
* Add `AssetTypeMcp` enum value for [marketplace.AssetType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#AssetType).
* Add `ComputeMode` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
* Add `ExpectedWorkspaceStatus` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
* Add `DependencyStorageLocations` field for [sharing.TableInternalAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributes).
* Add `GitEmail` field for [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
* Add `GitEmail` field for [workspace.CreateCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsResponse).
* Add `GitEmail` field for [workspace.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CredentialInfo).
* Add `GitEmail` field for [workspace.GetCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsResponse).
* Add `GitEmail` field for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).
* Add `GermanyTisax` enum value for [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
* [Breaking] Remove `PredictionProbabilityColumn` field for [dataquality.InferenceLogConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dataquality#InferenceLogConfig).


## Release v0.85.0 (2025-09-23)

### New Features and Improvements

* Add native support for authentication through Azure DevOps OIDC.


## Release v0.84.0 (2025-09-18)

### New Features and Improvements
* Increased rate limiter burst setting from 1 to rate\_limit value.

### API Changes
* Added [w.FeatureEngineering](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureEngineeringAPI) workspace-level service.
* Added `JavaDependencies` field for [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
* Added `FollowupQuestions` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Added `EffectiveCapacity` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `CreateTime` and `UpdateTime` fields for [tags.TagPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPolicy).
* Added `TableDeltaUniformIcebergForeignDeltasharing` enum value for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* Added `InternalCatalogPathOverlapException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* [Breaking] Removed `DefaultDataSecurityMode` and `EffectiveDefaultDataSecurityMode` fields for [settingsv2.Setting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#Setting).
* Added [a.AccountGroupsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountGroupsV2API) account-level service, [a.AccountServicePrincipalsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountServicePrincipalsV2API) account-level service, [a.AccountUsersV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccountUsersV2API) account-level service, [w.GroupsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GroupsV2API) workspace-level service, [w.ServicePrincipalsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipalsV2API) workspace-level service and [w.UsersV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersV2API) workspace-level service.
* Added `NetsuiteJarPath` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `WorkdayReportParameters` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `InternalCatalogMissingUcPathException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `ListShares` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service.
* [Breaking] Removed `List` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service.
* Added `SuggestedQuestions` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Added `WarehouseId` field for [dashboards.GenieSpace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSpace).
* Added `Palantir` enum value for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Added `TableMetricViewDeltasharing` and `TableForeignPalantir` enum values for [catalog.SecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableKind).
* Added `NoActivatedK8sTestingTag` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `MetricView` enum value for [sharing.TableInternalAttributesSharedTableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributesSharedTableType).
* [Breaking] Removed `FollowupQuestions` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* [Breaking] Removed `Comment` field for [dashboards.GenieFeedback](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieFeedback).
* [Breaking] Removed `Comment` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* Added `ExternalUseSchema` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).


## Release v0.83.0 (2025-09-08)

### API Changes
* Added [iamv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2) package.
* Added [a.AccountIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#AccountIamV2API) account-level service and [w.WorkspaceIamV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iamv2#WorkspaceIamV2API) workspace-level service.
* Added `Feedback` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage).
* Added `Disabled` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `AuxiliaryManagedLocation` field for [sharing.TableInternalAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributes).
* Added `Alerts` field for [sql.ListAlertsV2Response](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsV2Response).
* Added `NoActivatedK8s` and `UsagePolicyEntitlementDenied` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `ForeignCatalog` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `ForeignIcebergTable` enum value for [sharing.TableInternalAttributesSharedTableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributesSharedTableType).
* [Breaking] Removed `Disabled` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).


## Release v0.82.0 (2025-08-31)

### Bug Fixes

- Fixed Azure CLI authentication for federated token service principals in AKS 
  workload identity environments. The SDK now properly detects service principals 
  with GUID-like names as federated tokens and skips tenant ID parameters that 
  cause authentication failures.

### API Changes
* Added `Comment` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* [Breaking] Added `Rating` field for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).
* Added `EffectiveEnablePgNativeLogin` and `EnablePgNativeLogin` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `TaskRetryMode` field for [jobs.Continuous](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Continuous).
* Added `SourceConfigurations` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `AppId`, `AppIdSet`, `AuthSecret`, `AuthSecretSet`, `ChannelUrl`, `ChannelUrlSet`, `TenantId` and `TenantIdSet` fields for [settings.MicrosoftTeamsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#MicrosoftTeamsConfig).
* Added `EnsureRerankerCompatible` field for [vectorsearch.GetIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#GetIndexRequest).
* Added `Reranker` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* [Breaking] Changed `CreateCleanRoomAssetReview` method for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `SendMessageFeedback` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service with new required argument order.
* Changed `NotebookReview` field for [cleanrooms.CreateCleanRoomAssetReviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CreateCleanRoomAssetReviewRequest) to no longer be required.
* [Breaking] Changed `Features` field for [ml.FeatureList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureList) to type [][ml.LinkedFeature](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LinkedFeature).
* [Breaking] Removed `FeedbackRating` and `FeedbackText` fields for [dashboards.GenieSendMessageFeedbackRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieSendMessageFeedbackRequest).


## Release v0.81.0 (2025-08-20)

### Bug Fixes

- Update the `U2M` port selection mechanism to try port `8020` first and fall
  back incrementally (to port `8021`, `8022`, and so on...) if a port is not 
  available. This fixes an issue with the Databricks CLI which is only 
  allowlisted on port `8020`. 

### API Changes
* Added [settingsv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2) and [tags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags) packages.
* Added [w.AppsSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsSettingsAPI) workspace-level service.
* Added [w.EntityTagAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EntityTagAssignmentsAPI) workspace-level service and [w.Rfa](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RfaAPI) workspace-level service.
* Added [a.AccountSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#AccountSettingsV2API) account-level service and [w.WorkspaceSettingsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settingsv2#WorkspaceSettingsV2API) workspace-level service.
* Added [w.TagPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/tags#TagPoliciesAPI) workspace-level service.
* Added `DeleteConversationMessage`, `ListConversationMessages` and `SendMessageFeedback` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `IncludeAll` field for [dashboards.GenieListConversationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieListConversationsRequest).
* Added `EffectiveUsagePolicyId` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
* Added `EffectiveUsagePolicyId` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
* Added `EffectiveUsagePolicyId` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
* Added `EffectiveUsagePolicyId` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
* Added `Tokens` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit).
* Added `UsagePolicyId` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `EffectiveRunAs` and `RunAs` fields for [sql.AlertV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertV2).
* Added `CacheQueryId` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `ModelEndpointNameForQuery` field for [vectorsearch.EmbeddingSourceColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EmbeddingSourceColumn).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* [Breaking] Removed `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* [Breaking] Removed `Comment`, `DisplayName` and `Tags` fields for [sharing.Share](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Share).


## Release v0.80.0 (2025-08-13)

### API Changes
* Added [w.Policies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PoliciesAPI) workspace-level service and [w.TemporaryPathCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryPathCredentialsAPI) workspace-level service.
* Added `Create` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
* Added `ListDatabaseCatalogs`, `ListSyncedDatabaseTables`, `UpdateDatabaseCatalog` and `UpdateSyncedDatabaseTable` methods for [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added `FirstOnDemand` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
* Added `UsagePolicyId` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
* Added `UsagePolicyId` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
* Added `UsagePolicyId` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
* Added `ClientRequestId` and `UsageContext` fields for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
* Added `ChannelId`, `ChannelIdSet`, `OauthToken` and `OauthTokenSet` fields for [settings.SlackConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SlackConfig).
* Added `Snapshot` enum value for [ml.PublishSpecPublishMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpecPublishMode).
* [Breaking] Changed `PublishMode` field for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec) to be required.


## Release v0.79.0 (2025-08-06)

### API Changes
* Added `StatementIdSignature` field for [dashboards.Result](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Result).
* Added `EffectiveDatabaseInstanceName` and `EffectiveLogicalDatabaseName` fields for [database.SyncedDatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedDatabaseTable).
* Added `Table` field for [jobs.TriggerStateProto](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerStateProto).
* Added `EmailNotifications` field for [serving.CreatePtEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreatePtEndpointRequest).
* Added `EmailNotifications` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `EmailNotifications` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* [Breaking] Changed `List` method for [w.ConsumerProviders](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ConsumerProvidersAPI) workspace-level service . New request type is [marketplace.ListConsumerProvidersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListConsumerProvidersRequest).
* [Breaking] Changed `Create` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service . New request type is [provisioning.CreatePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreatePrivateAccessSettingsRequest).
* [Breaking] Changed `Create` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service with new required argument order.
* [Breaking] Changed `Replace` method for [a.PrivateAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#PrivateAccessAPI) account-level service . New request type is [provisioning.ReplacePrivateAccessSettingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#ReplacePrivateAccessSettingsRequest).
* [Breaking] Removed `IsFeatured` field for [marketplace.ListProvidersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListProvidersRequest).


## Release v0.78.0 (2025-07-31)

### API Changes
* Added `Share` field for [sharing.ListProviderShareAssetsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderShareAssetsResponse).
* Added `ProjectedRemainingWallclockTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Added `MessageAttachmentTooLongError` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `Mask` enum value for [serving.AiGatewayGuardrailPiiBehaviorBehavior](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailPiiBehaviorBehavior).
* [Breaking] Added waiter for [CleanRoomsAPI.Create](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomsAPI.Create).
* [Breaking] Added waiter for [DatabaseAPI.CreateDatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI.CreateDatabaseInstance).


## Release v0.77.0 (2025-07-24)

### New Features and Improvements

* Add new error value `ErrHTMLContent` to differentiate parsing errors caused by
  HTML content from other parsing errors.
* Return more detailed error messages when OAuth endpoints cannot be resolved.
* Use a free port in `u2m` authentication flows rather than 8020.

### API Changes
* Added [agentbricks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/agentbricks) package.
* Added [w.CleanRoomAssetRevisions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetRevisionsAPI) workspace-level service and [w.CleanRoomAutoApprovalRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAutoApprovalRulesAPI) workspace-level service.
* Added `CreateCleanRoomAssetReview` method for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service.
* Added `LatestMonitorFailureMsg` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor).
* Added `LatestMonitorFailureMsg` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
* Added `ProvisioningPhase` field for [database.SyncedTablePipelineProgress](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTablePipelineProgress).
* Added `Unspecified` enum value for [catalog.MonitorCronSchedulePauseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedulePauseStatus).
* Added `Unknown` enum value for [catalog.MonitorRefreshInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoState).
* Added `UnknownTrigger` enum value for [catalog.MonitorRefreshInfoTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoTrigger).
* Added `Redshift` and `Sqldw` enum values for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `GermanyC5` enum value for [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
* [Breaking] Changed `CancelRefresh` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service to start returning [catalog.CancelRefreshResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshResponse).
* [Breaking] Changed `Create` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `Delete` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service to start returning [catalog.DeleteMonitorResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteMonitorResponse).
* [Breaking] Changed `RefreshId` field for [catalog.CancelRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshRequest) to type `int64`.
* [Breaking] Changed `RefreshId` field for [catalog.GetRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRefreshRequest) to type `int64`.
* [Breaking] Changed `MonitorVersion` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to type `int64`.
* Changed `OutputSchemaName` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
* [Breaking] Changed `AssetType` and `Name` fields for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset) to be required.
* Changed `AssetType` and `Name` fields for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetForeignTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetForeignTableLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetForeignTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetForeignTableLocalDetails) to be required.
* [Breaking] Changed `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to be required.
* Changed `NotebookContent` field for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetViewLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetViewLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetViewLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetViewLocalDetails) to be required.
* [Breaking] Changed `LocalName` field for [cleanrooms.CleanRoomAssetVolumeLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetVolumeLocalDetails) to be required.
* Changed `LocalName` field for [cleanrooms.CleanRoomAssetVolumeLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetVolumeLocalDetails) to be required.
* [Breaking] Removed [aibuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder) package.


## Release v0.76.0 (2025-07-17)

### API Changes
* Added `EnvironmentSettings` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
* Added `EnvironmentSettings` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
* Added `EnvironmentSettings` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
* Added `QueryBasedConnectorConfig` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `MetaMarketing` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* [Breaking] Removed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service.


## Release v0.75.0 (2025-07-10)

### Bug Fixes

- Default values for `ApiClient` are now properly set in the type constructor rather than in the config constructor. 

### API Changes
* Added [w.ServicePrincipalSecretsProxy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ServicePrincipalSecretsProxyAPI) workspace-level service.
* Added [w.DefaultWarehouseId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultWarehouseIdAPI) workspace-level service.
* Added `Database` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Added `ReadReplicaCount` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore).
* Added `PageSize` field for [oauth2.ListServicePrincipalSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsRequest).
* Added `ProjectedRemainingTaskTotalTimeMs`, `RemainingTaskCount`, `RunnableTasks` and `WorkToBeDone` fields for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CreateCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsResponse).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CredentialInfo).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.GetCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsResponse).
* Added `IsDefaultForProvider` and `Name` fields for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).
* Added `Databricks` enum value for [catalog.SystemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemType).
* Added `DriverDnsResolutionFailure` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Confluence` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `DeltaIcebergTable` enum value for [sharing.TableInternalAttributesSharedTableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#TableInternalAttributesSharedTableType).
* [Breaking] Changed `Delete` method for [w.TableConstraints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableConstraintsAPI) workspace-level service to return `any`.
* [Breaking] Changed `Lifetime` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.DeleteServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#DeleteServicePrincipalSecretRequest) to type `string`.
* [Breaking] Changed `ServicePrincipalId` field for [oauth2.ListServicePrincipalSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsRequest) to type `string`.
* [Breaking] Changed `ExpireTime` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo) to type `string`.
* [Breaking] Changed `Calls` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit) to no longer be required.
* Changed `Calls` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit) to no longer be required.


## Release v0.74.0 (2025-07-03)

### API Changes
* Added [w.ExternalLineage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLineageAPI) workspace-level service and [w.ExternalMetadata](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalMetadataAPI) workspace-level service.
* Added [w.MaterializedFeatures](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#MaterializedFeaturesAPI) workspace-level service.
* Added `DeleteConversation`, `ListConversations` and `TrashSpace` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `CreateDatabaseInstanceRole`, `DeleteDatabaseInstanceRole`, `GetDatabaseInstanceRole` and `ListDatabaseInstanceRoles` methods for [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added `Connection` and `Credential` fields for [catalog.Dependency](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Dependency).
* Added `Rely` field for [catalog.ForeignKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ForeignKeyConstraint).
* Added `Rely` field for [catalog.PrimaryKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PrimaryKeyConstraint).
* Added `SecurableKindManifest` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
* Added `SecurableKindManifest` field for [catalog.TableSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableSummary).
* Added `ChildInstanceRefs`, `EffectiveEnableReadableSecondaries`, `EffectiveNodeCount`, `EffectiveRetentionWindowInDays`, `EnableReadableSecondaries`, `NodeCount`, `ParentInstanceRef`, `ReadOnlyDns` and `RetentionWindowInDays` fields for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `Claims` field for [database.GenerateDatabaseCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#GenerateDatabaseCredentialRequest).
* Added `LastSync` field for [database.SyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableStatus).
* Added `Activity` field for [ml.DeleteTransitionRequestResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestResponse).
* Added `MaxResults` field for [ml.ListWebhooksRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListWebhooksRequest).
* Added `Body` and `StatusCode` fields for [ml.TestRegistryWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TestRegistryWebhookResponse).
* Added `ModelVersionDatabricks` field for [ml.TransitionStageResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionStageResponse).
* Added `RegisteredModel` field for [ml.UpdateModelResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateModelResponse).
* Added `ModelVersion` field for [ml.UpdateModelVersionResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateModelVersionResponse).
* Added `Webhook` field for [ml.UpdateWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateWebhookResponse).
* Added `RunAs` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
* Added `Principal` field for [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit).
* Added `Description` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `ServedEntityName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route).
* Added `AnyStaticCredential` enum value for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Added `DatabricksRowStoreFormat`, `DeltaUniformHudi`, `DeltaUniformIceberg`, `Hive`, `Iceberg`, `MongodbFormat`, `OracleFormat`, `SalesforceDataCloudFormat` and `TeradataFormat` enum values for [catalog.DataSourceFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DataSourceFormat).
* Added `MetricView` enum value for [catalog.TableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableType).
* Added `SecurityAgentsFailedInitialVerification` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `CanCreateRegisteredModel` enum value for [ml.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PermissionLevel).
* Added `Bigquery` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `AppendOnly` enum value for [pipelines.TableSpecificConfigScdType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfigScdType).
* Added `ServicePrincipal` and `UserGroup` enum values for [serving.AiGatewayRateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimitKey).
* [Breaking] Changed `DeleteTransitionRequest`, `UpdateModel`, `UpdateModelVersion` and `UpdateWebhook` methods for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service return type to become non-empty.
* [Breaking] Changed `DeleteWebhook` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service with new required argument order.
* [Breaking] Changed `List` method for [w.AlertsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsLegacyAPI) workspace-level service . New request type is `any`.
* [Breaking] Changed `Update` method for [w.DashboardWidgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardWidgetsAPI) workspace-level service . New request type is [sql.UpdateWidgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateWidgetRequest).
* [Breaking] Changed `List` method for [w.DataSources](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DataSourcesAPI) workspace-level service . New request type is `any`.
* [Breaking] Changed `Create` method for [w.QueryVisualizationsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsLegacyAPI) workspace-level service with new required argument order.
* [Breaking] Changed `FromStage` and `ToStage` fields for [ml.Activity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Activity) to type `string`.
* [Breaking] Changed `Stage` field for [ml.ApproveTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ApproveTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.CreateTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.DeleteTransitionRequestRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestRequest) to type `string`.
* [Breaking] Changed `Id` field for [ml.DeleteWebhookRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteWebhookRequest) to be required.
* Changed `Capacity` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore) to be required.
* [Breaking] Changed `Capacity` field for [ml.OnlineStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#OnlineStore) to be required.
* [Breaking] Changed `OnlineTableName` field for [ml.PublishSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#PublishSpec) to be required.
* [Breaking] Changed `Stage` field for [ml.RejectTransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RejectTransitionRequest) to type `string`.
* [Breaking] Changed `Stage` field for [ml.TransitionModelVersionStageDatabricks](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionModelVersionStageDatabricks) to type `string`.
* [Breaking] Changed `ToStage` field for [ml.TransitionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionRequest) to type `string`.
* Changed `ServedModelName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route) to no longer be required.
* [Breaking] Changed `ServedModelName` field for [serving.Route](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Route) to no longer be required.
* Changed pagination for [TablesAPI.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI.List).
* Changed pagination for [TablesAPI.ListSummaries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI.ListSummaries).
* [Breaking] Removed `GenerateDownloadFullQueryResult` and `GetDownloadFullQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* [Breaking] Removed `IncludeDeltaMetadata` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
* [Breaking] Removed `Webhook` field for [ml.TestRegistryWebhookResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TestRegistryWebhookResponse).
* [Breaking] Removed `ModelVersion` field for [ml.TransitionStageResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#TransitionStageResponse).
* [Breaking] Removed `UnknownCatalogType` enum value for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* [Breaking] Removed `HiveCustom` and `HiveSerde` enum values for [catalog.DataSourceFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DataSourceFormat).
* [Breaking] Removed `UnknownSecurableType` enum value for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Removed `Archived`, `None`, `Production` and `Staging` enum values for [ml.DeleteTransitionRequestStage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteTransitionRequestStage).
* [Breaking] Removed `Archived`, `None`, `Production` and `Staging` enum values for [ml.Stage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Stage).


## Release v0.73.1 (2025-06-20)

### Bug Fixes
* Revert custom serialization which introduced a bug 
  impacting nested query paramters in v0.73.0


## Release v0.73.0 (2025-06-19)

### API Changes
* Added [w.AiBuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#AiBuilderAPI) workspace-level service.
* Added [w.FeatureStore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#FeatureStoreAPI) workspace-level service.
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Added `RemoteDiskThroughput` and `TotalInitialRemoteDiskSize` fields for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
* Added `ExpirationTime` field for [database.DatabaseCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseCredential).
* Added `EffectiveStopped` field for [database.DatabaseInstance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseInstance).
* Added `ExistingPipelineId` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* Added `PipelineId` field for [database.SyncedTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableStatus).
* Added `DbtPlatformOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DbtPlatformTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DbtPlatformTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DbtPlatformTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `Environment` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `Environment` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `Environment` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `Description` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `Description` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* Added `R` enum value for [compute.Language](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Language).
* Added `Cancelled`, `Error`, `Queued`, `Running`, `Starting` and `Success` enum values for [jobs.DbtPlatformRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtPlatformRunStatus).
* Added `Continuous` and `ContinuousRestart` enum values for [jobs.TriggerType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerType).
* [Breaking] Changed `Status` field for [jobs.DbtCloudJobRunStep](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtCloudJobRunStep) to type [jobs.DbtPlatformRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtPlatformRunStatus).
* [Breaking] Removed [w.CustomLlms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#CustomLlmsAPI) workspace-level service.
* [Breaking] Removed `TableServingUrl` field for [database.DatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseTable).
* [Breaking] Removed `TableServingUrl` field for [database.SyncedDatabaseTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedDatabaseTable).
* [Breaking] Removed `PipelineId` field for [database.SyncedTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#SyncedTableSpec).
* [Breaking] Removed `Cancelled`, `Error`, `Queued`, `Running`, `Starting` and `Success` enum values for [jobs.DbtCloudRunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtCloudRunStatus).


## Release v0.72.0 (2025-06-05)

### New Features and Improvements

- Enums now have a `Values` function that returns the list of all possible
  enum values ([PR #1228](https://github.com/databricks/databricks-sdk-go/pull/1228))

### API Changes
* Added [aibuilder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder) package, [database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database) package and [qualitymonitorv2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2) package.
* Added [w.CustomLlms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/aibuilder#CustomLlmsAPI) workspace-level service.
* Added [w.DashboardEmailSubscriptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DashboardEmailSubscriptionsAPI) workspace-level service and [w.SqlResultsDownload](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SqlResultsDownloadAPI) workspace-level service.
* Added [w.Database](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/database#DatabaseAPI) workspace-level service.
* Added [w.QualityMonitorV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/qualitymonitorv2#QualityMonitorV2API) workspace-level service.
* Added `UpdatePrivateEndpointRule` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* Added `ListSpaces` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `PageToken` field for [billing.ListLogDeliveryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListLogDeliveryRequest).
* Added `NextPageToken` field for [billing.WrappedLogDeliveryConfigurations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedLogDeliveryConfigurations).
* Added `NextPageToken` field for [catalog.EffectivePermissionsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePermissionsList).
* Added `MaxResults` and `PageToken` fields for [catalog.GetEffectiveRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetEffectiveRequest).
* Added `MaxResults` and `PageToken` fields for [catalog.GetGrantRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetGrantRequest).
* Added `NextPageToken` field for [catalog.ListMetastoresResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMetastoresResponse).
* Added `CleanRoomName` field for [cleanrooms.CleanRoomAsset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAsset).
* [Breaking] Added `Name` field for [cleanrooms.DeleteCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#DeleteCleanRoomAssetRequest).
* [Breaking] Added `Name` field for [cleanrooms.GetCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#GetCleanRoomAssetRequest).
* Added `TriggerState` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
* Added `TriggerState` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
* Added `DbtCloudOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DbtCloudTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DbtCloudTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DbtCloudTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `Tags` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `Tags` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `Tags` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Added `MaxProvisionedConcurrency` and `MinProvisionedConcurrency` fields for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Added `EndpointService` and `ResourceNames` fields for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule).
* Added `AwsPrivateEndpointRules` field for [settings.NccEgressTargetRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressTargetRules).
* Added `TaskTimeOverTimeRange` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
* Added `DeltasharingCatalog`, `ForeignCatalog`, `InternalCatalog`, `ManagedCatalog`, `ManagedOnlineCatalog`, `SystemCatalog` and `UnknownCatalogType` enum values for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Ga4RawData`, `PowerBi`, `Salesforce`, `SalesforceDataCloud`, `Servicenow`, `UnknownConnectionType` and `WorkdayRaas` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Added `OauthAccessToken`, `OauthM2m`, `OauthRefreshToken`, `OauthResourceOwnerPassword`, `OauthU2m`, `OauthU2mMapping`, `OidcToken`, `PemPrivateKey`, `ServiceCredential` and `UnknownCredentialType` enum values for [catalog.CredentialType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialType).
* Added `Internal` and `InternalAndExternal` enum values for [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* Added `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Added `ClusterMigrated` enum value for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
* Added `DriverUnhealthy` enum value for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Teradata` enum value for [pipelines.IngestionSourceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionSourceType).
* Added `OidcFederation` enum value for [sharing.AuthenticationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#AuthenticationType).
* [Breaking] Changed `Create` method for [a.LogDelivery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryAPI) account-level service with new required argument order.
* [Breaking] Changed `Get` method for [a.LogDelivery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryAPI) account-level service to return [billing.GetLogDeliveryConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetLogDeliveryConfigurationResponse).
* [Breaking] Changed `CreatePrivateEndpointRule`, `DeletePrivateEndpointRule` and `GetPrivateEndpointRule` methods for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service to return [settings.NccPrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccPrivateEndpointRule).
* [Breaking] Changed `ListPrivateEndpointRules` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service to return [settings.ListPrivateEndpointRulesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPrivateEndpointRulesResponse).
* [Breaking] Changed `Delete` and `Get` methods for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `Delete` and `Get` methods for [w.CleanRoomAssets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetsAPI) workspace-level service . Method path has changed.
* [Breaking] Changed `Get` method for [w.Grants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GrantsAPI) workspace-level service to return [catalog.GetPermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetPermissionsResponse).
* [Breaking] Changed `Update` method for [w.Grants](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GrantsAPI) workspace-level service to return [catalog.UpdatePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissionsResponse).
* [Breaking] Changed `List` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service to require request of [catalog.ListMetastoresRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListMetastoresRequest).
* Changed `AccountId`, `CredentialsId`, `LogType`, `OutputFormat` and `StorageConfigurationId` fields for [billing.LogDeliveryConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryConfiguration) to be required.
* Changed `Message` and `Status` fields for [billing.LogDeliveryStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#LogDeliveryStatus) to be required.
* [Breaking] Changed `LogDeliveryConfiguration` field for [billing.WrappedCreateLogDeliveryConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedCreateLogDeliveryConfiguration) to be required.
* [Breaking] Changed `SecurableType` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo) to type [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Changed `SecurableType` field for [catalog.GetEffectiveRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetEffectiveRequest) to type `string`.
* [Breaking] Changed `SecurableType` field for [catalog.GetGrantRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetGrantRequest) to type `string`.
* [Breaking] Changed `DeltaSharingScope` field for [catalog.GetMetastoreSummaryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMetastoreSummaryResponse) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `DeltaSharingScope` field for [catalog.MetastoreInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoreInfo) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `CatalogType` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to type [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* [Breaking] Changed `DeltaSharingScope` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore) to type [catalog.DeltaSharingScopeEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeltaSharingScopeEnum).
* [Breaking] Changed `SecurableType` field for [catalog.UpdatePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePermissions) to type `string`.
* Changed `ResourceId` field for [settings.CreatePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRule) to no longer be required.
* [Breaking] Changed pagination for [NetworkConnectivityAPI.ListPrivateEndpointRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI.ListPrivateEndpointRules).
* [Breaking] Removed [w.DatabaseInstances](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabaseInstancesAPI) workspace-level service.
* [Breaking] Removed [w.QueryExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryExecutionAPI) workspace-level service.
* [Breaking] Removed `UpdateNccAzurePrivateEndpointRulePublic` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* [Breaking] Removed `GetCredentialsForTraceDataDownload`, `GetCredentialsForTraceDataUpload` and `ListLoggedModelArtifacts` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* [Breaking] Removed `GetPublishedDashboardEmbedded` method for [w.LakeviewEmbedded](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewEmbeddedAPI) workspace-level service.
* [Breaking] Removed `AssetFullName` field for [cleanrooms.DeleteCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#DeleteCleanRoomAssetRequest).
* [Breaking] Removed `AssetFullName` field for [cleanrooms.GetCleanRoomAssetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#GetCleanRoomAssetRequest).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.GetMetastoreSummaryResponseDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMetastoreSummaryResponseDeltaSharingScope).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.MetastoreInfoDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoreInfoDeltaSharingScope).
* [Breaking] Removed `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Removed `Internal` and `InternalAndExternal` enum values for [catalog.UpdateMetastoreDeltaSharingScope](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastoreDeltaSharingScope).


## Release v0.71.0 (2025-05-26)

### Bug Fixes

- Properly cache tokens when using the `metadata-service` auth type ([#1225](https://github.com/databricks/databricks-sdk-go/pull/1225)).


## Release v0.70.0 (2025-05-21)

### Bug Fixes

- Stop swallowing authentication errors when a non-default auth type is 
  explicitly set ([#1223](https://github.com/databricks/databricks-sdk-go/pull/1223))

### API Changes
* Added [a.NetworkPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkPoliciesAPI) account-level service and [a.WorkspaceNetworkConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#WorkspaceNetworkConfigurationAPI) account-level service.
* Added [w.RecipientFederationPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientFederationPoliciesAPI) workspace-level service.
* Added `CreateLoggedModel`, `DeleteLoggedModel`, `DeleteLoggedModelTag`, `FinalizeLoggedModel`, `GetLoggedModel`, `ListLoggedModelArtifacts`, `LogLoggedModelParams`, `LogOutputs`, `SearchLoggedModels` and `SetLoggedModelTags` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* Added `UcSecurable` field for [apps.AppResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppResource).
* Added `TimeseriesColumns` field for [catalog.PrimaryKeyConstraint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PrimaryKeyConstraint).
* Added `ReviewState`, `Reviews` and `RunnerCollaboratorAliases` fields for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook).
* Added `NotebookEtag` and `NotebookUpdatedAt` fields for [cleanrooms.CleanRoomNotebookTaskRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomNotebookTaskRun).
* Added `RootPath` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `RootPath` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `RootPath` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `MaterializationNamespace` field for [sharing.Table](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Table).
* Added `OmitPermissionsList` field for [sharing.UpdateSharePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissions).
* Added `AutoResolveDisplayName` field for [sql.UpdateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequest).
* Added `AutoResolveDisplayName` field for [sql.UpdateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequest).
* Added `InternalCatalog`, `ManagedOnlineCatalog` and `UnknownCatalogType` enum values for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Catalog`, `CleanRoom`, `Connection`, `Credential`, `ExternalLocation`, `ExternalMetadata`, `Function`, `Metastore`, `Pipeline`, `Provider`, `Recipient`, `Schema`, `Share`, `StagingTable`, `StorageCredential`, `Table`, `UnknownSecurableType` and `Volume` enum values for [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* [Breaking] Changed `SecurableType` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo) to type [catalog.SecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableType).
* Changed `Etag` and `Name` fields for [iam.RuleSetResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#RuleSetResponse) to be required.
* Added `EnableFileEvents` and `FileEventQueue` fields for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* Added `EnableFileEvents` and `FileEventQueue` fields for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* Added `EnableFileEvents` and `FileEventQueue` fields for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
* Added `PolicyId` and `ServicePrincipalId` fields for [oauth2.FederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#FederationPolicy).
* [Breaking] Removed `AccessPoint` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
* [Breaking] Removed `AccessPoint` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
* [Breaking] Removed `AccessPoint` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).


## Release v0.69.0 (2025-05-12)

### New Features and Improvements

- Add support to authenticate with Account-wide token federation from the 
  following auth methods: `env-oidc`, `file-oidc`, and `github-oidc`.  

### API Changes
* Added [a.LlmProxyPartnerPoweredAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredAccountAPI) account-level service, [a.LlmProxyPartnerPoweredEnforce](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredEnforceAPI) account-level service and [w.LlmProxyPartnerPoweredWorkspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#LlmProxyPartnerPoweredWorkspaceAPI) workspace-level service.
* Added [w.DatabaseInstances](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabaseInstancesAPI) workspace-level service.
* Added `CreateProvisionedThroughputEndpoint` and `UpdateProvisionedThroughputEndpointConfig` methods for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
* Added `CatalogName` field for [catalog.EnableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableRequest).
* Added `SourceType` field for [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
* Added `Glob` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
* Added `ProvisionedModelUnits` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
* Added `ProvisionedModelUnits` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
* Added `ProvisionedModelUnits` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
* Added `ProvisionedModelUnits` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
* Added `DescribeQueryInvalidSqlError`, `DescribeQueryTimeout`, `DescribeQueryUnexpectedFailure`, `InvalidChatCompletionArgumentsJsonException`, `InvalidSqlMultipleDatasetReferencesException`, `InvalidSqlMultipleStatementsException` and `InvalidSqlUnknownTableException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `CanCreate` and `CanMonitorOnly` enum values for [iam.PermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionLevel).
* Added `SuccessWithFailures` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* Added `InfrastructureMaintenance` enum value for [pipelines.StartUpdateCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdateCause).
* Added `InfrastructureMaintenance` enum value for [pipelines.UpdateInfoCause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfoCause).
* [Breaking] Changed `CreateAlert` and `UpdateAlert` methods for [w.AlertsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API) workspace-level service with new required argument order.
* [Breaking] Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.SetObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#SetObjectPermissions).
* [Breaking] Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service . New request type is [iam.UpdateObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UpdateObjectPermissions).
* [Breaking] Changed `Get` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.GetCatalogWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogWorkspaceBindingsResponse).
* [Breaking] Changed `GetBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.GetWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetWorkspaceBindingsResponse).
* [Breaking] Changed `Update` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.UpdateCatalogWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalogWorkspaceBindingsResponse).
* [Breaking] Changed `UpdateBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service to return [catalog.UpdateWorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsResponse).
* [Breaking] Changed `SecurableType` field for [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest) to type `string`.
* Changed `Schema` and `State` fields for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to be required.
* [Breaking] Changed `State` field for [catalog.SystemSchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfo) to type `string`.
* [Breaking] Changed `SecurableType` field for [catalog.UpdateWorkspaceBindingsParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsParameters) to type `string`.
* [Breaking] Changed `WorkspaceId` field for [catalog.WorkspaceBinding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBinding) to be required.
* [Breaking] Changed `GpuNodePoolId` field for [jobs.ComputeConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeConfig) to no longer be required.
* Changed `GpuNodePoolId` field for [jobs.ComputeConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ComputeConfig) to no longer be required.
* [Breaking] Changed `Alert` field for [sql.CreateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertV2Request) to be required.
* [Breaking] Changed `Alert` field for [sql.UpdateAlertV2Request](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertV2Request) to be required.
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* [Breaking] Removed `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* [Breaking] Removed `Catalog`, `Credential`, `ExternalLocation` and `StorageCredential` enum values for [catalog.GetBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsSecurableType).
* [Breaking] Removed `Available`, `DisableInitialized`, `EnableCompleted`, `EnableInitialized` and `Unavailable` enum values for [catalog.SystemSchemaInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SystemSchemaInfoState).
* [Breaking] Removed `Catalog`, `Credential`, `ExternalLocation` and `StorageCredential` enum values for [catalog.UpdateBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateBindingsSecurableType).


## Release v0.68.0 (2025-05-06)

### New Features and Improvements

- Add support for OIDC ID token authentication using a file
  ([PR #1213](https://github.com/databricks/databricks-sdk-go/pull/1213)).
- Add support for OIDC ID token authentication using an environment variable
  ([PR #1215](https://github.com/databricks/databricks-sdk-go/pull/1215)).


## Release v0.67.0 (2025-05-02)

### Bug Fixes
* Fixed the deserialization of responses in VectorSearchAPI's `QueryIndex()` method ([#1214](https://github.com/databricks/databricks-sdk-py/pull/1214)).

### API Changes
* Added `FutureFeatureDataPath` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `ExcludeColumns` and `IncludeColumns` fields for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
* Added `NetworkCheckControlPlaneFailure`, `NetworkCheckDnsServerFailure`, `NetworkCheckMetadataEndpointFailure`, `NetworkCheckMultipleComponentsFailure`, `NetworkCheckNicFailure`, `NetworkCheckStorageFailure` and `SecretPermissionDenied` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* [Breaking] Changed [vectorsearch.ListValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ListValue) to.
* [Breaking] Changed `PipelineId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline) to be required.
* Changed `ConnectionName`, `GatewayStorageCatalog` and `GatewayStorageSchema` fields for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition) to be required.
* [Breaking] Changed `ConnectionName`, `GatewayStorageCatalog` and `GatewayStorageSchema` fields for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition) to be required.
* [Breaking] Changed `Kind` field for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment) to be required.
* Changed `Kind` field for [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceUrl` fields for [pipelines.ReportSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ReportSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceUrl` fields for [pipelines.ReportSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ReportSpec) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceSchema` fields for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceSchema` fields for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec) to be required.
* Changed `DestinationCatalog`, `DestinationSchema` and `SourceTable` fields for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to be required.
* [Breaking] Changed `DestinationCatalog`, `DestinationSchema` and `SourceTable` fields for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec) to be required.
* [Breaking] Changed pagination for [AlertsV2API.ListAlerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API.ListAlerts).
* [Breaking] Changed waiter for [GenieAPI.CreateMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI.CreateMessage).


## Release v0.66.0 (2025-04-30)

### Bug Fixes
* Tolerate trailing slashes in hostnames in `Config`.

### API Changes
* Added [w.AlertsV2](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsV2API) workspace-level service.
* Added `UpdateNccAzurePrivateEndpointRulePublic` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
* Added `CreatedAt`, `CreatedBy` and `MetastoreId` fields for [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist).
* [Breaking] Added `NetworkConnectivityConfig` field for [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest).
* [Breaking] Added `PrivateEndpointRule` field for [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest).
* Added `DomainNames` field for [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule).
* Added `AutoResolveDisplayName` field for [sql.CreateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequest).
* Added `AutoResolveDisplayName` field for [sql.CreateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryRequest).
* Added `CreateCleanRoom`, `ExecuteCleanRoomTask` and `ModifyCleanRoom` enum values for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Added `DnsResolutionError` and `GcpDeniedByOrgPolicy` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Expired` enum value for [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState).
* [Breaking] Changed `CreateNetworkConnectivityConfiguration` and `CreatePrivateEndpointRule` methods for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service with new required argument order.
* [Breaking] Changed `WorkloadSize` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to type `string`.
* [Breaking] Changed `GroupId` field for [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule) to type `string`.
* [Breaking] Changed `TargetServices` field for [settings.NccAzureServiceEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzureServiceEndpointRule) to type [settings.EgressResourceTypeList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressResourceTypeList).
* [Breaking] Removed `Name` and `Region` fields for [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest).
* [Breaking] Removed `GroupId` and `ResourceId` fields for [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest).
* [Breaking] Removed `Large`, `Medium` and `Small` enum values for [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
* [Breaking] Removed `Blob`, `Dfs`, `MysqlServer` and `SqlServer` enum values for [settings.NccAzurePrivateEndpointRuleGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleGroupId).
* [Breaking] Field `AppDeployment` of `CreateAppDeploymentRequest` is changed from `*AppDeployment` to `AppDeployment`.
* [Breaking] Field `App` of `CreateAppRequest` is changed from `*App` to `App`.
* [Breaking] Field `App` of `UpdateAppRequest` is changed from `*App` to `App`.
* [Breaking] Field `BudgetPolicy` of `UpdateBudgetPolicyRequest` is changed from `*BudgetPolicy` to `BudgetPolicy`.
* [Breaking] Field `OnlineTable` of `CreateOnlineTableRequest` is changed from `*OnlineTable` to `OnlineTable`.
* [Breaking] Field `CleanRoomAsset` of `CreateCleanRoomAssetRequest` is changed from `*CleanRoomAsset` to `CleanRoomAsset`.
* [Breaking] Field `CleanRoom` of `CreateCleanRoomRequest` is changed from `*CleanRoom` to `CleanRoom`.
* [Breaking] Field `CleanRoomAsset` of `UpdateCleanRoomAssetRequest` is changed from `*CleanRoomAsset` to `CleanRoomAsset`.
* [Breaking] Field `Dashboard` of `CreateDashboardRequest` is changed from `*Dashboard` to `Dashboard`.
* [Breaking] Field `Schedule` of `CreateScheduleRequest` is changed from `*Schedule` to `Schedule`.
* [Breaking] Field `Subscription` of `CreateSubscriptionRequest` is changed from `*Subscription` to `Subscription`.
* [Breaking] Field `Dashboard` of `UpdateDashboardRequest` is changed from `*Dashboard` to `Dashboard`.


## Release v0.65.0 (2025-04-29)

### New Features and Improvements
* Introduce support for Databricks Workload Identity Federation in GitHub workflows ([1177](https://github.com/databricks/databricks-sdk-go/pull/1177)).
  See README.md for instructions.
* [Breaking] Users running their worklows in GitHub Actions, which use Cloud native authentication and also have a `DATABRICKS_CLIENT_ID` and `DATABRICKS_HOST` 
  environment variables set may see their authentication start failing due to the order in which the SDK tries different authentication methods.
  In such case, the `DATABRICKS_AUTH_TYPE` environment variable must be set to match the previously used authentication method.


## Release v0.64.0 (2025-04-24)

### New Features and Improvements
* Enabled asynchronous token refreshes by default ([#1208](https://github.com/databricks/databricks-sdk-go/pull/1208)).

### API Changes
* Added `UpdateEndpointBudgetPolicy` and `UpdateEndpointCustomTags` methods for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service.
* Added `NodeTypeFlexibility` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
* Added `PageSize` and `PageToken` fields for [compute.GetEvents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetEvents).
* Added `NextPageToken` and `PrevPageToken` fields for [compute.GetEventsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetEventsResponse).
* Added `NodeTypeFlexibility` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
* Added `NodeTypeFlexibility` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
* Added `EffectivePerformanceTarget` field for [jobs.RepairHistoryItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairHistoryItem).
* Added `PerformanceTarget` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
* Added `BudgetPolicyId` field for [vectorsearch.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateEndpoint).
* Added `CustomTags` and `EffectiveBudgetPolicyId` fields for [vectorsearch.EndpointInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EndpointInfo).
* Added `Disabled` enum value for [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode).
* [Breaking] Changed `CreateIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service to return [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
* [Breaking] Changed `DeleteDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service . HTTP method/verb has changed.
* [Breaking] Changed `DeleteDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service with new required argument order.
* [Breaking] Changed `DataArray` field for [vectorsearch.ResultData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ResultData) to type [vectorsearch.ListValueList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ListValueList).
* [Breaking] Changed waiter for [VectorSearchEndpointsAPI.CreateEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI.CreateEndpoint).
* [Breaking] Removed `NullValue` field for [vectorsearch.Value](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#Value).


## Release v0.63.0 (2025-04-14)

### API Changes
* Added [w.EnableExportNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableExportNotebookAPI) workspace-level service, [w.EnableNotebookTableClipboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableNotebookTableClipboardAPI) workspace-level service and [w.EnableResultsDownloading](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableResultsDownloadingAPI) workspace-level service.
* Added `GetCredentialsForTraceDataDownload` and `GetCredentialsForTraceDataUpload` methods for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
* Added `GetDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `GetPublishedDashboardTokenInfo` method for [w.LakeviewEmbedded](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewEmbeddedAPI) workspace-level service.
* Added `BindingWorkspaceIds` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy).
* Added `DownloadId` field for [dashboards.GenieGenerateDownloadFullQueryResultResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGenerateDownloadFullQueryResultResponse).
* Added `DashboardOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `DashboardTask` and `PowerBiTask` fields for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `IncludeFeatures` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `Models` field for [ml.LogInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogInputs).
* Added `DatasetDigest`, `DatasetName` and `ModelId` fields for [ml.LogMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogMetric).
* Added `DatasetDigest`, `DatasetName`, `ModelId` and `RunId` fields for [ml.Metric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Metric).
* Added `ModelInputs` field for [ml.RunInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RunInputs).
* Added `ClientApplication` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `Geography` and `Geometry` enum values for [catalog.ColumnTypeName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ColumnTypeName).
* Added `AllocationTimeoutNoHealthyAndWarmedUpClusters`, `DockerContainerCreationException`, `DockerImageTooLargeForInstanceException` and `DockerInvalidOsException` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `Standard` enum value for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* Added `CanView` enum value for [sql.WarehousePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionLevel).
* [Breaking] Changed `GenerateDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service . Method path has changed.
* [Breaking] Changed waiter for [CommandExecutionAPI.Create](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CommandExecutionAPI.Create).
* [Breaking] Changed waiter for [CommandExecutionAPI.Execute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CommandExecutionAPI.Execute).
* [Breaking] Removed `Error`, `Status` and `TransientStatementId` fields for [dashboards.GenieGenerateDownloadFullQueryResultResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGenerateDownloadFullQueryResultResponse).
* [Breaking] Removed `Balanced` and `CostOptimized` enum values for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* [Breaking] Removed [PipelinesAPI.WaitGetPipelineRunning](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI.WaitGetPipelineRunning) method.


## Release v0.62.0 (2025-04-10)

### Internal Changes
* Remove `TestAccDashboards` as legacy dashboard creation on has been discontinued ([#1205](https://github.com/databricks/databricks-sdk-go/pull/1205)).


## Release v0.61.0 (2025-03-26)

### New Features and Improvements

* Support user-to-machine authentication in the SDK ([#1108](https://github.com/databricks/databricks-sdk-go/pull/1108)).
- Instances of `ApiClient` now share the same connection pool by default ([PR #1190](https://github.com/databricks/databricks-sdk-go/pull/1190)).

### Internal Changes

- Stop recommending users to report an issue when the SDK encounters an unknown
  error ([PR #1189](https://github.com/databricks/databricks-sdk-go/pull/1189)).

### API Changes
* Added `GenerateDownloadFullQueryResult` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `EffectiveUserApiScopes`, `Oauth2AppClientId`, `Oauth2AppIntegrationId` and `UserApiScopes` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Added `Abfss`, `Dbfs`, `ErrorMessage`, `ExecutionDurationSeconds`, `File`, `Gcs`, `S3`, `Status`, `Volumes` and `Workspace` fields for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* [Breaking] Added `ForecastGranularity` field for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* Added `JwksUri` field for [oauth2.OidcFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OidcFederationPolicy).
* Added `EventLog` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
* Added `EventLog` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
* Added `EventLog` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
* Added `FallbackConfig` field for [serving.AiGatewayConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayConfig).
* Added `CustomProviderConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
* Added `FallbackConfig` field for [serving.PutAiGatewayRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayRequest).
* Added `FallbackConfig` field for [serving.PutAiGatewayResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayResponse).
* Added `Aliases`, `Comment`, `DataType`, `DependencyList`, `FullDataType`, `Id`, `InputParams`, `Name`, `Properties`, `RoutineDefinition`, `Schema`, `SecurableKind`, `Share`, `ShareId`, `StorageLocation` and `Tags` fields for [sharing.DeltaSharingFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeltaSharingFunction).
* Added `QuerySource` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
* Added `ForeignCatalog` enum value for [catalog.CatalogType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogType).
* Added `Browse` enum value for [catalog.Privilege](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#Privilege).
* Added `AccessTokenFailure`, `AllocationTimeout`, `AllocationTimeoutNodeDaemonNotReady`, `AllocationTimeoutNoHealthyClusters`, `AllocationTimeoutNoMatchedClusters`, `AllocationTimeoutNoReadyClusters`, `AllocationTimeoutNoUnallocatedClusters`, `AllocationTimeoutNoWarmedUpClusters`, `AwsInaccessibleKmsKeyFailure`, `AwsInstanceProfileUpdateFailure`, `AwsInvalidKeyPair`, `AwsInvalidKmsKeyState`, `AwsResourceQuotaExceeded`, `AzurePackedDeploymentPartialFailure`, `BootstrapTimeoutDueToMisconfig`, `BudgetPolicyLimitEnforcementActivated`, `BudgetPolicyResolutionFailure`, `CloudAccountSetupFailure`, `CloudOperationCancelled`, `CloudProviderInstanceNotLaunched`, `CloudProviderLaunchFailureDueToMisconfig`, `CloudProviderResourceStockoutDueToMisconfig`, `ClusterOperationThrottled`, `ClusterOperationTimeout`, `ControlPlaneRequestFailureDueToMisconfig`, `DataAccessConfigChanged`, `DisasterRecoveryReplication`, `DriverEviction`, `DriverLaunchTimeout`, `DriverNodeUnreachable`, `DriverOutOfDisk`, `DriverOutOfMemory`, `DriverPodCreationFailure`, `DriverUnexpectedFailure`, `DynamicSparkConfSizeExceeded`, `EosSparkImage`, `ExecutorPodUnscheduled`, `GcpApiRateQuotaExceeded`, `GcpForbidden`, `GcpIamTimeout`, `GcpInaccessibleKmsKeyFailure`, `GcpInsufficientCapacity`, `GcpIpSpaceExhausted`, `GcpKmsKeyPermissionDenied`, `GcpNotFound`, `GcpResourceQuotaExceeded`, `GcpServiceAccountAccessDenied`, `GcpServiceAccountNotFound`, `GcpSubnetNotReady`, `GcpTrustedImageProjectsViolated`, `GkeBasedClusterTermination`, `InitContainerNotFinished`, `InstancePoolMaxCapacityReached`, `InstancePoolNotFound`, `InstanceUnreachableDueToMisconfig`, `InternalCapacityFailure`, `InvalidAwsParameter`, `InvalidInstancePlacementProtocol`, `InvalidWorkerImageFailure`, `InPenaltyBox`, `LazyAllocationTimeout`, `MaintenanceMode`, `NetvisorSetupTimeout`, `NoMatchedK8s`, `NoMatchedK8sTestingTag`, `PodAssignmentFailure`, `PodSchedulingFailure`, `ResourceUsageBlocked`, `SecretCreationFailure`, `ServerlessLongRunningTerminated`, `SparkImageDownloadThrottled`, `SparkImageNotFound`, `SshBootstrapFailure`, `StorageDownloadFailureDueToMisconfig`, `StorageDownloadFailureSlow`, `StorageDownloadFailureThrottled`, `UnexpectedPodRecreation`, `UserInitiatedVmTermination` and `WorkspaceUpdate` enum values for [compute.TerminationReasonCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#TerminationReasonCode).
* Added `GeneratedSqlQueryTooLongException` and `MissingSqlQueryException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `Balanced` enum value for [jobs.PerformanceTarget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PerformanceTarget).
* Added `ListingResource` enum value for [marketplace.FileParentType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#FileParentType).
* Added `App` enum value for [marketplace.MarketplaceFileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#MarketplaceFileType).
* Added `Custom` enum value for [serving.ExternalModelProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelProvider).
* Added `ArclightMultiTenantAzureExchangeToken` and `ArclightMultiTenantAzureExchangeTokenWithUserDelegationKey` enum values for [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
* [Breaking] Changed `CreateExperiment` method for [w.Forecasting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ForecastingAPI) workspace-level service with new required argument order.
* Changed `InstanceTypeId` field for [compute.NodeInstanceType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeInstanceType) to be required.
* Changed `Category` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType) to be required.
* [Breaking] Changed `Functions` field for [sharing.ListProviderShareAssetsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderShareAssetsResponse) to type [sharing.DeltaSharingFunctionList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeltaSharingFunctionList).
* [Breaking] Removed `ExecutionDetails` and `Script` fields for [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails).
* [Breaking] Removed `SupportsElasticDisk` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType).
* [Breaking] Removed `DataGranularityQuantity` and `DataGranularityUnit` fields for [ml.CreateForecastingExperimentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateForecastingExperimentRequest).
* [Breaking] Removed `Aliases`, `Comment`, `DataType`, `DependencyList`, `FullDataType`, `Id`, `InputParams`, `Name`, `Properties`, `RoutineDefinition`, `Schema`, `SecurableKind`, `Share`, `ShareId`, `StorageLocation` and `Tags` fields for [sharing.Function](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#Function).


## Release v0.60.0 (2025-03-11)

### API Changes
* Added [w.Forecasting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ForecastingAPI) workspace-level service.
* Added `ExecuteMessageAttachmentQuery` and `GetMessageAttachmentQueryResult` methods for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `StatementId` field for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Added `BudgetPolicyId` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
* Added `BudgetPolicyId` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
* Added `BudgetPolicyId` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
* Added `CouldNotGetModelDeploymentsException` enum value for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).


## Release v0.59.0 (2025-03-03)

### Bug Fixes

* Fix unlikely issue due to conflicting error details in `APIError`.

### Internal Changes

* Update Jobs ListRuns API to support paginated responses ([#1151](https://github.com/databricks/databricks-sdk-go/pull/1151))
* Add `poll.SimpleError` to mock waiter objects returning errors  ([#1155](https://github.com/databricks/databricks-sdk-go/pull/1155))
* Refactor `APIError` to expose different types of error details ([#1153](https://github.com/databricks/databricks-sdk-go/pull/1153)). 
* Update Jobs ListJobs API to support paginated responses ([#1150](https://github.com/databricks/databricks-sdk-go/pull/1150))
* Introduce automated tagging ([#1148](https://github.com/databricks/databricks-sdk-go/pull/1148)).
* Update Jobs GetJob API to support paginated responses  ([#1133](https://github.com/databricks/databricks-sdk-go/pull/1133)).
* Update Jobs GetRun API to support paginated responses  ([#1132](https://github.com/databricks/databricks-sdk-go/pull/1132)).

### API Changes
* Added `GetSpace` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
* Added `ListProviderShareAssets` method for [w.Providers](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ProvidersAPI) workspace-level service.
* Added `BudgetPolicyId` and `EffectiveBudgetPolicyId` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
* Added `Policy` field for [billing.CreateBudgetPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetPolicyRequest).
* Added `DatabricksGcpServiceAccount` field for [catalog.ValidateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateCredentialRequest).
* Added `AttachmentId` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment).
* Added `ConversationId` field for [dashboards.GenieConversation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversation).
* Added `MessageId` field for [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage).
* Added `Description`, `Id`, `LastUpdatedTimestamp`, `Query`, `QueryResultMetadata` and `Title` fields for [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Added `GenAiComputeTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
* Added `GenAiComputeTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
* Added `GenAiComputeTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
* Added `RunName` field for [ml.CreateRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#CreateRun).
* Added `RunName` field for [ml.RunInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RunInfo).
* Added `RunName` field for [ml.UpdateRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#UpdateRun).
* Added `Lifetime` field for [oauth2.CreateServicePrincipalSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretRequest).
* Added `ExpireTime` field for [oauth2.CreateServicePrincipalSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateServicePrincipalSecretResponse).
* Added `ExpireTime` field for [oauth2.SecretInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#SecretInfo).
* Added `InstanceProfileArn` field for [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig).
* Added `Add`, `Principal` and `Remove` fields for [sharing.PermissionsChange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#PermissionsChange).
* Added `ColumnsToRerank` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
* Added `Oracle` and `Teradata` enum values for [catalog.ConnectionType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionType).
* Added `FunctionArgumentsInvalidTypeException` and `MessageCancelledWhileExecutingException` enum values for [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType).
* Added `Waiting` enum value for [jobs.RunLifecycleStateV2State](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunLifecycleStateV2State).
* Added `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* Added `OauthClientCredentials` enum value for [sharing.AuthenticationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#AuthenticationType).
* Added `Raw` enum value for [workspace.ExportFormat](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ExportFormat).
* [Breaking] Changed `GetByName` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentByNameResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentByNameResponse).
* [Breaking] Changed `LogInputs` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service with new required argument order.
* [Breaking] Changed `SharePermissions` and `UpdatePermissions` methods for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service return type to become non-empty.
* [Breaking] Changed `SharePermissions` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service to return [sharing.GetSharePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetSharePermissionsResponse).
* [Breaking] Changed `UpdatePermissions` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service to return [sharing.UpdateSharePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissionsResponse).
* Changed `PolicyId` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy) to no longer be required.
* [Breaking] Changed `PolicyId` field for [billing.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicy) to no longer be required.
* [Breaking] Changed `Partitions` field for [cleanrooms.CleanRoomAssetTableLocalDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetTableLocalDetails) to type [cleanrooms.PartitionList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#PartitionList).
* [Breaking] Changed `Query` field for [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment) to type [dashboards.GenieQueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieQueryAttachment).
* Changed `Digest`, `Name`, `Source` and `SourceType` fields for [ml.Dataset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Dataset) to be required.
* [Breaking] Changed `Digest`, `Name`, `Source` and `SourceType` fields for [ml.Dataset](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#Dataset) to be required.
* [Breaking] Changed `Dataset` field for [ml.DatasetInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DatasetInput) to be required.
* Changed `Dataset` field for [ml.DatasetInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DatasetInput) to be required.
* Changed `Key` and `Value` fields for [ml.InputTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#InputTag) to be required.
* [Breaking] Changed `Key` and `Value` fields for [ml.InputTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#InputTag) to be required.
* [Breaking] Changed `ViewType` field for [ml.ListExperimentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ListExperimentsRequest) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Changed `RunId` field for [ml.LogInputs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#LogInputs) to be required.
* [Breaking] Changed `ViewType` field for [ml.SearchExperiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchExperiments) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Changed `RunViewType` field for [ml.SearchRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchRuns) to type [ml.ViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ViewType).
* [Breaking] Removed `CustomTags` and `PolicyName` fields for [billing.CreateBudgetPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetPolicyRequest).
* [Breaking] Removed `CachedQuerySchema`, `Description`, `Id`, `InstructionId`, `InstructionTitle`, `LastUpdatedTimestamp`, `Query`, `StatementId` and `Title` fields for [dashboards.QueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryAttachment).
* [Breaking] Removed `MaxResults` and `PageToken` fields for [sharing.UpdateSharePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissions).
* [Breaking] Removed `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.SearchExperimentsViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchExperimentsViewType).
* [Breaking] Removed `ActiveOnly`, `All` and `DeletedOnly` enum values for [ml.SearchRunsRunViewType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#SearchRunsRunViewType).


## Release v0.58.1 (2025-02-12)

### Internal Changes

* Do not send ForceSendFields as query parameters.

## Release v0.58.0 (2025-02-11)

### New Features and Improvements

 * Enable async refreshes for OAuth tokens ([#1143](https://github.com/databricks/databricks-sdk-go/pull/1143)).


### Internal Changes

 * Add support for asynchronous data plane token refreshes ([#1142](https://github.com/databricks/databricks-sdk-go/pull/1142)).
 * Introduce new TokenSource interface that takes a `context.Context` ([#1141](https://github.com/databricks/databricks-sdk-go/pull/1141)).


### API Changes:

 * Added `GetMessageQueryResultByAttachment` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
 * Added `Id` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Added `LimitConfig` field for [billing.UpdateBudgetPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetPolicyRequest).
 * Added `Volumes` field for [compute.ClusterLogConf](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterLogConf).
 * Added .
 * Removed `ReviewState`, `Reviews` and `RunnerCollaborators` fields for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook).

OpenAPI SHA: 99f644e72261ef5ecf8d74db20f4b7a1e09723cc, Date: 2025-02-11

## Release v0.57.0 (2025-02-03)

### New Features and Improvements

 * Add support for async OAuth token refreshes ([#1135](https://github.com/databricks/databricks-sdk-go/pull/1135)).


### API Changes:

 * Added [a.BudgetPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetPolicyAPI) account-level service.
 * Added [a.EnableIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnableIpAccessListsAPI) account-level service.
 * Added [w.LakeviewEmbedded](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewEmbeddedAPI) workspace-level service and [w.QueryExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryExecutionAPI) workspace-level service.
 * Added [w.RedashConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RedashConfigAPI) workspace-level service.
 * Added `GcpOauthToken` field for [catalog.TemporaryCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryCredentials).
 * Added `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Added `StatementId` field for [dashboards.QueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryAttachment).
 * Added `EffectivePerformanceTarget` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `PerformanceTarget` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `PerformanceTarget` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `EffectivePerformanceTarget` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `PerformanceTarget` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
 * Added `Disabled` and `EffectivePerformanceTarget` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `UserAuthorizedScopes` field for [oauth2.CreateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateCustomAppIntegration).
 * Added `UserAuthorizedScopes` field for [oauth2.GetCustomAppIntegrationOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetCustomAppIntegrationOutput).
 * Added `UserAuthorizedScopes` field for [oauth2.UpdateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateCustomAppIntegration).
 * Added `Contents` field for [serving.HttpRequestResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#HttpRequestResponse).
 * Changed `HttpRequest` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service to type `HttpRequest` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Changed `HttpRequest` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service to return [serving.HttpRequestResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#HttpRequestResponse).
 * Removed `SecurableKind` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Removed `SecurableKind` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `StatusCode` and `Text` fields for [serving.ExternalFunctionResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalFunctionResponse).

OpenAPI SHA: c72c58f97b950fcb924a90ef164bcb10cfcd5ece, Date: 2025-02-03

## Release v0.56.1 (2025-01-23)

### Bug Fixes

 * Do not send query parameters when set to zero value ([#1136](https://github.com/databricks/databricks-sdk-go/pull/1136)).



## Release v0.56.0 (2025-01-22)

### Bug Fixes

 * Support Query parameters for all HTTP operations ([#1124](https://github.com/databricks/databricks-sdk-go/pull/1124)).


### Internal Changes

 * Add download target to MakeFile ([#1125](https://github.com/databricks/databricks-sdk-go/pull/1125)).
 * Delete examples/mocking module ([#1126](https://github.com/databricks/databricks-sdk-go/pull/1126)).
 * Scope the traversing directory in the Recursive list workspace test ([#1120](https://github.com/databricks/databricks-sdk-go/pull/1120)).


### API Changes:

 * Added [w.AccessControl](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccessControlAPI) workspace-level service.
 * Added `HttpRequest` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `ReviewState`, `Reviews` and `RunnerCollaborators` fields for [cleanrooms.CleanRoomAssetNotebook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetNotebook).
 * Added `CleanRoomsNotebookOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `RunAsRepl` field for [jobs.SparkJarTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SparkJarTask).
 * Added `Scopes` field for [oauth2.UpdateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateCustomAppIntegration).
 * Added `Contents` field for [serving.GetOpenApiResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetOpenApiResponse).
 * Added `Activated`, `ActivationUrl`, `AuthenticationType`, `Cloud`, `Comment`, `CreatedAt`, `CreatedBy`, `DataRecipientGlobalMetastoreId`, `IpAccessList`, `MetastoreId`, `Name`, `Owner`, `PropertiesKvpairs`, `Region`, `SharingCode`, `Tokens`, `UpdatedAt` and `UpdatedBy` fields for [sharing.RecipientInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientInfo).
 * Added `ExpirationTime` field for [sharing.RecipientInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientInfo).
 * Added `Pending` enum value for [cleanrooms.CleanRoomAssetStatusEnum](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomAssetStatusEnum).
 * Added `AddNodesFailed`, `AutomaticClusterUpdate`, `AutoscalingBackoff` and `AutoscalingFailed` enum values for [compute.EventType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventType).
 * Added `PendingWarehouse` enum value for [dashboards.MessageStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageStatus).
 * Added `Cpu`, `GpuLarge`, `GpuMedium`, `GpuSmall` and `MultigpuMedium` enum values for [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
 * Changed `Update` method for [w.Recipients](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientsAPI) workspace-level service to return [sharing.RecipientInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientInfo).
 * Changed `Update` method for [w.Recipients](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientsAPI) workspace-level service return type to become non-empty.
 * Changed `Update` method for [w.Recipients](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientsAPI) workspace-level service to type `Update` method for [w.Recipients](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#RecipientsAPI) workspace-level service.
 * Changed `Create` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service with new required argument order.
 * Changed `GetOpenApi` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service return type to become non-empty.
 * Changed `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service to type `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Changed `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service to return [serving.EndpointTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointTags).
 * Changed [serving.EndpointTagList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointTagList) to.
 * Changed `CollaboratorAlias` field for [cleanrooms.CleanRoomCollaborator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomCollaborator) to be required.
 * Changed `CollaboratorAlias` field for [cleanrooms.CleanRoomCollaborator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms#CleanRoomCollaborator) to be required.
 * Changed `Behavior` field for [serving.AiGatewayGuardrailPiiBehavior](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailPiiBehavior) to no longer be required.
 * Changed `Behavior` field for [serving.AiGatewayGuardrailPiiBehavior](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailPiiBehavior) to no longer be required.
 * Changed `Config` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint) to no longer be required.
 * Changed `ProjectId` and `Region` fields for [serving.GoogleCloudVertexAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GoogleCloudVertexAiConfig) to be required.
 * Changed `ProjectId` and `Region` fields for [serving.GoogleCloudVertexAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GoogleCloudVertexAiConfig) to be required.
 * Changed `WorkloadType` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput) to type [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
 * Changed `WorkloadType` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput) to type [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
 * Changed `WorkloadType` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput) to type [serving.ServingModelWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingModelWorkloadType).
 * Changed waiter for [ServingEndpointsAPI.Create](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI.Create).
 * Changed waiter for [ServingEndpointsAPI.UpdateConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI.UpdateConfig).


OpenAPI SHA: 0be1b914249781b5e903b7676fd02255755bc851, Date: 2025-01-22

## Release v0.55.0 (2025-01-13)

### Internal Changes

 * Bump staticcheck to 0.5.1 and add go 1.23 test coverage ([#1106](https://github.com/databricks/databricks-sdk-go/pull/1106)).
 * Bump x/net, x/crypto dependencies ([#1107](https://github.com/databricks/databricks-sdk-go/pull/1107)).
 * Create custom codeql.yml ([#1114](https://github.com/databricks/databricks-sdk-go/pull/1114)).
 * Decouple serving and oauth2 package ([#1110](https://github.com/databricks/databricks-sdk-go/pull/1110)).
 * Migrate workflows that need write access to use hosted runners ([#1112](https://github.com/databricks/databricks-sdk-go/pull/1112)).
 * Move package credentials in config ([#1115](https://github.com/databricks/databricks-sdk-go/pull/1115)).
 * Update Queries test ([#1104](https://github.com/databricks/databricks-sdk-go/pull/1104)).


### API Changes:

 * Added `NoCompute` field for [apps.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppRequest).
 * Added `HasMore` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
 * Added `HasMore` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `PageToken` field for [jobs.GetJobRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobRequest).
 * Added `HasMore` and `NextPageToken` fields for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
 * Added `HasMore` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `RunAs` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `RunAs` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `AuthorizationDetails` and `EndpointUrl` fields for [serving.DataPlaneInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DataPlaneInfo).
 * [Breaking] Changed `Update` method for [a.AccountFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#AccountFederationPolicyAPI) account-level service with new required argument order.
 * [Breaking] Changed `Update` method for [a.ServicePrincipalFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ServicePrincipalFederationPolicyAPI) account-level service with new required argument order.
 * Changed `UpdateMask` field for [oauth2.UpdateAccountFederationPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateAccountFederationPolicyRequest) to no longer be required.
 * Changed `UpdateMask` field for [oauth2.UpdateServicePrincipalFederationPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#UpdateServicePrincipalFederationPolicyRequest) to no longer be required.
 * [Breaking] Changed `DaysOfWeek` field for [pipelines.RestartWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#RestartWindow) to type [pipelines.DayOfWeekList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DayOfWeekList).

OpenAPI SHA: 779817ed8d63031f5ea761fbd25ee84f38feec0d, Date: 2025-01-08

## Release v0.54.0 (2024-12-17)

### API Changes:

 * Added [a.AccountFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#AccountFederationPolicyAPI) account-level service and [a.ServicePrincipalFederationPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ServicePrincipalFederationPolicyAPI) account-level service.
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
 * Added `IsSingleNode`, `Kind` and `UseMlRuntime` fields for [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource).
 * Added `UpdateParameterSyntax` field for [dashboards.MigrateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MigrateDashboardRequest).
 * Added `CleanRoomsNotebookTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `CleanRoomsNotebookTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `CleanRoomsNotebookTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Changed `DaysOfWeek` field for [pipelines.RestartWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#RestartWindow) to type [pipelines.RestartWindowDaysOfWeekList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#RestartWindowDaysOfWeekList).

OpenAPI SHA: a6a317df8327c9b1e5cb59a03a42ffa2aabeef6d, Date: 2024-12-16

## Release v0.53.0 (2024-12-05)

### Bug Fixes

 * Update Changelog file ([#1091](https://github.com/databricks/databricks-sdk-go/pull/1091)).


### Internal Changes

 * Update to latest OpenAPI spec ([#1098](https://github.com/databricks/databricks-sdk-go/pull/1098)).

Note: This release contains breaking changes, please see the API changes below for more details.

### API Changes:

 * Added [cleanrooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/cleanrooms) package.
 * Added `DeletePublicWorkspaceSetting` method for [w.AibiDashboardEmbeddingAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AibiDashboardEmbeddingAccessPolicyAPI) workspace-level service.
 * Added `DeletePublicWorkspaceSetting` method for [w.AibiDashboardEmbeddingApprovedDomains](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AibiDashboardEmbeddingApprovedDomainsAPI) workspace-level service.
 * Added [jobs.CleanRoomTaskRunLifeCycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CleanRoomTaskRunLifeCycleState), [jobs.CleanRoomTaskRunResultState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CleanRoomTaskRunResultState) and [jobs.CleanRoomTaskRunState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CleanRoomTaskRunState).
 * Added [dashboards.DataType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#DataType), [dashboards.QuerySchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QuerySchema) and [dashboards.QuerySchemaColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QuerySchemaColumn).
 * Added [catalog.DatabricksGcpServiceAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DatabricksGcpServiceAccount) and [catalog.GenerateTemporaryServiceCredentialGcpOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryServiceCredentialGcpOptions).
 * Added [files.ContentLength](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ContentLength) and [files.ContentRange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ContentRange).
 * Added [settings.DeleteAibiDashboardEmbeddingAccessPolicySettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAibiDashboardEmbeddingAccessPolicySettingRequest), [settings.DeleteAibiDashboardEmbeddingAccessPolicySettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAibiDashboardEmbeddingAccessPolicySettingResponse), [settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest), [settings.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse), [settings.EgressNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicy), [settings.EgressNetworkPolicyInternetAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicy), [settings.EgressNetworkPolicyInternetAccessPolicyInternetDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyInternetDestination), [settings.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol), [settings.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType), [settings.EgressNetworkPolicyInternetAccessPolicyLogOnlyMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyLogOnlyMode), [settings.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType), [settings.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType), [settings.EgressNetworkPolicyInternetAccessPolicyRestrictionMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyRestrictionMode), [settings.EgressNetworkPolicyInternetAccessPolicyStorageDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyStorageDestination) and [settings.EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType).
 * Added [sharing.PartitionSpecificationPartition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#PartitionSpecificationPartition).
 * Added `DatabricksGcpServiceAccount` field for [catalog.CreateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCredentialRequest).
 * Added `DatabricksGcpServiceAccount` field for [catalog.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialInfo).
 * Added `GcpOptions` field for [catalog.GenerateTemporaryServiceCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryServiceCredentialRequest).
 * Added `DatabricksGcpServiceAccount` field for [catalog.UpdateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCredentialRequest).
 * Added `CachedQuerySchema` field for [dashboards.QueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryAttachment).
 * [Breaking] Changed `ContentLength` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse) to [files.ContentLength](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ContentLength).
 * [Breaking] Changed `ContentLength` field for [files.GetMetadataResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataResponse) to [files.ContentLength](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ContentLength).
 * [Breaking] Removed [catalog.GcpServiceAccountKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GcpServiceAccountKey).
 * [Breaking] Removed [files.FileSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FileSize).
 * [Breaking] Removed `GcpServiceAccountKey` field for [catalog.CreateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCredentialRequest).

OpenAPI SHA: 7016dcbf2e011459416cf408ce21143bcc4b3a25, Date: 2024-12-05

## Release v0.52.0 (2024-11-19)

### Internal Changes

 * Update Jobs GetRun API to support paginated responses for jobs and ForEach tasks ([#1089](https://github.com/databricks/databricks-sdk-go/pull/1089)).


### API Changes:

 * Added `ServicePrincipalClientId` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Added `AzureServicePrincipal`, `GcpServiceAccountKey` and `ReadOnly` fields for [catalog.CreateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCredentialRequest).
 * Added `AzureServicePrincipal`, `ReadOnly` and `UsedForManagedStorage` fields for [catalog.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialInfo).
 * Added `AzureServicePrincipal` and `ReadOnly` fields for [catalog.UpdateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCredentialRequest).
 * Added `ExternalLocationName`, `ReadOnly` and `Url` fields for [catalog.ValidateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateCredentialRequest).
 * Added `IsDir` field for [catalog.ValidateCredentialResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateCredentialResponse).
 * Changed `CreateCredential` and `GenerateTemporaryServiceCredential` methods for [w.Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialsAPI) workspace-level service with new required argument order.
 * Changed `AccessConnectorId` field for [catalog.AzureManagedIdentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentity) to be required.
 * Changed `AccessConnectorId` field for [catalog.AzureManagedIdentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentity) to be required.
 * Changed `Name` field for [catalog.CreateCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCredentialRequest) to be required.
 * Changed `CredentialName` field for [catalog.GenerateTemporaryServiceCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryServiceCredentialRequest) to be required.

OpenAPI SHA: f2385add116e3716c8a90a0b68e204deb40f996c, Date: 2024-11-15

## Release v0.51.0 (2024-11-07)

### Internal Changes

 * Always write message for manual test execution ([#1079](https://github.com/databricks/databricks-sdk-go/pull/1079)).
 * Use error names instead of codes in `errors.go` ([#1080](https://github.com/databricks/databricks-sdk-go/pull/1080)).


### API Changes:

 * Added [w.Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CredentialsAPI) workspace-level service.
 * Added `AzureAad` field for [catalog.GenerateTemporaryTableCredentialResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryTableCredentialResponse).
 * Added `OmitUsername` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
 * Added `FullName` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo).
 * Added `WarehouseId` field for [dashboards.Schedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Schedule).
 * Added `Only` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
 * Added `RestartWindow` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `RestartWindow` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `ConnectionName` field for [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition).
 * Added `RestartWindow` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added `PrivateAccessSettingsId` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Removed [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service.

OpenAPI SHA: d25296d2f4aa7bd6195c816fdf82e0f960f775da, Date: 2024-11-07

## Release v0.50.0 (2024-10-31)

### Internal Changes

 * Add DCO guidelines ([#1047](https://github.com/databricks/databricks-sdk-go/pull/1047)).
 * Add test instructions for external contributors ([#1073](https://github.com/databricks/databricks-sdk-go/pull/1073)).
 * Automatically trigger integration tests ([#1067](https://github.com/databricks/databricks-sdk-go/pull/1067)).
 * Move templates in the code generator ([#1075](https://github.com/databricks/databricks-sdk-go/pull/1075)).
 * Remove unnecessary test ([#1071](https://github.com/databricks/databricks-sdk-go/pull/1071)).


### API Changes:

 * Added [w.AibiDashboardEmbeddingAccessPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AibiDashboardEmbeddingAccessPolicyAPI) workspace-level service and [w.AibiDashboardEmbeddingApprovedDomains](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AibiDashboardEmbeddingApprovedDomainsAPI) workspace-level service.
 * Added `AppDeployment` field for [apps.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppDeploymentRequest).
 * Added `App` field for [apps.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppRequest).
 * Added `App` field for [apps.UpdateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#UpdateAppRequest).
 * Added `Table` field for [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added `Dashboard` field for [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added `Schedule` field for [dashboards.CreateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateScheduleRequest).
 * Added `Subscription` field for [dashboards.CreateSubscriptionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateSubscriptionRequest).
 * Added `Dashboard` field for [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `Schedule` field for [dashboards.UpdateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateScheduleRequest).
 * Added `PageToken` field for [oauth2.ListServicePrincipalSecretsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsRequest).
 * Added `NextPageToken` field for [oauth2.ListServicePrincipalSecretsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListServicePrincipalSecretsResponse).
 * Added `IsNoPublicIpEnabled` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `ExternalCustomerInfo` and `IsNoPublicIpEnabled` fields for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added `LastUsedDay` field for [settings.TokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenInfo).
 * Changed `Create` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service with new required argument order.
 * Changed `ExecuteMessageQuery` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service . New request type is [dashboards.GenieExecuteMessageQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieExecuteMessageQueryRequest).
 * Changed `ExecuteMessageQuery` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service to type `ExecuteMessageQuery` method for [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
 * Changed `Create`, `CreateSchedule`, `CreateSubscription` and `UpdateSchedule` methods for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service with new required argument order.
 * Removed `DeploymentId`, `Mode` and `SourceCodePath` fields for [apps.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppDeploymentRequest).
 * Removed `Description`, `Name` and `Resources` fields for [apps.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppRequest).
 * Removed `Description` and `Resources` fields for [apps.UpdateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#UpdateAppRequest).
 * Removed `Name` and `Spec` fields for [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed `DisplayName`, `ParentPath`, `SerializedDashboard` and `WarehouseId` fields for [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Removed `CronSchedule`, `DisplayName` and `PauseStatus` fields for [dashboards.CreateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateScheduleRequest).
 * Removed `Subscriber` field for [dashboards.CreateSubscriptionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateSubscriptionRequest).
 * Removed `DisplayName`, `Etag`, `SerializedDashboard` and `WarehouseId` fields for [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Removed `CronSchedule`, `DisplayName`, `Etag` and `PauseStatus` fields for [dashboards.UpdateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateScheduleRequest).
 * Removed `PrevPageToken` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).

OpenAPI SHA: 25b2478e5a18c888f0d423249abde5499dc58424, Date: 2024-10-31

## Release v0.49.0 (2024-10-14)

### API Changes:

 * Added [w.DisableLegacyDbfs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DisableLegacyDbfsAPI) workspace-level service.
 * Added `UnityCatalogProvisioningState` field for [catalog.OnlineTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTable).
 * Added `IsTruncated` field for [dashboards.Result](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Result).
 * Added `EffectiveBudgetPolicyId` field for [jobs.BaseJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseJob).
 * Added `BudgetPolicyId` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `EffectiveBudgetPolicyId` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
 * Added `BudgetPolicyId` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `BudgetPolicyId` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `Report` field for [pipelines.IngestionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionConfig).
 * Added `SequenceBy` field for [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
 * Added `NotifyOnOk` field for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert).
 * Added `NotifyOnOk` field for [sql.CreateAlertRequestAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequestAlert).
 * Added `NotifyOnOk` field for [sql.ListAlertsResponseAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsResponseAlert).
 * Added `NotifyOnOk` field for [sql.UpdateAlertRequestAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequestAlert).

OpenAPI SHA: cf9c61453990df0f9453670f2fe68e1b128647a2, Date: 2024-10-14

## Release v0.48.0 (2024-10-03)

### Internal Changes

 * Update SDK to latest OpenAPI spec ([#1057](https://github.com/databricks/databricks-sdk-go/pull/1057)).

Note: This release contains breaking changes, please see the API changes below for more details.

### API Changes:

 * Added `DefaultSourceCodePath` and `Resources` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Added `Resources` field for [apps.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppRequest).
 * Added `Resources` field for [apps.UpdateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#UpdateAppRequest).
 * Added `Schema` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `Schema` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `Schema` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * [Breaking] Changed `Create` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service . New request type is [workspace.CreateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsRequest).
 * [Breaking] Changed `Create` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to type `Create` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service.
 * [Breaking] Changed `Delete` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service . New request type is [workspace.DeleteCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#DeleteCredentialsRequest).
 * [Breaking] Changed `Delete` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to return `any`.
 * Changed `Delete` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to type `Delete` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service.
 * [Breaking] Changed `Get` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service . New request type is [workspace.GetCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsRequest).
 * Changed `Get` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to type `Get` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service.
 * [Breaking] Changed `Get` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to return [workspace.GetCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsResponse).
 * [Breaking] Changed `List` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to return [workspace.ListCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ListCredentialsResponse).
 * Changed `List` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to type `List` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service.
 * Changed `Update` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to type `Update` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service.
 * [Breaking] Changed `Update` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service to return `any`.
 * [Breaking] Changed `Update` method for [w.GitCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GitCredentialsAPI) workspace-level service . New request type is [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest).
 * Changed `Create` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to type `Create` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * [Breaking] Changed `Create` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service . New request type is [workspace.CreateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateRepoRequest).
 * [Breaking] Changed `Create` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to return [workspace.CreateRepoResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateRepoResponse).
 * [Breaking] Changed `Delete` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to return `any`.
 * Changed `Delete` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to type `Delete` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * [Breaking] Changed `Get` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to return [workspace.GetRepoResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoResponse).
 * Changed `Get` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to type `Get` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Changed `Update` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to return `any`.
 * [Breaking] Changed `Update` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service . New request type is [workspace.UpdateRepoRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateRepoRequest).
 * Changed `Update` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service to type `Update` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * [Breaking] Changed `CredentialId` and `GitProvider` fields for [workspace.CreateCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CreateCredentialsResponse) to be required.
 * Changed `CredentialId` field for [workspace.CredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#CredentialInfo) to be required.
 * Changed `CredentialId` field for [workspace.GetCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetCredentialsResponse) to be required.
 * Changed `Patterns` field for [workspace.SparseCheckout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#SparseCheckout) to type [workspace.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#List).
 * Changed `Patterns` field for [workspace.SparseCheckoutUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#SparseCheckoutUpdate) to type [workspace.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#List).
 * [Breaking] Changed `GitProvider` field for [workspace.UpdateCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#UpdateCredentialsRequest) to be required.

OpenAPI SHA: 0c86ea6dbd9a730c24ff0d4e509603e476955ac5, Date: 2024-10-02

## Release v0.47.0 (2024-10-01)

### Internal Changes

 * Update SDK to latest OpenAPI spec ([#1042](https://github.com/databricks/databricks-sdk-go/pull/1042)).


### API Changes:

 * Added [w.DisableLegacyAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DisableLegacyAccessAPI) workspace-level service and [a.DisableLegacyFeatures](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DisableLegacyFeaturesAPI) account-level service.
 * Added [w.TemporaryTableCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TemporaryTableCredentialsAPI) workspace-level service.
 * Added `PutAiGateway` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added [apps.ApplicationState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ApplicationState), [apps.ApplicationStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ApplicationStatus), [apps.ComputeState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeState) and [apps.ComputeStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#ComputeStatus).
 * Added [catalog.AwsCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsCredentials), [catalog.AzureUserDelegationSas](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureUserDelegationSas), [catalog.GcpOauthToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GcpOauthToken), [catalog.GenerateTemporaryTableCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryTableCredentialRequest), [catalog.GenerateTemporaryTableCredentialResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GenerateTemporaryTableCredentialResponse), [catalog.R2Credentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#R2Credentials) and [catalog.TableOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableOperation).
 * Added [serving.AiGatewayConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayConfig), [serving.AiGatewayGuardrailParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailParameters), [serving.AiGatewayGuardrailPiiBehavior](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailPiiBehavior), [serving.AiGatewayGuardrailPiiBehaviorBehavior](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrailPiiBehaviorBehavior), [serving.AiGatewayGuardrails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayGuardrails), [serving.AiGatewayInferenceTableConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayInferenceTableConfig), [serving.AiGatewayRateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimit), [serving.AiGatewayRateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimitKey), [serving.AiGatewayRateLimitRenewalPeriod](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayRateLimitRenewalPeriod), [serving.AiGatewayUsageTrackingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AiGatewayUsageTrackingConfig), [serving.PutAiGatewayRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayRequest) and [serving.PutAiGatewayResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutAiGatewayResponse).
 * Added [settings.BooleanMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#BooleanMessage), [settings.DeleteDisableLegacyAccessRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDisableLegacyAccessRequest), [settings.DeleteDisableLegacyAccessResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDisableLegacyAccessResponse), [settings.DeleteDisableLegacyFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDisableLegacyFeaturesRequest), [settings.DeleteDisableLegacyFeaturesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDisableLegacyFeaturesResponse), [settings.DisableLegacyAccess](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DisableLegacyAccess), [settings.DisableLegacyFeatures](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DisableLegacyFeatures), [settings.GetDisableLegacyAccessRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDisableLegacyAccessRequest), [settings.GetDisableLegacyFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDisableLegacyFeaturesRequest), [settings.UpdateDisableLegacyAccessRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDisableLegacyAccessRequest) and [settings.UpdateDisableLegacyFeaturesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDisableLegacyFeaturesRequest).
 * Added `AppStatus` and `ComputeStatus` fields for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Added `DeploymentId` field for [apps.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppDeploymentRequest).
 * Added `ExternalAccessEnabled` field for [catalog.GetMetastoreSummaryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetMetastoreSummaryResponse).
 * Added `IncludeManifestCapabilities` field for [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest).
 * Added `IncludeManifestCapabilities` field for [catalog.ListSummariesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSummariesRequest).
 * Added `IncludeManifestCapabilities` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
 * Added `ExternalAccessEnabled` field for [catalog.MetastoreInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoreInfo).
 * Added `BudgetPolicyId` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `BudgetPolicyId` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `EffectiveBudgetPolicyId` field for [pipelines.GetPipelineResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelineResponse).
 * Added `BudgetPolicyId` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added `AiGateway` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Added `AiGateway` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `AiGateway` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added `WorkspaceId` field for [settings.TokenInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenInfo).
 * Changed `Delete`, `Start` and `Stop` methods for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service to return [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Changed `Deploy` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppsAPI) workspace-level service with new required argument order.
 * Changed `SourceCodePath` field for [apps.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppDeployment) to no longer be required.
 * Changed `SourceCodePath` field for [apps.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#CreateAppDeploymentRequest) to no longer be required.
 * Changed `ReturnParams` and `RoutineDependencies` fields for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) to no longer be required.
 * Removed [apps.AppState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppState), [apps.AppStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#AppStatus), `any` and `any`.
 * Removed [sql.ClientCallContext](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ClientCallContext), [sql.EncodedText](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EncodedText), [sql.EncodedTextEncoding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EncodedTextEncoding), [sql.QuerySource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySource), [sql.QuerySourceDriverInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceDriverInfo), [sql.QuerySourceEntryPoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceEntryPoint), [sql.QuerySourceJobManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceJobManager), [sql.QuerySourceTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceTrigger) and [sql.ServerlessChannelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ServerlessChannelInfo).
 * Removed `Status` field for [apps.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps#App).
 * Removed `QuerySource` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).

OpenAPI SHA: 6f6b1371e640f2dfeba72d365ac566368656f6b6, Date: 2024-09-19

## Release v0.46.0 (2024-09-04)

### Bug Fixes

 * Fail fast when authenticating if host is not configured ([#1033](https://github.com/databricks/databricks-sdk-go/pull/1033)).
 * Improve non-JSON error handling ([#1031](https://github.com/databricks/databricks-sdk-go/pull/1031)).


### Internal Changes

 * Add TestAccCreateOboTokenOnAws to flaky test list ([#1029](https://github.com/databricks/databricks-sdk-go/pull/1029)).
 * Add workflows manage integration tests checks ([#1032](https://github.com/databricks/databricks-sdk-go/pull/1032)).
 * Fix TestMwsAccWorkspaces cleanup ([#1028](https://github.com/databricks/databricks-sdk-go/pull/1028)).
 * Improve integration test comment ([#1035](https://github.com/databricks/databricks-sdk-go/pull/1035)).
 * Temporary ignore Metastore test failures ([#1027](https://github.com/databricks/databricks-sdk-go/pull/1027)).
 * Update test to support new accounts ([#1026](https://github.com/databricks/databricks-sdk-go/pull/1026)).
 * Use statuses instead of checks ([#1036](https://github.com/databricks/databricks-sdk-go/pull/1036)).


### API Changes:

 * Added `RegenerateDashboard` method for [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI) workspace-level service.
 * Added [catalog.RegenerateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegenerateDashboardRequest) and [catalog.RegenerateDashboardResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegenerateDashboardResponse).
 * Added [jobs.QueueDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#QueueDetails), [jobs.QueueDetailsCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#QueueDetailsCodeCode), [jobs.RunLifecycleStateV2State](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunLifecycleStateV2State), [jobs.RunStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunStatus), [jobs.TerminationCodeCode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationCodeCode), [jobs.TerminationDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationDetails) and [jobs.TerminationTypeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TerminationTypeType).
 * Added `Status` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `Status` field for [jobs.RepairHistoryItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairHistoryItem).
 * Added `Status` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `Status` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `MaxProvisionedThroughput` and `MinProvisionedThroughput` fields for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `ColumnsToSync` field for [vectorsearch.DeltaSyncVectorIndexSpecRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecRequest).
 * Changed `WorkloadSize` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to no longer be required.

OpenAPI SHA: d05898328669a3f8ab0c2ecee37db2673d3ea3f7, Date: 2024-09-04

## Release v0.45.0 (2024-08-21)

### Bug Fixes

 * Add INVALID_STATE to error code mapping ([#1014](https://github.com/databricks/databricks-sdk-go/pull/1014)).
 * Do not specify `--tenant` flag when fetching managed identity access token from the CLI ([#1021](https://github.com/databricks/databricks-sdk-go/pull/1021)).


### Internal Changes

 * Add terraform aliases to Entity ([#1017](https://github.com/databricks/databricks-sdk-go/pull/1017)).
 * Added Service.NamedIdMap ([#1016](https://github.com/databricks/databricks-sdk-go/pull/1016)).
 * Fix billing test for budget configuration update ([#1019](https://github.com/databricks/databricks-sdk-go/pull/1019)).


### API Changes:

 * Added [w.PolicyComplianceForClusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#PolicyComplianceForClustersAPI) workspace-level service.
 * Added [w.PolicyComplianceForJobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PolicyComplianceForJobsAPI) workspace-level service.
 * Added [w.ResourceQuotas](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ResourceQuotasAPI) workspace-level service.
 * Added [catalog.GetQuotaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetQuotaRequest), [catalog.GetQuotaResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetQuotaResponse), [catalog.ListQuotasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListQuotasRequest), [catalog.ListQuotasResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListQuotasResponse) and [catalog.QuotaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QuotaInfo).
 * Added [compute.ClusterCompliance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterCompliance), [compute.ClusterSettingsChange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSettingsChange), [compute.EnforceClusterComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceRequest), [compute.EnforceClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EnforceClusterComplianceResponse), [compute.GetClusterComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterComplianceRequest), [compute.GetClusterComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterComplianceResponse), [compute.ListClusterCompliancesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClusterCompliancesRequest) and [compute.ListClusterCompliancesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClusterCompliancesResponse).
 * Added [jobs.EnforcePolicyComplianceForJobResponseJobClusterSettingsChange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#EnforcePolicyComplianceForJobResponseJobClusterSettingsChange), [jobs.EnforcePolicyComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#EnforcePolicyComplianceRequest), [jobs.EnforcePolicyComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#EnforcePolicyComplianceResponse), [jobs.GetPolicyComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetPolicyComplianceRequest), [jobs.GetPolicyComplianceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetPolicyComplianceResponse), [jobs.JobCompliance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCompliance), [jobs.ListJobComplianceForPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ListJobComplianceForPolicyResponse) and [jobs.ListJobComplianceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ListJobComplianceRequest).
 * Added `Fallback` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Added `Fallback` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `Fallback` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `JobRunId` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `JobRunId` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `IncludeMetrics` field for [sql.ListQueryHistoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryHistoryRequest).
 * Added `StatementIds` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Removed [sql.ContextFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ContextFilter).
 * Removed `ContextFilter` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Removed `PipelineId` and `PipelineUpdateId` fields for [sql.QuerySource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySource).

OpenAPI SHA: 3eae49b444cac5a0118a3503e5b7ecef7f96527a, Date: 2024-08-21

## Release v0.44.0 (2024-08-13)

### New Features and Improvements

 * Remove deprecated `WithImpl` and `Impl` service methods ([#1003](https://github.com/databricks/databricks-sdk-go/pull/1003)).


### Bug Fixes

 * Allowed filtering for graviton DBRs in `ClusterAPI.SelectSparkVersion` ([#1004](https://github.com/databricks/databricks-sdk-go/pull/1004)).
 * Avoid loading the response bodies twice in memory when parsing `bytes.Buffer` ([#984](https://github.com/databricks/databricks-sdk-go/pull/984)).
 * Bump google.golang.org/grpc from 1.64.0 to 1.64.1 ([#974](https://github.com/databricks/databricks-sdk-go/pull/974)).
 * Fix default-auth example when less than 10 clusters ([#1012](https://github.com/databricks/databricks-sdk-go/pull/1012)).
 * Infer Azure tenant ID if not set ([#910](https://github.com/databricks/databricks-sdk-go/pull/910)).


### Internal Changes

 * Add comment to validate workflow ([#987](https://github.com/databricks/databricks-sdk-go/pull/987)).
 * Add error mapping for GetRun ([#1006](https://github.com/databricks/databricks-sdk-go/pull/1006)).
 * Add missing Package to Entity ([#995](https://github.com/databricks/databricks-sdk-go/pull/995)).
 * Add prefix to Dependabot commit messages ([#976](https://github.com/databricks/databricks-sdk-go/pull/976)).
 * Configure Dependabot for security updates only ([#988](https://github.com/databricks/databricks-sdk-go/pull/988)).
 * Enable mixins via struct embedding ([#1000](https://github.com/databricks/databricks-sdk-go/pull/1000)).
 * Fix GetWorkspaceClient test & GCP SQL Warehouse Creation test ([#1010](https://github.com/databricks/databricks-sdk-go/pull/1010)).
 * Fix integration tests ([#1008](https://github.com/databricks/databricks-sdk-go/pull/1008)).
 * Fix processing of `quoted` titles ([#989](https://github.com/databricks/databricks-sdk-go/pull/989)).
 * Generate SDK from Open API  ([#997](https://github.com/databricks/databricks-sdk-go/pull/997)).
 * Log a warning when declaring inline entities  ([#994](https://github.com/databricks/databricks-sdk-go/pull/994)).
 * Move PR message validation to a separate workflow ([#983](https://github.com/databricks/databricks-sdk-go/pull/983)).
 * Pin jobs APIs to 2.1 in SDKs ([#993](https://github.com/databricks/databricks-sdk-go/pull/993)).
 * Prefix all extension files with `ext_` ([#1001](https://github.com/databricks/databricks-sdk-go/pull/1001)).
 * Trigger the `validate` workflow in the merge queue ([#986](https://github.com/databricks/databricks-sdk-go/pull/986)).
 * Update OpenAPI spec ([#991](https://github.com/databricks/databricks-sdk-go/pull/991)).
 * Use HTTP status text as message if the error response body is empty ([#990](https://github.com/databricks/databricks-sdk-go/pull/990)).


### API Changes:

 * Added [apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/apps) package.
 * Added [a.UsageDashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UsageDashboardsAPI) account-level service.
 * Added [w.AlertsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsLegacyAPI) workspace-level service, [w.QueriesLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesLegacyAPI) workspace-level service and [w.QueryVisualizationsLegacy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsLegacyAPI) workspace-level service.
 * Added [w.Genie](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAPI) workspace-level service.
 * Added [w.NotificationDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NotificationDestinationsAPI) workspace-level service.
 * Added `Update` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `ListVisualizations` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service.
 * Added [catalog.GetBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsSecurableType) and [catalog.UpdateBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateBindingsSecurableType).
 * Added [billing.ActionConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ActionConfiguration), [billing.ActionConfigurationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ActionConfigurationType), [billing.AlertConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfiguration), [billing.AlertConfigurationQuantityType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfigurationQuantityType), [billing.AlertConfigurationTimePeriod](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfigurationTimePeriod), [billing.AlertConfigurationTriggerType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#AlertConfigurationTriggerType), [billing.BudgetConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfiguration), [billing.BudgetConfigurationFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfigurationFilter), [billing.BudgetConfigurationFilterClause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfigurationFilterClause), [billing.BudgetConfigurationFilterOperator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfigurationFilterOperator), [billing.BudgetConfigurationFilterTagClause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfigurationFilterTagClause), [billing.BudgetConfigurationFilterWorkspaceIdClause](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetConfigurationFilterWorkspaceIdClause), [billing.CreateBillingUsageDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBillingUsageDashboardRequest), [billing.CreateBillingUsageDashboardResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBillingUsageDashboardResponse), [billing.CreateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudget), [billing.CreateBudgetConfigurationBudgetActionConfigurations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudgetActionConfigurations), [billing.CreateBudgetConfigurationBudgetAlertConfigurations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationBudgetAlertConfigurations), [billing.CreateBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationRequest), [billing.CreateBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationResponse), [billing.DeleteBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DeleteBudgetConfigurationRequest), `any`, [billing.GetBillingUsageDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBillingUsageDashboardRequest), [billing.GetBillingUsageDashboardResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBillingUsageDashboardResponse), [billing.GetBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBudgetConfigurationRequest), [billing.GetBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBudgetConfigurationResponse), [billing.ListBudgetConfigurationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListBudgetConfigurationsRequest), [billing.ListBudgetConfigurationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListBudgetConfigurationsResponse), [billing.UpdateBudgetConfigurationBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationBudget), [billing.UpdateBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationRequest), [billing.UpdateBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationResponse) and [billing.UsageDashboardType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UsageDashboardType).
 * Added [compute.ListClustersFilterBy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersFilterBy), [compute.ListClustersSortBy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersSortBy), [compute.ListClustersSortByDirection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersSortByDirection), [compute.ListClustersSortByField](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersSortByField), [compute.UpdateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateCluster), [compute.UpdateClusterResource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#UpdateClusterResource) and `any`.
 * Added [dashboards.ExecuteMessageQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ExecuteMessageQueryRequest), [dashboards.GenieAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieAttachment), [dashboards.GenieConversation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieConversation), [dashboards.GenieCreateConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieCreateConversationMessageRequest), [dashboards.GenieGetConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetConversationMessageRequest), [dashboards.GenieGetMessageQueryResultRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetMessageQueryResultRequest), [dashboards.GenieGetMessageQueryResultResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieGetMessageQueryResultResponse), [dashboards.GenieMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieMessage), [dashboards.GenieStartConversationMessageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieStartConversationMessageRequest), [dashboards.GenieStartConversationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GenieStartConversationResponse), [dashboards.MessageError](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageError), [dashboards.MessageErrorType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageErrorType), [dashboards.MessageStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MessageStatus), [dashboards.QueryAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#QueryAttachment), [dashboards.Result](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Result) and [dashboards.TextAttachment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TextAttachment).
 * Added `any`, [iam.MigratePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MigratePermissionsRequest) and [iam.MigratePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MigratePermissionsResponse).
 * Added [oauth2.ListCustomAppIntegrationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListCustomAppIntegrationsRequest) and [oauth2.ListPublishedAppIntegrationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListPublishedAppIntegrationsRequest).
 * Added [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition) and [pipelines.PipelineStateInfoHealth](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineStateInfoHealth).
 * Added [serving.GoogleCloudVertexAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GoogleCloudVertexAiConfig).
 * Added [settings.Config](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#Config), [settings.CreateNotificationDestinationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNotificationDestinationRequest), [settings.DeleteNotificationDestinationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteNotificationDestinationRequest), [settings.DestinationType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DestinationType), [settings.EmailConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EmailConfig), `any`, [settings.GenericWebhookConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GenericWebhookConfig), [settings.GetNotificationDestinationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetNotificationDestinationRequest), [settings.ListNotificationDestinationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNotificationDestinationsRequest), [settings.ListNotificationDestinationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNotificationDestinationsResponse), [settings.ListNotificationDestinationsResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNotificationDestinationsResult), [settings.MicrosoftTeamsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#MicrosoftTeamsConfig), [settings.NotificationDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NotificationDestination), [settings.PagerdutyConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PagerdutyConfig), [settings.SlackConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SlackConfig) and [settings.UpdateNotificationDestinationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateNotificationDestinationRequest).
 * Added [sql.AlertCondition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertCondition), [sql.AlertConditionOperand](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertConditionOperand), [sql.AlertConditionThreshold](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertConditionThreshold), [sql.AlertOperandColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOperandColumn), [sql.AlertOperandValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOperandValue), [sql.AlertOperator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOperator), [sql.ClientCallContext](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ClientCallContext), [sql.ContextFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ContextFilter), [sql.CreateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequest), [sql.CreateAlertRequestAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequestAlert), [sql.CreateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryRequest), [sql.CreateQueryRequestQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryRequestQuery), [sql.CreateQueryVisualizationsLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryVisualizationsLegacyRequest), [sql.CreateVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateVisualizationRequest), [sql.CreateVisualizationRequestVisualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateVisualizationRequestVisualization), [sql.DatePrecision](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DatePrecision), [sql.DateRange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DateRange), [sql.DateRangeValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DateRangeValue), [sql.DateRangeValueDynamicDateRange](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DateRangeValueDynamicDateRange), [sql.DateValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DateValue), [sql.DateValueDynamicDate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DateValueDynamicDate), [sql.DeleteAlertsLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteAlertsLegacyRequest), [sql.DeleteQueriesLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueriesLegacyRequest), [sql.DeleteQueryVisualizationsLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryVisualizationsLegacyRequest), [sql.DeleteVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteVisualizationRequest), `any`, [sql.EncodedText](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EncodedText), [sql.EncodedTextEncoding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EncodedTextEncoding), [sql.EnumValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#EnumValue), [sql.GetAlertsLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetAlertsLegacyRequest), [sql.GetQueriesLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetQueriesLegacyRequest), [sql.LegacyAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyAlert), [sql.LegacyAlertState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyAlertState), [sql.LegacyQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyQuery), [sql.LegacyVisualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyVisualization), [sql.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LifecycleState), [sql.ListAlertsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsRequest), [sql.ListAlertsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsResponse), [sql.ListAlertsResponseAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsResponseAlert), [sql.ListQueriesLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueriesLegacyRequest), [sql.ListQueryObjectsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryObjectsResponse), [sql.ListQueryObjectsResponseQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryObjectsResponseQuery), [sql.ListVisualizationsForQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListVisualizationsForQueryRequest), [sql.ListVisualizationsForQueryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListVisualizationsForQueryResponse), [sql.NumericValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#NumericValue), [sql.QueryBackedValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryBackedValue), [sql.QueryParameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryParameter), [sql.QuerySource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySource), [sql.QuerySourceDriverInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceDriverInfo), [sql.QuerySourceEntryPoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceEntryPoint), [sql.QuerySourceJobManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceJobManager), [sql.QuerySourceTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QuerySourceTrigger), [sql.RestoreQueriesLegacyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RestoreQueriesLegacyRequest), [sql.RunAsMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RunAsMode), [sql.ServerlessChannelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ServerlessChannelInfo), [sql.StatementResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementResponse), [sql.TextValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TextValue), [sql.TrashAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashAlertRequest), [sql.TrashQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashQueryRequest), [sql.UpdateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequest), [sql.UpdateAlertRequestAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequestAlert), [sql.UpdateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequest), [sql.UpdateQueryRequestQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequestQuery), [sql.UpdateVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateVisualizationRequest) and [sql.UpdateVisualizationRequestVisualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateVisualizationRequestVisualization).
 * Added `Force` field for [catalog.DeleteSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteSchemaRequest).
 * Added `MaxResults` and `PageToken` fields for [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest).
 * Added `IncludeAliases` field for [catalog.GetByAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetByAliasRequest).
 * Added `IncludeAliases` field for [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest).
 * Added `IncludeAliases` field for [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest).
 * Added `MaxResults` and `PageToken` fields for [catalog.ListSystemSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSystemSchemasRequest).
 * Added `NextPageToken` field for [catalog.ListSystemSchemasResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSystemSchemasResponse).
 * Added `Aliases` field for [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Added `NextPageToken` field for [catalog.WorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsResponse).
 * Added `Version` field for [compute.GetPolicyFamilyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetPolicyFamilyRequest).
 * Added `FilterBy`, `PageSize`, `PageToken` and `SortBy` fields for [compute.ListClustersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersRequest).
 * Added `NextPageToken` and `PrevPageToken` fields for [compute.ListClustersResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersResponse).
 * Added `PageToken` field for [jobs.GetRunRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetRunRequest).
 * Added `Iterations`, `NextPageToken` and `PrevPageToken` fields for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `CreateTime`, `CreatedBy`, `CreatorUsername` and `Scopes` fields for [oauth2.GetCustomAppIntegrationOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetCustomAppIntegrationOutput).
 * Added `NextPageToken` field for [oauth2.GetCustomAppIntegrationsOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetCustomAppIntegrationsOutput).
 * Added `CreateTime` and `CreatedBy` fields for [oauth2.GetPublishedAppIntegrationOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetPublishedAppIntegrationOutput).
 * Added `NextPageToken` field for [oauth2.GetPublishedAppIntegrationsOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetPublishedAppIntegrationsOutput).
 * Added `EnableLocalDiskEncryption` field for [pipelines.PipelineCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineCluster).
 * Added `Whl` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
 * Added `Health` field for [pipelines.PipelineStateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineStateInfo).
 * Added `Ai21labsApiKeyPlaintext` field for [serving.Ai21LabsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Ai21LabsConfig).
 * Added `AwsAccessKeyIdPlaintext` and `AwsSecretAccessKeyPlaintext` fields for [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig).
 * Added `AnthropicApiKeyPlaintext` field for [serving.AnthropicConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AnthropicConfig).
 * Added `CohereApiBase` and `CohereApiKeyPlaintext` fields for [serving.CohereConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CohereConfig).
 * Added `DatabricksApiTokenPlaintext` field for [serving.DatabricksModelServingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DatabricksModelServingConfig).
 * Added `GoogleCloudVertexAiConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `MicrosoftEntraClientSecretPlaintext` and `OpenaiApiKeyPlaintext` fields for [serving.OpenAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#OpenAiConfig).
 * Added `PalmApiKeyPlaintext` field for [serving.PaLmConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PaLmConfig).
 * Added `ExpirationTime` field for [sharing.CreateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateRecipient).
 * Added `NextPageToken` field for [sharing.GetRecipientSharePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetRecipientSharePermissionsResponse).
 * Added `NextPageToken` field for [sharing.ListProviderSharesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProviderSharesResponse).
 * Added `MaxResults` and `PageToken` fields for [sharing.ListProvidersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProvidersRequest).
 * Added `NextPageToken` field for [sharing.ListProvidersResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListProvidersResponse).
 * Added `MaxResults` and `PageToken` fields for [sharing.ListRecipientsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListRecipientsRequest).
 * Added `NextPageToken` field for [sharing.ListRecipientsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListRecipientsResponse).
 * Added `MaxResults` and `PageToken` fields for [sharing.ListSharesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListSharesRequest).
 * Added `NextPageToken` field for [sharing.ListSharesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListSharesResponse).
 * Added `MaxResults` and `PageToken` fields for [sharing.SharePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharePermissionsRequest).
 * Added `ExpirationTime` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient).
 * Added `MaxResults` and `PageToken` fields for [sharing.UpdateSharePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateSharePermissions).
 * Added `Condition`, `CreateTime`, `CustomBody`, `CustomSubject`, `DisplayName`, `LifecycleState`, `OwnerUserName`, `ParentPath`, `QueryId`, `SecondsToRetrigger`, `TriggerTime` and `UpdateTime` fields for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert).
 * Added `Id` field for [sql.GetAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetAlertRequest).
 * Added `Id` field for [sql.GetQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetQueryRequest).
 * Added `PageToken` field for [sql.ListQueriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueriesRequest).
 * Added `ApplyAutoLimit`, `Catalog`, `CreateTime`, `DisplayName`, `LastModifierUserName`, `LifecycleState`, `OwnerUserName`, `Parameters`, `ParentPath`, `QueryText`, `RunAsMode`, `Schema`, `UpdateTime` and `WarehouseId` fields for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query).
 * Added `ContextFilter` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Added `QuerySource` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
 * Added `CreateTime`, `DisplayName`, `QueryId`, `SerializedOptions`, `SerializedQueryPlan` and `UpdateTime` fields for [sql.Visualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Visualization).
 * Changed `Create` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to return [billing.CreateBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationResponse).
 * Changed `Create` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service . New request type is [billing.CreateBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#CreateBudgetConfigurationRequest).
 * Changed `Delete` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service . New request type is [billing.DeleteBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DeleteBudgetConfigurationRequest).
 * Changed `Delete` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to return `any`.
 * Changed `Get` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to return [billing.GetBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBudgetConfigurationResponse).
 * Changed `Get` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service . New request type is [billing.GetBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBudgetConfigurationRequest).
 * Changed `List` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to return [billing.ListBudgetConfigurationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListBudgetConfigurationsResponse).
 * Changed `List` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to require request of [billing.ListBudgetConfigurationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#ListBudgetConfigurationsRequest).
 * Changed `Update` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service . New request type is [billing.UpdateBudgetConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationRequest).
 * Changed `Update` method for [a.Budgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetsAPI) account-level service to return [billing.UpdateBudgetConfigurationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#UpdateBudgetConfigurationResponse).
 * Changed `Create` method for [a.CustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CustomAppIntegrationAPI) account-level service with new required argument order.
 * Changed `List` method for [a.CustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CustomAppIntegrationAPI) account-level service to require request of [oauth2.ListCustomAppIntegrationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListCustomAppIntegrationsRequest).
 * Changed `List` method for [a.PublishedAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#PublishedAppIntegrationAPI) account-level service to require request of [oauth2.ListPublishedAppIntegrationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListPublishedAppIntegrationsRequest).
 * Changed `Delete` method for [a.WorkspaceAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#WorkspaceAssignmentAPI) account-level service to return `any`.
 * Changed `Update` method for [a.WorkspaceAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#WorkspaceAssignmentAPI) account-level service with new required argument order.
 * Changed `Create` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service . New request type is [sql.CreateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateAlertRequest).
 * Changed `Delete` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service . New request type is [sql.TrashAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashAlertRequest).
 * Changed `Delete` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service to return `any`.
 * Changed `Get` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service with new required argument order.
 * Changed `List` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service to require request of [sql.ListAlertsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsRequest).
 * Changed `List` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service to return [sql.ListAlertsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListAlertsResponse).
 * Changed `Update` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service . New request type is [sql.UpdateAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateAlertRequest).
 * Changed `Update` method for [w.Alerts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertsAPI) workspace-level service to return [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert).
 * Changed `Create` and `Edit` methods for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.ModelVersions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionsAPI) workspace-level service to return [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Changed `MigratePermissions` method for [w.PermissionMigration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationAPI) workspace-level service to return [iam.MigratePermissionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MigratePermissionsResponse).
 * Changed `MigratePermissions` method for [w.PermissionMigration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationAPI) workspace-level service . New request type is [iam.MigratePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#MigratePermissionsRequest).
 * Changed `Create` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service . New request type is [sql.CreateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryRequest).
 * Changed `Delete` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service . New request type is [sql.TrashQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TrashQueryRequest).
 * Changed `Delete` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service to return `any`.
 * Changed `Get` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service with new required argument order.
 * Changed `List` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service to return [sql.ListQueryObjectsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryObjectsResponse).
 * Changed `Update` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service . New request type is [sql.UpdateQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateQueryRequest).
 * Changed `Create` method for [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service . New request type is [sql.CreateVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateVisualizationRequest).
 * Changed `Delete` method for [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service . New request type is [sql.DeleteVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteVisualizationRequest).
 * Changed `Delete` method for [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service to return `any`.
 * Changed `Update` method for [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service . New request type is [sql.UpdateVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UpdateVisualizationRequest).
 * Changed `List` method for [w.Shares](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharesAPI) workspace-level service to require request of [sharing.ListSharesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListSharesRequest).
 * Changed `ExecuteStatement` and `GetStatement` methods for [w.StatementExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementExecutionAPI) workspace-level service to return [sql.StatementResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementResponse).
 * Changed `SecurableType` field for [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest) to [catalog.GetBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsSecurableType).
 * Changed `SecurableType` field for [catalog.UpdateWorkspaceBindingsParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsParameters) to [catalog.UpdateBindingsSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateBindingsSecurableType).
 * Changed `Name` field for [compute.CreatePolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreatePolicy) to no longer be required.
 * Changed `Name` field for [compute.EditPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditPolicy) to no longer be required.
 * Changed `PolicyFamilyId` field for [compute.GetPolicyFamilyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetPolicyFamilyRequest) to `string`.
 * Changed `PolicyFamilies` field for [compute.ListPolicyFamiliesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListPolicyFamiliesResponse) to no longer be required.
 * Changed `Definition`, `Description`, `Name` and `PolicyFamilyId` fields for [compute.PolicyFamily](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#PolicyFamily) to no longer be required.
 * Changed `Permissions` field for [iam.UpdateWorkspaceAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UpdateWorkspaceAssignments) to no longer be required.
 * Changed `AccessControlList` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) to [jobs.JobAccessControlRequestList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlRequestList).
 * Changed `AccessControlList` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun) to [jobs.JobAccessControlRequestList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlRequestList).
 * Changed `Name` and `RedirectUrls` fields for [oauth2.CreateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateCustomAppIntegration) to no longer be required.
 * Changed `IngestionDefinition` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline) to [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
 * Changed `IngestionDefinition` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline) to [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
 * Changed `IngestionDefinition` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec) to [pipelines.IngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionPipelineDefinition).
 * Changed `Ai21labsApiKey` field for [serving.Ai21LabsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Ai21LabsConfig) to no longer be required.
 * Changed `AwsAccessKeyId` and `AwsSecretAccessKey` fields for [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig) to no longer be required.
 * Changed `AnthropicApiKey` field for [serving.AnthropicConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AnthropicConfig) to no longer be required.
 * Changed `CohereApiKey` field for [serving.CohereConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CohereConfig) to no longer be required.
 * Changed `DatabricksApiToken` field for [serving.DatabricksModelServingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DatabricksModelServingConfig) to no longer be required.
 * Changed `PalmApiKey` field for [serving.PaLmConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PaLmConfig) to no longer be required.
 * Changed `Tags` field for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query) to [sql.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#List).
 * Changed `UserIds` and `WarehouseIds` fields for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter) to [sql.List](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#List).
 * Changed `Results` field for [sql.QueryList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryList) to [sql.LegacyQueryList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyQueryList).
 * Changed `Visualization` field for [sql.Widget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Widget) to [sql.LegacyVisualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#LegacyVisualization).
 * Removed [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed `Restore` method for [w.Queries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueriesAPI) workspace-level service.
 * Removed [marketplace.FilterType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#FilterType), [marketplace.ProviderIconFile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderIconFile), [marketplace.ProviderIconType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderIconType), [marketplace.ProviderListingSummaryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderListingSummaryInfo), [marketplace.SortBy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#SortBy) and [marketplace.VisibilityFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#VisibilityFilter).
 * Removed [billing.Budget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#Budget), [billing.BudgetAlert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetAlert), [billing.BudgetList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetList), [billing.BudgetWithStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetWithStatus), [billing.BudgetWithStatusStatusDailyItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BudgetWithStatusStatusDailyItem), [billing.DeleteBudgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DeleteBudgetRequest), `any`, [billing.GetBudgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#GetBudgetRequest), `any`, [billing.WrappedBudget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedBudget) and [billing.WrappedBudgetWithStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#WrappedBudgetWithStatus).
 * Removed `any`, [iam.PermissionMigrationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationRequest) and [iam.PermissionMigrationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationResponse).
 * Removed [pipelines.ManagedIngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ManagedIngestionPipelineDefinition).
 * Removed [serving.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#App), [serving.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeployment), [serving.AppDeploymentArtifacts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentArtifacts), [serving.AppDeploymentMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentMode), [serving.AppDeploymentState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentState), [serving.AppDeploymentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentStatus), [serving.AppEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppEnvironment), [serving.AppState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppState), [serving.AppStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppStatus), [serving.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppDeploymentRequest), [serving.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppRequest), [serving.DeleteAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeleteAppRequest), [serving.EnvVariable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EnvVariable), [serving.GetAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppDeploymentRequest), [serving.GetAppEnvironmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppEnvironmentRequest), [serving.GetAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppRequest), [serving.ListAppDeploymentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppDeploymentsRequest), [serving.ListAppDeploymentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppDeploymentsResponse), [serving.ListAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsRequest), [serving.ListAppsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsResponse), [serving.StartAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#StartAppRequest), [serving.StopAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#StopAppRequest), `any` and [serving.UpdateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#UpdateAppRequest).
 * Removed [sql.CreateQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryVisualizationRequest), [sql.DeleteAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteAlertRequest), [sql.DeleteQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryRequest), [sql.DeleteQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryVisualizationRequest), [sql.ExecuteStatementResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementResponse), [sql.GetStatementResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetStatementResponse), [sql.RestoreQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RestoreQueryRequest), [sql.StatementId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementId), [sql.UserId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#UserId) and [sql.WarehouseId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseId).
 * Removed [compute.PolicyFamilyId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#PolicyFamilyId).
 * Removed `CanUseClient` field for [compute.ListClustersRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ListClustersRequest).
 * Removed `IsAscending` and `SortBy` fields for [marketplace.ListListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListListingsRequest).
 * Removed `ProviderSummary` field for [marketplace.Listing](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#Listing).
 * Removed `Filters` field for [marketplace.ListingSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListingSetting).
 * Removed `MetastoreId` field for [marketplace.ListingSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListingSummary).
 * Removed `IsAscending` and `SortBy` fields for [marketplace.SearchListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#SearchListingsRequest).
 * Removed `CreatedAt`, `LastTriggeredAt`, `Name`, `Options`, `Parent`, `Query`, `Rearm`, `UpdatedAt` and `User` fields for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert).
 * Removed `AlertId` field for [sql.GetAlertRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetAlertRequest).
 * Removed `QueryId` field for [sql.GetQueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetQueryRequest).
 * Removed `Order`, `Page` and `Q` fields for [sql.ListQueriesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueriesRequest).
 * Removed `IncludeMetrics` field for [sql.ListQueryHistoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ListQueryHistoryRequest).
 * Removed `CanEdit`, `CreatedAt`, `DataSourceId`, `IsArchived`, `IsDraft`, `IsFavorite`, `IsSafe`, `LastModifiedBy`, `LastModifiedById`, `LatestQueryDataId`, `Name`, `Options`, `Parent`, `PermissionTier`, `Query`, `QueryHash`, `RunAsRole`, `UpdatedAt`, `User`, `UserId` and `Visualizations` fields for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query).
 * Removed `StatementIds` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Removed `CanSubscribeToLiveQuery` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
 * Removed `MetadataTimeMs`, `PlanningTimeMs` and `QueryExecutionTimeMs` fields for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `CreatedAt`, `Description`, `Name`, `Options`, `Query` and `UpdatedAt` fields for [sql.Visualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Visualization).

OpenAPI SHA: f98c07f9c71f579de65d2587bb0292f83d10e55d, Date: 2024-08-12

## Release v0.43.2 (2024-07-09)

### Internal Changes
 * Enforce Tag on PRs ([#969](https://github.com/databricks/databricks-sdk-go/pull/969)).
 * Generate SDK for `apierr` changes ([#970](https://github.com/databricks/databricks-sdk-go/pull/970)).
 * Add Release tag and Workflow Fix ([#972](https://github.com/databricks/databricks-sdk-go/pull/972)).



## Release v0.43.1 (2024-07-08)

### Major Changes and Improvements:
 * Add a credentials provider for Github Azure OIDC ([#965](https://github.com/databricks/databricks-sdk-go/pull/965)).
 * Add DataPlane API Support ([#936](https://github.com/databricks/databricks-sdk-go/pull/936)).
 * Added more error messages for retriable errors (timeouts, etc.) ([#963](https://github.com/databricks/databricks-sdk-go/pull/963)).

### Internal Changes
 * Add ChangelogConfig to Generator struct ([#967](https://github.com/databricks/databricks-sdk-go/pull/967)).
 * Improve Changelog by grouping changes ([#962](https://github.com/databricks/databricks-sdk-go/pull/962)).
 * Parse API Error messages with `int` error codes ([#960](https://github.com/databricks/databricks-sdk-go/pull/960)).


## Release v0.43.0 (2024-06-24)

Major Changes and Improvements:
 * Support partners in user agent for SDK ([#925](https://github.com/databricks/databricks-sdk-go/pull/925)).
 * Add `serverless_compute_id` field to the config ([#952](https://github.com/databricks/databricks-sdk-go/pull/952)).

Other Changes:

 * Generate from latest spec ([#944](https://github.com/databricks/databricks-sdk-go/pull/944)) and ([#947](https://github.com/databricks/databricks-sdk-go/pull/947)).

API Changes:

 * Changed `IsolationMode` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo) to [catalog.CatalogIsolationMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogIsolationMode).
 * Added `IsolationMode` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `MaxResults` and `PageToken` fields for [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Added `NextPageToken` field for [catalog.ListCatalogsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsResponse).
 * Added `TableServingUrl` field for [catalog.OnlineTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTable).
 * Added `IsolationMode` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo).
 * Changed `IsolationMode` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog) to [catalog.CatalogIsolationMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogIsolationMode).
 * Added `IsolationMode` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `IsolationMode` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential).
 * Added [catalog.CatalogIsolationMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogIsolationMode).
 * Added `CreateSchedule`, `CreateSubscription`, `DeleteSchedule`, `DeleteSubscription`, `GetSchedule`, `GetSubscription`, `List`, `ListSchedules`, `ListSubscriptions` and `UpdateSchedule` methods for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateScheduleRequest), [dashboards.CreateSubscriptionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateSubscriptionRequest), [dashboards.CronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CronSchedule), [dashboards.DashboardView](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#DashboardView), [dashboards.DeleteScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#DeleteScheduleRequest), [dashboards.DeleteSubscriptionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#DeleteSubscriptionRequest), [dashboards.GetScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetScheduleRequest), [dashboards.GetSubscriptionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetSubscriptionRequest), [dashboards.ListDashboardsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListDashboardsRequest), [dashboards.ListDashboardsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListDashboardsResponse), [dashboards.ListSchedulesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListSchedulesRequest), [dashboards.ListSchedulesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListSchedulesResponse), [dashboards.ListSubscriptionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListSubscriptionsRequest), [dashboards.ListSubscriptionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#ListSubscriptionsResponse), [dashboards.Schedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Schedule), [dashboards.SchedulePauseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#SchedulePauseStatus), [dashboards.Subscriber](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Subscriber), [dashboards.Subscription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Subscription), [dashboards.SubscriptionSubscriberDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#SubscriptionSubscriberDestination), [dashboards.SubscriptionSubscriberUser](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#SubscriptionSubscriberUser) and [dashboards.UpdateScheduleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateScheduleRequest) structs.
 * Added `OnStreamingBacklogExceeded` field for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
 * Added `EnvironmentKey` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Removed `ConditionTask`, `DbtTask`, `NotebookTask`, `PipelineTask`, `PythonWheelTask`, `RunJobTask`, `SparkJarTask`, `SparkPythonTask`, `SparkSubmitTask`, `SqlTask` and `Environments` fields for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `DbtTask` and `EnvironmentKey` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `OnStreamingBacklogExceeded` field for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
 * Added `Periodic` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
 * Added `OnStreamingBacklogExceeded` field for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
 * Added [jobs.PeriodicTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PeriodicTriggerConfiguration).
 * Added [jobs.PeriodicTriggerConfigurationTimeUnit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PeriodicTriggerConfigurationTimeUnit).
 * Added `ProviderSummary` field for [marketplace.Listing](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#Listing).
 * Added [marketplace.ProviderIconFile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderIconFile).
 * Added [marketplace.ProviderIconType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderIconType).
 * Added [marketplace.ProviderListingSummaryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ProviderListingSummaryInfo).
 * Added `Start` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added [w.ServingEndpointsDataPlane](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsDataPlaneAPI) workspace-level service.
 * Added `ServicePrincipalId` field for [serving.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#App).
 * Added `ServicePrincipalName` field for [serving.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#App).
 * Added [serving.StartAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#StartAppRequest).
 * Added `QueryNextPage` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service.
 * Added `QueryType` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).
 * Added `NextPageToken` field for [vectorsearch.QueryVectorIndexResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexResponse).
 * Added [vectorsearch.QueryVectorIndexNextPageRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexNextPageRequest).

OpenAPI SHA: 7437dabb9dadee402c1fc060df4c1ce8cc5369f0, Date: 2024-06-25
## Release v0.42.0 (2024-06-03)

* Ignore additional flaky test ([#930](https://github.com/databricks/databricks-sdk-go/pull/930)).
* Ignore DataPlane Services during generation ([#933](https://github.com/databricks/databricks-sdk-go/pull/933)).
* Update OpenAPI spec ([#934](https://github.com/databricks/databricks-sdk-go/pull/934)).

API Changes:

 * Changed `List` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.ListAccountStorageCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountStorageCredentialsResponse).
 * Added [catalog.ListAccountStorageCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountStorageCredentialsResponse).
 * Added `TerminationCategory` field for [jobs.ForEachTaskErrorMessageStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskErrorMessageStats).
 * Added [oauth2.DataPlaneInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#DataPlaneInfo).
 * Removed `CreateDeployment` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Deploy` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Mode` field for [serving.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeployment).
 * Added `Mode` field for [serving.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppDeploymentRequest).
 * Added `DataPlaneInfo` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.AppDeploymentMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentMode).
 * Added [serving.ModelDataPlaneInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ModelDataPlaneInfo).

OpenAPI SHA: 37b925eba37dfb3d7e05b6ba2d458454ce62d3a0, Date: 2024-06-03

Dependency updates:

 * Bump golang.org/x/mod from 0.16.0 to 0.17.0 ([#879](https://github.com/databricks/databricks-sdk-go/pull/879)).
 * Bump golang.org/x/oauth2 from 0.18.0 to 0.20.0 ([#911](https://github.com/databricks/databricks-sdk-go/pull/911)).
 * Bump golang.org/x/net from 0.24.0 to 0.25.0 ([#912](https://github.com/databricks/databricks-sdk-go/pull/912)).
 * Bump google.golang.org/api from 0.169.0 to 0.182.0 ([#932](https://github.com/databricks/databricks-sdk-go/pull/932)).

## Release v0.41.0 (2024-05-21)

### Backward incompatible changes
* Renamed `CredentialsProvider` to `CredentialsStrategy`.

### Improvements and new features

* Create a method to generate OAuth tokens ([#886](https://github.com/databricks/databricks-sdk-go/pull/886)).
* Better error message when private link enabled workspaces reject requests ([#924](https://github.com/databricks/databricks-sdk-go/pull/924)).
* Update OpenAPI spec ([#926](https://github.com/databricks/databricks-sdk-go/pull/926)).

### API Changes:

 * Changed `List` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service to require request of [catalog.ListConnectionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListConnectionsRequest).
 * Renamed [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service to [w.QualityMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#QualityMonitorsAPI).
 * Renamed [catalog.DeleteLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteLakehouseMonitorRequest).
 * Changed `SchemaName` field for [catalog.DisableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DisableRequest) to `string`.
 * Removed [catalog.DisableSchemaName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DisableSchemaName) to [catalog.DeleteQualityMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteQualityMonitorRequest).
 * Changed `SchemaName` field for [catalog.EnableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableRequest) to `string`.
 * Removed [catalog.EnableSchemaName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableSchemaName).
 * Renamed [catalog.GetLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetLakehouseMonitorRequest) to [catalog.GetQualityMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetQualityMonitorRequest).
 * Added `NextPageToken` field for [catalog.ListConnectionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListConnectionsResponse).
 * Added `DashboardId` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Added [catalog.ListConnectionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListConnectionsRequest).
 * Added [catalog.MonitorRefreshListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshListResponse).
 * Changed `ClusterStatus` method for [w.Libraries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibrariesAPI) workspace-level service to return [compute.ClusterLibraryStatuses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterLibraryStatuses).
 * Removed `ClusterSource` field for [compute.ClusterAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAttributes).
 * Changed `Spec` field for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails) to [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
 * Removed `CloneFrom` and `ClusterSource` fields for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
 * Removed [compute.ClusterStatusResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatusResponse).
 * Removed `ClusterSource` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Removed `CloneFrom` and `ClusterSource` fields for [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
 * Rename `SortBySpec` field to `SortBy` for [marketplace.ListListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListListingsRequest).
 * Added `IsAscending` field for [marketplace.ListListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#ListListingsRequest).
 * Added `IsAscending` field for [marketplace.SearchListingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#SearchListingsRequest).
 * Removed [marketplace.SortBySpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#SortBySpec).
 * Removed [marketplace.SortOrder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace#SortOrder).
 * Added `GatewayDefinition` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `GatewayDefinition` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `TableConfiguration` field for [pipelines.ManagedIngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ManagedIngestionPipelineDefinition).
 * Added `GatewayDefinition` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added `TableConfiguration` field for [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec).
 * Added `TableConfiguration` field for [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec).
 * Added [pipelines.IngestionGatewayPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionGatewayPipelineDefinition).
 * Added [pipelines.TableSpecificConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfig).
 * Added [pipelines.TableSpecificConfigScdType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpecificConfigScdType).
 * Added `DeploymentArtifacts` field for [serving.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeployment).
 * Added `RouteOptimized` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Added `Contents` field for [serving.ExportMetricsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExportMetricsResponse).
 * Changed `OpenaiApiKey` field for [serving.OpenAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#OpenAiConfig) to no longer be required.
 * Added `MicrosoftEntraClientId`, `MicrosoftEntraClientSecret` and `MicrosoftEntraTenantId` fields for [serving.OpenAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#OpenAiConfig).
 * Added `EndpointUrl` and `RouteOptimized` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.AppDeploymentArtifacts](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentArtifacts).
 * Added `StorageRoot` field for [sharing.CreateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateShare).
 * Added `StorageLocation` and `StorageRoot` field for [sharing.ShareInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ShareInfo).
 * Added `StorageRoot` field for [sharing.UpdateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateShare).
 * Added `ScanIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service.
 * Added `EmbeddingWritebackTable` field for [vectorsearch.DeltaSyncVectorIndexSpecRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecRequest).
 * Added `EmbeddingWritebackTable` field for [vectorsearch.DeltaSyncVectorIndexSpecResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeltaSyncVectorIndexSpecResponse).
 * Added [vectorsearch.ListValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ListValue).
 * Added [vectorsearch.MapStringValueEntry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#MapStringValueEntry).
 * Added [vectorsearch.ScanVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ScanVectorIndexRequest).
 * Added [vectorsearch.ScanVectorIndexResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#ScanVectorIndexResponse).
 * Added [vectorsearch.Struct](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#Struct).
 * Added [vectorsearch.Value](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#Value).

OpenAPI SHA: 7eb5ad9a2ed3e3f1055968a2d1014ac92c06fe92, Date: 2024-05-21

## Release v0.40.1 (2024-05-14)

* Fixed codecov for repository ([#909](https://github.com/databricks/databricks-sdk-go/pull/909)).
* Add traceparent header to enable distributed tracing. ([#914](https://github.com/databricks/databricks-sdk-go/pull/914)).
* Log cancelled and failed requests ([#919](https://github.com/databricks/databricks-sdk-go/pull/919)).

Dependency updates:

 * Bump golang.org/x/net from 0.22.0 to 0.24.0 ([#884](https://github.com/databricks/databricks-sdk-go/pull/884)).
 * Bump golang.org/x/net from 0.17.0 to 0.23.0 in /examples/zerolog ([#896](https://github.com/databricks/databricks-sdk-go/pull/896)).
 * Bump golang.org/x/net from 0.21.0 to 0.23.0 in /examples/slog ([#897](https://github.com/databricks/databricks-sdk-go/pull/897)).

## Release v0.40.0 (2024-05-02)

* Allow unlimited timeouts in retries ([#904](https://github.com/databricks/databricks-sdk-go/pull/904)). By setting RETRY_TIMEOUT_SECONDS to a negative value, WorkspaceClient and AccountClient will retry retriable failures indefinitely. As a reminder, without setting this parameter, the default retry timeout is 5 minutes.

API Changes:

 * Changed `Create` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service . New request type is [serving.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppRequest).
 * Changed `Create` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service to return [serving.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#App).
 * Removed `DeleteApp` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed `GetApp` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed `GetAppDeploymentStatus` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed `GetApps` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed `GetEvents` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `CreateDeployment` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Delete` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Get` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `GetDeployment` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `GetEnvironment` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `List` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `ListDeployments` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Stop` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added `Update` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed [serving.AppEvents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppEvents).
 * Removed [serving.AppManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppManifest).
 * Removed [serving.AppServiceStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppServiceStatus).
 * Removed [serving.DeleteAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeleteAppResponse).
 * Removed [serving.DeployAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeployAppRequest).
 * Removed [serving.DeploymentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeploymentStatus).
 * Removed [serving.DeploymentStatusState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeploymentStatusState).
 * Removed [serving.GetAppDeploymentStatusRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppDeploymentStatusRequest).
 * Removed [serving.GetAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppResponse).
 * Removed [serving.GetEventsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetEventsRequest).
 * Removed [serving.ListAppEventsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppEventsResponse).
 * Changed `Apps` field for [serving.ListAppsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsResponse) to [serving.AppList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppList).
 * Added [serving.App](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#App).
 * Added [serving.AppDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeployment).
 * Added [serving.AppDeploymentState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentState).
 * Added [serving.AppDeploymentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppDeploymentStatus).
 * Added [serving.AppEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppEnvironment).
 * Added [serving.AppState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppState).
 * Added [serving.AppStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppStatus).
 * Added [serving.CreateAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppDeploymentRequest).
 * Added [serving.CreateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateAppRequest).
 * Added [serving.EnvVariable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EnvVariable).
 * Added [serving.GetAppDeploymentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppDeploymentRequest).
 * Added [serving.GetAppEnvironmentRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppEnvironmentRequest).
 * Added [serving.ListAppDeploymentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppDeploymentsRequest).
 * Added [serving.ListAppDeploymentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppDeploymentsResponse).
 * Added [serving.ListAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsRequest).
 * Added [serving.StopAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#StopAppRequest).
 * Added `any`.
 * Added [serving.UpdateAppRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#UpdateAppRequest).
 * Removed [w.CspEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAPI) workspace-level service.
 * Removed [w.EsmEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAPI) workspace-level service.
 * Added [w.ComplianceSecurityProfile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceSecurityProfileAPI) workspace-level service.
 * Added [w.EnhancedSecurityMonitoring](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnhancedSecurityMonitoringAPI) workspace-level service.
 * Removed [settings.CspEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablement).
 * Removed [settings.CspEnablementSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementSetting).
 * Removed [settings.EsmEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablement).
 * Removed [settings.EsmEnablementSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementSetting).
 * Removed [settings.GetCspEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementSettingRequest).
 * Removed [settings.GetEsmEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementSettingRequest).
 * Removed [settings.UpdateCspEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateCspEnablementSettingRequest).
 * Removed [settings.UpdateEsmEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEsmEnablementSettingRequest).
 * Added [settings.ComplianceSecurityProfile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceSecurityProfile).
 * Added [settings.ComplianceSecurityProfileSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceSecurityProfileSetting).
 * Added [settings.EnhancedSecurityMonitoring](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnhancedSecurityMonitoring).
 * Added [settings.EnhancedSecurityMonitoringSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EnhancedSecurityMonitoringSetting).
 * Added [settings.GetComplianceSecurityProfileSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetComplianceSecurityProfileSettingRequest).
 * Added [settings.GetEnhancedSecurityMonitoringSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEnhancedSecurityMonitoringSettingRequest).
 * Added [settings.UpdateComplianceSecurityProfileSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateComplianceSecurityProfileSettingRequest).
 * Added [settings.UpdateEnhancedSecurityMonitoringSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEnhancedSecurityMonitoringSettingRequest).
 * Added `Tags` field for [sql.DashboardEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardEditContent).
 * Added `Tags` field for [sql.QueryEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryEditContent).
 * Added `Catalog` field for [sql.QueryOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryOptions).
 * Added `Schema` field for [sql.QueryOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryOptions).
 * Added `Tags` field for [sql.QueryPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryPostContent).
 * Added `Query` field for [sql.Visualization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Visualization).

OpenAPI SHA: master, Date: 2024-05-02
Dependency updates:

 * Bump version of mockery ([#907](https://github.com/databricks/databricks-sdk-go/pull/907)).

## Release v0.39.0 (2024-04-24)

* Ignored flaky integration tests ([#894](https://github.com/databricks/databricks-sdk-go/pull/894)).
* Added retries for "worker env WorkerEnvId(workerenv-XXXXX) not found" ([#890](https://github.com/databricks/databricks-sdk-go/pull/890)).
* Updated SDK to OpenAPI spec ([#899](https://github.com/databricks/databricks-sdk-go/pull/899)).

Note: This release contains breaking changes, please see the API changes below for more details.

API Changes:

 * Added `IngestionDefinition` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline), [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline) and [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added `Deployment` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline), [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline) and [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added [compute.ClusterStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatus).
 * Added [compute.ClusterStatusResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatusResponse).
 * Added [compute.LibraryInstallStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibraryInstallStatus).
 * Added `WarehouseId` field for [jobs.NotebookTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#NotebookTask).
 * Added `RunAs` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added [pipelines.DeploymentKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#DeploymentKind).
 * Added [pipelines.IngestionConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#IngestionConfig).
 * Added [pipelines.ManagedIngestionPipelineDefinition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ManagedIngestionPipelineDefinition).
 * Added [pipelines.PipelineDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineDeployment).
 * Added [pipelines.SchemaSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#SchemaSpec).
 * Added [pipelines.TableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#TableSpec).
 * Added `GetOpenApi` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added [serving.GetOpenApiRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetOpenApiRequest).
 * Added `SchemaId` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Added `Operation` field for [catalog.ValidationResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResult).
 * Added [catalog.ValidationResultOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultOperation).
 * Added `Requirements` field for [compute.Library](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Library).
 * Removed `AwsOperation` field for [catalog.ValidationResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResult).
 * Removed `AzureOperation` field for [catalog.ValidationResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResult).
 * Removed `GcpOperation` field for [catalog.ValidationResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResult).
 * Removed [catalog.ValidationResultAwsOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultAwsOperation).
 * Removed [catalog.ValidationResultAzureOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultAzureOperation).
 * Removed [catalog.ValidationResultGcpOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultGcpOperation).
 * Removed [compute.ClusterStatusRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatusRequest).
 * Removed [compute.LibraryFullStatusStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibraryFullStatusStatus).
 * Changed `ClusterStatus` method for [w.Libraries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibrariesAPI) workspace-level service . New request type is [compute.ClusterStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatus).
 * Changed `ClusterStatus` method for [w.Libraries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibrariesAPI) workspace-level service to return [compute.ClusterStatusResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterStatusResponse).
 * Changed `Status` field for [compute.LibraryFullStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibraryFullStatus) to [compute.LibraryInstallStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LibraryInstallStatus).


OpenAPI SHA: 21f9f1482f9d0d15228da59f2cd9f0863d2a6d55, Date: 2024-04-23

## Release v0.38.0 (2024-04-11)

### Behavior Changes
* Override INVALID_PARAMETER_VALUE on fetching non-existent job/cluster ([#864](https://github.com/databricks/databricks-sdk-go/pull/864)). The error returned when fetching a non-existent job or cluster has been changed from `INVALID_PARAMETER_VALUE` to `RESOURCE_DOES_NOT_EXIST`. Update your error handling code to check for `databricks.ErrResourceDoesNotExist` instead of `databricks.ErrInvalidParameterValue`. For example, if you are using the `Jobs.GetById` method, you should update your error handling code to:
```go
_, err := w.Jobs.GetById(ctx, "id")
if errors.Is(err, databricks.ErrResourceDoesNotExist) {
    // handle the error
}
```
Note that the original error code is still accessible in the `ErrorCode` field of `APIError`.

### Other Improvements
* Do not leak secondary authorization tokens in debug logs ([#882](https://github.com/databricks/databricks-sdk-go/pull/882)).
* Fix logging of request bodies containing percent characters. ([#881](https://github.com/databricks/databricks-sdk-go/pull/881)).
* Added clientId and clientSecret to oauth-m2m auth_types ([#885](https://github.com/databricks/databricks-sdk-go/pull/885)).

### Internal Changes
* Support custom AuthVisitors ([#874](https://github.com/databricks/databricks-sdk-go/pull/874)).

API Changes:

 * Replaced [catalog.AzureManagedIdentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentity) with [catalog.AzureManagedIdentityRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentityRequest) and [catalog.AzureManagedIdentityResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentityResponse).
 * Renamed `FullName` field  to `TableName` for [catalog.CancelRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshRequest), [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor), [catalog.DeleteLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteLakehouseMonitorRequest), [catalog.GetLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetLakehouseMonitorRequest), [catalog.GetRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRefreshRequest), [catalog.ListRefreshesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRefreshesRequest).
 * Changed `CustomMetrics` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorMetricList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorMetricList).
 * Changed `InferenceLog` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorInferenceLog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLog).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotifications).
 * Changed `Snapshot` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to `any`.
 * Changed `TimeSeries` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorTimeSeries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeries).
 * Changed `QuartzCronExpression` field for [catalog.MonitorCronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedule) to be required.
 * Changed `TimezoneId` field for [catalog.MonitorCronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedule) to be required.
 * Renamed [catalog.MonitorCustomMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetric) to [catalog.MonitorMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorMetric).
 * Renamed [catalog.MonitorCustomMetricType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetricType) to [catalog.MonitorMetricType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorMetricType).
 * Renamed [catalog.MonitorDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDestinations) to [catalog.MonitorDestination](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDestination).
 * Renamed [catalog.MonitorInferenceLogProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileType) to [catalog.MonitorInferenceLog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLog).
 * Renamed [catalog.MonitorInferenceLogProfileTypeProblemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileTypeProblemType) to [catalog.MonitorInferenceLogProblemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProblemType).
 * Renamed [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig) to [catalog.MonitorNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotifications).
 * Changed `CustomMetrics` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorMetricList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorMetricList).
 * Changed `DriftMetricsTableName` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
 * Changed `InferenceLog` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorInferenceLog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLog).
 * Changed `MonitorVersion` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotifications).
 * Changed `ProfileMetricsTableName` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
 * Changed `Snapshot` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to `any`.
 * Changed `Status` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
 * Changed `TableName` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to be required.
 * Changed `TimeSeries` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorTimeSeries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeries).
 * Changed `RefreshId` field for [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo) to be required.
 * Changed `StartTimeMs` field for [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo) to be required.
 * Changed `State` field for [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo) to be required.
 * Added `Trigger` field for [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo).
 * Removed [catalog.MonitorTimeSeriesProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeriesProfileType).
 * Removed `FullName` field for [catalog.RunRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RunRefreshRequest).
 * Added `TableName` field for [catalog.RunRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RunRefreshRequest).
 * Changed `AzureManagedIdentity` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to .
 * Removed `Name` field for [catalog.TableRowFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableRowFilter).
 * Added `FunctionName` field for [catalog.TableRowFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableRowFilter).
 * Changed `CustomMetrics` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorMetricList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorMetricList).
 * Removed `FullName` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Changed `InferenceLog` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorInferenceLog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLog).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotifications).
 * Changed `Snapshot` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to `any`.
 * Changed `TimeSeries` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorTimeSeries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeries).
 * Added `TableName` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Changed `AzureManagedIdentity` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AzureManagedIdentityResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentityResponse).
 * Changed `AzureManagedIdentity` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AzureManagedIdentityRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AzureManagedIdentityRequest).
 * Replaced `Operation` field for [catalog.ValidationResult](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResult) with `AwsOperation`, `AzureOperation` and `GcpOperation`.
 * Replaced [catalog.ValidationResultOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultOperation) with [catalog.ValidationResultAwsOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultAwsOperation), [catalog.ValidationResultAzureOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultAzureOperation) and [catalog.ValidationResultGcpOperation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidationResultGcpOperation).
 * Added [catalog.MonitorRefreshInfoTrigger](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoTrigger).
 * Added [catalog.MonitorTimeSeries](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeries).
 * Removed [compute.ComputeSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ComputeSpec).
 * Removed [compute.ComputeSpecKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ComputeSpecKind).
 * Added `CloneFrom` field to [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster), [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec), and [compute.EditCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditCluster).
 * Added [compute.CloneCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CloneCluster).
 * Added [compute.Environment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Environment).
 * Changed `Update` method for [a.WorkspaceAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#WorkspaceAssignmentAPI) account-level service to return [iam.PermissionAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionAssignment).
 * Removed `ComputeKey` field for [jobs.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ClusterSpec).
 * Removed `Compute` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Environments` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Removed [jobs.JobCompute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCompute).
 * Removed `Compute` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `Environments` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Removed `ComputeKey` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Removed [jobs.TableTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableTriggerConfiguration).
 * Removed `ComputeKey` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `EnvironmentKey` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Changed `Table` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings) to [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration).
 * Changed `TableUpdate` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings) to [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration).
 * Added [jobs.JobEnvironment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEnvironment).
 * Added [jobs.TableUpdateTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableUpdateTriggerConfiguration).
 * Added [marketplace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/marketplace) package.

OpenAPI SHA: 94684175b8bd65f8701f89729351f8069e8309c9, Date: 2024-04-11

## ## 0.37.0

### Other Changes

* Fix integration test naming for UC Files ([#868](https://github.com/databricks/databricks-sdk-go/pull/868)).

### API Changes

 * Added `Migrate` and `Unpublish` methods for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.MigrateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#MigrateDashboardRequest).
 * Added [dashboards.UnpublishDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UnpublishDashboardRequest).
 * Added `Description`, `QueueDuration` and `RepairHistory` fields for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `ComputeKey` and `JobClusterKey` fields for [jobs.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ClusterSpec).
 * Changed `Left`, `Op` and `Right` fields for [jobs.ConditionTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ConditionTask) to be required.
 * Changed `EditMode` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) to [jobs.JobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEditMode).
 * Replaced [jobs.CreateJobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobEditMode) to [jobs.JobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEditMode).
 * Changed `Url` field for [jobs.FileArrivalTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#FileArrivalTriggerConfiguration) to be required.
 * Changed `ErrorMessageStats` field for [jobs.ForEachStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachStats) to [jobs.ForEachTaskErrorMessageStatsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskErrorMessageStatsList).
 * Changed `NewCluster` field for [jobs.JobCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobCluster) to be required.
 * Changed `EditMode` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings) to [jobs.JobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEditMode).
 * Replaced [jobs.JobSettingsEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsEditMode) by [jobs.JobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEditMode).
 * Changed `Metric`, `Op` and `Value` fields for [jobs.JobsHealthRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRule) to be required.
 * Changed `RunType` field for [jobs.ListRunsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ListRunsRequest) to [jobs.RunType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunType).
 * Removed [jobs.ListRunsRunType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ListRunsRunType).
 * Removed [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Changed `PipelineId` field for [jobs.PipelineTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PipelineTask) to be required.
 * Changed `EntryPoint` and `PackageName` fields for [jobs.PythonWheelTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#PythonWheelTask) to be required.
 * Changed `JobParameters` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun) to map[string]`string`.
 * Changed `BaseParameters` field for [jobs.ResolvedNotebookTaskValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedNotebookTaskValues) to map[string]`string`.
 * Changed `Parameters` field for [jobs.ResolvedParamPairValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedParamPairValues) to map[string]`string`.
 * Changed `NamedParameters` field for [jobs.ResolvedPythonWheelTaskValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedPythonWheelTaskValues) to map[string]`string`.
 * Removed `NamedParameters` field for [jobs.ResolvedRunJobTaskValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedRunJobTaskValues).
 * Changed `Parameters` field for [jobs.ResolvedRunJobTaskValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedRunJobTaskValues) to map[string]`string`.
 * Added `JobParameters` field for [jobs.ResolvedRunJobTaskValues](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ResolvedRunJobTaskValues).
 * Added `Description` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `QueueDuration` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Changed `Op` field for [jobs.RunConditionTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunConditionTask) to [jobs.ConditionTaskOp](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ConditionTaskOp).
 * Removed [jobs.RunConditionTaskOp](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunConditionTaskOp).
 * Changed `Inputs` and `Task` fields for [jobs.RunForEachTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunForEachTask) to be required.
 * Changed `JobParameters` field for [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask) to map[string]`string`.
 * Added `DbtCommands`, `JarParams`, `NotebookParams`, `PipelineParams`, `PythonNamedParams`, `PythonParams`, `SparkSubmitParams` and `SqlParams` fields for [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask).
 * Changed `JobParameters` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow) to map[string]`string`.
 * Added `Info` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Removed `JobParameters` field for [jobs.RunParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunParameters).
 * Changed `TaskKey` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask) to be required.
 * Added `ComputeKey`,`EmailNotifications`, `JobClusterKey`, `NotificatioSettings`, `RunDuration`, `RunPageUrl`, `TimeoutSeconds` and `WebhookNotifications` fields for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `EndpointId` field for [jobs.SqlQueryOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlQueryOutput).
 * Added `ConditionTask` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `DbtCommands`, `JarParams`, `NotebookParams`, `PipelineParams`, `PythonNamedParams`, `PythonParams`, `SparkSubmitParams` and `SqlParams` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `Description` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `DisableAutoOptimization` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `NoAlertForSkippedRuns` field for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
 * Added `TableUpdate` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
 * Changed `Id` field for [jobs.Webhook](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Webhook) to be required.
 * Changed `OnDurationWarningThresholdExceeded` field for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications) to [jobs.WebhookList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookList).
 * Removed [jobs.WebhookNotificationsOnDurationWarningThresholdExceededItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotificationsOnDurationWarningThresholdExceededItem).
 * Added [jobs.JobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEditMode).
 * Removed [serving.AwsBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfig).
 * Removed [serving.AwsBedrockConfigBedrockProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfigBedrockProvider).
 * Removed `AwsBedrockConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `AmazonBedrockConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added [serving.AmazonBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfig).
 * Added [serving.AmazonBedrockConfigBedrockProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AmazonBedrockConfigBedrockProvider).
 * Changed `Get` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service . New request type is [settings.GetIpAccessListRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListRequest).
 * Renamed [settings.GetIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessList) to [settings.GetIpAccessListRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListRequest).

OpenAPI SHA: d38528c3e47dd81c9bdbd918272a3e49d36e09ce, Date: 2024-03-27

## Release v0.36.0 (2024-03-20)


API Changes:

 * Changed `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service . New request type is [dashboards.GetDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetDashboardRequest).
 * Changed `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service . New request type is [dashboards.GetPublishedDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedDashboardRequest).
 * Changed `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service . New request type is [dashboards.TrashDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashDashboardRequest).
 * Removed [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Removed [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Removed [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.GetDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetDashboardRequest).
 * Added [dashboards.GetPublishedDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedDashboardRequest).
 * Added [dashboards.TrashDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashDashboardRequest).
 * Added `AutoCaptureConfig` field for [serving.EndpointPendingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointPendingConfig).
 * Changed `Get` method for [w.AutomaticClusterUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AutomaticClusterUpdateAPI) workspace-level service . New request type is [settings.GetAutomaticClusterUpdateSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetAutomaticClusterUpdateSettingRequest).
 * Changed `Get` method for [w.CspEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAPI) workspace-level service . New request type is [settings.GetCspEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementSettingRequest).
 * Changed `Get` method for [a.CspEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAccountAPI) account-level service . New request type is [settings.GetCspEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementAccountSettingRequest).
 * Changed `Delete` method for [w.DefaultNamespace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceAPI) workspace-level service . New request type is [settings.DeleteDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingRequest).
 * Changed `Get` method for [w.DefaultNamespace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceAPI) workspace-level service . New request type is [settings.GetDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceSettingRequest).
 * Changed `Get` method for [w.EsmEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAPI) workspace-level service . New request type is [settings.GetEsmEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementSettingRequest).
 * Changed `Get` method for [a.EsmEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAccountAPI) account-level service . New request type is [settings.GetEsmEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementAccountSettingRequest).
 * Changed `Get` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service . New request type is [settings.GetIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessList).
 * Changed `Delete` method for [a.PersonalCompute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PersonalComputeAPI) account-level service . New request type is [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest).
 * Changed `Get` method for [a.PersonalCompute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PersonalComputeAPI) account-level service . New request type is [settings.GetPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeSettingRequest).
 * Changed `Delete` method for [w.RestrictWorkspaceAdmins](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsAPI) workspace-level service . New request type is [settings.DeleteRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingRequest).
 * Changed `Get` method for [w.RestrictWorkspaceAdmins](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsAPI) workspace-level service . New request type is [settings.GetRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminsSettingRequest).
 * Removed [settings.DeleteDefaultNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceRequest).
 * Removed [settings.DeletePersonalComputeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeRequest).
 * Removed [settings.DeleteRestrictWorkspaceAdminRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminRequest).
 * Removed [settings.GetAutomaticClusterUpdateRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetAutomaticClusterUpdateRequest).
 * Removed [settings.GetCspEnablementAccountRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementAccountRequest).
 * Removed [settings.GetCspEnablementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementRequest).
 * Removed [settings.GetDefaultNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceRequest).
 * Removed [settings.GetEsmEnablementAccountRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementAccountRequest).
 * Removed [settings.GetEsmEnablementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementRequest).
 * Removed [settings.GetIpAccessListRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListRequest).
 * Removed [settings.GetPersonalComputeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeRequest).
 * Removed [settings.GetRestrictWorkspaceAdminRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminRequest).
 * Added [settings.DeleteDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingRequest).
 * Added [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest).
 * Added [settings.DeleteRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingRequest).
 * Added [settings.GetAutomaticClusterUpdateSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetAutomaticClusterUpdateSettingRequest).
 * Added [settings.GetCspEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementAccountSettingRequest).
 * Added [settings.GetCspEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementSettingRequest).
 * Added [settings.GetDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceSettingRequest).
 * Added [settings.GetEsmEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementAccountSettingRequest).
 * Added [settings.GetEsmEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementSettingRequest).
 * Added [settings.GetIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessList).
 * Added [settings.GetPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeSettingRequest).
 * Added [settings.GetRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminsSettingRequest).
 * Changed `DataObjectType` field for [sharing.SharedDataObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObject) to [sharing.SharedDataObjectDataObjectType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObjectDataObjectType).
 * Added `Content` field for [sharing.SharedDataObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObject).
 * Added [sharing.SharedDataObjectDataObjectType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObjectDataObjectType).
 * Added `EmbeddingSourceColumns` field for [vectorsearch.DirectAccessVectorIndexSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DirectAccessVectorIndexSpec).

OpenAPI SHA: 93763b0d7ae908520c229c786fff28b8fd623261, Date: 2024-03-20

## Release v0.35.0 (2024-03-20)

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo), [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), and [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

##
Release v0.35.0

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo), [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), and [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

## Release v0.35.0 (2024-03-20)

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo), [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), and [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).


## Release v0.35.0 (2024-03-20)

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo), [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), and [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

##
Release v0.35.0

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo), [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), and [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

## Release v0.35.0 (2024-03-20)

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).
* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).
* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added `BrowseOnly` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `BrowseOnly` field for [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest).
 * Added `IncludeBrowse` field for [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest).
 * Added `IncludeBrowse` field for [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest).
 * Added `IncludeBrowse` field for [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest).
 * Added `IncludeBrowse` field for [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest).
 * Added `IncludeBrowse` field for [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest).
 * Added `IncludeBrowse` field for [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest).
 * Added `IncludeBrowse` field for [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest).
 * Added `IncludeBrowse` field for [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest).
 * Added `IncludeBrowse` field for [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest).
 * Added `IncludeBrowse` field for [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest).
 * Added `IncludeBrowse` field for [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest).
 * Added `IncludeBrowse` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
 * Added `IncludeBrowse` field for [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest).
 * Added `BrowseOnly` field for [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Added `IncludeBrowse` field for [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo).
 * Added `BrowseOnly` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added `BrowseOnly` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added `BrowseOnly` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Removed `any`.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `any`.
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

## Release v0.35.0 (2024-03-20)

* Added Config.GetAuthDetails ([#838](https://github.com/databricks/databricks-sdk-go/pull/838)).
* Support DATABRICKS_SDK_UPSTREAM and DATABRICKS_SDK_UPSTREAM_VERSION ([#854](https://github.com/databricks/databricks-sdk-go/pull/854)).

### Test Fixes

* Fix TestUcAccShares ([#858](https://github.com/databricks/databricks-sdk-go/pull/858)).

### Internal Changes

* Add telemetry for SDK usage from DBR ([#851](https://github.com/databricks/databricks-sdk-go/pull/851)).

API Changes:

 * Changed `List` method for [w.Catalogs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogsAPI) workspace-level service to require request of [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Create` method for [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service . New request type is [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Removed [catalog.AwsIamRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRole).
 * Changed `Notifications` field for [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `AwsIamRole` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added `IncludeBrowse` field for [catalog.GetCatalogRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetCatalogRequest), [catalog.GetExternalLocationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetExternalLocationRequest), [catalog.GetFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetFunctionRequest), [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest), [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest), [catalog.GetSchemaRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetSchemaRequest), [catalog.GetTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetTableRequest), [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest), [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest), [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest), [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest), [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest), [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest), and [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `BrowseOnly` field for [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo), [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo), [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo), [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo), [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo), [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo), and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo).
 * Changed `Notifications` field for [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `AwsIamRole` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo) to [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Changed `Notifications` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor) to [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig).
 * Changed `AwsIamRole` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Changed `AwsIamRole` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential) to [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Removed [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added [catalog.AwsIamRoleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleRequest).
 * Added [catalog.AwsIamRoleResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AwsIamRoleResponse).
 * Added [catalog.CreateOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateOnlineTableRequest).
 * Added [catalog.ListCatalogsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListCatalogsRequest).
 * Changed `Publish` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service to return [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `Create` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Get` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `GetPublished` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Trash` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Added `Update` method for [w.Lakeview](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LakeviewAPI) workspace-level service.
 * Removed `any`.
 * Added [dashboards.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#CreateDashboardRequest).
 * Added [dashboards.Dashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#Dashboard).
 * Added [dashboards.GetLakeviewRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetLakeviewRequest).
 * Added [dashboards.GetPublishedRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#GetPublishedRequest).
 * Added [dashboards.LifecycleState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#LifecycleState).
 * Added [dashboards.PublishedDashboard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#PublishedDashboard).
 * Added `any`.
 * Added [dashboards.TrashRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#TrashRequest).
 * Added [dashboards.UpdateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards#UpdateDashboardRequest).
 * Added `ScoreThreshold` field for [vectorsearch.QueryVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#QueryVectorIndexRequest).

OpenAPI SHA: 3821dc51952c5cf1c276dd84967da011b191e64a, Date: 2024-03-19
Dependency updates:

 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/zerolog ([#855](https://github.com/databricks/databricks-sdk-go/pull/855)).
 * Bump google.golang.org/protobuf from 1.31.0 to 1.33.0 in /examples/slog ([#856](https://github.com/databricks/databricks-sdk-go/pull/856)).
 * Bump google.golang.org/protobuf from 1.32.0 to 1.33.0 ([#857](https://github.com/databricks/databricks-sdk-go/pull/857)).
 * Bump github.com/stretchr/testify from 1.8.4 to 1.9.0 ([#840](https://github.com/databricks/databricks-sdk-go/pull/840)).
 * Bump golang.org/x/mod from 0.15.0 to 0.16.0 ([#843](https://github.com/databricks/databricks-sdk-go/pull/843)).
 * Bump golang.org/x/oauth2 from 0.17.0 to 0.18.0 ([#845](https://github.com/databricks/databricks-sdk-go/pull/845)).
 * Bump google.golang.org/api from 0.166.0 to 0.169.0 ([#849](https://github.com/databricks/databricks-sdk-go/pull/849)).

## Release v0.34.0 (2024-03-05)

### New Features and Improvements
* Fixed GetWorkspaceClient for GCP ([#803](https://github.com/databricks/databricks-sdk-go/pull/803)).
* Adaptive request timeouts ([#837](https://github.com/databricks/databricks-sdk-go/pull/837)).
* Added HTTP proxy example ([#825](https://github.com/databricks/databricks-sdk-go/pull/825)).
* Note: Backwards incompatible changes - Settings are now nested, please see the API changes below.

### API Changes:
* Added:
  - [w.PermissionMigration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationAPI) workspace-level service.
  - [iam.PermissionMigrationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationRequest).
  - [iam.PermissionMigrationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionMigrationResponse).
  - [w.AutomaticClusterUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AutomaticClusterUpdateAPI) workspace-level service.
  - [w.CspEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAPI) workspace-level service.
  - [a.CspEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAccountAPI) account-level service.
  - [w.DefaultNamespace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceAPI) workspace-level service.
  - [w.EsmEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAPI) workspace-level service.
  - [a.EsmEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAccountAPI) account-level service.
  - [a.PersonalCompute](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PersonalComputeAPI) account-level service.
  - [w.RestrictWorkspaceAdmins](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsAPI) workspace-level service.
  - `AwsStableIpRule` field for [settings.NccEgressDefaultRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressDefaultRules).
  - [settings.AutomaticClusterUpdateSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AutomaticClusterUpdateSetting).
  - [settings.ClusterAutoRestartMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessage).
  - [settings.ClusterAutoRestartMessageEnablementDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageEnablementDetails).
  - [settings.ClusterAutoRestartMessageMaintenanceWindow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageMaintenanceWindow).
  - [settings.ClusterAutoRestartMessageMaintenanceWindowDayOfWeek](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageMaintenanceWindowDayOfWeek).
  - [settings.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule).
  - [settings.ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency).
  - [settings.ClusterAutoRestartMessageMaintenanceWindowWindowStartTime](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ClusterAutoRestartMessageMaintenanceWindowWindowStartTime).
  - [settings.ComplianceStandard](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ComplianceStandard).
  - [settings.CspEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablement).
  - [settings.CspEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAccount).
  - [settings.CspEnablementAccountSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementAccountSetting).
  - [settings.CspEnablementSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CspEnablementSetting).
  - [settings.DeleteDefaultNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceRequest).
  - [settings.DeletePersonalComputeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeRequest).
  - [settings.DeleteRestrictWorkspaceAdminRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminRequest).
  - [settings.EsmEnablement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablement).
  - [settings.EsmEnablementAccount](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAccount).
  - [settings.EsmEnablementAccountSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementAccountSetting).
  - [settings.EsmEnablementSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#EsmEnablementSetting).
  - [settings.GetAutomaticClusterUpdateRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetAutomaticClusterUpdateRequest).
  - [settings.GetCspEnablementAccountRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementAccountRequest).
  - [settings.GetCspEnablementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetCspEnablementRequest).
  - [settings.GetDefaultNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceRequest).
  - [settings.GetEsmEnablementAccountRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementAccountRequest).
  - [settings.GetEsmEnablementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetEsmEnablementRequest).
  - [settings.GetPersonalComputeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeRequest).
  - [settings.GetRestrictWorkspaceAdminRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminRequest).
  - [settings.NccAwsStableIpRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAwsStableIpRule).
  - [settings.UpdateAutomaticClusterUpdateSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAutomaticClusterUpdateSettingRequest).
  - [settings.UpdateCspEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateCspEnablementAccountSettingRequest).
  - [settings.UpdateCspEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateCspEnablementSettingRequest).
  - [settings.UpdateEsmEnablementAccountSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEsmEnablementAccountSettingRequest).
  - [settings.UpdateEsmEnablementSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateEsmEnablementSettingRequest).
  - `IndexName` field for [vectorsearch.DeleteDataVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeleteDataVectorIndexRequest).
  - `EmbeddingModelEndpointName` field for [vectorsearch.EmbeddingSourceColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EmbeddingSourceColumn).
  - `IndexName` field for [vectorsearch.UpsertDataVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#UpsertDataVectorIndexRequest).
  - `DeltaSyncIndexSpec` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
  - `DirectAccessIndexSpec` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
* Changed:
  - `Version` field for [serving.AppManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppManifest) to [serving.AnyValue](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AnyValue).
  - `DeleteEndpoint` method for [w.VectorSearchEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchEndpointsAPI) workspace-level service with new required argument order.
  - `CreateIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service with new required argument order.
  - `DeleteDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service with new required argument order.
  - `UpsertDataVectorIndex` method for [w.VectorSearchIndexes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorSearchIndexesAPI) workspace-level service with new required argument order.
  - `EndpointName` field for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest) to be required.
* Removed:
  - `DeletePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
  - `GetPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
  - `UpdatePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
  - `DeleteDefaultNamespaceSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - `DeleteRestrictWorkspaceAdminsSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - `GetDefaultNamespaceSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - `GetRestrictWorkspaceAdminsSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - `UpdateDefaultNamespaceSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - `UpdateRestrictWorkspaceAdminsSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
  - [settings.DeleteDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingRequest).
  - [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest).
  - [settings.DeleteRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingRequest).
  - [settings.GetDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceSettingRequest).
  - [settings.GetPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeSettingRequest).
  - [settings.GetRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminsSettingRequest).
  - `PlanningPhases` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
  - `Name` field for [vectorsearch.DeleteDataVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeleteDataVectorIndexRequest).
  - `Name` field for [vectorsearch.DeleteEndpointRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#DeleteEndpointRequest).
  - [vectorsearch.EmbeddingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EmbeddingConfig).
  - `EmbeddingConfig` field for [vectorsearch.EmbeddingSourceColumn](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#EmbeddingSourceColumn).
  - `Name` field for [vectorsearch.UpsertDataVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#UpsertDataVectorIndexRequest).
  - `DeltaSyncVectorIndexSpec` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).
  - `DirectAccessVectorIndexSpec` field for [vectorsearch.VectorIndex](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#VectorIndex).

### Internal Changes:
* Differentiate between empty structures and components that can take on any value ([#821](https://github.com/databricks/databricks-sdk-go/pull/821)).
* Added integration tests for files API ([#818](https://github.com/databricks/databricks-sdk-go/pull/818)).
* Added `HasByteStreamField()` helper function ([#824](https://github.com/databricks/databricks-sdk-go/pull/824)).
* Update SDK to latest OpenAPI spec ([#839](https://github.com/databricks/databricks-sdk-go/pull/839)).
* Added tokei.rs badge ([#836](https://github.com/databricks/databricks-sdk-go/pull/836)).
* Updated isObject to consider empty objects ([#834](https://github.com/databricks/databricks-sdk-go/pull/834)).
* Treat empty entities as entities ([#831](https://github.com/databricks/databricks-sdk-go/pull/831)).
* Do not transpile getWorkspaceClient ([#830](https://github.com/databricks/databricks-sdk-go/pull/830)).
* Include Go 1.22 in test matrix ([#841](https://github.com/databricks/databricks-sdk-go/pull/841)).
* Added support for subservices ([#826](https://github.com/databricks/databricks-sdk-go/pull/826)).
* Bump google.golang.org/api from 0.161.0 to 0.166.0 ([#829](https://github.com/databricks/databricks-sdk-go/pull/829)).
* Bump exp & mod library ([#832](https://github.com/databricks/databricks-sdk-go/pull/832)).

OpenAPI SHA: d855b30f25a06fe84f25214efa20e7f1fffcdf9e, Date: 2024-03-04


## Release v0.33.0 (2024-02-19)

Internal Changes:

* Add helper function to get header fields ([#822](https://github.com/databricks/databricks-sdk-go/pull/822)).
* Add Int64 to header type injection ([#819](https://github.com/databricks/databricks-sdk-go/pull/819)).

API Changes:

 * Changed `Update` method for [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service with new required argument order.
 * Added [w.OnlineTables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTablesAPI) workspace-level service.
 * Removed `AssetsDir` field for [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Added [catalog.ContinuousUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ContinuousUpdateStatus).
 * Added [catalog.DeleteOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteOnlineTableRequest).
 * Added [catalog.FailedStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FailedStatus).
 * Added [catalog.GetOnlineTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetOnlineTableRequest).
 * Added [catalog.OnlineTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTable).
 * Added [catalog.OnlineTableSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableSpec).
 * Added [catalog.OnlineTableState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableState).
 * Added [catalog.OnlineTableStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#OnlineTableStatus).
 * Added [catalog.PipelineProgress](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#PipelineProgress).
 * Added [catalog.ProvisioningStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningStatus).
 * Added [catalog.TriggeredUpdateStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TriggeredUpdateStatus).
 * Added [catalog.ViewData](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ViewData).
 * Added `ContentLength` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added `ContentType` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added `LastModified` field for [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Changed `LastModified` field for [files.GetMetadataResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataResponse) to [files.LastModifiedHttpDate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#LastModifiedHttpDate).
 * Added [files.LastModifiedHttpDate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#LastModifiedHttpDate).
 * Removed `Config` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `Ai21labsConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `AnthropicConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `AwsBedrockConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `CohereConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `DatabricksModelServingConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `OpenaiConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added `PalmConfig` field for [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Removed [serving.ExternalModelConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelConfig).
 * Added `MaxProvisionedThroughput` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added `MinProvisionedThroughput` field for [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added `MaxProvisionedThroughput` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
 * Added `MinProvisionedThroughput` field for [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).

OpenAPI SHA: cdd76a98a4fca7008572b3a94427566dd286c63b, Date: 2024-02-19

## Release v0.32.0 (2024-02-14)

Major Changes:

* Generate fields for headers in responses ([#812](https://github.com/databricks/databricks-sdk-go/pull/812)).

Other Changes:

* Add `CurrentWorkspaceID()` method on WorkspaceClient ([#808](https://github.com/databricks/databricks-sdk-go/pull/808)).
* Update OpenAPI spec ([#816](https://github.com/databricks/databricks-sdk-go/pull/816)).

Internal Changes:

* Do not generate erroneous "r" and "w" samples ([#811](https://github.com/databricks/databricks-sdk-go/pull/811)).

API Changes:

 * Changed `Delete` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Delete` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Changed `Read` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
 * Removed `NameArg` field for [catalog.DeleteConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteConnectionRequest).
 * Added `Name` field for [catalog.DeleteConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteConnectionRequest).
 * Removed `FullNameArg` field for [catalog.DeleteVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteVolumeRequest).
 * Added `Name` field for [catalog.DeleteVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteVolumeRequest).
 * Removed `NameArg` field for [catalog.GetConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetConnectionRequest).
 * Added `Name` field for [catalog.GetConnectionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetConnectionRequest).
 * Added `MaxResults` field for [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest).
 * Added `PageToken` field for [catalog.ListVolumesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesRequest).
 * Added `NextPageToken` field for [catalog.ListVolumesResponseContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListVolumesResponseContent).
 * Removed `FullNameArg` field for [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Added `Name` field for [catalog.ReadVolumeRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ReadVolumeRequest).
 * Removed `NameArg` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Removed `FullNameArg` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Added `Name` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Removed `GetStatus` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `GetDirectoryMetadata` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `GetMetadata` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added [files.FileSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FileSize).
 * Added [files.GetDirectoryMetadataRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetDirectoryMetadataRequest).
 * Added [files.GetMetadataRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataRequest).
 * Added [files.GetMetadataResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#GetMetadataResponse).
 * Removed `TriggerHistory` field for [jobs.Job](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Job).
 * Removed [jobs.TriggerEvaluation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerEvaluation).
 * Removed [jobs.TriggerHistory](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerHistory).
 * Added `Table` field for [jobs.TriggerSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerSettings).
 * Added [jobs.Condition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Condition).
 * Added [jobs.TableTriggerConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TableTriggerConfiguration).
 * Changed `Delete` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service with new required argument order.
 * Removed `NameArg` field for [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Added `Name` field for [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Removed `NameArg` field for [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Added `Name` field for [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Removed `NameArg` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Added `Name` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Added `EnumOptions` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added `MultiValuesOptions` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added `QueryId` field for [sql.Parameter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Parameter).
 * Added [sql.MultiValuesOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#MultiValuesOptions).

OpenAPI SHA: c40670f5a2055c92cf0a6aac92a5bccebfb80866, Date: 2024-02-14

## Release v0.31.0 (2024-02-08)

* Support creating a new workspace client from an account client ([#792](https://github.com/databricks/databricks-sdk-go/pull/792)). Please see the example:
```
// GetWorkspaceClient returns a WorkspaceClient for the given workspace. The
// workspace can be fetched by calling w.Workspaces.Get() or w.Workspaces.List().
//
// The config used for the workspace is identical to that used for the account,
// except that the host is set to the workspace host, and the account ID is
// not set.

a, err := databricks.NewAccountClient()
if err != nil {
  panic(err)
}
ctx := context.Background()
workspaces, err := a.Workspaces.List(ctx)
if err != nil {
  panic(err)
}
w, err := a.GetWorkspaceClient(workspaces[0])
if err != nil {
  panic(err)
}
me, err := w.CurrentUser.Me(ctx)
```
* Added support to select Spark version with Photon ([#799](https://github.com/databricks/databricks-sdk-go/pull/799)). Please Note: Photon selection is disabled by default. To enable it, please use `Photon: true` in request. Example:
````
	version, err := sparkVersions.Select(compute.SparkVersionRequest{
		Photon:          true,
	})
````


API Changes:

Additions:

 * Added the following Requests:
    - [catalog.CancelRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CancelRefreshRequest).
    - [catalog.GetRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRefreshRequest).
    - [catalog.ListRefreshesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRefreshesRequest).
    - [settings.DeleteRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingRequest).
    - [settings.GetDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetDefaultNamespaceSettingRequest).
    - [settings.GetPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPersonalComputeSettingRequest).
    - [settings.GetRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetRestrictWorkspaceAdminsSettingRequest).
    - [settings.UpdateDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultNamespaceSettingRequest).
    - [settings.UpdateRestrictWorkspaceAdminsSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateRestrictWorkspaceAdminsSettingRequest).
    - [catalog.RunRefreshRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RunRefreshRequest).
    - [settings.DeleteDefaultNamespaceSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingRequest).
    - [files.CreateDirectoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#CreateDirectoryRequest).
    - [files.DeleteDirectoryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DeleteDirectoryRequest).
    - [files.ListDirectoryContentsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ListDirectoryContentsRequest).
 * Added the following Responses:
    - [settings.DeleteRestrictWorkspaceAdminsSettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteRestrictWorkspaceAdminsSettingResponse).
    - [settings.DeleteDefaultNamespaceSettingResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultNamespaceSettingResponse).
    - [files.ListDirectoryResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#ListDirectoryResponse).
 * Added `CancelRefresh`, `GetRefresh`, `ListRefreshes` and `RunRefresh` method for [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service.
 * Added `Abfss` and `Gcs` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added `CreateDirectory`, `DeleteDirectory` and `ListDirectoryContents` method for [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added `Source` field for [jobs.DbtTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#DbtTask) and [jobs.SqlTaskFile](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SqlTaskFile).
 * Added [settings.RestrictWorkspaceAdminsMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessage).
 * Jobs:
    - [jobs.ForEachStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachStats).
    - [jobs.ForEachTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTask).
    - [jobs.ForEachTaskErrorMessageStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskErrorMessageStats).
    - [jobs.ForEachTaskTaskRunStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ForEachTaskTaskRunStats).
    - [jobs.RunForEachTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunForEachTask).
 * Pipelines:
    - [pipelines.PipelineClusterAutoscale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscale).
    - [pipelines.PipelineClusterAutoscaleMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscaleMode).
 * Settings:
    - [settings.RestrictWorkspaceAdminsMessageStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsMessageStatus).
    - [settings.RestrictWorkspaceAdminsSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#RestrictWorkspaceAdminsSetting).
 * Catalog:
    - [catalog.MonitorRefreshInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfo).
    - [catalog.MonitorRefreshInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorRefreshInfoState).
 * Added `GetPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
 * Added the following fields:
  - `DeltaSyncIndexSpec` for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest).
  - `FileType` for [workspace.ExportResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ExportResponse).
  - `ResourceId` for [workspace.ObjectInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ObjectInfo).
 * Added `ZoneId` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
 * Added `ForEachTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask), [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask) and [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `DeleteDefaultNamespaceSetting`, `DeleteRestrictWorkspaceAdminsSetting`, `GetDefaultNamespaceSetting`, `GetRestrictWorkspaceAdminsSetting`, `UpdateDefaultNamespaceSetting` and `UpdateRestrictWorkspaceAdminsSetting` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Added `FieldMask` field for [settings.UpdatePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePersonalComputeSettingRequest).
 * Added `UsePreemptibleExecutors` field for [compute.GcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcpAttributes).
 * Misc:
    - [compute.Adlsgen2Info](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Adlsgen2Info).
    - [compute.GcsStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GcsStorageInfo).
    - [files.DirectoryEntry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DirectoryEntry).
    - [files.PageToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#PageToken).

Changes:
 * Changed `MaxWorkers` and `MinWorkers`  field for [compute.AutoScale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#AutoScale) to no longer be required.
 * Changed `Destination` field for [compute.LocalFileInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LocalFileInfo), [compute.S3StorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#S3StorageInfo), [compute.VolumesStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#VolumesStorageInfo), [compute.WorkspaceStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#WorkspaceStorageInfo) to be required.
 * Changed `Destination` field for [compute.DbfsStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#DbfsStorageInfo) to be required.
 * Changed `Clients` field for [compute.WorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#WorkloadType) to be required.
 * Changed `Autoscale` field for [pipelines.PipelineCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineCluster) to [pipelines.PipelineClusterAutoscale](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineClusterAutoscale).
 * Changed `DeletePersonalComputeSetting` and `UpdatePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `AllowMissing` and `Setting` field for [settings.UpdatePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdatePersonalComputeSettingRequest) to be required.
 * Changed `Etag` field for [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest) to no longer be required.

Removals:
 * Removed `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection), [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore), [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest), [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema), [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Settings:
    - [settings.DeleteDefaultWorkspaceNamespaceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceResponse).
    - [settings.DeleteDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceRequest).
    - [settings.ReadDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadDefaultWorkspaceNamespaceRequest).
    - [settings.UpdateDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWorkspaceNamespaceRequest).
    - [settings.ReadPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadPersonalComputeSettingRequest).
 * Removed `ReadDefaultWorkspaceNamespace`, `UpdateDefaultWorkspaceNamespace` and `DeleteDefaultWorkspaceNamespace` method for [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Removed `Reset` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Removed [pipelines.ResetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#ResetRequest).
 * Removed `ReadPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service.
 * Removed `DeltaSyncVectorIndexSpec` field for [vectorsearch.CreateVectorIndexRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch#CreateVectorIndexRequest).

 Internal Changes:
* Added Support for HEAD operations ([#802](https://github.com/databricks/databricks-sdk-go/pull/802)).
* Updated actions/setup-go to v5 ([#784](https://github.com/databricks/databricks-sdk-go/pull/784)).
* Retry update of catalog in test to avoid flakiness ([#788](https://github.com/databricks/databricks-sdk-go/pull/788)).
* Skip AccountClient_GetWorkspaceClient() test in Azure/GCP ([#798](https://github.com/databricks/databricks-sdk-go/pull/798)).
* SDK Generation + Fix backwards incompatible changes ([#806](https://github.com/databricks/databricks-sdk-go/pull/806)).
* Fixed `any` references when generating code with circular dependencies ([#805](https://github.com/databricks/databricks-sdk-go/pull/805)).
* Fixed stack overflow on recursive schemas ([#801](https://github.com/databricks/databricks-sdk-go/pull/801)).

Dependency updates:
 * Bump google.golang.org/api from 0.154.0 to 0.161.0 ([#794](https://github.com/databricks/databricks-sdk-go/pull/794)).

OpenAPI SHA: cadf1693527b365728a55ff06a0e38ce5740c9f7, Date: 2024-02-08


## Release v0.30.1 (2024-01-25)

Major changes:

* Enable Databricks CLI auth for all clouds ([#783](https://github.com/databricks/databricks-sdk-go/pull/783)).

Other changes:

* Update codecov/codecov-action to v3 ([#785](https://github.com/databricks/databricks-sdk-go/pull/785)).
* MustUseJson must be true if request is a map ([#786](https://github.com/databricks/databricks-sdk-go/pull/786)).


## Release v0.30.0 (2024-01-24)

Other changes:

* Added RequiredPositionalArguments method for codegen ([#773](https://github.com/databricks/databricks-sdk-go/pull/773)).
* Support locking in integration tests ([#776](https://github.com/databricks/databricks-sdk-go/pull/776)).
* Update OpenAPI spec and fix incompatible changes ([#778](https://github.com/databricks/databricks-sdk-go/pull/778)).

API Changes:

 * Added `Exists` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
 * Added [w.LakehouseMonitors](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#LakehouseMonitorsAPI) workspace-level service.
 * Added [catalog.CreateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMonitor), [catalog.DeleteLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteLakehouseMonitorRequest), [catalog.ExistsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExistsRequest), [catalog.GetLakehouseMonitorRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetLakehouseMonitorRequest), [catalog.MonitorCronSchedule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedule), [catalog.MonitorCronSchedulePauseStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCronSchedulePauseStatus), [catalog.MonitorCustomMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetric), [catalog.MonitorCustomMetricType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorCustomMetricType), [catalog.MonitorDataClassificationConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDataClassificationConfig), [catalog.MonitorDestinations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorDestinations), [catalog.MonitorInferenceLogProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileType), [catalog.MonitorInferenceLogProfileTypeProblemType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInferenceLogProfileTypeProblemType), [catalog.MonitorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfo), [catalog.MonitorInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorInfoStatus), [catalog.MonitorNotificationsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorNotificationsConfig), [catalog.MonitorTimeSeriesProfileType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MonitorTimeSeriesProfileType), [catalog.TableExistsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableExistsResponse) and [catalog.UpdateMonitor](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMonitor).
 * Added `InitScripts` field for [pipelines.PipelineCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineCluster).
 * Added `ValidateOnly` field for [pipelines.StartUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#StartUpdate).
 * Added `ValidateOnly` field for [pipelines.UpdateInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#UpdateInfo).
 * Changed `CreateOboToken` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service with new required argument order.
 * Changed `Get` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service to return [settings.GetTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenResponse).
 * Changed `LifetimeSeconds` field for [settings.CreateOboTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateOboTokenRequest) to no longer be required.
 * Added [settings.GetTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenResponse).
 * Changed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service . New request type is [sql.DashboardPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardPostContent).
 * Added `Update` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service.
 * Removed [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
 * Added `HttpHeaders` field for [sql.ExternalLink](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExternalLink).
 * Added `RunAsRole` field for [sql.QueryEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryEditContent).
 * Added [sql.DashboardEditContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardEditContent).
 * Added [sql.DashboardPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardPostContent).

OpenAPI SHA: e05401ed5dd4974c5333d737ec308a7d451f749f, Date: 2024-01-23

## Release v0.29.1 (2024-01-16)

This patch release contains two small changes:
* Retry on Status Code 503 ([#733](https://github.com/databricks/databricks-sdk-go/pull/733)), improving the stability of the SDK in light of transient API unavailability.
* Simplify mocking of iterator and waiter objects ([#769](https://github.com/databricks/databricks-sdk-go/pull/769), [#770](https://github.com/databricks/databricks-sdk-go/pull/770)). See the [Testing section of the README.md](https://github.com/databricks/databricks-sdk-go#testing) for usage information and examples.
## Release v0.29.0 (2024-01-10)

* Extract API interfaces for all services and generate mock clients ([#740](https://github.com/databricks/databricks-sdk-go/pull/740)).
* Handle empty response for errors ([#756](https://github.com/databricks/databricks-sdk-go/pull/756)).
* Update SDK to OpenAPI spec + Fix tests ([#755](https://github.com/databricks/databricks-sdk-go/pull/755)).
* Add utility to retry on specific errors ([#757](https://github.com/databricks/databricks-sdk-go/pull/757)).
* Integration test fixes for TestMwsAccWorkspaces ([#763](https://github.com/databricks/databricks-sdk-go/pull/763)) and TestMwsAccUsageDownload ([#764](https://github.com/databricks/databricks-sdk-go/pull/764)).

Note: This release contains breaking changes, please see below.

API Changes:

 * [Breaking] Changed `List` method for [w.ExternalLocations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationsAPI) workspace-level service to require request of [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest), [w.StorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialsAPI) workspace-level service to require request of [catalog.ListStorageCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsRequest) and [w.Tokens](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokensAPI) workspace-level service to return [settings.ListPublicTokensResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPublicTokensResponse).
 * Added `NextPageToken` field for [catalog.ListExternalLocationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsResponse), [catalog.ListFunctionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsResponse), [catalog.ListSchemasResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasResponse) and [catalog.ListStorageCredentialsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsResponse).
 * Added `MaxResults` field for [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest) and [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest).
 * Added `PageToken` field for [catalog.ListFunctionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListFunctionsRequest) and [catalog.ListSchemasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSchemasRequest).
 * Added `OmitColumns` and `OmitProperties` field for [catalog.ListTablesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListTablesRequest).
 * Added [catalog.ListExternalLocationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListExternalLocationsRequest).
 * Added [catalog.ListStorageCredentialsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListStorageCredentialsRequest).
 * Added [settings.ListPublicTokensResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPublicTokensResponse).
 * Added [dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/dashboards) package.
 * Added [vectorsearch](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/vectorsearch) package.

OpenAPI SHA: a7a9dc025bb80303e676bf3708942c6aa06689f1, Date: 2024-01-04
Dependency updates:

 * Bump google.golang.org/api from 0.153.0 to 0.154.0 ([#741](https://github.com/databricks/databricks-sdk-go/pull/741)).
 * Bump golang.org/x/crypto from 0.14.0 to 0.17.0 in /examples/slog ([#747](https://github.com/databricks/databricks-sdk-go/pull/747)) and /examples/zerolog ([#748](https://github.com/databricks/databricks-sdk-go/pull/748)).
 * Bump golang.org/x/crypto from 0.16.0 to 0.17.0 ([#749](https://github.com/databricks/databricks-sdk-go/pull/749)).

## Release v0.28.1 (2023-12-19)

This is a bugfix release that improves debuggability of unexpected errors ([#744](https://github.com/databricks/databricks-sdk-go/pull/744), [#750](https://github.com/databricks/databricks-sdk-go/pull/750)). When the SDK cannot deserialize a response from the Databricks REST API, the resulting error will include debugging information and instructions on how to submit a bug report to the SDK.


## Release v0.28.0 (2023-12-15)

* Consolidate usage of named sorting ([#736](https://github.com/databricks/databricks-sdk-go/pull/736)).
* Remove redundant retries for token refreshes as they're already handled in `httpclient` ([#729](https://github.com/databricks/databricks-sdk-go/pull/729)).
* Skip auth, config loading, and rate limits when fixture transport is present in the config ([#739](https://github.com/databricks/databricks-sdk-go/pull/739)).
* Generate SDK to latest OpenAPI specification with bugfix ([#742](https://github.com/databricks/databricks-sdk-go/pull/742)).

API Changes:

 * Added `AzureWorkspaceInfo` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added [provisioning.AzureWorkspaceInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#AzureWorkspaceInfo).
 * Changed `UpdateConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service with new required argument order.
 * Changed `ServedEntities` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput) to no longer be required.
 * Changed `Create` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Replace` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Update` method for [a.AccountIpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountIpAccessListsAPI) account-level service with new required argument order.
 * Changed `Create` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `Replace` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service with new required argument order.
 * Changed `IpAddresses` field for [settings.CreateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateIpAccessList) to no longer be required.
 * Changed `IpAddresses` field for [settings.ReplaceIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReplaceIpAccessList) to no longer be required.
 * Removed `ListId` field for [settings.ReplaceIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReplaceIpAccessList).
 * Changed `Enabled` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Changed `IpAddresses` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Changed `Label` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.
 * Removed `ListId` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList).
 * Changed `ListType` field for [settings.UpdateIpAccessList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateIpAccessList) to no longer be required.

OpenAPI SHA: d3853c8dee5806d04da2ae8910f273ffb35719a5, Date: 2023-12-14

## Release v0.27.0 (2023-12-12)


API Changes:

 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Added `CloudflareApiToken` field for [catalog.CreateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateStorageCredential).
 * Added `CloudflareApiToken` field for [catalog.StorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfo).
 * Changed `Name` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog) to be required.
 * Added `NewName` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Changed `Name` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection) to no longer be required.
 * Added `NewName` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Changed `Name` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation) to be required.
 * Added `NewName` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `NewName` field for [catalog.UpdateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateMetastore).
 * Added `NewName` field for [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
 * Added `NewName` field for [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
 * Changed `Name` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential) to be required.
 * Added `CloudflareApiToken` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential).
 * Added `NewName` field for [catalog.UpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateStorageCredential).
 * Added `NewName` field for [catalog.UpdateVolumeRequestContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateVolumeRequestContent).
 * Added `CloudflareApiToken` field for [catalog.ValidateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ValidateStorageCredential).
 * Added [catalog.CloudflareApiToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CloudflareApiToken).
 * Removed `Continuous` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Removed `Continuous` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Changed `JobParameters` field for [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask) to [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Added `RunIf` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `RunJobTask` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Changed `UpdateConfig` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service with new required argument order.
 * Added `Put` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `RateLimits` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Changed `ServedModels` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput) to no longer be required.
 * Added `AutoCaptureConfig` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigInput).
 * Added `AutoCaptureConfig` field for [serving.EndpointCoreConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigOutput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigOutput).
 * Added `ServedEntities` field for [serving.EndpointCoreConfigSummary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointCoreConfigSummary).
 * Added `ServedEntities` field for [serving.EndpointPendingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointPendingConfig).
 * Added `ExtraParams` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Input` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `MaxTokens` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Messages` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `N` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Prompt` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Stop` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Stream` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added `Temperature` field for [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Changed `Predictions` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse) to no longer be required.
 * Added `Choices` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Created` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Data` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Id` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Model` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Object` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Added `Usage` field for [serving.QueryEndpointResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponse).
 * Changed `WorkloadSize` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
 * Changed `WorkloadType` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput) to [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
 * Added `Task` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `Task` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.Ai21LabsConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#Ai21LabsConfig).
 * Added [serving.AnthropicConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AnthropicConfig).
 * Added [serving.AutoCaptureConfigInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureConfigInput).
 * Added [serving.AutoCaptureConfigOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureConfigOutput).
 * Added [serving.AutoCaptureState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AutoCaptureState).
 * Added [serving.AwsBedrockConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfig).
 * Added [serving.AwsBedrockConfigBedrockProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AwsBedrockConfigBedrockProvider).
 * Added [serving.ChatMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ChatMessage).
 * Added [serving.ChatMessageRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ChatMessageRole).
 * Added [serving.CohereConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CohereConfig).
 * Added [serving.DatabricksModelServingConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DatabricksModelServingConfig).
 * Added [serving.EmbeddingsV1ResponseEmbeddingElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EmbeddingsV1ResponseEmbeddingElement).
 * Added [serving.EmbeddingsV1ResponseEmbeddingElementObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EmbeddingsV1ResponseEmbeddingElementObject).
 * Added [serving.ExternalModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModel).
 * Added [serving.ExternalModelConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelConfig).
 * Added [serving.ExternalModelProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelProvider).
 * Added [serving.ExternalModelUsageElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ExternalModelUsageElement).
 * Added [serving.FoundationModel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#FoundationModel).
 * Added [serving.OpenAiConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#OpenAiConfig).
 * Added [serving.PaLmConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PaLmConfig).
 * Added [serving.PayloadTable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PayloadTable).
 * Added [serving.PutRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutRequest).
 * Added [serving.PutResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PutResponse).
 * Added [serving.QueryEndpointResponseObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointResponseObject).
 * Added [serving.RateLimit](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimit).
 * Added [serving.RateLimitKey](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimitKey).
 * Added [serving.RateLimitRenewalPeriod](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#RateLimitRenewalPeriod).
 * Added [serving.ServedEntityInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityInput).
 * Added [serving.ServedEntityOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntityOutput).
 * Added [serving.ServedEntitySpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedEntitySpec).
 * Added [serving.ServedModelInputWorkloadSize](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadSize).
 * Added [serving.ServedModelInputWorkloadType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInputWorkloadType).
 * Added [serving.V1ResponseChoiceElement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#V1ResponseChoiceElement).
 * Removed [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Removed [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Removed [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Removed [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Removed [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Removed [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).
 * Removed `Name` field for [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Changed `Name` field for [sharing.UpdateProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateProvider) to be required.
 * Added `NewName` field for [sharing.UpdateProvider](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateProvider).
 * Changed `Name` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient) to be required.
 * Added `NewName` field for [sharing.UpdateRecipient](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateRecipient).
 * Changed `Name` field for [sharing.UpdateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateShare) to be required.
 * Added `NewName` field for [sharing.UpdateShare](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateShare).
 * Added `StatementIds` field for [sql.QueryFilter](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryFilter).
 * Added [sql.StatementId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementId).

OpenAPI SHA: 63caa3cb0c05045e81d3dcf2451fa990d8670f36, Date: 2023-12-12
Dependency updates:

 * Bumped google.golang.org/api from 0.152.0 to 0.153.0 ([#731](https://github.com/databricks/databricks-sdk-go/pull/731)).

## Release v0.26.2 (2023-12-07)

This is a bugfix release, including a fix correcting issues with OAuth flows, due to a bug with the propagation of the response status in `httpclient`'s `RoundTrip()` implementation. This fixes the `failed during request visitor: token: oauth2: cannot fetch token: Response: {...}` error.

All fixes:

* Migrate Azure MSI & Metadata Service token sources to `httpclient` and add 100% test coverage ([#709](https://github.com/databricks/databricks-sdk-go/pull/709)).
* Added `config.NewAzureCliTokenSource` and `config.NewAzureMsiTokenSource` constructors ([#727](https://github.com/databricks/databricks-sdk-go/pull/727)).
* Use per-config refresh context for OAuth tokens ([#728](https://github.com/databricks/databricks-sdk-go/pull/728)).


## Release v0.26.1 (2023-11-30)

Minor changes:
* Support overriding DatabricksEnvironment ([#723](https://github.com/databricks/databricks-sdk-go/pull/723)).
* Detect `Accept` header in `httpclient.WithResponseUnmarshal` ([#710](https://github.com/databricks/databricks-sdk-go/pull/710)).
* Detect `Content-Type` header in `newRequestBody` for `httpclient` ([#711](https://github.com/databricks/databricks-sdk-go/pull/711)).

Bug fixes:
* Retry request on `REQUEST_LIMIT_EXCEEDED` error returned by the SCIM API ([#721](https://github.com/databricks/databricks-sdk-go/pull/721)).
* Match retry logic of pre-refactor SDK ([#722](https://github.com/databricks/databricks-sdk-go/pull/722)).


## Release v0.26.0 (2023-11-29)

Major changes:

* There has been a major overhaul of error handling. Users can now compare errors in API responses to the well-known error responses defined in the `apierr` package and reexported in the `databricks` package. Users can check whether a specific error was returned, for example `errors.Is(err, databricks.ErrResourceAlreadyExists)`, rather than converting the error to `*APIError` to check the status code and error code. This change is backwards-compatible; users do not need to modify existing error-handling code when upgrading the SDK. See [#682](https://github.com/databricks/databricks-sdk-go/pull/682) and [#703](https://github.com/databricks/databricks-sdk-go/pull/703) for the changes and https://github.com/databricks/databricks-sdk-go/blob/main/error_alias.go for the full set of errors.

Bug fixes:

* Handle "no configuration file found at" error during databricks-cli authentication ([#707](https://github.com/databricks/databricks-sdk-go/pull/707)).
* Introduce `DatabricksEnvironment` and fix Azure MSI auth from ACR, where IMDS doesn't give host environment information ([#700](https://github.com/databricks/databricks-sdk-go/pull/700)).
* Fix SCIM Pagination default parameters in the Go SDK ([#717](https://github.com/databricks/databricks-sdk-go/pull/717)).

Other changes:

* Update `slog` example with the correct interface ([#694](https://github.com/databricks/databricks-sdk-go/pull/694)).
* Fixed typo in error message for unknown azure environment ([#701](https://github.com/databricks/databricks-sdk-go/pull/701)).
* Allow injection of HTTP transport to enable HTTP replayer pattern ([#697](https://github.com/databricks/databricks-sdk-go/pull/697)).
* Decouple HTTP retries and error mapping mechanics from `DatabricksClient` into `httpclient.ApiClient` ([#699](https://github.com/databricks/databricks-sdk-go/pull/699), [#702](https://github.com/databricks/databricks-sdk-go/pull/702), [#712](https://github.com/databricks/databricks-sdk-go/pull/712)).
* Port `qa.HTTPFixtures` to faster transport-level stubs ([#708](https://github.com/databricks/databricks-sdk-go/pull/708)).

API Changes:

 * Removed `EnableOptimization` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `PipelineId` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `EnablePredictiveOptimization` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog) and [catalog.UpdateSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSchema).
 * Removed [catalog.UpdatePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimization) and [catalog.UpdatePredictiveOptimizationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimizationResponse).
 * Added `Description` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `ListNetworkConnectivityConfigurations` and `ListPrivateEndpointRules` method for [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
 * Added [settings.ListNccAzurePrivateEndpointRulesResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNccAzurePrivateEndpointRulesResponse), [settings.ListNetworkConnectivityConfigurationsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNetworkConnectivityConfigurationsRequest), [settings.ListNetworkConnectivityConfigurationsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListNetworkConnectivityConfigurationsResponse), and [settings.ListPrivateEndpointRulesRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListPrivateEndpointRulesRequest).
 * Added `StringSharedAs` field for [sharing.SharedDataObject](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#SharedDataObject).

Internal changes:

* Added `contains` method in OpenAPI Generator ([#690](https://github.com/databricks/databricks-sdk-go/pull/690)).
* Skip recipients tests in Azure ([#692](https://github.com/databricks/databricks-sdk-go/pull/692)).
* Allow Files API tests to run in UC environments ([#695](https://github.com/databricks/databricks-sdk-go/pull/695)).
* More cleanup in Unity Catalog integration test ([#719](https://github.com/databricks/databricks-sdk-go/pull/719)).

OpenAPI SHA: 22f09783eb8a84d52026f856be3b2068f9498db3, Date: 2023-11-23
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.13.0 to 0.14.0 ([#689](https://github.com/databricks/databricks-sdk-go/pull/689)).
 * Bump google.golang.org/api from 0.150.0 to 0.151.0 ([#698](https://github.com/databricks/databricks-sdk-go/pull/698)).
 * Bump the OpenAPI Spec ([#706](https://github.com/databricks/databricks-sdk-go/pull/706)).
 * Bump golang.org/x/oauth2 from 0.14.0 to 0.15.0 ([#715](https://github.com/databricks/databricks-sdk-go/pull/715)).
 * Bump golang.org/x/time from 0.4.0 to 0.5.0 ([#714](https://github.com/databricks/databricks-sdk-go/pull/714)).
 * Bump google.golang.org/api from 0.151.0 to 0.152.0 ([#716](https://github.com/databricks/databricks-sdk-go/pull/716)).


## Release v0.25.0 (2023-11-10)

* Make sure path parameters are first in order in RequiredFields ([#669](https://github.com/databricks/databricks-sdk-go/pull/669)).
* Added Field.IsRequestBodyField method for code generation ([#670](https://github.com/databricks/databricks-sdk-go/pull/670)).
* Added regressions question to the issue template ([#676](https://github.com/databricks/databricks-sdk-go/pull/676)).
* Added telemetry for CI/CD platform to useragent ([#665](https://github.com/databricks/databricks-sdk-go/pull/665)).
* Skiped GCP Integration Tests using Statement Execution API ([#678](https://github.com/databricks/databricks-sdk-go/pull/678)).
* Added more detailed error message on default credentials not found error ([#679](https://github.com/databricks/databricks-sdk-go/pull/679)).
* Updated SDK to latest OpenAPI Spec ([#685](https://github.com/databricks/databricks-sdk-go/pull/685)).

API Changes:

 * Changed `Create` method for [w.Functions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionsAPI) and [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service with new required argument order.
 * Changed `InputParams` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos).
 * Changed `Properties` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to `string`.
 * Changed `ReturnParams` field for [catalog.CreateFunction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunction) and [catalog.FunctionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionInfo) to [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos)
 * Changed `StorageRoot` field for [catalog.CreateMetastore](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastore) to no longer be required.
 * Added `SkipValidation` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `Libraries` field for [compute.CreatePolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreatePolicy), [compute.EditPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditPolicy) and [compute.Policy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#Policy).
 * Added [compute.InitScriptEventDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptEventDetails).
 * Added `InitScripts` field for [compute.EventDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EventDetails).
 * Added `File` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added `ZoneId` field for [compute.InstancePoolGcpAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolGcpAttributes).
 * Added `IncludeResolvedValues` field for [jobs.GetRunRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetRunRequest).
 * Added `EditMode` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `NetworkConnectivityConfigId` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Added `ContainerLogs` and `ExtraInfo` field for [serving.DeploymentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeploymentStatus).
 * Added [catalog.CreateFunctionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateFunctionRequest), [catalog.DependencyList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DependencyList) and [catalog.FunctionParameterInfos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#FunctionParameterInfos).
 * Added [compute.InitScriptExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptExecutionDetails), [compute.InitScriptExecutionDetailsStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptExecutionDetailsStatus), [compute.InitScriptInfoAndExecutionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfoAndExecutionDetails), [compute.LocalFileInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#LocalFileInfo).
 * Added [jobs.CreateJobEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobEditMode) and [jobs.JobSettingsEditMode](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsEditMode).
 * Added `DeleteApp`, `GetApp`, `GetAppDeploymentStatus`, `GetApps` and `GetEvents` method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Added [serving.AppEvents](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppEvents), [serving.AppServiceStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppServiceStatus), [serving.DeleteAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DeleteAppResponse), [serving.GetAppDeploymentStatusRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppDeploymentStatusRequest), [serving.GetAppResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetAppResponse), [serving.GetEventsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetEventsRequest), [serving.ListAppEventsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppEventsResponse) and [serving.ListAppsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ListAppsResponse).
 * Added [a.NetworkConnectivity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityAPI) account-level service.
 * Added [settings.CreateNetworkConnectivityConfigRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreateNetworkConnectivityConfigRequest), [settings.CreatePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequest), [settings.CreatePrivateEndpointRuleRequestGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CreatePrivateEndpointRuleRequestGroupId), [settings.DeleteNetworkConnectivityConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteNetworkConnectivityConfigurationRequest), [settings.DeletePrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePrivateEndpointRuleRequest), [settings.GetNetworkConnectivityConfigurationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetNetworkConnectivityConfigurationRequest), [settings.GetPrivateEndpointRuleRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetPrivateEndpointRuleRequest), [settings.NccAzurePrivateEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRule), [settings.NccAzurePrivateEndpointRuleConnectionState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleConnectionState), [settings.NccAzurePrivateEndpointRuleGroupId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzurePrivateEndpointRuleGroupId), [settings.NccAzureServiceEndpointRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccAzureServiceEndpointRule), [settings.NccEgressConfig](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressConfig), [settings.NccEgressDefaultRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressDefaultRules), [settings.NccEgressTargetRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NccEgressTargetRules), [settings.NetworkConnectivityConfiguration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#NetworkConnectivityConfiguration).
 * Removed `Delete`, `Get`,  method for [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service.
 * Removed [jobs.JobSettingsUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsUiState) and [jobs.CreateJobUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobUiState).
 * Removed [a.OAuthEnrollment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthEnrollmentAPI) account-level service.
 * Removed [oauth2.CreateOAuthEnrollment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateOAuthEnrollment) and [oauth2.OAuthEnrollmentStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthEnrollmentStatus).
 * Removed `UiState` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob) and [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).

OpenAPI SHA: e7b127cb07af8dd4d8c61c7cc045c8910cdbb02a, Date: 2023-11-08
Dependency updates:

 * Bump google.golang.org/api from 0.146.0 to 0.150.0 ([#683](https://github.com/databricks/databricks-sdk-go/pull/683)).
 * Bump golang.org/x/mod from 0.13.0 to 0.14.0 ([#681](https://github.com/databricks/databricks-sdk-go/pull/681)).
 * Bump google.golang.org/grpc from 1.58.2 to 1.58.3 in /examples/slog ([#672](https://github.com/databricks/databricks-sdk-go/pull/672)).
 * Bump google.golang.org/grpc to 1.58.3 in /examples/zerolog ([#684](https://github.com/databricks/databricks-sdk-go/pull/684)).
 * Bump golang.org/x/time from 0.3.0 to 0.4.0 ([#680](https://github.com/databricks/databricks-sdk-go/pull/680)).

## Release v0.24.0 (2023-10-24)

* Implemented Iterator support for paginated endpoints or endpoints returning a list ([#543](https://github.com/databricks/databricks-sdk-go/pull/543)). The Iterator interface allows users to get the next resource in the iterator and to check whether another resource is available in the iterator. Iterators for paginated endpoints fetch pages lazily, allowing users to only fetch the pages needed for their use case.
* Removed `photon` and `graviton` selectors in `compute.SparkVersionRequ…` ([#622](https://github.com/databricks/databricks-sdk-go/pull/622)). Going forward, photon is determined by the `RuntimeEngine` field in `compute.CreateCluster`, and graviton is chosen depending on the `aws_instance_type` field in `compute.CreateCluster`.

API Changes:

 * Added `Attributes`, `Count`, `ExcludedAttributes`, `Filter`, `SortBy`, `SortOrder`, and `StartIndex` fields to [iam.GetAccountUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetAccountUserRequest) and [iam.GetUserRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetUserRequest).
 * Added `Schemas` field to [iam.Group](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#Group), [iam.ListGroupsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListGroupsResponse), [iam.ListServicePrincipalResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListServicePrincipalResponse), [iam.ListUsersResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListUsersResponse), [iam.ServicePrincipal](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipal), and [iam.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#User).
 * Added [iam.GetSortOrder](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetSortOrder).
 * Added [iam.GroupSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GroupSchema).
 * Added [iam.ListResponseSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ListResponseSchema).
 * Added [iam.ServicePrincipalSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ServicePrincipalSchema).
 * Added [iam.UserSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UserSchema).
 * Added `WebhookNotifications` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added [w.Apps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#AppsAPI) workspace-level service and related methods.
 * Added [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service and related methods.

Internal SDK Changes:
* Update to actions/checkout@v4 ([#650](https://github.com/databricks/databricks-sdk-go/pull/650)).
* Skip unshallow step in test workflow ([#649](https://github.com/databricks/databricks-sdk-go/pull/649)).
* Add integration tests for `Jobs`: `ListRuns` ([#645](https://github.com/databricks/databricks-sdk-go/pull/645)).
* Only log pkg.Load after checking whether the operation is tagged ([#655](https://github.com/databricks/databricks-sdk-go/pull/655)).
* Prefix library paths with the target directory to construct absolute paths ([#656](https://github.com/databricks/databricks-sdk-go/pull/656)).
* Fixed HasRequiredNonBodyField method ([#660](https://github.com/databricks/databricks-sdk-go/pull/660)).
* Added CanSetRequiredFieldsFromJson method for OpenAPI generator ([#661](https://github.com/databricks/databricks-sdk-go/pull/661)).
* Add integration tests for `ServicePrincipals`: `Patch` ([#662](https://github.com/databricks/databricks-sdk-go/pull/662)).
* Add integration tests for `Users`: `Patch`, `Update` ([#663](https://github.com/databricks/databricks-sdk-go/pull/663)).
* Enforce running `az login --service-principal` on nightly runs ([#659](https://github.com/databricks/databricks-sdk-go/pull/659)).
* Add integration tests for `Connections`: `Create`, `Delete`, `Get`, `List`, `Update` ([#653](https://github.com/databricks/databricks-sdk-go/pull/653)).

OpenAPI SHA: 5903bb39137fd76ac384b2044e425f9c56840e00, Date: 2023-10-23
## Release v0.23.0 (2023-10-12)

Note: this release includes custom JSON marshalling that captures whether or
not a value is set and should be included in the marshalled output, if its
value is equal to Go's zero value (i.e. `0` for an int field).

* Add ForceSendFields and a custom marshaller ([#615](https://github.com/databricks/databricks-sdk-go/pull/615)).
* Support text/plain response as streaming request body ([#638](https://github.com/databricks/databricks-sdk-go/pull/638)).
* Added log statement if OIDC response is non 200 ([#644](https://github.com/databricks/databricks-sdk-go/pull/644)).

API Changes:

 * Changed `Download` method for [a.BillableUsage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#BillableUsageAPI) account-level service to start returning [billing.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DownloadResponse).
 * Added [billing.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/billing#DownloadResponse).
 * Changed `Delete` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Changed `Get` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Changed `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service with new required argument order.
 * Added `GetBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service.
 * Added `UpdateBindings` method for [w.WorkspaceBindings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsAPI) workspace-level service.
 * Removed `Name` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential).
 * Added `StorageCredentialName` field for [catalog.AccountsUpdateStorageCredential](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsUpdateStorageCredential).
 * Removed `Name` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Added `StorageCredentialName` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Removed `Name` field for [catalog.GetAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetAccountStorageCredentialRequest).
 * Added `StorageCredentialName` field for [catalog.GetAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetAccountStorageCredentialRequest).
 * Added `Owner` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added [catalog.GetBindingsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetBindingsRequest).
 * Added [catalog.UpdateWorkspaceBindingsParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateWorkspaceBindingsParameters).
 * Added [catalog.WorkspaceBinding](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBinding).
 * Added [catalog.WorkspaceBindingBindingType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingBindingType).
 * Added [catalog.WorkspaceBindingsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceBindingsResponse).
 * Added `Spec` field for [compute.ClusterDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterDetails).
 * Added `ApplyPolicyDefaultValues` field for [compute.ClusterSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterSpec).
 * Removed `AwsAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `AzureAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `DiskSpec` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `EnableElasticDisk` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `GcpAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `PreloadedDockerImages` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed `PreloadedSparkVersions` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Added `Deployment` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `UiState` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Deployment` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `UiState` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Removed `ConditionTask` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `WebhookNotifications` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added [jobs.CreateJobUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJobUiState).
 * Added [jobs.JobDeployment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeployment).
 * Added [jobs.JobDeploymentKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobDeploymentKind).
 * Added [jobs.JobSettingsUiState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettingsUiState).
 * Added `WorkloadType` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `WorkloadType` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
 * Removed [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Changed `List` method for [w.IpAccessLists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#IpAccessListsAPI) workspace-level service to return [settings.ListIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListIpAccessListResponse).
 * Removed [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Removed [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Removed [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Removed `IpAccessLists` field for [settings.GetIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListResponse).
 * Added `IpAccessList` field for [settings.GetIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetIpAccessListResponse).
 * Removed [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Removed [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).
 * Added [settings.ListIpAccessListResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ListIpAccessListResponse).

OpenAPI SHA: 493a76554afd3afdd15dc858773d01643f80352a, Date: 2023-10-12

Dependency updates:

 * Bump golang.org/x/mod from 0.12.0 to 0.13.0 ([#639](https://github.com/databricks/databricks-sdk-go/pull/639)).
 * Bump google.golang.org/api from 0.140.0 to 0.146.0 ([#643](https://github.com/databricks/databricks-sdk-go/pull/643)).
 * Bump golang.org/x/net from 0.15.0 to 0.17.0 ([#646](https://github.com/databricks/databricks-sdk-go/pull/646)).
 * Bump golang.org/x/oauth2 from 0.12.0 to 0.13.0 ([#642](https://github.com/databricks/databricks-sdk-go/pull/642)).
 * Bump golang.org/x/net from 0.8.0 to 0.17.0 in /examples/zerolog ([#647](https://github.com/databricks/databricks-sdk-go/pull/647)).
 * Bump golang.org/x/net from 0.8.0 to 0.17.0 in /examples/slog ([#648](https://github.com/databricks/databricks-sdk-go/pull/648)).

## Release v0.22.0 (2023-10-03)

Breaking API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

API Changes:

 * Added [catalog.ListAccountMetastoreAssignmentsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListAccountMetastoreAssignmentsResponse).

Internal SDK Changes:

* Add support for template libraries ([#635](https://github.com/databricks/databricks-sdk-go/pull/635)).

OpenAPI SHA: bcbf6e851e3d82fd910940910dd31c10c059746c, Date: 2023-10-02

## Release v0.21.0 (2023-09-27)

Breaking API Changes:

 * Changed `ArtifactMatchers` field for [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Owner` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `ArtifactMatchers` field for [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist) to [catalog.ArtifactMatcherList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcherList).
 * Removed `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Changed `CancelAllRuns` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service with new required argument order.
 * Changed `JobId` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns) to no longer be required.
 * Changed `JobParameters` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow) to [jobs.ParamPairs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#ParamPairs).
 * Changed `Query` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service . New request type is [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Removed [serving.QueryRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryRequest).
 * Changed `ExchangeToken` method for [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service with new required argument order.
 * Removed [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Changed `TokenType` field for [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken) to [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Removed `CredentialPartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Changed `List` method for [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service to require request of [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

API Changes:

 * Added `AllQueuedRuns` field for [jobs.CancelAllRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CancelAllRuns).
 * Added `Queue` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `Queue` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `JobParameters` field for [jobs.RepairRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RepairRun).
 * Added `Queue` field for [jobs.RunNow](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunNow).
 * Added `JobParameters` field for [jobs.RunParameters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunParameters).
 * Added `QueueReason` field for [jobs.RunState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunState).
 * Added `QueueDuration` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `Queue` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added [jobs.QueueSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#QueueSettings).
 * Added `Notifications` field for [pipelines.CreatePipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#CreatePipeline).
 * Added `Notifications` field for [pipelines.EditPipeline](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#EditPipeline).
 * Added `Notifications` field for [pipelines.PipelineSpec](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineSpec).
 * Added [pipelines.Notifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#Notifications).
 * Added [serving.DataframeSplitInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#DataframeSplitInput).
 * Added [serving.QueryEndpointInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#QueryEndpointInput).
 * Added [w.Settings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#SettingsAPI) workspace-level service.
 * Added `PartitionId` field for [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.DefaultNamespaceSetting](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DefaultNamespaceSetting).
 * Added [settings.DeleteDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceRequest).
 * Added [settings.DeleteDefaultWorkspaceNamespaceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteDefaultWorkspaceNamespaceResponse).
 * Added [settings.PartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#PartitionId).
 * Added [settings.ReadDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadDefaultWorkspaceNamespaceRequest).
 * Added [settings.StringMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#StringMessage).
 * Added [settings.UpdateDefaultWorkspaceNamespaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateDefaultWorkspaceNamespaceRequest).
 * Added `NextPageToken` field for [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.ListCleanRoomsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsRequest).

SDK Internal Changes:

* Remove use of cancel by job id in integration test ([#632](https://github.com/databricks/databricks-sdk-go/pull/632)).


OpenAPI SHA: 94ddf8ff02df271daebdc5f001075e1ca2ce080d, Date: 2023-09-27

## Release v0.20.0 (2023-09-21)

* Adjust token expiry window to 40 seconds on Azure ([#617](https://github.com/databricks/databricks-sdk-go/pull/617)).
* Rename variables in Azure code to match across auth types ([#618](https://github.com/databricks/databricks-sdk-go/pull/618)).
* Add details to APIError ([#613](https://github.com/databricks/databricks-sdk-go/pull/613)).
* Fixed code generation of primitive types ([#623](https://github.com/databricks/databricks-sdk-go/pull/623)).
* Updated SDK to changes in OpenAPI specification ([#625](https://github.com/databricks/databricks-sdk-go/pull/625)).

API Changes:

 * Changed `List` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to return [catalog.WorkspaceIdList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceIdList).
 * Added [catalog.WorkspaceId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#WorkspaceId).
 * Added [a.OAuthPublishedApps](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#OAuthPublishedAppsAPI) account-level service.
 * Added [oauth2.GetPublishedAppsOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#GetPublishedAppsOutput).
 * Added [oauth2.ListOAuthPublishedAppsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#ListOAuthPublishedAppsRequest).
 * Added [oauth2.PublishedAppOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#PublishedAppOutput).
 * Added `Patch` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `Tags` field for [serving.CreateServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#CreateServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpoint](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpoint).
 * Added `Tags` field for [serving.ServingEndpointDetailed](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointDetailed).
 * Added [serving.EndpointTag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#EndpointTag).
 * Added [serving.PatchServingEndpointTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#PatchServingEndpointTags).
 * Added [w.CredentialsManager](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialsManagerAPI) workspace-level service.
 * Added [settings.CredentialPartitionId](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#CredentialPartitionId).
 * Added [settings.ExchangeToken](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeToken).
 * Added [settings.ExchangeTokenRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenRequest).
 * Added [settings.ExchangeTokenResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ExchangeTokenResponse).
 * Added [settings.TokenType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenType).
 * Changed `ExecuteStatement` method for [w.StatementExecution](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementExecutionAPI) workspace-level service with new required argument order.
 * Added `EmptyResultState` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions).
 * Removed [sql.ChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ChunkInfo).
 * Changed `OnWaitTimeout` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).
 * Changed `Statement` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `WarehouseId` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest) to be required.
 * Changed `Chunks` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest) to [sql.BaseChunkInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfoList).
 * Added `Truncated` field for [sql.ResultManifest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ResultManifest).
 * Removed [sql.TimeoutAction](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#TimeoutAction).
 * Added [sql.AlertOptionsEmptyResultState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptionsEmptyResultState).
 * Added [sql.BaseChunkInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#BaseChunkInfo).
 * Added [sql.ExecuteStatementRequestOnWaitTimeout](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequestOnWaitTimeout).

OpenAPI SHA: 51e3d27c0553c891bc418bd0cd07c9240e0476dd, Date: 2023-09-13
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.11.0 to 0.12.0 ([#606](https://github.com/databricks/databricks-sdk-go/pull/606)).
 * Bump google.golang.org/api from 0.138.0 to 0.140.0 ([#612](https://github.com/databricks/databricks-sdk-go/pull/612)).

## Release v0.19.2 (2023-09-13)

* Fixed case where retry of a request without body fails ([#614](https://github.com/databricks/databricks-sdk-go/pull/614)).


## Release v0.19.1 (2023-09-07)

* Handled Azure authentication when WorkspaceResourceID is provided ([#597](https://github.com/databricks/databricks-sdk-go/pull/597)).
* Reverted error message changes from client.go ([#604](https://github.com/databricks/databricks-sdk-go/pull/604)).
* Reused tokens after first call to Azure CLI ([#605](https://github.com/databricks/databricks-sdk-go/pull/605)).
* Updated error message assertions ([#607](https://github.com/databricks/databricks-sdk-go/pull/607)).


## Release v0.19.0 (2023-09-04)

* Added `ErrNotWorkspaceClient` ([#596](https://github.com/databricks/databricks-sdk-go/pull/596)).
* Fix loading of databrickscfg with a password containing a hash ([#595](https://github.com/databricks/databricks-sdk-go/pull/595)).
* Use an operation's request type name if specified ([#598](https://github.com/databricks/databricks-sdk-go/pull/598)).
* Update OpenAPI spec ([#600](https://github.com/databricks/databricks-sdk-go/pull/600)).

API Changes:

* Renamed permissions APIs to no longer include the service name, for example:
  * `GetJobPermissionLevels` -> `GetPermissionLevels`
  * `GetJobPermissions` -> `GetPermissions`
  * `SetJobPermissions` -> `SetPermissions`
  * `UpdateJobPermissions` -> `UpdatePermissions`
* Changed `Create` method for [w.Volumes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumesAPI) workspace-level service with new required argument order.
* Added `SupportsElasticDisk` field for [compute.NodeType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#NodeType).
* Changed `Create` method for [w.Dashboards](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardsAPI) workspace-level service with new required argument order.
* Added [w.DashboardWidgets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DashboardWidgetsAPI) workspace-level service.
* Added [w.QueryVisualizations](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryVisualizationsAPI) workspace-level service.
* Changed `Name` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest) to be required.
* Added `DashboardFiltersEnabled` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.CreateDashboardRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateDashboardRequest).
* Added `RunAsRole` field for [sql.Query](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Query).
* Added `RunAsRole` field for [sql.QueryPostContent](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryPostContent).
* Removed `DashboardId` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Changed `Position` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions) to [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).
* Removed `Text` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Description` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added `Title` field for [sql.WidgetOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetOptions).
* Added [sql.CreateQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateQueryVisualizationRequest).
* Added [sql.CreateWidget](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#CreateWidget).
* Added [sql.DeleteDashboardWidgetRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteDashboardWidgetRequest).
* Added [sql.DeleteQueryVisualizationRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#DeleteQueryVisualizationRequest).
* Added [sql.RunAsRole](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#RunAsRole).
* Added [sql.WidgetPosition](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WidgetPosition).

OpenAPI SHA: 09a7fa63d9ae243e5407941f200960ca14d48b07, Date: 2023-09-04

Dependency updates:

* Bump google.golang.org/api from 0.136.0 to 0.138.0 ([#587](https://github.com/databricks/databricks-sdk-go/pull/587)).
## Release v0.18.0 (2023-08-29)

Breaking Changes:
* Added support Files API (application/octet-stream) in OpenAPI Spec ([#572](https://github.com/databricks/databricks-sdk-go/pull/572)). The signatures of `Upload`, `Download` and `Delete` have changed; these methods now take `UploadRequest`, `DownloadRequest`, and `DeleteRequest` structures, respectively. Shortcut methods are generated for `DownloadByFileName` and `DeleteByFileName` for convenience. The `WriteFile` and `ReadFile` methods are removed.
* Propagated Request Headers to client.go ([#589](https://github.com/databricks/databricks-sdk-go/pull/589)). The `Do` method of `Client` now requires an additional parameter for request headers. Add headers to requests using this parameter, or pass `nil` if no headers are needed.

Breaking API Changes:
 * Removed [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service and all associated structures.
 * Removed [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service and all associated structures.
 * Renamed `ProvisioningState` field to `ProvisioningInfo` for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Changed [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState) to [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Removed `InstancePoolFleetAttributes` field for [compute.CreateInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.EditInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#EditInstancePool).
 * Removed [compute.FleetLaunchTemplateOverride](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetLaunchTemplateOverride).
 * Removed [compute.FleetOnDemandOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOption).
 * Removed [compute.FleetOnDemandOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetOnDemandOptionAllocationStrategy).
 * Removed [compute.FleetSpotOption](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOption).
 * Removed [compute.FleetSpotOptionAllocationStrategy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#FleetSpotOptionAllocationStrategy).
 * Removed `InstancePoolFleetAttributes` field for [compute.GetInstancePool](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePool).
 * Removed `InstancePoolFleetAttributes` field for [compute.InstancePoolAndStats](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAndStats).
 * Removed [compute.InstancePoolFleetAttributes](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolFleetAttributes).
 * Changed `GetByName` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Changed `GetExperiment` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service to return [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).
 * Renamed [ml.GetExperimentByNameResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentByNameResponse) to [ml.GetExperimentResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentResponse).

API Changes:
 * Changed `List` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.StorageCredentialInfoList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#StorageCredentialInfoList).
 * Added [w.ModelVersions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionsAPI) workspace-level service.
 * Added [w.RegisteredModels](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelsAPI) workspace-level service.
 * Added `BrowseOnly` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `FullName` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `ProvisioningInfo` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableKind` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `SecurableType` field for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Added `Options` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `Options` field for [catalog.UpdateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateCatalog).
 * Added [catalog.CreateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateRegisteredModelRequest).
 * Added [catalog.DeleteAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAliasRequest).
 * Added [catalog.DeleteModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteModelVersionRequest).
 * Added [catalog.DeleteRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteRegisteredModelRequest).
 * Added [catalog.GetByAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetByAliasRequest).
 * Added [catalog.GetModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetModelVersionRequest).
 * Added [catalog.GetRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetRegisteredModelRequest).
 * Added [catalog.ListModelVersionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsRequest).
 * Added [catalog.ListModelVersionsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListModelVersionsResponse).
 * Added [catalog.ListRegisteredModelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsRequest).
 * Added [catalog.ListRegisteredModelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListRegisteredModelsResponse).
 * Added [catalog.ModelVersionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfo).
 * Added [catalog.ModelVersionInfoStatus](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ModelVersionInfoStatus).
 * Added [catalog.ProvisioningInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfo).
 * Added [catalog.ProvisioningInfoState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningInfoState).
 * Added [catalog.RegisteredModelAlias](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelAlias).
 * Added [catalog.RegisteredModelInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#RegisteredModelInfo).
 * Added [catalog.SetRegisteredModelAliasRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetRegisteredModelAliasRequest).
 * Added [catalog.UpdateModelVersionRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateModelVersionRequest).
 * Added [catalog.UpdateRegisteredModelRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateRegisteredModelRequest).
 * Added `Volumes` field for [compute.InitScriptInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InitScriptInfo).
 * Added [compute.VolumesStorageInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#VolumesStorageInfo).
 * Added [w.Files](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#FilesAPI) workspace-level service.
 * Added [files.DeleteFileRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DeleteFileRequest).
 * Added [files.DownloadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadRequest).
 * Added [files.DownloadResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#DownloadResponse).
 * Added [files.UploadRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/files#UploadRequest).
 * Added `CustomTags` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.UpdateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#UpdateWorkspaceRequest).
 * Added `CustomTags` field for [provisioning.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#Workspace).
 * Added [provisioning.CustomTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CustomTags).
 * Added `Parameters` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added `RowLimit` field for [sql.ExecuteStatementRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#ExecuteStatementRequest).
 * Added [sql.StatementParameterListItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#StatementParameterListItem).

SDK Internal Changes:
* Added support for io.ReadCloser for request and response bodies ([#590](https://github.com/databricks/databricks-sdk-go/pull/590)).
* Fixed nightly test failures for Go SDK ([#592](https://github.com/databricks/databricks-sdk-go/pull/592)).

OpenAPI Generator Changes:
* Improved error handling for ill-defined parameters ([#582](https://github.com/databricks/databricks-sdk-go/pull/582)).
* Switched to logging consistently rather than printing to stdout ([#588](https://github.com/databricks/databricks-sdk-go/pull/588)).

OpenAPI SHA: 5d0ccbb790d341eae8e85321a685a9e9e2d5bf24, Date: 2023-08-29

## Release v0.17.0 (2023-08-18)

* Introduced Artifact Allowlist, Securable Tags, and Subentity Tags services.
* Introduced DeleteRuns and RestoreRuns methods in the Experiments API.
* Introduced the GetSecret method in the Secrets API.
* Renamed Auto Maintenance to Predictive Optimization.
* Set necessary headers when authenticating via Azure CLI ([#584](https://github.com/databricks/databricks-sdk-go/pull/584)).

New Services:

 * Added [w.ArtifactAllowlists](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistsAPI) workspace-level service.
 * Added [w.SecurableTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SecurableTagsAPI) workspace-level service.
 * Added [w.SubentityTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SubentityTagsAPI) workspace-level service.
 * Added [catalog.ArtifactAllowlistInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactAllowlistInfo).
 * Added [catalog.ArtifactMatcher](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactMatcher).
 * Added [catalog.ArtifactType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ArtifactType).
 * Added [catalog.GetArtifactAllowlistRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#GetArtifactAllowlistRequest).
 * Added [catalog.ListSecurableTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableTagsRequest).
 * Added [catalog.ListSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSecurableType).
 * Added [catalog.ListSubentityTagsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ListSubentityTagsRequest).
 * Added [catalog.MatchType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MatchType).
 * Added [catalog.SetArtifactAllowlist](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SetArtifactAllowlist).
 * Added [catalog.TagChanges](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagChanges).
 * Added [catalog.TagKeyValuePair](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagKeyValuePair).
 * Added [catalog.TagSecurable](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurable).
 * Added [catalog.TagSecurableAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignment).
 * Added [catalog.TagSecurableAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSecurableAssignmentsList).
 * Added [catalog.TagSubentity](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentity).
 * Added [catalog.TagSubentityAssignmentsList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagSubentityAssignmentsList).
 * Added [catalog.TagsSubentityAssignment](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TagsSubentityAssignment).
 * Added [catalog.UpdateSecurableType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateSecurableType).
 * Added [catalog.UpdateTags](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTags).

New APIs:

 * Added `DeleteRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `RestoreRuns` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added [ml.DeleteRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRuns).
 * Added [ml.DeleteRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#DeleteRunsResponse).
 * Added [ml.RestoreRuns](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRuns).
 * Added [ml.RestoreRunsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RestoreRunsResponse).
 * Added `GetSecret` method for [w.Secrets](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#SecretsAPI) workspace-level service.
 * Added [workspace.GetSecretRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretRequest).
 * Added [workspace.GetSecretResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetSecretResponse).

Service Renames:

 * Renamed `EffectiveAutoMaintenanceFlag` field to `EffectivePredictiveOptimizationFlag` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed `EnableAutoMaintenance` field to `EnablePredictiveOptimization` for [catalog.CatalogInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CatalogInfo).
 * Renamed [catalog.EffectiveAutoMaintenanceFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlag) to [catalog.EffectivePredictiveOptimizationFlag](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlag).
 * Renamed [catalog.EffectiveAutoMaintenanceFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectiveAutoMaintenanceFlagInheritedFromType) to [catalog.EffectivePredictiveOptimizationFlagInheritedFromType](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EffectivePredictiveOptimizationFlagInheritedFromType).
 * Renamed [catalog.EnableAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnableAutoMaintenance) to [catalog.EnablePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EnablePredictiveOptimization).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo) to `EnablePredictiveOptimization` field for [catalog.SchemaInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SchemaInfo).
 * Renamed `EffectiveAutoMaintenanceFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EffectivePredictiveOptimizationFlag` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Renamed `EnableAutoMaintenance` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo) to `EnablePredictiveOptimization` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).

OpenAPI SHA: beff621d7b3e1d59244e2e34fc53a496f310e130, Date: 2023-08-17

## Release v0.16.0 (2023-08-14)

* Added ability to generate flat names ([#568](https://github.com/databricks/databricks-sdk-go/pull/568)).

API Changes:

 * Changed `Create` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Get` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Update` method for [a.AccountStorageCredentials](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountStorageCredentialsAPI) account-level service to return [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Changed `Create` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Changed `Update` method for [w.Connections](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionsAPI) workspace-level service with new required argument order.
 * Removed `OptionsKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `PropertiesKvpairs` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Options` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `Properties` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `ProvisioningState` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableKind` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Added `SecurableType` field for [catalog.ConnectionInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfo).
 * Removed `OptionsKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Removed `PropertiesKvpairs` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Options` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Added `Properties` field for [catalog.CreateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateConnection).
 * Changed `Algorithm` field for [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails) to no longer be required.
 * Removed `OptionsKvpairs` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added `Options` field for [catalog.UpdateConnection](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateConnection).
 * Added [catalog.AccountsStorageCredentialInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountsStorageCredentialInfo).
 * Added [catalog.ConnectionInfoSecurableKind](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ConnectionInfoSecurableKind).
 * Added [catalog.ProvisioningState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ProvisioningState).
 * Added `DataSecurityMode` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `DockerImage` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Added `SingleUserName` field for [compute.CreateCluster](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#CreateCluster).
 * Removed `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added `Schemas` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).

OpenAPI SHA: 1e3533f94335f0e6c5d9262bc1fea95b3ddcb0e1, Date: 2023-08-11
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.10.0 to 0.11.0 ([#574](https://github.com/databricks/databricks-sdk-go/pull/574)).
 * Bump google.golang.org/api from 0.134.0 to 0.135.0 ([#575](https://github.com/databricks/databricks-sdk-go/pull/575)).
 * Bump google.golang.org/api from 0.135.0 to 0.136.0 ([#576](https://github.com/databricks/databricks-sdk-go/pull/576)).

## Release v0.15.0 (2023-08-07)

To simplify documentation and management of object permissions, this release features a major reorganization of how permissions APIs are structured in the SDK. Rather than using a single `Permissions.Get()` API for all services, each service supporting permissions has its own permissions APIs. Follow these steps to migrate to the current SDK:

 * Change `w.Permissions.Get()` and `w.Permissions.GetByRequestObjectIdAndRequestObjectType()` to `w.<Service>.Get<Service>Permissions()`
 * Change `w.Permissions.GetPermissionLevels()` to `w.<Service>.GetServicePermissionLevels()`
 * Change `w.Permissions.Set()` to `w.<Service>.Set<Service>Permissions()`
 * Change `w.Permissions.Update()` to `w.<Service>.Update<Service>Permissions()`

API Changes:

 * Added `GetClusterPolicyPermissionLevels` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `SetClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `UpdateClusterPolicyPermissions` method for [w.ClusterPolicies](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPoliciesAPI) workspace-level service.
 * Added `GetClusterPermissionLevels` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `SetClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `UpdateClusterPermissions` method for [w.Clusters](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClustersAPI) workspace-level service.
 * Added `GetInstancePoolPermissionLevels` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `GetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `SetInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added `UpdateInstancePoolPermissions` method for [w.InstancePools](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolsAPI) workspace-level service.
 * Added [compute.ClusterAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlRequest).
 * Added [compute.ClusterAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterAccessControlResponse).
 * Added [compute.ClusterPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermission).
 * Added [compute.ClusterPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionLevel).
 * Added [compute.ClusterPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissions).
 * Added [compute.ClusterPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsDescription).
 * Added [compute.ClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPermissionsRequest).
 * Added [compute.ClusterPolicyAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlRequest).
 * Added [compute.ClusterPolicyAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyAccessControlResponse).
 * Added [compute.ClusterPolicyPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermission).
 * Added [compute.ClusterPolicyPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionLevel).
 * Added [compute.ClusterPolicyPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissions).
 * Added [compute.ClusterPolicyPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsDescription).
 * Added [compute.ClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#ClusterPolicyPermissionsRequest).
 * Added [compute.GetClusterPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsRequest).
 * Added [compute.GetClusterPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionLevelsResponse).
 * Added [compute.GetClusterPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPermissionsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsRequest).
 * Added [compute.GetClusterPolicyPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionLevelsResponse).
 * Added [compute.GetClusterPolicyPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetClusterPolicyPermissionsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsRequest).
 * Added [compute.GetInstancePoolPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionLevelsResponse).
 * Added [compute.GetInstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#GetInstancePoolPermissionsRequest).
 * Added [compute.InstancePoolAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlRequest).
 * Added [compute.InstancePoolAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolAccessControlResponse).
 * Added [compute.InstancePoolPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermission).
 * Added [compute.InstancePoolPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionLevel).
 * Added [compute.InstancePoolPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissions).
 * Added [compute.InstancePoolPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsDescription).
 * Added [compute.InstancePoolPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/compute#InstancePoolPermissionsRequest).
 * Changed `Set` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Changed `Update` method for [w.Permissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PermissionsAPI) workspace-level service to start returning [iam.ObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#ObjectPermissions).
 * Added `GetPasswordPermissionLevels` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `GetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `SetPasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `UpdatePasswordPermissions` method for [w.Users](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#UsersAPI) workspace-level service.
 * Added `DisplayName` field for [iam.AccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#AccessControlResponse).
 * Changed `Roles` field for [iam.GetAssignableRolesForResourceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetAssignableRolesForResourceResponse) to [iam.RoleList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#RoleList).
 * Added [iam.GetPasswordPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#GetPasswordPermissionLevelsResponse).
 * Added [iam.PasswordAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlRequest).
 * Added [iam.PasswordAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordAccessControlResponse).
 * Added [iam.PasswordPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermission).
 * Added [iam.PasswordPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionLevel).
 * Added [iam.PasswordPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissions).
 * Added [iam.PasswordPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsDescription).
 * Added [iam.PasswordPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PasswordPermissionsRequest).
 * Added [iam.Role](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#Role).
 * Added `GetJobPermissionLevels` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `GetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `SetJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added `UpdateJobPermissions` method for [w.Jobs](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsAPI) workspace-level service.
 * Added [jobs.GetJobPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsRequest).
 * Added [jobs.GetJobPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionLevelsResponse).
 * Added [jobs.GetJobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GetJobPermissionsRequest).
 * Added [jobs.JobAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlRequest).
 * Added [jobs.JobAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobAccessControlResponse).
 * Added [jobs.JobPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermission).
 * Added [jobs.JobPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionLevel).
 * Added [jobs.JobPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissions).
 * Added [jobs.JobPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsDescription).
 * Added [jobs.JobPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobPermissionsRequest).
 * Added `GetExperimentPermissionLevels` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `SetExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `UpdateExperimentPermissions` method for [w.Experiments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentsAPI) workspace-level service.
 * Added `GetRegisteredModelPermissionLevels` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `GetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `SetRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added `UpdateRegisteredModelPermissions` method for [w.ModelRegistry](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ModelRegistryAPI) workspace-level service.
 * Added [ml.ExperimentAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlRequest).
 * Added [ml.ExperimentAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentAccessControlResponse).
 * Added [ml.ExperimentPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermission).
 * Added [ml.ExperimentPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionLevel).
 * Added [ml.ExperimentPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissions).
 * Added [ml.ExperimentPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsDescription).
 * Added [ml.ExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#ExperimentPermissionsRequest).
 * Added [ml.GetExperimentPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsRequest).
 * Added [ml.GetExperimentPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionLevelsResponse).
 * Added [ml.GetExperimentPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetExperimentPermissionsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsRequest).
 * Added [ml.GetRegisteredModelPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionLevelsResponse).
 * Added [ml.GetRegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#GetRegisteredModelPermissionsRequest).
 * Added [ml.RegisteredModelAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlRequest).
 * Added [ml.RegisteredModelAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelAccessControlResponse).
 * Added [ml.RegisteredModelPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermission).
 * Added [ml.RegisteredModelPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionLevel).
 * Added [ml.RegisteredModelPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissions).
 * Added [ml.RegisteredModelPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsDescription).
 * Added [ml.RegisteredModelPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/ml#RegisteredModelPermissionsRequest).
 * Added `Scopes` field for [oauth2.CreateCustomAppIntegration](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/oauth2#CreateCustomAppIntegration).
 * Added `GetPipelinePermissionLevels` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `GetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `SetPipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added `UpdatePipelinePermissions` method for [w.Pipelines](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinesAPI) workspace-level service.
 * Added [pipelines.GetPipelinePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsRequest).
 * Added [pipelines.GetPipelinePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionLevelsResponse).
 * Added [pipelines.GetPipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#GetPipelinePermissionsRequest).
 * Added [pipelines.PipelineAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlRequest).
 * Added [pipelines.PipelineAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineAccessControlResponse).
 * Added [pipelines.PipelinePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermission).
 * Added [pipelines.PipelinePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionLevel).
 * Added [pipelines.PipelinePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissions).
 * Added [pipelines.PipelinePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsDescription).
 * Added [pipelines.PipelinePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelinePermissionsRequest).
 * Added `GcpManagedNetworkConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GkeConfig` field for [provisioning.CreateWorkspaceRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/provisioning#CreateWorkspaceRequest).
 * Added `GetServingEndpointPermissionLevels` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `GetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `SetServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `UpdateServingEndpointPermissions` method for [w.ServingEndpoints](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointsAPI) workspace-level service.
 * Added `InstanceProfileArn` field for [serving.ServedModelInput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelInput).
 * Added `InstanceProfileArn` field for [serving.ServedModelOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServedModelOutput).
 * Added [serving.GetServingEndpointPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsRequest).
 * Added [serving.GetServingEndpointPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionLevelsResponse).
 * Added [serving.GetServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#GetServingEndpointPermissionsRequest).
 * Added [serving.ServingEndpointAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlRequest).
 * Added [serving.ServingEndpointAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointAccessControlResponse).
 * Added [serving.ServingEndpointPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermission).
 * Added [serving.ServingEndpointPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionLevel).
 * Added [serving.ServingEndpointPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissions).
 * Added [serving.ServingEndpointPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsDescription).
 * Added [serving.ServingEndpointPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/serving#ServingEndpointPermissionsRequest).
 * Added `GetTokenPermissionLevels` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `GetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `SetTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added `UpdateTokenPermissions` method for [w.TokenManagement](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenManagementAPI) workspace-level service.
 * Added [settings.GetTokenPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#GetTokenPermissionLevelsResponse).
 * Added [settings.TokenAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlRequest).
 * Added [settings.TokenAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenAccessControlResponse).
 * Added [settings.TokenPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermission).
 * Added [settings.TokenPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionLevel).
 * Added [settings.TokenPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissions).
 * Added [settings.TokenPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsDescription).
 * Added [settings.TokenPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#TokenPermissionsRequest).
 * Added `GetWarehousePermissionLevels` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `GetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `SetWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `UpdateWarehousePermissions` method for [w.Warehouses](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousesAPI) workspace-level service.
 * Added `CanSubscribeToLiveQuery` field for [sql.QueryInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryInfo).
 * Removed `QueuedOverloadTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `QueuedProvisioningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Removed `TotalPartitionsCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `MetadataTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `OverloadingQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningPhases` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PlanningTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `ProvisioningQueueStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedBytes` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `PrunedFilesCount` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryCompilationStartTimestamp` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added `QueryExecutionTimeMs` field for [sql.QueryMetrics](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#QueryMetrics).
 * Added [sql.GetWarehousePermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsRequest).
 * Added [sql.GetWarehousePermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionLevelsResponse).
 * Added [sql.GetWarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#GetWarehousePermissionsRequest).
 * Added [sql.WarehouseAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlRequest).
 * Added [sql.WarehouseAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehouseAccessControlResponse).
 * Added [sql.WarehousePermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermission).
 * Added [sql.WarehousePermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionLevel).
 * Added [sql.WarehousePermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissions).
 * Added [sql.WarehousePermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsDescription).
 * Added [sql.WarehousePermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#WarehousePermissionsRequest).
 * Added `GetRepoPermissionLevels` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `SetRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `UpdateRepoPermissions` method for [w.Repos](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#ReposAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissionLevels` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `GetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `SetWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added `UpdateWorkspaceObjectPermissions` method for [w.Workspace](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceAPI) workspace-level service.
 * Added [workspace.GetRepoPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsRequest).
 * Added [workspace.GetRepoPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionLevelsResponse).
 * Added [workspace.GetRepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetRepoPermissionsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsRequest).
 * Added [workspace.GetWorkspaceObjectPermissionLevelsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionLevelsResponse).
 * Added [workspace.GetWorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#GetWorkspaceObjectPermissionsRequest).
 * Added [workspace.RepoAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlRequest).
 * Added [workspace.RepoAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoAccessControlResponse).
 * Added [workspace.RepoPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermission).
 * Added [workspace.RepoPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionLevel).
 * Added [workspace.RepoPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissions).
 * Added [workspace.RepoPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsDescription).
 * Added [workspace.RepoPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#RepoPermissionsRequest).
 * Added [workspace.WorkspaceObjectAccessControlRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlRequest).
 * Added [workspace.WorkspaceObjectAccessControlResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectAccessControlResponse).
 * Added [workspace.WorkspaceObjectPermission](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermission).
 * Added [workspace.WorkspaceObjectPermissionLevel](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionLevel).
 * Added [workspace.WorkspaceObjectPermissions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissions).
 * Added [workspace.WorkspaceObjectPermissionsDescription](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsDescription).
 * Added [workspace.WorkspaceObjectPermissionsRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/workspace#WorkspaceObjectPermissionsRequest).

OpenAPI SHA: 386b65ecdc825b9c3ed4aa7ca88e2e5baf9d87df, Date: 2023-08-07
Dependency updates:

 * Bump google.golang.org/api from 0.131.0 to 0.134.0 ([#564](https://github.com/databricks/databricks-sdk-go/pull/564)).

## Release v0.14.1 (2023-07-27)

* Handle nested query parameters in Client.attempt ([#559](https://github.com/databricks/databricks-sdk-go/pull/559)).
* Support x-databricks-path-style overrides at the operation level ([#562](https://github.com/databricks/databricks-sdk-go/pull/562)).


## Release v0.14.0 (2023-07-26)

* Fixed names of keyword arguments in examples ([#560](https://github.com/databricks/databricks-sdk-go/pull/560)).

API Changes:

 * Changed `Create` method for [a.AccountMetastoreAssignments](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#AccountMetastoreAssignmentsAPI) account-level service to no longer return [catalog.CreateMetastoreAssignmentsResponseItemList](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItemList).
 * Added `ConnectionName` field for [catalog.CreateCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateCatalog).
 * Added `AccessPoint` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.CreateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateExternalLocation).
 * Removed [catalog.CreateMetastoreAssignmentsResponseItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#CreateMetastoreAssignmentsResponseItem).
 * Added `AccessPoint` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `EncryptionDetails` field for [catalog.ExternalLocationInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#ExternalLocationInfo).
 * Added `AccessPoint` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `EncryptionDetails` field for [catalog.TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo).
 * Added `AccessPoint` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `EncryptionDetails` field for [catalog.UpdateExternalLocation](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateExternalLocation).
 * Added `AccessPoint` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added `EncryptionDetails` field for [catalog.VolumeInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#VolumeInfo).
 * Added [catalog.EncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#EncryptionDetails).
 * Added [catalog.SseEncryptionDetails](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetails).
 * Added [catalog.SseEncryptionDetailsAlgorithm](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#SseEncryptionDetailsAlgorithm).
 * Added [a.AccountNetworkPolicy](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyAPI) account-level service.
 * Added [settings.AccountNetworkPolicyMessage](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountNetworkPolicyMessage).
 * Added [settings.DeleteAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyRequest).
 * Added [settings.DeleteAccountNetworkPolicyResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeleteAccountNetworkPolicyResponse).
 * Added [settings.ReadAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadAccountNetworkPolicyRequest).
 * Added [settings.UpdateAccountNetworkPolicyRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#UpdateAccountNetworkPolicyRequest).

OpenAPI SHA: fbdd0fa3e83fed2c798a58d376529bdb1285b915, Date: 2023-07-26

## Release v0.13.0 (2023-07-18)

* Add issue templates ([#539](https://github.com/databricks/databricks-sdk-go/pull/539)).
* Added HasRequiredNonBodyField method ([#536](https://github.com/databricks/databricks-sdk-go/pull/536)).
* Make Azure MSI auth account compatible ([#544](https://github.com/databricks/databricks-sdk-go/pull/544)).
* Refactor Handling of Name<->ID Mapping in OpenAPI Generator ([#547](https://github.com/databricks/databricks-sdk-go/pull/547)).
* Regenerate Go SDK from current OpenAPI Specification ([#549](https://github.com/databricks/databricks-sdk-go/pull/549)).
* Parse Camel Case and Pascal Case Enum Values ([#550](https://github.com/databricks/databricks-sdk-go/pull/550)).
* Prepare for auto-releaser infra ([#554](https://github.com/databricks/databricks-sdk-go/pull/554)).
* Added SCIM Patch Acceptance Tests ([#540](https://github.com/databricks/databricks-sdk-go/pull/540)).

API Changes:

 * Removed `Maintenance` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `EnableOptimization` method for [w.Metastores](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#MetastoresAPI) workspace-level service.
 * Added `Update` method for [w.Tables](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TablesAPI) workspace-level service.
 * Added `Force` field for [catalog.DeleteAccountMetastoreRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountMetastoreRequest).
 * Added `Force` field for [catalog.DeleteAccountStorageCredentialRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#DeleteAccountStorageCredentialRequest).
 * Removed [catalog.UpdateAutoMaintenance](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenance).
 * Removed [catalog.UpdateAutoMaintenanceResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateAutoMaintenanceResponse).
 * Added [catalog.UpdatePredictiveOptimization](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimization).
 * Added [catalog.UpdatePredictiveOptimizationResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdatePredictiveOptimizationResponse).
 * Added [catalog.UpdateTableRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#UpdateTableRequest).
 * Added `Schema` field for [iam.PartialUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PartialUpdate).
 * Added [iam.PatchSchema](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/iam#PatchSchema).
 * Added `TriggerInfo` field for [jobs.BaseRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#BaseRun).
 * Added `Health` field for [jobs.CreateJob](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#CreateJob).
 * Added `JobSource` field for [jobs.GitSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#GitSource).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.JobEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobEmailNotifications).
 * Added `Health` field for [jobs.JobSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSettings).
 * Added `TriggerInfo` field for [jobs.Run](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Run).
 * Added `RunJobOutput` field for [jobs.RunOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunOutput).
 * Added `RunJobTask` field for [jobs.RunTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunTask).
 * Added `EmailNotifications` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `Health` field for [jobs.SubmitRun](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitRun).
 * Added `EmailNotifications` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `NotificationSettings` field for [jobs.SubmitTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#SubmitTask).
 * Added `Health` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `RunJobTask` field for [jobs.Task](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#Task).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.TaskEmailNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TaskEmailNotifications).
 * Added `OnDurationWarningThresholdExceeded` field for [jobs.WebhookNotifications](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotifications).
 * Added [jobs.JobSource](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSource).
 * Added [jobs.JobSourceDirtyState](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobSourceDirtyState).
 * Added [jobs.JobsHealthMetric](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthMetric).
 * Added [jobs.JobsHealthOperator](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthOperator).
 * Added [jobs.JobsHealthRule](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRule).
 * Added [jobs.JobsHealthRules](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#JobsHealthRules).
 * Added [jobs.RunJobOutput](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobOutput).
 * Added [jobs.RunJobTask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#RunJobTask).
 * Added [jobs.TriggerInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#TriggerInfo).
 * Added [jobs.WebhookNotificationsOnDurationWarningThresholdExceededItem](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/jobs#WebhookNotificationsOnDurationWarningThresholdExceededItem).
 * Removed `Whl` field for [pipelines.PipelineLibrary](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/pipelines#PipelineLibrary).
 * Changed `DeletePersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `ReadPersonalComputeSetting` method for [a.AccountSettings](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#AccountSettingsAPI) account-level service with new required argument order.
 * Changed `Etag` field for [settings.DeletePersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#DeletePersonalComputeSettingRequest) to be required.
 * Changed `Etag` field for [settings.ReadPersonalComputeSettingRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/settings#ReadPersonalComputeSettingRequest) to be required.
 * Added [w.CleanRooms](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomsAPI) workspace-level service.
 * Added [sharing.CentralCleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CentralCleanRoomInfo).
 * Added [sharing.CleanRoomAssetInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomAssetInfo).
 * Added [sharing.CleanRoomCatalog](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalog).
 * Added [sharing.CleanRoomCatalogUpdate](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCatalogUpdate).
 * Added [sharing.CleanRoomCollaboratorInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomCollaboratorInfo).
 * Added [sharing.CleanRoomInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomInfo).
 * Added [sharing.CleanRoomNotebookInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomNotebookInfo).
 * Added [sharing.CleanRoomTableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CleanRoomTableInfo).
 * Added [sharing.ColumnInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnInfo).
 * Added [sharing.ColumnMask](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnMask).
 * Added [sharing.ColumnTypeName](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ColumnTypeName).
 * Added [sharing.CreateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#CreateCleanRoom).
 * Added [sharing.DeleteCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#DeleteCleanRoomRequest).
 * Added [sharing.GetCleanRoomRequest](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#GetCleanRoomRequest).
 * Added [sharing.ListCleanRoomsResponse](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#ListCleanRoomsResponse).
 * Added [sharing.UpdateCleanRoom](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sharing#UpdateCleanRoom).
 * Changed `Query` field for [sql.Alert](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#Alert) to [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).
 * Changed `Value` field for [sql.AlertOptions](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertOptions) to `any`.
 * Removed `IsDbAdmin` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Removed `ProfileImageUrl` field for [sql.User](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#User).
 * Added [sql.AlertQuery](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/sql#AlertQuery).

OpenAPI SHA: 0a1949ba96f71680dad30e06973eaae85b1307bb, Date: 2023-07-18
Dependency updates:

 * Bump golang.org/x/oauth2 from 0.9.0 to 0.10.0 ([#545](https://github.com/databricks/databricks-sdk-go/pull/545)).
 * Bump golang.org/x/mod from 0.11.0 to 0.12.0 ([#541](https://github.com/databricks/databricks-sdk-go/pull/541)).
 * Bump google.golang.org/api from 0.129.0 to 0.130.0 ([#542](https://github.com/databricks/databricks-sdk-go/pull/542)).
 * Bump google.golang.org/api from 0.130.0 to 0.131.0 ([#551](https://github.com/databricks/databricks-sdk-go/pull/551)).

## Release v0.12.0 (2023-06-28)

 * Change release type
 * Regenerate SDK according to OpenAPI spec ([#534](https://github.com/databricks/databricks-sdk-go/pull/534)).
 * Bump google.golang.org/api from 0.128.0 to 0.129.0 ([#533](https://github.com/databricks/databricks-sdk-go/pull/533)).

## Release v0.11.0 (2023-06-28)

 * Regenerated from OpenAPI spec ([#527](https://github.com/databricks/databricks-sdk-go/pull/527)). Includes bugfixes for System Tables.
 * Support accounts-dod as account endpoint ([#530](https://github.com/databricks/databricks-sdk-go/pull/530)).

Breaking API Changes:

 - Renamed compute.ClusterInfo to compute.ClusterDetails
 - Renamed compue.BaseClusterInfo to compute.ClusterSpec
 - Renamed jobs.RunSubmitTaskSettings to jobs.SubmitTask
 - Renamed jobs.JobTaskSettings to jobs.Task
 - Renamed jobs.CreateJobFormat to jobs.Format
 - Renamed jobs.JobsWebhookNotifications to jobs.WebhookNotifications
 - Renamed jobs.CronSchedulePauseStatus to jobs.PauseStatus
 - Renamed jobs.FileArrivalTriggerSettings to jobs.FileArrivalTriggerConfiguration
 - Renamed jobs.GitSourceGitProvider to jobs.GitProvider
 - Consolidated jobs.NotebookTaskSource, jobs.SparkPythonTaskSource to jobs.Source
 - Consolidated jobs.JobWebhookNotificationsOnFailureItem, jobs.JobWebhookNotificationsOnStartItem, jobs.JobWebhookNotificationsOnSuccessItem to jobs.Webhook
 - Renamed jobs.TaskDependenciesItem to job.TaskDependency
 - jobs.SparkSubmitTask.Widgets changed from *SqlDashboardWidgetOutput to []SqlDashboardWidgetOutput
 - (a *SystemSchemasAPI) DisableByMetastoreIdAndSchemaName changed from func(context.Context, string, string) error to func(context.Context, string, DisableSchemaName) error
 - (a *SystemSchemasAPI) Enable changed from func(context.Context) error to func(context.Context, EnableRequest) error
 - catalog.CatalogInfo.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.CreateConnection.OptionsKvpairs changed from *OptionsKvPairs to map[string]string
 - catalog.UpdateConnection.OptionsKvpairs changed from OptionsKvPairs to map[string]string
 - catalog.DisableRequest.SchemaName changed from string to DisableSchemaName
 - catalog.OptionsKvPairs removed
 - catalog.SystemSchemaInfoStateDisableinitialized renamed to catalog.SystemSchemaInfoStateDisabledInitialized
 - catalog.SystemSchemaInfoStateEnabledcompleted renamed to catalog.SystemSchemaInfoStateEnabledCompleted
 - catalog.SystemSchemaInfoStateEnabledinitialized renamed to catalog.SystemSchemaInfoStateEnabledInitialized
 - settings.ReadPersonalComputeSettingsResponse renamed to settings.PersonalComputeSetting
 - workpace.CreateScope.KeyvaultMetadata removed

Other API Changes:

 - catalog.CatalogInfo.ConnectionName added
 - catalog.CatalogInfo.Options added
 - catalog.DisableSchemaName added
 - catalog.EnableRequest added
 - catalog.PrivilegeUseMarketplaceAssets added
 - compute.ClusterEvent added
 - compute.ComputeSpec added
 - compute.ComputeSpecKind added
 - jobs.ConditionTask added
 - jobs.ConditionTaskOp added
 - jobs.CreateJob.Compute added
 - jobs.JobCompute added
 - jobs.JobSettings.Compute added
 - jobs.RunConditionTask added
 - jobs.RunOutput.ConditionTask added
 - jobs.RunTask.ConditionTask added
 - serving.ServedModelInput.EnvironmentVars added
 - serving.ServedModelOutput.EnvironmentVars added
 - settings.DeletePersonalComputeSetting added
 - settings.UpdatePersonalComputeSetting added
 - sharing.PrivilegeUseMarketplaceAssets added
 - workspace.CreateScope.BackendAzureKeyVault added

## Release v0.10.1 (2023-06-21)

 * Added ClusterID field for Config ([#524](https://github.com/databricks/databricks-sdk-go/pull/524)).
 * Added OnlyPayload property for method ([#522](https://github.com/databricks/databricks-sdk-go/pull/522)).
 * Renamed IsJsonOnly method ([#523](https://github.com/databricks/databricks-sdk-go/pull/523)).

Dependency updates:

* Bump google.golang.org/api from 0.127.0 to 0.128.0 ([#521](https://github.com/databricks/databricks-sdk-go/pull/521)).

## Release v0.10.0 (2023-06-16)

 * Added log level support to SimpleLogger. Default logger now prints only `INFO` level messages. To replicate more verbose behavior from the previous versions, set the `DEBUG` level in `SimpleLogger` explicitly ([#426](https://github.com/databricks/databricks-sdk-go/pull/426)).
 * Added `Upload` and `Download` methods where applicable ([#423](https://github.com/databricks/databricks-sdk-go/pull/423)).
 * Added auth tests for a databrickscfg file with an empty DEFAULT profile ([#496](https://github.com/databricks/databricks-sdk-go/pull/496)).
 * Added more integration tests and examples ([#503](https://github.com/databricks/databricks-sdk-go/pull/503)).
 * Added preview level functions for services and methods to OpenAPI code ([#497](https://github.com/databricks/databricks-sdk-go/pull/497)).
 * Added IsAllRequiredFieldsPrimitive helper function ([#404](https://github.com/databricks/databricks-sdk-go/pull/404)).
 * Added isCrudCreate method ([#416](https://github.com/databricks/databricks-sdk-go/pull/416)).
 * Added retries for getting refreshable token ([#500](https://github.com/databricks/databricks-sdk-go/pull/500)).
 * Added type update when entity fields are updated ([#411](https://github.com/databricks/databricks-sdk-go/pull/411)).
 * Don't panic but return error from NewAccountClient ([#422](https://github.com/databricks/databricks-sdk-go/pull/422)).
 * Error when opening a DBFS directory for reading ([#415](https://github.com/databricks/databricks-sdk-go/pull/415)).
 * Fixed cleanup logic in SQL warehouse integration test ([#400](https://github.com/databricks/databricks-sdk-go/pull/400)).
 * Fixed error handling in cluster utility function ([#399](https://github.com/databricks/databricks-sdk-go/pull/399)).
 * Fixed example in README.md ([#494](https://github.com/databricks/databricks-sdk-go/pull/494)).
 * Fixed nondeterminism in workspace filesystem integration test ([#401](https://github.com/databricks/databricks-sdk-go/pull/401)).
 * Improve command execution interface ([#410](https://github.com/databricks/databricks-sdk-go/pull/410)).
 * Introduce waiters as top-level methods ([#408](https://github.com/databricks/databricks-sdk-go/pull/408)).
 * Regenerate examples if configured ([#518](https://github.com/databricks/databricks-sdk-go/pull/518)).
 * Respect limit field passed in ListRequests ([#407](https://github.com/databricks/databricks-sdk-go/pull/407)).
 * Updated from OpenAPI spec ([#412](https://github.com/databricks/databricks-sdk-go/pull/412), [#413](https://github.com/databricks/databricks-sdk-go/pull/413), [#421](https://github.com/databricks/databricks-sdk-go/pull/421), [#519](https://github.com/databricks/databricks-sdk-go/pull/519), [#424](https://github.com/databricks/databricks-sdk-go/pull/424)).
 * Updated API test template with page tokens ([#417](https://github.com/databricks/databricks-sdk-go/pull/417)).
 * Use constants instead of hardcoding strings ([#402](https://github.com/databricks/databricks-sdk-go/pull/402)).
 * Use `x-databricks-is-accounts` flag to determine whether a service is an account level service ([#420](https://github.com/databricks/databricks-sdk-go/pull/420)).

API Changes:

github.com/databricks/databricks-sdk-go/service/serving
  - (*ServingEndpointsAPI).Create: changed from func(context.Context, CreateServingEndpoint) (*ServingEndpointDetailed, error) to func(context.Context, CreateServingEndpoint) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  - (*ServingEndpointsAPI).List: removed
  - (*ServingEndpointsAPI).UpdateConfig: changed from func(context.Context, EndpointCoreConfigInput) (*ServingEndpointDetailed, error) to func(context.Context, EndpointCoreConfigInput) (*WaitGetServingEndpointNotUpdating[ServingEndpointDetailed], error)
  Compatible changes:
  - (*ServingEndpointsAPI).ListAll: added
  - (*ServingEndpointsAPI).WaitGetServingEndpointNotUpdating: added
  - WaitGetServingEndpointNotUpdating: added

github.com/databricks/databricks-sdk-go
  - WorkspaceClient.CommandExecutor: removed
  - WorkspaceClient.CommandExecution: added
  - WorkspaceClient.Connections: added
  - WorkspaceClient.Files: added
  - WorkspaceClient.SystemSchemas: added

github.com/databricks/databricks-sdk-go/service/sql
  - (*WarehousesAPI).Create: changed from func(context.Context, CreateWarehouseRequest) (*CreateWarehouseResponse, error) to func(context.Context, CreateWarehouseRequest) (*WaitGetWarehouseRunning[CreateWarehouseResponse], error)
  - (*WarehousesAPI).DeleteAndWait: removed
  - (*WarehousesAPI).DeleteByIdAndWait: removed
  - (*WarehousesAPI).Edit: changed from func(context.Context, EditWarehouseRequest) error to func(context.Context, EditWarehouseRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).GetAndWait: removed
  - (*WarehousesAPI).GetByIdAndWait: removed
  - (*WarehousesAPI).Start: changed from func(context.Context, StartRequest) error to func(context.Context, StartRequest) (*WaitGetWarehouseRunning[any], error)
  - (*WarehousesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetWarehouseStopped[any], error)
  - (*StatementExecutionAPI).ExecuteAndWait: added
  - (*WarehousesAPI).WaitGetWarehouseRunning: added
  - (*WarehousesAPI).WaitGetWarehouseStopped: added
  - FormatCsv: added
  - WaitGetWarehouseRunning: added
  - WaitGetWarehouseStopped: added

github.com/databricks/databricks-sdk-go/service/jobs
  - (*JobsAPI).CancelRun: changed from func(context.Context, CancelRun) error to func(context.Context, CancelRun) (*WaitGetRunJobTerminatedOrSkipped[any], error)
  - (*JobsAPI).GetRunAndWait: removed
  - (*JobsAPI).RepairRun: changed from func(context.Context, RepairRun) (*RepairRunResponse, error) to func(context.Context, RepairRun) (*WaitGetRunJobTerminatedOrSkipped[RepairRunResponse], error)
  - (*JobsAPI).RunNow: changed from func(context.Context, RunNow) (*RunNowResponse, error) to func(context.Context, RunNow) (*WaitGetRunJobTerminatedOrSkipped[RunNowResponse], error)
  - (*JobsAPI).Submit: changed from func(context.Context, SubmitRun) (*SubmitRunResponse, error) to func(context.Context, SubmitRun) (*WaitGetRunJobTerminatedOrSkipped[SubmitRunResponse], error)
  - (*JobsAPI).WaitGetRunJobTerminatedOrSkipped: added
  - CreateJob.RunAs: added
  - JobRunAs: added
  - JobSettings.RunAs: added
  - WaitGetRunJobTerminatedOrSkipped: added

github.com/databricks/databricks-sdk-go/service/compute
  - (*ClustersAPI).Create: changed from func(context.Context, CreateCluster) (*CreateClusterResponse, error) to func(context.Context, CreateCluster) (*WaitGetClusterRunning[CreateClusterResponse], error)
  - (*ClustersAPI).Delete: changed from func(context.Context, DeleteCluster) error to func(context.Context, DeleteCluster) (*WaitGetClusterTerminated[any], error)
  - (*ClustersAPI).Edit: changed from func(context.Context, EditCluster) error to func(context.Context, EditCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).GetAndWait: removed
  - (*ClustersAPI).GetByClusterIdAndWait: removed
  - (*ClustersAPI).Resize: changed from func(context.Context, ResizeCluster) error to func(context.Context, ResizeCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Restart: changed from func(context.Context, RestartCluster) error to func(context.Context, RestartCluster) (*WaitGetClusterRunning[any], error)
  - (*ClustersAPI).Start: changed from func(context.Context, StartCluster) error to func(context.Context, StartCluster) (*WaitGetClusterRunning[any], error)
  - (*CommandExecutionAPI).Cancel: changed from func(context.Context, CancelCommand) error to func(context.Context, CancelCommand) (*WaitCommandStatusCommandExecutionCancelled[any], error)
  - (*CommandExecutionAPI).Create: changed from func(context.Context, CreateContext) (*Created, error) to func(context.Context, CreateContext) (*WaitContextStatusCommandExecutionRunning[Created], error)
  - (*CommandExecutionAPI).Execute: changed from func(context.Context, Command) (*Created, error) to func(context.Context, Command) (*WaitCommandStatusCommandExecutionFinishedOrError[Created], error)
  - (*ClustersAPI).WaitGetClusterRunning: added
  - (*ClustersAPI).WaitGetClusterTerminated: added
  - (*CommandExecutionAPI).Start: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionCancelled: added
  - (*CommandExecutionAPI).WaitCommandStatusCommandExecutionFinishedOrError: added
  - (*CommandExecutionAPI).WaitContextStatusCommandExecutionRunning: added
  - BaseClusterInfo.DataSecurityMode: added
  - BaseClusterInfo.DockerImage: added
  - BaseClusterInfo.SingleUserName: added
  - ClusterAttributes.DataSecurityMode: added
  - ClusterAttributes.DockerImage: added
  - ClusterAttributes.SingleUserName: added
  - ClusterInfo.DockerImage: added
  - CommandExecutorV2: added
  - EditCluster.DataSecurityMode: added
  - EditCluster.DockerImage: added
  - EditCluster.SingleUserName: added
  - WaitCommandStatusCommandExecutionCancelled: added
  - WaitCommandStatusCommandExecutionFinishedOrError: added
  - WaitContextStatusCommandExecutionRunning: added
  - WaitGetClusterRunning: added
  - WaitGetClusterTerminated: added

github.com/databricks/databricks-sdk-go/service/pipelines
  - (*PipelinesAPI).GetAndWait: removed
  - (*PipelinesAPI).GetByPipelineIdAndWait: removed
  - (*PipelinesAPI).Reset: changed from func(context.Context, ResetRequest) error to func(context.Context, ResetRequest) (*WaitGetPipelineRunning[any], error)
  - (*PipelinesAPI).Stop: changed from func(context.Context, StopRequest) error to func(context.Context, StopRequest) (*WaitGetPipelineIdle[any], error)
  - (*PipelinesAPI).WaitGetPipelineIdle: added
  - (*PipelinesAPI).WaitGetPipelineRunning: added
  - WaitGetPipelineIdle: added
  - WaitGetPipelineRunning: added

github.com/databricks/databricks-sdk-go/service/sharing
  - (*ProvidersAPI).ListShares: removed
  - (*ProvidersAPI).ListSharesAll: added

github.com/databricks/databricks-sdk-go/service/workspace
  - ExportFormatAuto: removed
  - ExportRequest.DirectDownload: removed
  - Import.Format: changed from ExportFormat to ImportFormat
  - (*WorkspaceAPI).Download: added
  - (*WorkspaceAPI).ReadFile: added
  - (*WorkspaceAPI).Upload: added
  - (*WorkspaceAPI).WriteFile: added
  - DownloadFormat: added
  - DownloadOption: added
  - ImportFormat: added
  - ImportFormatAuto: added
  - ImportFormatDbc: added
  - ImportFormatHtml: added
  - ImportFormatJupyter: added
  - ImportFormatRMarkdown: added
  - ImportFormatSource: added
  - UploadFormat: added
  - UploadLanguage: added
  - UploadOption: added
  - UploadOverwrite: added

github.com/databricks/databricks-sdk-go/service/catalog
  - (*FunctionsAPI).List: removed
  - (*MetastoresAPI).Assign: changed from func(context.Context, CreateMetastoreAssignment) error to func(context.Context, CreateMetastoreAssignment) error
  - (*MetastoresAPI).Create: changed from func(context.Context, CreateMetastore) (*MetastoreInfo, error) to func(context.Context, CreateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).Update: changed from func(context.Context, UpdateMetastore) (*MetastoreInfo, error) to func(context.Context, UpdateMetastore) (*MetastoreInfo, error)
  - (*MetastoresAPI).UpdateAssignment: changed from func(context.Context, UpdateMetastoreAssignment) error to func(context.Context, UpdateMetastoreAssignment) error
  - (*StorageCredentialsAPI).Create: changed from func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, CreateStorageCredential) (*StorageCredentialInfo, error)
  - (*StorageCredentialsAPI).Update: changed from func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error) to func(context.Context, UpdateStorageCredential) (*StorageCredentialInfo, error)
  - (*TablesAPI).ListSummaries: removed
  - CreateMetastore.Name: removed
  - CreateMetastore.Region: removed
  - CreateMetastore.StorageRoot: removed
  - CreateMetastore: changed from CreateMetastore to CreateMetastore
  - CreateMetastoreAssignment.DefaultCatalogName: removed
  - CreateMetastoreAssignment: changed from CreateMetastoreAssignment to CreateMetastoreAssignment
  - CreateStorageCredential.AwsIamRole: removed
  - CreateStorageCredential.AzureServicePrincipal: removed
  - CreateStorageCredential.Comment: removed
  - CreateStorageCredential.GcpServiceAccountKey: removed
  - CreateStorageCredential.Name: removed
  - CreateStorageCredential.ReadOnly: removed
  - CreateStorageCredential.SkipValidation: removed
  - CreateStorageCredential: changed from CreateStorageCredential to CreateStorageCredential
  - GcpServiceAccountKey: removed
  - ListFunctionsResponse.Schemas: removed
  - MetastoreAssignment.WorkspaceId: changed from string to int64
  - StorageCredentialInfo.GcpServiceAccountKey: removed
  - UpdateMetastore.DeltaSharingOrganizationName: removed
  - UpdateMetastore.DeltaSharingRecipientTokenLifetimeInSeconds: removed
  - UpdateMetastore.DeltaSharingScope: removed
  - UpdateMetastore.Id: removed
  - UpdateMetastore.Name: removed
  - UpdateMetastore.Owner: removed
  - UpdateMetastore.PrivilegeModelVersion: removed
  - UpdateMetastore.StorageRootCredentialId: removed
  - UpdateMetastore: changed from UpdateMetastore to UpdateMetastore
  - UpdateMetastoreAssignment.DefaultCatalogName: removed
  - UpdateMetastoreAssignment: changed from UpdateMetastoreAssignment to UpdateMetastoreAssignment
  - UpdateStorageCredential.AwsIamRole: removed
  - UpdateStorageCredential.AzureServicePrincipal: removed
  - UpdateStorageCredential.Comment: removed
  - UpdateStorageCredential.Force: removed
  - UpdateStorageCredential.GcpServiceAccountKey: removed
  - UpdateStorageCredential.Owner: removed
  - UpdateStorageCredential.ReadOnly: removed
  - UpdateStorageCredential.SkipValidation: removed
  - UpdateStorageCredential: changed from UpdateStorageCredential to UpdateStorageCredential
  - ValidateStorageCredential.GcpServiceAccountKey: removed
  - (*FunctionsAPI).FunctionInfoNameToFullNameMap: added
  - (*FunctionsAPI).ListAll: added
  - (*TablesAPI).ListSummariesAll: added
  - AccountsCreateMetastore: added
  - AccountsCreateMetastoreAssignment: added
  - AccountsCreateStorageCredential: added
  - AccountsUpdateMetastore: added
  - AccountsUpdateMetastoreAssignment: added
  - AccountsUpdateStorageCredential: added
  - AzureManagedIdentity: added
  - ConnectionInfo: added
  - ConnectionType: added
  - ConnectionTypeDatabricks: added
  - ConnectionTypeMysql: added
  - ConnectionTypePostgresql: added
  - ConnectionTypeRedshift: added
  - ConnectionTypeSnowflake: added
  - ConnectionTypeSqldw: added
  - ConnectionTypeSqlserver: added
  - ConnectionsAPI: added
  - ConnectionsService: added
  - CreateConnection: added
  - CreateMetastore.MetastoreInfo: added
  - CreateMetastoreAssignment.MetastoreAssignment: added
  - CreateStorageCredential.CredentialInfo: added
  - CredentialType: added
  - CredentialTypeUsernamePassword: added
  - DatabricksGcpServiceAccountResponse: added
  - DeleteConnectionRequest: added
  - DisableRequest: added
  - GetConnectionRequest: added
  - ListConnectionsResponse: added
  - ListFunctionsResponse.Functions: added
  - ListSystemSchemasRequest: added
  - ListSystemSchemasResponse: added
  - NewConnections: added
  - NewSystemSchemas: added
  - OptionsKvPairs: added
  - PropertiesKvPairs: added
  - StorageCredentialInfo.AzureManagedIdentity: added
  - StorageCredentialInfo.DatabricksGcpServiceAccount: added
  - SystemSchemaInfo: added
  - SystemSchemaInfoState: added
  - SystemSchemaInfoStateDisableinitialized: added
  - SystemSchemaInfoStateEnablecompleted: added
  - SystemSchemaInfoStateEnableinitialized: added
  - SystemSchemaInfoStateUnavailable: added
  - SystemSchemasAPI: added
  - SystemSchemasService: added
  - UpdateConnection: added
  - UpdateMetastore.MetastoreInfo: added
  - UpdateMetastoreAssignment.MetastoreAssignment: added
  - UpdateStorageCredential.CredentialInfo: added
  - ValidateStorageCredential.AzureManagedIdentity: added
  - ValidateStorageCredential.DatabricksGcpServiceAccount: added

github.com/databricks/databricks-sdk-go/service/ml
  - GetModelResponse.RegisteredModel: removed
  - GetModelResponse.RegisteredModelDatabricks: added

github.com/databricks/databricks-sdk-go/service/files
  - FilesAPI: added
  - FilesService: added
  - NewFiles: added

github.com/databricks/databricks-sdk-go/logger
  - SimpleLogger.Level: added

github.com/databricks/databricks-sdk-go/service/provisioning
  - (*WorkspacesAPI).Create: changed from func(context.Context, CreateWorkspaceRequest) (*Workspace, error) to func(context.Context, CreateWorkspaceRequest) (*WaitGetWorkspaceRunning[Workspace], error)
  - (*WorkspacesAPI).Update: changed from func(context.Context, UpdateWorkspaceRequest) error to func(context.Context, UpdateWorkspaceRequest) (*WaitGetWorkspaceRunning[any], error)
  - (*WorkspacesAPI).WaitGetWorkspaceRunning: added
  - WaitGetWorkspaceRunning: added

github.com/databricks/databricks-sdk-go/service/iam
  - AccountAccessControlAPI: added
  - AccountAccessControlProxyAPI: added
  - AccountAccessControlProxyService: added
  - AccountAccessControlService: added
  - GetAssignableRolesForResourceRequest: added
  - GetAssignableRolesForResourceResponse: added
  - GetRuleSetRequest: added
  - GrantRule: added
  - NewAccountAccessControl: added
  - NewAccountAccessControlProxy: added
  - RuleSetResponse: added
  - RuleSetUpdateRequest: added
  - UpdateRuleSetRequest: added

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.3 to 1.8.4 ([#406](https://github.com/databricks/databricks-sdk-go/pull/406)).
 * Bump golang.org/x/mod from 0.10.0 to 0.11.0 ([#515](https://github.com/databricks/databricks-sdk-go/pull/515)).
 * Bump golang.org/x/oauth2 from 0.8.0 to 0.9.0 ([#498](https://github.com/databricks/databricks-sdk-go/pull/498)).
 * Bump google.golang.org/api from 0.123.0 to 0.127.0 ([#405](https://github.com/databricks/databricks-sdk-go/pull/405), [#425](https://github.com/databricks/databricks-sdk-go/pull/425), [#429](https://github.com/databricks/databricks-sdk-go/pull/429)).

## Release v0.9.0 (2023-05-22)

 * Added more usage examples for `go doc` and Go Packages ([#389](https://github.com/databricks/databricks-sdk-go/pull/389)).
 * Make u2m authentication work with new CLI ([#394](https://github.com/databricks/databricks-sdk-go/pull/394)).
 * Update from OpenAPI spec (19 may) ([#390](https://github.com/databricks/databricks-sdk-go/pull/390)).

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.2 to 1.8.3 ([#393](https://github.com/databricks/databricks-sdk-go/pull/393)).
 * Bump google.golang.org/api from 0.122.0 to 0.123.0 ([#392](https://github.com/databricks/databricks-sdk-go/pull/392)).

## Release v0.8.1 (2023-05-11)

 * Added `in` codegen function ([#387](https://github.com/databricks/databricks-sdk-go/pull/387)).
 * Fixed mlflow acceptance tests ([#378](https://github.com/databricks/databricks-sdk-go/pull/378)).
 * Fixed MLflow integration test and removed workaround for `DELETE` query parameter ([#380](https://github.com/databricks/databricks-sdk-go/pull/380)).
 * Make clusters acceptance tests robust to duplicate cluster names ([#381](https://github.com/databricks/databricks-sdk-go/pull/381)).
 * Remove dead code from apierr/errors.go ([#376](https://github.com/databricks/databricks-sdk-go/pull/376)).
 * Serialize params to request body on delete ([#383](https://github.com/databricks/databricks-sdk-go/pull/383)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.7.0 to 0.8.0 ([#385](https://github.com/databricks/databricks-sdk-go/pull/385)).
 * Bump google.golang.org/api from 0.118.0 to 0.122.0 ([#382](https://github.com/databricks/databricks-sdk-go/pull/382), [#386](https://github.com/databricks/databricks-sdk-go/pull/386)).

## Release v0.8.0 (2023-04-20)

 * Added more code generation utilities ([#369](https://github.com/databricks/databricks-sdk-go/pull/369)).
 * Body logger for non-JSON payload as well ([#365](https://github.com/databricks/databricks-sdk-go/pull/365)).
 * Cleanup ephemeral groups in integration tests ([#368](https://github.com/databricks/databricks-sdk-go/pull/368)).
 * Fixed external entity generation ([#372](https://github.com/databricks/databricks-sdk-go/pull/372)).
 * Skip loading default profile if host is already configured ([#363](https://github.com/databricks/databricks-sdk-go/pull/363)).
 * Update debug messages in config loader to display correct path ([#362](https://github.com/databricks/databricks-sdk-go/pull/362)).

Dependency updates:

 * Bump golang.org/x/oauth2 from 0.6.0 to 0.7.0 ([#364](https://github.com/databricks/databricks-sdk-go/pull/364)).
 * Bump google.golang.org/api from 0.115.0 to 0.116.0 ([#361](https://github.com/databricks/databricks-sdk-go/pull/361)).
 * Bump google.golang.org/api from 0.116.0 to 0.118.0 ([#367](https://github.com/databricks/databricks-sdk-go/pull/367)).

API changes:

 * Moved `clusterpolicies` APIs to `compute` package.
 * Moved `clusters` APIs to `compute` package.
 * Moved `commands` APIs to `compute` package.
 * Moved `globalinitscripts` APIs to `compute` package.
 * Moved `instancepools` APIs to `compute` package.
 * Moved `scim` APIs to `iam` package.
 * Moved `permissions` APIs to `iam` package.
 * Moved `ipaccesslists` APIs to `settings` package.
 * Moved `tokenmanagement` APIs to `settings` package.
 * Moved `tokens` APIs to `settings` package.
 * Moved `workspaceconf` APIs to `settings` package.
 * Moved `gitcredentials` APIs to `workspace` package.
 * Moved `repos` APIs to `workspace` package.
 * Moved `secrets` APIs to `workspace` package.
 * Split `unitcatalog` package to `catalog` and `sharing`.
 * Renamed `mlflow` package to `ml`.
 * Renamed `dbfs` package to `files`.
 * Renamed `deployment` package to `provisioning`.
 * Renamed `endpoints` package to `serving`.
 * Renamed `clusters.List` type to `compute.ListClustersRequest`.
 * Renamed `jobs.ListRuns` type to `jobs.ListRunsRequest`.
 * Renamed `jobs.ExportRun` type to `jobs.ExportRunRequest`.
 * Renamed `clusterpolicies.List` type to `compute.ListClusterPoliciesRequest`.
 * Renamed `jobs.List` type to `jobs.ListJobsRequest`.
 * Renamed `permissions.GetPermissionLevels` type to `iam.GetPermissionLevelsRequest`.
 * Renamed `pipelines.ListPipelineEvents` type to `pipelines.ListPipelineEventsRequest`.
 * Renamed `pipelines.ListPipelines` type to `pipelines.ListPipelinesRequest`.
 * Renamed `workspaceconf.GetStatus` type to `settings.GetStatusRequest`.
 * Renamed `repos.List` type to `workspace.ListReposRequest`.
 * Renamed `tokenmanagement.List` type to `settings.ListTokenManagementRequest`.
 * Renamed `workspace.Export` type to `workspace.ExportRequest`.
 * Renamed `workspace.List` type to `workspace.ListWorkspaceRequest`.

## Release v0.7.0 (2023-04-05)

 * Update from OpenAPI ([#359](https://github.com/databricks/databricks-sdk-go/pull/359)).
 * Experimental credentials provider via local server ([#340](https://github.com/databricks/databricks-sdk-go/pull/340)).
 * Added `isTesting` marker to support `ResourceFixture` in Terraform ([#358](https://github.com/databricks/databricks-sdk-go/pull/358)).

Dependency updates:

 * Bump golang.org/x/mod from 0.9.0 to 0.10.0 ([#356](https://github.com/databricks/databricks-sdk-go/pull/356)).
 * Bump google.golang.org/api from 0.114.0 to 0.115.0 ([#357](https://github.com/databricks/databricks-sdk-go/pull/357)).

## Release v0.6.0 (2023-04-03)

 * Added type to represent a loaded configuration file ([#349](https://github.com/databricks/databricks-sdk-go/pull/349)).
 * Added named `Wait` the level of services ([#348](https://github.com/databricks/databricks-sdk-go/pull/348)).
 * Pass `azure_client_id` to Azure MSI if specified ([#354](https://github.com/databricks/databricks-sdk-go/pull/354)).

Dependency updates:

 * Bump google.golang.org/api from 0.112.0 to 0.114.0 ([#344](https://github.com/databricks/databricks-sdk-go/pull/344)).

## Release v0.5.0 (2023-03-14)

 * Added Delta Live Tables events method to `pipelines.PipelinesAPI` ([#339](https://github.com/databricks/databricks-sdk-go/pull/339)).
 * Added `id` field to the `SparkVersionRequest` for easier integration with Terraform's [databricks_spark_version](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/spark_version) data source ([#335](https://github.com/databricks/databricks-sdk-go/pull/335)).
 * Improve auth fixture capturing ([#336](https://github.com/databricks/databricks-sdk-go/pull/336), [#337](https://github.com/databricks/databricks-sdk-go/pull/337)).
 * Make logger context aware and added examples of usage together with [`zerolog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/zerolog) and [`slog`](https://github.com/databricks/databricks-sdk-go/tree/main/examples/slog) ([#333](https://github.com/databricks/databricks-sdk-go/pull/333)).
 * Update from OpenAPI spec and added new APIs ([#338](https://github.com/databricks/databricks-sdk-go/pull/338)).

Dependency updates:

 * Bump google.golang.org/api from 0.111.0 to 0.112.0 ([#334](https://github.com/databricks/databricks-sdk-go/pull/334)).

## Release v0.4.1 (2023-03-09)

 * Added an option to configure the location of the bricks CLI ([#330](https://github.com/databricks/databricks-sdk-go/pull/330)).
 * Added support for Azure CLI authentication on Azure China and Azure GovCloud ([#331](https://github.com/databricks/databricks-sdk-go/pull/331)).

## Release v0.4.0 (2023-03-08)

 * Added `id` field to the `NodeTypeRequest` for easier integration with Terraform's [databricks_node_type](https://registry.terraform.io/providers/databricks/databricks/latest/docs/data-sources/node_type) data source ([#325](https://github.com/databricks/databricks-sdk-go/pull/325)).
 * Don't load `~/.databrickscfg` if `azure-cli` auth is implicitly configured ([#324](https://github.com/databricks/databricks-sdk-go/pull/324)).
 * Fixed newline in codegen comments ([#326](https://github.com/databricks/databricks-sdk-go/pull/326)).
 * Update from OpenAPI ([#320](https://github.com/databricks/databricks-sdk-go/pull/320), [#328](https://github.com/databricks/databricks-sdk-go/pull/328)).

API changes:

* Renamed `deployment.AwsCredentials` to `deployment.CreateCredentialAwsCredentials`.
* Renamed `deployment.StsRole` to `deployment.CreateCredentialStsRole`.
* Removed schedules support from `sql.AlertsAPI`.

Dependency updates:

 * Bump github.com/stretchr/testify from 1.8.1 to 1.8.2 ([#318](https://github.com/databricks/databricks-sdk-go/pull/318)).
 * Bump golang.org/x/mod from 0.6.0-dev.0.20220419223038-86c51ed26bb4 to 0.8.0 ([#316](https://github.com/databricks/databricks-sdk-go/pull/316)).
 * Bump golang.org/x/mod from 0.8.0 to 0.9.0 ([#323](https://github.com/databricks/databricks-sdk-go/pull/323)).
 * Bump golang.org/x/oauth2 from 0.5.0 to 0.6.0 ([#322](https://github.com/databricks/databricks-sdk-go/pull/322)).
 * Bump golang.org/x/time from 0.0.0-20210723032227-1f47c861a9ac to 0.3.0 ([#317](https://github.com/databricks/databricks-sdk-go/pull/317)).
 * Bump google.golang.org/api from 0.110.0 to 0.111.0 ([#319](https://github.com/databricks/databricks-sdk-go/pull/319)).

## Release v0.3.3 (2023-02-24)

 * Allow AAD SPN authentication on Databricks Account level ([#311](https://github.com/databricks/databricks-sdk-go/pull/311)).
 * Port auth tests from the JS SDK ([#313](https://github.com/databricks/databricks-sdk-go/pull/313)).
 * Skip loading `~/.databrickscfg` when not required ([#314](https://github.com/databricks/databricks-sdk-go/pull/314)).

Dependency updates:

 * Bump golang.org/x/net from 0.6.0 to 0.7.0 ([#312](https://github.com/databricks/databricks-sdk-go/pull/312)).

## Release v0.3.2 (2023-02-17)

 * Always use the latest value for user agent key ([#309](https://github.com/databricks/databricks-sdk-go/pull/309)).

## Release v0.3.1 (2023-02-17)

 * Change APIError to use pointer receiver ([#298](https://github.com/databricks/databricks-sdk-go/pull/298)).
 * Drop duplicate prefix in randomized email ([#299](https://github.com/databricks/databricks-sdk-go/pull/299)).
 * Expand semver pattern to be compliant with https://semver.org ([#302](https://github.com/databricks/databricks-sdk-go/pull/302)).
 * Fix `apierr.APIError` pointer receivers ([#307](https://github.com/databricks/databricks-sdk-go/pull/307)).
 * Skip loading config if auth is already explicitly configured directly ([#306](https://github.com/databricks/databricks-sdk-go/pull/306)).
 * Sync fixes for smallest node selection from Terraform ([#301](https://github.com/databricks/databricks-sdk-go/pull/301)).
 * Updated from OpenAPI spec ([#305](https://github.com/databricks/databricks-sdk-go/pull/305)).

Dependency updates:

 * Bump google.golang.org/api from 0.109.0 to 0.110.0 ([#303](https://github.com/databricks/databricks-sdk-go/pull/303)).

## Release v0.3.0 (2023-02-03)

 * Added support for GCP deployment APIs in Public Preview.
 * Added new node type selector ([#287](https://github.com/databricks/databricks-sdk-go/pull/287)).
 * Added Azure MSI auth ([#279](https://github.com/databricks/databricks-sdk-go/pull/279)).
 * Added validation of strings in user agent package ([#291](https://github.com/databricks/databricks-sdk-go/pull/291)).
 * Added logging for error responses ([#289](https://github.com/databricks/databricks-sdk-go/pull/289)).
 * Update implementation of smallest node selection to match terraform's ([#273](https://github.com/databricks/databricks-sdk-go/pull/273)).
 * Added experimental OAuth support ([#276](https://github.com/databricks/databricks-sdk-go/pull/276), [#266](https://github.com/databricks/databricks-sdk-go/pull/266), [#292](https://github.com/databricks/databricks-sdk-go/pull/292)).
 * Fixed `google-credentials` to take precedence over `google-id` ([#283](https://github.com/databricks/databricks-sdk-go/pull/283)).
 * Test with Go 1.20 ([#295](https://github.com/databricks/databricks-sdk-go/pull/295)).
 * Regenerate from OpenAPI ([#293](https://github.com/databricks/databricks-sdk-go/pull/293), [#294](https://github.com/databricks/databricks-sdk-go/pull/294), [#296](https://github.com/databricks/databricks-sdk-go/pull/296), [#282](https://github.com/databricks/databricks-sdk-go/pull/282), [#269](https://github.com/databricks/databricks-sdk-go/pull/269), [#272](https://github.com/databricks/databricks-sdk-go/pull/272)).
 * Added `replaceAll` and `lowerFirst` template functions ([#288](https://github.com/databricks/databricks-sdk-go/pull/288), [#277](https://github.com/databricks/databricks-sdk-go/pull/277)).

Dependency updates:

 * Bump google.golang.org/api from 0.105.0 to 0.109.0 ([#271](https://github.com/databricks/databricks-sdk-go/pull/271), [#280](https://github.com/databricks/databricks-sdk-go/pull/280), [#284](https://github.com/databricks/databricks-sdk-go/pull/284), [#290](https://github.com/databricks/databricks-sdk-go/pull/290)).

API changes:

 * Renamed `clusters.CreateCluster` to `clusters.BaseClusterInfo`.
 * Renamed `jobs.Job` to `jobs.BaseJob`.
 * Renamed `jobs.Run` to `jobs.BaseRun`.
 * Fixed `commands.Results` `Schema` type from `[][]any` to `[]map[string]any`.

## Release v0.2.0 (2022-12-23)

 * Added `DATABRICKS_AUTH_TYPE` environment variable ([#248](https://github.com/databricks/databricks-sdk-go/pull/248)).
 * Added Policy Families API ([#263](https://github.com/databricks/databricks-sdk-go/pull/263)).
 * Added experimental `ErrCannotConfigureAuth` and `ErrNotAccountClient` ([#237](https://github.com/databricks/databricks-sdk-go/pull/237), [#238](https://github.com/databricks/databricks-sdk-go/pull/238)).
 * Added DBFS file handle that supports both reading and writing ([#261](https://github.com/databricks/databricks-sdk-go/pull/261)).
 * Added `io.WriterTo` for DBFS file reader ([#249](https://github.com/databricks/databricks-sdk-go/pull/249)).
 * Added `pflag.Value` interfaces for enums ([#234](https://github.com/databricks/databricks-sdk-go/pull/234)).
 * Added support for adding custom HTTP visitors per request ([#230](https://github.com/databricks/databricks-sdk-go/pull/230)).
 * Added support for raw body as byte slice if requested ([#247](https://github.com/databricks/databricks-sdk-go/pull/247)).
 * Improved callbacks for polling the status of long-running operations ([#258](https://github.com/databricks/databricks-sdk-go/pull/258)).
 * Improved rendering of HTTP links in godoc ([#229](https://github.com/databricks/databricks-sdk-go/pull/229)).
 * Updated field types in the Jobs API from spec ([#259](https://github.com/databricks/databricks-sdk-go/pull/259)).
 * Multiple OpenAPI consistency passes ([#254](https://github.com/databricks/databricks-sdk-go/pull/254), [#241](https://github.com/databricks/databricks-sdk-go/pull/241), [#243](https://github.com/databricks/databricks-sdk-go/pull/243), [#255](https://github.com/databricks/databricks-sdk-go/pull/255), [#236](https://github.com/databricks/databricks-sdk-go/pull/236)).

API changes:
 * Renamed `IsAccountsClient` to `IsAccountClient` ([#231](https://github.com/databricks/databricks-sdk-go/pull/231)).
 * `w.ClusterPolicies.ListAll` now takes `clusterpolicies.List` as an argument.
 * `github.com/databricks/databricks-sdk-go/service/dbsql` package is renamed to `github.com/databricks/databricks-sdk-go/service/sql`.
 * `w.DataSources.ListDataSources` is renamed to `w.DataSources.List`.
 * `w.Queries.CreateQuery` is renamed to `w.Queries.CreateQuery`.
 * `w.Queries.DeleteQueryByQueryId` is renamed to `w.Queries.DeleteByQueryId`.
 * `w.Queries.GetQueryByQueryId` is renamed to `w.Queries.GetByQueryId`.
 * `w.Queries.UpdateQuery` is renamed to `w.Queries.Update`.
 * `w.Alerts.DeleteAlertByAlertId` is renamed to `w.Alerts.DeleteByAlertId`.
 * `w.Alerts.UpdateAlert` is renamed to `w.Alerts.Update`.
 * `w.Alerts.GetAlertByAlertId` is renamed to `w.Alerts.GetByAlertId`.
 * `w.Alerts.ListAlerts` is renamed to `w.Alerts.List`.
 * `w.Dashboards.CreateDashboard` is renamed to `w.Dashboards.Create`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.GetByDashboardId`.
 * `w.Dashboards.ListDashboardsAll` is renamed to `w.Dashboards.ListAll`.
 * `w.Dashboards.DeleteDashboardByDashboardId` is renamed to `w.Dashboards.DeleteByDashboardId`.
 * `w.Dashboards.RestoreDashboard` is renamed to `w.Dashboards.Restore`.
 * `deployment.CreateCustomerManagedKeyRequest` now takes `deployment.KeyUseCase` enum.
 * `w.GlobalInitScripts.CreateScript` is renamed to `w.GlobalInitScripts.Create`.
 * `w.GlobalInitScripts.DeleteScriptByScriptId` is renamed to `w.GlobalInitScripts.DeleteByScriptId`.
 * `w.GlobalInitScripts.UpdateScript` is renamed to `w.GlobalInitScripts.Update`.
 * `w.GlobalInitScripts.GetScriptByScriptId` is renamed to `w.GlobalInitScripts.GetByScriptId`.
 * `w.GlobalInitScripts.ListScriptsAll` is renamed to `w.GlobalInitScripts.ListAll`.
 * `jobs.ResetJob.NewSettings` is now required field.
 * `w.Pipelines.CreatePipeline` is renamed to `w.Pipelines.Create`.
 * `w.Pipelines.DeletePipelineByPipelineId` is renamed to `w.Pipelines.DeleteByPipelineId`.
 * `w.Pipelines.UpdatePipeline` is renamed to `w.Pipelines.Update`.
 * `w.Pipelines.GetPipelineByPipelineId` is renamed to `w.Pipelines.GetByPipelineId`.
 * `w.StorageCredentials.Update` now also returns an entity.
 * `w.ExternalLocations.Update` now also returns an entity.
 * `w.Metastores.Update` now also returns an entity.
 * `unitycatalog.CreateMetastoreAssignment.WorkspaceId` type changed from `int` to `int64`.
 * `unitycatalog.UnassignRequest.WorkspaceId` type changed from `int` to `int64`.
 * `w.Catalogs.Update` now also returns an entity.
 * `w.Schemas.Update` now also returns an entity.
 * `w.Providers.Update` now also returns an entity.
 * `w.Shares.Update` now also returns an entity.
 * `WarehousesAPI` service moved to `github.com/databricks/databricks-sdk-go/service/sql` package.
 * `w.Warehouses.CreateWarehouseAndWait` renamed to `w.Warehouses.CreateAndWait`.
 * `w.Warehouses.DeleteWarehouseByIdAndWait` renamed to `w.Warehouses.DeleteByIdAndWait`.
 * `w.Warehouses.EditWarehouse` renamed to `w.Warehouses.Edit`.
 * `w.Warehouses.GetWarehouseById` renamed to `w.Warehouses.GetById`.
 * `w.Warehouses.ListWarehousesAll` renamed to `w.Warehouses.ListAll`.
 * Removed `w.Dbfs.Overwrite` in favor of `w.Dbfs.Open("....", dbfs.FileModeOverwrite|dbfs.FileModeWrite)`.
 * Added third required argument to `w.Dbfs.Open`.

Code generation:

 * Added concept of `main` service for the package ([#239](https://github.com/databricks/databricks-sdk-go/pull/239)).
 * Added entity primitives check ([#242](https://github.com/databricks/databricks-sdk-go/pull/242)).
 * Added helpers for CRUD generation ([#246](https://github.com/databricks/databricks-sdk-go/pull/246)).
 * Added more entity-generation utils ([#257](https://github.com/databricks/databricks-sdk-go/pull/257)).
 * Dynamically generate `.gitattributes` ([#244](https://github.com/databricks/databricks-sdk-go/pull/244)).
 * Fixed required order fields ([#245](https://github.com/databricks/databricks-sdk-go/pull/245)).
 * Parse summary from descriptions ([#228](https://github.com/databricks/databricks-sdk-go/pull/228)).
 * Print error on formatter failure ([#235](https://github.com/databricks/databricks-sdk-go/pull/235)).
 * Update usage string in generator ([#260](https://github.com/databricks/databricks-sdk-go/pull/260)).
 * Fixed order of host completion ([#233](https://github.com/databricks/databricks-sdk-go/pull/233)).

Dependency updates:

 * Bump google.golang.org/api from 0.103.0 to 0.105.0 ([#232](https://github.com/databricks/databricks-sdk-go/pull/232), [#252](https://github.com/databricks/databricks-sdk-go/pull/252)).

## Release v0.1.1 (2022-12-02)

* Improved usage instructions and added more documentation for services.

## Release v0.1.0 (2022-11-30)

 * Initial release of the Databricks SDK for Go.
