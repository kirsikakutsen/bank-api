# Bank API with Go

Simple banking REST API built with Go, Echo, GORM, PostgreSQL, and JWT authentication.

## Features

- User signup and login
- JWT-protected endpoints
- Account balance retrieval
- Deposit and withdrawal operations
- Transfers between users
- Transaction history with pagination and filtering
- Username lookup

## Tech Stack

- Go
- Echo v5
- GORM
- PostgreSQL
- JWT (HS256)

## Project Structure

```text
cmd/api/main.go                  # app entrypoint
internal/routes/routes.go        # route registration
internal/handlers/               # HTTP handlers
internal/auth/                   # JWT sign/verify + middleware
internal/database/database.go    # PostgreSQL connection + migrations
internal/models/                 # database models
internal/dto/                    # request/response DTOs
```

## Environment Variables

Create a .env file in the project root.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=gobank
JWT_SECRET=secret
```

## Getting Started

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Start PostgreSQL

Use a local PostgreSQL instance and make sure the database in DB_NAME exists.

### 3. Run the API

```bash
go run cmd/api/main.go
```

Server starts on:

```text
http://localhost:8080
```

## API Endpoints

Base path: /api

### Public

- GET /ping
- POST /auth/signup
- POST /auth/login

### Protected (Authorization: Bearer <token>)

- POST /auth/logout
- GET /account
- POST /transfer
- GET /transfers
- GET /accounts/lookup?username=<username>
- POST /withdraw
- POST /deposit

## Request Examples

### Signup

```bash
{
    "username": "account Name",
    "password": "password"
}
```

### Login

```bash
{
    "username": "test5",
    "password": "pass1234"
}
```

### Deposit

```bash
{
    "amountCents": 5000
}
```

### Withdraw

```bash
{
    "amountCents": 300
}
```

### Transfer

```bash
{
    "toUsername": "username",
    "amountCents": 300
}
```

### Transactions

```bash
{{BaseURL}}/api/transfers?page=1&limit=5&type=transfer
{{BaseURL}}/api/transfers?page=1&limit=5
{{BaseURL}}/api/transfers?id=18
```

## Notes

- Amount values use cents (int64).
- Transaction type filter values are: transfer, deposit, withdrawal.
- Database tables are auto-migrated on startup.


