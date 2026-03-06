# DoIt API

A RESTful task planner API built with Go, Gin, PostgreSQL, and sqlc.

## Tech Stack

- **Go** — primary language
- **Gin** — HTTP framework
- **PostgreSQL** — database
- **sqlc** — type-safe SQL query generation
- **goose** — database migrations
- **squirrel** — dynamic query builder for filtering
- **JWT** — authentication

## Project Structure

```
doit-api/
├── cmd/
│   └── api/
│       ├── main.go        # entry point
│       ├── app.go         # dependency wiring
│       └── routes.go      # route registration
│
├── internal/
│   ├── config/            # app configuration
│   ├── database/          # postgres connection
│   ├── sqlc/              # sqlc generated code
│   ├── migrations/        # goose SQL migrations
│   ├── repository/        # data access layer
│   ├── service/           # business logic
│   ├── handler/           # HTTP handlers
│   ├── middleware/        # auth middleware
│   ├── dto/               # request/response types
│   └── utils/             # shared helpers
│
├── queries/               # sqlc SQL queries
├── static/                # uploaded files (avatars)
├── sqlc.yaml
├── Makefile
├── .air.toml
├── .env
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- sqlc
- goose

### Installation

```bash
git clone https://github.com/wasilisk/doit-api
cd doit-api
go mod download
```

### Environment Variables

Create a `.env` file in the project root:

```env
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=doit-db
DB_SSLMODE=disable

JWT_SECRET=your_secret_key
```

### Database Setup

Create the database:

```bash
psql -U postgres -c 'CREATE DATABASE "doit-db";'
psql -U postgres -c 'GRANT ALL PRIVILEGES ON DATABASE "doit-db" TO your_user;'
psql -d doit-db -U postgres -c 'GRANT ALL ON SCHEMA public TO your_user;'
```

Run migrations:

```bash
make migrate-up
```

### Run

Always run from the project root:

```bash
make run
```

## API Reference

All timestamps are ISO 8601 format: `2026-03-06T17:41:24Z`

### Auth

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/auth/register` | Register a new user |
| POST | `/auth/login` | Login and receive JWT token |

**Register**
```json
POST /auth/register
{
  "email": "user@example.com",
  "password": "secret123",
  "full_name": "John Doe"
}
```

**Login**
```json
POST /auth/login
{
  "email": "user@example.com",
  "password": "secret123"
}
```

Response:
```json
{ "token": "eyJhbGci..." }
```

### Authentication

All `/api/*` endpoints require a Bearer token:

```
Authorization: Bearer <token>
```

### Profile

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/profile` | Get current user profile |
| PUT | `/api/profile` | Update profile (multipart/form-data) |

**Update Profile** — `multipart/form-data`
```
full_name: John Doe
avatar: <file>
```

### Tags

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tags` | Get all tags |
| POST | `/api/tags` | Create a tag |
| PUT | `/api/tags/:id` | Update a tag |
| DELETE | `/api/tags/:id` | Delete a tag |

**Create/Update Tag**
```json
{
  "name": "Work",
  "color": "#FF5733"
}
```

### Tasks

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/tasks` | Get tasks (with optional filters) |
| POST | `/api/tasks` | Create a task |
| GET | `/api/tasks/:id` | Get task by ID |
| PATCH | `/api/tasks/:id` | Update a task |
| DELETE | `/api/tasks/:id` | Soft delete a task |
| POST | `/api/tasks/:id/restore` | Restore a deleted task |

**Create Task**
```json
{
  "name": "Write API docs",
  "description": "Complete OpenAPI specs",
  "date": "2026-03-06T00:00:00Z",
  "time_start": "2026-03-06T10:30:00Z",
  "time_end": "2026-03-06T11:00:00Z",
  "is_completed": false,
  "is_favourite": false,
  "tag_ids": ["uuid-1", "uuid-2"]
}
```

**Filter Tasks**
```
GET /api/tasks?date=2026-03-06T00:00:00Z&is_completed=false&tag=uuid-1
```

| Query Param | Type | Description |
|-------------|------|-------------|
| `date` | string | Filter by day |
| `is_completed` | bool | Filter by completion status |
| `is_deleted` | bool | Show soft-deleted tasks |
| `tag` | string | Filter by tag ID |

## Development

### Regenerate sqlc

```bash
sqlc generate
```

### Run development mode

This will start the API with hot-reload (using Air) so your changes are applied automatically

```bash
make dev
```