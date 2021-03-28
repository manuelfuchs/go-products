create-script = sql/create.sql
drop-script = sql/drop.sql
start-sql-server = brew services restart postgresql
stop-sql-server = brew services stop postgresql

build:
	@echo "todo"

.all: build sql-up sql-down sql-start sql-stop sql-setup sql-teardown
.PHONY: .all 

sql-up: sql-start sql-setup
sql-down: sql-teardown sql-stop

sql-start:
	@echo "Starting SQL-server"
	@${start-sql-server} > /dev/null
	@sleep 2

sql-stop:
	@echo "Stopping SQL-server"
	@${stop-sql-server} > /dev/null

sql-setup: sql-start
	@echo "Creating schema"
	@cat ${create-script} | psql postgres > /dev/null

sql-teardown:
	@echo "Cleaning database"
	@cat ${drop-script} | psql postgres > /dev/null