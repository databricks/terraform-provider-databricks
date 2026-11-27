# Code Review

Review the given PR or code changes for the following concerns, in order of priority:

1. **Backward compatibility** — Does this change break any existing public API?
   Check for removed or renamed exported types, functions, or methods.
2. **Correctness** — Are there logic errors, race conditions, or unhandled edge cases?
3. **Error handling** — Are errors wrapped with caller context? Are sentinel errors
   used correctly? See `.agent/rules/error-handling.md`.
4. **Test coverage** — Are new code paths tested? Do tests follow the conventions
   in `.agent/rules/testing.md`?
5. **API design** — Does the change follow the principles in `.agent/rules/api-design.md`?
   Is the exported surface minimal?
6. **Style** — Does the code follow `.agent/rules/style-guide.md`? Are names descriptive?
   Are doc comments present on exported symbols?
7. **Generated code boundaries** — Does the change accidentally modify files in
   `service/` or `experimental/mocks/`?

For each issue found, state the file and line, the concern category, and a
concrete suggestion for how to fix it.
