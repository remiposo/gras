GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: build
build:
	docker compose build --no-cache

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down

.PHONY: logs
logs:
	docker compose logs -f

.PHONY: ps
ps:
	docker compose ps

.PHONY: test
test:
	go test -race -shuffle=on ./...

.PHONY: migrate
migrate: $(GOBIN)/mysqldef
	mysqldef -u gras -p gras -h 127.0.0.1 -P 33306 gras < ./_tools/mysql/schema.sql

$(GOBIN)/mysqldef:
	go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
