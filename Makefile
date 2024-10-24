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
	chmod +x ./pkg/models/generate.sh
	go generate ./...
deps:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/charmbracelet/gum@latest
	go install golang.org/x/tools/cmd/goimports@latest
models:
	chmod +x ./pkg/models/generate_models.sh
	go generate ./pkg/models
apis:
	go generate ./pkg/codegen/...
