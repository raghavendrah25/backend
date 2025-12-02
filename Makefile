run:
	@echo "Run Triggered"
	go run api/main.go

build:
	@echo "Building for Linux"
	GOOS=linux GOARCH=amd64 go build -o bin/ api/main.go
	GOOS=linux GOARCH=amd64 go build -o bootstrap testLambda/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin api/main.go
