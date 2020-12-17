package main

import (
	"fmt"
	"os"

	"github.com/dhermes/instrumented-go/addition"
	"github.com/dhermes/instrumented-go/multiplication"
)

var (
	stdout *os.File = os.Stdout
	stderr *os.File = os.Stderr
)

func run() error {
	sum := addition.AddFloat64(2.5, 39.5)
	fmt.Fprintf(stdout, "2.5 + 39.5 = %f\n", sum)
	product := multiplication.MultiplyInt(3, 14)
	fmt.Fprintf(stdout, "3 × 14 = %d\n", product)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		os.Exit(1)
	}
}
