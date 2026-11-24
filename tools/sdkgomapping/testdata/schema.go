package testdata

// testVolume is a schema struct with fields that the generator can't map to the
// VolumeInfo of sdk-go uc/volumes/v1.
type testVolume struct {
	Name string `json:"name"`
	// NotInSdkGo has no VolumeInfo field.
	NotInSdkGo string `json:"not_in_sdk_go,omitempty"`
	// BrowseOnly is a bool, which the generator doesn't support.
	BrowseOnly bool `json:"browse_only,omitempty"`
	// CreatedAt maps to an *int64, which the generator doesn't support.
	CreatedAt string `json:"created_at,omitempty"`
	Skipped   string `json:"-"`
}
