# testframeworkV2 — Makefile shortcuts
#
# Wraps the existing CLI (`go run ./cmd/tfv2 run`) and `go test` entry
# points so day-to-day fixture runs are a one-liner. No new behaviour;
# every target dispatches to commands you can also invoke by hand.
#
# Usage:
#
#   make test <path>     Run a single test (single fixture; CLI path)
#   make test-all        Run every fixture via `go test` (TFV2_RUN=1)
#   make unit            Unit tests only — no cloud auth required
#   make build           Build the tfv2 binary at ./tfv2
#   make clean           Remove the built binary
#   make help            (default) print this banner
#
# `make test <path>` uses the `%: @:` catch-all idiom so the second
# argument (the test directory path) doesn't trigger
# "No rule to make target". `MAKECMDGOALS` carries every goal-name
# the user typed; `filter-out test,...` drops the literal "test"
# leaving only the path.
#
# Trailing slash is tolerated (the CLI normalises via filepath.Abs).

.PHONY: help test test-all unit build clean
.DEFAULT_GOAL := help

# Treat any extra command-line goals (i.e. the path arg to `make test <path>`)
# as phony. Without this, when the path corresponds to an existing directory
# (e.g. `issues-repro/issue_5672`), make filesystem-checks it and prints
# `make: 'issues-repro/issue_5672' is up to date.` after the test output.
.PHONY: $(filter-out test test-all unit build clean help,$(MAKECMDGOALS))

# Catch-all so the path-arg goals passed to `make test <path>` are
# silent no-ops rather than "No rule to make target" errors.
%:
	@:

help:
	@echo "testframeworkV2 — Makefile targets"
	@echo ""
	@echo "  make test <path>     Run a single test (e.g. make test issues-repro/issue_5672)"
	@echo "  make test-all        Run all fixtures via 'go test' (sets TFV2_RUN=1)"
	@echo "  make unit            Run unit tests only (no cloud auth required)"
	@echo "  make build           Build the tfv2 binary at ./tfv2"
	@echo "  make clean           Remove the built binary"

test:
	@PATH_ARG="$(filter-out test,$(MAKECMDGOALS))"; \
	if [ -z "$$PATH_ARG" ]; then \
		echo "Usage: make test <path>"; \
		echo "Example: make test issues-repro/issue_5672"; \
		echo "Or: make test-all  (runs every fixture)"; \
		exit 1; \
	fi; \
	go run ./cmd/tfv2 run $$PATH_ARG

test-all:
	TFV2_RUN=1 go test -run TestFixtures -v ./...

unit:
	go test ./internal/... ./cmd/...

build:
	go build -o tfv2 ./cmd/tfv2

clean:
	rm -f tfv2
