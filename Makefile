dev:
	go run main.go
test:
	go test ./... -coverprofile out
lint: 
	golangci-lint run -v