# Go Products

[![Build Status](https://travis-ci.com/manuelfuchs/go-products.svg?branch=main)](https://travis-ci.com/manuelfuchs/go-products)

This was written with the instructions of [this post](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql).

## Setup

Execute the following commands:
* `make db-build`
    * This command builds the db 
* `make db-start`
    * This command starts a postgresql-server docker container, creates a new role and products table.
* `make start`
    * Starts the webserver