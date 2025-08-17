cli:
    go run cmd/cli/main.go

build:
    go build -v -o build/app cmd/cli/main.go

clean:
    rm -rf build
    go clean

update:
    go get -u ./...
    go mod tidy

test:
    go clean -testcache
    go mod tidy
    go test -cover ./...

format:
    gofmt -l .

info:
    go vet ./...
    go list ./...
    go version
