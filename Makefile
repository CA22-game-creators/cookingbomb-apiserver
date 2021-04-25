.PHONY: run
run:
	docker compose -f docker/docker-compose.yml up

.PHONY: re-run
re-run:
	docker compose -f docker/docker-compose.yml up --build
