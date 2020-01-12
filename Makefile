.PHONY: build-docker

include .env

# TODO: figure this out with generating versions from git tags?
VERSION=0.1

build:
	go build -v .

build-docker:
	docker build --tag gogoapp:$(VERSION) .

# TODO remove --rm
run-docker:
	docker run --rm gogoapp:$(VERSION)

compose-up:
	docker-compose up

compose-up-build:
	docker-compose up --force-recreate --build

compose-up-d:
	docker-compose up -d

compose-down:
	docker-compose down

compose-down-purge:
	docker-compose down --rmi local --volumes