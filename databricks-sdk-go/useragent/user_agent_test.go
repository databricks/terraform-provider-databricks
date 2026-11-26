package useragent

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/version"
	"github.com/stretchr/testify/assert"
)

func TestGoVersion(t *testing.T) {
	// The function `goVersion` returns a semver canonicalized Go version,
	// but `runtime.Version` may or may not include a patch version.
	// Correct for that in the test below.
	parts := strings.SplitN(strings.TrimPrefix(runtime.Version(), "go"), ".", 3)
	if len(parts) == 2 {
		parts = append(parts, "0")
	}
	assert.Equal(t, goVersion(), strings.Join(parts, "."))
}

func TestFromContext_Default(t *testing.T) {
	userAgent := FromContext(context.Background())
	// tests may be developed and run on different versions of different things
	expectedFormat := "unknown/0.0.0 databricks-sdk-go/%s go/%s os/%s"
	expected := fmt.Sprintf(expectedFormat, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)
}

func TestFromContext_Custom(t *testing.T) {
	WithProduct("unit-tests", "0.0.1")
	WithUserAgentExtra("pulumi", "3.8.4")
	WithUserAgentExtra("terraform", "1.3.4")
	WithUserAgentExtra("terraform", "1.3.5")
	WithUserAgentExtra("terraform", "1.3.6")

	ctx := InContext(context.Background(), "a", "b")
	ctx = InContext(ctx, "a", "d")
	ctx = InContext(ctx, "a", "e")

	ctx2 := InContext(ctx, "a", "f")
	ctx2 = InContext(ctx2, "foo", "bar")

	userAgent := FromContext(ctx)
	userAgent2 := FromContext(ctx2)

	// tests may be developed and run on different versions of different things
	expectedFormat := "unit-tests/0.0.1 databricks-sdk-go/%s go/%s os/%s pulumi/3.8.4 terraform/1.3.4 terraform/1.3.5 terraform/1.3.6 a/b a/d a/e"
	expected := fmt.Sprintf(expectedFormat, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)

	expectedFormat2 := "unit-tests/0.0.1 databricks-sdk-go/%s go/%s os/%s pulumi/3.8.4 terraform/1.3.4 terraform/1.3.5 terraform/1.3.6 a/b a/d a/e a/f foo/bar"
	expected2 := fmt.Sprintf(expectedFormat2, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected2, userAgent2)

	// upstream and upstream-version only appear when both are set
	assert.NotContains(t, userAgent, "upstream/")
	os.Setenv("DATABRICKS_SDK_UPSTREAM", "my-upstream-sdk")
	defer os.Unsetenv("DATABRICKS_SDK_UPSTREAM")
	userAgent3 := FromContext(ctx)
	assert.NotContains(t, userAgent3, "upstream/")
	assert.NotContains(t, userAgent3, "upstream-version/")

	os.Setenv("DATABRICKS_SDK_UPSTREAM_VERSION", "1.2.3")
	defer os.Unsetenv("DATABRICKS_SDK_UPSTREAM_VERSION")
	userAgent4 := FromContext(context.Background())
	assert.Contains(t, userAgent4, "upstream/my-upstream-sdk upstream-version/1.2.3")
}

func TestDefaultsAreValid(t *testing.T) {
	WithProduct(productName, productVersion)
}

func TestMultiplePartners(t *testing.T) {
	WithPartner("partner1")
	WithPartner("partner2")
	userAgent := FromContext(context.Background())
	assert.Contains(t, userAgent, "partner/partner1")
	assert.Contains(t, userAgent, "partner/partner2")
}

func TestInContext_DeduplicatesIdenticalPairs(t *testing.T) {
	// Repeatedly injecting the same key/value pair onto a reused context must
	// not grow the user agent: this is what happened when a long-lived caller
	// threaded one context across many operations.
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		ctx = InContext(ctx, "resource", "cluster")
		ctx = InContext(ctx, "sdk", "sdkv2")
	}
	userAgent := FromContext(ctx)
	assert.Equal(t, 1, strings.Count(userAgent, "resource/cluster"))
	assert.Equal(t, 1, strings.Count(userAgent, "sdk/sdkv2"))

	// Distinct values for the same key are still preserved.
	ctx2 := InContext(context.Background(), "resource", "a")
	ctx2 = InContext(ctx2, "resource", "b")
	userAgent2 := FromContext(ctx2)
	assert.Contains(t, userAgent2, "resource/a")
	assert.Contains(t, userAgent2, "resource/b")
}
