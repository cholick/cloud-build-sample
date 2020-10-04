default: all

GO_FILES = $$(find . -name "*.go" | grep -v vendor | uniq)

bootstrap:
	go get golang.org/x/tools/cmd/goimports

fmt:
	goimports -l -w  $(GO_FILES)

test:
	go vet ./...
	go test ./...

run:
	go run cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
	cp main deployment/docker/tmp/main

clean:
	go clean
	rm main
	rm -rf vendor

all: fmt build-linux
