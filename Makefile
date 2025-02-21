#Makefile для создания миграций

#Переменные которые будут использоваться в наших командах (Таргетах)
DB_DSN := "postgres://postgres:bebra@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

#Таргет для создания новой миграции
migrate-force:
	$(MIGRATE) force ${VERSION}

version:
	$(MIGRATE) version

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

#Применение миграций
migrate:
	$(MIGRATE) up

#Откат миграций
migrate-down:
	$(MIGRATE) down

migrate-delete:
	migrate-down 1
	rm -f ./migrations/${NAME}.up.sql ./migrations/${NAME}.down.sql

run:
	go run cmd/app/main.go

gen-tasks:
	oapi-codegen -config openapi/.openapi.tasks -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

gen-users:
	oapi-codegen -config openapi/.openapi.users -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint cache clean
	golangci-lint run --timeout=5m --out-format=colored-line-number
