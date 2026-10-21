# CodeQL

CodeQL uses the shared Go build environment to resolve dependencies through
JFrog. Fork and Dependabot PRs use the pre-warmed module cache offline, like the
unit tests. If a dependency change is missing from that cache, a maintainer must
run **Warm Go Cache** with the PR number, then re-run CodeQL.

Dependency resolution, package loading (including test imports), and an explicit
Go build must succeed before analysis can complete. These failures fail the job,
which prevents `setup-go` from saving an incomplete cache in its success-only
post step. Successful runs continue to share the normal Go cache with tests.

Failures on main, from pushes or the weekly schedule, notify
`#eng-dev-ecosystem-team`, the same Slack channel as the Terraform nightlies.
PR failures remain visible in the CodeQL check. Slack posts only on the first run
attempt, like the nightly notifier.

Before enabling these notifications, make the existing nightly `SLACK_WEBHOOK`
secret available to this repository as a repository secret or selected
organization secret. A missing secret or notification HTTP error fails the
notification job. The nightly workflow helper is in a private repository, so
this public workflow sends the webhook directly.
