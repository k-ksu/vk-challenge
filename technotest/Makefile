include .env

ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP := user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=5432 sslmode=disable
endif

MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down

.PHONY: install-goose
install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: db-up
db-up:
	docker-compose up -d
	make install-goose
	make migration-up

.PHONY: run
run:
	make db-up
	brew install http-server
	http-server -p 8080 &
	go run $(CURDIR)/cmd/technotest &
	go run $(CURDIR)/cmd/pinger-cron
