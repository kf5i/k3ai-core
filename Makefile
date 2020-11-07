# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

TEST_FLAGS ?=

build-cli:
	go build -o bin/k3ai-cli

.PHONY: lint
lint: check-format
	go get golang.org/x/lint/golint
	go vet ./...
	golint -set_exit_status=1 ./...

.PHONY: check-format
check-format:
	@echo "Running gofmt..."
	$(eval unformatted=$(shell find . -name '*.go' | grep -v ./.git | grep -v vendor | xargs gofmt -s -l))
	$(if $(strip $(unformatted)),\
		$(error $(\n) Some files are ill formatted! Run: \
			$(foreach file,$(unformatted),$(\n)    gofmt -w -s $(file))$(\n)),\
		@echo All files are well formatted.\
	)

.PHONY: test
test:
	go test $(TEST_FLAGS) -coverprofile=coverage.txt -covermode=atomic -race ./...

integration-test:
	go test -tags integration -coverprofile=coverage.txt -covermode=atomic -race ./...
