### Intro
Here is my helper for Golang master classes on hexlet.io platform
https://hexlet.io/u/whatafunc
Automatic code reviews - paased example: https://hexlet.io/code_reviews/1769592 // for
https://hexlet.io/code_reviews/1771898 //sort slices

### Linters
1. downloaded & executed
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.2
golangci/golangci-lint info checking GitHub for tag 'v1.62.2'
golangci/golangci-lint info found version: 1.62.2 for v1.62.2/darwin/amd64
golangci/golangci-lint info installed /Users/Dmitriy/go/bin/golangci-lint

2. checking
/Users/Dmitriy/go/bin/golangci-lint version
golangci-lint has version 1.62.2 built with go1.23.3 from 89476e7a on 2024-11-25T14:16:01Z

3. running
~/go/bin/golangci-lint run ./cmd/

### Tests
# Test a specific package:
go test ./cmd/...

# Test all packages recursively:
go test ./...

# Verbose output (show details):
go test -v ./cmd/for/...

# Run a specific test:
go test -v ./cmd/for/... -run TestMap
go test -v ./cmd/... -run TestUniqueSortedUserIDs