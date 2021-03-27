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

mysql:
	@docker exec -ti mysql mysql --user=dev --password=dev --database=dev
