## 2024-05-20 - Remove Hardcoded Secrets
**Vulnerability:** Hardcoded JWT secret keys were found in `internal/auth/auth.go` and `cmd/api/main.go`.
**Learning:** Hardcoded keys are a critical vulnerability and pose a severe security risk if the codebase is exposed. Instead of hardcoding keys, they should be initialized dynamically via environment variables (e.g., `JWT_SECRET`) to ensure keys can be managed securely across different environments.
**Prevention:** Always use environment variables for sensitive configuration options like secret keys and use checks early in application startup to fail fast if required secrets are missing.
