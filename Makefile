db_folder = ./app/db/
backend_folder = ./app/backend/
backend_source_folder = ./app/backend/cmd/

bin_folder = bin
build_artifact_name = products_server

go_cmd = go
go_run = ${go_cmd} run
go_test = ${go_cmd} test -v
go_build = ${go_cmd} build

docker = docker
docker_container = ${docker} container
docker_container_build = ${docker} build
docker_container_run = ${docker_container} run
docker_container_stop = ${docker_container} stop
docker_container_rm = ${docker_container} rm

backend_container = go-products-backend
backend_image = manueltfuchs/${backend_container}:latest
build_backend = ${docker_container_build} -t ${backend_image} .
start_backend = ${docker_container_run} --name ${backend_container} -p 8080:80 -d ${backend_image}
stop_backend = ${docker_container_stop} ${backend_container}
rm_backend = ${docker_container_rm} ${backend_container}

db_container = go-products-db
db_image = manueltfuchs/${db_container}:latest
build_db = ${docker_container_build} -t ${db_image} .
start_db = ${docker_container_run} --name ${db_container} -p 5432:5432 -d ${db_image}
stop_db = ${docker_container_stop} ${db_container}
rm_db = ${docker_container_rm} ${db_container}

default: start
.PHONY: run test sql-up sql-down sql-build sql-setup

build:
	@echo "Building backend image"
	@cd ${backend_folder} && ${build_backend} > /dev/null

start:
	@echo "Starting backend"
	@${start_backend} > /dev/null

stop:
	@echo "Stopping backend"
	@${stop_backend} > /dev/null
	@echo "Removing backend"
	@${rm_backend} > /dev/null

db-build:
	@echo "Building database image"
	@cd ${db_folder} && ${build_db} > /dev/null

db-start:
	@echo "Starting database"
	@${start_db} > /dev/null

db-stop:
	@echo "Stopping database"
	@${stop_db} > /dev/null
	@echo "Removing database"
	@${rm_db} > /dev/null