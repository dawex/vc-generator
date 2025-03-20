# Validation Plugin - VC Generator

VC Generator distribution for UPCAST European project.

## OpenAPI documentation

[Visit Documentation hosted on GitHub pages](https://dawex.github.io/vc-generator/)

Build & Run the documentation:
```bash
cd documentation/
yarn i
yarn oas:bundle
```

## Installation

### Prerequisites

- Go (v1.24 or later)
- Docker

### Steps

1. Run database with docker compose:
    ```bash
    make docker_compose_up
    ```

Create DB with dbname=vc-generator

2. Install Go modules:
    ```bash
    make install
    ```

3. Generate API interface from oas documentation:
    ```bash
    make generate_api_interfaces
    ```

4. Run service:
    ```bash
    make run
    ```

5. Build binarie:
    ```bash
    make build
    ```

6. Run test:
    ```bash
    make test_go
    ```
