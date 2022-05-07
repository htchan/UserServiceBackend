## help: show available command and description
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed  -e 's/^/ /'

## compile_grpc: compile the protobuf to grpc go file
compile_grpc: ../assets/user_service.proto
	export PATH=${PATH}:${HOME}/go/bin
	protoc --go_out=../internal/grpc --go_opt=paths=source_relative \
		--go-grpc_out=../internal/grpc --go-grpc_opt=paths=source_relative \
		--proto_path=../assets/ user_service.proto

## apply_db_schema: construct ./assets/template.db by sql writern in ./assets/schema/schema.sql
apply_db_schema: ../assets/schema/schema.sql
	rm -f ../assets/template.db
	cat ../assets/schema/schema.sql | sqlite3 ../assets/template.db

## generate_test_data: construct ./assets/template.db by sql writern in ./assets/schema/schema.sql
generate_test_data: ../assets/schema/schema.sql ../assets/schema/test_data.sql
	rm -f ../assets/test_data.db
	cat ../assets/schema/schema.sql ../assets/schema/test_data.sql  | sqlite3 ../assets/test_data.db

## test: test packages and show coverage
test:
	clear
	# go get github.com/htchan/UserService/backend/${package}
	go test github.com/htchan/UserService/backend/pkg/token -coverprofile ./profile.out
	go tool cover -html=profile.out -o coverage.html

poc_orm:
	clear
	cd ../; go build ./POC/orm/sqlite_main.go; ./sqlite_main
	sleep 10
	cd ../; go build ./POC/orm/gorm_main.go; ./gorm_main
	ls -l ../gorm_main ../sqlite_main
	rm ../sqlite_main ../gorm_main

build_server: ../pkg ../cmd/server ../internal
	go build github.com/htchan/UserService/backend/cmd/server
	mv server ../../main

build_controller: ../pkg ../cmd/controller ../internal
	go build github.com/htchan/UserService/backend/cmd/controller
	mv controller ../../controller

build: build_server build_controller