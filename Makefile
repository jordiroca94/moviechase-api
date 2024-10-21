build:
	go build -o bin/user-auth-api cmd/main.go

test:
	go test -v ./...

run: build 
	./bin/user-auth-api