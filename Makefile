run:
	go run cmd/main.go

docker-build:
	docker build -t fizz-buzz-api .

docker-run:
	docker run -p 8080:8080 --env-file .env fizzbuzz-api