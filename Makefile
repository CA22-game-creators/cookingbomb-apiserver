ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	$(ENV_LOCAL) docker-compose -f docker/docker-compose.yml up

.PHONY: re-run
re-run:
	$(ENV_LOCAL) docker-compose -f docker/docker-compose.yml up --build
