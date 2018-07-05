all: build

build:
	go fmt ./pkg/...
	go vet ./pkg/...
	go install -v ./pkg/...

test:
	ginkgo ./pkg/...

dep:
	dep ensure -v

clean-dep:
	rm -f ./Gopkg.lock
	rm -rf ./vendor

.PHONY: build test dep clean-dep
