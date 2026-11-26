package useragent

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/databricks/databricks-sdk-go/version"
	"golang.org/x/mod/semver"
)

const (
	RuntimeKey     = "runtime"
	CicdKey        = "cicd"
	AuthKey        = "auth"
	AgentKey       = "agent"
	MetaHarnessKey = "meta-harness"
)

// WithProduct sets the product name and product version globally.
// It should be called by developers to differentiate their application from others.
func WithProduct(name, version string) {
	if err := matchAlphanum(name); err != nil {
		panic(err)
	}
	if err := matchSemVer(version); err != nil {
		panic(err)
	}
	productName = name
	productVersion = version
}

var productName = "unknown"
var productVersion = "0.0.0"

const sdkName = "databricks-sdk-go"
const sdkVersion = version.Version

// extra holds per-process static extra information
// bits, that are agreed upfront with Databricks.
var extra data

// private type for user agent info in the context
type key int

// context key for holding volatile part of user agent data
const ctxAgent key = 5

// WithUserAgentExtra sets per-process extra user agent data,
// which integration developers have agreed with Databricks.
func WithUserAgentExtra(key, value string) {
	extra = extra.With(key, value)
}

func WithPartner(partner string) {
	WithUserAgentExtra("partner", partner)
}

func getUpstreamUserAgentInfo() []info {
	product := os.Getenv("DATABRICKS_SDK_UPSTREAM")
	version := os.Getenv("DATABRICKS_SDK_UPSTREAM_VERSION")
	if product == "" || version == "" {
		return nil
	}
	return []info{
		{"upstream", product},
		{"upstream-version", version},
	}
}

// InContext populates context with user agent dimension,
// usually to differentiate subsets of functionality and
// agreed with Databricks.
func InContext(ctx context.Context, key, value string) context.Context {
	uac, _ := ctx.Value(ctxAgent).(data)
	uac = uac.With(key, value)
	return context.WithValue(ctx, ctxAgent, uac)
}

// FromContext gets compliant user-agent string in a given context
func FromContext(ctx context.Context) string {
	base := data{
		{productName, productVersion},
		{sdkName, sdkVersion},
		{"go", goVersion()},
		{"os", runtime.GOOS},
	}
	base = append(base, extra...)
	base = append(base, getUpstreamUserAgentInfo()...)
	uac, _ := ctx.Value(ctxAgent).(data)
	return append(base, uac...).String()
}

func goVersion() string {
	gv := runtime.Version()
	ssv := strings.ReplaceAll(gv, "go", "v")
	sv := semver.Canonical(ssv)
	return strings.TrimPrefix(sv, "v")
}

type info struct {
	Key   string
	Value string
}

func (u info) String() string {
	return fmt.Sprintf("%s/%s", u.Key, u.Value)
}

type data []info

// With appends the key/value pair, skipping it when an identical key/value
// pair is already present. De-duplicating exact pairs keeps the user agent from
// growing without bound when the same dimension is injected repeatedly onto a
// reused context. Distinct values for the same key are still preserved.
// Panics if key or value don't satisfy alphanumeric or semver format.
func (d data) With(key, value string) data {
	if err := matchAlphanum(key); err != nil {
		panic(err)
	}
	if err := matchAlphanumOrSemVer(value); err != nil {
		panic(err)
	}
	for _, e := range d {
		if e.Key == key && e.Value == value {
			return d
		}
	}
	return append(d, info{key, value})
}

func (d data) String() string {
	pairs := []string{}
	for _, v := range d {
		pairs = append(pairs, v.String())
	}
	return strings.Join(pairs, " ")
}
