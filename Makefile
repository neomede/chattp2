PACKAGES=$(shell go list ./... | grep -v vendor/)
FGT := fgt

build-linux:
	@echo "$(WARN_COLOR)+ $@$(NO_COLOR)"
	GOOS=linux GOARCH=amd64 go build .
.PHONY: build-linux

build:
	@echo "$(WARN_COLOR)+ $@$(NO_COLOR)"
	go build .
.PHONY: build

test:
	go test -v $(PACKAGES)

test-ci:
	GORACE="halt_on_error=1" go test -race -v $(PACKAGES)

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 $(FGT) golint

linters: lint
	$(FGT) go fmt $(PACKAGES)
	$(FGT) go vet $(PACKAGES)
	$(FGT) errcheck -ignore Close $(PACKAGES)

ci-check: test-ci linters
