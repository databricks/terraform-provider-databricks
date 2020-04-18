package service

import (
	db "github.com/databrickslabs/databricks-terraform/client"
)

var scimHeaders = map[string]string{
	"Content-Type": "application/scim+json",
}

type DBApiClient struct {
	config *db.DBApiClientConfig
}

// SetConfig initializes the client
func (c *DBApiClient) SetConfig(clientConfig *db.DBApiClientConfig) DBApiClient {
	c.config = clientConfig
	clientConfig.Setup()
	return *c
}

// Clusters returns an instance of ClustersAPI
func (c DBApiClient) Clusters() ClustersAPI {
	return ClustersAPI{Client: c}
}

func (c DBApiClient) Secrets() SecretsAPI {
	return SecretsAPI{Client: c}
}

func (c DBApiClient) SecretScopes() SecretScopesAPI {
	return SecretScopesAPI{Client: c}
}

func (c DBApiClient) SecretAcls() SecretAclsAPI {
	return SecretAclsAPI{Client: c}
}

func (c *DBApiClient) Tokens() TokensAPI {
	return TokensAPI{Client: c}
}

func (c DBApiClient) Users() UsersAPI {
	return UsersAPI{Client: c}
}

func (c DBApiClient) Groups() GroupsAPI {
	return GroupsAPI{Client: c}
}

func (c DBApiClient) Notebooks() NotebooksAPI {
	return NotebooksAPI{Client: c}
}

func (c DBApiClient) Jobs() JobsAPI {
	return JobsAPI{Client: c}
}

func (c DBApiClient) DBFS() DBFSAPI {
	return DBFSAPI{Client: c}
}

func (c DBApiClient) Libraries() LibrariesAPI {
	return LibrariesAPI{Client: c}
}

func (c DBApiClient) InstancePools() InstancePoolsAPI {
	return InstancePoolsAPI{Client: c}
}

func (c DBApiClient) InstanceProfiles() InstanceProfilesAPI {
	return InstanceProfilesAPI{Client: c}
}

func (c DBApiClient) Commands() CommandsAPI {
	return CommandsAPI{Client: c}
}

func (c DBApiClient) performQuery(method, path string, apiVersion string, headers map[string]string, data interface{}) ([]byte, error) {
	return db.PerformQuery(c.config, method, path, apiVersion, headers, true, false, data)
}
