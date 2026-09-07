# Go MongoDB REST API

A small layered REST API built with Go, Fiber, and MongoDB. It models users and
hobbies, provides CRUD operations for both resources, and demonstrates a MongoDB
aggregation that returns users associated with a hobby.

## What it demonstrates

- Clear routing, handler, service, model, and database layers
- MongoDB CRUD operations with the official Go driver
- Resource lookup by MongoDB ObjectID
- Aggregation with `$lookup`
- Environment-based configuration and explicit connection lifecycle

## Stack

- Go 1.18+
- Fiber
- MongoDB Go Driver
- MongoDB 7 for local development

## Run locally

1. Copy the example environment file:

   ```bash
   cp .env.example .env
   ```

2. Start MongoDB:

   ```bash
   docker compose up -d
   ```

3. Download dependencies and start the API:

   ```bash
   go mod download
   go run .
   ```

The API listens on `http://localhost:3000` by default.

## Configuration

| Variable | Default | Purpose |
|---|---|---|
| `APP_PORT` | `3000` | HTTP server port |
| `MONGODB_URI` | `mongodb://localhost:27017` | MongoDB connection URI |
| `MONGODB_DATABASE` | `go-mongodb` | Database name |
| `MONGODB_USERS_COLLECTION` | `users` | User collection |
| `MONGODB_HOBBIES_COLLECTION` | `hobbies` | Hobby collection |

Do not commit credentials in a populated `.env` file. Use authentication and a
secure connection URI outside local development.

## API routes

| Method | Route | Description |
|---|---|---|
| `GET` | `/users` | List users |
| `GET` | `/users/:id` | Get a user by ObjectID |
| `POST` | `/users` | Create a user |
| `PUT` | `/users/:id` | Update a user |
| `DELETE` | `/users/:id` | Delete a user |
| `GET` | `/hobbies` | List hobbies |
| `GET` | `/hobbies/:id` | Get a hobby by ObjectID |
| `POST` | `/hobbies` | Create a hobby |
| `PUT` | `/hobbies/:id` | Update a hobby |
| `DELETE` | `/hobbies/:id` | Delete a hobby |
| `GET` | `/users/hobbies/:hobby` | List users associated with a hobby |

Example requests:

```bash
curl -X POST http://localhost:3000/hobbies \
  -H 'Content-Type: application/json' \
  -d '{"name":"cycling"}'

curl -X POST http://localhost:3000/users \
  -H 'Content-Type: application/json' \
  -d '{"name":"Ada","surname":"Lovelace","hobby":"cycling","age":36}'

curl http://localhost:3000/users/hobbies/cycling
```

## Project structure

```text
.
├── database/   # MongoDB configuration and connection lifecycle
├── handlers/   # HTTP request and response handling
├── models/     # MongoDB documents
├── router/     # Route registration
├── service/    # Database operations and aggregation
└── viewmodel/  # Aggregation result models
```

## Verification

```bash
go test ./...
go build ./...
```

## Scope

This repository is an educational backend example. Authentication,
authorization, advanced validation, pagination, and production observability are
outside its current scope.
