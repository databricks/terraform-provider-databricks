// Package common contains the changes used by both internal/providers/sdkv2 and internal/providers/pluginfw packages.
//
// Note: This is different from internal/providers which contains the changes that *depends* on both:
// internal/providers/sdkv2 and internal/providers/pluginfw packages. Whereas, internal/providers/common package contains
// the changes *used* by both internal/providers/sdkv2 and internal/providers/pluginfw packages.
package common

import (
	"context"

	"github.com/databricks/databricks-sdk-go/useragent"
)

const ProviderName = "databricks-tf-provider"

func SetSDKInContext(ctx context.Context, sdkUsed string) context.Context {
	return useragent.InContext(ctx, "sdk", sdkUsed)
}
