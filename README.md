# RSS Aggregator

An RSS aggregator server written in Go. It periodically checks for new posts based on the list of URLs in the database. Users can subscribe to RSS feeds and receive a personalized feed.

## Features

- User management
- Feed management
- Subscription to feeds
- Personalized post feeds

## Endpoints

- `GET /users` - Retrieve user information *(authentication required)*
- `POST /users` - Create a new user
- `GET /feeds` - Retrieve all feeds
- `POST /feeds` - Create a new feed *(authentication required)*
- `GET /feed_follows` - Retrieve feed follows *(authentication required)*
- `POST /feed_follows` - Follow a feed *(authentication required)*
- `DELETE /feed_follows/{id}` - Unfollow a feed *(authentication required)*
- `GET /posts` - Retrieve posts for the authenticated user *(authentication required)*

## Setup

### Prerequisites

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Goose](https://github.com/pressly/goose) - a database migration tool *(optional)*
- [sqlc](https://github.com/sqlc-dev/sqlc) - generates type-safe code from SQL *(optional)*

### Environmental Variables

Create a .env file with the following content:

```env
# Port on which the server will be listening
PORT=8080

# Database connection string (make sure to replace username and password)
DB_URL=postgres://YOUR_USERNAME:YOUR_PASSWORD@localhost:5432/rss-aggregator?sslmode=disable
```

### Database

This project uses PostgreSQL.

Create a new database called **rss-aggregator**. Below are the steps to set up the database:

#### Using Migration Tool

> [!NOTE]
> This section can be ignored if you wish to [set up your database manually](#using-queries) or using different tool.

[Goose](https://github.com/pressly/goose) is used for managing database migrations. Install Goose globally:

```shell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Navigate to the [`/sql/schema/`](/sql/schema/) directory and apply migrations:

```shell
goose postgres postgres://YOUR_USERNAME:YOUR_PASSWORD@localhost:5432/rss-aggregator up
```

#### Using Queries

Alternatively, you can manually run the following SQL queries to set up the database:

```psql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (encode(sha256(random()::text::bytea), 'hex'))
);

CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  last_fetched_at TIMESTAMP,
  UNIQUE(user_id, feed_id)
);

CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  url TEXT UNIQUE NOT NULL,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
```

## Start the Application

Navigate to the root directory of the project and run the following commands to build and start the application:

- Build the application:

```shell
go build
```

- Start the application:

```shell
./rss-aggregator
```
