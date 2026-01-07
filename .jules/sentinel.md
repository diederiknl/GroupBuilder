## 2026-01-07 - [CRITICAL] Hardcoded JWT Secret
**Vulnerability:** The JWT signing key was hardcoded as `"your_secret_key"` in `internal/auth/auth.go`.
**Learning:** Hardcoded secrets in source code allow anyone with access to the repo (or a leaked version) to generate valid authentication tokens, effectively bypassing all authentication.
**Prevention:** Use environment variables (e.g., `JWT_SECRET`) to inject secrets at runtime. Fallback to a clear "dev-only" default with a warning log for local development to maintain developer experience without compromising production security.
