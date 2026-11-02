# NCC initial state

These checkpoints are unchanged state files produced by Terraform 1.13 against
local placeholder HTTP responses, using the AWS/Azure create configurations in
the sibling network-connectivity-config scenarios. No state JSON was edited.
They include both the NCC and endpoint rule, including the API-populated account
ID. The regression's configuration omits that computed-only field and requires
an empty plan before automatic cleanup. These fixtures test account-ID stability,
not migration from a historical provider schema.
