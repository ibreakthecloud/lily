.PHONY: build run test clean

build:
	docker build -t atlan-lily:latest .

run:
	docker run -p 8080:8080 your-app

test:
	go test ./...

clean:
	docker rmi your-app