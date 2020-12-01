package common

import "context"

var (
	version = "0.3.0"
	// ResourceName ...
	ResourceName contextKey = 1
	// TerraformVersion ...
	TerraformVersion contextKey = 2
)

type contextKey int

func (k contextKey) GetOrUnknown(ctx context.Context) string {
	rn, ok := ctx.Value(k).(string)
	if !ok {
		return "unknown"
	}
	return rn
}

// Version returns version of provider
func Version() string {
	return version
}
