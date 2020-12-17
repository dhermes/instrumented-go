// +build instrumented_main

// NOTE: The `instrumented_main` build tag ensures this won't get picked up
//       as part of the "regularly" built part of the binary. This file will
//       only get compiled if the build tag is explicitly supplied.

package main

import (
	"os"
	"testing"
)

// TestStub is a unit test that invokes `main()` in the context of a
// `testing.T` so `go test` can be run and coverage measured.
func TestStub(t *testing.T) {
	main()
}

func TestMain(m *testing.M) {
	// Swap out `STDOUT` and `STDERR` so our binary does not have output
	// polluted by the `go test` output.
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	os.Stdout = devNull
	os.Stderr = devNull

	os.Exit(m.Run())
}
