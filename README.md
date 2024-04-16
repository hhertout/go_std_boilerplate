# Go API Boilerplate - Std lib

## Getting Started

## Docker

### Build the image
```bash
docker build -t go-api-boilerplate .
```

### Run the container
```bash
docker run -p 8080:8080 go-api-boilerplate
```

### Run the container with environment variables
```bash
docker run -p 8080:8080 -e PORT=8080 go-api-boilerplate
```

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```
