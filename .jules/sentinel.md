## 2024-05-22 - [Hardcoded JWT Secrets]
**Vulnerability:** Found hardcoded JWT signing keys ("your_secret_key" and "neinneinnein") directly in the source code (`internal/auth/auth.go` and `cmd/api/main.go`).
**Learning:** Hardcoded secrets are a critical risk because they allow anyone with code access (or binary analysis) to forge authentication tokens. They are often left during early development and forgotten.
**Prevention:** Use environment variables (e.g., `JWT_SECRET`) to inject secrets at runtime. Fail securely or warn loudly if secrets are missing. Never commit secrets to version control.
