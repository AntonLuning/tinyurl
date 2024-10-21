.PHONY: run
run: build
	@docker compose up -d

.PHONY: build
build:
	@docker build --force-rm -t tinyurl:example ../../.

.PHONY: down
down:
	@docker compose down
