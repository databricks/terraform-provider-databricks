package pluginfw

// This file contains all of the utils for controlling the plugin framework rollout.
// For migrated resources and data sources, we can add them to the two maps below to have them registered with the plugin framework.
// Users can manually specify resources and data sources to use SDK V2 instead of the plugin framework by setting the USE_SDK_V2_RESOURCES and USE_SDK_V2_DATA_SOURCES environment variables.
//
// Example: USE_SDK_V2_RESOURCES="databricks_library" would force the library resource to use SDK V2 instead of the plugin framework.

import (
	"context"
	"os"
	"slices"
	"strings"

	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/resources/cluster"
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
	// TODO: Add DataSourceCluster into migratedDataSources after fixing unit tests.
	cluster.DataSourceCluster, // This is still using the staging name (with pluginframework suffix)
}

type sdkV2FallbackOptions struct {
	resourceFallbacks   []string
	dataSourceFallbacks []string
}

// SdkV2FallbackOption is an interface for acceptance tests to specify resources / data sources to fallback to SDK V2
type SdkV2FallbackOption interface {
	Apply(*sdkV2FallbackOptions)
}

type sdkV2ResourceFallback struct {
	resourceFallbacks []string
}

func (o *sdkV2ResourceFallback) Apply(options *sdkV2FallbackOptions) {
	options.resourceFallbacks = o.resourceFallbacks
}

// WithSdkV2ResourceFallbacks is a helper function to specify resources to fallback to SDK V2
func WithSdkV2ResourceFallbacks(fallbacks []string) SdkV2FallbackOption {
	return &sdkV2ResourceFallback{resourceFallbacks: fallbacks}
}

type sdkv2DataSourceFallback struct {
	dataSourceFallbacks []string
}

func (o *sdkv2DataSourceFallback) Apply(options *sdkV2FallbackOptions) {
	options.dataSourceFallbacks = o.dataSourceFallbacks
}

// WithSdkV2DataSourceFallbacks is a helper function to specify data sources to fallback to SDK V2
func WithSdkV2DataSourceFallbacks(fallbacks []string) SdkV2FallbackOption {
	return &sdkv2DataSourceFallback{dataSourceFallbacks: fallbacks}
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
func getPluginFrameworkResourcesToRegister(sdkV2Fallbacks ...SdkV2FallbackOption) []func() resource.Resource {
	fallbackOption := sdkV2FallbackOptions{}
	for _, o := range sdkV2Fallbacks {
		o.Apply(&fallbackOption)
	}

	var resources []func() resource.Resource

	// Loop through the map and add resources if they're not specifically marked to use the SDK V2
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if !shouldUseSdkV2Resource(name) && !slices.Contains(fallbackOption.resourceFallbacks, name) {
			resources = append(resources, resourceFunc)
		}
	}

	return append(resources, onboardedResources...)
}

// getPluginFrameworkDataSourcesToRegister is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkDataSourcesToRegister(sdkV2Fallbacks ...SdkV2FallbackOption) []func() datasource.DataSource {
	fallbackOption := sdkV2FallbackOptions{}
	for _, o := range sdkV2Fallbacks {
		o.Apply(&fallbackOption)
	}

	var dataSources []func() datasource.DataSource

	// Loop through the map and add data sources if they're not specifically marked to use the SDK V2
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if !shouldUseSdkV2DataSource(name) && !slices.Contains(fallbackOption.dataSourceFallbacks, name) {
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
func GetSdkV2ResourcesToRemove(sdkV2Fallbacks ...SdkV2FallbackOption) []string {
	fallbackOption := sdkV2FallbackOptions{}
	for _, o := range sdkV2Fallbacks {
		o.Apply(&fallbackOption)
	}

	resourcesToRemove := []string{}
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if !shouldUseSdkV2Resource(name) && !slices.Contains(fallbackOption.resourceFallbacks, name) {
			resourcesToRemove = append(resourcesToRemove, name)
		}
	}
	return resourcesToRemove
}

// GetSdkV2DataSourcesToRemove is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func GetSdkV2DataSourcesToRemove(sdkV2Fallbacks ...SdkV2FallbackOption) []string {
	fallbackOption := sdkV2FallbackOptions{}
	for _, o := range sdkV2Fallbacks {
		o.Apply(&fallbackOption)
	}

	dataSourcesToRemove := []string{}
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if !shouldUseSdkV2DataSource(name) && !slices.Contains(fallbackOption.dataSourceFallbacks, name) {
			dataSourcesToRemove = append(dataSourcesToRemove, name)
		}
	}
	return dataSourcesToRemove
}
