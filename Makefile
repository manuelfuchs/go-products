create_table_script = `cat sql/create_table.sql`
create_role_script = `cat sql/create_role.sql`
drop_table_script = `cat sql/drop_table.sql`
drop_role_script = `cat sql/drop_role.sql`

go_cmd = go
go_run = ${go_cmd} run
go_test = ${go_cmd} test -v

docker = docker
docker_container = ${docker} container
docker_container_run = ${docker_container} run
docker_container_stop = ${docker_container} stop
docker_container_rm = ${docker_container} rm
docker_container_exec = ${docker_container} exec

sql_server = postgres

start_sql_server = ${docker_container_run} --name ${sql_server} -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
stop_sql_server = ${docker_container_stop} ${sql_server}
rm_sql_server = ${docker_container_rm} ${sql_server}
execute_on_sql_server = ${docker_container_exec} ${sql_server} 
psql_command = psql -U postgres

run:
	@${go_run} ./cmd/...

test:
	@${go_test} ./cmd/...

.PHONY: run test sql_up sql_down sql_start sql_setup

sql_up: sql_start sql_setup

sql_down:
	@echo "Stopping SQL-server"
	@${stop_sql_server} > /dev/null
	@echo "Removing SQL_server"
	@${rm_sql_server} > /dev/null

sql_start:
	@echo "Starting SQL-server"
	@${start_sql_server} > /dev/null
	@sleep 2

sql_setup: sql_start
	@echo "Creating role"
	@${execute_on_sql_server} ${psql_command} -c "${create_role_script}" > /dev/null
	@echo "Creating table"
	@${execute_on_sql_server} ${psql_command} -c "${create_table_script}" > /dev/null