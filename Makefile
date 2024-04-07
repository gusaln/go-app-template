all: bin/migrate

bin/migrate: schemas/main.go schemas/migrations/*.sql
	go build -o bin/migrate schemas/main.go