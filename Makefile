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
	go build github.com/htchan/UserService/backend/cmd/controller

local_frontend:
	cd frontend ; flutter run -d chrome

clean:
	rm controller