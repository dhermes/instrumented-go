.PHONY: help
help:
	@echo 'Makefile for `instrumented-go` project'
	@echo ''
	@echo 'Usage:'
	@echo '   make instrumented-main    Make a custom `main` binary that is instrumented for coverage'
	@echo '   make cover.out            Generate a cover profile for code paths invoked by `cmd/main.go`'
	@echo '   make coverage.html        Generate HTML for the `cmd/main.go` cover profile'
	@echo '   make clean                Remove all generated files'
	@echo ''

instrumented-main: addition/float.go addition/integer.go cmd/instrumented_test.go cmd/main.go indirect/identity.go multiplication/float.go multiplication/integer.go
	go test \
	  -c \
	  -o ./instrumented-main \
	  -covermode atomic \
	  -coverpkg=./... \
	  -tags instrumented_main \
	  ./cmd

cover.out: instrumented-main
	./instrumented-main -test.coverprofile ./cover.out

coverage.html: cover.out
	go tool cover -html=./cover.out -o ./coverage.html

.PHONY: clean
clean:
	rm -f instrumented-main cover.out coverage.html
