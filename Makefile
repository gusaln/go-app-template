all: bin/migrate

bin/migrate: cmd/migrate/* schemas/migrations/*.sql
	go build -o bin/migrate cmd/migrate/*

.PHONY : sqlc
sqlc: schemas/queries/*.sql
	go tool sqlc generate