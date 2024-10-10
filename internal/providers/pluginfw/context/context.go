package common

import (
	"context"

	"github.com/databricks/databricks-sdk-go/useragent"
)

func SetResourceNameInContext(ctx context.Context, resourceName string) context.Context {
	return useragent.InContext(ctx, "resource", resourceName)
}

func SetDataSourceNameInContext(ctx context.Context, dataSourceName string) context.Context {
	return useragent.InContext(ctx, "data", dataSourceName)
}
