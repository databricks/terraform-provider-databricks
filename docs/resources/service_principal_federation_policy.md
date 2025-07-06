---
subcategory: "Security"
---

# databricks_service_principal_federation_policy Resource

Service principal federation, also known as Workload Identity Federation, allows your automated workloads running outside of Databricks to securely access Databricks APIs without the need for Databricks secrets. With Workload Identity Federation, your application (or workload) authenticates to Databricks as a Databricks service principal, using tokens provided by the workload runtime.

-> This resource can only be used with an account-level provider!

## Example Usage

Creating service principal federation policy:

```hcl

resource "databricks_service_principal" "sp" {
  display_name = "Admin SP"
}

resource "databricks_service_principal_federation_policy" "dspfp" {
  service_principal_id = databricks_service_principal.sp.id
  oidc_policy = {
    issuer = "https://idp.mycompany.com/oidc"
    audiences = ["2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"]
    jwks_uri: "https://idp.mycompany.com/jwks.json"
  }
}

```

The following arguments are available:

- `service_principal_id` - (Required) The service principal id for the federation policy.
- `description` - (Optional) Description of the federation policy.
- `name` - (Optional) Resource name for the federation policy.

### oidc_policy Configuration Block (Required) Specifies the policy to use for validating OIDC claims in your federated tokens.

- `audiences` - (Optional) The allowed token audiences, as specified in the 'aud' claim of federated tokens. The audience identifier is intended to represent the recipient of the token. Can be any non-empty string value. As long as the audience in the token matches at least one audience in the policy, the token is considered a match. If audiences is unspecified, defaults to your Databricks account id.
- `issuer` - (Required) The required token issuer, as specified in the 'iss' claim of federated tokens.
- `jwks_json` - (Optional) The public keys used to validate the signature of federated tokens, in JWKS format. Most use cases should not need to specify this field. If jwks_uri and jwks_json are both unspecified (recommended), Databricks automatically fetches the public keys from your issuer’s well known endpoint. Databricks strongly recommends relying on your issuer’s well known endpoint for discovering public keys.
- `jwks_uri` - (Optional) URL of the public keys used to validate the signature of federated tokens, in JWKS format. Most use cases should not need to specify this field. If jwks_uri and jwks_json are both unspecified (recommended), Databricks automatically fetches the public keys from your issuer’s well known endpoint. Databricks strongly recommends relying on your issuer’s well known endpoint for discovering public keys.
- `subject` - (Required) The required token subject, as specified in the subject claim of federated tokens. Must be specified for service principal federation policies. Must not be specified for account federation policies.
- `subject_claim` - (Optional) The claim that contains the subject of the token. If unspecified, the default value is 'sub'.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `uid` - Unique, immutable id of the federation policy.
- `create_time` - Creation time of the federation policy.
- `update_time` - Last update time of the federation policy.

## Related Resources

The following resources are often used in the same context:

- [databricks_service_principal](service_principal.md) to Directly manage service principal
