package apierr

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/common/environment"
)

// Metadata about the private link product. Private link redirects users to the
// login page with a query parameter that indicates the error. This struct
// contains information about the private link service, the endpoint name, and a
// reference page for more information.
//
// Eventually, the REST API should return an error directly when a request is
// made from a network that does not have access to the workspace. Once that
// happens, this struct can be removed.
type privateLinkInfo struct {
	// The name of the private link service (e.g. AWS PrivateLink, Azure Private
	// Link, etc.)
	serviceName string

	// The name of the private link endpoint (e.g. AWS VPC endpoint, Azure Private
	// Link endpoint, etc.)
	endpointName string

	// A reference page for more information about the private link service.
	referencePage string
}

func (p privateLinkInfo) errorMessage() string {
	privateLinkValidationError := fmt.Sprintf(
		`The requested workspace has %[1]s enabled and is not accessible from
the current network. Ensure that %[1]s is properly configured and that your
device has access to the %s. For more information, see %s.`,
		p.serviceName, p.endpointName, p.referencePage)
	return strings.ReplaceAll(privateLinkValidationError, "\n", " ")
}

// Map of private link information by cloud.
var privateLinkInfoMap = map[environment.Cloud]privateLinkInfo{
	environment.CloudAWS: {
		serviceName:   "AWS PrivateLink",
		endpointName:  "AWS VPC endpoint",
		referencePage: "https://docs.databricks.com/en/security/network/classic/privatelink.html",
	},
	environment.CloudAzure: {
		serviceName:   "Azure Private Link",
		endpointName:  "Azure Private Link endpoint",
		referencePage: "https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting",
	},
	environment.CloudGCP: {
		serviceName:   "Private Service Connect",
		endpointName:  "GCP VPC endpoint",
		referencePage: "https://docs.gcp.databricks.com/en/security/network/classic/private-service-connect.html",
	},
}

func isPrivateLinkRedirect(url *url.URL) bool {
	return strings.Contains(url.RawQuery, "error=private-link-validation-error") && url.EscapedPath() == "/login.html"
}

func privateLinkValidationError(url *url.URL) *APIError {
	env := environment.GetEnvironmentForHostname(url.Host)
	info := privateLinkInfoMap[env.Cloud]
	return &APIError{
		ErrorCode:  "PRIVATE_LINK_VALIDATION_ERROR",
		StatusCode: 403,
		Message:    info.errorMessage(),
		unwrap:     ErrPermissionDenied,
	}
}
