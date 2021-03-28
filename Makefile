create-table-script = sql/create_table.sql
create-role-script = sql/create_role.sql
drop-table-script = sql/drop_table.sql
drop-role-script = sql/drop_role.sql

sql-server = postgres

start-sql-server = docker container run --name ${sql-server} -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
stop-sql-server = docker container stop ${sql-server}
rm-sql-server = docker container rm ${sql-server}
psql-command = docker container exec ${sql-server} psql -U postgres

build:
	@echo "todo"

.all: build setup sql-up sql-down sql-start sql-stop sql-setup
.PHONY: .all 

setup: sql-up

brew:
	@brew bundle

sql-up: sql-start sql-setup

sql-down:
	@echo "Stopping SQL-server"
	@${stop-sql-server} > /dev/null
	@echo "Removing SQL-server"
	@${rm-sql-server} > /dev/null

sql-start:
	@echo "Starting SQL-server"
	@${start-sql-server}
	@sleep 2

sql-setup: sql-start
	@echo "Creating role"
	@cat ${create-role-script} | ${psql-command} > /dev/null
	@echo "Creating table"
	@cat ${create-table-script} | ${psql-command} > /dev/null