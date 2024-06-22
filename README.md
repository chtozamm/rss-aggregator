# RSS Aggregator

RSS Aggregator written in Go. Features database, authentication.

## Setup

### Environmental Variables

```env
PORT=   # Port on which server will be listening
DB_URL= # Database connection string (i.e. postgres://user:password@localhost:5432/rss-aggregator?sslmode=disable)
```

### Database

This project uses PostgreSQL.

Create new database called **rss-aggregator**. Below you'll find two ways of populating the database: using [migration tool](#using-migration-tool) and [manual](#manual-table-initialization).

### Using Migration Tool

> [!NOTE]
> This section could be ignored if you wish to [setup your database manually](#manual-table-initialization) or using different tool.

In this project, [Goose](https://github.com/pressly/goose) is used for managing database migrations. To install it globally, run the command:

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Migrations are located at [`/sql/schema/`](/sql/schema/). To apply migrations, navigate to that location and run the following command:

```shell
goose postgres postgres://YOUR_USERNAME:YOUR_PASSWORD@localhost:5432/rss-aggregator up
```

### Manual Table Initialization

Run the following queries to populate database with data:

```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL
);
```

### Start the Application

Make sure to change directory to the root of the project and run the following commands to build and start your application:

- Linux/MacOS:

```go
go build && ./rss-aggregator
```

- Windows:

```go
go build && ./rss-aggregator.exe
```

### Note on sqlc

[slqc](https://github.com/sqlc-dev/sqlc) is used to generate type-safe Go code from SQL.
