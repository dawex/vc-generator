# Docker Compose Up
docker_compose_up:
	docker compose up -d

# Docker Compose Up
docker_compose_down:
	docker compose down

# Install all Go dependencies
install:
	go mod download
	go install tool

# Generate OAS bundled document in /dist
generate_api_interfaces:
	oapi-codegen -version
	go generate -x $$(go list ./... | grep -v '/ports\|/adapters|/core' | tr '\n' ' ')

# Build binary
build:
	go build -o cmd/vc-generator cmd/main.go

run: generate_api_interfaces
	go run cmd/main.go

update_deps:
	go get -u ./...
	go mod tidy

test_go:
	ginkgo -v -r -cover -coverpkg=APIs/internal/... -coverprofile=coverage.txt ./...