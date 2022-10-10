.PHONY: build install test clean tool

GO=go

all: test tool

build:
	@$(GO) build -o ./plato ./cmd/plato/main.go

install:
	@$(GO) get github.com/i0Ek3/plato

test:
	@$(GO) test -v .
	@$(GO) test -bench .

tool:
	gofmt -w .

clean:
	rm plato