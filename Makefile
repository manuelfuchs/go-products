create-table-script = `cat sql/create_table.sql`
create-role-script = `cat sql/create_role.sql`
drop-table-script = `cat sql/drop_table.sql`
drop-role-script = `cat sql/drop_role.sql`

sql-server = postgres

docker = docker container
start-sql-server = ${docker} run --name ${sql-server} -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
stop-sql-server = ${docker} stop ${sql-server}
rm-sql-server = ${docker} rm ${sql-server}
execute-on-sql-server = ${docker} exec ${sql-server} 
psql-command = psql -U postgres

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
	@${start-sql-server} > /dev/null
	@sleep 2

sql-setup: sql-start
	@echo "Creating role"
	@${execute-on-sql-server} ${psql-command} -c "${create-role-script}" > /dev/null
	@echo "Creating table"
	@${execute-on-sql-server} ${psql-command} -c "${create-table-script}" > /dev/null