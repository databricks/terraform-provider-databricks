package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

type tokenError struct {
	message string
	err     *httpclient.HttpError
}

func (e *tokenError) Error() string {
	return e.message
}

func (e *tokenError) Unwrap() error {
	sdkErr, ok := apierr.ByStatusCode(e.err.StatusCode)
	if ok {
		// this is how we distinguish between bad requests and permission denies:
		// if clients sets a wrong AAD tenant ID, client code would expect to handle
		// apierr.ErrBadRequest - no matter if it's our API or AAD's API, same for
		// apierr.ErrUnauthorized.
		return sdkErr
	}
	return e.err
}

func (c *Config) mapAzureError(defaultErr *httpclient.HttpError) error {
	env := c.Environment()
	switch defaultErr.Request.Host {
	case c.hostOrEmpty(env.AzureActiveDirectoryEndpoint()):
		return c.mapAzureActiveDirectoryError(defaultErr)
	case c.hostOrEmpty(env.AzureResourceManagerEndpoint()):
		return c.mapAzureResourceManagerError(defaultErr)
	default:
		// Azure MSI endpoint returns not so typed error bodies: `404 page not found`
		return &tokenError{
			message: defaultErr.Message,
			err:     defaultErr,
		}
	}
}

func (c *Config) hostOrEmpty(endpoint string) string {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return ""
	}
	return parsedURL.Host
}

type azureActiveDirectoryErrorResponse struct {
	CorrelationID    string `json:"correlation_id,omitempty"`
	ErrorType        string `json:"error"`
	ErrorCodes       []int  `json:"error_codes"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}

func (c *Config) mapAzureActiveDirectoryError(defaultErr *httpclient.HttpError) error {
	var aadError azureActiveDirectoryErrorResponse
	err := json.Unmarshal([]byte(defaultErr.Message), &aadError)
	if err != nil {
		return defaultErr
	}
	// remove Trace ID / Correlation ID from error description, as we're adding
	// aadError.ErrorURI to the message - The most useful portion of info to
	// resolve the issue. If user is looking at debug logs -  Trace ID and
	// Correlation ID will appear there anyway (in pretty-printed JSON).
	msg, _, ok := strings.Cut(aadError.ErrorDescription, ". Trace ID")
	if ok {
		aadError.ErrorDescription = msg
	}
	if aadError.ErrorURI != "" {
		msg = fmt.Sprintf("%s. See %s", strings.TrimSuffix(msg, "."), aadError.ErrorURI)
	}
	return &tokenError{
		message: msg,
		err:     defaultErr,
	}
}

type azureResourceManagerErrorError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type azureResourceManagerErrorResponse struct {
	Error azureResourceManagerErrorError `json:"error"`
}

func (c *Config) mapAzureResourceManagerError(defaultErr *httpclient.HttpError) error {
	var rmError azureResourceManagerErrorResponse
	err := json.Unmarshal([]byte(defaultErr.Message), &rmError)
	if err != nil {
		return defaultErr
	}
	return &tokenError{
		message: strings.TrimSuffix(rmError.Error.Message, "."),
		err:     defaultErr,
	}
}

type azureHostResolver interface {
	tokenSourceFor(ctx context.Context, cfg *Config, aadEndpoint, resource string) auth.TokenSource
}

func (c *Config) azureEnsureWorkspaceUrl(ctx context.Context, ahr azureHostResolver) error {
	if c.AzureResourceID == "" || c.Host != "" {
		return nil
	}
	azureEnv := c.Environment().AzureEnvironment
	// azure resource ID can also be used in lieu of host by some of the clients, like Terraform
	management := ahr.tokenSourceFor(ctx, c, azureEnv.ActiveDirectoryEndpoint, azureEnv.ResourceManagerEndpoint)
	var workspaceMetadata struct {
		Properties struct {
			WorkspaceURL string `json:"workspaceUrl"`
		} `json:"properties"`
	}
	requestURL := strings.TrimSuffix(azureEnv.ResourceManagerEndpoint, "/") + c.AzureResourceID + "?api-version=2018-04-01"
	err := c.refreshClient.Do(ctx, "GET", requestURL,
		httpclient.WithResponseUnmarshal(&workspaceMetadata),
		httpclient.WithAuthTokenSource(management),
	)
	if err != nil {
		return fmt.Errorf("resolve workspace: %w", err)
	}
	c.Host = fmt.Sprintf("https://%s", workspaceMetadata.Properties.WorkspaceURL)
	logger.Debugf(ctx, "Discovered workspace url: %s", c.Host)
	return nil
}

// loadAzureTenantId fetches the Azure tenant ID from the Azure AD endpoint.
// The tenant ID isn't directly exposed by any API today, but it can be inferred
// from the URL that a user is redirected to after accessing /aad/auth (the
// Azure Databricks login endpoint). Here, the redirect is not followed, but the
// tenant ID is extracted from the URL.
func (c *Config) loadAzureTenantId(ctx context.Context) error {
	if !c.IsAzure() || c.AzureTenantID != "" || c.Host == "" {
		return nil
	}
	req, err := http.NewRequestWithContext(ctx, "GET", c.CanonicalHostName()+"/aad/auth", nil)
	if err != nil {
		return err
	}
	res, err := c.azureTenantIdFetchClient.Do(req)
	if err != nil && !errors.Is(err, http.ErrUseLastResponse) {
		return err
	}
	location := res.Header.Get("Location")
	parsedUrl, err := url.ParseRequestURI(location)
	if err != nil {
		return err
	}
	// Response URL is of the form https://login.microsoftonline.com/<tenantID>/oauth2/authorize?...
	// or corresponding in other Azure clouds
	splitPath := strings.SplitN(parsedUrl.Path, "/", 3)
	c.AzureTenantID = splitPath[1]
	return nil
}
