# performance_test_go

Go stdlib JWT benchmark backend.

## Endpoints

- `POST /api/login`
- `GET /api/userinfo`

`/api/userinfo` accepts both:

- `GoToken: <token>`
- `Authorization: Bearer <token>`

## Shared DB dependency

This backend uses the shared PostgreSQL 17 instance from `performance_test_db`.

Start DB first:

```bash
cd ../performance_test_db
cp .env.example .env
docker compose up -d
```

## Run backend with Docker

```bash
docker compose up --build
```

## Run backend locally

```bash
go run .
```

Default DB connection for local run:

- host: `localhost`
- port: `5432`
- database: `testdb`
- user: `testuser`
- password: `testpassword`
