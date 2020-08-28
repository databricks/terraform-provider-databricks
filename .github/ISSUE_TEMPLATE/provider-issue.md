---
name: Provider Issue
about: Use this to identify a issue or a bug with the provider.
title: "[ISSUE] Provider bug"
labels: bug
---

Hi there,

Thank you for opening an issue. Please note that we try to keep the Databricks Provider issue tracker reserved for bug reports and feature requests. For general usage questions, please see: https://www.terraform.io/community.html.

### Terraform Version
Run `terraform -v` to show the version. If you are not running the latest version of Terraform, please upgrade because your issue may have already been fixed.

### Affected Resource(s)
Please list the resources as a list, for example:
- databricks_cluster
- databricks_job

If this issue appears to affect multiple resources, it may be an issue with Terraform's core, so please mention this.

### Terraform Configuration Files & Environment Variable Names
```hcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key.
```

### Debug Output
Please add turn on logging, e.g. `TF_LOG=DEBUG terraform apply` and run command again, paste it to gist & provide the link to gist. If you're still willing to paste in log output, make sure you provide only relevant log lines with requests.

### Panic Output
If Terraform produced a panic, please provide a link to a GitHub Gist containing the output of the `crash.log`.

### Expected Behavior
What should have happened?

### Actual Behavior
What actually happened?

### Steps to Reproduce
Please list the steps required to reproduce the issue, for example:
1. `terraform apply`

### Important Factoids
Are there anything atypical about your accounts that we should know? 