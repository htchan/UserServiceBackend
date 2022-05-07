.PHONY: backend frontend test controller

service ?= all

## help: show available command and description
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed  -e 's/^/ /'

## test: deploy test container
test:
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile test up --build
	rm ./backend/go.*

## build service=<service>: build docker image of specified service (default all)
build:
	cp go.mod ./backend
	cp go.sum ./backend
	DOCKER_BUILDKIT=1 docker-compose --profile ${service} build
	rm ./backend/go.*

## backend: deploy backend container
backend:
	docker-compose --profile backend up -d

## frontend: compile flutter frontend
frontend:
	docker-compose --profile frontend up

## controller: deploy controller container
controller:
	docker-compose --profile controller up

local_build:
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile local build
	rm ./backend/go.*

local_frontend:
	docker-compose --profile local-frontend up

local_backend:
	docker-compose --profile local-backend up -d

clean:
	rm controller
