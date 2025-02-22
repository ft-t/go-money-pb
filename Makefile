.PHONY: build
build:
	docker compose up --force-recreate

.PHONY: generate
generate: build

.PHONY: update
update:
	go get -u ./...
