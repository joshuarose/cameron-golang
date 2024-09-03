This is a skeleton Golang project with no dependencies yet.

One of the great things about Go is how few dependencies are needed to build really functional apis.

The app currently has a main function that calls a sub-function add in main.go and some unit test cases in main_test.go

Go 1.23 or greater required.

I added a docker-compose that only starts postgres.
```
docker compose up
```

For building a rest API you'll need an HTTP Muxer/Server

For HTTP:
- you can use standard library by importing "net/http"
- you can use a third party as well. Gin/Echo/Fiber are all popular. https://github.com/jmalloc/echo-server
- You'll need a Postgres Driver connecting to the database, pgx is recommended: https://github.com/jackc/pgx
- You can just manually run migrations and schema updates and use pgx to run queries but I like sql-c for these interactions: https://sqlc.dev/, some folks really like GORM
