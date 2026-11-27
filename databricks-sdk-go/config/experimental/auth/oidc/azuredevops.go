package oidc

import (
	"context"
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

var (
	errMissingAccessToken = fmt.Errorf("SYSTEM_ACCESSTOKEN env var not found, " +
		"if calling from Azure DevOps Pipeline, please set this env var following " +
		"https://learn.microsoft.com/en-us/azure/devops/pipelines/build/variables?view=azure-devops&tabs=yaml#systemaccesstoken")

	errNotCallingFromAzureDevOps = fmt.Errorf("not calling from Azure DevOps Pipeline")
)

// NewAzureDevOpsIDTokenSource returns a new IDTokenSource that retrieves an
// IDToken from an Azure DevOps environment.
//
// This IDTokenSource is only valid when running in Azure DevOps Pipelines.
func NewAzureDevOpsIDTokenSource(client *httpclient.ApiClient) (IDTokenSource, error) {
	ts := &azureDevOpsIDTokenSource{Client: client}

	// Environment variable SYSTEM_ACCESSTOKEN is treated differently from
	// other environment variables because it is controlled by users within
	// their Azure DevOps Pipelines.
	if v := os.Getenv("SYSTEM_ACCESSTOKEN"); v != "" {
		ts.AccessToken = v
	} else {
		return nil, errMissingAccessToken
	}

	// Environment variables that should always be set when calling from an
	// Azure DevOps Pipeline.
	if err := setFromAzureDevOpsEnv(&ts.TeamFoundationCollectionURI, "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.PlanID, "SYSTEM_PLANID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.JobID, "SYSTEM_JOBID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.TeamProjectID, "SYSTEM_TEAMPROJECTID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.HostType, "SYSTEM_HOSTTYPE"); err != nil {
		return nil, err
	}

	return ts, nil
}

func setFromAzureDevOpsEnv(s *string, envVar string) error {
	v := os.Getenv(envVar)
	if v == "" {
		return fmt.Errorf("%w: missing env var %s", errNotCallingFromAzureDevOps, envVar)
	}
	*s = v
	return nil
}

// azureDevOpsIDTokenSource retrieves JWT Tokens from Azure DevOps Pipelines.
type azureDevOpsIDTokenSource struct {
	Client                      *httpclient.ApiClient
	AccessToken                 string
	TeamFoundationCollectionURI string
	PlanID                      string
	JobID                       string
	TeamProjectID               string
	HostType                    string
}

// IDToken returns a JWT Token for the specified audience. For Azure DevOps OIDC,
// the audience parameter is ignored as Azure DevOps tokens always use "api://AzureADTokenExchange".
// It will return an error if not running in Azure DevOps Pipelines.
func (a *azureDevOpsIDTokenSource) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	// Azure DevOps OIDC endpoint format
	// Reference: https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/oidctoken/create?view=azure-devops-rest-7.1
	// Use azureDevOpsHostType to determine the hub name dynamically (e.g., "build", "release", etc.)
	requestUrl := fmt.Sprintf("%s/%s/_apis/distributedtask/hubs/%s/plans/%s/jobs/%s/oidctoken?api-version=7.2-preview.1",
		a.TeamFoundationCollectionURI,
		a.TeamProjectID,
		a.HostType,
		a.PlanID,
		a.JobID,
	)

	// Azure DevOps returns {"oidcToken":"***"}, not {"value":"***"}.
	var azureResp struct {
		OIDCToken string `json:"oidcToken"`
	}

	err := a.Client.Do(ctx, "POST", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken)),
		httpclient.WithResponseUnmarshal(&azureResp),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request ID token from Azure DevOps: %w", err)
	}
	if azureResp.OIDCToken == "" {
		return nil, fmt.Errorf("empty OIDC token received from Azure DevOps")
	}

	return &IDToken{Value: azureResp.OIDCToken}, nil
}
