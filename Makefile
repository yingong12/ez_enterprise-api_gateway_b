os?=linux
port?=8686


run: 
	swag init -g http/router.go
	go run main.go bootstrap.go

build:export GOOS=$(os)
build:export GOARCH=amd64
build:
	@echo "building binary for $(GOOS)..."
	go build -o ./api_gateway_b
	@echo "done!"
	