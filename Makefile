dev:
	go run main.go
test:
	go test ./... -coverprofile cover.out -v
lint: 
	golangci-lint run -v