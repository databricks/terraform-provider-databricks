package context

import (
	"context"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/databricks/terraform-provider-databricks/internal/providers/common"
)

func SetUserAgentInEphemeralResourceContext(ctx context.Context, ephemeralResourceName string) context.Context {
	ctx = common.SetSDKInContext(ctx, sdkName)
	return useragent.InContext(ctx, "ephemeral", ephemeralResourceName)
}
