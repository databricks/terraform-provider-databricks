# HTTP fixture replay

`Handler` serves stateless request/response exchanges. `StartStep` installs a
step's consumable fixtures and returns any failures from the previous step.
`Verify` checks the current step for unexpected calls and unused fixtures.

Matching includes the method, escaped path, query parameters, JSON body, and
listed request headers. JSON object key order and whitespace are ignored;
array order and values matter. An absent body differs from JSON `null`.

Independent request signatures may interleave. Identical signatures consume
responses FIFO. Requests and fixture consumption are protected by a mutex.
Test-wide reusable exchanges are intended for invariant client initialization,
such as host metadata; they are not required to be consumed.
