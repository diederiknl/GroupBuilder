## 2026-01-18 - Hardcoded JWT Secret
**Vulnerability:** Found a hardcoded JWT secret `"your_secret_key"` in `internal/auth/auth.go`.
**Learning:** Secrets should never be hardcoded in source code as they can be easily extracted.
**Prevention:** Use environment variables to inject secrets at runtime, and provide safe defaults (with warnings) or fail fast for development.
