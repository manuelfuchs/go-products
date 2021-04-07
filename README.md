# Go Products

This was written with the instructions of [this post](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql).

## Setup

Execute the following commands:
* `make sql-up`
    * This command starts a postgresql-server docker container, creates a new role and products table.
* `make run`
    * Starts the webserver