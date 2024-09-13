---
page_title: "Authenticate with OpenID Connect: Google Cloud"
---

# Authenticate with OpenID Connect

OpenID Connect (OIDC) is an authentication protocol allowing users to authenticate to applications without managing long-lived credentials. The Terraform Provider for Databricks can leverage OIDC to authenticate to Databricks accounts and workspaces. For Databricks on Google Cloud, the provider can authenticate leveraging OIDC using workload identity pools. This guide will walk you through the steps to authenticate to Databricks using OIDC on GitHub Actions.

This guide assumes that you have an existing GCP Databricks workspace.

## GitHub Actions

### Configure your service account and workload identity pool

First, you need to create a service account and a workload identity pool. The pool is configured to allow clients using OIDC to assume the identity of the service account. The service account will be used to authenticate to Azure Databricks. You can create a service account using the `google` Terraform provider.

```hcl
provider "google" {
  features {}
}

resource "google_service_account" "github_actions" {
  project      = "<project>"
  account_id   = "github-actions"
  display_name = "GitHub Actions Service Account"
}
```

Then, create the workload identity pool, and configure it to use the service account.

```hcl
resource "google_iam_workload_identity_pool" "github_pool" {
  project                   = "<project>"
  workload_identity_pool_id = "github-pool"
  display_name              = "GitHub Actions Pool"
  description               = "Identity pool for GitHub Actions"
}

resource "google_iam_workload_identity_pool_provider" "github_provider" {
  project                            = "<project>"
  workload_identity_pool_id          = google_iam_workload_identity_pool.github_pool.workload_identity_pool_id
  workload_identity_pool_provider_id = "github-provider"
  display_name                       = "GitHub Actions Provider"
  attribute_mapping = {
    "google.subject"       = "assertion.sub"
    "attribute.actor"      = "assertion.actor"
    "attribute.repository" = "assertion.repository"
    "attribute.ref"        = "assertion.ref"
    "attribute.event_name" = "assertion.event_name"
  }
  oidc {
    issuer_uri = "https://token.actions.githubusercontent.com"
  }
  attribute_condition = "assertion.sub == 'repo:<organization>/<repo>:environment:<environment>'"
}

resource "google_service_account_iam_binding" "workload_identity_user" {
  service_account_id = google_service_account.github_actions.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "principalSet://iam.googleapis.com/${google_iam_workload_identity_pool.github_pool.name}/attribute.repository/<organization>/<repo>"
  ]
}

resource "google_project_iam_member" "token_creator_binding" {
  project = module.defaults.google_project_isolated
  role    = "roles/iam.serviceAccountTokenCreator"
  member  = "serviceAccount:${google_service_account.github_actions.email}"
}
```

Finally, grant the service principal access to the workspace by following the instructions in the [Databricks documentation](https://docs.gcp.databricks.com/en/dev-tools/google-id-auth.html#step-2-assign-your-google-cloud-service-account-to-your-databricks-account).

### Configure the Databricks provider to use the service principal

In your Terraform configuration, configure the Databricks provider to use the service principal.

```hcl
# account-level provider
provider "databricks" {
  host                   = "https://accounts.gcp.databricks.com"
  account_id             = "<databricks-account-id>"
  google_service_account = google_service_account.github_actions.email
}

# workspace-level provider
provider "databricks" {
  host                   = "https://<workspace-url>"
  google_service_account = google_service_account.github_actions.email
}
```

### Create a GitHub Action that authenticates to Databricks on Google Cloud

To create a GitHub Action, make a `.github/workflows/deploy.yml` file in your repository.

To authenticate to Databricks using OIDC, ensure that your action has the `id-token: write` permission. You can then authenticate to Google using the `google-github-actions/auth` action. Finally, run `terraform apply`.

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

      - name: "Authenticate to Google Cloud"
        uses: "google-github-actions/auth@v2"
        with:
          token_format: "access_token"
          workload_identity_provider: "projects/<project-id>/locations/global/workloadIdentityPools/github-pool/providers/github-provider"
          service_account: "<service-account-email>"

      - name: Set up Terraform
        uses: hashicorp/setup-terraform@v3

      - name: Terraform Init
        run: terraform init
        working-directory: path/to/terraform/module
```

### (Optional) GitHub Actions Details

The `subject` field is used to scope the federated credentials to a specific GitHub Actions environment. The `subject` field is a string in the format `repo:<organization>/<repo>:environment:<environment>`. The `organization`, `repo`, and `environment` fields should be replaced with the appropriate values.

If the action runs without an environment context, the `subject` field should be set to `repo:<organization>/<repo>:ref:refs/heads/<BRANCH-NAME>` if the workflow is triggered from a branch, or `repo:<organization>/<repo>:ref:refs/tags/<TAG-NAME>` when triggered from a tag.

If needed, it is also possible to configure the `subject` field for your organization or repository. See the [GitHub Actions OIDC documentation](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect) for more information about how to configure the `subject` field.

### (Optional) Restrict access to the workload identity pool

The workload identity pool provider can be configured to restrict access to specific repositories, branches, or tags. The `attribute_condition` field in the `google_iam_workload_identity_pool_provider` resource specifies the conditions under which the provider will issue tokens. See [the Google Cloud reference](https://cloud.google.com/iam/docs/workload-identity-federation#conditions) for more information.

## References

For more information about OIDC and the above OIDC providers, see the following resources:

- [GitHub Actions OIDC documentation](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect)
- [Google Cloud Workload Identity documentation](https://cloud.google.com/iam/docs/workload-identity-federation)
