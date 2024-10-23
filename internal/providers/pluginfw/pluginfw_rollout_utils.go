package pluginfw

// This file contains all of the utils for controlling the plugin framework rollout.
// For migrated resources and data sources, we can add them to the two maps below to have them registered with the plugin framework.
// Users can manually specify resources and data sources to use SDK V2 instead of the plugin framework by setting the USE_SDK_V2_RESOURCES and USE_SDK_V2_DATA_SOURCES environment variables.
//
// Example: USE_SDK_V2_RESOURCES="databricks_library" would force the library resource to use SDK V2 instead of the plugin framework.

import (
	"context"
	"os"
	"strings"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/library"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/notificationdestinations"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/qualitymonitor"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/registered_model"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/volume"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// List of resources that have been migrated from SDK V2 to plugin framework
var migratedResources = []func() resource.Resource{
	qualitymonitor.ResourceQualityMonitor,
	library.ResourceLibrary,
}

// List of data sources that have been migrated from SDK V2 to plugin framework
var migratedDataSources = []func() datasource.DataSource{
	// TODO: Add DataSourceCluster back in after fixing unit tests.
	// cluster.DataSourceCluster,
	volume.DataSourceVolumes,
}

// List of resources that have been onboarded to the plugin framework - not migrated from sdkv2.
var onboardedResources = []func() resource.Resource{
	// TODO Add resources here
}

// List of data sources that have been onboarded to the plugin framework - not migrated from sdkv2.
var onboardedDataSources = []func() datasource.DataSource{
	registered_model.DataSourceRegisteredModel,
	notificationdestinations.DataSourceNotificationDestinations,
}

// GetUseSdkV2DataSources is a helper function to get name of resources that should use SDK V2 instead of plugin framework
func getUseSdkV2Resources(ctx context.Context) []string {
	useSdkV2 := os.Getenv("USE_SDK_V2_RESOURCES")
	useSdkV2Ctx := ctx.Value("USE_SDK_V2_RESOURCES")
	combinedNames := ""
	if useSdkV2 != "" && useSdkV2Ctx != "" {
		combinedNames = useSdkV2 + "," + useSdkV2Ctx.(string)
	} else {
		combinedNames = useSdkV2 + useSdkV2Ctx.(string)
	}
	return strings.Split(combinedNames, ",")
}

// GetUseSdkV2DataSources is a helper function to get name of data sources that should use SDK V2 instead of plugin framework
func getUseSdkV2DataSources(ctx context.Context) []string {
	useSdkV2 := os.Getenv("USE_SDK_V2_DATA_SOURCES")
	useSdkV2Ctx := ctx.Value("USE_SDK_V2_DATA_SOURCES")
	combinedNames := ""
	if useSdkV2 != "" && useSdkV2Ctx != "" {
		combinedNames = useSdkV2 + "," + useSdkV2Ctx.(string)
	} else {
		combinedNames = useSdkV2 + useSdkV2Ctx.(string)
	}
	return strings.Split(combinedNames, ",")
}

// Helper function to check if a resource should use be in SDK V2 instead of plugin framework
func shouldUseSdkV2Resource(ctx context.Context, resourceName string) bool {
	useSdkV2Resources := getUseSdkV2Resources(ctx)
	for _, sdkV2Resource := range useSdkV2Resources {
		if resourceName == sdkV2Resource {
			return true
		}
	}
	return false
}

// Helper function to check if a data source should use be in SDK V2 instead of plugin framework
func shouldUseSdkV2DataSource(ctx context.Context, dataSourceName string) bool {
	sdkV2DataSources := getUseSdkV2DataSources(ctx)
	for _, sdkV2DataSource := range sdkV2DataSources {
		if dataSourceName == sdkV2DataSource {
			return true
		}
	}
	return false
}

// getPluginFrameworkResourcesToRegister is a helper function to get the list of resources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkResourcesToRegister(ctx context.Context) []func() resource.Resource {
	var resources []func() resource.Resource

	// Loop through the map and add resources if they're not specifically marked to use the SDK V2
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if !shouldUseSdkV2Resource(ctx, name) {
			resources = append(resources, resourceFunc)
		}
	}

	return append(resources, onboardedResources...)
}

// getPluginFrameworkDataSourcesToRegister is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkDataSourcesToRegister(ctx context.Context) []func() datasource.DataSource {
	var dataSources []func() datasource.DataSource

	// Loop through the map and add data sources if they're not specifically marked to use the SDK V2
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if !shouldUseSdkV2DataSource(ctx, name) {
			dataSources = append(dataSources, dataSourceFunc)
		}
	}

	return append(dataSources, onboardedDataSources...)
}

func getResourceName(resourceFunc func() resource.Resource) string {
	resp := resource.MetadataResponse{}
	resourceFunc().Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "databricks"}, &resp)
	return resp.TypeName
}

func getDataSourceName(dataSourceFunc func() datasource.DataSource) string {
	resp := datasource.MetadataResponse{}
	dataSourceFunc().Metadata(context.Background(), datasource.MetadataRequest{ProviderTypeName: "databricks"}, &resp)
	return resp.TypeName
}

// GetSdkV2ResourcesToRemove is a helper function to get the list of resources that are migrated away from sdkv2 to plugin framework
func GetSdkV2ResourcesToRemove(ctx context.Context) []string {
	resourcesToRemove := []string{}
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if !shouldUseSdkV2Resource(ctx, name) {
			resourcesToRemove = append(resourcesToRemove, name)
		}
	}
	return resourcesToRemove
}

// GetSdkV2DataSourcesToRemove is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func GetSdkV2DataSourcesToRemove(ctx context.Context) []string {
	dataSourcesToRemove := []string{}
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if !shouldUseSdkV2DataSource(ctx, name) {
			dataSourcesToRemove = append(dataSourcesToRemove, name)
		}
	}
	return dataSourcesToRemove
}
