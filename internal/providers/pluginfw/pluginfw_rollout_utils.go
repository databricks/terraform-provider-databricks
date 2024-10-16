package pluginfw

// This file contains all of the utils for controlling the plugin framework rollout.
// For migrated resources and data sources, we can add them to the two maps below to have them registered with the plugin framework.
// Users can manually specify resources and data sources to use SDK V2 instead of the plugin framework by setting the USE_SDK_V2_RESOURCES and USE_SDK_V2_DATA_SOURCES environment variables.
//
// Example: USE_SDK_V2_RESOURCES="databricks_library" would force the library resource to use SDK V2 instead of the plugin framework.

import (
	"os"
	"strings"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/cluster"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/library"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/qualitymonitor"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/registered_model"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/volume"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Map of resources that have been migrated from SDK V2 to plugin framework
var migratedResourceMap = map[string]func() resource.Resource{
	"databricks_qualitymonitor": qualitymonitor.ResourceQualityMonitor,
	"databricks_library":        library.ResourceLibrary,
}

// Map of data sources that have been migrated from SDK V2 to plugin framework
var migratedDataSourceMap = map[string]func() datasource.DataSource{
	"databricks_qualitymonitor":   cluster.DataSourceCluster,
	"databricks_volumes":          volume.DataSourceVolumes,
	"databricks_registered_model": registered_model.DataSourceRegisteredModel,
}

// GetUseSdkV2DataSources is a helper function to get name of resources that should use SDK V2 instead of plugin framework
func getUseSdkV2Resources() []string {
	useSdkV2 := os.Getenv("USE_SDK_V2_RESOURCES")
	if useSdkV2 == "" {
		return []string{}
	}
	return strings.Split(useSdkV2, ",")
}

// GetUseSdkV2DataSources is a helper function to get name of data sources that should use SDK V2 instead of plugin framework
func getUseSdkV2DataSources() []string {
	useSdkV2 := os.Getenv("USE_SDK_V2_DATA_SOURCES")
	if useSdkV2 == "" {
		return []string{}
	}
	return strings.Split(useSdkV2, ",")
}

// Helper function to check if a resource should use be in SDK V2 instead of plugin framework
func shouldUseSdkV2Resource(resourceName string) bool {
	useSdkV2Resources := getUseSdkV2Resources()
	for _, sdkV2Resource := range useSdkV2Resources {
		if resourceName == sdkV2Resource {
			return true
		}
	}
	return false
}

// Helper function to check if a data source should use be in SDK V2 instead of plugin framework
func shouldUseSdkV2DataSource(dataSourceName string) bool {
	sdkV2DataSources := getUseSdkV2DataSources()
	for _, sdkV2DataSource := range sdkV2DataSources {
		if dataSourceName == sdkV2DataSource {
			return true
		}
	}
	return false
}

// getPluginFrameworkResourcesToRegister is a helper function to get the list of resources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkResourcesToRegister() []func() resource.Resource {
	var resources []func() resource.Resource

	// Loop through the map and add resources if they're not specifically marked to use the SDK V2
	for name, resourceFunc := range migratedResourceMap {
		if !shouldUseSdkV2Resource(name) {
			resources = append(resources, resourceFunc)
		}
	}

	return resources
}

// getPluginFrameworkDataSourcesToRegister is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkDataSourcesToRegister() []func() datasource.DataSource {
	var dataSources []func() datasource.DataSource

	// Loop through the map and add data sources if they're not specifically marked to use the SDK V2
	for name, dataSourceFunc := range migratedDataSourceMap {
		if !shouldUseSdkV2DataSource(name) {
			dataSources = append(dataSources, dataSourceFunc)
		}
	}

	return dataSources
}

// GetSdkV2ResourcesToRemove is a helper function to get the list of resources that are migrated away from sdkv2 to plugin framework
func GetSdkV2ResourcesToRemove() []string {
	resourcesToRemove := []string{}
	for name, _ := range migratedResourceMap {
		if !shouldUseSdkV2Resource(name) {
			resourcesToRemove = append(resourcesToRemove, name)
		}
	}
	return resourcesToRemove
}

// GetSdkV2DataSourcesToRemove is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func GetSdkV2DataSourcesToRemove() []string {
	dataSourcesToRemove := []string{}
	for name, _ := range migratedDataSourceMap {
		if !shouldUseSdkV2DataSource(name) {
			dataSourcesToRemove = append(dataSourcesToRemove, name)
		}
	}
	return dataSourcesToRemove
}
