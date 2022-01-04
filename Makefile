.PHONY: backend frontend test controller

test:
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile test up --build
	rm ./backend/go.*

build:
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile all build
	rm ./backend/go.*

backend:
	docker-compose --profile backend up -d

frontend:
	docker-compose --profile frontend up

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