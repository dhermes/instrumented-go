package addition

import (
	"github.com/dhermes/instrumented-go/indirect"
)

// AddInt adds two integers together.
func AddInt(a, b int) int {
	return a + indirect.IdentityInt(b)
}
