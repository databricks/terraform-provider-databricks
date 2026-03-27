# Conftest policies for GitHub Actions

This directory contains [Conftest](https://www.conftest.dev/) policies that
validate GitHub Actions [workflows] and [composite actions]. They are evaluated
by the [conftest workflow](../workflows/conftest.yml) on every push and pull
request that touches `.github/`.

## Adding a new rule

1. Create a new `.rego` file under `rules/`.
2. Use `package main` and add violations to `deny`.
3. Include a comment block at the top of the file explaining the rule and how
   to fix violations.
4. Push — the conftest workflow picks up new rules automatically.

Note that workflows and composite actions have different YAML schemas.
Workflows define jobs under `jobs.<name>.steps`, while composite actions define
steps under `runs.steps`. Rules that inspect steps must handle both.

## Running locally

```bash
# Install conftest (macOS)
brew install conftest

# Run all policies against workflows and composite actions
conftest test \
  .github/workflows/*.yml \
  .github/actions/*/action.yml \
  --policy .github/conftest/rules
```

## References

- [Conftest](https://www.conftest.dev/) — policy testing tool for configuration files
- [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/) — the policy language used by Conftest and OPA
- [Workflow syntax](https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions) — YAML schema for `.github/workflows/*.yml`
- [Composite actions](https://docs.github.com/en/actions/sharing-automations/creating-actions/creating-a-composite-action) — YAML schema for `action.yml` in composite actions
- [Security hardening](https://docs.github.com/en/actions/security-for-github-actions/security-guides/security-hardening-for-github-actions) — GitHub's guide to securing workflows
- [Using third-party actions](https://docs.github.com/en/actions/security-for-github-actions/security-guides/security-hardening-for-github-actions#using-third-party-actions) — why pinning to commit SHAs matters

[workflows]: https://docs.github.com/en/actions/writing-workflows/workflow-syntax-for-github-actions
[composite actions]: https://docs.github.com/en/actions/sharing-automations/creating-actions/creating-a-composite-action
