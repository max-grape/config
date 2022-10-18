WORKDIR         = /usr/src/max-grape/config
IMAGE_GOLANG_CI = golangci/golangci-lint:v1.49-alpine
IMAGE_GOLANG    = golang:1.19.1

lint:
	@docker run --rm -t -w $(WORKDIR) -v $(CURDIR):$(WORKDIR) $(IMAGE_GOLANG_CI) golangci-lint run

unit:
	@docker run --rm -w $(WORKDIR) -v $(CURDIR):$(WORKDIR) $(IMAGE_GOLANG) go test -v ./...

test: lint unit
