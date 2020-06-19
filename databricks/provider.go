package databricks

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/databrickslabs/databricks-terraform/client/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

// Provider returns the entire terraform provider object
func Provider(version string) terraform.ResourceProvider {
	provider := &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"databricks_default_user_roles": dataSourceDefaultUserRoles(),
			"databricks_notebook":           dataSourceNotebook(),
			"databricks_notebook_paths":     dataSourceNotebookPaths(),
			"databricks_dbfs_file":          dataSourceDBFSFile(),
			"databricks_dbfs_file_paths":    dataSourceDBFSFilePaths(),
			"databricks_zones":              dataSourceClusterZones(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"databricks_token":         resourceToken(),
			"databricks_secret_scope":  resourceSecretScope(),
			"databricks_secret":        resourceSecret(),
			"databricks_secret_acl":    resourceSecretACL(),
			"databricks_instance_pool": resourceInstancePool(),
			"databricks_scim_user":     resourceScimUser(),
			"databricks_scim_group":    resourceScimGroup(),
			// Scim Group is split into multiple components for flexibility to pick and choose
			"databricks_group":                  resourceGroup(),
			"databricks_group_instance_profile": resourceGroupInstanceProfile(),
			"databricks_group_member":           resourceGroupMember(),
			"databricks_notebook":               resourceNotebook(),
			"databricks_cluster":                resourceCluster(),
			"databricks_cluster_policy":         resourceClusterPolicy(),
			"databricks_job":                    resourceJob(),
			"databricks_dbfs_file":              resourceDBFSFile(),
			"databricks_dbfs_file_sync":         resourceDBFSFileSync(),
			"databricks_instance_profile":       resourceInstanceProfile(),
			"databricks_aws_s3_mount":           resourceAWSS3Mount(),
			"databricks_azure_blob_mount":       resourceAzureBlobMount(),
			"databricks_azure_adls_gen1_mount":  resourceAzureAdlsGen1Mount(),
			"databricks_azure_adls_gen2_mount":  resourceAzureAdlsGen2Mount(),
			//	MWS (multiple workspaces) resources are only limited to AWS as azure already has a built in concept of MWS
			"databricks_mws_credentials":            resourceMWSCredentials(),
			"databricks_mws_storage_configurations": resourceMWSStorageConfigurations(),
			"databricks_mws_networks":               resourceMWSNetworks(),
			"databricks_mws_workspaces":             resourceMWSWorkspaces(),
		},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_HOST", nil),
			},
			"token": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				DefaultFunc:   schema.EnvDefaultFunc("DATABRICKS_TOKEN", nil),
				ConflictsWith: []string{"basic_auth"},
			},
			"basic_auth": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_USERNAME", nil),
						},
						"password": {
							Type:        schema.TypeString,
							Sensitive:   true,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_PASSWORD", nil),
						},
					},
				},
				ConflictsWith: []string{"token"},
			},
			"config_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_CONFIG_FILE", "~/.databrickscfg"),
				Description: "Location of the Databricks CLI credentials file, that is created\n" +
					"by `databricks configure --token` command. By default, it is located\n" +
					"in ~/.databrickscfg. Check  https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for docs. Config\n" +
					"file credetials will only be used when host/token are not provided.",
			},
			"profile": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEFAULT",
				Description: "Connection profile specified within ~/.databrickscfg. Please check\n" +
					"https://docs.databricks.com/dev-tools/cli/index.html#connection-profiles for documentation.",
			},
			"azure_auth": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_resource_group": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP", nil),
						},
						"azure_region": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("AZURE_REGION", nil),
						},
						"workspace_name": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_WORKSPACE_NAME", nil),
						},
						"resource_group": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("DATABRICKS_AZURE_RESOURCE_GROUP", nil),
						},
						"subscription_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_SUBSCRIPTION_ID", "ARM_SUBSCRIPTION_ID"}, nil),
						},
						"client_secret": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_SECRET", "ARM_CLIENT_SECRET"}, nil),
						},
						"client_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_CLIENT_ID", "ARM_CLIENT_ID"}, nil),
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.MultiEnvDefaultFunc([]string{"DATABRICKS_AZURE_TENANT_ID", "ARM_TENANT_ID"}, nil),
						},
					},
				},
			},
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		//terraformVersion := provider.TerraformVersion
		//if terraformVersion == "" {
		// Terraform 0.12 introduced this field to the protocol
		// We can therefore assume that if it's missing it's 0.10 or 0.11
		//terraformVersion = "0.11+compatible"
		//}
		return providerConfigure(d, version)
	}

	return provider
}

func providerConfigureAzureClient(d *schema.ResourceData, config *service.DBApiClientConfig) (interface{}, error) {
	log.Println("Creating db client via azure auth!")
	azureAuth, _ := d.GetOk("azure_auth")
	azureAuthMap := azureAuth.(map[string]interface{})
	//azureAuth AzureAuth{}
	tokenPayload := TokenPayload{}
	// The if else is required for the reason that "azure_auth" schema object is not a block but a map
	// Maps do not inherently auto populate defaults from environment variables unless we explicitly assign values
	// This makes it very difficult to test
	if managedResourceGroup, ok := azureAuthMap["managed_resource_group"].(string); ok {
		tokenPayload.ManagedResourceGroup = managedResourceGroup
	} else if os.Getenv("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP") != "" {
		tokenPayload.ManagedResourceGroup = os.Getenv("DATABRICKS_AZURE_MANAGED_RESOURCE_GROUP")
	}
	if azureRegion, ok := azureAuthMap["azure_region"].(string); ok {
		tokenPayload.AzureRegion = azureRegion
	} else if os.Getenv("AZURE_REGION") != "" {
		tokenPayload.AzureRegion = os.Getenv("AZURE_REGION")
	}
	if resourceGroup, ok := azureAuthMap["resource_group"].(string); ok {
		tokenPayload.ResourceGroup = resourceGroup
	} else if os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP") != "" {
		tokenPayload.ResourceGroup = os.Getenv("DATABRICKS_AZURE_RESOURCE_GROUP")
	}
	if workspaceName, ok := azureAuthMap["workspace_name"].(string); ok {
		tokenPayload.WorkspaceName = workspaceName
	} else if os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME") != "" {
		tokenPayload.WorkspaceName = os.Getenv("DATABRICKS_AZURE_WORKSPACE_NAME")
	}

	// This provider takes DATABRICKS_AZURE_* for client ID etc
	// The azurerm provider uses ARM_* for the same values
	// To make it easier to use the two providers together we use the following sources in order:
	//  - provider config
	//  - DATABRICKS_AZURE_* environment variables
	//  - ARM_* environment variables

	subscriptionID, ok := azureAuthMap["subscription_id"].(string)
	switch {
	case ok:
		tokenPayload.SubscriptionID = subscriptionID
	case os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID") != "":
		tokenPayload.SubscriptionID = os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID")
	case os.Getenv("ARM_SUBSCRIPTION_ID") != "":
		tokenPayload.SubscriptionID = os.Getenv("DATABRICKS_AZURE_SUBSCRIPTION_ID")
	}

	clientSecret, ok := azureAuthMap["client_secret"].(string)
	switch {
	case ok:
		tokenPayload.ClientSecret = clientSecret
	case os.Getenv("DATABRICKS_AZURE_CLIENT_SECRET") != "":
		tokenPayload.ClientSecret = os.Getenv("DATABRICKS_AZURE_CLIENT_SECRET")
	case os.Getenv("ARM_CLIENT_SECRET") != "":
		tokenPayload.ClientSecret = os.Getenv("ARM_CLIENT_SECRET")
	}

	clientID, ok := azureAuthMap["client_id"].(string)
	switch {
	case ok:
		tokenPayload.ClientID = clientID
	case os.Getenv("DATABRICKS_AZURE_CLIENT_ID") != "":
		tokenPayload.ClientID = os.Getenv("DATABRICKS_AZURE_CLIENT_ID")
	case os.Getenv("ARM_CLIENT_ID") != "":
		tokenPayload.ClientID = os.Getenv("ARM_CLIENT_ID")
	}

	tenantID, ok := azureAuthMap["tenant_id"].(string)
	switch {
	case ok:
		tokenPayload.TenantID = tenantID
	case os.Getenv("DATABRICKS_AZURE_TENANT_ID") != "":
		tokenPayload.TenantID = os.Getenv("DATABRICKS_AZURE_TENANT_ID")
	case os.Getenv("ARM_TENANT_ID") != "":
		tokenPayload.TenantID = os.Getenv("ARM_TENANT_ID")
	}

	azureAuthSetup := AzureAuth{
		TokenPayload:           &tokenPayload,
		ManagementToken:        "",
		AdbWorkspaceResourceID: "",
		AdbAccessToken:         "",
		AdbPlatformToken:       "",
	}

	// Setup the CustomAuthorizer Function to be called at API invoke rather than client invoke
	config.CustomAuthorizer = azureAuthSetup.initWorkspaceAndGetClient
	var dbClient service.DBApiClient
	dbClient.SetConfig(config)
	return &dbClient, nil
}

// tryDatabricksCliConfigFile sets Host and Token from ~/.databrickscfg file if it exists
func tryDatabricksCliConfigFile(d *schema.ResourceData, config *service.DBApiClientConfig) error {
	configFile, err := homedir.Expand(d.Get("config_file").(string))
	if err != nil {
		return err
	}
	cfg, err := ini.Load(configFile)
	if err != nil {
		return fmt.Errorf("Authentication is not configured for provider. Please configure it\n"+
			"through one of the following options:\n"+
			"1. DATABRICKS_HOST + DATABRICKS_TOKEN environment variables.\n"+
			"2. host + token provider arguments.\n"+
			"3. Run `databricks configure --token` that will create %s file.\n\n"+
			"Please check https://docs.databricks.com/dev-tools/cli/index.html#set-up-authentication for details", configFile)
	}
	if profile, ok := d.GetOk("profile"); ok {
		dbcliConfig := cfg.Section(profile.(string))
		token := dbcliConfig.Key("token").String()
		if "" == token {
			return fmt.Errorf("config file %s is corrupt: cannot find token in %s profile",
				configFile, profile)
		}
		config.Token = token

		host := dbcliConfig.Key("host").String()
		if "" == host {
			return fmt.Errorf("config file %s is corrupt: cannot find host in %s profile",
				configFile, profile)
		}
		config.Host = host
	}

	return nil
}

func providerConfigure(d *schema.ResourceData, providerVersion string) (interface{}, error) {
	var config service.DBApiClientConfig
	// Call setup to configure retryable httpclient
	config.Setup()

	//version information from go-releaser using -ldflags to tell the golang linker to send semver info
	config.UserAgent = fmt.Sprintf("databricks-tf-provider/%s", providerVersion)

	if _, ok := d.GetOk("azure_auth"); !ok {
		if host, ok := d.GetOk("host"); ok {
			config.Host = host.(string)
		}
		if token, ok := d.GetOk("token"); ok {
			config.Token = token.(string)
		}

		// Basic authentication setup via username and password
		if _, ok := d.GetOk("basic_auth"); ok {
			username, userOk := d.GetOk("basic_auth.0.username")
			password, passOk := d.GetOk("basic_auth.0.password")
			if userOk && passOk {
				tokenUnB64 := fmt.Sprintf("%s:%s", username.(string), password.(string))
				config.Token = base64.StdEncoding.EncodeToString([]byte(tokenUnB64))
				config.AuthType = service.BasicAuth
			}
		}

		// Final catch all in case basic_auth/token + host is not setup
		if config.Host == "" || config.Token == "" {
			if err := tryDatabricksCliConfigFile(d, &config); err != nil {
				return nil, fmt.Errorf("failed to get credentials from config file; error msg: %w", err)
			}
		}
	} else {
		// Abstracted logic to another function that returns a interface{}, error to inject directly
		// for the providers during cloud integration testing
		return providerConfigureAzureClient(d, &config)
	}

	var dbClient service.DBApiClient
	dbClient.SetConfig(&config)
	return &dbClient, nil
}
