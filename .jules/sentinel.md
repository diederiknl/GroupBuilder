## 2024-10-24 - Hardcoded Secrets & Broken Build
**Vulnerability:** Hardcoded JWT secret keys in `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** The project was in a state where critical secrets were committed directly to source code, likely for convenience during initial development. Additionally, the codebase was broken (references to undefined types and functions), masking the security issue and complicating verification.
**Prevention:** Use environment variables for all secrets from day one. Ensure CI/CD pipelines run tests and builds to catch broken code early. Use pre-commit hooks or linters (e.g., `gosec`) to detect hardcoded credentials.
