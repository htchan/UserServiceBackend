test: ./go.mod ./go.sum ./backend
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile test up --build
	rm ./backend/go.*

build: ./go.mod ./go.sum ./backend
	cp go.mod ./backend
	cp go.sum ./backend
	docker-compose --profile all build
	rm ./backend/go.*

server: ./go.mod ./go.sum ./backend
	echo "to be finish"

controller: ./go.mod ./go.sum ./backend
	echo "to be finish"