# GroupBuilder

A web application for building student groups.

## Getting Started

### Prerequisites

- Go 1.16+
- SQLite3

### Configuration

You must set the `JWT_SECRET` environment variable to run the application securely.

1.  Copy `.env.example` to `.env` (if using a tool to load .env files) or just export the variable.
    ```bash
    export JWT_SECRET=your_super_secret_key
    ```

### Running the Application

```bash
go run cmd/api/main.go
```

### Running Tests

```bash
go test ./...
```
