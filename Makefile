COVERAGE_FILE ?= coverage.out

TARGET ?= hangman # CHANGE THIS TO YOUR BINARY NAME/NAMES

.PHONY: build
build:
	@echo "Выполняется go build для таргета ${TARGET}"
	@mkdir -p .bin
	@go build -o ./bin/${TARGET} ./cmd/${TARGET}

## test: hangman all tests
.PHONY: test
test:
	@go test -coverpkg='github.com/k6zma/cli-hangman/...' --race -count=1 -coverprofile='$(COVERAGE_FILE)' ./...
	@go tool cover -func='$(COVERAGE_FILE)' | grep ^total | tr -s '\t'
