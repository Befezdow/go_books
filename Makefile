.PHONY: build
build_gateway:
	go build -o main -v ./cmd/gateway

run_gateway:
	go run ./cmd/gateway/main.go config_path=configs/gateway-config.toml
