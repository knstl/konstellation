include Makefile.ledger
# The below include contains the tools target.
include contrib/devtools/Makefile
all: lint install

install: go.sum
		go install $(BUILD_FLAGS) ./client/konstellation
		go install $(BUILD_FLAGS) ./client/konstellationcli
		go install $(BUILD_FLAGS) ./client/konstellationlcd

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify