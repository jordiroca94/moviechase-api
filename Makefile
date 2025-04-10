build:
	go build -o bin/moviechase-api main.go

test:
	go test -v ./...

run: build 
	./bin/moviechase-api

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $$@,$(MAKECMDGOALS))

add-user-table-up:
	@go run cmd/migrate/main.go up

add-user-table-down:
	@go run cmd/migrate/main.go down