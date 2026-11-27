# Contributing to the Databricks Terraform Provider

We welcome community contributions through the [issue tracker](https://github.com/databricks/terraform-provider-databricks/issues). Bug reports, root-cause analyses, feature requests, and documentation fixes help us understand what matters most and improve the provider.

**We won't merge external pull requests.** The provider is developed in Databricks' internal repository and published here. Pull requests opened in this repository will not be merged.

## Why we don't accept pull requests

Thank you to everyone who has contributed PRs, reviews, and feedback over the years. Your work has helped shape this provider. Open source development has changed a lot in the last year, though. Writing a PR is now cheap, but reviewing one and making sure it fits the internal architecture of Databricks is not, and our small team can't keep up with the growing number of issues and PRs.

Using issues as the entry point works better for everyone:

- **Most of the provider is generated.** Resources and data sources are generated from the Databricks API specification, which lives in our internal repository. Changes to generated code were never possible in this repo.
- **Handwritten code needs internal context.** PRs on the handwritten parts of the provider often depend on context that is only available inside Databricks. That leads to long, frustrating review cycles for contributors.
- **Changes ship with the API.** Provider changes and the API changes they depend on can be built, tested, and released together.
- **API teams own their resources.** APIs changing in ways that break declarative tools like Terraform has been a long-standing frustration. With the provider next to the APIs, these regressions can be caught before they roll out.

Incoming issues are triaged and either fixed in the provider or routed to the owning product team when the problem is in the API itself rather than the resource.

The best way to help is with detailed bug reports, reproduction steps, debug logs, root-cause analysis, and design discussion in issues. Understanding the problem and picking the right fix is usually the hard part. Once that is clear, the implementation is the easy bit.

## Reporting bugs

Before opening a new issue, check the [troubleshooting guide](https://registry.terraform.io/providers/databricks/databricks/latest/docs/guides/troubleshooting) and search the [existing issues](https://github.com/databricks/terraform-provider-databricks/issues). If someone already reported the problem, add your details to that issue instead of opening a new one.

Open a [provider issue](https://github.com/databricks/terraform-provider-databricks/issues/new?template=provider-issue.md) and include as much of the following as you can:

- The provider version and the output of `terraform version`. If you aren't on the [latest release](https://github.com/databricks/terraform-provider-databricks/releases), please try it first.
- A minimal Terraform configuration that reproduces the problem.
- The steps to reproduce, plus the expected and actual behavior.
- The debug log from `TF_LOG_PROVIDER=DEBUG terraform apply` (or `plan`). **Remove tokens, passwords, workspace URLs, and other sensitive information first.**
- Whether this is a regression, and if so, the last version where it worked.
- Any root-cause analysis, technical observations, or ideas for a fix.

## Requesting features or missing fields

Open a [feature request](https://github.com/databricks/terraform-provider-databricks/issues/new?template=feature-request.md), or add a 👍 to an existing one that describes the same need. Explain your use case and the behavior you'd like. If the request is for a new resource, data source, or field, link to the relevant page in the [Databricks REST API reference](https://docs.databricks.com/api/).

## Reporting documentation issues

Open a [docs issue](https://github.com/databricks/terraform-provider-databricks/issues/new?template=docs-issue.md) with a link to the affected page (including the section anchor) and a description of what is wrong or missing.

## Questions

The issue tracker is for bug reports and feature requests. For other questions:

- Databricks Terraform provider usage: [Databricks Community](https://community.databricks.com) or [Stack Overflow](https://stackoverflow.com/questions/tagged/terraform-provider-databricks) with the `terraform-provider-databricks` tag.
- General Databricks questions: [Databricks Help Center](https://help.databricks.com/).
- General Terraform questions: [Terraform community](https://www.terraform.io/community).

## Security

If you find a security vulnerability, don't open a public issue. Follow the [security policy](SECURITY.md) and report it to [security@databricks.com](mailto:security@databricks.com).
