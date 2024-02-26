.PHONY: all

all:
	go build -v ./cmd/api

m_up:
	migrate -path migrations -database "postgres://localhost:8080/restapi?sslmode=disable&user=postgres&password=admin" up

m_down:
	migrate -path migrations -database "postgres://localhost:8080/restapi?sslmode=disable&user=postgres&password=admin" down