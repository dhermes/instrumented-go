package multiplication

import (
	"github.com/dhermes/instrumented-go/indirect"
)

// MultiplyInt multiplies two integers together.
func MultiplyInt(a, b int) int {
	return a * indirect.IdentityInt(b)
}
