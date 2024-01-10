build:
	@go build -o ./dist/api main.go
run: build
	@./dist/api
dev:
	@/go/bin/reflex -s -r '\.go' -R '^vendor/.' -R '^_.*' go run main.go
docker-dev:
	@docker rm -f authentication-service || true && docker build -t authentication-service . && docker run --name authentication-service -v $$(pwd):/app -p 5050:5050  authentication-service
