BIN_DIR=./bin

help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## deps: get dependency
deps:
	@cp .hooks/* .git/hooks
	go build ./...
	go test ./...

## godoc: run godoc.
# maybe you need install godoc. $ go install golang.org/x/tools/cmd/godoc@latest
godoc:
	@echo "http://localhost:6060"
	godoc -http=:6060

## lint: run golangci-lint.
# maybe you need install golangci-lint. $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint:
	golangci-lint run ./...

## test: run test. not cache
test:
	go test -count=1 -race -coverprofile=coverage.out -covermode=atomic ./...

## view_test: view test coverage
view_test:
	go tool cover -html=coverage.out

## all_test: run all test(include bench). not cache
all_test:
	go test -count=1 -bench=. -v ./...

## all_generate: run all go generate
all_generate:
	go generate ./...

## mod_tidy: go mod tidy
mod_tidy:
	go mod tidy
