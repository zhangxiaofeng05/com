BIN_DIR=./bin

help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## godoc: run godoc.
# maybe you need install godoc. $ go install golang.org/x/tools/cmd/godoc@latest
godoc:
	@echo "http://localhost:6060"
	godoc -http=:6060

## lint: run golangci-lint.
# maybe you need install golangci-lint. $ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint:
	golangci-lint run ./...

## test: run test. view result:$ go tool cover -html=coverage.txt
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic ./...

## all_test: run test(include bench)
all_test:
	go test -bench=. -v ./...

## mod_tidy: go mod tidy
mod_tidy:
	go mod tidy
