package pluginfw

// This file contains all of the utils for controlling the plugin framework rollout.
// For migrated resources and data sources, we can add them to the two maps below to have them registered with the plugin framework.
// Users can manually specify resources and data sources to use SDK V2 instead of the plugin framework by setting the USE_SDK_V2_RESOURCES and USE_SDK_V2_DATA_SOURCES environment variables.
//
// Example: USE_SDK_V2_RESOURCES="databricks_library" would force the library resource to use SDK V2 instead of the plugin framework.

import (
	"context"
	"log"
	"os"
	"slices"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/app"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/catalog"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/cluster"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/dashboards"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/library"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/notificationdestinations"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/qualitymonitor"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/registered_model"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/serving"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/sharing"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/user"
	"github.com/databricks/terraform-provider-databricks/internal/providers/pluginfw/products/volume"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// migratedResources lists resources that were originally implemented on SDKv2
// and have been re-implemented on the Plugin Framework. Entries here are served
// by the Plugin Framework by default; users can opt back to SDKv2 by listing
// the resource name in the USE_SDK_V2_RESOURCES environment variable.
//
// Deprecation contract for entries in this list:
//  1. The SDKv2 implementation must stay present and functional for the rest of
//     this major version so that USE_SDK_V2_RESOURCES still works.
//  2. The resource's documentation must carry a `~> **Deprecation**` banner
//     warning that the SDKv2 implementation will be removed in the next major.
//  3. The PR that adds an entry must announce the deprecation in NEXT_CHANGELOG.
//  4. In the next major release, the SDKv2 source file and test must be deleted
//     and the entry moved from migratedResources to pluginFwOnlyResources.
//
// getPluginFrameworkResourcesToRegister emits a one-time runtime warning whenever
// a user steers a name in this list back to SDKv2, so entries inherit the
// deprecation signal automatically.
//
// Keep this list sorted.
var migratedResources = []func() resource.Resource{
	library.ResourceLibrary,
	qualitymonitor.ResourceQualityMonitor,
	sharing.ResourceShare,
}

// migratedDataSources lists data sources that were originally implemented on
// SDKv2 and have been re-implemented on the Plugin Framework. The same
// deprecation contract as migratedResources applies; see the comment above
// migratedResources for details.
//
// Keep this list sorted.
var migratedDataSources = []func() datasource.DataSource{
	sharing.DataSourceShare,
	sharing.DataSourceShares,
	volume.DataSourceVolumes,
}

// warnedFallbackNames tracks resource/data-source names for which we have
// already emitted the SDKv2-fallback deprecation warning, so we warn at most
// once per name per provider process.
var warnedFallbackNames sync.Map

// emitSdkV2FallbackWarning logs a one-time deprecation warning when a name in
// migratedResources / migratedDataSources is steered back to the SDKv2
// implementation, either via the USE_SDK_V2_RESOURCES / USE_SDK_V2_DATA_SOURCES
// environment variables or via the test-only WithSdkV2ResourceFallbacks /
// WithSdkV2DataSourceFallbacks options. kind is "resource" or "data source".
//
// The warning is sent through log.Printf with a [WARN] prefix; the Terraform
// CLI surfaces this to users running with TF_LOG=WARN or higher. This is the
// same channel SDKv2 schema deprecations use.
func emitSdkV2FallbackWarning(name, kind string) {
	if _, loaded := warnedFallbackNames.LoadOrStore(name, struct{}{}); loaded {
		return
	}
	log.Printf("[WARN] %s %q is being served by the deprecated SDKv2 implementation "+
		"(selected via USE_SDK_V2_RESOURCES / USE_SDK_V2_DATA_SOURCES or a test-only "+
		"fallback option). The SDKv2 implementation will be removed in the next major "+
		"release of the provider; remove the override to use the Plugin Framework version.",
		kind, name)
}

// List of resources that have been onboarded to the plugin framework - not migrated from sdkv2.
// Keep this list sorted.
var pluginFwOnlyResources = append(
	[]func() resource.Resource{
		app.ResourceApp,
	},
	autoGeneratedResources...,
)

// List of data sources that have been onboarded to the plugin framework - not migrated from sdkv2.
// Keep this list sorted.
var pluginFwOnlyDataSources = append(
	[]func() datasource.DataSource{
		app.DataSourceApp,
		app.DataSourceApps,
		catalog.DataSourceFunctions,
		dashboards.DataSourceDashboards,
		notificationdestinations.DataSourceNotificationDestinations,
		registered_model.DataSourceRegisteredModel,
		registered_model.DataSourceRegisteredModelVersions,
		serving.DataSourceServingEndpoints,
		user.DataSourceUsers,
		// TODO: Add DataSourceCluster into migratedDataSources after fixing unit tests.
		cluster.DataSourceCluster, // Using the staging name (with pluginframework suffix)
	},
	autoGeneratedDataSources...,
)

type pluginFrameworkOptions struct {
	resourceFallbacks   []string
	dataSourceFallbacks []string
	configCustomizer    func(*config.Config) error
}

// PluginFrameworkOption is an interface for acceptance tests to specify resources / data sources to fallback to SDK V2
type PluginFrameworkOption interface {
	Apply(*pluginFrameworkOptions)
}

type sdkV2ResourceFallback struct {
	resourceFallbacks []string
}

func (o *sdkV2ResourceFallback) Apply(options *pluginFrameworkOptions) {
	options.resourceFallbacks = o.resourceFallbacks
}

// WithSdkV2ResourceFallbacks is a helper function to specify resources to fallback to SDK V2
func WithSdkV2ResourceFallbacks(fallbacks []string) PluginFrameworkOption {
	return &sdkV2ResourceFallback{resourceFallbacks: fallbacks}
}

type sdkv2DataSourceFallback struct {
	dataSourceFallbacks []string
}

func (o *sdkv2DataSourceFallback) Apply(options *pluginFrameworkOptions) {
	options.dataSourceFallbacks = o.dataSourceFallbacks
}

// WithSdkV2DataSourceFallbacks is a helper function to specify data sources to fallback to SDK V2
func WithSdkV2DataSourceFallbacks(fallbacks []string) PluginFrameworkOption {
	return &sdkv2DataSourceFallback{dataSourceFallbacks: fallbacks}
}

type configCustomizer struct {
	configCustomizer func(*config.Config) error
}

func (o *configCustomizer) Apply(options *pluginFrameworkOptions) {
	options.configCustomizer = o.configCustomizer
}

// WithConfigCustomizer allows the caller to customize the SDK config before config resolution,
// so customizer-set fields (e.g. Host) participate in resolveHostMetadata and auth.
func WithConfigCustomizer(customizer func(*config.Config) error) PluginFrameworkOption {
	return &configCustomizer{configCustomizer: customizer}
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
	return slices.Contains(useSdkV2Resources, resourceName)
}

// Helper function to check if a data source should use be in SDK V2 instead of plugin framework
func shouldUseSdkV2DataSource(dataSourceName string) bool {
	sdkV2DataSources := getUseSdkV2DataSources()
	return slices.Contains(sdkV2DataSources, dataSourceName)
}

// getPluginFrameworkResourcesToRegister is a helper function to get the list of resources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkResourcesToRegister(resourceFallbacks []string) []func() resource.Resource {
	var resources []func() resource.Resource

	// Loop through the map and add resources if they're not specifically marked to use the SDK V2
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if shouldUseSdkV2Resource(name) || slices.Contains(resourceFallbacks, name) {
			emitSdkV2FallbackWarning(name, "resource")
			continue
		}
		resources = append(resources, resourceFunc)
	}

	return append(resources, pluginFwOnlyResources...)
}

// getPluginFrameworkDataSourcesToRegister is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func getPluginFrameworkDataSourcesToRegister(dataSourceFallbacks []string) []func() datasource.DataSource {
	var dataSources []func() datasource.DataSource

	// Loop through the map and add data sources if they're not specifically marked to use the SDK V2
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if shouldUseSdkV2DataSource(name) || slices.Contains(dataSourceFallbacks, name) {
			emitSdkV2FallbackWarning(name, "data source")
			continue
		}
		dataSources = append(dataSources, dataSourceFunc)
	}

	return append(dataSources, pluginFwOnlyDataSources...)
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
func GetSdkV2ResourcesToRemove(resourceFallbacks []string) []string {
	resourcesToRemove := []string{}
	for _, resourceFunc := range migratedResources {
		name := getResourceName(resourceFunc)
		if !shouldUseSdkV2Resource(name) && !slices.Contains(resourceFallbacks, name) {
			resourcesToRemove = append(resourcesToRemove, name)
		}
	}
	return resourcesToRemove
}

// GetSdkV2DataSourcesToRemove is a helper function to get the list of data sources that are migrated away from sdkv2 to plugin framework
func GetSdkV2DataSourcesToRemove(dataSourceFallbacks []string) []string {
	dataSourcesToRemove := []string{}
	for _, dataSourceFunc := range migratedDataSources {
		name := getDataSourceName(dataSourceFunc)
		if !shouldUseSdkV2DataSource(name) && !slices.Contains(dataSourceFallbacks, name) {
			dataSourcesToRemove = append(dataSourcesToRemove, name)
		}
	}
	return dataSourcesToRemove
}
