---
name: Provider Issue
about: Use this to identify a issue with the provider.
title: "[ISSUE] Issue with `databricks_XXX` resource"
---

<!--
Hi there,

Please make sure that you checked Troubleshooting Guide first:
https://registry.terraform.io/providers/databricks/databricks/latest/docs/guides/troubleshooting

Thank you for opening an issue.
Please note that we try to keep the Databricks Provider issue tracker reserved for bug reports and feature requests.
For general Terraform usage questions, please see: <https://www.terraform.io/community.html>.
Questions about Databricks Provider please post on https://community.databricks.com
-->

### Configuration
<!-- Please provide a minimal reproducible configuration for the issue -->
```hcl
# Copy-paste your Terraform configuration here
```

### Expected Behavior
<!-- What should have happened? -->

### Actual Behavior
<!-- What actually happened? -->

### Steps to Reproduce
<!-- Please list the steps required to reproduce the issue, for example:
1. `terraform apply`-->

### Terraform and provider versions
<!-- Please paste the output of `terraform version`. If version of `databricks` provider is not the latest (https://github.com/databricks/terraform-provider-databricks/releases), please make sure to use the latest one. -->

### Is it a regression?
<!-- Did this work in a previous version of the provider? If so, which versions did you try? -->

### Debug Output
<!-- Please add turn on logging, e.g. `TF_LOG=DEBUG terraform apply -no-color` and run command again, paste it to gist & provide the link to gist. If you're still willing to paste in log output, make sure you provide only relevant log lines with requests. It would make it more readable, if you pipe the log through `| grep databricks | sed -E 's/^.* plugin[^:]+: (.*)$/\1/'`, e.g.: `TF_LOG=DEBUG terraform apply -no-color 2>&1 | grep databricks | sed -E 's/^.* plugin[^:]+: (.*)$/\1/' 2>&1 |tee tf-debug.log`. If Terraform produced a panic, please provide a link to a GitHub Gist containing the output of the `crash.log`. -->

### Important Factoids
<!-- Are there anything atypical about your accounts that we should know? -->

### Would you like to implement a fix?
<!-- If you plan to implement a fix for this, let the maintainers and community know -->
