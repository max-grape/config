CWD            = /go/src/github.com/max-grape/config
GOLANG_CI_IMG  = golangci/golangci-lint:v1.42-alpine
GOLANG_IMG     = golang:1.17.1

lint:
	@docker run --rm -t -w $(CWD) -v $(CURDIR):$(CWD) -e GOFLAGS=-mod=vendor $(GOLANG_CI_IMG) golangci-lint run

unit:
	@docker run --rm -w $(CWD) -v $(CURDIR):$(CWD) -e GOFLAGS=-mod=vendor $(GOLANG_IMG) go test -v ./...

test: lint unit
