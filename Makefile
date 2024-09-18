dbcontainer:
	docker run -p 5432:5432 --name course -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pass -d postgres:16.4-alpine3.20

startdb:
	docker run course

stopdb:
	docker stop course

createdb:
	docker exec -it course createdb --username=root --owner=root courses

dropdb:
	docker exec -it course dropdb  -U root courses

migrateup:
	migrate -path db/migration/ -database "postgresql://root:pass@localhost:5432/courses?sslmode=disable" -verbose  up

migratedown:
	migrate -path db/migration/ -database "postgresql://root:pass@localhost:5432/courses?sslmode=disable" -verbose  down

sqlc:
	sqlc generate

test:
	go test -v --cover ./...

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

server:
	go run main.go

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: dbcontainer startdb stopdb createdb migrateup migratedown sqlc test proto server evans
