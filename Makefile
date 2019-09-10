PACKAGE := github.com/quantumghost/borg-tm
COMMIT_ID := master
VERSION := master
BUILD_ARGS := -ldflags "-X $(PACKAGE)/consts.version=$(VERSION) -X $(PACKAGE)/consts.commitID=$(COMMIT_ID)"
EXTRA_BUILD_ARGS =
OUTPUT_FILE := out/borg-tm
MAIN_FILE := cmd/main.go

.PHONY: fotmat build check-style lint check-error build-image

format:
	@find $(CURDIR) -mindepth 1 -maxdepth 1 -type d -not -name vendor -not -name .git -print0 | xargs -0 gofmt -s -w
	@find $(CURDIR) -maxdepth 1 -type f -name '*.go' -print0 | xargs -0 gofmt -s -w

build:
	go build $(BUILD_ARGS) $(EXTRA_BUILD_ARGS) -o $(OUTPUT_FILE) $(MAIN_FILE)

check-style:
	@golangci-lint run --disable-all -E gofmt ./...

lint:
	@golangci-lint run ./...

check-error:
	@golangci-lint run --disable-all -E errcheck ./...
