# Beer Production API

  

This is a Go project for managing beer production. Follow the steps below to get started.

  

## Prerequisites

  

1.  **Go**: You need to have Go installed on your machine. You can download and install it from the [official website](https://golang.org/dl/)

2.  **PostgreSQL**: You should have PostgreSQL installed and running on your machine. You can download and install PostgreSQL from the [official website](https://www.postgresql.org/download/linux/)

 

## Installation

  

Clone the repository:


```bash
	git  clone  https://github.com/MatheusMorais2/beer-production-api.git
```

  

Change to the project directory:

  

```bash
	cd  beer-production-api
```

  

Install project dependencies using Go modules:
```bash
	go  mod  tidy
```

Create a PostgreSQL database for the project:

 
```bash
	createdb  beer_production
```

Migrate the PostgreSQL database:

```bash
	make  migration_up
```

  

Setup Documentation API

```bash
	swag  init
```

## Running the project
Start the server with:
```bash
	go run cmd/main.go
```

## Documentation
Run the project with
```bash
	go run cmd/main.go
```

Then, access:
	http://localhost:8080/docs/index.html
