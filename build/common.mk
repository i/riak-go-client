.PHONY: all install-deps lint unit-test integration-test test fmt help

all: install-deps lint test

install-deps:
	go get -t github.com/basho/riak-go-client/...

lint: install-deps
	go vet github.com/basho/riak-go-client/...

unit-test: lint
	go test -v github.com/basho/riak-go-client/...

integration-test: lint
	go test -v -tags=integration github.com/basho/riak-go-client/...

test: unit-test integration-test

fmt:
	gofmt -s -w .

help:
	@echo ''
	@echo ' Targets:'
	@echo '--------------------------------------------------'
	@echo ' all              - Run everything                '
	@echo ' fmt              - Format code                   '
	@echo ' lint             - Run jshint                    '
	@echo ' test             - Run unit & integration tests  '
	@echo ' unit-test        - Run unit tests                '
	@echo ' integration-test - Run integration tests         '
	@echo '--------------------------------------------------'
	@echo ''
