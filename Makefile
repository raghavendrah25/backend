run:
	@echo "Run Triggered"
	go run main.go

build:
	@echo "Building for Linux"
	GOOS=linux GOARCH=amd64 go build -o bin/ main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin main.go
