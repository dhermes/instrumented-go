// +build instrumented_main

// NOTE: The `instrumented_main` build tag ensures this won't get picked up
//       as part of the "regularly" built part of the binary. This file will
//       only get compiled if the build tag is explicitly supplied.

package main

import (
	"testing"
)

// TestStub is a unit test that invokes `main()` in the context of a
// `testint.T` so `go test` can be run and coverage measured.
func TestStub(t *testing.T) {
	main()
}
