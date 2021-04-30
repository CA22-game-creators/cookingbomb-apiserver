ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	$(ENV_LOCAL) docker-compose up

.PHONY: re-run
re-run:
	$(ENV_LOCAL) docker-compose up --build

.PHONY: migration
migration:
	$(ENV_LOCAL) sh ./infrastructure/mysql/schemas/migrator.sh
