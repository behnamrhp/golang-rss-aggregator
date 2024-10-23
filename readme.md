# Go Rss Aggregator

## Table of Contents
- [Overview](#overview)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
  - [Migrate postgres](#migrate-postgres)
  - [env file](#env-file)
  - [Start project](#start-project)

## Overview
Go RSS Aggregator is a small training project designed to allow users to manage their own RSS feeds. Users can create feeds, follow them, and receive the latest posts scraped repeatedly in each minuets, from these sources.

The application is powered by REST APIs, which provide the following functionality for managing users and their RSS feeds.

- User Management
  - Create a new user.
  - Authenticate users with middleware.
  - Fetch user details.

- Feed Management
  - Create and manage RSS feeds.
  - Users can follow/unfollow feeds.
  - Scrape and store posts from feeds.

- Post Management
  - Retrieve all posts from feeds a user follows.

## Technologies
- Language: [Golang](https://go.dev/)
- Database: [Postgress](https://postgresql.org/)
- Migration: [Goose](https://github.com/pressly/goose)
- Sql compiler: [Sqlc](https://github.com/sqlc-dev/sqlc)
- Router: [Chi](https://github.com/go-chi/chi)

## Getting Started

### Migrate postgres
```
goose postgres posgtres://[username]:[password]@[db address]/[db name] up
```

### Env file
Uncommenct .env.example and configure it

### Start project
```
make build
```

### Run tidy with vendor
```
make refresh
```

### Clear built file
```
make clean
```