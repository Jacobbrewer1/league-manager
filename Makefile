# Define variables
hash = $(shell git rev-parse --short HEAD)
DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

pr-approval:
	@echo "Running PR CI"
	go build ./...
	go vet ./...
	go test ./...
codegen: deps
	@echo "Generating code"
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

	go generate ./...
deps:
	go install github.com/charmbracelet/gum@latest
	go install golang.org/x/tools/cmd/goimports@latest
