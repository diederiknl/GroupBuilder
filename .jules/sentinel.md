## 2025-03-22 - Hardcoded JWT Secret Key in Codebase
**Vulnerability:** A hardcoded `jwtKey` variable `[]byte("your_secret_key")` was found in `internal/auth/auth.go`, and another copy in `cmd/api/main.go`. This allows attackers to forge valid JWT tokens if they obtain the source code.
**Learning:** Hardcoding secrets directly in the codebase is a critical vulnerability that bypasses secure secret management and makes key rotation nearly impossible without redeploying the application.
**Prevention:** Always retrieve secrets, such as JWT signing keys, from a secure environment variables (e.g., `os.Getenv("JWT_SECRET")`) or a dedicated secret management service during runtime. Return secure, descriptive errors if secrets are missing.
