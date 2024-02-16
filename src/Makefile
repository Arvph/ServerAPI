.PHONY: all m_up m_down

all:
	go build -v -o ../build ./cmd/api 

m_up:
	migrate -path migrations -database "postgres://localhost:5434/restapi?sslmode=disable&user=postgres&password=admin" up

m_down:
	migrate -path migrations -database "postgres://localhost:5434/restapi?sslmode=disable&user=postgres&password=admin" down