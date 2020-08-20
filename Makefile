all: build

build:
	go fmt ./pkg/...
	go vet ./pkg/...
	go install -v ./pkg/...
	go mod tidy

test:
	ginkgo ./pkg/...

dep:
	go get -t -u ./...
	go mod vendor

.PHONY: build test dep