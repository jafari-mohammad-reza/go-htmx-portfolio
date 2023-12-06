build:
	@go build -o ./dist/api main.go
run: build
	@./dist/api
dev:
	@~/go/bin/reflex -s -r '\.go' -R '^vendor/.' -R '^_.*' go run main.go
docker-dev:
	@docker rm -f portfolio || true && docker build -t portfolio . && docker run --name portfolio -v $$(pwd):/app -p 5050:5050  portfolio