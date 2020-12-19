package indirect_test

import (
	"testing"

	"github.com/dhermes/instrumented-go/indirect"
)

func TestIdentityInt(t *testing.T) {
	got := indirect.IdentityInt(3)
	want := 3
	if got != want {
		t.Errorf("IdentityInt(3) = %d; want %d", got, want)
	}
}

func TestIdentityInt64(t *testing.T) {
	got := indirect.IdentityInt64(13)
	want := int64(13)
	if got != want {
		t.Errorf("IdentityInt64(13) = %d; want %d", got, want)
	}
}

func TestIdentityFloat32(t *testing.T) {
	got := indirect.IdentityFloat32(11.75)
	want := float32(11.75)
	if got != want {
		t.Errorf("IdentityFloat32(11.75) = %f; want %f", got, want)
	}
}

func TestIdentityFloat64(t *testing.T) {
	got := indirect.IdentityFloat64(117.5)
	want := 117.5
	if got != want {
		t.Errorf("IdentityFloat64(117.5) = %f; want %f", got, want)
	}
}
