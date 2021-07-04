ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	if [ `docker network ls | grep shared-local |wc -l` -eq 0 ]; then docker network create shared-local; fi
	$(ENV_LOCAL) docker-compose up


.PHONY: migrate
migrate:
	$(ENV_LOCAL) sh ./infrastructure/mysql/schemas/migrator.sh

.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --enable=gosec,prealloc,gocognit

.PHONY: test
test:
	$(ENV_LOCAL) go test ./test/... -count=1

.PHONY: generate
generate:
	rm -rf mock/
	go generate ./...
