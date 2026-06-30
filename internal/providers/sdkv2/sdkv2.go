// Package sdkv2 contains the changes specific to the SDKv2
//
// Note: This package shouldn't depend on internal/providers/pluginfw or internal/providers
package sdkv2

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"unicode"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/useragent"

	"github.com/databricks/terraform-provider-databricks/common"
	"github.com/databricks/terraform-provider-databricks/internal/providers/client"
	providercommon "github.com/databricks/terraform-provider-databricks/internal/providers/common"
	"github.com/databricks/terraform-provider-databricks/logger"
	"github.com/databricks/terraform-provider-databricks/settings"
)

func init() {
	// IMPORTANT: this line cannot be changed, because it's used for
	// internal purposes at Databricks.
	useragent.WithProduct(providercommon.ProviderName, common.Version())

	userAgentExtraEnv := os.Getenv("DATABRICKS_USER_AGENT_EXTRA")
	out, err := ParseUserAgentExtra(userAgentExtraEnv)
	if err != nil {
		panic(fmt.Errorf("failed to parse DATABRICKS_USER_AGENT_EXTRA: %s", err))
	}

	for _, extra := range out {
		useragent.WithUserAgentExtra(extra.Key, extra.Value)
	}
}

var terraformVersionOnce sync.Once

type sdkV2ProviderOptions struct {
	configCustomizer func(*config.Config) error
}

// SdkV2ProviderOption is a functional option for configuring the SDK V2 provider.
type SdkV2ProviderOption func(*sdkV2ProviderOptions)

// WithConfigCustomizer allows the caller to customize the SDK config before config resolution,
// so customizer-set fields (e.g. Host) participate in resolveHostMetadata and auth.
func WithConfigCustomizer(customizer func(*config.Config) error) SdkV2ProviderOption {
	return func(o *sdkV2ProviderOptions) {
		o.configCustomizer = customizer
	}
}

// DatabricksProvider returns the entire terraform provider object
func DatabricksProvider(opts ...SdkV2ProviderOption) *schema.Provider {
	providerOptions := &sdkV2ProviderOptions{}
	for _, optFunc := range opts {
		optFunc(providerOptions)
	}

	dataSourceMap := DataSources()
	resourceMap := Resources()

	p := &schema.Provider{
		DataSourcesMap: dataSourceMap,
		ResourcesMap:   resourceMap,
		Schema:         providerSchema(),
	}
	for name, resource := range settings.AllSettingsResources() {
		p.ResourcesMap[fmt.Sprintf("databricks_%s_setting", name)] = resource.ToResource()
	}
	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		if p.TerraformVersion != "" {
			terraformVersionOnce.Do(func() {
				useragent.WithUserAgentExtra("terraform", p.TerraformVersion)
			})
		}
		logger.SetTfLogger(logger.NewTfLogger(ctx))
		return ConfigureDatabricksClient(ctx, d, providerOptions.configCustomizer)
	}
	common.AddContextToAllResources(p, "databricks")
	return p
}

func providerSchema() map[string]*schema.Schema {
	ps := map[string]*schema.Schema{}
	for _, attr := range config.ConfigAttributes {
		switch attr.Kind {
		case reflect.String:
			ps[attr.Name] = &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		case reflect.Bool:
			ps[attr.Name] = &schema.Schema{
				Type:      schema.TypeBool,
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		case reflect.Int:
			ps[attr.Name] = &schema.Schema{
				Type:      schema.TypeInt,
				Optional:  true,
				Sensitive: attr.Sensitive,
			}
		case reflect.Slice:
			ps[attr.Name] = &schema.Schema{
				Type:      schema.TypeList,
				Optional:  true,
				Sensitive: attr.Sensitive,
				Elem:      &schema.Schema{Type: schema.TypeString},
			}
		}
	}
	return ps
}

func ConfigureDatabricksClient(ctx context.Context, d *schema.ResourceData, configCustomizer func(*config.Config) error) (any, diag.Diagnostics) {
	cfg := &config.Config{}
	attrsUsed := []string{}
	for _, attr := range config.ConfigAttributes {
		if value, ok := d.GetOk(attr.Name); ok {
			// SDKv2's GetOk returns []interface{} for lists, but the SDK expects []string.
			if attr.Kind == reflect.Slice {
				rawList, ok := value.([]interface{})
				if !ok {
					return nil, diag.Errorf("unexpected type for attribute %s: expected []interface{}, got %T", attr.Name, value)
				}
				strList := make([]string, len(rawList))
				for i, v := range rawList {
					strList[i] = v.(string)
				}
				value = strList
			}
			err := attr.Set(cfg, value)
			if err != nil {
				return nil, diag.FromErr(err)
			}
			attrsUsed = append(attrsUsed, attr.Name)
		}
	}
	if len(attrsUsed) > 0 {
		sort.Strings(attrsUsed)
		tflog.Info(ctx, fmt.Sprintf("(sdkv2) Attributes specified in provider configuration: %s", strings.Join(attrsUsed, ", ")))
	} else {
		tflog.Info(ctx, "(sdkv2) No attributes specified in provider configuration")
	}
	databricksClient, err := client.PrepareDatabricksClient(ctx, cfg, configCustomizer)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	return databricksClient, nil
}

type UserAgentExtra struct {
	Key   string
	Value string
}

// Regex for product strings. See RFC 9110.
//
// product = token ["/" product-version]
// product-version = token
// token = 1*tchar
// tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." / "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
var productRegexRfc9110 = regexp.MustCompile("^([!#$%&'*+\\-.^_`|~0-9A-Za-z]+)(/([!#$%&'*+\\-.^_`|~0-9A-Za-z]+))?$")

func ParseUserAgentExtra(env string) ([]UserAgentExtra, error) {
	out := []UserAgentExtra{}

	products := strings.FieldsFunc(env, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	for _, product := range products {
		match := productRegexRfc9110.FindStringSubmatch(product)

		if len(match) != 4 {
			return nil, fmt.Errorf("product string must follow RFC 9110: %s", product)
		}

		if match[3] == "" {
			return nil, fmt.Errorf("product string must include version: %s", product)
		}

		out = append(out, UserAgentExtra{
			Key:   match[1],
			Value: match[3],
		})
	}

	return out, nil
}
