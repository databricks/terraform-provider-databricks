---
name: write-pr-description
description: >
  Write or improve a GitHub pull request description. Use when the user asks
  to "write a PR description", "improve the PR description", "update the PR
  body", or provides a PR URL and asks for a better description.
  Keywords: PR description, pull request summary, PR body, PR writeup.
---

# Write PR Description

Generate a structured PR description that explains **why** the change exists,
not just what files were touched.

## When to use

- The user provides a PR URL and asks to write or improve its description.
- The user asks to draft a PR description for the current branch.
- An agent needs to open a PR and wants a high-quality description.

## Workflow

### Phase 1: Gather context

Collect all the information needed to understand the change:

1. **Read the PR metadata** — title, current description, author, branch name.
2. **Read the full diff** — understand every file changed, every function added
   or removed, every signature change. Do not skip files.
3. **Read surrounding code when needed** — if the diff modifies an interface or
   a struct, read the full file to understand how the change fits into the
   existing architecture.
4. **Check for linked issues or docs** — the PR or commit messages may
   reference issues, design docs, or RFCs that explain motivation.

### Phase 2: Analyze the change

Before writing, answer these questions internally:

- **What was the status quo before this PR?** What limitation, bug, or missing
  capability existed?
- **Why is this change needed now?** What concrete problem does it solve? Who
  benefits?
- **What are the key design decisions?** Why was this approach chosen over
  alternatives?
- **What is the new API surface?** Any new public types, functions, or
  configuration options?
- **What are the architectural changes?** How does the internal structure
  change? What moves where? What gets refactored?
- **Are there behavioral changes?** If not, say so explicitly — this reassures
  reviewers.

### Phase 3: Write the description

Use the structure defined in `.github/PULL_REQUEST_TEMPLATE.md` as the
template. The tone should be direct and technical. Write for a reviewer who is
familiar with the codebase but has not seen this change before.

**Key principles:**

- **Lead with why, not what.** The diff already shows the what. The description
  should explain the reasoning that is not visible in the code.
- **Be specific.** Instead of "improves extensibility", say "allows internal
  tools to compose their own auth chain from individual credential strategies".
- **Name things.** Reference actual types, functions, files, and config fields.
  Use backticks for code references.
- **State non-changes explicitly.** If the PR is a refactor with no behavioral
  change, say "No behavioral changes. Existing users are unaffected." This is
  valuable information for reviewers.
- **Keep the summary to one or two sentences.** It should be scannable.
- **Use the motivation section to tell a story.** What was the problem? Why
  couldn't it be solved before? What does this PR unlock?

### Phase 4: Update the PR

Use the GitHub MCP tools or `gh` CLI to update the PR body with the new
description. Confirm with the user before pushing if unsure.
