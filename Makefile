ENV_LOCAL = $(shell cat env.local)

.PHONY: run
run:
	$(ENV_LOCAL) docker-compose up

.PHONY: re-run
re-run:
	$(ENV_LOCAL) docker-compose up --build

.PHONY: lint
lint:
	go mod tidy
	golangci-lint run --enable=golint,gosec,prealloc,gocognit
