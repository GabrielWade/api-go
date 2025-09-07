include .env
export

MIGRATE_URL=postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE_PATH=database/migration

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(MIGRATE_URL)" -verbose up

migrate-create:
	migrate create -ext sql -dir $(MIGRATE_PATH) -seq init_mg

create-user:
	