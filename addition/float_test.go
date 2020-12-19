package addition_test

import (
	"testing"

	"github.com/dhermes/instrumented-go/addition"
)

func TestAddFloat64(t *testing.T) {
	got := addition.AddFloat64(3.5, 10.75)
	want := 14.25
	if got != want {
		t.Errorf("AddFloat64(3.5, 10.75) = %f; want %f", got, want)
	}
}
