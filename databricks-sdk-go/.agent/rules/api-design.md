---
description: "API design principles for the Databricks SDK."
globs: "**/*.go"
alwaysApply: false
---

# API Design

## Public API Surface

- Minimize exported API surface: if a type is only used through an interface, keep
  the concrete type unexported
- **Return interfaces from constructors** when the concrete type is private and adds
  no methods beyond the interface — this prevents direct construction and keeps
  implementation details hidden
- Prefer new types over modifying existing exported types (additive changes only)
- Adding a required field to a public type is a breaking change — use nil checks
  and defaults to preserve existing behavior
- Unstable APIs go under `experimental/` with `// Experimental: subject to change.` doc comments
- Remove dead or unused code — don't leave unexported functions that nothing calls

## Interface Design

- **Keep interfaces small and focused**: avoid coupling unrelated concerns in one interface.
  If an interface does "too many things", split it
- **Encapsulate domain knowledge**: requirements and validation should live in the component
  that owns the behavior, not in an external orchestrator. Every consumer should not have
  to re-implement the same mapping logic
- Pass specific values, not entire config objects — functions should depend on the narrowest
  interface possible

## Constructors and Defaults

- Validate required fields in constructors — validate early, fail fast
- Validation logic should live near its point of use, not in generic/shared layers
- **Components should own their default values** — don't let callers hard-code defaults
  that belong to the component
- **Zero values should be the default behavior**: use negative field names (`Disable` not
  `Enable`) so that the zero value corresponds to the intended default. A non-zero value
  always corresponds to a special case
- When setting default values (timeouts, retries, cache durations), comment the rationale
- **Use variadic options** (`...Option`) for backward-compatible extensibility

## Immutability and Concurrency

- **Treat config objects as immutable** after construction — do not mutate config fields
  after a type has been built from them
- **Keep concurrency primitives internal**: locks, mutexes, and synchronization should be
  implementation details of the type, not exposed to or managed by callers
- **Separate mechanism from policy**: e.g., separate *how* to iterate strategies from
  *which* strategies to try

## Code Structure

- Place helper functions near their callers, not at the top of the file
- Use `switch/case` for type dispatching, not `if/else` chains
- Split functions with multiple concerns (e.g., create temp file + atomic rename)
  so each has one clear responsibility. Use `defer` to enforce cleanup invariants
- Extract shared logic into helpers at 3+ duplications, not before
- **Consolidate redundant code paths**: when two functions produce the same result,
  have one delegate to the other rather than duplicating logic
