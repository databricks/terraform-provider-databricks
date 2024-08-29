# Introduction
This document contains the good practices for any new resource or data source that will be introduced or migrated to the plugin framework.  


## Migrating resource to plugin framework
Ideally there shouldn't be any behaviour change when migrating a resource or data source to either Go SDk or Plugin Framework. - Please make sure there are no breaking differences due to changes in schema by running: `make diff-schema`. 
- Integration tests shouldn't require any major changes.   


## Code Organization
Each resource should go into it's separate package eg: `volume` package will contain both resource, data sources and other utils specific to volumes. Tests (both unit and integration tests) will also remain in this package. 

Note: Only Docs will stay under root docs/ directory.


## Code Conventions
1. Make sure the resource or data source implemented is of the right type:
    ```golang
    var _ resource.ResourceWithConfigure = &QualityMonitorResource{}
    var _ datasource.DataSourceWithConfigure = &VolumesDataSource{}
    ```
2. To get the databricks client, `func (*common.DatabricksClient).GetWorkspaceClient()` or `func (*common.DatabricksClient).GetAccountClient()` will be used instead of directly using the underlying `WorkspaceClient()`, `AccountClient()` functions respectively.  
3. Any method that returns the diagnostics should be called inline while appending diagnostics in response. Example:
    ```golang
    resp.Diagnostics.Append(req.Plan.Get(ctx, &monitorInfoTfSDK)...)
    if resp.Diagnostics.HasError() {
        return
    }
    ```
    is preferred over the following:
    ```golang
    diags := req.Plan.Get(ctx, &monitorInfoTfSDK)
    if diags.HasError() {
        resp.Diagnostics.Append(diags...)
        return
    }
    ```
4. Any method returning an error should directly be followed by appending that to the diagnostics. 