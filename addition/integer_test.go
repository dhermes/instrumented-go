package addition_test

import (
	"testing"

	"github.com/dhermes/instrumented-go/addition"
)

func TestAddInt(t *testing.T) {
	got := addition.AddInt(11, 111)
	want := 122
	if got != want {
		t.Errorf("AddInt(11, 111) = %d; want %d", got, want)
	}
}
