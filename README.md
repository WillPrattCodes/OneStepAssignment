
## Prerequisites

- [SQLite3](https://sqlite.org/download.html) (Only if you want to interact with the database)

## Backend Setup

- cd backend
- go mod tidy
- create .env file
- go run cmd/main.go

Server will run on http://localhost:8080

## Frontend Setup

- cd frontend
- npm install
- create .env file
- npm run dev

## Some reasons for certain decisions

* Although I know your company uses MySQL, I decided to use SQLite instead. This is because I am much more familiar with Postgres. I was originally going to use pgx but chose SQLite for ease of use. While I know that pgx is frequently used and you certainly are familiar with it, I thought it would be more streamlined and easy for you to test given the simplicity and small scale of this application. I also did not want to add unneeded complexity until it was needed.
* Decided against using a SQL query builder like Squirrel because of the small scale of the application and the simplicity of the SQL queries. I didn’t want to overcomplicate it.









