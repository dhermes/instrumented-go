# instrumented-go

> A Small Experiment Showcasing a `main` Binary Instrumented for Coverage

## Usage

```
$ make
Makefile for `instrumented-go` project

Usage:
   make instrumented-main    Make a custom `main` binary that is instrumented for coverage
   make cover.out            Generate a cover profile for code paths invoked by `cmd/main.go`
   make coverage.html        Generate HTML for the `cmd/main.go` cover profile

```

## References

-   [Accurate Code Coverage in Go][1]
-   [The cover story][2]
-   [Measuring Code Coverage of Golang Binaries with Bincover][3]
-   `bincover` [repository][4]

[1]: https://www.ory.sh/golang-go-code-coverage-accurate/
[2]: https://blog.golang.org/cover
[3]: https://www.confluent.io/blog/measure-go-code-coverage-with-bincover/
[4]: https://github.com/confluentinc/bincover
