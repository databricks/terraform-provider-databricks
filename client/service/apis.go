package service

var scimHeaders = map[string]string{
	"Content-Type": "application/scim+json",
}

// DBApiClient is the client struct that contains clients for all the services available on Databricks
type DBApiClient struct {
	config *DBApiClientConfig
}

// SetConfig initializes the client
func (c *DBApiClient) SetConfig(clientConfig *DBApiClientConfig) DBApiClient {
	c.config = clientConfig
	clientConfig.Setup()
	return *c
}

// Clusters returns an instance of ClustersAPI
func (c DBApiClient) Clusters() ClustersAPI {
	return ClustersAPI{Client: c}
}

// Secrets returns an instance of SecretsAPI
func (c DBApiClient) Secrets() SecretsAPI {
	return SecretsAPI{Client: c}
}

// SecretScopes returns an instance of SecretScopesAPI
func (c DBApiClient) SecretScopes() SecretScopesAPI {
	return SecretScopesAPI{Client: c}
}

// SecretAcls returns an instance of SecretAclsAPI
func (c DBApiClient) SecretAcls() SecretAclsAPI {
	return SecretAclsAPI{Client: c}
}

// Tokens returns an instance of TokensAPI
func (c *DBApiClient) Tokens() TokensAPI {
	return TokensAPI{Client: c}
}

// Users returns an instance of UsersAPI
func (c DBApiClient) Users() UsersAPI {
	return UsersAPI{Client: c}
}

// Groups returns an instance of GroupsAPI
func (c DBApiClient) Groups() GroupsAPI {
	return GroupsAPI{Client: c}
}

// Notebooks returns an instance of NotebooksAPI
func (c DBApiClient) Notebooks() NotebooksAPI {
	return NotebooksAPI{Client: c}
}

// Jobs returns an instance of JobsAPI
func (c DBApiClient) Jobs() JobsAPI {
	return JobsAPI{Client: c}
}

// DBFS returns an instance of DBFSAPI
func (c DBApiClient) DBFS() DBFSAPI {
	return DBFSAPI{Client: c}
}

// Libraries returns an instance of LibrariesAPI
func (c DBApiClient) Libraries() LibrariesAPI {
	return LibrariesAPI{Client: c}
}

// InstancePools returns an instance of InstancePoolsAPI
func (c DBApiClient) InstancePools() InstancePoolsAPI {
	return InstancePoolsAPI{Client: c}
}

// InstanceProfiles returns an instance of InstanceProfilesAPI
func (c DBApiClient) InstanceProfiles() InstanceProfilesAPI {
	return InstanceProfilesAPI{Client: c}
}

// Commands returns an instance of CommandsAPI
func (c DBApiClient) Commands() CommandsAPI {
	return CommandsAPI{Client: c}
}

func (c DBApiClient) performQuery(method, path string, apiVersion string, headers map[string]string, data interface{}, secretsMask *SecretsMask) ([]byte, error) {
	return PerformQuery(c.config, method, path, apiVersion, headers, true, false, data, secretsMask)
}
