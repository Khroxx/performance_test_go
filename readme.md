# performance_test_go

Go stdlib JWT benchmark backend.

## Role in project

- Handles `/api/login` and `/api/userinfo` for backend tag `go`.
- Uses shared PostgreSQL 17 data from `performance_test_db`.

## Endpoints

- `POST /api/login`
- `GET /api/userinfo`

`/api/userinfo` accepts both:

- `GoToken: <token>`
- `Authorization: Bearer <token>`

## Repository map and clone URLs

- `performance_test_angular` (frontend): `https://github.com/Khroxx/performance_test_angular.git`
- `performance_test_go` (Go backend): `https://github.com/Khroxx/performance_test_go.git`
- `performance_test_django` (Django Ninja backend): `https://github.com/Khroxx/performance_test_django.git`
- `performance_test_java` (Spring Boot backend): `https://github.com/Khroxx/performance_test_java.git`
- `performance_test_db` (shared PostgreSQL 17): `https://github.com/Khroxx/performance_test_db.git`

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
