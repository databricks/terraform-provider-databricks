// Package internal contains the changes used by both internal/providers/sdkv2 and internal/providers/pluginfw packages.
//
// Note: This is different from internal/providers which contains the changes that *depends* on both:
// internal/providers/sdkv2 and internal/providers/pluginfw packages. Whereas, internal/providers/internal package contains
// the changes *used* by both internal/providers/sdkv2 and internal/providers/pluginfw packages.
package internal

const ProviderName = "databricks-tf-provider"
