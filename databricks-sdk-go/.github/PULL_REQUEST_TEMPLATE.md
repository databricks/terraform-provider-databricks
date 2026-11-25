<!--
  This template provides a recommended structure for PR descriptions.
  Adapt it freely — the goal is clarity, not rigid compliance.
  The three-section format (Summary / Why / What Changed) helps reviewers
  understand the change quickly and makes the PR easier to revisit later.
-->

## Summary

<!--
  One or two sentences describing what this PR changes and what it enables.
  Focus on the effect, not the implementation details.

  Example:
    Extracts the credentials chain iteration logic into a reusable
    `NewCredentialsChain` constructor so that internal tools can compose
    their own authentication chains from individual credential strategies.
-->

## Why

<!--
  Explain the problem that motivated this change. A reviewer who reads only
  this section should understand why the PR exists and what problem it solves.

  - Start with the status quo: how things work today and what limitation exists.
  - Explain who is affected and what they cannot do (or must work around).
  - If alternatives were considered and rejected, briefly mention why.
  - End with how this PR addresses the gap.

  The "why" is the most important part of a PR description — it usually
  cannot be inferred from the code itself.
-->

## What changed

### Interface changes

<!--
  New or modified public API surface: types, functions, configuration options.
  Use backticks for code references. Write "None." if there are no changes.

  Example:
    - **`NewCredentialsChain(...CredentialsStrategy) CredentialsStrategy`** —
      Takes an ordered list of credential strategies and returns a strategy
      that tries them in sequence.
-->

### Behavioral changes

<!--
  User-visible behavior changes: different defaults, changed error messages,
  new side effects, performance characteristics. Write "None." if this is a
  pure refactor — this explicitly reassures reviewers.
-->

### Internal changes

<!--
  Refactoring, file moves, implementation details, test infrastructure.
  Things that don't affect the public API or user-visible behavior.
-->

## How is this tested?

<!--
  Describe any tests you have done, especially verification that is not part
  of the unit tests (e.g. local checks or manual verification).

  ALWAYS ANSWER THIS QUESTION: answer with "N/A" if tests are not applicable
  to your PR (e.g. if the PR only modifies comments). Do not be afraid of
  answering "Not tested" if the PR has not been tested. Being clear about what
  has been done and not done provides important context to the reviewers.
-->
