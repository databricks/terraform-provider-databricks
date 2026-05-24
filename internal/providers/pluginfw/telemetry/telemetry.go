package telemetry

import (
	"context"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/internal/providers/common"
)

const sdkName = "pluginframework"

func WithResource(ctx context.Context, resourceName string) context.Context {
	ctx = common.SetSDKInContext(ctx, sdkName)
	return useragent.InContext(ctx, "resource", resourceName)
}

func WithDataSource(ctx context.Context, dataSourceName string) context.Context {
	ctx = common.SetSDKInContext(ctx, sdkName)
	return useragent.InContext(ctx, "data", dataSourceName)
}
