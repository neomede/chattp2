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
.PHONY: test

test-ci:
	echo "mode: count" > c.out
	$(foreach pkg,$(PACKAGES),\
		GORACE="halt_on_error=1" go test -v -race -cover -coverprofile=coverage.out $(pkg) || exit 1;\
		tail -n +2 coverage.out >> c.out;)
.PHONY: test-ci

lint:
	$(FGT) go fmt ./...
	$(FGT) go list ./... | grep -v vendor/ | xargs -L1 $(FGT) golint
	$(FGT) go vet ./...
	$(FGT) errcheck -ignore Close ./...
	$(FGT) staticcheck ./...
	$(FGT) interfacer ./...
	$(FGT) gosimple ./...
.PHONY: lint

ci-check: test-ci lint
.PHONY: ci-check

# Install tools
install-tools:
	go get -u github.com/GeertJohan/fgt
	go get -u golang.org/x/tools/cmd/cover
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/golang/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u mvdan.cc/interfacer
	go get -u honnef.co/go/tools/cmd/staticcheck
.PHONY: install-tools
