package common

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/databricks/databricks-sdk-go/useragent"
)

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

// ParseUserAgentExtra parses a whitespace-separated list of RFC 9110 product
// strings (e.g. "databricks-sra/1.2.0 my-project/0.0.1") as accepted by both
// the DATABRICKS_USER_AGENT_EXTRA environment variable and the
// user_agent_extra provider attribute.
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

// ApplyUserAgentExtra parses a whitespace-separated list of RFC 9110 product
// strings and registers each of them as a user agent extra. Registration is
// process-global and deduplicated by the SDK, so applying the same product
// string from multiple sources (environment variable, provider attribute,
// muxed provider instances) is safe.
func ApplyUserAgentExtra(products string) error {
	extras, err := ParseUserAgentExtra(products)
	if err != nil {
		return err
	}
	for _, extra := range extras {
		useragent.WithUserAgentExtra(extra.Key, extra.Value)
	}
	return nil
}
