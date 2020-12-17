package main

import (
	"fmt"
	"os"

	"github.com/dhermes/instrumented-go/addition"
	"github.com/dhermes/instrumented-go/multiplication"
)

func run() error {
	sum := addition.AddFloat64(2.5, 39.5)
	fmt.Printf("2.5 + 39.5 = %f\n", sum)
	product := multiplication.MultiplyInt(3, 14)
	fmt.Printf("3 × 14 = %d\n", product)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
