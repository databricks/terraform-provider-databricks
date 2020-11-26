package common

var (
	version = "0.3.0"
	// ResourceName ...
	ResourceName contextKey = 1
	// TerraformVersion ...
	TerraformVersion contextKey = 2
)

type contextKey int

// Version returns version of provider
func Version() string {
	return version
}
