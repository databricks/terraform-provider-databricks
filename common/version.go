package common

import "context"

var (
	version = "1.6.5"
	// ResourceName is resource name without databricks_ prefix
	ResourceName contextKey = 1
	// Provider is the current instance of provider
	Provider contextKey = 2
	// Current is the current name of integration test
	Current contextKey = 3
	// If current resource is data
	IsData contextKey = 4
	// apiVersion
	Api contextKey = 5
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
