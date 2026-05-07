// Package testframeworkv2 is a placeholder package whose only purpose is
// to give fixtures_test.go a package to live in. The framework's runtime
// code lives under cmd/tfv2 and internal/...; this file lets `go test`
// pick up TestFixtures (which programmatically drives every test.yaml
// under issues-repro/ and tests/).
//
// See DESIGN.md §12.7 for the go-test integration design.
package testframeworkv2
