## 2024-05-18 - Hardcoded Secrets
**Vulnerability:** A hardcoded JWT secret ("your_secret_key" and "neinneinnein") was used in the `internal/auth` package and `cmd/api/main.go`.
**Learning:** Hardcoded secrets present a critical vulnerability because anyone with access to the codebase can impersonate any user or forge authentication tokens. In this project, this was exposed directly in `auth.go` as a global variable.
**Prevention:** Always use environment variables (e.g., `os.Getenv("JWT_SECRET")`) or a secure secrets manager to fetch keys at runtime, and ensure that they are not committed to version control. Let the application fail securely if the secret is missing.
