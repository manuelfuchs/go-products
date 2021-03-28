all: build sql-setup
.PHONY: all

build:
	@echo "todo"

sql-setup:
	@echo "Restarting postgresql-server"
	@brew services restart postgresql
	# todo

sql-clean:
	@echo "Cleaning database"
	@echo "Removing database"
	@echo "Stopping postgresql-server"