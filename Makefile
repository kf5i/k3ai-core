# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

build-cli:
	go build -o bin/k3ai-cli

.PHONY: lint
lint: check-format
	go vet ./...
	# Add addiotional linters here

.PHONY: check-format
check-format:
	@echo "Running gofmt..."
	$(eval unformatted=$(shell find . -name '*.go' | grep -v ./.git | grep -v vendor | xargs gofmt -l))
	$(if $(strip $(unformatted)),\
		$(error $(\n) Some files are ill formatted! Run: \
			$(foreach file,$(unformatted),$(\n)    gofmt -w $(file))$(\n)),\
		@echo All files are well formatted.\
	)
.PHONY: test
test:
	go test -coverprofile=coverage.txt -covermode=atomic -race ./...
