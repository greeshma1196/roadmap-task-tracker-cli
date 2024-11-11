build:
	go build -o task-cli
test:
	go test ./... -coverprofile cover.out -v
lint: 
	golangci-lint run -v