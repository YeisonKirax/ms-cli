
start-dev:
	air

start-prod:
	./main

build:
	go build -o ./main .

linter:
	golangci-lint run

format:
	gofmt -w .

install:
	go mod tidy

test-cov:
	go test -coverprofile=coverage.out ./src/modules/.../domain/... ./src/modules/.../application/... && go tool cover
	-html=coverage.out

test:
	go test -v ./pkg/modules/.../domain/... ./pkg/modules/.../application/...

gen-mocks:
	mockery --all