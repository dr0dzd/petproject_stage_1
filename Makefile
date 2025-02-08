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

run:
	go run cmd/app/main.go