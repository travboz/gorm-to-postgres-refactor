.PHONY: compose-up compose-down run test

include .env

ENTRYPOINT_DIR=./cmd/api
MIGRATIONS_PATH=./migrations

run:
	@go run $(ENTRYPOINT_DIR)/*.go

setup: compose-up migrate/up

quick-run: setup run

# docker commands

.PHONY: compose-up
compose-up:	
	@echo "Starting containers..."
	docker compose up -d

.PHONY: compose-down
compose-down:
	@echo "Stopping containers..."
	@docker compose down -v

# postgres connect

.PHONY: connect-pg connect-pg-long
connect-pg:
	@psql $(QUEST_DB_DSN)
# password for above is: pa55word

echo-dsn:
	@echo $(QUEST_DB_DSN)


# Migrations
.PHONY: migrate/create migrate/up migrate/down migrate/version migrate/force

migrate/create:
	@migrate create -seq -ext=.sql -dir=$(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate/up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(QUEST_DB_DSN) up

migrate/down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(QUEST_DB_DSN) down $(filter-out $@,$(MAKECMDGOALS))

# e.g. migrate/goto/version 1 -> rolls back to migration 1
migrate/goto/version:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(QUEST_DB_DSN) goto $(filter-out $@,$(MAKECMDGOALS))

migrate/version:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(QUEST_DB_DSN) version

# Used for cleaning a dirty database.
# 1st: Manually roll back partial changes to DB - i.e. fix errors in migration in question. 
# 2nd: Run the below rule with the DB version you want. # eg: migrate -path=./migrations -database=$EXAMPLE_DSN force 1
migrate/force:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(QUEST_DB_DSN) force $(filter-out $@,$(MAKECMDGOALS))
