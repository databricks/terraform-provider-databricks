# Action pinning — supply-chain protection
#
# External actions must be pinned to a full 40-character commit SHA.
# Mutable tags like @v1 can be reassigned to point at malicious commits.
# Local composite actions (./...) are exempt.
#
# Good:  actions/checkout@11bd71901bbe5b1630ceea73d27597364c9af683
# Bad:   actions/checkout@v4
# Bad:   actions/checkout@main
#
# How to fix:
#   1. Find the tag you want to pin (e.g. v4.3.1).
#   2. Look up the commit SHA:
#        git ls-remote --tags https://github.com/<owner>/<action>.git '<tag>^{}' '<tag>'
#   3. Replace the tag with the SHA and add a comment with the tag name:
#        uses: actions/checkout@<sha> # v4.3.1
#
# Always include the "# <tag>" suffix comment so humans can tell which
# version is pinned. This cannot be enforced by conftest (YAML strips
# comments during parsing), so it is a convention to follow manually.

package main

import rego.v1

_is_pinned(ref) if {
	regex.match(`^[^@]+@[0-9a-f]{40}$`, ref)
}

_is_local(ref) if {
	startswith(ref, "./")
}

# Workflow files: jobs.<name>.steps[].uses
deny contains msg if {
	some job_name, job in input.jobs
	some i, step in job.steps
	step.uses
	not _is_local(step.uses)
	not _is_pinned(step.uses)
	msg := sprintf("%s: step %d: action '%s' must be pinned to a full commit SHA", [job_name, i, step.uses])
}

# Composite actions: runs.steps[].uses
deny contains msg if {
	some i, step in input.runs.steps
	step.uses
	not _is_local(step.uses)
	not _is_pinned(step.uses)
	msg := sprintf("step %d: action '%s' must be pinned to a full commit SHA", [i, step.uses])
}
