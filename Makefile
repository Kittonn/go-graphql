gql-generate:
	go run github.com/99designs/gqlgen generate

run:
	go run ./cmd/api/main.go

build:
	go build -o main ./cmd/api/main.go