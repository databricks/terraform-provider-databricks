---
page_title: "Authenticate with OpenID Connect: Azure"
---

# Authenticate with OpenID Connect

OpenID Connect (OIDC) is an authentication protocol allowing users to authenticate to applications without managing long-lived credentials. The Terraform Provider for Databricks can leverage OIDC to authenticate to Databricks accounts and workspaces. For Azure Databricks, the provider uses the Azure CLI to authenticate using OIDC. This guide will walk you through the steps to authenticate to Azure Databricks using OIDC on GitHub Actions and Azure DevOps.

This guide assumes that you have an existing Azure Databricks workspace.

## GitHub Actions

### Configure your service principal with federated credentials

First, you need to create a service principal with federated credentials. This service principal will be used to authenticate to Azure Databricks. You can create a service principal using the `azuread` Terraform provider.

```hcl
provider "azurerm" {
  features {}
}

resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application_registration.example.application_id
}
```

Then, configure the service principal to use federated credentials issued by GitHub Actions.

```hcl
resource "azuread_application_federated_identity_credential" "example" {
  application_id = azuread_application_registration.example.id
  display_name   = "my-repo-deploy"
  description    = "Deployments for my-repo"
  audiences      = ["api://AzureADTokenExchange"]
  issuer         = "https://token.actions.githubusercontent.com"
  subject        = "repo:<organization>/<repo>:environment:<environment>"
}
```

Finally, grant the service principal access to the workspace.

```hcl
resource "azurerm_role_assignment" "example" {
    scope                = "/subscriptions/<subscription-id>/resourceGroups/<resource-group>/providers/Microsoft.Databricks/workspaces/<workspace>"
    role_definition_name = "Contributor"
    principal_id         = azuread_service_principal.example.id
}
```

### Configure the Databricks provider to use the service principal

In your Terraform configuration, configure the Databricks provider to use the service principal.

```hcl
provider "databricks" {
  azure_client_id = "<application-id>"
  azure_tenant_id = "<tenant-id>"
  host = "https://<workspace-url>"
}
```

### Create a GitHub Action that authenticates to Azure Databricks

To create a GitHub Action, make a `.github/workflows/deploy.yml` file in your repository.

To authenticate to Azure Databricks using OIDC, ensure that your action has the `id-token: write` permission. You can then authenticate to Azure using the `azure/login` action. Finally, run `terraform apply` with the `azure/cli` action.

```yaml
name: Deploy to Azure Databricks
jobs:
  deploy:
    runs-on: ubuntu-latest
    environments: production
    permissions:
      id-token: write
      contents: read
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Set up Terraform
        uses: hashicorp/setup-terraform@v3

      - name: Azure login
        uses: azure/login@v2
        with:
          client-id: ${{ secrets.AZURE_CLIENT_ID }}
          tenant-id: ${{ secrets.AZURE_TENANT_ID }}
          subscription-id: ${{ secrets.AZURE_SUBSCRIPTION_ID }}

      - name: Azure CLI script
        uses: azure/cli@v2
        with:
          azcliversion: latest
          inlineScript: |
            cd path/to/terraform/module
            terraform init
            terraform plan
            terraform apply
```

### (Optional) GitHub Actions Details

The `subject` field is used to scope the federated credentials to a specific GitHub Actions environment. The `subject` field is a string in the format `repo:<organization>/<repo>:environment:<environment>`. The `organization`, `repo`, and `environment` fields should be replaced with the appropriate values.

If the action runs without an environment context, the `subject` field should be set to `repo:<organization>/<repo>:ref:refs/heads/<BRANCH-NAME>` if the workflow is triggered from a branch, or `repo:<organization>/<repo>:ref:refs/tags/<TAG-NAME>` when triggered from a tag.

If needed, it is also possible to configure the `subject` field for your organization or repository. See the [GitHub Actions OIDC documentation](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect) for more information about how to configure the `subject` field.

## Azure DevOps

### Configure a service connection for your DevOps pipeline

First, you need to create a service connection with federated credentials. This service principal will be used to authenticate to Azure Databricks. You can create a service principal using the `azuread` Terraform provider.

```hcl
provider "azurerm" {
  features {}
}

resource "azuread_application_registration" "example" {
  display_name = "example"
}

resource "azuread_service_principal" "example" {
  application_id = azuread_application_registration.example.application_id
}
```

Then, configure the service principal to use federated credentials issued by Azure DevOps.

```hcl
resource "azuread_application_federated_identity_credential" "example" {
  application_id = azuread_application_registration.example.id
  display_name   = "my-repo-deploy"
  description    = "Deployments for my-repo"
  audiences      = ["api://AzureADTokenExchange"]
  issuer         = "https://vstoken.dev.azure.com/<organization-id>"
  subject        = "sc://<organisation-name>/<project-name>/<service-connection-name>"
}
```

Finally, grant the service principal access to the workspace.

```hcl
resource "azurerm_role_assignment" "example" {
    scope                = "/subscriptions/<subscription-id>/resourceGroups/<resource-group>/providers/Microsoft.Databricks/workspaces/<workspace>"
    role_definition_name = "Contributor"
    principal_id         = azuread_service_principal.example.id
}
```

In Azure DevOps, navigate to the project settings and create a new service connection. Select `Azure Resource Manager`, then `Workload Identity federation (manual)` and enter the subscription ID, subscription name, service principal ID and tenant ID in the dialog. Note that the Issuer and Subject Identifier fields must match the `issuer` and `subject` attributes of the `azuread_application_federated_identity_credential` resource.

### Configure the Databricks provider to use the service principal

In your Terraform configuration, configure the Databricks provider to use the service principal.

```hcl
provider "databricks" {
  azure_client_id = "<application-id>"
  azure_tenant_id = "<tenant-id>"
  host = "https://<workspace-url>"
}
```

### Create a DevOps Pipeline that authenticates to Azure Databricks

To create a pipeline, make a `pipelines/deploy.yml` file in your repository.

To authenticate to Azure Databricks using OIDC, use the `AzureCLI@2` task. This automatically authenticates the Azure CLI using the service connection you created earlier. The Terraform Provider for Databricks will detect the authenticated CLI and use it to authenticate to Azure Databricks.

```yaml
variables:
  DATABRICKS_HOST: "https://test-shard-dbc-6baa92c2-7251.dev.databricks.com/"
  DATABRICKS_CLIENT_ID: "7aacab6a-d0be-4581-ac9c-2875e7796a18"

steps:
  - task: Checkout@1
    displayName: "Checkout repository"
    inputs:
      repository: "self"
      path: "main"

  - task: TerraformInstaller@0
    inputs:
      terraformVersion: "latest"

  - task: AzureCLI@2
    displayName: "TF init"
    inputs:
      addSpnToEnvironment: true
      azureSubscription: <service-connection-name>
      scriptType: bash
      scriptLocation: inlineScript
      workingDirectory: "$(Pipeline.Workspace)/main/<repo>/path/to/terraform/module"
      inlineScript: |
        terraform init
```

## References

For more information about OIDC and the above OIDC providers, see the following resources:

- [GitHub Actions OIDC documentation](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect)
- [Azure DevOps Workload federation blog post](https://devblogs.microsoft.com/devops/introduction-to-azure-devops-workload-identity-federation-oidc-with-terraform/)
