include Makefile.ledger
all: lint install

install: go.sum
		go install $(BUILD_FLAGS) ./cmd/konstellation
		test -f ~/go/bin/kd && echo "kd exists" || ln -s ~/go/bin/konstellation ~/go/bin/kd
		go install $(BUILD_FLAGS) ./cmd/konstellationcli
		test -f ~/go/bin/kcli && echo "kcli exists" || ln -s ~/go/bin/konstellationcli ~/go/bin/kcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify