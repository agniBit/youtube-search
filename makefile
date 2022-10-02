run:
	go run cmd/api/main.go

run_worker:
	go run cmd/worker/main.go

migrate:
	go run cmd/migration/migration.go

build:
	go build -o bin/api/main cmd/api/main.go
	go build -o bin/cmd/worker/main cmd/worker/main.go