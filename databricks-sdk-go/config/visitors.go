package config

import (
	"net/http"

	"github.com/databricks/databricks-sdk-go/useragent"
)

type AuthType struct{}

func authInUserAgentVisitor(c *Config) func(*http.Request) error {
	return func(r *http.Request) error {
		authType := c.AuthType
		if t := r.Context().Value(AuthType{}); t != nil {
			authType = t.(string)
		}
		ctx := useragent.InContext(r.Context(), useragent.AuthKey, authType)
		*r = *r.WithContext(ctx) // replace request
		return nil
	}
}
