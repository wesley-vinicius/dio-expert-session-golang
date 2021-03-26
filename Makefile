build:
	docker build -t wesley-vinicius/finance .

up:
	docker-compose up -d

down:
	docker-compose down

test:
	go test ./...

lint:
	golint ./...
	go fmt -n ./...

code_format:
	go fmt ./...