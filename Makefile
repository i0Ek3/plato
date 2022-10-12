.PHONY: install test clean tool

GO=go

all: test tool

install:
	@$(GO) get github.com/i0Ek3/plato

test:
	@$(GO) test -v .
	@$(GO) test -bench .

tool:
	gofmt -w .

clean:
	rm plato
