create_table_script = `cat sql/create_table.sql`
create_role_script = `cat sql/create_role.sql`
drop_table_script = `cat sql/drop_table.sql`
drop_role_script = `cat sql/drop_role.sql`
source_folder = ./app/cmd/
db_folder = ./app/db/

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

sql_server = postgres
sql_server_image = manueltfuchs/go-products-db:latest

build_sql_server_image = ${docker_container_build} -t ${sql_server_image} .
start_sql_server = ${docker_container_run} --name ${sql_server} -p 5432:5432 -d ${sql_server_image}
stop_sql_server = ${docker_container_stop} ${sql_server}
rm_sql_server = ${docker_container_rm} ${sql_server}

run:
	@${go_run} ${source_folder}/...

build:
	@mkdir ${bin_folder} > /dev/null
	@${go_build} -o ${bin_folder}/${build_artifact_name} ${source_folder}/...

test:
	@${go_test} ${source_folder}/...

clean:
	@echo "Removing build artifact"
	@rm -r ${bin_folder}

.PHONY: run test sql-up sql-down sql-build sql-setup

sql-build:
	@echo "Building SQL-Server image"
	@cd ${db_folder} && ${build_sql_server_image} > /dev/null

sql-up:
	@echo "Starting SQL-server"
	@${start_sql_server} > /dev/null

sql-down:
	@echo "Stopping SQL-server"
	@${stop_sql_server} > /dev/null
	@echo "Removing SQL_server"
	@${rm_sql_server} > /dev/null