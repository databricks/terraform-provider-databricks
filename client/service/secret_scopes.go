package service

import (
	"fmt"
	"net/http"

	"github.com/databrickslabs/databricks-terraform/client/model"
)

// SecretScopesAPI exposes the Secret Scopes API
type SecretScopesAPI struct {
	client *DatabricksClient
}

// Create creates a new secret scope
func (a SecretScopesAPI) Create(scope string, initialManagePrincipal string) error {
	return a.client.post("/secrets/scopes/create", map[string]string{
		"scope":                    scope,
		"initial_manage_principal": initialManagePrincipal,
	}, nil)
}

// Delete deletes a secret scope
func (a SecretScopesAPI) Delete(scope string) error {
	return a.client.post("/secrets/scopes/delete", map[string]string{
		"scope": scope,
	}, nil)
}

// List lists all secret scopes available in the workspace
func (a SecretScopesAPI) List() ([]model.SecretScope, error) {
	var listSecretScopesResponse model.SecretScopeList
	err := a.client.get("/secrets/scopes/list", nil, &listSecretScopesResponse)
	return listSecretScopesResponse.Scopes, err
}

// Read will return the metadata for the secret scope
func (a SecretScopesAPI) Read(scopeName string) (model.SecretScope, error) {
	var secretScope model.SecretScope
	scopes, err := a.List()
	if err != nil {
		return secretScope, err
	}
	for _, scope := range scopes {
		if scope.Name == scopeName {
			return scope, nil
		}
	}
	return secretScope, APIError{
		ErrorCode:  "NOT_FOUND",
		Message:    fmt.Sprintf("no Secret Scope found with scope name %s", scopeName),
		Resource:   "/api/2.0/secrets/scopes/list",
		StatusCode: http.StatusNotFound,
	}
}
